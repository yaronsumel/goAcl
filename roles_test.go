package goAcl

import (
	"testing"
	"reflect"
)

func TestNewRoles(t *testing.T){
	r := newRoles()
	if reflect.TypeOf(r).String() != "goAcl.roles"{
		t.FailNow()
	}
}

func TestAppendRole(t *testing.T){
	r := newRoles()
	r.appendRole("test")
	if r.rolesMap["test"].name != "test"{
		t.FailNow()
	}
}

func TestGetRoleFromRoles(t *testing.T){
	r := newRoles()
	r.appendRole("test")
	role,err := r.getRole("test")
	if err!=nil{
		t.FailNow()
	}
	if r.rolesMap["test"] != role{
		t.FailNow()
	}
}