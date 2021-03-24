package gs

// New creates instance of services container
func New() *Container {
	return &Container{
		services: map[interface{}]interface{}{},
	}
}
