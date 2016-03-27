package acl

type resource struct {
	name   string
	active bool
}

// returns new resource pointer
func newResource(name string) *resource {
	return &resource{
		name: name,
		active:true,
	}
}

// check if active flag is on
func (r *resource) isActive() bool {
	return r.active
}

// turn on the active flag
func (r *resource) activate() {
	r.active = true
}

// turn off the active flag
func (r *resource) deactivate() {
	r.active = false
}
