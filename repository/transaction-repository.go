package repository

import (
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

//BookRepository is a ....
type TransactionRepository interface {
	InsertTransaction(b entity.Transaction) entity.Transaction
	UpdateTransaction(b entity.Transaction) entity.Transaction
	DeleteTransaction(b entity.Transaction)
	AllTransaction() []entity.Transaction
	FindTransactionByID(transactionID uint64) entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}

func (db *transactionConnection) InsertTransaction(b entity.Transaction) entity.Transaction {
	db.connection.Save(&b)
	// db.connection.Preload("User").Preload("Book").Find(&b)
	db.connection.Preload("User").Joins("Book").Find(&b)

	return b
}

func (db *transactionConnection) UpdateTransaction(b entity.Transaction) entity.Transaction {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *transactionConnection) DeleteTransaction(b entity.Transaction) {
	db.connection.Delete(&b)
}

func (db *transactionConnection) AllTransaction() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("User").Find(&transactions)
	return transactions
}

func (db *transactionConnection) FindTransactionByID(transactionID uint64) entity.Transaction {
	var transaction entity.Transaction
	db.connection.Preload("User").Find(&transaction, transactionID)
	return transaction
}
