package goAcl

import (
	"testing"
	"reflect"
)

func TestNewAcl(t *testing.T){
	acl := NewAcl()
	if reflect.TypeOf(acl.roles.rolesMap).String()  != "map[string]*goAcl.role" {
		t.FailNow()
	}
	if reflect.TypeOf(acl.resources.resourceMap).String()  != "map[string]*goAcl.resource" {
		t.FailNow()
	}
}

func TestAddRole(t *testing.T){
	acl := NewAcl()
	acl.AddRole("test")
	if acl.roles.rolesMap["test"].name != "test"{
		t.FailNow()
	}
}

func TestGetRole(t *testing.T){
	acl := NewAcl()
	acl.AddRole("test")
	role,err := acl.GetRole("test")
	if err != nil{
		t.FailNow()
	}
	if acl.roles.rolesMap["test"] != role{
		t.FailNow()
	}
}

func TestAddResource(t *testing.T){
	acl := NewAcl()
	acl.AddResource("test")
	if acl.resources.resourceMap["test"].name != "test"{
		t.FailNow()
	}
}