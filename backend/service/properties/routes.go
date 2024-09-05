package properties

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/properties", h.handleGetProperties).Methods(http.MethodGet)
	router.HandleFunc("/properties/{propertiesID}", h.handleGetProperty).Methods(http.MethodGet)
}
func (h *Handler) handleGetProperties(w http.ResponseWriter, r *http.Request) {
	log.Println("Get the Properties")
}
func (h *Handler) handleGetProperty(w http.ResponseWriter, r *http.Request) {
	log.Println("Get the Property by ID")
}
