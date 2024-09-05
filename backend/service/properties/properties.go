package properties

import (
	"database/sql"
	"fmt"

	"github.com/mik-dmi/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) GetPropertyByNamme(name string) (*types.Properties, error) {
	rows, err := s.db.Query("Select * From properties WHERE name= ?", name)
	if err != nil {
		return nil, err
	}
	property := new(types.Properties)
	for rows.Next() {
		err = rows.Scan(
			&property.Name,
			&property.Unit,
		)
		if err != nil {
			return nil, err
		}
		if property.Name == "" {
			return nil, fmt.Errorf("user not found")
		}
	}
	return property, nil
}

func (s *Repository) CreateProperty(name string, units []string) error {
	return nil
}
func (s *Repository) GetAllProperties() error {
	return nil
}
func (s *Repository) DeleteProperty(name string) error {
	return nil
}
func (s *Repository) GetPropertiesByNumberOfBedrooms(numberBedrooms string) error {
	return nil
}
