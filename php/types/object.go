package types

// Object
type Object struct {
	// todo
	ce *Class
}

func (o Object) Ce() *Class { return o.ce }

func (o Object) CanCast() bool {
	// todo
	return false
}

func (o Object) Cast(typ ZvalType) (*Zval, bool) {
	// todo
	return nil, false
}
