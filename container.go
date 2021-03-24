package gs

import (
	"reflect"
)

const initMethodName = "GoService"

// Container struct
type Container struct {
	services map[interface{}]interface{}
}

// Get returns existing or creates new instance of specified type
func (c *Container) Get(i interface{}) interface{} {
	service := c.services[reflect.TypeOf(i)]
	if service == nil {
		service = c.Create(i)
		c.services[reflect.TypeOf(i)] = service
	}
	return service
}

// Set value of instance of specified type
func (c *Container) Set(i interface{}, v interface{}) {
	c.services[reflect.TypeOf(i)] = v
}

// Create creates new instance of specified type
func (c *Container) Create(i interface{}) interface{} {
	instanceType := reflect.TypeOf(i).Elem()
	svc := reflect.New(instanceType)
	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		fieldType := field.Type
		if fieldType.Kind() != reflect.Ptr {
			continue
		}
		fieldTypeElem := fieldType.Elem()
		if fieldTypeElem.Kind() != reflect.Struct {
			continue
		}
		f := svc.Elem().Field(i)
		if !f.CanSet() {
			continue
		}
		val := c.Get(reflect.Zero(fieldType).Interface())
		f.Set(reflect.ValueOf(val))
	}
	method, ok := svc.Type().MethodByName(initMethodName)
	if ok {
		argumentsCount := method.Type.NumIn()
		arguments := make([]reflect.Value, argumentsCount-1)
		for i := 1; i < argumentsCount; i++ {
			s := c.Get(reflect.New(method.Type.In(i).Elem()).Interface())
			arguments[i-1] = reflect.ValueOf(s)
		}
		svc.Method(method.Index).Call(arguments)
	}
	service := svc.Interface()
	return service
}
