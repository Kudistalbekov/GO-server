package models

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

type User struct {
	Name         string      `json:"name" db:"name"`
	Surname      zero.String `json:"surname" db:"surname"`
	Username     zero.String `json:"username" db:"username"`
	ID           zero.String `json:"id" db:"id"`
	Email        string      `json:"email" db:"email"`
	RegisterDate *time.Time  `db:"registerdate"`
	Password     string      `json:"password" db:"password"`
}
