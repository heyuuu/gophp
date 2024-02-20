package php

import (
	"strconv"
)

// --- constants
type IniModifiable uint8

const (
	IniUser   IniModifiable = 1 << 0
	IniPerDir IniModifiable = 1 << 1
	IniSystem IniModifiable = 1 << 2
	IniAll    IniModifiable = IniUser | IniPerDir | IniSystem
)

func (mod IniModifiable) Match(target IniModifiable) bool { return mod&target == target }

type IniStage uint8

const (
	IniStageStartup    IniStage = 1 << 0
	IniStageShutdown   IniStage = 1 << 1
	IniStageActivate   IniStage = 1 << 2
	IniStageDeactivate IniStage = 1 << 3
	IniStageRuntime    IniStage = 1 << 4
	IniStageHtaccess   IniStage = 1 << 5
)

// -- types

type onModifyFunc func(ctx *Context, entry *IniEntry, newValue string, hasValue bool, stage IniStage) bool

// IniEntryDef
type IniEntryDef struct {
	name       string `get:""`
	onModify   onModifyFunc
	value      string
	hasValue   bool
	modifiable IniModifiable
}

func NewIniEntryDef(name string, modifiable IniModifiable) *IniEntryDef {
	return &IniEntryDef{name: name, modifiable: modifiable}
}
func (d *IniEntryDef) Value(value string) *IniEntryDef {
	d.value, d.hasValue = value, true
	return d
}
func (d *IniEntryDef) OnModify(onModify onModifyFunc) *IniEntryDef {
	d.onModify = onModify
	return d
}
func (d *IniEntryDef) OnModifyBool(handler func(*Context, bool)) *IniEntryDef {
	return d.OnModify(func(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
		val := IniStringParseBool(newValue)
		handler(ctx, val)
		return true
	})
}
func (d *IniEntryDef) OnModifyLong(handler func(*Context, int)) *IniEntryDef {
	return d.OnModify(func(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
		val, ok := ParseLongWithUnit(newValue)
		if !ok || val < 0 {
			return false
		}
		handler(ctx, val)
		return true
	})
}
func (d *IniEntryDef) OnModifyLongGEZero(handler func(*Context, int)) *IniEntryDef {
	return d.OnModify(func(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
		val, ok := ParseLongWithUnit(newValue)
		if !ok || val < 0 {
			return false
		}
		handler(ctx, val)
		return true
	})
}
func (d *IniEntryDef) OnModifyString(handler func(*Context, string)) *IniEntryDef {
	return d.OnModify(func(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
		handler(ctx, newValue)
		return true
	})
}
func (d *IniEntryDef) OnModifyStringNotEmpty(handler func(*Context, string)) *IniEntryDef {
	return d.OnModify(func(ctx *Context, entry *IniEntry, newValue string, hasValue bool, stage IniStage) bool {
		if hasValue && newValue == "" {
			return false
		}
		handler(ctx, newValue)
		return true
	})
}

// IniEntry: immutable ini entry
type IniEntry struct {
	moduleNumber int    `get:""`
	name         string `get:""`
	onModify     onModifyFunc
	value        string        `get:""`
	hasValue     bool          `get:""`
	modifiable   IniModifiable `get:""`
}

func NewIniEntry(def *IniEntryDef, moduleNumber int) *IniEntry {
	return &IniEntry{
		name:         def.name,
		onModify:     def.onModify,
		value:        def.value,
		hasValue:     def.hasValue,
		modifiable:   def.modifiable,
		moduleNumber: moduleNumber,
	}
}

func (ini *IniEntry) EmitOnModify(ctx *Context, value string, hasValue bool, stage IniStage) bool {
	if ini.onModify == nil {
		return true
	}
	return ini.onModify(ctx, ini, value, hasValue, stage)
}
func (ini *IniEntry) EmitOnModifyBy(ctx *Context, by *IniEntry, stage IniStage) bool {
	if ini.onModify == nil {
		return true
	}
	return ini.onModify(ctx, ini, by.value, by.hasValue, stage)
}

func (ini *IniEntry) WithModifiable(value IniModifiable) *IniEntry {
	if ini.modifiable == value {
		return ini
	}
	dup := *ini
	dup.modifiable = value
	return &dup
}

func (ini *IniEntry) WithNewValue(value string, hasValue bool) *IniEntry {
	if !hasValue {
		value = ""
	}
	if ini.value == value && ini.hasValue == hasValue {
		return ini
	}
	dup := *ini
	dup.value = value
	dup.hasValue = hasValue
	return &dup
}

// --- functions

func IniStringParseBool(str string) bool {
	if str == "true" || str == "yes" || str == "on" {
		return true
	} else {
		v, _ := strconv.Atoi(str)
		return v != 0
	}
}
