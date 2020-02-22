package handlers
import (
	"net/http"
"projects/server/models"
)

type Myhandle(http.ResponseWriter, *http.Request)(string,int,error)

func handleError(a Myhandle)http.HandleFunc{
	resp := &models.Response{
			

	}
	return (w http.ResponseWriter,r *http.Request){
	err := a(w,r)
	if err:=nil{

	}
	}
}