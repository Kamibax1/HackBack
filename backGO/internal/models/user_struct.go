package models

type UserLogin struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
}

type UserRegister struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
}

type UserDelete struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserUpdate struct {
	UserName        string `json:"user_name"`
	UserPassword    string `json:"user_password"`
	NewUserName     string `json:"new_user_name"`
	NewUserPassword string `json:"new_user_password"`
}

type UserProfile struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	UserName     string `gorm:"uniqueIndex;not null"`
	UserEmail    string `gorm:"uniqueIndex;not null"`
	UserRole     string
	PasswordHash string `gorm:"not null"`
	Salt         string `gorm:"not null"`
}
