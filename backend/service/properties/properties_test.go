package properties

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mik-dmi/types"
)

func TestPropertiesServiceHandlers(t *testing.T) {
	properties := &mockProperties{}
	handler := NewHandler(properties)
	t.Run("should handle creating a property", func(t *testing.T) {
		payload := types.Properties{
			Name:  "create_property_test",
			Units: []string{"bathroom", "bedroom"},
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/properties", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/properties", handler.handleCreateProperty).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
	t.Run("should handle get properties by ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/properties/test_property", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/properties/{propertyID}", handler.handleGetProperty).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should handle get all the properties", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/properties", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/properties", handler.handleGetAllProperties).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should delete a property using an ID", func(t *testing.T) {
		payload := types.Properties{
			Name: "delete_property_test",
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodDelete, "/properties", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/properties", handler.handleDeleteProperty).Methods(http.MethodDelete)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

type mockProperties struct{}

func (h *mockProperties) CreateProperty(name types.Properties) error {
	return nil
}
func (s *mockProperties) GetPropertyByName(name string) (*types.Properties, error) {

	if name == "create_property_test" {
		return nil, fmt.Errorf("property does not exit")
	}

	return &types.Properties{}, nil
}

func (s *mockProperties) DeleteProperty(property string) error {
	return nil
}

func (s *mockProperties) GetPropertiesByNumberOfBedrooms(numberBedrooms string) error {
	return nil
}
func (s *mockProperties) GetAllProperties() ([]*types.Properties, error) {
	return []*types.Properties{}, nil
}
