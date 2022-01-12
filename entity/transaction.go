package entity

//Book struct represents books table in database
type Transaction struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	BookID uint64 `gorm:"not null" json:"-"`
	Book   Book   `gorm:"foreignkey:BookID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"book"`
	UserID uint64 `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Amount int64  `json:"amount" form:"amount" binding:"required"`
}
