package php

import "github.com/heyuuu/gophp/php/types"

// properties for Context
func (c *Context) Executor() *Executor {
	return c.executor
}

// properties for IniEntry
func (ini *IniEntry) ModuleNumber() int {
	return ini.moduleNumber
}
func (ini *IniEntry) Name() string {
	return ini.name
}
func (ini *IniEntry) Value() string {
	return ini.value
}
func (ini *IniEntry) HasValue() bool {
	return ini.hasValue
}
func (ini *IniEntry) Modifiable() IniModifiable {
	return ini.modifiable
}

// properties for IniEntryDef
func (d *IniEntryDef) Name() string {
	return d.name
}

// properties for StdInternObject
func (o *StdInternObject) Ctx() *Context {
	return o.ctx
}
func (o *StdInternObject) Obj() *types.Object {
	return o.obj
}
func (o *StdInternObject) Class() *types.Class {
	return o.class
}
