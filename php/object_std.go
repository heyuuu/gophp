package php

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// StdInternObject
type StdInternObject struct {
	ctx             *Context              `get:""`
	obj             *types.Object         `get:""`
	class           *types.Class          `get:""`
	propertiesTable []types.Zval          // 静态属性
	properties      map[string]types.Zval // 动态属性
}

var _ types.IObject = (*StdInternObject)(nil)
var _ types.ObjectBindable = (*StdInternObject)(nil)

func NewStdInternObject(ctx *Context, ce *types.Class) *StdInternObject {
	o := &StdInternObject{
		ctx:             ctx,
		class:           ce,
		propertiesTable: make([]types.Zval, ce.PropertyTable().Len()),
		properties:      make(map[string]types.Zval),
	}
	return o
}

func MakeStdInternObject(ctx *Context, ce *types.Class) StdInternObject {
	return StdInternObject{ctx: ctx, class: ce}
}

func (o *StdInternObject) PropertiesTable() []types.Zval     { return o.propertiesTable }
func (o *StdInternObject) Properties() map[string]types.Zval { return o.properties }
func (o *StdInternObject) StdInternObject() *StdInternObject { return o }

func (o *StdInternObject) Bind(obj *types.Object) { o.obj = obj }
func (o *StdInternObject) CanClone() bool         { return true }
func (o *StdInternObject) Clone() *types.Object {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) getPropertyInfo(member string, silent bool) (*types.PropertyInfo, bool) {
	var propInfo *types.PropertyInfo
	lcMember := ascii.StrToLower(member)
	if propInfo = o.class.PropertyTable().Get(lcMember); propInfo == nil {
		return nil, true
	}
	// todo 可访问性检查
	return propInfo, true
}

func (o *StdInternObject) ReadProperty(member types.Zval, typ int) types.Zval {
	ctx := o.ctx

	name, ok := ZvalTryGetStr(ctx, member)
	if !ok {
		return UninitializedZval()
	}

	/* make zend_get_property_info silent if we have getter - we may want to use it */
	propInfo, ok := o.getPropertyInfo(name, typ == BP_VAR_IS || o.class.GetGet() != nil)
	if ok {
		if propInfo != nil {
			return o.propertiesTable[propInfo.Offset()]
		} else {
			value := o.properties[name]
			if value.IsNotUndef() {
				return value
			}
		}
	} else {
		return UninitializedZval()
	}

	if typ != BP_VAR_IS {
		if propInfo != nil {
			ThrowError(ctx, nil, fmt.Sprintf("Typed property %s::$%s must not be accessed before initialization", propInfo.Ce().Name(), name))
		} else {
			Error(ctx, perr.E_NOTICE, fmt.Sprintf("Undefined property: %s::$%s", o.class.Name(), name))
		}
	}
	return UninitializedZval()
}

func (o *StdInternObject) WriteProperty(member types.Zval, value types.Zval) {
	ctx := o.ctx
	name, ok := ZvalTryGetStr(ctx, member)
	if !ok {
		return
	}

	propInfo, ok := o.getPropertyInfo(name, false)
	if propInfo != nil {
		o.propertiesTable[propInfo.Offset()] = value
	} else {
		o.properties[name] = value
	}
}

func (o *StdInternObject) HasProperty(member types.Zval, hasSetExists int) bool {
	name := ZvalGetStrVal(o.ctx, member)
	_, exists := o.properties[name]
	return exists
}

func (o *StdInternObject) UnsetProperty(member types.Zval) {
	name := ZvalGetStrVal(o.ctx, member)
	delete(o.properties, name)
}

func (o *StdInternObject) GetPropertyPtr(member types.Zval, typ int) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) GetPropertiesArray() *types.Array {
	arr := types.NewArrayCap(len(o.properties))
	o.class.PropertyTable().Each(func(_ string, info *types.PropertyInfo) {
		arr.KeyAdd(info.Name(), o.propertiesTable[info.Offset()])
	})
	for key, value := range o.properties {
		arr.KeyAdd(key, value)
	}
	return arr
}

func (o *StdInternObject) GetPropertiesFor(purpose types.PropPurposeType) *types.Array {
	return nil
}

func (o *StdInternObject) ReadDimension(offset types.Zval, typ int) types.Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) WriteDimension(offset types.Zval, value types.Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) HasDimension(offset types.Zval, checkEmpty int) bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) UnsetDimension(offset types.Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) CountElements() (int, bool) {
	return 0, false
}

func (o *StdInternObject) GetMethod(method string) *types.Function {
	lcName := ascii.StrToLower(method)
	return o.class.FunctionTable().Get(lcName)
}

func (o *StdInternObject) GetConstructor(object *types.Object) *types.Function {
	//TODO implement me
	panic("implement me")
}

// see: zend_std_cast_object_tostring
func (o *StdInternObject) Cast(typ types.ZvalType) (types.Zval, bool) {
	ctx := o.ctx
	obj := o.obj
	switch typ {
	case types.IsString:
		if str, ok := StdCastObjectToString(ctx, obj); ok {
			return String(str), true
		}
		return types.Null, false
	case types.IsBool:
		return types.True, true
	case types.IsLong:
		className := obj.ClassName()
		Error(ctx, perr.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to int", className))
		return Long(1), true
	case types.IsDouble:
		className := obj.ClassName()
		Error(ctx, perr.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to float", className))
		return Double(1), true
	case types.IsNumber:
		className := obj.ClassName()
		Error(ctx, perr.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to number", className))
		return Long(1), true
	default:
		return types.Null, false
	}
}

func (o *StdInternObject) CompareObjectTo(another *types.Object) (int, bool) {
	//TODO implement me
	panic("implement me")
}

func (o *StdInternObject) CompareTo(another types.Zval) (int, bool) {
	return 0, false
}

func (o *StdInternObject) GetClosure() (*types.ClosureData, bool) {
	return nil, false
}
