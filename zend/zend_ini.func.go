package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func INI_INT(name string) ZendLong {
	return ZendIniLong(name, 0)
}
func INI_STR(name string) *byte {
	return ZendIniStringEx(name, 0, nil)
}
func REGISTER_INI_ENTRIES(module_number int) int {
	return ZendRegisterIniEntries(IniEntries, module_number)
}
func UNREGISTER_INI_ENTRIES(module_number int)     { ZendUnregisterIniEntries(module_number) }
func DISPLAY_INI_ENTRIES(zend_module *ModuleEntry) { DisplayIniEntries(zend_module) }
func REGISTER_INI_DISPLAYER(name *byte, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	return ZendIniRegisterDisplayer(name, b.SizeOf("name")-1, displayer)
}
func REGISTER_INI_BOOLEAN(name *byte) int {
	return REGISTER_INI_DISPLAYER(name, ZendIniBooleanDisplayerCb)
}
func ZendRemoveIniEntries(el *types.Zval, arg any) int {
	var ini_entry *ZendIniEntry = (*ZendIniEntry)(el.Ptr())
	var module_number int = *((*int)(arg))
	return ini_entry.GetModuleNumber() == module_number
}
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
			return k1.IndexKey() < k2.IndexKey()
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
		if defaultValue != nil && p.EmitOnModify(defaultValue.String(), ZEND_INI_STAGE_STARTUP) {
			p.SetValue(types.NewString(defaultValue.StringVal()))
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
	var modified types.ZendBool

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
func ZendRestoreIniEntry(name *types.String, stage int) int {
	var ini_entry = EG__().IniDirectives().Get(name.GetStr())
	if ini_entry == nil || stage == ZEND_INI_STAGE_RUNTIME && (ini_entry.GetModifiable()&ZEND_INI_USER) == 0 {
		return types.FAILURE
	}
	if EG__().ModifiedIniDirectives() != nil {
		if ZendRestoreIniEntryCb(ini_entry, stage) == 0 {
			EG__().ModifiedIniDirectives().Del(name.GetStr())
		} else {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZendIniRegisterDisplayer(name *byte, name_length uint32, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	var ini_entry *ZendIniEntry = RegisteredZendIniDirectives.Get(b.CastStr(name, name_length))
	if ini_entry == nil {
		return types.FAILURE
	}
	ini_entry.SetDisplayer(displayer)
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
func ZendIniStringEx(name string, orig int, exists *types.ZendBool) *byte {
	var ini_entry *ZendIniEntry = EG__().IniDirectives().Get(name)
	if ini_entry != nil {
		if exists != nil {
			*exists = 1
		}
		if orig != 0 && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil {
				return ini_entry.GetOrigValue().GetVal()
			} else {
				return nil
			}
		} else {
			if ini_entry.GetValue() != nil {
				return ini_entry.GetValue().GetVal()
			} else {
				return nil
			}
		}
	} else {
		if exists != nil {
			*exists = 0
		}
		return nil
	}
}
func ZendIniString(name string, orig int) *byte {
	var exists types.ZendBool = 1
	var return_value *byte
	return_value = ZendIniStringEx(name, orig, &exists)
	if exists == 0 {
		return nil
	} else if return_value == nil {
		return_value = ""
	}
	return return_value
}
func ZendIniGetValueEx(name string) (string, bool) {
	var ini_entry *ZendIniEntry = EG__().IniDirectives().Get(name)
	if ini_entry != nil {
		if ini_entry.GetValue() != nil {
			return ini_entry.GetValue().GetStr(), true
		} else {
			return "", true
		}
	} else {
		return "", false
	}
}
func ZendIniGetValue(name string) *types.String {
	var ini_entry *ZendIniEntry = EG__().IniDirectives().Get(name)
	if ini_entry != nil {
		if ini_entry.GetValue() != nil {
			return ini_entry.GetValue()
		} else {
			return types.NewString("")
		}
	} else {
		return nil
	}
}
func ZendIniStringParseBool(str string) bool {
	if str == "true" || str == "yes" || str == "on" {
		return true
	} else {
		return b.Atoi(str) != 0
	}
}
func ZendIniParseBool(str *types.String) types.ZendBool {
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
		ZEND_PUTS("On")
	} else {
		ZEND_PUTS("Off")
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
		if ZendUv.GetHtmlErrors() != 0 {
			ZendPrintf("<font style=\"color: %s\">%s</font>", value, value)
		} else {
			ZEND_PUTS(value)
		}
	} else {
		if ZendUv.GetHtmlErrors() != 0 {
			ZEND_PUTS(NO_VALUE_HTML)
		} else {
			ZEND_PUTS(NO_VALUE_PLAINTEXT)
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
	var p *types.ZendBool
	var base *byte = (*byte)(mh_arg2)
	p = (*types.ZendBool)(base + int(mh_arg1))
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
	if new_value != nil && !(new_value.GetVal()[0]) {
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
