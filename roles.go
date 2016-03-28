package goAcl

import (
	"errors"
)

type roles struct {
	rolesMap map[string]*role
}

// returns new roles pointer
func newRoles()roles{
	return roles{
		rolesMap: make(map[string]*role),
	}
}

// adds new role to roleMap if its not exist yet
func (r *roles)appendRole(roleName string)*role{
	role,err := r.getRole(roleName)
	if err!=nil{
		r.rolesMap[roleName] = newRole(roleName)
		role = r.rolesMap[roleName]
	}
	return role
}

// get role from roleMap by name
func (r *roles)getRole(roleName string) (*role, error) {
	if role, ok := r.rolesMap[roleName]; ok {
		return role, nil
	}
	return nil, errors.New("no such role:" + roleName)
}