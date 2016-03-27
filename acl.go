package acl

type Acl struct {
	roles     roles
	resources resources
}

// returns new acl instance pointer
func NewAcl() *Acl {
	return &Acl{
		roles:newRoles(),
		resources:newResources(),
	}
}

// add new role to acl instance
func (a *Acl)AddRole(roleName string)*role{
	return a.roles.appendRole(roleName)
}

// get role from the acl instance
func (a *Acl)GetRole(roleName string) (*role, error){
	return a.roles.getRole(roleName)
}

// add resource to acl instance
func (a *Acl)AddResource(resourceName string)*resource{
	return a.resources.appendResource(resourceName)
}