package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/server/models"
)

type Response struct {
	Success bool
	Error   string
	data    interface{}
}

//Reg - registers with the Post requests
func Reg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method :%v ", r.Method, r.RemoteAddr)
	if "application/json" == r.Header.Get("Content-Type") { //will determine i what format request is
		fmt.Println("json received")
		a := &models.User{}

		resp := &Response{
			Success: true,
			Error:   "no error",
			data:    nil,
		}
		w.Header().Add("Content-type", "application/json")
		err := json.NewDecoder(r.Body).Decode(a)

		if err != nil {
			fmt.Println(err)
			resp.Error = err.Error()
			resp.Success = false
			resp.data = nil
		}

		json.NewEncoder(w).Encode(resp)
		fmt.Println(a)
	}
	fmt.Println("connected")
}
