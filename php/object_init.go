package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/types"
)

// 跳过对 classType 的检查，直接创建 object 实例
func ObjectInitDirect(ctx *Context, classType *types.Class) *types.Object {
	handle := ctx.EG().NextObjectHandle()
	objectData := NewStdObjectData(ctx, classType)
	return types.NewObject(classType, handle, objectData)
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

	// todo check properties
	return ObjectInitDirect(ctx, classType)
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
