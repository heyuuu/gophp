// <<generate>>

package zend

import b "sik/builtin"

// Source: <Zend/zend_language_scanner.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_SCANNER_H

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
	input_filter         ZendEncodingFilter
	output_filter        ZendEncodingFilter
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
	label                 *byte
	length                int
	indentation           int
	indentationUsesSpaces bool
}

func NewHeredocLabel(label string) *ZendHeredocLabel {
	return &ZendHeredocLabel{
		label:       label,
		indentation: 0,
	}
}

func (l *ZendHeredocLabel) GetLabel() string { return "" } // todo
func (l *ZendHeredocLabel) Copy() *ZendHeredocLabel {
	newLabel := *l
	return &newLabel
}
