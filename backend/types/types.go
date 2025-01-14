package types

type Properties struct {
	Name  string   `json:"name"`
	Units []string `json:"units"`
}

type DeletePropertyPayload struct {
	Name string `json:"name"`
}

type PropertiesRepository interface {
	GetPropertyByName(name string) (*Properties, error)
	CreateProperty(property Properties) error
	GetAllProperties() ([]*Properties, error)
	DeleteProperty(name string) error
	GetPropertiesByNumberOfBedrooms(numberBedrooms string) error
}

type MockProperties struct {
	Name  string   `json:"name"`
	Units []string `json:"units"`
}
