package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/server/models"
)

type Response struct {
	Success bool        `json:"succes"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

//Reg - registers with the Post requests
func Reg(w http.ResponseWriter, r *http.Request) {
	op := "handlers/Registration"
	fmt.Println("Method :%v ", r.Method, r.RemoteAddr)
	if "application/json" == r.Header.Get("Content-Type") { //will determine what format request is
		user := &models.User{} ///creating empty user
		//ask bejan
		//setting to json in order to help client understand
		w.Header().Add("Content-type", "application/json")
		//r.Body is what client sent and (a is user)
		//Unmarshal from r.Body to a(User)
		err := json.NewDecoder(r.Body).Decode(user)
		//creating response for ok
		resp := &Response{
			Success: true,
			Error:   "no error",
			Data:    user, //nil
		}

		if err != nil {
			log.Fatalf("%s %v", op, err)
			//changing for error,and success false
			resp.Error = err.Error()
			resp.Success = false
			resp.Data = nil
			w.WriteHeader(http.StatusBadRequest)
		}
		//converting our responsse into json
		//writing to responsewriter
		json.NewEncoder(w).Encode(resp)
		db := conn.DBconnect()
		defer db.Close()
	}
	fmt.Println("connected")
}
