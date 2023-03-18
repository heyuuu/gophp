// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func INI_INT(name string) ZendLong {
	return ZendIniLong(name, b.SizeOf("name")-1, 0)
}
func INI_STR(name string) *byte {
	return ZendIniStringEx(name, b.SizeOf("name")-1, 0, nil)
}
func REGISTER_INI_ENTRIES(module_number int) int {
	return ZendRegisterIniEntries(IniEntries, module_number)
}
func UNREGISTER_INI_ENTRIES(module_number int)         { ZendUnregisterIniEntries(module_number) }
func DISPLAY_INI_ENTRIES(zend_module *ZendModuleEntry) { DisplayIniEntries(zend_module) }
func REGISTER_INI_DISPLAYER(name *byte, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	return ZendIniRegisterDisplayer(name, b.SizeOf("name")-1, displayer)
}
func REGISTER_INI_BOOLEAN(name *byte) int {
	return REGISTER_INI_DISPLAYER(name, ZendIniBooleanDisplayerCb)
}
func ZendRemoveIniEntries(el *types.Zval, arg any) int {
	var ini_entry *ZendIniEntry = (*ZendIniEntry)(el.GetPtr())
	var module_number int = *((*int)(arg))
	return ini_entry.GetModuleNumber() == module_number
}
func ZendRestoreIniEntryCb(ini_entry *ZendIniEntry, stage int) int {
	var result = false
	if ini_entry.GetModified() != 0 {
		if ini_entry.HasOnModify() {
			var __orig_bailout *JMP_BUF = EG__().GetBailout()
			var __bailout JMP_BUF
			EG__().SetBailout(&__bailout)
			if SETJMP(__bailout) == 0 {

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

				result = ini_entry.EmitOnModify(ini_entry.GetOrigValue(), stage)

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

			}
			EG__().SetBailout(__orig_bailout)
		}
		if stage == ZEND_INI_STAGE_RUNTIME && result == false {
			/* runtime failure is OK */
			return 1
		}
		if ini_entry.GetValue() != ini_entry.GetOrigValue() {
			types.ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(ini_entry.GetOrigValue())
		ini_entry.SetModifiable(ini_entry.GetOrigModifiable())
		ini_entry.SetModified(0)
		ini_entry.SetOrigValue(nil)
		ini_entry.SetOrigModifiable(0)
	}
	return 0
}
func FreeIniEntry(zv *types.Zval) {
	var entry *ZendIniEntry = (*ZendIniEntry)(zv.GetPtr())
	if entry.GetValue() != nil {
		types.ZendStringRelease(entry.GetValue())
	}
	if entry.GetOrigValue() != nil {
		types.ZendStringReleaseEx(entry.GetOrigValue(), 1)
	}
	Free(entry)
}
func ZendIniStartup() int {
	RegisteredZendIniDirectives = (*types.HashTable)(Malloc(b.SizeOf("HashTable")))
	EG__().SetIniDirectives(RegisteredZendIniDirectives)
	EG__().SetModifiedIniDirectives(nil)
	EG__().SetErrorReportingIniEntry(nil)
	ZendHashInitEx(RegisteredZendIniDirectives, 128, nil, FreeIniEntry, 1, 0)
	return types.SUCCESS
}
func ZendIniShutdown() int {
	ZendIniDtor(EG__().GetIniDirectives())
	return types.SUCCESS
}
func ZendIniDtor(ini_directives *types.HashTable) {
	ini_directives.Destroy()
	Free(ini_directives)
}
func ZendIniGlobalShutdown() int {
	RegisteredZendIniDirectives.Destroy()
	Free(RegisteredZendIniDirectives)
	return types.SUCCESS
}
func ZendIniDeactivate() int {
	if EG__().GetModifiedIniDirectives() != nil {
		var ini_entry *ZendIniEntry
		var __ht *types.HashTable = EG__().GetModifiedIniDirectives()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			ini_entry = _z.GetPtr()
			ZendRestoreIniEntryCb(ini_entry, ZEND_INI_STAGE_DEACTIVATE)
		}
		EG__().GetModifiedIniDirectives().Destroy()
		FREE_HASHTABLE(EG__().GetModifiedIniDirectives())
		EG__().SetModifiedIniDirectives(nil)
	}
	return types.SUCCESS
}
func IniKeyCompare(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	if f.GetKey() == nil && s.GetKey() == nil {
		if f.GetH() > s.GetH() {
			return -1
		} else if f.GetH() < s.GetH() {
			return 1
		}
		return 0
	} else if f.GetKey() == nil {
		return -1
	} else if s.GetKey() == nil {
		return 1
	} else {
		return ZendBinaryStrcasecmp(f.GetKey().GetVal(), f.GetKey().GetLen(), s.GetKey().GetVal(), s.GetKey().GetLen())
	}
}
func ZendIniSortEntries() {
	EG__().GetIniDirectives().SortCompatible(IniKeyCompare, 0)
}
func ZendRegisterIniEntries(iniEntryDefs []ZendIniEntryDef, moduleNumber int) int {
	var directives *types.HashTable = RegisteredZendIniDirectives
	for i := range iniEntryDefs {
		iniEntryDef := &iniEntryDefs[i]
		p := NewZendIniEntry(iniEntryDef, moduleNumber)
		if ZendHashAddPtr(directives, p.GetName(), any(p)) == nil {
			if p.GetName() != nil {
				types.ZendStringReleaseEx(p.GetName(), 1)
			}
			ZendUnregisterIniEntries(moduleNumber)
			return types.FAILURE
		}

		var defaultValue *types.Zval = ZendGetConfigurationDirective(p.GetName())
		if defaultValue != nil && p.EmitOnModify(defaultValue.GetStr(), ZEND_INI_STAGE_STARTUP) {
			p.SetValue(types.ZendNewInternedString(defaultValue.GetStr().Copy()))
		} else {
			p.SetValueStr(iniEntryDef.GetValueStr())
			p.EmitOnModify(p.GetValue(), ZEND_INI_STAGE_STARTUP)
		}
	}
	return types.SUCCESS
}
func ZendUnregisterIniEntries(module_number int) {
	ZendHashApplyWithArgument(RegisteredZendIniDirectives, ZendRemoveIniEntries, any(&module_number))
}
func ZendAlterIniEntry(name *types.ZendString, new_value *types.ZendString, modify_type int, stage int) int {
	return ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
}
func ZendAlterIniEntryChars(name *types.ZendString, value string, value_length int, modify_type int, stage int) int {
	var ret int
	var new_value *types.ZendString
	new_value = types.ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
	types.ZendStringRelease(new_value)
	return ret
}
func ZendAlterIniEntryCharsEx(
	name *types.ZendString,
	value *byte,
	value_length int,
	modify_type int,
	stage int,
	force_change int,
) int {
	var ret int
	var new_value *types.ZendString
	new_value = types.ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, force_change)
	types.ZendStringRelease(new_value)
	return ret
}
func ZendAlterIniEntryEx(name *types.ZendString, new_value *types.ZendString, modify_type int, stage int, force_change int) int {
	var ini_entry *ZendIniEntry
	var duplicate *types.ZendString
	var modifiable uint8
	var modified types.ZendBool
	if b.Assign(&ini_entry, ZendHashFindPtr(EG__().GetIniDirectives(), name)) == nil {
		return types.FAILURE
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
	if EG__().GetModifiedIniDirectives() == nil {
		ALLOC_HASHTABLE(EG__().GetModifiedIniDirectives())
		ZendHashInit(EG__().GetModifiedIniDirectives(), 8, nil, nil, 0)
	}
	if modified == 0 {
		ini_entry.SetOrigValue(ini_entry.GetValue())
		ini_entry.SetOrigModifiable(modifiable)
		ini_entry.SetModified(1)
		ZendHashAddPtr(EG__().GetModifiedIniDirectives(), ini_entry.GetName(), ini_entry)
	}
	duplicate = new_value.Copy()
	if ini_entry.EmitOnModify(duplicate, stage) {
		if modified != 0 && ini_entry.GetOrigValue() != ini_entry.GetValue() {
			types.ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(duplicate)
	} else {
		types.ZendStringRelease(duplicate)
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendRestoreIniEntry(name *types.ZendString, stage int) int {
	var ini_entry *ZendIniEntry
	if b.Assign(&ini_entry, ZendHashFindPtr(EG__().GetIniDirectives(), name)) == nil || stage == ZEND_INI_STAGE_RUNTIME && (ini_entry.GetModifiable()&ZEND_INI_USER) == 0 {
		return types.FAILURE
	}
	if EG__().GetModifiedIniDirectives() != nil {
		if ZendRestoreIniEntryCb(ini_entry, stage) == 0 {
			ZendHashDel(EG__().GetModifiedIniDirectives(), name)
		} else {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZendIniRegisterDisplayer(name *byte, name_length uint32, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(RegisteredZendIniDirectives, name, name_length)
	if ini_entry == nil {
		return types.FAILURE
	}
	ini_entry.SetDisplayer(displayer)
	return types.SUCCESS
}
func ZendIniLong(name *byte, name_length int, orig int) ZendLong {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(EG__().GetIniDirectives(), name, name_length)
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
func ZendIniDouble(name *byte, name_length int, orig int) float64 {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(EG__().GetIniDirectives(), name, name_length)
	if ini_entry != nil {
		if orig != 0 && ini_entry.GetModified() != 0 {
			return float64(b.CondF1(ini_entry.GetOrigValue() != nil, func() float64 { return ZendStrtod(ini_entry.GetOrigValue().GetVal(), nil) }, 0.0))
		} else {
			return float64(b.CondF1(ini_entry.GetValue() != nil, func() float64 { return ZendStrtod(ini_entry.GetValue().GetVal(), nil) }, 0.0))
		}
	}
	return 0.0
}
func ZendIniStringEx(name *byte, name_length int, orig int, exists *types.ZendBool) *byte {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(EG__().GetIniDirectives(), name, name_length)
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
func ZendIniString(name string, name_length int, orig int) *byte {
	var exists types.ZendBool = 1
	var return_value *byte
	return_value = ZendIniStringEx(name, name_length, orig, &exists)
	if exists == 0 {
		return nil
	} else if return_value == nil {
		return_value = ""
	}
	return return_value
}
func ZendIniGetValue(name *types.ZendString) *types.ZendString {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashFindPtr(EG__().GetIniDirectives(), name)
	if ini_entry != nil {
		if ini_entry.GetValue() != nil {
			return ini_entry.GetValue()
		} else {
			return types.ZSTR_EMPTY_ALLOC()
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
func ZendIniParseBool(str *types.ZendString) types.ZendBool {
	if str.GetLen() == 4 && strcasecmp(str.GetVal(), "true") == 0 || str.GetLen() == 3 && strcasecmp(str.GetVal(), "yes") == 0 || str.GetLen() == 2 && strcasecmp(str.GetVal(), "on") == 0 {
		return 1
	} else {
		return atoi(str.GetVal()) != 0
	}
}
func ZendIniBooleanDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value int
	var tmp_value *types.ZendString
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
	new_value *types.ZendString,
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
	new_value *types.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	*p = ZendAtol(new_value.GetVal(), new_value.GetLen())
	return types.SUCCESS
}
func OnUpdateLongGEZero(
	entry *ZendIniEntry,
	new_value *types.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p *ZendLong
	var tmp ZendLong
	var base *byte = (*byte)(mh_arg2)
	tmp = ZendAtol(new_value.GetVal(), new_value.GetLen())
	if tmp < 0 {
		return types.FAILURE
	}
	p = (*ZendLong)(base + int(mh_arg1))
	*p = tmp
	return types.SUCCESS
}
func OnUpdateReal(
	entry *ZendIniEntry,
	new_value *types.ZendString,
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
	new_value *types.ZendString,
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
	new_value *types.ZendString,
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
