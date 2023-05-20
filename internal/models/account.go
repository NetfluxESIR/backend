package models

type Account struct {
	BaseModel
	Email          string `json:"email" gorm:"unique"`
	HashedPassword string `json:"hashed_password"`
	Role           string `json:"role"`
}
