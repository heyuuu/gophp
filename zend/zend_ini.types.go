// <<generate>>

package zend

/**
 * ZendIniEntryDef
 */
type ZendIniEntryDef struct {
	name         *byte
	on_modify    func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int
	mh_arg1      any
	mh_arg2      any
	mh_arg3      any
	value        *byte
	displayer    func(ini_entry *ZendIniEntry, type_ int)
	value_length uint32
	name_length  uint16
	modifiable   uint8
}

// func NewZendIniEntryDef(name *byte, on_modify func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int, mh_arg1 any, mh_arg2 any, mh_arg3 any, value *byte, displayer func(ini_entry *ZendIniEntry, type_ int), value_length uint32, name_length uint16, modifiable uint8) *ZendIniEntryDef {
//     return &ZendIniEntryDef{
//         name:name,
//         on_modify:on_modify,
//         mh_arg1:mh_arg1,
//         mh_arg2:mh_arg2,
//         mh_arg3:mh_arg3,
//         value:value,
//         displayer:displayer,
//         value_length:value_length,
//         name_length:name_length,
//         modifiable:modifiable,
//     }
// }
// func MakeZendIniEntryDef(name *byte, on_modify func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int, mh_arg1 any, mh_arg2 any, mh_arg3 any, value *byte, displayer func(ini_entry *ZendIniEntry, type_ int), value_length uint32, name_length uint16, modifiable uint8) ZendIniEntryDef {
//     return ZendIniEntryDef{
//         name:name,
//         on_modify:on_modify,
//         mh_arg1:mh_arg1,
//         mh_arg2:mh_arg2,
//         mh_arg3:mh_arg3,
//         value:value,
//         displayer:displayer,
//         value_length:value_length,
//         name_length:name_length,
//         modifiable:modifiable,
//     }
// }
func (this *ZendIniEntryDef) GetName() *byte { return this.name }

// func (this *ZendIniEntryDef) SetName(value *byte) { this.name = value }
func (this *ZendIniEntryDef) GetOnModify() func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	return this.on_modify
}

// func (this *ZendIniEntryDef) SetOnModify(value func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int) { this.on_modify = value }
func (this *ZendIniEntryDef) GetMhArg1() any { return this.mh_arg1 }

// func (this *ZendIniEntryDef) SetMhArg1(value any) { this.mh_arg1 = value }
func (this *ZendIniEntryDef) GetMhArg2() any { return this.mh_arg2 }

// func (this *ZendIniEntryDef) SetMhArg2(value any) { this.mh_arg2 = value }
func (this *ZendIniEntryDef) GetMhArg3() any { return this.mh_arg3 }

// func (this *ZendIniEntryDef) SetMhArg3(value any) { this.mh_arg3 = value }
func (this *ZendIniEntryDef) GetValue() *byte { return this.value }

// func (this *ZendIniEntryDef) SetValue(value *byte) { this.value = value }
func (this *ZendIniEntryDef) GetDisplayer() func(ini_entry *ZendIniEntry, type_ int) {
	return this.displayer
}

// func (this *ZendIniEntryDef) SetDisplayer(value func(ini_entry *ZendIniEntry, type_ int)) { this.displayer = value }
func (this *ZendIniEntryDef) GetValueLength() uint32 { return this.value_length }

// func (this *ZendIniEntryDef) SetValueLength(value uint32) { this.value_length = value }
func (this *ZendIniEntryDef) GetNameLength() uint16 { return this.name_length }

// func (this *ZendIniEntryDef) SetNameLength(value uint16) { this.name_length = value }
func (this *ZendIniEntryDef) GetModifiable() uint8 { return this.modifiable }

// func (this *ZendIniEntryDef) SetModifiable(value uint8) { this.modifiable = value }

/**
 * ZendIniEntry
 */
type ZendIniEntry struct {
	name            *ZendString
	on_modify       func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int
	mh_arg1         any
	mh_arg2         any
	mh_arg3         any
	value           *ZendString
	orig_value      *ZendString
	displayer       func(ini_entry *ZendIniEntry, type_ int)
	module_number   int
	modifiable      uint8
	orig_modifiable uint8
	modified        uint8
}

// func NewZendIniEntry(name *ZendString, on_modify func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int, mh_arg1 any, mh_arg2 any, mh_arg3 any, value *ZendString, orig_value *ZendString, displayer func(ini_entry *ZendIniEntry, type_ int), module_number int, modifiable uint8, orig_modifiable uint8, modified uint8) *ZendIniEntry {
//     return &ZendIniEntry{
//         name:name,
//         on_modify:on_modify,
//         mh_arg1:mh_arg1,
//         mh_arg2:mh_arg2,
//         mh_arg3:mh_arg3,
//         value:value,
//         orig_value:orig_value,
//         displayer:displayer,
//         module_number:module_number,
//         modifiable:modifiable,
//         orig_modifiable:orig_modifiable,
//         modified:modified,
//     }
// }
// func MakeZendIniEntry(name *ZendString, on_modify func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int, mh_arg1 any, mh_arg2 any, mh_arg3 any, value *ZendString, orig_value *ZendString, displayer func(ini_entry *ZendIniEntry, type_ int), module_number int, modifiable uint8, orig_modifiable uint8, modified uint8) ZendIniEntry {
//     return ZendIniEntry{
//         name:name,
//         on_modify:on_modify,
//         mh_arg1:mh_arg1,
//         mh_arg2:mh_arg2,
//         mh_arg3:mh_arg3,
//         value:value,
//         orig_value:orig_value,
//         displayer:displayer,
//         module_number:module_number,
//         modifiable:modifiable,
//         orig_modifiable:orig_modifiable,
//         modified:modified,
//     }
// }
func (this *ZendIniEntry) GetName() *ZendString      { return this.name }
func (this *ZendIniEntry) SetName(value *ZendString) { this.name = value }
func (this *ZendIniEntry) GetOnModify() func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	return this.on_modify
}
func (this *ZendIniEntry) SetOnModify(value func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int) {
	this.on_modify = value
}
func (this *ZendIniEntry) GetMhArg1() any                 { return this.mh_arg1 }
func (this *ZendIniEntry) SetMhArg1(value any)            { this.mh_arg1 = value }
func (this *ZendIniEntry) GetMhArg2() any                 { return this.mh_arg2 }
func (this *ZendIniEntry) SetMhArg2(value any)            { this.mh_arg2 = value }
func (this *ZendIniEntry) GetMhArg3() any                 { return this.mh_arg3 }
func (this *ZendIniEntry) SetMhArg3(value any)            { this.mh_arg3 = value }
func (this *ZendIniEntry) GetValue() *ZendString          { return this.value }
func (this *ZendIniEntry) SetValue(value *ZendString)     { this.value = value }
func (this *ZendIniEntry) GetOrigValue() *ZendString      { return this.orig_value }
func (this *ZendIniEntry) SetOrigValue(value *ZendString) { this.orig_value = value }
func (this *ZendIniEntry) GetDisplayer() func(ini_entry *ZendIniEntry, type_ int) {
	return this.displayer
}
func (this *ZendIniEntry) SetDisplayer(value func(ini_entry *ZendIniEntry, type_ int)) {
	this.displayer = value
}
func (this *ZendIniEntry) GetModuleNumber() int          { return this.module_number }
func (this *ZendIniEntry) SetModuleNumber(value int)     { this.module_number = value }
func (this *ZendIniEntry) GetModifiable() uint8          { return this.modifiable }
func (this *ZendIniEntry) SetModifiable(value uint8)     { this.modifiable = value }
func (this *ZendIniEntry) GetOrigModifiable() uint8      { return this.orig_modifiable }
func (this *ZendIniEntry) SetOrigModifiable(value uint8) { this.orig_modifiable = value }
func (this *ZendIniEntry) GetModified() uint8            { return this.modified }
func (this *ZendIniEntry) SetModified(value uint8)       { this.modified = value }

/**
 * ZendIniParserParam
 */
type ZendIniParserParam struct {
	ini_parser_cb ZendIniParserCbT
	arg           any
}

// func NewZendIniParserParam(ini_parser_cb ZendIniParserCbT, arg any) *ZendIniParserParam {
//     return &ZendIniParserParam{
//         ini_parser_cb:ini_parser_cb,
//         arg:arg,
//     }
// }
// func MakeZendIniParserParam(ini_parser_cb ZendIniParserCbT, arg any) ZendIniParserParam {
//     return ZendIniParserParam{
//         ini_parser_cb:ini_parser_cb,
//         arg:arg,
//     }
// }
func (this *ZendIniParserParam) GetIniParserCb() ZendIniParserCbT      { return this.ini_parser_cb }
func (this *ZendIniParserParam) SetIniParserCb(value ZendIniParserCbT) { this.ini_parser_cb = value }
func (this *ZendIniParserParam) GetArg() any                           { return this.arg }
func (this *ZendIniParserParam) SetArg(value any)                      { this.arg = value }
