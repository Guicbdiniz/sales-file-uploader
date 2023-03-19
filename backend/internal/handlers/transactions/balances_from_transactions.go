package transactions

import "guicbdiniz/hubla/backend/internal/models"

// ProcessBalancesFromTransactions updates clients balances from received
// transactions.
//
// If a transaction with a new seller is received, a new client is
// inserted in the models.
func ProcessBalancesFromTransactions(m *models.Models, transactions []models.Transaction) error {
	var clientsToUpdate []models.Client
	for _, transaction := range transactions {
		clientsToUpdate = updateClientsForTransaction(clientsToUpdate, transaction)
	}

	for _, client := range clientsToUpdate {
		err := m.Clients.AddOrUpdateClient(client)
		if err != nil {
			return err
		}
	}
	return nil
}

// updateClientsForTransaction updates the clients array using
// the data from the current analyzed transaction.
func updateClientsForTransaction(clientsToUpdate []models.Client, transaction models.Transaction) []models.Client {
	balance, clientIndex := getOldClientInfo(clientsToUpdate, transaction)
	balance, isProducer := getClientInfoFromTransactionType(transaction.Type, transaction.Value, balance)
	if clientIndex != -1 {
		clientsToUpdate[clientIndex].Balance = balance
	} else {
		clientsToUpdate = append(clientsToUpdate, models.Client{
			Balance:    balance,
			IsProducer: isProducer,
			Name:       transaction.Seller,
		})
	}
	return clientsToUpdate
}

// getOldClientInfo checks if the passed transaction uses
// data from an already used client.
func getOldClientInfo(clientsToUpdate []models.Client, transaction models.Transaction) (float64, int) {
	var balance float64
	var clientIndex int = -1
	for i, client := range clientsToUpdate {
		if client.Name == transaction.Seller {
			balance = client.Balance
			clientIndex = i
			break
		}
	}
	return balance, clientIndex
}

// getClientInfoFromTransactionType analyzes the transaction data
// and returns the updated clients balance and if it is a producer.
func getClientInfoFromTransactionType(transactionType string, transactionValue float64, balance float64) (float64, bool) {
	var isProducer bool = false
	switch transactionType {
	case "ProducerSale":
		balance += transactionValue
		isProducer = true
	case "AffiliatedSale":
		balance += transactionValue
	case "PaidCommission":
		balance -= transactionValue
	case "ReceivedCommission":
		balance += transactionValue
		isProducer = true
	}
	return balance, isProducer
}
