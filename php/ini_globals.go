package php

import (
	"github.com/heyuuu/gophp/kits/mapkit"
	"github.com/heyuuu/gophp/php/types"
	"maps"
	"slices"
	"strings"
)

type IniDirectives = map[string]*IniEntry

type IniGlobals struct {
	ctx *Context

	iniEntries      string             `prop:""`
	iniIgnore       bool               `prop:""`
	iniIgnoreCwd    bool               `prop:""`
	iniPathOverride string             `prop:""`
	iniDefaultsFunc func(*types.Array) `prop:""`

	// config
	hash             *types.Array
	hasPerDirConfig  bool     `get:""`
	hasPerHostConfig bool     `get:""`
	zendExtensions   []string `get:""`
	phpExtensions    []string `get:""`
	iniOpenedPath    string   `prop:""`
	iniScannedPath   string   `prop:""`
	iniScannedFiles  string   `get:""`

	// directives
	iniDirectives         IniDirectives
	modifiedIniDirectives IniDirectives
}

func (ig *IniGlobals) Init(ctx *Context, base *IniGlobals) {
	ig.ctx = ctx
	if base != nil {
		ig.iniDirectives = base.iniDirectives
		ig.modifiedIniDirectives = make(IniDirectives)
	} else {
		ig.iniDirectives = make(IniDirectives)
		ig.modifiedIniDirectives = make(IniDirectives)
	}
}

// ini config

func (ig *IniGlobals) AppendIniEntries(iniEntries string) {
	if strings.TrimSpace(iniEntries) == "" {
		return
	}

	ig.iniEntries += "\n" + iniEntries
}

func (ig *IniGlobals) ConfigInit() {
	ig.hash = types.NewArray()
	if ig.iniDefaultsFunc != nil {
		ig.iniDefaultsFunc(ig.hash)
	}
	ig.phpExtensions = nil
	ig.zendExtensions = nil
}
func (ig *IniGlobals) ConfigShutdown() {
	ig.hash.Clean()
	ig.iniOpenedPath = ""
	ig.iniScannedFiles = ""
}
func (ig *IniGlobals) ConfigSet(key string, value string) {
	ig.hash.KeyUpdate(key, types.ZvalString(value))
}
func (ig *IniGlobals) ConfigGet(key string) types.Zval {
	return ig.hash.KeyFind(key)
}
func (ig *IniGlobals) ConfigGetStr(key string) (string, bool) {
	conf := ig.hash.KeyFind(key)
	if conf.IsString() {
		return conf.String(), true
	}
	return "", false
}
func (ig *IniGlobals) ConfigGetLong(key string) (int, bool) {
	conf := ig.hash.KeyFind(key)
	if conf.IsUndef() {
		return 0, false
	}
	return ZvalGetLong(ig.ctx, conf), true
}
func (ig *IniGlobals) ConfigHash() *types.Array {
	return ig.hash
}

func (ig *IniGlobals) LoadIniStr(s string) bool {
	cb := NewDefaultIniParserCb(ig.ctx, ig.hash)
	return IniParse(ig.ctx, s, IniScanNormal, cb)
}

func (ig *IniGlobals) LoadIniFile(fh *FileHandle) bool {
	str, ok := fh.ReadAll()
	if !ok {
		return false
	}
	return ig.LoadIniStr(str)
}

// ini directives

func (ig *IniGlobals) GlobalIni(name string) *IniEntry {
	return ig.iniDirectives[name]
}
func (ig *IniGlobals) LocalIni(name string) *IniEntry {
	ini := ig.modifiedIniDirectives[name]
	if ini == nil {
		ini = ig.iniDirectives[name]
	}
	return ini
}

func (ig *IniGlobals) RegisterIniEntries(moduleNumber int, iniEntryDefs []*IniEntryDef) bool {
	for _, iniEntryDef := range iniEntryDefs {
		// check name conflict
		_, exists := ig.iniDirectives[iniEntryDef.Name()]
		if exists {
			ig.UnregisterIniEntries(moduleNumber)
			return false
		}

		// register
		p := NewIniEntry(iniEntryDef, moduleNumber)
		ig.iniDirectives[p.Name()] = p

		// apply default value
		var defaultValue, ok = ig.ConfigGetStr(p.Name())
		if ok && p.EmitOnModify(ig.ctx, defaultValue, true, IniStageStartup) {
			p = p.WithNewValue(defaultValue, true)
			ig.iniDirectives[p.Name()] = p
		} else {
			p.EmitOnModifyBy(ig.ctx, p, IniStageStartup)
		}
	}
	return true
}

func (ig *IniGlobals) UnregisterIniEntries(moduleNumber int) {
	maps.DeleteFunc(ig.iniDirectives, func(_ string, ini *IniEntry) bool {
		return ini.ModuleNumber() == moduleNumber
	})
}

func (ig *IniGlobals) AlterIni(name string, newValue string) bool {
	return ig.AlterIniEx(name, newValue, IniUser, IniStageRuntime)
}
func (ig *IniGlobals) AlterIniEx(name string, newValue string, modifyType IniModifiable, stage IniStage) bool {
	// find
	ini := ig.GlobalIni(name)
	if ini == nil || !ini.Modifiable().Match(modifyType) {
		return false
	}

	// modify
	if stage == IniStageActivate && modifyType == IniSystem {
		ini = ini.WithModifiable(IniSystem)
	}
	if ini.EmitOnModify(ig.ctx, newValue, true, stage) {
		ini = ini.WithNewValue(newValue, true)
	} else {
		return false
	}

	// save
	ig.modifiedIniDirectives[name] = ini

	return true
}
func (ig *IniGlobals) RestoreIni(name string, stage IniStage) {
	modifiedIni := ig.modifiedIniDirectives[name]
	if modifiedIni == nil || stage == IniStageRuntime && !modifiedIni.Modifiable().Match(IniUser) {
		return
	}

	origIni := ig.iniDirectives[name]
	modifiedIni.EmitOnModifyBy(ig.ctx, origIni, stage)
	delete(ig.modifiedIniDirectives, name)
}

func (ig *IniGlobals) EachIniDirective(sorted bool, h func(global *IniEntry, local *IniEntry)) {
	names := mapkit.Keys(ig.iniDirectives)
	if sorted {
		slices.Sort(names)
	}
	var global, local *IniEntry
	for _, name := range names {
		global = ig.iniDirectives[name]
		local = ig.modifiedIniDirectives[name]
		if local == nil {
			local = global
		}
		h(global, local)
	}
}

// fast functions

func (ig *IniGlobals) GetStr(name string) (string, bool) {
	ini := ig.LocalIni(name)
	if ini == nil {
		return "", false
	}

	return ini.Value(), true
}
func (ig *IniGlobals) GetStrVal(name string) string {
	ini := ig.LocalIni(name)
	if ini == nil {
		return ""
	}

	return ini.Value()
}
func (ig *IniGlobals) GetInt(name string) int {
	value := ig.GetStrVal(name)
	if value == "" {
		return 0
	}
	return ParseLong(value, 0)
}
