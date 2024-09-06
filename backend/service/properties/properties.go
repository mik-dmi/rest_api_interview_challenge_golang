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
	/*log.Printf("Inserting property: Name = %s, Units = %v\n", property.Name, property.Units)
	if len(property.Units) == 0 {
		return fmt.Errorf("units array is empty or nil")
	}*/

	_, err := s.db.Exec("INSERT INTO properties (name, units) VALUES($1, $2)", property.Name, pq.Array(property.Units))
	if err != nil {
		return err
	}
	return nil
}
func (s *Repository) GetAllProperties() ([]*types.Properties, error) {
	rows, err := s.db.Query("Select * From properties")
	if err != nil {
		return nil, err
	}

	properties := make([]*types.Properties, 0)
	for rows.Next() {
		property := new(types.Properties)
		err := rows.Scan(
			&property.Name,
			pq.Array(&property.Units), // Handle array with pq.Array
		)
		if err != nil {
			return nil, err
		}

		properties = append(properties, property)
	}

	return properties, nil
}
func (s *Repository) DeleteProperty(property string) error {
	_, err := s.db.Exec("DELETE FROM properties WHERE name = $1", property)
	if err != nil {
		return err
	}
	return nil
}
func (s *Repository) GetPropertiesByNumberOfBedrooms(numberBedrooms string) error {
	return nil
}
