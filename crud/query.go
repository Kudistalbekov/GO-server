package crud

import (
	"database/sql"
	"projects/server/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var UserDoesNotExist = errors.New("user does not exist")
var UserExistAlready = errors.New("user already exist")

//RegisterUser date is going to register user,
//If found error ,error returned to function caller
//If no error ,returned nil
func RegisterUser(user *models.User, db *sqlx.DB) error {
	op := "crud/RegisterUser"
	if err := UserExist(user, db); err != nil {
		return err
	}
	key := "myverystrongpasswordo32bitlength"
	c := myencrypt([]byte(key), user.Password+" 8gwifi.org")
	_, err := db.Query("insert into users (name,email,registerdate,password,surname,username) values($1,$2,$3,$4,$5,$6)", user.Name, user.Email, time.Now(), c, user.Surname, user.Username)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

//DeleteUser deletes user from database
//true nil returned
//if not returned error
func DeleteUser(email string, db *sqlx.DB) error {
	_, err := db.Query("DELETE from users where email=$1", email)
	if err != nil {
		return err
	}
	return nil
}

//GetUserByEmail returns row with error,
//If no error nil returned
func GetUserByEmail(email string, db *sqlx.DB) *sqlx.Row {
	row := db.QueryRowx("select * from users where email=$1", email)
	return row
}

//UserExist return nil if at most one exist,
//else returns UserExistAlready
func UserExist(user *models.User, db *sqlx.DB) error {
	err := db.QueryRow("select * from users where email=$1", user.Email).Scan()
	if errors.Cause(err) == sql.ErrNoRows {
		return nil
	}
	return UserExistAlready
}
