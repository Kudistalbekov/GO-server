package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/server/conn"
	"projects/server/models"
	"time"
)

//Response for client
type Response struct {
	Success bool        `json:"succes"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

//RegPost - registers with the Post requests
func RegPost(w http.ResponseWriter, r *http.Request) {
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
	op := "handlers/ReqGet"

}
