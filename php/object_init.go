package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/types"
)

// 跳过对 classType 的检查，直接创建 object 实例
func StdObjectInitDirect(ctx *Context, classType *types.Class) *types.Object {
	handle := ctx.EG().NextObjectHandle()
	intern := NewStdInternObject(ctx, classType)
	return types.NewObject(handle, intern)
}

func IsStdGetProperties(o *types.Object) bool {
	// 目前没有很好的方法判断继承(组合)了 StdInternObject 的类型是否保留了 GetProperties 方法，所以目前只判断是否直接使用 StdInternObject
	//std := php.ZendStdGetProperties
	return IsStdObject(o)
}

func IsStdObject(o *types.Object) bool {
	_, ok := o.Intern().(*StdInternObject)
	return ok
}

// 常规对象初始化入口

func ObjectInit(ctx *Context, classType *types.Class) *types.Object {
	if classType.HasFlags(types.AccInterface | types.AccTrait | types.AccImplicitAbstractClass | types.AccExplicitAbstractClass) {
		if classType.IsInterface() {
			ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate interface %s", classType.Name()))
		} else if classType.IsTrait() {
			ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate trait %s", classType.Name()))
		} else {
			ThrowError(ctx, nil, fmt.Sprintf("Cannot instantiate abstract class %s", classType.Name()))
		}
		return nil
	}
	if !classType.IsLinked() {
		DoLinkClass(ctx, classType)
	}
	if !classType.IsConstantsUpdated() {

	}
	if classType.GetCreateObject() == nil {
		return StdObjectInitDirect(ctx, classType)
	} else {
		return classType.GetCreateObject()(classType)
	}
}

func ObjectInitZval(ctx *Context, arg *types.Zval, classType *types.Class) bool {
	obj := ObjectInit(ctx, classType)
	if obj == nil {
		arg.SetNull()
		return false
	}

	arg.SetObject(obj)
	return true
}
