package goAcl

import (
	"testing"
	"reflect"
)

func TestNewResources(t *testing.T){
	r := newResources()
	if reflect.TypeOf(r).String() != "goAcl.resources"{
		t.FailNow()
	}
}

func TestAppendResource(t *testing.T){
	r := newResources()
	r.appendResource("test")
	resource,err := r.getResource("test")
	if err != nil{
		t.FailNow()
	}
	if resource.name != "test"{
		t.FailNow()
	}
}

func TestGetResource(t *testing.T){
	r := newResources()
	r.appendResource("test")
	resource,err := r.getResource("test")
	if err != nil{
		t.FailNow()
	}
	if resource.name != r.resourceMap["test"].name{
		t.FailNow()
	}
}
