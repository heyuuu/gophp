package types

type ObjectData interface {
	Class() *Class

	// property
	ReadProperty(member Zval) Zval
	WriteProperty(member Zval, value Zval)
	HasProperty(member Zval, hasSetExists int) bool
	UnsetProperty(member Zval)

	// properties
	GetPropertiesArray() *Array
	GetPropertiesFor(purpose PropPurposeType) *Array

	// dimension
	ReadDimension(offset Zval, typ int) Zval
	WriteDimension(offset Zval, value Zval)
	HasDimension(offset Zval, checkEmpty int) bool
	UnsetDimension(offset Zval)

	// elements
	CountElements() (int, bool)

	// methods
	GetMethod(method string) *Function
	GetConstructor(object *Object) *Function

	// mixed
	Cast(typ ZvalType) (Zval, bool)
}
