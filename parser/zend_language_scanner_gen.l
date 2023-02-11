package zend

/*!header:re2c:on*/
package zend
/*!header:re2c:off*/

func _lexScan(zendlval *Zval, elem *ZendParserStackElem, sc *LangScanner) (int, bool) {
/*!re2c
re2c:yyfill:check = 0;
re2c:define:YYCTYPE = byte;
re2c:define:YYPEEK = "sc.peek()";
re2c:define:YYSKIP = "sc.skip()";
re2c:define:YYBACKUP = "sc.backup()";
re2c:define:YYRESTORE = "sc.restore()";
re2c:define:YYDEBUG = "sc.debug";
re2c:define:YYFILL = "sc.yyfill()";
re2c:define:YYGETCONDITION = "sc.getState()";
re2c:yyfill:check = 0;

LNUM	[0-9]+(_[0-9]+)*
DNUM	({LNUM}?"."{LNUM})|({LNUM}"."{LNUM}?)
EXPONENT_DNUM	(({LNUM}|{DNUM})[eE][+-]?{LNUM})
HNUM	"0x"[0-9a-fA-F]+(_[0-9a-fA-F]+)*
BNUM	"0b"[01]+(_[01]+)*
LABEL	[a-zA-Z_\x80-\xff][a-zA-Z0-9_\x80-\xff]*
WHITESPACE [ \n\r\t]+
TABS_AND_SPACES [ \t]*
TOKENS [;:,.\[\]()|^&+-/*=%!~$<>?@]
ANY_CHAR [^]
NEWLINE ("\r"|"\n"|"\r\n")

/* compute sc.len_ before each rule */
<!*> := sc.resetLen()

<ST_IN_SCRIPTING>"exit" {
	return sc.token(T_EXIT);
}

<ST_IN_SCRIPTING>"die" {
	return sc.token(T_EXIT);
}

<ST_IN_SCRIPTING>"fn" {
	return sc.token(T_FN);
}

<ST_IN_SCRIPTING>"function" {
	return sc.token(T_FUNCTION);
}

<ST_IN_SCRIPTING>"const" {
	return sc.token(T_CONST);
}

<ST_IN_SCRIPTING>"return" {
	return sc.token(T_RETURN);
}

<ST_IN_SCRIPTING>"yield"{WHITESPACE}"from"[^a-zA-Z0-9_\x80-\xff] {
	sc.setLen(sc.len_ - 1);
	sc.handleNewlines(sc.yyText, sc.len_);
	return sc.token(T_YIELD_FROM);
}

<ST_IN_SCRIPTING>"yield" {
	return sc.token(T_YIELD);
}

<ST_IN_SCRIPTING>"try" {
	return sc.token(T_TRY);
}

<ST_IN_SCRIPTING>"catch" {
	return sc.token(T_CATCH);
}

<ST_IN_SCRIPTING>"finally" {
	return sc.token(T_FINALLY);
}

<ST_IN_SCRIPTING>"throw" {
	return sc.token(T_THROW);
}

<ST_IN_SCRIPTING>"if" {
	return sc.token(T_IF);
}

<ST_IN_SCRIPTING>"elseif" {
	return sc.token(T_ELSEIF);
}

<ST_IN_SCRIPTING>"endif" {
	return sc.token(T_ENDIF);
}

<ST_IN_SCRIPTING>"else" {
	return sc.token(T_ELSE);
}

<ST_IN_SCRIPTING>"while" {
	return sc.token(T_WHILE);
}

<ST_IN_SCRIPTING>"endwhile" {
	return sc.token(T_ENDWHILE);
}

<ST_IN_SCRIPTING>"do" {
	return sc.token(T_DO);
}

<ST_IN_SCRIPTING>"for" {
	return sc.token(T_FOR);
}

<ST_IN_SCRIPTING>"endfor" {
	return sc.token(T_ENDFOR);
}

<ST_IN_SCRIPTING>"foreach" {
	return sc.token(T_FOREACH);
}

<ST_IN_SCRIPTING>"endforeach" {
	return sc.token(T_ENDFOREACH);
}

<ST_IN_SCRIPTING>"declare" {
	return sc.token(T_DECLARE);
}

<ST_IN_SCRIPTING>"enddeclare" {
	return sc.token(T_ENDDECLARE);
}

<ST_IN_SCRIPTING>"instanceof" {
	return sc.token(T_INSTANCEOF);
}

<ST_IN_SCRIPTING>"as" {
	return sc.token(T_AS);
}

<ST_IN_SCRIPTING>"switch" {
	return sc.token(T_SWITCH);
}

<ST_IN_SCRIPTING>"endswitch" {
	return sc.token(T_ENDSWITCH);
}

<ST_IN_SCRIPTING>"case" {
	return sc.token(T_CASE);
}

<ST_IN_SCRIPTING>"default" {
	return sc.token(T_DEFAULT);
}

<ST_IN_SCRIPTING>"break" {
	return sc.token(T_BREAK);
}

<ST_IN_SCRIPTING>"continue" {
	return sc.token(T_CONTINUE);
}

<ST_IN_SCRIPTING>"goto" {
	return sc.token(T_GOTO);
}

<ST_IN_SCRIPTING>"echo" {
	return sc.token(T_ECHO);
}

<ST_IN_SCRIPTING>"print" {
	return sc.token(T_PRINT);
}

<ST_IN_SCRIPTING>"class" {
	return sc.token(T_CLASS);
}

<ST_IN_SCRIPTING>"interface" {
	return sc.token(T_INTERFACE);
}

<ST_IN_SCRIPTING>"trait" {
	return sc.token(T_TRAIT);
}

<ST_IN_SCRIPTING>"extends" {
	return sc.token(T_EXTENDS);
}

<ST_IN_SCRIPTING>"implements" {
	return sc.token(T_IMPLEMENTS);
}

<ST_IN_SCRIPTING>"->" {
	sc.pushState(yycST_LOOKING_FOR_PROPERTY);
	return sc.token(T_OBJECT_OPERATOR);
}

<ST_IN_SCRIPTING,ST_LOOKING_FOR_PROPERTY>{WHITESPACE}+ {
	sc.handleNewlinesEx(sc.seg())
	return sc.returnOrSkipToken(T_WHITESPACE)
}

<ST_LOOKING_FOR_PROPERTY>"->" {
	return sc.token(T_OBJECT_OPERATOR);
}

<ST_LOOKING_FOR_PROPERTY>{LABEL} {
	sc.popState();
	return sc.tokenWithStr(T_STRING, 0);
}

<ST_LOOKING_FOR_PROPERTY>{ANY_CHAR} {
	sc.setLen(0);
	sc.popState();
	return 0, true
}

<ST_IN_SCRIPTING>"::" {
	return sc.token(T_PAAMAYIM_NEKUDOTAYIM);
}

<ST_IN_SCRIPTING>"\\" {
	return sc.token(T_NS_SEPARATOR);
}

<ST_IN_SCRIPTING>"..." {
	return sc.token(T_ELLIPSIS);
}

<ST_IN_SCRIPTING>"??" {
	return sc.token(T_COALESCE);
}

<ST_IN_SCRIPTING>"new" {
	return sc.token(T_NEW);
}

<ST_IN_SCRIPTING>"clone" {
	return sc.token(T_CLONE);
}

<ST_IN_SCRIPTING>"var" {
	return sc.token(T_VAR);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}("int"|"integer"){TABS_AND_SPACES}")" {
	return sc.token(T_INT_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}("double"|"float"){TABS_AND_SPACES}")" {
	return sc.token(T_DOUBLE_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}"real"{TABS_AND_SPACES}")" {
	if sc.isParserMode() {
		ZendError(E_DEPRECATED, "The (real) cast is deprecated, use (float) instead")
	}
	return sc.token(T_DOUBLE_CAST)
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}("string"|"binary"){TABS_AND_SPACES}")" {
	return sc.token(T_STRING_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}"array"{TABS_AND_SPACES}")" {
	return sc.token(T_ARRAY_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}"object"{TABS_AND_SPACES}")" {
	return sc.token(T_OBJECT_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}("bool"|"boolean"){TABS_AND_SPACES}")" {
	return sc.token(T_BOOL_CAST);
}

<ST_IN_SCRIPTING>"("{TABS_AND_SPACES}("unset"){TABS_AND_SPACES}")" {
	return sc.token(T_UNSET_CAST);
}

<ST_IN_SCRIPTING>"eval" {
	return sc.token(T_EVAL);
}

<ST_IN_SCRIPTING>"include" {
	return sc.token(T_INCLUDE);
}

<ST_IN_SCRIPTING>"include_once" {
	return sc.token(T_INCLUDE_ONCE);
}

<ST_IN_SCRIPTING>"require" {
	return sc.token(T_REQUIRE);
}

<ST_IN_SCRIPTING>"require_once" {
	return sc.token(T_REQUIRE_ONCE);
}

<ST_IN_SCRIPTING>"namespace" {
	return sc.token(T_NAMESPACE);
}

<ST_IN_SCRIPTING>"use" {
	return sc.token(T_USE);
}

<ST_IN_SCRIPTING>"insteadof" {
    return sc.token(T_INSTEADOF);
}

<ST_IN_SCRIPTING>"global" {
	return sc.token(T_GLOBAL);
}

<ST_IN_SCRIPTING>"isset" {
	return sc.token(T_ISSET);
}

<ST_IN_SCRIPTING>"empty" {
	return sc.token(T_EMPTY);
}

<ST_IN_SCRIPTING>"__halt_compiler" {
	return sc.token(T_HALT_COMPILER);
}

<ST_IN_SCRIPTING>"static" {
	return sc.token(T_STATIC);
}

<ST_IN_SCRIPTING>"abstract" {
	return sc.token(T_ABSTRACT);
}

<ST_IN_SCRIPTING>"final" {
	return sc.token(T_FINAL);
}

<ST_IN_SCRIPTING>"private" {
	return sc.token(T_PRIVATE);
}

<ST_IN_SCRIPTING>"protected" {
	return sc.token(T_PROTECTED);
}

<ST_IN_SCRIPTING>"public" {
	return sc.token(T_PUBLIC);
}

<ST_IN_SCRIPTING>"unset" {
	return sc.token(T_UNSET);
}

<ST_IN_SCRIPTING>"=>" {
	return sc.token(T_DOUBLE_ARROW);
}

<ST_IN_SCRIPTING>"list" {
	return sc.token(T_LIST);
}

<ST_IN_SCRIPTING>"array" {
	return sc.token(T_ARRAY);
}

<ST_IN_SCRIPTING>"callable" {
	return sc.token(T_CALLABLE);
}

<ST_IN_SCRIPTING>"++" {
	return sc.token(T_INC);
}

<ST_IN_SCRIPTING>"--" {
	return sc.token(T_DEC);
}

<ST_IN_SCRIPTING>"===" {
	return sc.token(T_IS_IDENTICAL);
}

<ST_IN_SCRIPTING>"!==" {
	return sc.token(T_IS_NOT_IDENTICAL);
}

<ST_IN_SCRIPTING>"==" {
	return sc.token(T_IS_EQUAL);
}

<ST_IN_SCRIPTING>"!="|"<>" {
	return sc.token(T_IS_NOT_EQUAL);
}

<ST_IN_SCRIPTING>"<=>" {
	return sc.token(T_SPACESHIP);
}

<ST_IN_SCRIPTING>"<=" {
	return sc.token(T_IS_SMALLER_OR_EQUAL);
}

<ST_IN_SCRIPTING>">=" {
	return sc.token(T_IS_GREATER_OR_EQUAL);
}

<ST_IN_SCRIPTING>"+=" {
	return sc.token(T_PLUS_EQUAL);
}

<ST_IN_SCRIPTING>"-=" {
	return sc.token(T_MINUS_EQUAL);
}

<ST_IN_SCRIPTING>"*=" {
	return sc.token(T_MUL_EQUAL);
}

<ST_IN_SCRIPTING>"*\*" {
	return sc.token(T_POW);
}

<ST_IN_SCRIPTING>"*\*=" {
	return sc.token(T_POW_EQUAL);
}

<ST_IN_SCRIPTING>"/=" {
	return sc.token(T_DIV_EQUAL);
}

<ST_IN_SCRIPTING>".=" {
	return sc.token(T_CONCAT_EQUAL);
}

<ST_IN_SCRIPTING>"%=" {
	return sc.token(T_MOD_EQUAL);
}

<ST_IN_SCRIPTING>"<<=" {
	return sc.token(T_SL_EQUAL);
}

<ST_IN_SCRIPTING>">>=" {
	return sc.token(T_SR_EQUAL);
}

<ST_IN_SCRIPTING>"&=" {
	return sc.token(T_AND_EQUAL);
}

<ST_IN_SCRIPTING>"|=" {
	return sc.token(T_OR_EQUAL);
}

<ST_IN_SCRIPTING>"^=" {
	return sc.token(T_XOR_EQUAL);
}

<ST_IN_SCRIPTING>"??=" {
	return sc.token(T_COALESCE_EQUAL);
}

<ST_IN_SCRIPTING>"||" {
	return sc.token(T_BOOLEAN_OR);
}

<ST_IN_SCRIPTING>"&&" {
	return sc.token(T_BOOLEAN_AND);
}

<ST_IN_SCRIPTING>"OR" {
	return sc.token(T_LOGICAL_OR);
}

<ST_IN_SCRIPTING>"AND" {
	return sc.token(T_LOGICAL_AND);
}

<ST_IN_SCRIPTING>"XOR" {
	return sc.token(T_LOGICAL_XOR);
}

<ST_IN_SCRIPTING>"<<" {
	return sc.token(T_SL);
}

<ST_IN_SCRIPTING>">>" {
	return sc.token(T_SR);
}

<ST_IN_SCRIPTING>{TOKENS} {
	return sc.token(sc.yyText0());
}


<ST_IN_SCRIPTING>"{" {
	sc.pushState(yycST_IN_SCRIPTING);
	return sc.token('{');
}


<ST_DOUBLE_QUOTES,ST_BACKQUOTE,ST_HEREDOC>"${" {
	sc.pushState(yycST_LOOKING_FOR_VARNAME);
	return sc.token(T_DOLLAR_OPEN_CURLY_BRACES);
}


<ST_IN_SCRIPTING>"}" {
	RESET_DOC_COMMENT()
	sc.tryPopState()
	return sc.token('}')
}


<ST_LOOKING_FOR_VARNAME>{LABEL}[[}] {
	sc.setLen(sc.len_ - 1)
	sc.popState()
	sc.pushState(yycST_IN_SCRIPTING)
	return sc.tokenWithStr(T_STRING_VARNAME, 0)
}


<ST_LOOKING_FOR_VARNAME>{ANY_CHAR} {
	sc.setLen(0)
	sc.popState()
	sc.pushState(yycST_IN_SCRIPTING)
	return 0, true
}

<ST_IN_SCRIPTING>{BNUM} {
	/* skips "0b" */
    return sc.tokenWithNumStr(2, 2)
}

<ST_IN_SCRIPTING>{LNUM} {
	var isOctal = sc.peek() == '0'
	if isOctal {
		return sc.tokenWithNumStr(0, 8)
	} else {
		return sc.tokenWithNumStr(0, 10)
	}
}

<ST_IN_SCRIPTING>{HNUM} {
	/* skips "0x" */
	return sc.tokenWithNumStr(2, 16)
}

<ST_VAR_OFFSET>[0]|([1-9][0-9]*) { /* Offset could be treated as a long */
	str := sc.seg()
	lVal, _, overflow := sc.parseNumStr(str, 10)
	if !overflow {
		sc.zendlval.SetLong(lVal)
	} else {
		sc.zendlval.SetRawString(str)
	}
	return sc.tokenWithVal(T_NUM_STRING)
}

<ST_VAR_OFFSET>{LNUM}|{HNUM}|{BNUM} { /* Offset must be treated as a string */
    return sc.lexerOffsetNum()
}

<ST_IN_SCRIPTING>{DNUM}|{EXPONENT_DNUM} {
    return sc.lexerDNum()
}

<ST_IN_SCRIPTING>"__CLASS__" {
	return sc.token(T_CLASS_C);
}

<ST_IN_SCRIPTING>"__TRAIT__" {
	return sc.token(T_TRAIT_C);
}

<ST_IN_SCRIPTING>"__FUNCTION__" {
	return sc.token(T_FUNC_C);
}

<ST_IN_SCRIPTING>"__METHOD__" {
	return sc.token(T_METHOD_C);
}

<ST_IN_SCRIPTING>"__LINE__" {
	return sc.token(T_LINE);
}

<ST_IN_SCRIPTING>"__FILE__" {
	return sc.token(T_FILE);
}

<ST_IN_SCRIPTING>"__DIR__" {
	return sc.token(T_DIR);
}

<ST_IN_SCRIPTING>"__NAMESPACE__" {
	return sc.token(T_NS_C);
}

<SHEBANG>"#!" .* {NEWLINE} {
	sc.lineno++;
	sc.begin(yycINITIAL);
	return 0, true
}

<SHEBANG>{ANY_CHAR} {
	sc.setLen(0);
	sc.begin(yycINITIAL);
	return 0, true
}

<INITIAL>"<?=" {
	sc.begin(yycST_IN_SCRIPTING)
	if sc.isParserMode() {
		return sc.token(T_ECHO)
	}
	return sc.token(T_OPEN_TAG_WITH_ECHO)
}


<INITIAL>"<?php"([ \t]|{NEWLINE}) {
	sc.handleNewline(sc.yyText[sc.len_-1]);
	sc.begin(yycST_IN_SCRIPTING);
	return sc.returnOrSkipToken(0)
}

<INITIAL>"<?php" {
    return sc.lexerRule12()
}

<INITIAL>"<?" {
    return sc.lexerRule13()
}

<INITIAL>{ANY_CHAR} {
    return sc.lexerRule14()
}


/* Make sure a label character follows "->", otherwise there is no property
 * and "->" will be taken literally
 */
<ST_DOUBLE_QUOTES,ST_HEREDOC,ST_BACKQUOTE>"$"{LABEL}"->"[a-zA-Z_\x80-\xff] {
	sc.setLen(sc.len_ - 3);
	sc.pushState(yycST_LOOKING_FOR_PROPERTY);
	return sc.tokenWithStr(T_VARIABLE, 1);
}

/* A [ always designates a variable offset, regardless of what follows
 */
<ST_DOUBLE_QUOTES,ST_HEREDOC,ST_BACKQUOTE>"$"{LABEL}"[" {
	sc.setLen(sc.len_ - 1);
	sc.pushState(yycST_VAR_OFFSET);
	return sc.tokenWithStr(T_VARIABLE, 1);
}

<ST_IN_SCRIPTING,ST_DOUBLE_QUOTES,ST_HEREDOC,ST_BACKQUOTE,ST_VAR_OFFSET>"$"{LABEL} {
	return sc.tokenWithStr(T_VARIABLE, 1);
}

<ST_VAR_OFFSET>"]" {
	sc.popState();
	return sc.token(']');
}

<ST_VAR_OFFSET>{TOKENS}|[{}"`] {
	/* Only '[' or '-' can be valid, but returning other tokens will allow a more explicit parse error */
	return sc.token(sc.yyText0());
}

<ST_VAR_OFFSET>[ \n\r\t\\'#] {
	/* Invalid rule to return a more explicit parse error with proper line number */
	sc.setLen(0)
	sc.popState()
	sc.zendlval.SetNull()
	return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
}

<ST_IN_SCRIPTING,ST_VAR_OFFSET>{LABEL} {
	return sc.tokenWithStr(T_STRING, 0);
}


<ST_IN_SCRIPTING>"#"|"//" {
    return sc.lexerRule16()
}

<ST_IN_SCRIPTING>"/*"|"/**"{WHITESPACE} {
    return sc.lexerRule17()
}

<ST_IN_SCRIPTING>"?>"{NEWLINE}? {
    return sc.lexerRule18()
}


<ST_IN_SCRIPTING>b?['] {
    return sc.lexerRule19()
}


<ST_IN_SCRIPTING>b?["] {
    return sc.lexerRule20()
}


<ST_IN_SCRIPTING>b?"<<<"{TABS_AND_SPACES}({LABEL}|([']{LABEL}['])|(["]{LABEL}["])){NEWLINE} {
    return sc.lexerRule21()
}


<ST_IN_SCRIPTING>[`] {
	sc.begin(yycST_BACKQUOTE);
	return sc.token('`');
}


<ST_END_HEREDOC>{ANY_CHAR} {
    return sc.lexerRule22()
}


<ST_DOUBLE_QUOTES,ST_BACKQUOTE,ST_HEREDOC>"{$" {
	sc.pushState(yycST_IN_SCRIPTING);
	sc.setLen(1);
	return sc.token(T_CURLY_OPEN);
}


<ST_DOUBLE_QUOTES>["] {
	sc.begin(yycST_IN_SCRIPTING);
	return sc.token('"');
}

<ST_BACKQUOTE>[`] {
	sc.begin(yycST_IN_SCRIPTING);
	return sc.token('`');
}


<ST_DOUBLE_QUOTES>{ANY_CHAR} {
    return sc.lexerRule23()
}


<ST_BACKQUOTE>{ANY_CHAR} {
    return sc.lexerRule24()
}


<ST_HEREDOC>{ANY_CHAR} {
    return sc.lexerRule25()
}


<ST_NOWDOC>{ANY_CHAR} {
    return sc.lexerRule26()
}


<ST_IN_SCRIPTING,ST_VAR_OFFSET>{ANY_CHAR} {
    return sc.lexerRule27()
}

*/
}
