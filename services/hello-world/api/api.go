package api

import (
	"encoding/json"
	//"errors"
	"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/models"

	//auth "github.com/Stenny-io/stenny/services/auth/functions"
	//"github.com/Stenny-io/stenny/services/common/conf"
	. "github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/functions"
	. "github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/middlewares"
	//"github.com/Stenny-io/stenny/services/common/models"
	"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/response"
	//delete "github.com/Stenny-io/stenny/services/user_delete/functions"
	//profile "github.com/Stenny-io/stenny/services/user_profile/functions"
	"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"log"
	"net/http"
	//"strconv"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", SetContentTypeJson(RootEndpoint)).Methods(http.MethodGet)
	router.HandleFunc("/get-version", SetContentTypeJson(GetVersion)).Methods(http.MethodGet)
	router.HandleFunc("/post-hello", SetContentTypeJson(ParseJsonBody(LoginEndpointPOST))).Methods(http.MethodPost)
	router.NotFoundHandler = SetContentTypeJson(NotFoundEndpoint)
	router.MethodNotAllowedHandler = SetContentTypeJson(MethodNotAllowed)
	return router
}


// HTTP HANDLERS
func RootEndpoint(w http.ResponseWriter, _ *http.Request) {
	response.ServeSuccessResponse(w, "Hey there 👋", nil)
}

func GetVersion(w http.ResponseWriter, _ *http.Request) {
	VERSION := "v1.0"
	response.ServeSuccessResponse(w, "Service is running version: "+VERSION, nil)
}

func PostHelloEndpointPort(w http.ResponseWriter, r *http.Request){
	var posthelloreq models.PostHelloReqBody
	body := GetDataFromContext(r, "body")
	json.Unmarshal([]byte(body), &posthelloreq)

	response.ServeSuccessResponse(w, "Hello, I will mail you at ", posthelloreq.Name, posthelloreq.Email)


}

func LoginEndpointPOST(w http.ResponseWriter, r *http.Request) {
	var (
		cred models.UserCredentials
		u models.User
		medium string
	)
	body := GetDataFromContext(r, "body")
	json.Unmarshal([]byte(body), &cred)

	// Checking if blank data provided
	if auth.IsLoginBodyEmpty(w, &cred){
		return
	} else if !IsEmpty(cred.Email) {
		// Going with email login flow
		if !ValidEmail(w, cred.Email){
			return
		}

		err := auth.FindOrCreateOneByEmail(w, &u, cred.Email)
		if !IsNil(err) {
			return
		}
		medium = "email"
	} else {
		if !ValidMobileNumber(w, cred.CountryCode, cred.Mobile){
			return
		}

		err := auth.FindOrCreateOneByMobile(w, &u, cred.CountryCode, cred.Mobile)
		if !IsNil(err) {
			return
		}
		medium = "mobile"
	}
	if !IsUserActive(w, u) {
		return
	}
	ok, d := auth.DoAuthStateGenerationWorkForLogin(w, r, u, medium)
	if !ok{
		return
	}
	response.ServeSuccessResponse(w, "Verification Code Sent", d)
}

func MethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	response.ServeFailureResponse(w, http.StatusMethodNotAllowed, "The place you want to visit is over there, but you are entering through wrong gate! 😛", nil)
}

func NotFoundEndpoint(w http.ResponseWriter, _ *http.Request) {
	response.ServeFailureResponse(w, http.StatusNotFound, "The place you want to visit is not available! 😅", nil)
}
