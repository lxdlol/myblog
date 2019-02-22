package models

import (
	"time"
)

type User struct {
	Model
	Name   string `gorm:"unique_index" json:"name"`
	Email  string `gorm:"unique_index" json:"email"`
	Pwd    string `json:"-"`
	Avatar string `json:"avatar"`
	Role   int    `json:"role"`
}

func QuerryByEmail(email string) (user User, err error) {

	e := db.Where("email = ?", email).Take(&user).Error
	return user, e
}
func QuerryByName(name string) (user User, err error) {
	e := db.Where("name = ?", name).Take(&user).Error
	return user, e
}

func SaveUser(user User) error {
	return db.Save(&user).Error
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
