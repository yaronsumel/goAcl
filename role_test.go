package goAcl

import (
	"testing"
	"reflect"
)

func TestNewRole(t *testing.T){
	r := newRole("test")
	if reflect.TypeOf(r).String() != "*goAcl.role"{
		t.FailNow()
	}
}

func TestInherits(t *testing.T){
	a := NewAcl()
	role1:=a.AddRole("test1").AllowResource("r1")
	role2:=a.AddRole("test2").Inherits(role1)
	if role2.resources.resourceMap["r1"].name != "r1"{
		t.FailNow()
	}
}

func TestAllowResource(t *testing.T){
	a := NewAcl()
	role:=a.AddRole("test1").AllowResource("r1")
	resource, ok := role.resources.resourceMap["r1"]
	if !ok {
		t.FailNow()
	}
	if resource.name != "r1"{
		t.FailNow()
	}
}

func TestDenyResource(t *testing.T){
	a := NewAcl()
	role:=a.AddRole("test1").AllowResource("r1")
	if role.IsAllowed("r1") != true{
		t.FailNow()
	}
	role.DenyResource("r1")
	if role.IsAllowed("r1") != false{
		t.FailNow()
	}
}

func TestIsAllowed(t *testing.T){
	a := NewAcl()
	role:=a.AddRole("test1").AllowResource("r1")
	if role.IsAllowed("r1") != role.resources.resourceMap["r1"].isActive(){
		t.FailNow()
	}
}
