package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"` // Assuming UserID is the primary key of User
	Username string `json:"username"`
	Password string `json:"password"`
}

type Bookmarks struct {
	gorm.Model
	User  User   `gorm:"foreignKey:ID" json:"user"`
	Title string `json:"title"`
}
