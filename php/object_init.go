package php

import (
	"github.com/heyuuu/gophp/php/types"
)

func NewStdObject(ctx *Context, ce *types.Class) *types.Object {
	//return types.NewObjectEx(ce, StdObjectHandlersPtr, nil)
	handle := ctx.EG().NextObjectHandle()
	return types.NewObject(ce, handle)
}

//
//func NewStdObjectEx(ce *types.Class, properties *types.Array) *types.Object {
//	return types.NewObjectEx(ce, StdObjectHandlersPtr, properties)
//}

//func NewStdClassObject(properties *types.Array) *types.Object {
//	obj := NewStdObject(ZendStandardClassDef)
//	obj.SetProperties(properties)
//	return obj
//}

//
//func NewStdObjectSkipPropertiesInit(ce *types.Class) *types.Object {
//	return types.NewObjectSkipPropertiesInit(ce, StdObjectHandlersPtr)
//}

func ObjectAndPropertiesInit(ctx *Context, classType *types.Class, properties *types.Array) *types.Object {
	return _objectAndPropertiesInit(ctx, classType, properties)
}
func ObjectInitEx(ctx *Context, classType *types.Class) *types.Object {
	return _objectAndPropertiesInit(ctx, classType, nil)
}

//	func ObjectInit(arg *types.Zval) bool {
//		arg.SetObject(NewStdClassObject(nil))
//		return true
//	}
func _objectAndPropertiesInit(ctx *Context, classType *types.Class, properties *types.Array) *types.Object {
	var obj *types.Object
	obj = NewStdObject(ctx, classType)
	return obj

	//if classType.HasCeFlags(types.AccInterface | types.AccTrait | types.AccImplicitAbstractClass | types.AccExplicitAbstractClass) {
	//	if classType.IsInterface() {
	//		ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate interface %s", classType.Name()))
	//	} else if classType.IsTrait() {
	//		ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate trait %s", classType.Name()))
	//	} else {
	//		ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate abstract class %s", classType.Name()))
	//	}
	//	arg.SetNull()
	//	return false
	//}
	//if !classType.IsConstantsUpdated() {
	//	if ZendUpdateClassConstants(ctx, classType) != types.SUCCESS {
	//		arg.SetNull()
	//		return false
	//	}
	//}
	//if classType.GetCreateObject() == nil {
	//var obj *types.Object
	//if properties != nil {
	//	obj = NewStdObjectEx(classType, properties)
	//} else {
	//obj = NewStdObject(classType)
	//}
	//arg.SetObject(obj)
	//} else {
	//	arg.SetObject(classType.GetCreateObject()(classType))
	//}
}
