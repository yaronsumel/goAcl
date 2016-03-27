package acl

type role struct {
	name      string
	resources resources
}

// returns new role pointer
func newRole(roleName string) *role {
	return &role{
		name:roleName,
		resources:newResources(),
	}
}

// basically copies all resources from newRole to role
func (r *role)Inherits(newRole *role) *role {
	if r.name != newRole.name {
		for _, resource := range (newRole.resources.resourceMap) {
			r.AllowResource(resource.name)
		}
	}
	return r
}

// if resource is already exist change to active resource
// else will create new resource to this role
func (r *role)AllowResource(resourceName string) *role {
	resource, err := r.resources.getResource(resourceName)
	if err == nil {
		resource.activate()
	}else {
		r.resources.appendResource(resourceName)
	}
	return r
}

// get the resource and mark it as not active
func (r *role)DenyResource(resourceName string) {
	resource, err := r.resources.getResource(resourceName)
	if err == nil {
		resource.deactivate()
	}
}

// checks if resourceName is allowed for this role
func (r *role)IsAllowed(resourceName string) bool {
	resource, err := r.resources.getResource(resourceName)
	if err != nil {
		return false
	}
	return resource.isActive()
}
