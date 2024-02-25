package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func InitResourceType(typeName string) types.ResourceType {
	return types.ResourceType(typeName)
}

func RegisterResource(ctx *Context, resourcePtr any, resourceType types.ResourceType) *types.Resource {
	// resource 计数
	handle := 0
	return types.NewResource(handle, resourceType, resourcePtr)
}

func CloseResource(res *types.Resource) int {
	// todo 此处可能需要补充 unregister 相关逻辑
	return res.Close()
}

func iFetchResource(ctx *Context, res *types.Resource, resourceTypeName string, resourceType types.ResourceType) any {
	assert.Assert(res != nil)
	if resourceType == res.Type() {
		return res.Ptr()
	}
	if resourceTypeName != "" {
		Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): supplied resource is not a valid %s resource", ctx.CurrEX().CalleeName(), resourceTypeName))
	}
	return nil
}

func iFetchResource2(ctx *Context, res *types.Resource, resourceTypeName string, resourceType1 types.ResourceType, resourceType2 types.ResourceType) any {
	if res != nil {
		if resourceType1 == res.Type() || resourceType2 == res.Type() {
			return res.Ptr()
		}
	}
	if resourceTypeName != "" {
		Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): supplied resource is not a valid %s resource", ctx.CurrEX().CalleeName(), resourceTypeName))
	}
	return nil
}

func FetchResource[T any](ctx *Context, res *types.Resource, resourceTypeName string, resourceType types.ResourceType) *T {
	result := iFetchResource(ctx, res, resourceTypeName, resourceType)
	if v, ok := result.(*T); ok {
		return v
	}
	return nil
}

func FetchResourceEx[T any](ctx *Context, res *types.Zval, resourceTypeName string, resourceType types.ResourceType) *T {
	if res == nil {
		if resourceTypeName != "" {
			Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): no %s resource supplied", ctx.CurrEX().CalleeName(), resourceTypeName))
		}
		return nil
	}
	if !res.IsResource() {
		if resourceTypeName != "" {
			Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): supplied argument is not a valid %s resource", ctx.CurrEX().CalleeName(), resourceTypeName))
		}
		return nil
	}

	result := iFetchResource(ctx, res.Resource(), resourceTypeName, resourceType)
	if v, ok := result.(*T); ok {
		return v
	}
	return nil
}

func FetchResource2[T any](ctx *Context, res *types.Resource, resourceTypeName string, resourceType1 types.ResourceType, resourceType2 types.ResourceType) *T {
	result := iFetchResource2(ctx, res, resourceTypeName, resourceType1, resourceType2)
	if v, ok := result.(*T); ok {
		return v
	}
	return nil
}

func FetchResource2Ex[T any](ctx *Context, res *types.Zval, resourceTypeName string, resourceType1 types.ResourceType, resourceType2 types.ResourceType) *T {
	if res == nil {
		if resourceTypeName != "" {
			Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): no %s resource supplied", ctx.CurrEX().CalleeName(), resourceTypeName))
		}
		return nil
	}
	if !res.IsResource() {
		if resourceTypeName != "" {
			Error(ctx, perr.E_WARNING, fmt.Sprintf("%s(): supplied argument is not a valid %s resource", ctx.CurrEX().CalleeName(), resourceTypeName))
		}
		return nil
	}
	return FetchResource2[T](ctx, res.Resource(), resourceTypeName, resourceType1, resourceType2)
}
