package dto

type TransactionCreateDTO struct {
	UserID uint64 `json:"user_id,"  form:"user_id" binding:"required"`
	BookID uint64 `json:"book_id,"  form:"book_id" binding:"required"`
	Amount int64  `json:"amount,"  form:"amount" binding:"required"`
}
