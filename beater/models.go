package beater

type serviceMap map[string][]string

type ServiceHealth struct {
	Node Node
	Service Service
	Checks []Check
}

type TaggedAddresses struct {
	Lan string
	Wan string
}

type Node struct {
	Node string
	Address string
	TaggedAddresses TaggedAddresses
}

type Service struct {
	ID string
	Service string
	Tags []string
	Address string
	Port int
}

type Check struct {
	Node string
	CheckID string
	Name string
	Status string
	Notes string
	Output string
	ServiceID string
	ServiceName string
}