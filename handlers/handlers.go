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
	//below : will determine what format request is
	if "application/json" == r.Header.Get("Content-Type") {
		user := &models.User{} ///creating empty user
		//setting to json in order to help client understand
		w.Header().Add("Content-type", "application/json")
		//creating response for ok
		resp := &models.Response{
			Success: true,
			Error:   "no error",
			Data:    nil,
		}
		//r.Body is what client sent and (a is user)
		//Unmarshal from r.Body to a(User)
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			//changing for not ok
			resp.Error = err.Error()
			resp.Success = false
			resp.Data = nil
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalf("%s %v", op, err)
		}
		//converting our responsse into json
		//writing to responsewriter
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("%s %v", op, err)
		}
		db := conn.DBconnect()
		defer db.Close()
		//Registering user
		err = crud.RegisterUser(user, db)
		if err != nil {
			log.Fatalf("%s %v", op, err)
		}
	}
	fmt.Println("connected")
}

//RegGet going to send data to user
//using gmail
func RegGet(w http.ResponseWriter, r *http.Request) {
	//we are going to send to the client json
	w.Header().Add("content-type", "json/application")
	user := &models.User{}
	respuser := &models.ResponseUser{}
	//127.0.0.1:8080/user/?email=youremail@gmail.com
	email := r.FormValue("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	db := conn.DBconnect()
	defer db.Close()

	row := db.QueryRowx("select * from users where email=$1", email)
	err := row.StructScan(user)
	//
	respuser.Email = user.Email
	respuser.ID = user.ID
	respuser.Name = user.Name
	respuser.Surname = user.Surname
	respuser.Username = user.Username
	//
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
	op := "ReqDelete"
	resp := &models.Response{
		Success: true,
		Error:   "no error",
		Data:    nil,
	}
	email := r.FormValue("email")
	db := conn.DBconnect()
	defer db.Close()
	_, err := db.Query("DELETE from users where email=$1", email)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			resp.Success = false
			resp.Error = "user not exist"
			w.WriteHeader(http.StatusBadRequest)
		}
		resp.Success = false
		resp.Error = "error"
		log.Fatalf("%s %v", op, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//ReqPut updtates the data
func ReqPut(w http.ResponseWriter, r *http.Request) {
	op := "handlers/ReqPut"
	email := r.FormValue("email")
	if r.Header.Get("content-type") != "json/application" {
		newuser := &models.User{}
		dbuser := &models.User{}
		db := conn.DBconnect()
		defer db.Close()
		row := db.QueryRowx("select * from users where email=$1", email)
		err := row.StructScan(dbuser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("%s %v", op, err)
		}
		json.NewDecoder(r.Body).Decode(newuser)
		change(dbuser, newuser)
		//deleting old data
		_, err = db.Query("DELETE from users where email=$1", email)
		if err != nil {
			log.Fatalf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		//inserting new
		_, err = db.Query("insert into users (id,name,email,surname,username,password,registerdate) values($1,$2,$3,$4,$5,$6,$7)", dbuser.ID, dbuser.Name, dbuser.Email, dbuser.Surname, dbuser.Username, dbuser.Password, dbuser.RegisterDate)
		if err != nil {
			log.Fatalf("Query %v", err)
		}
		//fmt.Println(dbuser)
	} else {
		log.Fatalf("%s the type is not json", op)
		w.WriteHeader(http.StatusBadRequest)
	}
}
