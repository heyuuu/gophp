// <<generate>>

package zend

/**
 * ZendIniEntryDef
 */
type ZendIniEntryModifierEx = func(entry *ZendIniEntry, new_value *string, stage int) bool
type ZendIniEntryModifier = func(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int
type ZendIniEntryDisplayer = func(ini_entry *ZendIniEntry, type_ int)
type ZendIniEntryDef struct {
	name       string
	modifiable uint8
	value      *string
	onModify   ZendIniEntryModifierEx
	displayer  ZendIniEntryDisplayer
}

func NewZendIniEntryDef(name string, modifiable uint8) *ZendIniEntryDef {
	return &ZendIniEntryDef{name: name, modifiable: modifiable}
}
func (this *ZendIniEntryDef) Value(value string) *ZendIniEntryDef {
	this.value = &value
	return this
}
func (this *ZendIniEntryDef) OnModify(onModify ZendIniEntryModifierEx) *ZendIniEntryDef {
	this.onModify = onModify
	return this
}
func (this *ZendIniEntryDef) OnModifyArgs(onModify ZendIniEntryModifier, args ...any) *ZendIniEntryDef {
	args = append(args, nil, nil, nil)
	this.onModify = _buildOnModify(onModify, args[0], args[1], args[2])
	return this
}
func (this *ZendIniEntryDef) Displayer(displayer ZendIniEntryDisplayer) *ZendIniEntryDef {
	this.displayer = displayer
	return this
}

func _buildOnModify(onModify ZendIniEntryModifier, mh_arg1 any, mh_arg2 any, mh_arg3 any) ZendIniEntryModifierEx {
	if onModify == nil {
		return nil
	}

	return func(entry *ZendIniEntry, new_value *string, stage int) bool {
		result := onModify(entry, castZendStringPtr(new_value), mh_arg1, mh_arg2, mh_arg3, stage)
		return result == SUCCESS
	}
}

func (this *ZendIniEntryDef) GetValueStr() *string { return this.value }
func (this *ZendIniEntryDef) GetName() *byte       { return this.name }

/**
 * ZendIniEntry
 */
type ZendIniEntry struct {
	name            *string
	onModify        func(entry *ZendIniEntry, newValue *string, stage int) bool
	value           *string
	orig_value      *string
	displayer       func(ini_entry *ZendIniEntry, type_ int)
	module_number   int
	modifiable      uint8
	orig_modifiable uint8
	modified        uint8
}

func NewZendIniEntry(def *ZendIniEntryDef, module_number int) *ZendIniEntry {
	entry := &ZendIniEntry{
		name:            def.name,
		onModify:        def.onModify,
		value:           nil,
		orig_value:      nil,
		displayer:       def.displayer,
		modifiable:      def.modifiable,
		orig_modifiable: 0,
		modified:        0,
		module_number:   module_number,
	}
	return entry
}

func (this *ZendIniEntry) HasOnModify() bool {
	return this.onModify != nil
}

func (this *ZendIniEntry) EmitOnModify(new_value *ZendString, stage int) bool {
	if this.onModify == nil {
		return true
	}

	return this.onModify(this, castStrPtr(new_value), stage)
}

func (this *ZendIniEntry) EmitOnModifyCurrValue(stage int) bool {
	return this.EmitOnModify(this.GetValue(), stage)
}

/**
 * generate
 */
func (this *ZendIniEntry) SetValueStr(value *string) { this.value = value }

func (this *ZendIniEntry) GetName() *ZendString           { return castZendStringPtr(this.name) }
func (this *ZendIniEntry) GetValue() *ZendString          { return castZendStringPtr(this.value) }
func (this *ZendIniEntry) GetOrigValue() *ZendString      { return castZendStringPtr(this.orig_value) }
func (this *ZendIniEntry) SetValue(value *ZendString)     { this.value = castStrPtr(value) }
func (this *ZendIniEntry) SetOrigValue(value *ZendString) { this.orig_value = castStrPtr(value) }

func (this *ZendIniEntry) GetDisplayer() func(ini_entry *ZendIniEntry, type_ int) {
	return this.displayer
}
func (this *ZendIniEntry) SetDisplayer(value func(ini_entry *ZendIniEntry, type_ int)) {
	this.displayer = value
}

func (this *ZendIniEntry) GetModuleNumber() int          { return this.module_number }
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

func (this *ZendIniParserParam) GetIniParserCb() ZendIniParserCbT      { return this.ini_parser_cb }
func (this *ZendIniParserParam) SetIniParserCb(value ZendIniParserCbT) { this.ini_parser_cb = value }
func (this *ZendIniParserParam) GetArg() any                           { return this.arg }
func (this *ZendIniParserParam) SetArg(value any)                      { this.arg = value }
