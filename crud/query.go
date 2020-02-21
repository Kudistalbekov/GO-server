package crud

import (
	"projects/server/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//RegisterUser date is going to register user,
//If found error ,error returned to function caller
//If no error ,returned nil
func RegisterUser(user *models.User, db *sqlx.DB) error {
	op := "crud/RegisterUser"
	//key := "myverystrongpasswordo32bitlength"
	//c := encrypt([]byte(key), user.Password)
	_, err := db.Query("insert into users (name,email,registerdate,password) values($1,$2,$3,$4)", user.Name, user.Email, time.Now(), user.Password)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
