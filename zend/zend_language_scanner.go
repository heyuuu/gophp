package zend

import b "sik/builtin"

/**
 * ZendLexState
 */
type ZendLexState struct {
	len_              uint // LANG_SCNG__().yy_leng
	start             uint // LANG_SCNG__().yy_start *byte
	text              uint // LANG_SCNG__().yy_text *byte
	cursor            uint // LANG_SCNG__().yy_cursor *byte
	marker            uint // LANG_SCNG__().yy_marker *byte
	limit             uint // LANG_SCNG__().yy_limit *byte
	state             int  // LANG_SCNG__().yy_state
	stateStack        b.Stack[int]
	heredocLabelStack b.Stack[*ZendHeredocLabel]

	in                   *ZendFileHandle
	lineno               uint32
	filename             *ZendString
	script_org           *uint8
	script_org_size      int
	script_filtered      *uint8
	script_filtered_size int
	input_filter         func(string) string // LANG_SCNG__().input_filter  函数参数类型有差异
	output_filter        func(string) string // LANG_SCNG__().output_filter 函数参数类型有差异
	script_encoding      *ZendEncoding
	on_event             func(event ZendPhpScannerEvent, token int, line int, context any)
	on_event_context     any
	ast                  *ZendAst
	ast_arena            *ZendArena
}

/**
 * ZendHeredocLabel
 */
type ZendHeredocLabel struct {
	label                 string
	indentation           uint
	indentationUsesSpaces bool
}

func NewHeredocLabel(label string) *ZendHeredocLabel {
	return &ZendHeredocLabel{
		label:       label,
		indentation: 0,
	}
}

func (l *ZendHeredocLabel) Label() string { return l.label }
func (l *ZendHeredocLabel) Length() uint  { return uint(len(l.label)) }
func (l *ZendHeredocLabel) Copy() *ZendHeredocLabel {
	newLabel := *l
	return &newLabel
}
