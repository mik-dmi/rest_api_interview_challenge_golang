package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mik-dmi/service/properties"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	propertyRepository := properties.NewRepository(s.db)
	propertiesHandler := properties.NewHandler(propertyRepository)
	propertiesHandler.RegisterRoutes(router)

	corsOptions := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	cors.Default()

	handlerWithCors := corsOptions.Handler(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, handlerWithCors)
}
