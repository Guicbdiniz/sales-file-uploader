package models

import "database/sql"

// Client is a seller from one or more transactions
type Client struct {
	Id         int     `json:"id,omitempty"`
	Name       string  `json:"name"`
	IsProducer bool    `json:"isProducer"`
	Balance    float64 `json:"balance"`
}

// ClientsModel is an interface with only the necessary
// data methods related to Clients.
type ClientsModel interface {
	InitTable() error
	GetAllClients() ([]Client, error)
	AddNewClient(client Client) error
	UpdateClient(client Client) error
	AddOrUpdateClient(client Client) error
}

// ClientsPostgresModel implements the transactions models methods
// using SQL queries to retrieve and insert clients to a
// Postgres database.
type ClientsPostgresModel struct {
	db *sql.DB
}

func (model *ClientsPostgresModel) InitTable() error {
	query := `CREATE TABLE IF NOT EXISTS clients ( 
				id SERIAL NOT NULL, 
				name VARCHAR(255) NOT NULL UNIQUE, 
				is_producer BOOLEAN DEFAULT FALSE, 
				balance DOUBLE PRECISION DEFAULT 0.0, 
				PRIMARY KEY(ID)
	);`
	stmt, err := model.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (model *ClientsPostgresModel) GetAllClients() ([]Client, error) {
	query := "SELECT id, name, is_producer, balance FROM clients;"

	var clients []Client

	rows, err := model.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.Id, &client.Name, &client.IsProducer, &client.Balance)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (model *ClientsPostgresModel) AddNewClient(client Client) error {
	query := "INSERT INTO clients (name, is_producer, balance) VALUES ($1, $2, $3)"
	stmt, err := model.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Name, client.IsProducer, client.Balance)
	if err != nil {
		return err
	}

	return nil
}

func (model *ClientsPostgresModel) UpdateClient(client Client) error {
	query := "UPDATE clients SET name = $1, is_producer = $2, balance = $3 WHERE id = $4"
	stmt, err := model.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Name, client.IsProducer, client.Balance, client.Id)
	if err != nil {
		return err
	}

	return nil
}

func (model *ClientsPostgresModel) AddOrUpdateClient(client Client) error {
	query := "INSERT INTO clients (name, is_producer, balance) VALUES ($1, $2, $3) ON CONFLICT " +
		"(name) DO UPDATE SET balance = excluded.balance;"
	stmt, err := model.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Name, client.IsProducer, client.Balance)
	if err != nil {
		return err
	}

	return nil
}

// CreateClientsPostgresModel returns a ready to use
// ClientsModel that fetches and inserts data using
// the database instance passed.
func CreateClientsPostgresModel(db *sql.DB) *ClientsPostgresModel {
	return &ClientsPostgresModel{
		db: db,
	}
}

// ClientsTestModel implements the clients model's methods
// using a simple slice.
//
// It must be used in tests only.
type ClientsTestModel struct {
	clients []Client
}

func (model *ClientsTestModel) InitTable() error {
	return nil
}

func (model *ClientsTestModel) GetAllClients() ([]Client, error) {
	return model.clients, nil
}

func (model *ClientsTestModel) AddNewClient(client Client) error {
	model.clients = append(model.clients, client)
	return nil
}

func (model *ClientsTestModel) UpdateClient(client Client) error {
	var clientIndex int = -1
	for i, c := range model.clients {
		if c.Id == client.Id {
			clientIndex = i
			break
		}
	}

	if clientIndex != -1 {
		model.clients[clientIndex] = client
	}
	return nil
}

func (model *ClientsTestModel) AddOrUpdateClient(client Client) error {
	var clientIndex int = -1
	for i, c := range model.clients {
		if c.Name == client.Name {
			clientIndex = i
			break
		}
	}
	if clientIndex != -1 {
		model.clients[clientIndex] = client
	} else {
		model.clients = append(model.clients, client)
	}
	return nil
}

// CreateClientsTestModel returns a fresh in memory clients model
// for testing.
func CreateClientsTestModel() *ClientsTestModel {
	var clients []Client
	return &ClientsTestModel{
		clients: clients,
	}
}
