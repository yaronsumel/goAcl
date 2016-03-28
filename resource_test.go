package goAcl

import (
	"testing"
	"reflect"
)

func TestNewResource(t *testing.T){
	r := newResource("test")
	if reflect.TypeOf(r).String() != "*goAcl.resource"{
		t.FailNow()
	}
}

func TestIsActive(t *testing.T){
	r := newResource("test")
	r.active = true
	if r.isActive() == false{
		t.FailNow()
	}
	r.active = false
	if r.isActive() == true{
		t.FailNow()
	}
}

func TestActivate(t *testing.T){
	r := newResource("test")
	r.activate()
	if r.active == false{
		t.FailNow()
	}
}

func TestDeactivate(t *testing.T){
	r := newResource("test")
	r.deactivate()
	if r.active == true{
		t.FailNow()
	}
}
