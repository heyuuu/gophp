// <<generate>>

package zend

import (
	b "sik/builtin"
)

func INI_INT(name string) ZendLong {
	return ZendIniLong(name, b.SizeOf("name")-1, 0)
}
func INI_FLT(name *byte) float64 {
	return ZendIniDouble(name, b.SizeOf("name")-1, 0)
}
func INI_STR(name string) *byte {
	return ZendIniStringEx(name, b.SizeOf("name")-1, 0, nil)
}
func INI_BOOL(name *byte) ZendBool { return ZendBool(INI_INT(name)) }
func INI_ORIG_INT(name *byte) ZendLong {
	return ZendIniLong(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_FLT(name *byte) float64 {
	return ZendIniDouble(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_STR(name *byte) *byte {
	return ZendIniString(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_BOOL(name *byte) ZendBool { return ZendBool(INI_ORIG_INT(name)) }
func REGISTER_INI_ENTRIES() int {
	return ZendRegisterIniEntries(IniEntries, module_number)
}
func UNREGISTER_INI_ENTRIES() { ZendUnregisterIniEntries(module_number) }
func DISPLAY_INI_ENTRIES()    { DisplayIniEntries(zend_module) }
func REGISTER_INI_DISPLAYER(name *byte, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	return ZendIniRegisterDisplayer(name, b.SizeOf("name")-1, displayer)
}
func REGISTER_INI_BOOLEAN(name *byte) int {
	return REGISTER_INI_DISPLAYER(name, ZendIniBooleanDisplayerCb)
}
func ZendRemoveIniEntries(el *Zval, arg any) int {
	var ini_entry *ZendIniEntry = (*ZendIniEntry)(el.GetPtr())
	var module_number int = *((*int)(arg))
	return ini_entry.GetModuleNumber() == module_number
}
func ZendRestoreIniEntryCb(ini_entry *ZendIniEntry, stage int) int {
	var result int = FAILURE
	if ini_entry.GetModified() != 0 {
		if ini_entry.GetOnModify() != nil {
			var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
			var __bailout JMP_BUF
			ExecutorGlobals.SetBailout(&__bailout)
			if SETJMP(__bailout) == 0 {

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

				result = ini_entry.GetOnModify()(ini_entry, ini_entry.GetOrigValue(), ini_entry.GetMhArg1(), ini_entry.GetMhArg2(), ini_entry.GetMhArg3(), stage)

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

			}
			ExecutorGlobals.SetBailout(__orig_bailout)
		}
		if stage == ZEND_INI_STAGE_RUNTIME && result == FAILURE {

			/* runtime failure is OK */

			return 1

			/* runtime failure is OK */

		}
		if ini_entry.GetValue() != ini_entry.GetOrigValue() {
			ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(ini_entry.GetOrigValue())
		ini_entry.SetModifiable(ini_entry.GetOrigModifiable())
		ini_entry.SetModified(0)
		ini_entry.SetOrigValue(nil)
		ini_entry.SetOrigModifiable(0)
	}
	return 0
}
func FreeIniEntry(zv *Zval) {
	var entry *ZendIniEntry = (*ZendIniEntry)(zv.GetPtr())
	ZendStringReleaseEx(entry.GetName(), 1)
	if entry.GetValue() != nil {
		ZendStringRelease(entry.GetValue())
	}
	if entry.GetOrigValue() != nil {
		ZendStringReleaseEx(entry.GetOrigValue(), 1)
	}
	Free(entry)
}
func ZendIniStartup() int {
	RegisteredZendIniDirectives = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	ExecutorGlobals.SetIniDirectives(RegisteredZendIniDirectives)
	ExecutorGlobals.SetModifiedIniDirectives(nil)
	ExecutorGlobals.SetErrorReportingIniEntry(nil)
	RegisteredZendIniDirectives.InitEx(128, nil, FreeIniEntry, 1, 0)
	return SUCCESS
}
func ZendIniShutdown() int {
	ZendIniDtor(ExecutorGlobals.GetIniDirectives())
	return SUCCESS
}
func ZendIniDtor(ini_directives *HashTable) {
	ini_directives.Destroy()
	Free(ini_directives)
}
func ZendIniGlobalShutdown() int {
	RegisteredZendIniDirectives.Destroy()
	Free(RegisteredZendIniDirectives)
	return SUCCESS
}
func ZendIniDeactivate() int {
	if ExecutorGlobals.GetModifiedIniDirectives() != nil {
		var ini_entry *ZendIniEntry
		for {
			var __ht *HashTable = ExecutorGlobals.GetModifiedIniDirectives()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				ini_entry = _z.GetPtr()
				ZendRestoreIniEntryCb(ini_entry, ZEND_INI_STAGE_DEACTIVATE)
			}
			break
		}
		ExecutorGlobals.GetModifiedIniDirectives().Destroy()
		FREE_HASHTABLE(ExecutorGlobals.GetModifiedIniDirectives())
		ExecutorGlobals.SetModifiedIniDirectives(nil)
	}
	return SUCCESS
}
func IniKeyCompare(a any, b any) int {
	var f *Bucket
	var s *Bucket
	f = (*Bucket)(a)
	s = (*Bucket)(b)
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
	ExecutorGlobals.GetIniDirectives().Sort(IniKeyCompare, 0)
}
func ZendRegisterIniEntries(ini_entry *ZendIniEntryDef, module_number int) int {
	var p *ZendIniEntry
	var default_value *Zval
	var directives *HashTable = RegisteredZendIniDirectives
	for ini_entry.GetName() != nil {
		p = Pemalloc(b.SizeOf("zend_ini_entry"), 1)
		p.SetName(ZendStringInitInterned(ini_entry.GetName(), ini_entry.GetNameLength(), 1))
		p.SetOnModify(ini_entry.GetOnModify())
		p.SetMhArg1(ini_entry.GetMhArg1())
		p.SetMhArg2(ini_entry.GetMhArg2())
		p.SetMhArg3(ini_entry.GetMhArg3())
		p.SetValue(nil)
		p.SetOrigValue(nil)
		p.SetDisplayer(ini_entry.GetDisplayer())
		p.SetModifiable(ini_entry.GetModifiable())
		p.SetOrigModifiable(0)
		p.SetModified(0)
		p.SetModuleNumber(module_number)
		if directives.AddPtr(p.GetName(), any(p)) == nil {
			if p.GetName() != nil {
				ZendStringReleaseEx(p.GetName(), 1)
			}
			ZendUnregisterIniEntries(module_number)
			return FAILURE
		}
		if b.Assign(&default_value, ZendGetConfigurationDirective(p.GetName())) != nil && (p.GetOnModify() == nil || p.GetOnModify()(p, default_value.GetStr(), p.GetMhArg1(), p.GetMhArg2(), p.GetMhArg3(), ZEND_INI_STAGE_STARTUP) == SUCCESS) {
			p.SetValue(ZendNewInternedString(ZendStringCopy(default_value.GetStr())))
		} else {
			if ini_entry.GetValue() != nil {
				p.SetValue(ZendStringInitInterned(ini_entry.GetValue(), ini_entry.GetValueLength(), 1))
			} else {
				p.SetValue(nil)
			}
			if p.GetOnModify() != nil {
				p.GetOnModify()(p, p.GetValue(), p.GetMhArg1(), p.GetMhArg2(), p.GetMhArg3(), ZEND_INI_STAGE_STARTUP)
			}
		}
		ini_entry++
	}
	return SUCCESS
}
func ZendUnregisterIniEntries(module_number int) {
	RegisteredZendIniDirectives.ApplyWithArgument(ZendRemoveIniEntries, any(&module_number))
}
func ZendAlterIniEntry(name *ZendString, new_value *ZendString, modify_type int, stage int) int {
	return ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
}
func ZendAlterIniEntryChars(name *ZendString, value string, value_length int, modify_type int, stage int) int {
	var ret int
	var new_value *ZendString
	new_value = ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
	ZendStringRelease(new_value)
	return ret
}
func ZendAlterIniEntryCharsEx(name *ZendString, value *byte, value_length int, modify_type int, stage int, force_change int) int {
	var ret int
	var new_value *ZendString
	new_value = ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, force_change)
	ZendStringRelease(new_value)
	return ret
}
func ZendAlterIniEntryEx(name *ZendString, new_value *ZendString, modify_type int, stage int, force_change int) int {
	var ini_entry *ZendIniEntry
	var duplicate *ZendString
	var modifiable uint8
	var modified ZendBool
	if b.Assign(&ini_entry, ExecutorGlobals.GetIniDirectives().FindPtr(name)) == nil {
		return FAILURE
	}
	modifiable = ini_entry.GetModifiable()
	modified = ini_entry.GetModified()
	if stage == ZEND_INI_STAGE_ACTIVATE && modify_type == ZEND_INI_SYSTEM {
		ini_entry.SetModifiable(ZEND_INI_SYSTEM)
	}
	if force_change == 0 {
		if (ini_entry.GetModifiable() & modify_type) == 0 {
			return FAILURE
		}
	}
	if ExecutorGlobals.GetModifiedIniDirectives() == nil {
		ALLOC_HASHTABLE(ExecutorGlobals.GetModifiedIniDirectives())
		ExecutorGlobals.GetModifiedIniDirectives().Init(8, nil, nil, 0)
	}
	if modified == 0 {
		ini_entry.SetOrigValue(ini_entry.GetValue())
		ini_entry.SetOrigModifiable(modifiable)
		ini_entry.SetModified(1)
		ExecutorGlobals.GetModifiedIniDirectives().AddPtr(ini_entry.GetName(), ini_entry)
	}
	duplicate = ZendStringCopy(new_value)
	if ini_entry.GetOnModify() == nil || ini_entry.GetOnModify()(ini_entry, duplicate, ini_entry.GetMhArg1(), ini_entry.GetMhArg2(), ini_entry.GetMhArg3(), stage) == SUCCESS {
		if modified != 0 && ini_entry.GetOrigValue() != ini_entry.GetValue() {
			ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(duplicate)
	} else {
		ZendStringRelease(duplicate)
		return FAILURE
	}
	return SUCCESS
}
func ZendRestoreIniEntry(name *ZendString, stage int) int {
	var ini_entry *ZendIniEntry
	if b.Assign(&ini_entry, ExecutorGlobals.GetIniDirectives().FindPtr(name)) == nil || stage == ZEND_INI_STAGE_RUNTIME && (ini_entry.GetModifiable()&ZEND_INI_USER) == 0 {
		return FAILURE
	}
	if ExecutorGlobals.GetModifiedIniDirectives() != nil {
		if ZendRestoreIniEntryCb(ini_entry, stage) == 0 {
			ExecutorGlobals.GetModifiedIniDirectives().Del(name)
		} else {
			return FAILURE
		}
	}
	return SUCCESS
}
func ZendIniRegisterDisplayer(name *byte, name_length uint32, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	var ini_entry *ZendIniEntry
	ini_entry = RegisteredZendIniDirectives.StrFindPtr(name, name_length)
	if ini_entry == nil {
		return FAILURE
	}
	ini_entry.SetDisplayer(displayer)
	return SUCCESS
}
func ZendIniLong(name *byte, name_length int, orig int) ZendLong {
	var ini_entry *ZendIniEntry
	ini_entry = ExecutorGlobals.GetIniDirectives().StrFindPtr(name, name_length)
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
	ini_entry = ExecutorGlobals.GetIniDirectives().StrFindPtr(name, name_length)
	if ini_entry != nil {
		if orig != 0 && ini_entry.GetModified() != 0 {
			return float64(b.CondF1(ini_entry.GetOrigValue() != nil, func() float64 { return ZendStrtod(ini_entry.GetOrigValue().GetVal(), nil) }, 0.0))
		} else {
			return float64(b.CondF1(ini_entry.GetValue() != nil, func() float64 { return ZendStrtod(ini_entry.GetValue().GetVal(), nil) }, 0.0))
		}
	}
	return 0.0
}
func ZendIniStringEx(name *byte, name_length int, orig int, exists *ZendBool) *byte {
	var ini_entry *ZendIniEntry
	ini_entry = ExecutorGlobals.GetIniDirectives().StrFindPtr(name, name_length)
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
	var exists ZendBool = 1
	var return_value *byte
	return_value = ZendIniStringEx(name, name_length, orig, &exists)
	if exists == 0 {
		return nil
	} else if return_value == nil {
		return_value = ""
	}
	return return_value
}
func ZendIniGetValue(name *ZendString) *ZendString {
	var ini_entry *ZendIniEntry
	ini_entry = ExecutorGlobals.GetIniDirectives().FindPtr(name)
	if ini_entry != nil {
		if ini_entry.GetValue() != nil {
			return ini_entry.GetValue()
		} else {
			return ZSTR_EMPTY_ALLOC()
		}
	} else {
		return nil
	}
}
func ZendIniParseBool(str *ZendString) ZendBool {
	if str.GetLen() == 4 && strcasecmp(str.GetVal(), "true") == 0 || str.GetLen() == 3 && strcasecmp(str.GetVal(), "yes") == 0 || str.GetLen() == 2 && strcasecmp(str.GetVal(), "on") == 0 {
		return 1
	} else {
		return atoi(str.GetVal()) != 0
	}
}
func ZendIniBooleanDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value int
	var tmp_value *ZendString
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
func DisplayLinkNumbers(ini_entry *ZendIniEntry, type_ int) {
	var value *byte
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		value = ini_entry.GetOrigValue().GetVal()
	} else if ini_entry.GetValue() != nil {
		value = ini_entry.GetValue().GetVal()
	} else {
		value = nil
	}
	if value != nil {
		if atoi(value) == -1 {
			ZEND_PUTS("Unlimited")
		} else {
			ZendPrintf("%s", value)
		}
	}
}
func OnUpdateBool(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendBool
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendBool)(base + int(mh_arg1))
	*p = ZendIniParseBool(new_value)
	return SUCCESS
}
func OnUpdateLong(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	*p = ZendAtol(new_value.GetVal(), new_value.GetLen())
	return SUCCESS
}
func OnUpdateLongGEZero(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var tmp ZendLong
	var base *byte = (*byte)(mh_arg2)
	tmp = ZendAtol(new_value.GetVal(), new_value.GetLen())
	if tmp < 0 {
		return FAILURE
	}
	p = (*ZendLong)(base + int(mh_arg1))
	*p = tmp
	return SUCCESS
}
func OnUpdateReal(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *float64
	var base *byte = (*byte)(mh_arg2)
	p = (*float64)(base + int(mh_arg1))
	*p = ZendStrtod(new_value.GetVal(), nil)
	return SUCCESS
}
func OnUpdateString(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = new_value.GetVal()
	} else {
		*p = nil
	}
	return SUCCESS
}
func OnUpdateStringUnempty(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	if new_value != nil && !(new_value.GetVal()[0]) {
		return FAILURE
	}
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = new_value.GetVal()
	} else {
		*p = nil
	}
	return SUCCESS
}
