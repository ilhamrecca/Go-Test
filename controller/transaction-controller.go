package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
)

//BookController is a ...
type TransactionController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	// Update(context *gin.Context)
	Delete(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

//NewBookController create a new instances of BoookController
func NewTransactionController(tranServ service.TransactionService, jwtServ service.JWTService) TransactionController {
	return &transactionController{
		transactionService: tranServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) All(context *gin.Context) {
	var transactions []entity.Transaction = c.transactionService.All()
	res := helper.BuildResponse(true, "OK", transactions)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var transaction entity.Transaction = c.transactionService.FindByID(id)
	if (transaction == entity.Transaction{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", transaction)
		context.JSON(http.StatusOK, res)
	}
}

func (c *transactionController) Insert(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			transactionCreateDTO.UserID = convertedUserID
		}
		result := c.transactionService.Insert(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
func (c *transactionController) Delete(context *gin.Context) {
	var transaction entity.Transaction
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	transaction.ID = id
	// authHeader := context.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// userID := fmt.Sprintf("%v", claims["user_id"])

	c.transactionService.Delete(transaction)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}

func (c *transactionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

// func (c *bookController) Update(context *gin.Context) {
// 	var bookUpdateDTO dto.BookUpdateDTO
// 	errDTO := context.ShouldBind(&bookUpdateDTO)
// 	if errDTO != nil {
// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	authHeader := context.GetHeader("Authorization")
// 	token, errToken := c.jwtService.ValidateToken(authHeader)
// 	if errToken != nil {
// 		panic(errToken.Error())
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	userID := fmt.Sprintf("%v", claims["user_id"])
// 	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
// 		id, errID := strconv.ParseUint(userID, 10, 64)
// 		if errID == nil {
// 			bookUpdateDTO.UserID = id
// 		}
// 		result := c.bookService.Update(bookUpdateDTO)
// 		response := helper.BuildResponse(true, "OK", result)
// 		context.JSON(http.StatusOK, response)
// 	} else {
// 		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
// 		context.JSON(http.StatusForbidden, response)
// 	}
// }
