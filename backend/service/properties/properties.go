package properties

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/mik-dmi/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) GetPropertyByName(name string) (*types.Properties, error) {
	rows, err := s.db.Query("Select * From properties WHERE name=$1", name)
	if err != nil {
		return nil, err
	}
	property := new(types.Properties)
	for rows.Next() {
		err = rows.Scan(
			&property.Name,
			pq.Array(&property.Units),
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

func (s *Repository) CreateProperty(property types.Properties) error {
	_, err := s.db.Exec("INSERT INTO properties (name, units) VALUES($1, $2)", property.Name, pq.Array(property.Units))
	if err != nil {
		return err
	}
	return nil
}
func (s *Repository) GetAllProperties() ([]types.Properties, error) {
	rows, err := s.db.Query("Select * From properties")
	if err != nil {
		return nil, err
	}
	property := new(types.Properties)
	for rows.Next() {
		err = rows.Scan(
			&property.Name,
			pq.Array(&property.Units),
		)
		if err != nil {
			return nil, err
		}
		if property.Name == "" {
			return nil, fmt.Errorf("user not found")
		}
	}
	return nil, nil
}
func (s *Repository) DeleteProperty(name string) error {
	return nil
}
func (s *Repository) GetPropertiesByNumberOfBedrooms(numberBedrooms string) error {
	return nil
}
