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
func RegPost(w http.ResponseWriter, r *http.Request) {
	op := "handlers/RegPost"
	if "application/json" == r.Header.Get("Content-Type") { // will determine what format request is
		user := &models.User{} ///creating empty user

		w.Header().Add("Content-type", "application/json") //setting to json in order to help client understand

		resp := &models.Response{ //creating response for ok
			Success: true,
			Error:   "no error",
			Data:    nil,
		}

		err := json.NewDecoder(r.Body).Decode(user) //r.Body is what client sent and (a is user)
		if err != nil {                             //Unmarshal from r.Body to a(User)
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalf("%s %v", op, err)
		}

		db := conn.DBconnect()
		defer db.Close()

		err = crud.RegisterUser(user, db) //Registering user
		if err != nil {
			if err == crud.UserExistAlready {
				resp.Success = true
				resp.Error = "user exist already"
				json.NewEncoder(w).Encode(resp)
			}
			return
		}

		json.NewEncoder(w).Encode(resp) //converting our responsse into json and writing to responsewriter
	}
}

//RegGet going to send data to user
//using gmail
func RegGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "json/application") //we are going to send to the client json
	user := &models.User{}
	respuser := &models.ResponseUser{}
	email := r.FormValue("email") //127.0.0.1:8080/user/?email=youremail@gmail.com

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	db := conn.DBconnect()
	defer db.Close()

	//row := db.QueryRowx("select * from users where email=$1", email)
	//err := row.StructScan(user)
	err := crud.GetUserByEmail(email, db).StructScan(user)

	respuser.Email = user.Email
	respuser.ID = user.ID
	respuser.Name = user.Name
	respuser.Surname = user.Surname
	respuser.Username = user.Username

	resp := &models.Response{
		Success: true,
		Error:   "no error",
		Data:    respuser,
	}
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			resp.Error = "user not exist"
			resp.Success = true
			resp.Data = nil
		} else {
			log.Fatalf("StructScan %v", err)
			resp.Error = "error"
			resp.Success = false
			resp.Data = nil
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(resp)

}

//ReqDelete deletes the user by email
func ReqDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "json/application")

	op := "ReqDelete"
	resp := &models.Response{
		Success: true,
		Error:   "no error",
		Data:    nil,
	}

	email := r.FormValue("email")

	db := conn.DBconnect()
	defer db.Close()

	err := crud.DeleteUser(email, db)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			resp.Success = true
			resp.Error = "user not exist"
			w.WriteHeader(http.StatusBadRequest)
		}
		resp.Success = false
		resp.Error = "error"
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("%s %v", op, err)
	}

	json.NewEncoder(w).Encode(resp)
}

//ReqPut updtates the data
func ReqPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "json/application")
	op := "handlers/ReqPut"
	email := r.FormValue("email")
	resp := &models.Response{
		Success: true,
		Error:   "no error",
		Data:    nil,
	}
	if r.Header.Get("content-type") != "json/application" {

		newuser := &models.User{}
		dbuser := &models.User{}

		db := conn.DBconnect()
		defer db.Close()

		err := crud.GetUserByEmail(email, db).StructScan(dbuser)

		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				resp.Error = "user not exist"
				json.NewEncoder(w).Encode(resp)
				fmt.Println("%s %v", op, resp.Error)
				return
			}
			resp.Error = "error"
			resp.Success = false
			json.NewEncoder(w).Encode(resp)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("%s %v", op, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(newuser)
		if err != nil {
			log.Fatal(err)
		}

		change(dbuser, newuser)

		//deleting old data
		err = crud.DeleteUser(email, db)
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("%s %v", op, err)
		}

		//inserting new
		err = crud.RegisterUser(dbuser, db)
		if err != nil {
			if err == crud.UserExistAlready {
				resp.Success = true
				resp.Error = "user exist already"
				json.NewEncoder(w).Encode(resp)
			}
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(resp)

	} else {
		resp.Success = false
		resp.Error = "%s the type is not json"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
	}
}
