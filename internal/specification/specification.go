package specification

type Database struct {
	Name DatabaseName `json:"name" yaml:"name"`
}

type File struct {
	// Source when defined indicates that this configuration file shoudl be considered a child of this Source
	Source *Source `json:"Source" yaml:"Source"`
	// Name defines the name of this system specification
	Name string `json:"name" yaml:"name"`
	// Services defines the services within this system
	Services       []Service       `json:"services" yaml:"services"`
	Infrastructure *Infrastructure `json:"infrastructure" yaml:"infrastructure"`
}

type Lifecycle struct {
	// Type defines the type of this lifecycle
	Type LifecycleType `json:'type" yaml:"type"`
}

type Service struct {
	// Name indicates the name of this service
	Name string `json:"name" yaml:"name"`
	// Description indicates why this Service exists
	Description string `json:"description" yaml:"description"`
	// Services indicates other Services that this Service has knowledge of (connects to)
	Services []Service `json:"services" yaml:"services"`
	// Ref indicates that this service has already been defined and to refer to it
	Ref string `json:"ref" yaml:"ref"`
	// Type defines the type of the service
	Type Type `json:"type" yaml:"type"`
	// Database defines the properties of the service if it's a database instance (Type should equal TypeDatabase)
	Database *Database `json:"database" yaml:"database"`
	// Target defines where the Service will be deployed to
	Target *Target `json:"target" yaml:"target"`
	// Interface defines the Service's interface properties
	Interface *Interface `json:"interface" yaml:"interface"`
	// Lifecycle defines the Service's lifecycle properties
	Lifecycle *Lifecycle `json:"lifecycle" yaml:"lifecycle"`
}

type Source struct {
	URL  string `json:"url" yaml:"url"`
	Name string `json:"Source" yaml:"Source"`
}

type Infrastructure struct {
	// AzCount defines the number of availability zones this system should span across
	AzCount uint8 `json:"azCount" yaml:"azCount"`
}

type Interface struct {
	Type InterfaceType `json:"type" yaml:"type"`
	// Specification defines the Service's specification
	Specification *Specification `json:"specification" yaml:"specification"`
}

type Endpoint struct {
	Description string `json:"description" yaml:"description"`
	// Method defines the method at which the endpoint will respond
	Method string `json:"method" yaml:"method"`
	// Path defines the path at which the endpoint will respond
	Path       string `json:"path" yaml:"path"`
	Parameters map[string]string
}

type Specification []Endpoint

type Target struct {
	Name TargetName `json:"name" yaml:"name"`
}
