package gs

var global = Container{services: map[interface{}]interface{}{}}

// Get returns existing or creates new instance of specified type
func Get(i interface{}) interface{} {
	return global.Get(i)
}

// Set value of instance of specified type
func Set(i interface{}, v interface{}) {
	global.Set(i, v)
}

// Create creates new instance of specified type
func Create(i interface{}) interface{} {
	return global.Create(i)
}
