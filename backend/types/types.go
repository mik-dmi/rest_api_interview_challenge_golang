package types

type Properties struct {
	Name  string   `json:"name"`
	Units []string `json:"unit"`
}

type PropertiesRepository interface {
	GetPropertyByName(name string) (*Properties, error)
	CreateProperty(property Properties) error
	GetAllProperties() ([]Properties, error)
	DeleteProperty(name string) error
	GetPropertiesByNumberOfBedrooms(numberBedrooms string) error
}
