package models

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Models is the api's main and only way to access data.
//
// It contains all the models used in the service.
type Models struct {
	Transactions TransactionsModel
	Clients      ClientsModel
	db           *sql.DB
}

// InitDB creates the necessary database tables
// used by the models.
func (models *Models) InitDB() error {
	err := models.Clients.InitTable()
	if err != nil {
		return err
	}

	err = models.Transactions.InitTable()
	if err != nil {
		return err
	}

	return nil
}

// Clear closes the database connection used by the
// models (if it exists).
func (models *Models) Clear() error {
	if models.db != nil {
		return models.db.Close()
	}
	return nil
}

// CreatePostgresModels returns a reference to a Models
// object which can access data through a Postgres
// connection.
//
// The postgres data source name is taken from
// the environment variable POSTGRES_URL.
//
// Remember to always call Models.Clear after
// usage.
func CreatePostgresModels() (*Models, error) {
	dsn := os.Getenv("POSTGRES_URL")

	if len(dsn) == 0 {
		return nil, errors.New("missing postgres url to connect")
	}

	db, err := sql.Open("postgres", dsn)
	var triesLeft int = 5
	var connected bool = false
	for triesLeft > 0 {
		log.Printf("Attempting to connect to database. %d tries left...", triesLeft)
		err = db.Ping()
		if err == nil {
			connected = true
			break
		}
		log.Printf("Error while connecting to database: %v\n", err)
		log.Println("Reattempting connection in 1 second...")
		triesLeft -= 1
		time.Sleep(time.Second * 1)
	}
	if !connected {
		return nil, err
	}
	log.Println("Connection to database completed.")

	models := Models{
		Clients:      CreateClientsPostgresModel(db),
		Transactions: CreateTransactionsPostgresModel(db),
		db:           db,
	}

	return &models, nil
}

// CreateTestModels returns a reference to a Models
// object which uses in memory objects to mock
// a database connection and its methods.
func CreateTestModels() (*Models, error) {
	models := Models{
		Clients:      CreateClientsTestModel(),
		Transactions: CreateTransactionsTestModel(),
		db:           nil,
	}

	return &models, nil
}
