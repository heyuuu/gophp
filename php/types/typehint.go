package types

import b "github.com/heyuuu/gophp/builtin"

/**
 * TypeHint
 */
const (
	typeHintAllowNull = 1 << iota
	typeHintIsCode
	typeHintIsClass
	typeHintIsCe
	typeHintIsName
)

type TypeHint struct {
	flags uint8
	code  ZvalType
	ce    *ClassEntry
	name  string
}

func TypeHintCode(code ZvalType, allowNull bool) TypeHint {
	var flags uint8 = typeHintIsCode
	if allowNull {
		flags |= typeHintAllowNull
	}
	return TypeHint{flags: flags, code: code}
}
func TypeHintCe(ce *ClassEntry, allowNull bool) TypeHint {
	var flags uint8 = typeHintIsClass | typeHintIsCe
	if allowNull {
		flags |= typeHintAllowNull
	}
	return TypeHint{flags: flags, ce: ce}
}
func TypeHintClassName(className string, allowNull bool) TypeHint {
	var flags uint8 = typeHintIsClass | typeHintIsName
	if allowNull {
		flags |= typeHintAllowNull
	}
	return TypeHint{flags: flags, name: className}
}

func (t TypeHint) IsSet() bool     { return t.flags != 0 }
func (t TypeHint) AllowNull() bool { return t.flags&typeHintAllowNull != 0 }
func (t TypeHint) IsCode() bool    { return t.flags&typeHintIsCode != 0 }
func (t TypeHint) IsClass() bool   { return t.flags&typeHintIsClass != 0 }
func (t TypeHint) IsCe() bool      { return t.flags&typeHintIsCe != 0 }
func (t TypeHint) IsName() bool    { return t.flags&typeHintIsName != 0 }

func (t TypeHint) Ce() *ClassEntry  { return t.ce }
func (t TypeHint) Name() string     { return t.name }
func (t TypeHint) Code() ZvalType   { return t.code }
func (t TypeHint) CodeName() string { return ZendGetTypeByConst(t.code) }

func (t *TypeHint) FormatName() string {
	b.Assert(t.IsSet())
	if t.IsClass() {
		return t.Name()
	} else {
		return t.CodeName()
	}
}

func (t *TypeHint) FormatType() string {
	var typ string
	if t.IsClass() {
		if t.IsCe() {
			typ = t.Ce().Name()
		} else {
			typ = t.Name()
		}
	} else {
		typ = t.CodeName()
	}
	if t.AllowNull() {
		typ = "?" + typ
	}
	return typ
}
