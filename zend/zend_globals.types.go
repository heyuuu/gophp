package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

type ClassTable = *types.Table[*types.ClassEntry]
type FunctionTable = *types.Table[types.IFunction]
type ConstantTable = *types.Table[*ZendConstant]
type IniDirectives = *types.Table[*ZendIniEntry]
type ResourceTable = *types.Table[*types.Resource]

/**
 * ZendIniScannerGlobals
 */
type ZendIniScannerGlobals struct {
	yy_in        *FileHandle
	yy_out       *FileHandle
	yy_leng      uint
	yy_start     *uint8
	yy_text      *uint8
	yy_cursor    *uint8
	yy_marker    *uint8
	yy_limit     *uint8
	yy_state     int
	state_stack  []any
	filename     *byte
	lineno       int
	scanner_mode int
}

func (this *ZendIniScannerGlobals) GetYyIn() *FileHandle       { return this.yy_in }
func (this *ZendIniScannerGlobals) SetYyIn(value *FileHandle)  { this.yy_in = value }
func (this *ZendIniScannerGlobals) GetYyOut() *FileHandle      { return this.yy_out }
func (this *ZendIniScannerGlobals) SetYyOut(value *FileHandle) { this.yy_out = value }
func (this *ZendIniScannerGlobals) GetYyLeng() uint            { return this.yy_leng }
func (this *ZendIniScannerGlobals) SetYyLeng(value uint)       { this.yy_leng = value }
func (this *ZendIniScannerGlobals) GetYyStart() *uint8         { return this.yy_start }
func (this *ZendIniScannerGlobals) SetYyStart(value *uint8)    { this.yy_start = value }
func (this *ZendIniScannerGlobals) GetYyText() *uint8          { return this.yy_text }
func (this *ZendIniScannerGlobals) SetYyText(value *uint8)     { this.yy_text = value }
func (this *ZendIniScannerGlobals) GetYyCursor() *uint8        { return this.yy_cursor }
func (this *ZendIniScannerGlobals) SetYyCursor(value *uint8)   { this.yy_cursor = value }
func (this *ZendIniScannerGlobals) GetYyMarker() *uint8        { return this.yy_marker }
func (this *ZendIniScannerGlobals) SetYyMarker(value *uint8)   { this.yy_marker = value }
func (this *ZendIniScannerGlobals) GetYyLimit() *uint8         { return this.yy_limit }
func (this *ZendIniScannerGlobals) SetYyLimit(value *uint8)    { this.yy_limit = value }
func (this *ZendIniScannerGlobals) GetYyState() int            { return this.yy_state }
func (this *ZendIniScannerGlobals) SetYyState(value int)       { this.yy_state = value }
func (this *ZendIniScannerGlobals) GetStateStack() []any       { return this.state_stack }
func (this *ZendIniScannerGlobals) SetStateStack(value []any)  { this.state_stack = value }
func (this *ZendIniScannerGlobals) GetFilename() *byte         { return this.filename }
func (this *ZendIniScannerGlobals) SetFilename(value *byte)    { this.filename = value }
func (this *ZendIniScannerGlobals) GetLineno() int             { return this.lineno }
func (this *ZendIniScannerGlobals) SetLineno(value int)        { this.lineno = value }
func (this *ZendIniScannerGlobals) GetScannerMode() int        { return this.scanner_mode }
func (this *ZendIniScannerGlobals) SetScannerMode(value int)   { this.scanner_mode = value }

/**
 * ZendPhpScannerGlobals
 */
type ZendPhpScannerGlobals struct {
	yy_in                           *FileHandle
	yy_out                          *FileHandle
	yy_leng                         uint
	yy_start                        *uint8
	yy_text                         *uint8
	yy_cursor                       *uint8
	yy_marker                       *uint8
	yy_limit                        *uint8
	yy_state                        int
	state_stack                     b.Stack[int]
	heredoc_label_stack             b.Stack[*ZendHeredocLabel]
	heredoc_scan_ahead              bool
	heredoc_indentation             int
	heredoc_indentation_uses_spaces bool
	script_org                      *uint8
	script_org_size                 int
	script_filtered                 *uint8
	script_filtered_size            int
	script_encoding                 *ZendEncoding
	scanned_string_len              int
}
