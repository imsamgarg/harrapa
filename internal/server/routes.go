package server

import (
	"encoding/json"
	"fmt"
	"harrapa/internal/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	authRouter := chi.NewRouter()

	v1Router.Get("/", s.HelloWorldHandler)

	authRouter.Post("/login", s.LoginHandler)
	authRouter.Post("/register", s.RegisterHandler)
	// authRouter.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
	// 	s.AuthMiddleware(http.HandlerFunc(s.LogoutHandler)).ServeHTTP(w, r)
	// })

	v1Router.Mount("/auth", authRouter)

	r.Mount("/v1", v1Router)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.SendResponse(w, 404, utils.NewErrorResponse(fmt.Sprintf("Path %v not found", r.URL.Path)))
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
