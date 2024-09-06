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

	cors.Default()

	handlerWithCors := cors.Default().Handler(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, handlerWithCors)
}
