package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/server/conn"
	"projects/server/models"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3/zero"
)

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

//RegPost - registers with the Post requests
func RegPost(w http.ResponseWriter, r *http.Request) { //w kuda  ,r otkuda
	op := "handlers/RegPost"
	//below : will determine what format request is
	if "application/json" == r.Header.Get("Content-Type") {
		user := &models.User{} ///creating empty user
		//setting to json in order to help client understand
		w.Header().Add("Content-type", "application/json")
		//creating response for ok
		resp := &Response{
			Success: true,
			Error:   "no error",
			Data:    nil,
		}
		//r.Body is what client sent and (a is user)
		//Unmarshal from r.Body to a(User)
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			log.Fatalf("%s %v", op, err)
			//changing for not ok
			resp.Error = err.Error()
			resp.Success = false
			resp.Data = nil
			w.WriteHeader(http.StatusBadRequest)
		}
		//converting our responsse into json
		//writing to responsewriter
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Fatalf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		db := conn.DBconnect()
		defer db.Close()
		//query for inserrt
		db.Query("select count(email) from user where email = 1$", user.Email)

		_, err = db.Query("insert into users (name,email,registerdate,password) values($1,$2,$3,$4)", user.Name, user.Email, time.Now(), user.Password)
		if err != nil {
			log.Fatalf("Query %v", err)
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
	respuser := &ResponseUser{}
	//127.0.0.1:8080/user/?email=youremail@gmail.com
	email := r.FormValue("email")
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
	resp := &Response{
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
