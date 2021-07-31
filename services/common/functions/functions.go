package functions

import (
	//"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/conf"

	"encoding/json"
	//"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/conf"

	//	"errors"
	//"github.com/Stenny-io/stenny/services/common/conf"
	//"github.com/Stenny-io/stenny/services/common/database"
	//"github.com/Stenny-io/stenny/services/common/models"
	"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/response"
	//"github.com/Stenny-io/stenny/services/common/validator"
//	"github.com/dgrijalva/jwt-go"
	gcontext "github.com/gorilla/context"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"gopkg.in/gomail.v2"
	"io/ioutil"
	//"log"
	"net/http"
	//"strings"
)

// Below are common functions which will be needed to call across variable services //

// Sets ResponseWriter content type
func SetResponseContentType(w http.ResponseWriter, resType string) {
	w.Header().Set("Content-Type", resType)
}

// Is POST Body Nil
func IsNilReqBody(w http.ResponseWriter, r *http.Request) bool {
	if IsNil(r.Body) {
		response.ServeFailureResponse(w, http.StatusBadRequest, "Invalid POST Body", nil)
		return true
	}
	return false
}

// Returns true if nil
func IsNil(input interface{}) bool {
	return input == nil
}
/*
// Returns true if empty
func IsEmpty(input interface{}) bool {
	if input == "" {
		return true
	} else if input == 0 {
		return true
	}
	return false
}
*/
// Returns true if POST Body is JSON Parsable
func NotJsonReqBody(w http.ResponseWriter, r *http.Request) bool {
	bodyStr, _ := ioutil.ReadAll(r.Body)
	var tmp map[string]interface{}
	if json.Unmarshal(bodyStr, &tmp) != nil {
		response.ServeFailureResponse(w, http.StatusBadRequest, "Invalid POST Body", nil)
		return true
	}
	AddDataToContext(r, "body", string(bodyStr))
	return false
}

// Adds data in context
func AddDataToContext(r *http.Request, name string, value string) {
	gcontext.Set(r, name, value)
}


// Returns data from context
func GetDataFromContext(r *http.Request, name string) string {
	return gcontext.Get(r, name).(string)
}
/*
// Adds user into database and reports if any error occurs
func AddUser(w http.ResponseWriter, user *models.User, isActive bool) error {
	var uF models.UserFactory
	// Setting some default user creation params
	//FeedUserInstanceWithDefaults(user, isActive)
	(*user).PopulateDefaults(isActive)
	// Inserting user into DB
	err := uF.InsertOne(user)
	if !IsNil(err) {
		response.ServeFailureResponse(w, http.StatusInternalServerError, "Application error, please view logs", err)
		return err
	}
	return nil
}

// Validates is emails
func ValidEmail(w http.ResponseWriter, email string) bool {
	if !validator.IsValidEmail(email) || email == "" {
		response.ServeFailureResponse(w, http.StatusBadRequest, "Invalid email address", nil)
		return false
	}
	return true
}

// Validates mobile number
func ValidMobileNumber(w http.ResponseWriter, countryCode string, mobile int) bool {
	if !validator.IsValidMobileNumber(countryCode, mobile) || countryCode == "" || mobile == 0 {
		response.ServeFailureResponse(w, http.StatusBadRequest, "Invalid mobile number or country code", nil)
		return false
	}
	return true
}

// Validates if user is active
func IsUserActive(w http.ResponseWriter, user models.User) bool {
	if !user.IsActive {
		response.ServeFailureResponse(w, http.StatusForbidden, "Your profile has been disabled, please contact us for details. Mail us at: "+conf.CONF("STENNY_HELP_EMAIL"), nil)
	}
	return user.IsActive
}

// Returns http or https schema used in request
func GetSchema(r *http.Request) string {
	// small code function to get if server is running on http or https
	if r.TLS == nil {
		return "http://"
	}
	return "https://"
}

// Returns true if Request has valid Authentication
func IsValidAuth(w http.ResponseWriter, r *http.Request) bool {
	var (
		userId primitive.ObjectID
		auth   string
		claims jwt.MapClaims
		err    error
	)

	auth, err = CheckAuthHeaderPresent(r)
	if !IsNil(err) {
		response.ServeFailureResponse(w, http.StatusForbidden, strings.Title(err.Error()), nil)
		return false
	}

	auth, err = CheckIsBlankAuthHeader(auth)
	if !IsNil(err) {
		response.ServeFailureResponse(w, http.StatusForbidden, strings.Title(err.Error()), nil)
		return false
	}

	claims, err = ParseJWT(auth)
	if !IsNil(err) {
		response.ServeFailureResponse(w, http.StatusForbidden, strings.Title(err.Error()), nil)
		return false
	}

	userId = GetUserIDFromClaims(claims)

	err = models.UserFactory{}.FindOneByUserId(&models.User{}, userId)
	if err!=nil {
		response.ServeFailureResponse(w, http.StatusForbidden, strings.Title("invalid JWT"), nil)
		return false
	}

	AddDataToContext(r, "userId", userId.Hex())
	AddDataToContext(r, "auth", auth)
	return true
}

// Returns user id from JWT Claims
func GetUserIDFromClaims(claims jwt.MapClaims) primitive.ObjectID {
	userId, _ := primitive.ObjectIDFromHex(claims["UserId"].(string))
	return userId
}

// Checks if Auth Header is present of not
func CheckAuthHeaderPresent(r *http.Request) (string, error) {
	if len(r.Header.Values("Authorization")) == 0 {
		return "", errors.New("authorization header missing")
	}
	return r.Header.Values("Authorization")[0], nil
}

// Checks if Auth Header is blank or not
func CheckIsBlankAuthHeader(auth string) (string, error) {
	if auth == "" || strings.TrimSpace(strings.TrimPrefix(auth, "Bearer")) == "" {
		return "", errors.New("blank authorization header")
	}
	auth = strings.TrimPrefix(auth, "Bearer ")
	return auth, nil
}

// Is JWT Parsing and valid?
func ParseJWT(jwtStr string) (jwt.MapClaims, error) {
	expired, err := IsExpiredJWT(jwtStr)
	if !IsNil(err) {
		if !database.ErrorIs(err, "no document") && !database.ErrorIs(err, "token expired") {
			log.Println(err.Error())
			return jwt.MapClaims{}, errors.New("application error, please view logs")
		}
	}
	if expired {
		return jwt.MapClaims{}, err
	}

	tkn, err := ValidateJWT(jwtStr)
	if tkn == nil || !IsNil(err) {
		return jwt.MapClaims{}, errors.New("invalid JWT")
	}

	claims, ok := GiveClaimFromToken(tkn)
	if ok && tkn.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("invalid JWT")
	}
}

// Returns all claims from given JWT
func GiveClaimFromToken(tkn *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := tkn.Claims.(jwt.MapClaims)
	return claims, ok
}

// Validating JWT
func ValidateJWT(jwtStr string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.CONF("JWT_SECRET")), nil
	})
	return tkn, err
}

// Is JWT logged-out one?
func IsExpiredJWT(jwt string) (bool, error) {
	var expJWT models.ExpiredJWTFactory
	if expJWT.Exists(jwt) {
		return true, errors.New("token expired")
	}
	return false, nil
}

// Send mail function
func SendMail(message *gomail.Message, smtpAddress string, smtpPort int, smtpEmail string, smtpPassword string) error {
	d := gomail.NewDialer(smtpAddress, smtpPort, smtpEmail, smtpPassword)
	return d.DialAndSend(message)
}

 */