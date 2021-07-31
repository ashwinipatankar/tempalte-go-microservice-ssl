package response

import (
	"encoding/json"
	"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/models"
	"log"
	"net/http"
)

func ServeFailureResponse(w http.ResponseWriter, responseCode int, message string, err error) {
	if responseCode == http.StatusInternalServerError {
		log.Println(err.Error())
	}
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(models.NewResponse("failure", message, nil))

}

func ServeSuccessResponse(w http.ResponseWriter, message string, data map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.NewResponse("success", message, data))

}

func AddHeadersToResponse(w http.ResponseWriter, name string, value string){
	w.Header().Set(name, value)
}