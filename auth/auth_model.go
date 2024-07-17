package auth

type AccessTokenResponse struct {
	Status      uint   `json:"status" example:"200"`
	Message     string `json:"message" example:"OK"`
	AccessToken string `json:"accessToken" example:"Token"`
}

type User struct {
	ID       uint  `json:"id" gorm:"primary_key;auto_increment;column:ID;type:int4"`
	Username string `json:"username" gorm:"column:Username" binding:"required"`
	Email    string `json:"email" gorm:"column:Email" binding:"required"`
	Password string `json:"password" gorm:"column:Password" binding:"required"`
	IsActive bool `json:"isActive" gorm:"column:IsActive"`
}

type UserData struct {
	Username 	string	`json:"username" gorm:"column:Username"`
	AccessToken string 	`json:"accessToken" gorm:"column:Password"`
}

type UserLogin struct {
	Username string `json:"username" gorm:"column:Username" binding:"required"`
	Password string `json:"password" gorm:"column:Password" binding:"required"`
}