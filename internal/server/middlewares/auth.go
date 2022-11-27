package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ericzty/eve/internal/controllers"
	"github.com/ericzty/eve/internal/controllers/authentication"
	"github.com/rs/zerolog/log"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		if !strings.HasPrefix(reqToken, "v1") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		id, err := authentication.VerifyToken(ctx, reqToken)
		if err != nil {
			if errors.Is(err, authentication.TokenErr) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid token"))
				return
			}
			if errors.Is(err, authentication.TokenExpiredErr) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Token expired"))
				return
			}
			if errors.Is(err, authentication.ServerTokenErr) {
				log.Error().Err(err).Msg("Error parsing server token")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
				return
			}
		}
		isAdmin, err := controllers.IsAdmin(ctx, id)
		if err != nil {
			log.Error().Err(err).Msg("Error fetching isAdmin from db")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		fmt.Println(id)
		fmt.Println(isAdmin)
	})
}
