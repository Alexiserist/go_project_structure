package models

type User struct {
	ID       uint  `json:"id" gorm:"primary_key;auto_increment;column:ID;type:int4"`
	Username string `json:"username" gorm:"column:Username" binding:"required"`
	Email    string `json:"email" gorm:"column:Email" binding:"required"`
	Password string `json:"-" gorm:"column:Password" binding:"required"`
	IsActive bool `json:"isActive" gorm:"column:IsActive"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" gorm:"column:Password" binding:"required"`
	IsActive bool `json:"isActive" gorm:"column:IsActive"`
}

type UpdateUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" gorm:"column:Password" `
	IsActive bool `json:"isActive" gorm:"column:IsActive"`
}