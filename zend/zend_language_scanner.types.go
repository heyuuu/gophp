// <<generate>>

package zend

/**
 * ZendLexState
 */
type ZendLexState struct {
	yy_leng              uint
	yy_start             *uint8
	yy_text              *uint8
	yy_cursor            *uint8
	yy_marker            *uint8
	yy_limit             *uint8
	yy_state             int
	state_stack          ZendStack
	heredoc_label_stack  ZendPtrStack
	in                   *ZendFileHandle
	lineno               uint32
	filename             *ZendString
	script_org           *uint8
	script_org_size      int
	script_filtered      *uint8
	script_filtered_size int
	input_filter         ZendEncodingFilter
	output_filter        ZendEncodingFilter
	script_encoding      *ZendEncoding
	on_event             func(event ZendPhpScannerEvent, token int, line int, context any)
	on_event_context     any
	ast                  *ZendAst
	ast_arena            *ZendArena
}

func (this ZendLexState) GetYyLeng() uint                           { return this.yy_leng }
func (this *ZendLexState) SetYyLeng(value uint)                     { this.yy_leng = value }
func (this ZendLexState) GetYyStart() *uint8                        { return this.yy_start }
func (this *ZendLexState) SetYyStart(value *uint8)                  { this.yy_start = value }
func (this ZendLexState) GetYyText() *uint8                         { return this.yy_text }
func (this *ZendLexState) SetYyText(value *uint8)                   { this.yy_text = value }
func (this ZendLexState) GetYyCursor() *uint8                       { return this.yy_cursor }
func (this *ZendLexState) SetYyCursor(value *uint8)                 { this.yy_cursor = value }
func (this ZendLexState) GetYyMarker() *uint8                       { return this.yy_marker }
func (this *ZendLexState) SetYyMarker(value *uint8)                 { this.yy_marker = value }
func (this ZendLexState) GetYyLimit() *uint8                        { return this.yy_limit }
func (this *ZendLexState) SetYyLimit(value *uint8)                  { this.yy_limit = value }
func (this ZendLexState) GetYyState() int                           { return this.yy_state }
func (this *ZendLexState) SetYyState(value int)                     { this.yy_state = value }
func (this ZendLexState) GetStateStack() ZendStack                  { return this.state_stack }
func (this *ZendLexState) SetStateStack(value ZendStack)            { this.state_stack = value }
func (this ZendLexState) GetHeredocLabelStack() ZendPtrStack        { return this.heredoc_label_stack }
func (this *ZendLexState) SetHeredocLabelStack(value ZendPtrStack)  { this.heredoc_label_stack = value }
func (this ZendLexState) GetIn() *ZendFileHandle                    { return this.in }
func (this *ZendLexState) SetIn(value *ZendFileHandle)              { this.in = value }
func (this ZendLexState) GetLineno() uint32                         { return this.lineno }
func (this *ZendLexState) SetLineno(value uint32)                   { this.lineno = value }
func (this ZendLexState) GetFilename() *ZendString                  { return this.filename }
func (this *ZendLexState) SetFilename(value *ZendString)            { this.filename = value }
func (this ZendLexState) GetScriptOrg() *uint8                      { return this.script_org }
func (this *ZendLexState) SetScriptOrg(value *uint8)                { this.script_org = value }
func (this ZendLexState) GetScriptOrgSize() int                     { return this.script_org_size }
func (this *ZendLexState) SetScriptOrgSize(value int)               { this.script_org_size = value }
func (this ZendLexState) GetScriptFiltered() *uint8                 { return this.script_filtered }
func (this *ZendLexState) SetScriptFiltered(value *uint8)           { this.script_filtered = value }
func (this ZendLexState) GetScriptFilteredSize() int                { return this.script_filtered_size }
func (this *ZendLexState) SetScriptFilteredSize(value int)          { this.script_filtered_size = value }
func (this ZendLexState) GetInputFilter() ZendEncodingFilter        { return this.input_filter }
func (this *ZendLexState) SetInputFilter(value ZendEncodingFilter)  { this.input_filter = value }
func (this ZendLexState) GetOutputFilter() ZendEncodingFilter       { return this.output_filter }
func (this *ZendLexState) SetOutputFilter(value ZendEncodingFilter) { this.output_filter = value }
func (this ZendLexState) GetScriptEncoding() *ZendEncoding          { return this.script_encoding }
func (this *ZendLexState) SetScriptEncoding(value *ZendEncoding)    { this.script_encoding = value }
func (this ZendLexState) GetOnEvent() func(event ZendPhpScannerEvent, token int, line int, context any) {
	return this.on_event
}
func (this *ZendLexState) SetOnEvent(value func(event ZendPhpScannerEvent, token int, line int, context any)) {
	this.on_event = value
}
func (this ZendLexState) GetOnEventContext() any        { return this.on_event_context }
func (this *ZendLexState) SetOnEventContext(value any)  { this.on_event_context = value }
func (this ZendLexState) GetAst() *ZendAst              { return this.ast }
func (this *ZendLexState) SetAst(value *ZendAst)        { this.ast = value }
func (this ZendLexState) GetAstArena() *ZendArena       { return this.ast_arena }
func (this *ZendLexState) SetAstArena(value *ZendArena) { this.ast_arena = value }

/**
 * ZendHeredocLabel
 */
type ZendHeredocLabel struct {
	label                   *byte
	length                  int
	indentation             int
	indentation_uses_spaces ZendBool
}

func (this ZendHeredocLabel) GetLabel() *byte                    { return this.label }
func (this *ZendHeredocLabel) SetLabel(value *byte)              { this.label = value }
func (this ZendHeredocLabel) GetLength() int                     { return this.length }
func (this *ZendHeredocLabel) SetLength(value int)               { this.length = value }
func (this ZendHeredocLabel) GetIndentation() int                { return this.indentation }
func (this *ZendHeredocLabel) SetIndentation(value int)          { this.indentation = value }
func (this ZendHeredocLabel) GetIndentationUsesSpaces() ZendBool { return this.indentation_uses_spaces }
func (this *ZendHeredocLabel) SetIndentationUsesSpaces(value ZendBool) {
	this.indentation_uses_spaces = value
}
