package models

import "gopkg.in/guregu/null.v3/zero"

//Response for client
type Response struct {
	Success bool        `json:"succes"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

//ResponseUser for client without date and password
type ResponseUser struct {
	ID       zero.String `json:"id"`
	Name     string      `json:"name"`
	Surname  zero.String `json:"surname"`
	Email    string      `json:"email"`
	Username zero.String `json:"username"`
}
