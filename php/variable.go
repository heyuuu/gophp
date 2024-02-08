package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

type iVariable interface {
	Get() types.Zval
	Set(value types.Zval)
	Unset()
	MakeRef() *types.Reference
}

type symbolVariable struct {
	symbols ISymtable
	name    string
}

func newSymbolVariable(symbols ISymtable, name string) *symbolVariable {
	return &symbolVariable{symbols: symbols, name: name}
}

func (s symbolVariable) Get() types.Zval {
	return s.symbols.Get(s.name)
}

func (s symbolVariable) Set(value types.Zval) {
	if !value.IsRef() {
		raw := s.symbols.Get(s.name)
		if raw.IsRef() {
			raw.Ref().SetVal(value)
			return
		}
	}
	s.symbols.Set(s.name, value)
}

func (s symbolVariable) Unset() {
	s.symbols.Unset(s.name)
}

func (s symbolVariable) MakeRef() *types.Reference {
	raw := s.symbols.Get(s.name)
	if raw.IsRef() {
		return raw.Ref()
	} else {
		ref := types.NewReference(raw)
		s.symbols.Set(s.name, types.ZvalRef(ref))
		return ref
	}
}

type arrayAppendVariable struct {
	ctx *Context
	arr types.Zval
}

func newArrayAppendVariable(ctx *Context, arr types.Zval) *arrayAppendVariable {
	return &arrayAppendVariable{ctx: ctx, arr: arr}
}

func (v arrayAppendVariable) Get() types.Zval {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		return types.Undef
	default:
		panic(perr.Todof("unsupported arrayAppendVariable.Set arr type: %s", types.ZvalGetType(arr)))
	}
}

func (v arrayAppendVariable) Set(value types.Zval) {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Append(value) // todo ArrayAccess
	default:
		panic(perr.Todof("unsupported arrayAppendVariable.Set arr type: %s", types.ZvalGetType(arr)))
	}
}

func (v arrayAppendVariable) Unset() {
	panic(perr.Todof("$var[] cannot be unset"))
}

func (v arrayAppendVariable) MakeRef() *types.Reference {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		ref := types.NewReference(UninitializedZval())
		arr.Array().Append(types.ZvalRef(ref))
		return ref
	default:
		panic(perr.Todof("unsupported arrayAppendVariable.MakeRef arr type: %s", types.ZvalGetType(arr)))
	}
}

type arrayDimVariable struct {
	ctx *Context
	arr types.Zval
	key types.ArrayKey
}

func newArrayDimVariable(ctx *Context, arr types.Zval, key types.ArrayKey) *arrayDimVariable {
	return &arrayDimVariable{ctx: ctx, arr: arr, key: key}
}

func (v arrayDimVariable) Get() types.Zval {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		return arr.Array().Find(v.key)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported arrayDimVariable.Get arr type: %s", types.ZvalGetType(arr)))
	}
}

func (v arrayDimVariable) Set(value types.Zval) {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		if !value.IsRef() {
			raw := arr.Array().Find(v.key)
			if raw.IsRef() {
				raw.Ref().SetVal(value)
				return
			}
		}
		arr.Array().Update(v.key, value)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported arrayDimVariable.Set arr type: %s", types.ZvalGetType(arr)))
	}
}

func (v arrayDimVariable) Unset() {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Delete(v.key)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported arrayDimVariable.Set arr type: %s", types.ZvalGetType(arr)))
	}
}

func (v arrayDimVariable) MakeRef() *types.Reference {
	arr := v.arr
	switch arr.Type() {
	case types.IsArray:
		raw := arr.Array().Find(v.key)
		if raw.IsRef() {
			return raw.Ref()
		} else {
			if raw.IsUndef() {
				raw = UninitializedZval()
			}
			ref := types.NewReference(raw)
			arr.Array().Update(v.key, types.ZvalRef(ref))
			return ref
		}
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported arrayDimVariable.Set arr type: %s", types.ZvalGetType(arr)))
	}
}

type propertyVariable struct {
	obj    *types.Object
	member types.Zval
}

func newPropertyVariable(obj *types.Object, member types.Zval) *propertyVariable {
	return &propertyVariable{obj: obj, member: member}
}

func (v propertyVariable) Get() types.Zval {
	return v.obj.ReadProperty(v.member, BP_VAR_R)
}

func (v propertyVariable) Set(value types.Zval) {
	v.obj.WriteProperty(v.member, value)
}

func (v propertyVariable) Unset() {
	v.obj.UnsetProperty(v.member)
}

func (v propertyVariable) MakeRef() *types.Reference {
	raw := v.Get()
	if raw.IsRef() {
		return raw.Ref()
	} else {
		if raw.IsUndef() {
			raw = UninitializedZval()
		}
		ref := types.NewReference(raw)
		v.Set(types.ZvalRef(ref))
		return ref
	}
}
