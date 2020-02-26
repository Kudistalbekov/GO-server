package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/server/models"
)

//Myhandle does every work and returns the result
type Myhandle func(http.ResponseWriter, *http.Request) (string, int, interface{}, error)

//HandleError takes my func and handles his response
func HandleError(a Myhandle) http.HandlerFunc {
	op := "HandleError"
	return func(w http.ResponseWriter, r *http.Request) {
		kind, status, data, err := a(w, r)
		resp := &models.Response{ //creating response for ok
			Success: true,
			Error:   "no error",
			Data:    data,
		}
		w.Header().Add("content-type", "application/json")
		if err != nil {
			if kind == "user" {
				resp.Error = err.Error()
			} else if kind == "system" {
				resp.Success = false
				log.Printf("%s:%v->", op, err.Error())
			}
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(resp)
	}
}
