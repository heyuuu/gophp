package php

import (
	"github.com/heyuuu/gophp/php/types"
)

func visibilityString(flags uint32) string {
	if (flags & types.AccPublic) != 0 {
		return "public"
	} else if (flags & types.AccPrivate) != 0 {
		return "private"
	} else {
		Assert((flags & types.AccProtected) != 0)
		return "protected"
	}
}

func inheritProperty(ctx *Context, parentInfo *types.PropertyInfo, key string, ce *types.Class) {
	var childInfo = ce.PropertyTable().Get(key)
	if childInfo != nil {
		if parentInfo.HasFlags(types.AccPrivate | types.AccChanged) {
			childInfo.MarkIsChanged()
		}
		if !parentInfo.IsPrivate() {

		}
	} else {
		childInfo = parentInfo
		ce.PropertyTable().Add(key, childInfo)
	}
}

func doInheritanceEx(ctx *Context, ce *types.Class, parentCe *types.Class, checked bool) {
	ce.SetParent(parentCe)
	ce.SetIsResolvedParent(true)

	/* Inherit properties */
	if true {

	}

	ce.PropertyTable().Each(func(key string, propInfo *types.PropertyInfo) {
		if propInfo.Ce() == ce {

		}
	})
}

func DoLinkClass(ctx *Context, ce *types.Class) {
	if ce.IsLinked() {
		return
	}

	var parent *types.Class
	if ce.ParentName() != "" {
		parent = ZendFetchClassByName(ctx, ce.ParentName(), "", 0)
		if parent == nil {
			panic("parent is not found: " + ce.ParentName())
		}
		ce.SetParent(parent)
	}

	if parent != nil {
		parent.PropertyTable().Each(func(name string, parentPropInfo *types.PropertyInfo) {
			childPropInfo := ce.PropertyTable().Get(name)
			if childPropInfo == nil {
				ce.PropertyTable().Add(name, parentPropInfo)
			} else {
				// todo 检验 prop 是否兼容
			}
		})
	}

	// fix offset
	var propOffset uint32 = 0
	var staticOffset uint32 = 0
	ce.PropertyTable().Each(func(s string, info *types.PropertyInfo) {
		if info.IsStatic() {
			info.SetOffset(staticOffset)
			staticOffset++
		} else {
			info.SetOffset(propOffset)
			propOffset++
		}
	})
}
