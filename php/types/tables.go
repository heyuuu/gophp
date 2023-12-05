package types

import (
	"github.com/heyuuu/gophp/kits/ascii"
)

// ConstantTable
type ConstantTable struct {
	table *Table[*Constant]
}

func NewConstantTable() *ConstantTable {
	return &ConstantTable{table: NewTable[*Constant]()}
}

func (t *ConstantTable) Len() int                  { return t.table.Len() }
func (t *ConstantTable) Exists(name string) bool   { return t.table.Exists(name) }
func (t *ConstantTable) Get(name string) *Constant { return t.table.Get(name) }
func (t *ConstantTable) Each(handler func(c *Constant)) {
	t.table.Each(func(_ string, c *Constant) { handler(c) })
}
func (t *ConstantTable) Add(name string, c *Constant) bool { return t.table.Add(name, c) }

// FunctionTable
type FunctionTable struct {
	table *Table[*Function]
}

func NewFunctionTable() *FunctionTable {
	return &FunctionTable{table: NewTable[*Function]()}
}
func (t *FunctionTable) Len() int { return t.table.Len() }
func (t *FunctionTable) Exists(name string) bool {
	return t.table.Exists(ascii.StrToLower(name))
}
func (t *FunctionTable) Get(name string) *Function {
	return t.table.Get(ascii.StrToLower(name))
}
func (t *FunctionTable) Each(handler func(key string, fn *Function)) {
	t.table.Each(handler)
}
func (t *FunctionTable) EachEx(handler func(key string, fn *Function) error) error {
	return t.table.EachEx(handler)
}
func (t *FunctionTable) Add(name string, fn *Function) bool { return t.table.Add(name, fn) }
func (t *FunctionTable) Set(name string, fn *Function)      { t.table.Set(name, fn) }
func (t *FunctionTable) Del(name string)                    { t.table.Del(name) }

// FunctionTable
type ClassTable struct {
	table *Table[*Class]
}

func NewClassTable() *ClassTable {
	return &ClassTable{table: NewTable[*Class]()}
}
func (t *ClassTable) Len() int { return t.table.Len() }
func (t *ClassTable) Exists(name string) bool {
	return t.table.Exists(ascii.StrToLower(name))
}
func (t *ClassTable) Get(name string) *Class {
	return t.table.Get(ascii.StrToLower(name))
}
func (t *ClassTable) Each(handler func(key string, ce *Class)) {
	t.table.Each(handler)
}
func (t *ClassTable) EachEx(handler func(key string, ce *Class) error) error {
	return t.table.EachEx(handler)
}
func (t *ClassTable) Add(name string, ce *Class) bool { return t.table.Add(name, ce) }
func (t *ClassTable) Set(name string, ce *Class)      { t.table.Set(name, ce) }
func (t *ClassTable) Del(name string)                 { t.table.Del(name) }
