package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

type ISymtable interface {
	Isset(name string) bool
	Get(name string) *types.Zval
	Set(name string, value *types.Zval)
}

type Symtable struct {
	table map[string]*types.Zval
}

func NewSymtable() *Symtable {
	return &Symtable{
		table: make(map[string]*types.Zval),
	}
}

func (t *Symtable) Isset(name string) bool {
	_, exists := t.table[name]
	return exists
}

func (t *Symtable) Get(name string) *types.Zval {
	return t.table[name]
}

func (t *Symtable) Set(name string, value *types.Zval) {
	perr.Assert(value != nil)

	if v, exists := t.table[name]; exists {
		v.SetBy(value)
	} else {
		t.table[name] = value.Clone()
	}
}
