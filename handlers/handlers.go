package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/server/conn"
	"projects/server/models"

	"projects/server/crud"

	"github.com/pkg/errors"
)

//RegPost - registers with the Post requests
func RegPost(w http.ResponseWriter, r *http.Request) (string, int, interface{}, error) {
	//op := "handlers/RegPost"
	if "application/json" == r.Header.Get("Content-Type") { // will determine what format request is
		user := &models.User{} ///creating empty user

		w.Header().Add("Content-type", "application/json") //setting to json in order to help client understand

		err := json.NewDecoder(r.Body).Decode(user) //r.Body is what client sent
		if err != nil {                             //Unmarshal from r.Body to a(User)
			return "user", http.StatusBadRequest, nil, err
		}

		db := conn.DBconnect()
		defer db.Close()

		err = crud.RegisterUser(user, db) //Registering user
		if err != nil {
			if err == crud.UserExistAlready {
				return "user", http.StatusBadRequest, nil, errors.New("user already exist")
			}
			return "system", http.StatusInternalServerError, nil, err
		}

		return " ", http.StatusOK, "registered", nil

	}
	return "user", http.StatusBadRequest, r.Header.Get("content-type"), errors.New("format is not json")
}

//RegGet going to send data to user
//using gmail
func RegGet(w http.ResponseWriter, r *http.Request) (string, int, interface{}, error) {

	user := &models.User{}
	respuser := &models.ResponseUser{}
	email := r.FormValue("email") //127.0.0.1:8080/user/?email=youremail@gmail.com

	if email == "" {
		return "user", http.StatusBadRequest, nil, errors.New("email is empty")
	}
	db := conn.DBconnect()
	defer db.Close()

	//row := db.QueryRowx("select * from users where email=$1", email)
	//err := row.StructScan(user)
	err := crud.GetUserByEmail(email, db).StructScan(user)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return "user", http.StatusBadRequest, nil, errors.New("user does not exist")
		}
		return "system", http.StatusInternalServerError, nil, err
	}

	respuser.Email = user.Email
	respuser.ID = user.ID
	respuser.Name = user.Name
	respuser.Surname = user.Surname
	respuser.Username = user.Username

	return "", http.StatusOK, respuser, nil
}

//ReqDelete deletes the user by email
func ReqDelete(w http.ResponseWriter, r *http.Request) (string, int, interface{}, error) {
	w.Header().Add("content-type", "application/json")

	op := "ReqDelete"

	email := r.FormValue("email")

	db := conn.DBconnect()
	defer db.Close()

	err := crud.DeleteUser(email, db)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return "user", http.StatusBadRequest, nil, errors.New("user not exist")
			log.Fatalf("%s %v", op, err)
		}
		return "system", http.StatusInternalServerError, nil, err
		log.Fatalf("%s %v", op, err)
	}

	return "", http.StatusOK, "user deleted", nil
}

//ReqPut updtates the data
func ReqPut(w http.ResponseWriter, r *http.Request) (string, int, interface{}, error) {
	w.Header().Add("content-type", "application/json")
	op := "handlers/ReqPut"
	email := r.FormValue("email")
	if email == "" {
		return "user", http.StatusBadRequest, nil, errors.New("empty email")
	}
	if r.Header.Get("content-type") == "application/json" {

		newuser := &models.User{}
		dbuser := &models.User{}

		db := conn.DBconnect()
		defer db.Close()

		err := crud.GetUserByEmail(email, db).StructScan(dbuser)

		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				return "user", http.StatusBadRequest, nil, errors.New("user not exist")
			}
			fmt.Println("%s %v", op, err)
			return "system", http.StatusInternalServerError, nil, err
		}

		err = json.NewDecoder(r.Body).Decode(newuser)
		if err != nil {
			return "system", http.StatusInternalServerError, nil, err
		}

		change(dbuser, newuser)

		//deleting old data
		err = crud.DeleteUser(email, db)
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				return "user", http.StatusBadRequest, nil, errors.New("user not exist")
			}
			return "system", http.StatusInternalServerError, nil, err
		}

		//inserting new
		err = crud.RegisterUser(dbuser, db)
		if err != nil {
			if err == crud.UserExistAlready {
				return "user", http.StatusBadRequest, nil, errors.New("user already exist")
			}
			return "system", http.StatusInternalServerError, nil, err
		}
		return "", http.StatusOK, "user updated", nil

	}
	return "user", http.StatusBadRequest, r.Header.Get("content-type"), errors.New("request is not json type")
}
