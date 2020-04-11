package middlewares

import (
	"Food-Hub-API/internal/helpers"
	"net/http"
)

func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	err := helpers.TokenValid(r)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "unauthorised to perform this action," +
			" please signup/login")
		return
	}
	next(w, r)
	return
}
