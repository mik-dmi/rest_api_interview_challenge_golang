package properties

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mik-dmi/types"
	"github.com/mik-dmi/utils"
)

type Handler struct {
	repository types.PropertiesRepository
}

func NewHandler(repository types.PropertiesRepository) *Handler {
	return &Handler{repository: repository}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/properties", h.handleCreateProperty).Methods("POST")
	router.HandleFunc("/properties/{propertyID}", h.handleGetProperty).Methods("GET")
	router.HandleFunc("/properties", h.handleGetAllProperties).Methods("GET")
	router.HandleFunc("/properties", h.handleDeleteProperty).Methods("DELETE")
}
func (h *Handler) handleGetAllProperties(w http.ResponseWriter, r *http.Request) {
	properties, err := h.repository.GetAllProperties()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(""))
		return
	}
	utils.WriteJSON(w, http.StatusOK, properties)
}
func (h *Handler) handleGetProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	propertyID := vars["propertyID"]
	properties, err := h.repository.GetPropertyByName(propertyID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("property with the name %s does not exists", propertyID))
		return
	}
	utils.WriteJSON(w, http.StatusOK, properties)
}

func (h *Handler) handleCreateProperty(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		log.Panic("missing request body")
		return
	}

	var newProperty types.Properties
	err := json.NewDecoder(r.Body).Decode(&newProperty)
	//log.Println(newProperty)
	if err != nil {
		log.Println("error decoding the body from the request")
		return
	}

	_, err = h.repository.GetPropertyByName(newProperty.Name)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("property with the name %s already exists", newProperty.Name))
		return
	}

	err = h.repository.CreateProperty(newProperty)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

}

func (h *Handler) handleDeleteProperty(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing body request"))
		return
	}

	var deleteProperty types.DeletePropertyPayload
	err := json.NewDecoder(r.Body).Decode(&deleteProperty)
	//log.Println(newProperty)
	if err != nil {
		log.Println("error decoding the body from the request")
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request body"))

		return
	}

	_, err = h.repository.GetPropertyByName(deleteProperty.Name)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("property with the name %s does not exists", deleteProperty.Name))
		return
	}

	err = h.repository.DeleteProperty(deleteProperty.Name)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

}
