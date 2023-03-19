package models

import (
	"database/sql"
	"time"
)

// Transaction is a financial operation which happened in time.
//
// It represents a product sold by a seller by a certain value.
type Transaction struct {
	Id      int       `json:"id,omitempty"`
	Type    string    `json:"type"`
	Date    time.Time `json:"date"`
	Product string    `json:"product"`
	Value   float64   `json:"value"`
	Seller  string    `json:"seller"`
}

// TransactionsModel is an interface with only the necessary
// data methods related to Transactions.
type TransactionsModel interface {
	InitTable() error
	GetAllTransactions() ([]Transaction, error)
	AddNewTransaction(transaction Transaction) error
}

// TransactionsPostgresModel implements the transactions models methods
// using SQL queries to retrieve and insert transactions into
// a Postgres database.
type TransactionsPostgresModel struct {
	db *sql.DB
}

func (model *TransactionsPostgresModel) InitTable() error {
	query := `CREATE TABLE IF NOT EXISTS transactions ( 
		id SERIAL NOT NULL, 
		type VARCHAR(255) NOT NULL, 
		date TIMESTAMP with time zone, 
		product VARCHAR(255) NOT NULL,
		VALUE DOUBLE PRECISION NOT NULL,
		SELLER VARCHAR(255) NOT NULL, 
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

func (model *TransactionsPostgresModel) GetAllTransactions() ([]Transaction, error) {
	query := "SELECT id, type, date, product, value, seller FROM transactions;"

	var transactions []Transaction

	rows, err := model.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Transaction
		err := rows.Scan(&t.Id, &t.Type, &t.Date, &t.Product, &t.Value, &t.Seller)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (model *TransactionsPostgresModel) AddNewTransaction(t Transaction) error {
	query := "INSERT INTO transactions (type, date, product, value, seller) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := model.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Type, t.Date, t.Product, t.Value, t.Seller)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionsPostgresModel returns a ready to use
// TransactionsModel that fetches and inserts data using
// the database instance passed.
func CreateTransactionsPostgresModel(db *sql.DB) *TransactionsPostgresModel {
	return &TransactionsPostgresModel{
		db: db,
	}
}

// TransactionsTestModel implements the transactions model's methods
// using a simple slice.
//
// It must be used in tests only.
type TransactionsTestModel struct {
	transactions []Transaction
}

func (model *TransactionsTestModel) InitTable() error {
	return nil
}

func (model *TransactionsTestModel) GetAllTransactions() ([]Transaction, error) {
	return model.transactions, nil
}

func (model *TransactionsTestModel) AddNewTransaction(t Transaction) error {
	model.transactions = append(model.transactions, t)
	return nil
}

// CreateTransactionsTestModel returns a fresh in memory transactions model
// for testing.
func CreateTransactionsTestModel() *TransactionsTestModel {
	var transactions []Transaction
	return &TransactionsTestModel{
		transactions: transactions,
	}
}
