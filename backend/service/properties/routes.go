package properties

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mik-dmi/types"
)

type Handler struct {
	repository types.PropertiesRepository
}

func NewHandler(repository types.PropertiesRepository) *Handler {
	return &Handler{repository: repository}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/properties", h.handleCreateProperty).Methods("Post")
	router.HandleFunc("/properties/{propertiesID}", h.handleGetProperty).Methods(http.MethodGet)
}
func (h *Handler) handleGetProperties(w http.ResponseWriter, r *http.Request) {
	log.Println("Get the Properties")
}
func (h *Handler) handleGetProperty(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		log.Panic("missing request body")
		return
	}
	var payload types.Properties
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("error decoding the body form the request")
		return
	}
}

func (h *Handler) handleCreateProperty(w http.ResponseWriter, r *http.Request) {

}
