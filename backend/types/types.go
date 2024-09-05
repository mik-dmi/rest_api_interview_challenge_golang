package types

type Properties struct {
	Name string   `json:"name"`
	Unit []string `json:"unit"`
}

type PropertiesRepository interface {
	GetPropertyByNamme(name string) (*Properties, error)
	CreateProperty(name string, units []string) error
	GetAllProperties() error
	DeleteProperty(name string) error
	GetPropertiesByNumberOfBedrooms(numberBedrooms string) error
}
