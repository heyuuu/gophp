package php

import "github.com/heyuuu/gophp/php/types"

var _ types.ObjectData = (*StdObjectData)(nil)

type StdObjectData struct {
	ctx        *Context
	ce         *types.Class
	properties map[string]types.Zval
}

func NewStdObjectData(ctx *Context, ce *types.Class) *StdObjectData {
	return &StdObjectData{
		ctx:        ctx,
		ce:         ce,
		properties: make(map[string]types.Zval),
	}
}

func (o *StdObjectData) Class() *types.Class { return o.ce }
func (o *StdObjectData) ReadProperty(member types.Zval) types.Zval {
	name := ZvalGetStrVal(o.ctx, member)
	return o.properties[name]
}

func (o *StdObjectData) WriteProperty(member types.Zval, value types.Zval) {
	name := ZvalGetStrVal(o.ctx, member)
	o.properties[name] = value
}

func (o *StdObjectData) HasProperty(member types.Zval, hasSetExists int) bool {
	name := ZvalGetStrVal(o.ctx, member)
	_, exists := o.properties[name]
	return exists
}

func (o *StdObjectData) UnsetProperty(member types.Zval) {
	name := ZvalGetStrVal(o.ctx, member)
	delete(o.properties, name)
}

func (o *StdObjectData) GetPropertiesArray() *types.Array {
	arr := types.NewArrayCap(len(o.properties))
	for key, value := range o.properties {
		arr.KeyAdd(key, value)
	}
	return arr
}

func (o *StdObjectData) GetPropertiesFor(purpose types.PropPurposeType) *types.Array {
	return o.GetPropertiesArray()
}

func (o *StdObjectData) ReadDimension(offset types.Zval, typ int) types.Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) WriteDimension(offset types.Zval, value types.Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) HasDimension(offset types.Zval, checkEmpty int) bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) UnsetDimension(offset types.Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CountElements() (int, bool) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetMethod(method string) *types.Function {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetConstructor(object *types.Object) *types.Function {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) Cast(typ types.ZvalType) (types.Zval, bool) {
	//TODO implement me
	panic("implement me")
}
