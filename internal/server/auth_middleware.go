package server

import (
	"context"
	"harrapa/internal/database"
	"harrapa/internal/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type UserId int
type UserRole int

const (
	UserIdKey   UserId   = 0
	UserRoleKey UserRole = 1
)

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.ExtractJwtFromHeader(r.Header)
		if err != nil {
			log.Println(err)
			utils.SendResponse(w, http.StatusForbidden, utils.DefaultUnauthenticatedResponse)
			return
		}

		if err = token.Valid(); err != nil {
			log.Println(err)
			utils.SendResponse(w, http.StatusForbidden, utils.DefaultUnauthenticatedResponse)
			return
		}

		id, err := uuid.Parse(token.UserId)

		if err != nil {
			log.Println(err)
			utils.SendResponse(w, http.StatusForbidden, utils.DefaultUnauthenticatedResponse)
			return
		}

		// TODO(sam): implement this
		// _, err = s.db.ValidateAuthUser(r.Context(), database.ValidateAuthUserParams{
		// 	ID:   id,
		// 	Type: database.Usertype(token.UserType),
		// 	LoginHash: sql.NullString{
		// 		Valid:  true,
		// 		String: token.ApiKey,
		// 	},
		// })

		ctx := context.WithValue(r.Context(), UserIdKey, id)
		ctx = context.WithValue(ctx, UserRoleKey, database.UserRole(token.UserRole))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func (s *Server) AllowedUserWithRole(role database.UserRole, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value(UserRoleKey).(database.UserRole)

		if userRole != role {
			utils.SendResponse(w, http.StatusUnauthorized, utils.DefaultUnauthenticatedResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func (s *Server) AllowedOnlyAdmin(next http.Handler) http.Handler {
	return s.AllowedUserWithRole(database.UserRoleAdmin, next)
}
