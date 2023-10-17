package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func INI_INT(name string) ZendLong {
	return ZendIniLong(name, 0)
}
func INI_STR(name string) *byte {
	return ZendIniStringEx(name, 0, nil)
}
func INI_STRING(name string) string {
	value, _ := ZendIniStringExEx(name, false)
	return value
}
func REGISTER_INI_ENTRIES(module_number int) int {
	return ZendRegisterIniEntries(IniEntries, module_number)
}
func UNREGISTER_INI_ENTRIES(module_number int)     { ZendUnregisterIniEntries(module_number) }
func DISPLAY_INI_ENTRIES(zend_module *ModuleEntry) { DisplayIniEntries(zend_module) }
func ZendRestoreIniEntryCb(ini_entry *ZendIniEntry, stage int) int {
	var result = false
	if ini_entry.GetModified() != 0 {
		if ini_entry.HasOnModify() {
			faults.Try(func() {
				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */
				result = ini_entry.EmitOnModify(ini_entry.GetOrigValue(), stage)
			})
		}
		if stage == ZEND_INI_STAGE_RUNTIME && result == false {
			/* runtime failure is OK */
			return 1
		}
		if ini_entry.GetValue() != ini_entry.GetOrigValue() {
			// types.ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(ini_entry.GetOrigValue())
		ini_entry.SetModifiable(ini_entry.GetOrigModifiable())
		ini_entry.SetModified(0)
		ini_entry.SetOrigValue(nil)
		ini_entry.SetOrigModifiable(0)
	}
	return 0
}
func ZendIniStartup() int {
	EG__().InitIniDirectives()
	EG__().ModifiedIniDirectives().Destroy()
	EG__().SetErrorReportingIniEntry(nil)
	return types.SUCCESS
}
func ZendIniShutdown() int {
	EG__().IniDirectives().Destroy()
	return types.SUCCESS
}
func ZendIniDeactivate() int {
	if EG__().ModifiedIniDirectives() != nil {
		EG__().ModifiedIniDirectives().Foreach(func(_ string, ini_entry *ZendIniEntry) {
			ZendRestoreIniEntryCb(ini_entry, ZEND_INI_STAGE_DEACTIVATE)
		})
		EG__().ModifiedIniDirectives().Destroy()
	}
	return types.SUCCESS
}
func ZendIniSortEntries() {
	EG__().IniDirectives().SortByArrayKey(func(k1, k2 types.ArrayKey) bool {
		// 数字 < 字符串
		if k1.IsStrKey() && k2.IsStrKey() {
			return ascii.StrCaseCompare(k1.StrKey(), k2.StrKey()) < 0
		} else if k1.IsStrKey() {
			return false
		} else if k2.IsStrKey() {
			return true
		} else {
			return k1.IdxKey() < k2.IdxKey()
		}
	})
}
func ZendRegisterIniEntries(iniEntryDefs []ZendIniEntryDef, moduleNumber int) int {
	var directives = RegisteredZendIniDirectives
	for i := range iniEntryDefs {
		iniEntryDef := &iniEntryDefs[i]
		p := NewZendIniEntry(iniEntryDef, moduleNumber)
		if !directives.Add(p.GetName().GetStr(), p) {
			ZendUnregisterIniEntries(moduleNumber)
			return types.FAILURE
		}

		var defaultValue *types.Zval = ZendGetConfigurationDirective(p.GetName())
		if defaultValue != nil && p.EmitOnModify(defaultValue.StringEx(), ZEND_INI_STAGE_STARTUP) {
			p.SetValue(types.NewString(defaultValue.String()))
		} else {
			p.SetValueStr(iniEntryDef.GetValueStr())
			p.EmitOnModify(p.GetValue(), ZEND_INI_STAGE_STARTUP)
		}
	}
	return types.SUCCESS
}
func ZendUnregisterIniEntries(module_number int) {
	RegisteredZendIniDirectives.Filter(func(_ string, ini_entry *ZendIniEntry) bool {
		return ini_entry.GetModuleNumber() == module_number
	})
}
func ZendAlterIniEntryChars(name string, value string, modify_type int, stage int) bool {
	return ZendAlterIniEntryEx(types.NewString(name).GetStr(), types.NewString(value), modify_type, stage, 0)
}
func ZendAlterIniEntryEx(name string, new_value *types.String, modify_type int, stage int, force_change int) bool {
	var duplicate *types.String
	var modifiable uint8
	var modified bool

	var ini_entry *ZendIniEntry = EG__().IniDirectives().Get(name)
	if ini_entry == nil {
		return false
	}
	modifiable = ini_entry.GetModifiable()
	modified = ini_entry.GetModified()
	if stage == ZEND_INI_STAGE_ACTIVATE && modify_type == ZEND_INI_SYSTEM {
		ini_entry.SetModifiable(ZEND_INI_SYSTEM)
	}
	if force_change == 0 {
		if (ini_entry.GetModifiable() & modify_type) == 0 {
			return types.FAILURE
		}
	}
	if EG__().ModifiedIniDirectives() == nil {
		EG__().InitModifiedIniDirectives()
	}
	if modified == 0 {
		ini_entry.SetOrigValue(ini_entry.GetValue())
		ini_entry.SetOrigModifiable(modifiable)
		ini_entry.SetModified(1)
		EG__().ModifiedIniDirectives().Add(ini_entry.GetName().GetStr(), ini_entry)
	}
	duplicate = new_value.Copy()
	if ini_entry.EmitOnModify(duplicate, stage) {
		if modified != 0 && ini_entry.GetOrigValue() != ini_entry.GetValue() {
			// types.ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(duplicate)
	} else {
		// types.ZendStringRelease(duplicate)
		return false
	}
	return true
}
func ZendRestoreIniEntry(name string, stage int) int {
	var iniEntry = EG__().IniDirectives().Get(name)
	if iniEntry == nil || stage == ZEND_INI_STAGE_RUNTIME && (iniEntry.GetModifiable()&ZEND_INI_USER) == 0 {
		return types.FAILURE
	}
	if EG__().ModifiedIniDirectives() != nil {
		if ZendRestoreIniEntryCb(iniEntry, stage) == 0 {
			EG__().ModifiedIniDirectives().Del(name)
		} else {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZendIniLong(name string, orig int) ZendLong {
	var ini_entry *ZendIniEntry = EG__().IniDirectives().Get(name)
	if ini_entry != nil {
		if orig != 0 && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil {
				return ZEND_STRTOL(ini_entry.GetOrigValue().GetVal(), nil, 0)
			} else {
				return 0
			}
		} else {
			if ini_entry.GetValue() != nil {
				return ZEND_STRTOL(ini_entry.GetValue().GetVal(), nil, 0)
			} else {
				return 0
			}
		}
	}
	return 0
}
func ZendIniStringExEx(name string, orig bool) (value string, exists bool) {
	var iniEntry = EG__().IniDirectives().Get(name)
	if iniEntry != nil {
		exists = true
		if orig && iniEntry.GetModified() != 0 {
			if iniEntry.GetOrigValue() != nil {
				value = iniEntry.GetOrigValue().GetStr()
			} else {
				value = ""
			}
		} else {
			if iniEntry.GetValue() != nil {
				value = iniEntry.GetValue().GetStr()
			} else {
				value = ""
			}
		}
		return value, exists
	}
	return "", false
}

func getIniString(name string, orig bool) (value string, exists bool) {
	var iniEntry = EG__().IniDirectives().Get(name)
	if iniEntry != nil {
		if orig && iniEntry.GetModified() != 0 {
			if iniEntry.GetOrigValue() != nil {
				value = iniEntry.GetOrigValue().GetStr()
			} else {
				value = ""
			}
		} else {
			if iniEntry.GetValue() != nil {
				value = iniEntry.GetValue().GetStr()
			} else {
				value = ""
			}
		}
		return value, true
	}
	return "", false
}

func ZendIniStringEx(name string, orig int, exists *bool) *byte {
	value, exists_ := getIniString(name, orig != 0)
	if exists != nil {
		*exists = exists_
	}
	return b.CastStrPtr(value)
}
func ZendIniString(name string, orig int) *byte {
	value, exists_ := getIniString(name, orig != 0)
	if exists_ {
		return b.CastStrPtr(value)
	} else {
		return nil
	}
}
func ZendIniGetValueEx(name string) (string, bool) {
	return getIniString(name, false)
}
func ZendIniGetValue(name string) *types.String {
	value, exists := getIniString(name, false)
	if exists {
		return types.NewString(value)
	}
	return nil
}
func ZendIniStringParseBool(str string) bool {
	if str == "true" || str == "yes" || str == "on" {
		return true
	} else {
		return b.Atoi(str) != 0
	}
}
func ZendIniParseBool(str *types.String) bool {
	if str.GetLen() == 4 && strcasecmp(str.GetVal(), "true") == 0 || str.GetLen() == 3 && strcasecmp(str.GetVal(), "yes") == 0 || str.GetLen() == 2 && strcasecmp(str.GetVal(), "on") == 0 {
		return 1
	} else {
		return atoi(str.GetVal()) != 0
	}
}
func ZendIniBooleanDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value int
	var tmp_value *types.String
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		if ini_entry.GetOrigValue() != nil {
			tmp_value = ini_entry.GetOrigValue()
		} else {
			tmp_value = nil
		}
	} else if ini_entry.GetValue() != nil {
		tmp_value = ini_entry.GetValue()
	} else {
		tmp_value = nil
	}
	if tmp_value != nil {
		value = ZendIniParseBool(tmp_value)
	} else {
		value = 0
	}
	if value != 0 {
		ZendWrite("On")
	} else {
		ZendWrite("Off")
	}
}
func ZendIniColorDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value *byte
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		value = ini_entry.GetOrigValue().GetVal()
	} else if ini_entry.GetValue() != nil {
		value = ini_entry.GetValue().GetVal()
	} else {
		value = nil
	}
	if value != nil {
		if ZendUv.GetHtmlErrors() {
			ZendWrite(fmt.Sprintf(`<font style="color: %s">%s</font>`, value, value))
		} else {
			ZendWrite(value)
		}
	} else {
		if ZendUv.GetHtmlErrors() != 0 {
			ZendWrite(NO_VALUE_HTML)
		} else {
			ZendWrite(NO_VALUE_PLAINTEXT)
		}
	}
}
func OnUpdateBool(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *bool
	var base *byte = (*byte)(mh_arg2)
	p = (*bool)(base + int(mh_arg1))
	*p = ZendIniParseBool(new_value)
	return types.SUCCESS
}
func OnUpdateLong(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	*p = StrToLongWithUnit(new_value.GetStr())
	return types.SUCCESS
}
func OnUpdateLongGEZero(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *ZendLong
	var tmp ZendLong
	var base *byte = (*byte)(mh_arg2)
	tmp = StrToLongWithUnit(new_value.GetStr())
	if tmp < 0 {
		return types.FAILURE
	}
	p = (*ZendLong)(base + int(mh_arg1))
	*p = tmp
	return types.SUCCESS
}
func OnUpdateReal(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *float64
	var base *byte = (*byte)(mh_arg2)
	p = (*float64)(base + int(mh_arg1))
	*p = ZendStrtod(new_value.GetVal(), nil)
	return types.SUCCESS
}
func OnUpdateString(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = new_value.GetVal()
	} else {
		*p = nil
	}
	return types.SUCCESS
}
func OnUpdateStringUnempty(
	entry *ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	if new_value != nil && !(new_value.GetStr()[0]) {
		return types.FAILURE
	}
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = new_value.GetVal()
	} else {
		*p = nil
	}
	return types.SUCCESS
}
