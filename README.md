# goAcl
goAcl is simple application control list (ACl) package, provides basic ACL features in order to control and manage roles and resources in your app.


## Install
    $ go get github.com/yaronsumel/goAcl

## Usage
import and create new goAcl instance
```go
import "github.com/yaronsumel/goAcl"
acl := goAcl.NewAcl()
```

## Roles
first we need to create user roles
```go
guest 	:= acl.AddRole("guest")
member  := acl.AddRole("member")
```


## Resources
now we need to add possible app resources
```go
acl.AddResource("view")
acl.AddResource("edit")
```

## Allow Specific Resource
```go
guest.AllowResource("view")
member.AllowResource("edit").Inherits(guest)
```

## Deny Specific Resource
```go
member.DenyResource("view")
```

## Basic Usage
```go
//get the role from user
role,err := acl.GetRole("member")

if role.IsAllowed("edit"){
//user can edit things now
}
```
