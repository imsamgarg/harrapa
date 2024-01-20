package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"harrapa/internal/database"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

var validate = validator.New()

type Server struct {
	port  int
	db    *database.Queries
	sqldb *sql.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, sqldb := database.NewConn()

	NewServer := &Server{
		port:  port,
		db:    db,
		sqldb: sqldb,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
