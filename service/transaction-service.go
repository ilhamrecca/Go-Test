package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//BookService is a ....
type TransactionService interface {
	Insert(b dto.TransactionCreateDTO) entity.Transaction
	// Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Transaction)
	All() []entity.Transaction
	FindByID(transactionID uint64) entity.Transaction
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	bookRepository        repository.BookRepository
}

//NewBookService .....
func NewTransactionService(transRepo repository.TransactionRepository, bookRepo repository.BookRepository) TransactionService {
	return &transactionService{
		transactionRepository: transRepo,
		bookRepository:        bookRepo,
	}
}

func (service *transactionService) Insert(b dto.TransactionCreateDTO) entity.Transaction {
	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	var book = service.bookRepository.FindBookByID(transaction.BookID)
	book.Stock += transaction.Amount
	service.bookRepository.UpdateBook(book)
	res := service.transactionRepository.InsertTransaction(transaction)
	return res
}

func (service *transactionService) Delete(b entity.Transaction) {
	service.transactionRepository.DeleteTransaction(b)
}

func (service *transactionService) All() []entity.Transaction {
	return service.transactionRepository.AllTransaction()
}

func (service *transactionService) FindByID(transactionID uint64) entity.Transaction {
	return service.transactionRepository.FindTransactionByID(transactionID)
}

// func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
// 	book := entity.Book{}
// 	err := smapping.FillStruct(&book, smapping.MapFields(&b))
// 	if err != nil {
// 		log.Fatalf("Failed map %v: ", err)
// 	}
// 	res := service.bookRepository.UpdateBook(book)
// 	return res
// }

// func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
// 	b := service.bookRepository.FindBookByID(bookID)
// 	id := fmt.Sprintf("%v", b.UserID)
// 	return userID == id
// }
