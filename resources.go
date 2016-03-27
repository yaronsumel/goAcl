package acl

import (
	"errors"
)

type resources struct {
	resourceMap map[string]*resource
}

// returns new resources instance
func newResources()resources{
	return resources{
		resourceMap: make(map[string]*resource),
	}
}

// adds new resource to resourceMap if its not exist yet
func (r *resources)appendResource(resourceName string)*resource {
	resource,err := r.getResource(resourceName)
	if err != nil{
		resource = newResource(resourceName)
		r.resourceMap[resourceName] = resource
	}
	return resource
}

// get resource from resourceMap
func (r *resources)getResource(resourceName string) (*resource, error) {
	if resource, ok := r.resourceMap[resourceName]; ok {
		return resource, nil
	}
	return nil, errors.New("no such resource")
}
