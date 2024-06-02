package middleware

import (
	"errors"
	"goapi/api"
	"goapi/internal/tools"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		// write an error function inside api/api.go
		var err error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return //exiting the error check
		}

		var database *tools.DatabaseInterface // Need to create as internal/tools/database.go
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
		// Authorization -> next.ServeHTTP() -> GetCoinBalance
	})
}
