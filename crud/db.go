package crud

import (
	"projects/server/models"
	"time"

	"firebase.google.com/go/db"
	"github.com/pkg/errors"
)

//RegisterUser date is going to register user
//if found error ,error returned to function caller
//if no error ,returned nil
func RegisterUser(user *models.User) error {
	op := "crud/RegisterUser"
	_, err = db.Query("insert into users (name,email,registerdate,password) values($1,$2,$3,$4)", user.Name, user.Email, time.Now(), user.Password)
	if err != nil {
		return errors.Wrap(wrr, op)
	}
	return nil
}
