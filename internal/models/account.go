package models

type User struct {
	BaseModel
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RobotAccount struct {
	BaseModel
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
