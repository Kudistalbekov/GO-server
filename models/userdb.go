package models

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

type User struct {
	Name         string      `json:"name"`
	Surname      zero.String `json:"surname"`
	Username     zero.String `json:"username"`
	ID           zero.String `json:"id"`
	Email        string      `json:"email"`
	RegisterDate *time.Time
	Password     string `json:"password"`
}
