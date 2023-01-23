// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_ini_parser.h>

/* A Bison parser, made by GNU Bison 3.0.2.  */

// #define YY_INI_ZEND_ZEND_INI_PARSER_H_INCLUDED

/* Debug traces.  */

// #define YYDEBUG       0

/* Token type.  */

// #define YYTOKENTYPE

type Yytokentype = int

const (
	END              = 0
	TC_SECTION       = 258
	TC_RAW           = 259
	TC_CONSTANT      = 260
	TC_NUMBER        = 261
	TC_STRING        = 262
	TC_WHITESPACE    = 263
	TC_LABEL         = 264
	TC_OFFSET        = 265
	TC_DOLLAR_CURLY  = 266
	TC_VARNAME       = 267
	TC_QUOTED_STRING = 268
	BOOL_TRUE        = 269
	BOOL_FALSE       = 270
	NULL_NULL        = 271
	END_OF_LINE      = 272
)

/* Value type.  */

type YYSTYPE = int

// #define YYSTYPE_IS_TRIVIAL       1

// #define YYSTYPE_IS_DECLARED       1

// Source: <Zend/zend_ini_parser.c>

/* A Bison parser, made by GNU Bison 3.0.2.  */

// #define YYBISON       1

/* Bison version.  */

// #define YYBISON_VERSION       "3.0.2"

/* Skeleton name.  */

// #define YYSKELETON_NAME       "yacc.c"

/* Pure parsers.  */

// #define YYPURE       2

/* Push parsers.  */

// #define YYPUSH       0

/* Pull parsers.  */

// #define YYPULL       1

/* Substitute the variable and function names.  */

// #define yyparse       ini_parse

// #define yylex       ini_lex

// #define yyerror       ini_error

// #define yydebug       ini_debug

// #define yynerrs       ini_nerrs

/* Copy the first part of user declarations.  */

// #define DEBUG_CFG_PARSER       0

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_ini.h"

// # include "zend_constants.h"

// # include "zend_ini_scanner.h"

// # include "zend_extensions.h"

// #define YYSTYPE       zval

// #define ZEND_INI_PARSER_CB       ( CG ( ini_parser_param ) ) -> ini_parser_cb

// #define ZEND_INI_PARSER_ARG       ( CG ( ini_parser_param ) ) -> arg

// #define ZEND_SYSTEM_INI       CG ( ini_parser_unbuffered_errors )

func GetIntVal(op *Zval) int {
	switch op.GetType() {
	case 4:
		return op.GetValue().GetLval()
	case 5:
		return int(op.GetValue().GetDval())
	case 6:
		var val int = atoi(op.GetValue().GetStr().GetVal())
		ZendStringFree(op.GetValue().GetStr())
		return val
	default:
		break
	}
}

/* {{{ zend_ini_do_op()
 */

func ZendIniDoOp(type_ byte, result *Zval, op1 *Zval, op2 *Zval) {
	var i_result int
	var i_op1 int
	var i_op2 int
	var str_len int
	var str_result []byte
	i_op1 = GetIntVal(op1)
	if op2 != nil {
		i_op2 = GetIntVal(op2)
	} else {
		i_op2 = 0
	}
	switch type_ {
	case '|':
		i_result = i_op1 | i_op2
		break
	case '&':
		i_result = i_op1 & i_op2
		break
	case '^':
		i_result = i_op1 ^ i_op2
		break
	case '~':
		i_result = ^i_op1
		break
	case '!':
		i_result = !i_op1
		break
	default:
		i_result = 0
		break
	}
	str_len = sprintf(str_result, "%d", i_result)
	var __z *Zval = result
	var __s *ZendString = ZendStringInit(str_result, str_len, CG.GetIniParserUnbufferedErrors())
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
}

/* }}} */

func ZendIniInitString(result *Zval) {
	if CG.GetIniParserUnbufferedErrors() != 0 {
		var __z *Zval = result
		var __s *ZendString = ZendStringInit("", 0, 1)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else {
		var __z *Zval = result
		var __s *ZendString = ZendEmptyString
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
	}
}

/* }}} */

func ZendIniAddString(result *Zval, op1 *Zval, op2 *Zval) {
	var length int
	var op1_len int
	if op1.GetType() != 6 {

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

		if CG.GetIniParserUnbufferedErrors() != 0 {
			var tmp_str *ZendString
			var str *ZendString = ZvalGetTmpString(op1, &tmp_str)
			var __z *Zval = op1
			var __s *ZendString = ZendStringInit(str.GetVal(), str.GetLen(), 1)
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			ZendTmpStringRelease(tmp_str)
		} else {
			var __z *Zval = op1
			var __s *ZendString = ZvalGetStringFunc(op1)
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		}

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

	}
	op1_len = int(op1.GetValue().GetStr().GetLen())
	if op2.GetType() != 6 {
		if op2.GetType() != 6 {
			_convertToString(op2)
		}
	}
	length = op1_len + int(op2.GetValue().GetStr().GetLen())
	var __z *Zval = result
	var __s *ZendString = ZendStringExtend(op1.GetValue().GetStr(), length, CG.GetIniParserUnbufferedErrors())
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	memcpy(result.GetValue().GetStr().GetVal()+op1_len, op2.GetValue().GetStr().GetVal(), op2.GetValue().GetStr().GetLen()+1)
}

/* }}} */

func ZendIniGetConstant(result *Zval, name *Zval) {
	var c *Zval
	var tmp Zval

	/* If name contains ':' it is not a constant. Bug #26893. */

	if !(memchr(name.GetValue().GetStr().GetVal(), ':', name.GetValue().GetStr().GetLen())) && g.Assign(&c, ZendGetConstant(name.GetValue().GetStr())) != 0 {
		if c.GetType() != 6 {
			var _z1 *Zval = &tmp
			var _z2 *Zval = c
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
					ZendGcAddref(&_gc.gc)
				} else {
					ZvalCopyCtorFunc(_z1)
				}
			}
			if (tmp.GetTypeInfo() & 0xff) == 11 {
				ZvalUpdateConstantEx(&tmp, nil)
			}
			if &tmp.GetType() != 6 {
				_convertToString(&tmp)
			}
			c = &tmp
		}
		var __z *Zval = result
		var __s *ZendString = ZendStringInit(c.GetValue().GetStr().GetVal(), c.GetValue().GetStr().GetLen(), CG.GetIniParserUnbufferedErrors())
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if c == &tmp {
			ZendStringRelease(tmp.GetValue().GetStr())
		}
		ZendStringFree(name.GetValue().GetStr())
	} else {
		*result = *name
	}

	/* If name contains ':' it is not a constant. Bug #26893. */
}

/* }}} */

func ZendIniGetVar(result *Zval, name *Zval) {
	var curval *Zval
	var envvar *byte

	/* Fetch configuration option value */

	if g.Assign(&curval, ZendGetConfigurationDirective(name.GetValue().GetStr())) != nil {
		var __z *Zval = result
		var __s *ZendString = ZendStringInit(curval.GetValue().GetStr().GetVal(), curval.GetValue().GetStr().GetLen(), CG.GetIniParserUnbufferedErrors())
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else if g.Assign(&envvar, ZendGetenv(name.GetValue().GetStr().GetVal(), name.GetValue().GetStr().GetLen())) != nil || g.Assign(&envvar, getenv(name.GetValue().GetStr().GetVal())) != nil {
		var __z *Zval = result
		var __s *ZendString = ZendStringInit(envvar, strlen(envvar), CG.GetIniParserUnbufferedErrors())
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else {
		ZendIniInitString(result)
	}

	/* Fetch configuration option value */
}

/* }}} */

func IniError(msg *byte) {
	var error_buf *byte
	var error_buf_len int
	var currently_parsed_filename *byte
	currently_parsed_filename = ZendIniScannerGetFilename()
	if currently_parsed_filename != nil {
		error_buf_len = 128 + int(strlen(msg)+int(strlen(currently_parsed_filename)))
		error_buf = (*byte)(_emalloc(error_buf_len))
		sprintf(error_buf, "%s in %s on line %d\n", msg, currently_parsed_filename, ZendIniScannerGetLineno())
	} else {
		error_buf = _estrdup("Invalid configuration directive\n")
	}
	if CG.GetIniParserUnbufferedErrors() != 0 {
		r.Fprintf(stderr, "PHP:  %s", error_buf)
	} else {
		ZendError(1<<1, "%s", error_buf)
	}
	_efree(error_buf)
}

/* }}} */

func ZendParseIniFile(fh *ZendFileHandle, unbuffered_errors ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CG.SetIniParserParam(&ini_parser_param)
	if ZendIniOpenFileForScanning(fh, scanner_mode) == FAILURE {
		return FAILURE
	}
	CG.SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	ZendFileHandleDtor(fh)
	ShutdownIniScanner()
	if retval == 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendParseIniString(str *byte, unbuffered_errors ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CG.SetIniParserParam(&ini_parser_param)
	if ZendIniPrepareStringForScanning(str, scanner_mode) == FAILURE {
		return FAILURE
	}
	CG.SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	ShutdownIniScanner()
	if retval == 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZvalIniDtor(zv *Zval) {
	if zv.GetType() == 6 {
		ZendStringRelease(zv.GetValue().GetStr())
	}
}

/* }}} */

// #define YY_NULLPTR       0

/* Enabling verbose error messages.  */

// #define YYERROR_VERBOSE       1

/* In a future release of Bison, this section will be replaced
   by #include "zend_ini_parser.h".  */

// #define YY_INI_ZEND_ZEND_INI_PARSER_H_INCLUDED

/* Debug traces.  */

// #define YYDEBUG       0

/* Token type.  */

// #define YYTOKENTYPE

const (
	END              = 0
	TC_SECTION       = 258
	TC_RAW           = 259
	TC_CONSTANT      = 260
	TC_NUMBER        = 261
	TC_STRING        = 262
	TC_WHITESPACE    = 263
	TC_LABEL         = 264
	TC_OFFSET        = 265
	TC_DOLLAR_CURLY  = 266
	TC_VARNAME       = 267
	TC_QUOTED_STRING = 268
	BOOL_TRUE        = 269
	BOOL_FALSE       = 270
	NULL_NULL        = 271
	END_OF_LINE      = 272
)

/* Value type.  */

/* Copy the second part of user declarations.  */

type IniYytypeUint8 = uint8
type IniYytypeInt8 = signed__char
type IniYytypeUint16 = unsigned__short__int
type IniYytypeInt16 = short__int

// # include < stddef . h >

// #define YYSIZE_T       size_t

// #define YYSIZE_MAXIMUM       ( ( YYSIZE_T ) - 1 )

// #define YY_(Msgid) Msgid

// #define YY_ATTRIBUTE(Spec)

// #define YY_ATTRIBUTE_PURE       YY_ATTRIBUTE ( ( __pure__ ) )

// #define YY_ATTRIBUTE_UNUSED       YY_ATTRIBUTE ( ( __unused__ ) )

// #define _Noreturn       YY_ATTRIBUTE ( ( __noreturn__ ) )

/* Suppress unused-variable warnings by "using" E.  */

// #define YYUSE(E) ( ( void ) ( E ) )

// #define YY_INITIAL_VALUE(Value) Value

// #define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN

// #define YY_IGNORE_MAYBE_UNINITIALIZED_END

/* The parser invokes alloca or malloc; define the __special__  necessary symbols.  */

// #define YYSTACK_ALLOC       YYMALLOC

// #define YYSTACK_FREE       YYFREE

// #define YYSTACK_ALLOC_MAXIMUM       YYSIZE_MAXIMUM

// #define YYMALLOC       malloc

var Malloc func(int) any

// #define YYFREE       free

var Free func(any)

/* A type that is properly aligned for any stack member.  */

/* The size of the maximum gap between one aligned stack and the next.  */

// #define YYSTACK_GAP_MAXIMUM       ( sizeof ( union yyalloc ) - 1 )

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */

// #define YYSTACK_BYTES(N) ( ( N ) * ( sizeof ( yytype_int16 ) + sizeof ( YYSTYPE ) ) + YYSTACK_GAP_MAXIMUM )

// #define YYCOPY_NEEDED       1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */

// #define YYSTACK_RELOCATE(Stack_alloc,Stack) do { YYSIZE_T yynewbytes ; YYCOPY ( & yyptr -> Stack_alloc , Stack , yysize ) ; Stack = & yyptr -> Stack_alloc ; yynewbytes = yystacksize * sizeof ( * Stack ) + YYSTACK_GAP_MAXIMUM ; yyptr += yynewbytes / sizeof ( * yyptr ) ; } while ( 0 )

/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */

// #define YYCOPY(Dst,Src,Count) do { YYSIZE_T yyi ; for ( yyi = 0 ; yyi < ( Count ) ; yyi ++ ) ( Dst ) [ yyi ] = ( Src ) [ yyi ] ; } while ( 0 )

/* YYFINAL -- State number of the termination state.  */

// #define YYFINAL       2

/* YYLAST -- Last index in YYTABLE.  */

// #define YYLAST       123

/* YYNTOKENS -- Number of terminals.  */

// #define YYNTOKENS       44

/* YYNNTS -- Number of nonterminals.  */

// #define YYNNTS       13

/* YYNRULES -- Number of rules.  */

// #define YYNRULES       50

/* YYNSTATES -- Number of states.  */

// #define YYNSTATES       72

/* YYTRANSLATE[YYX] -- Symbol number corresponding to YYX as returned
   by yylex, with out-of-bounds checking.  */

// #define YYUNDEFTOK       2

// #define YYMAXUTOK       272

// #define YYTRANSLATE(YYX) ( ( unsigned int ) ( YYX ) <= YYMAXUTOK ? yytranslate [ YYX ] : YYUNDEFTOK )

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, without out-of-bounds checking.  */

var IniYytranslate []yytype_uint8 = []yytype_uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 40, 22, 2, 30, 29, 39, 23, 42, 43, 28, 25, 20, 26, 21, 27, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 19, 2, 32, 18, 33, 34, 35, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 41, 24, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 36, 38, 37, 31, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */

var IniYytname []*byte = []*byte{"\"end of file\"", "error", "$undefined", "TC_SECTION", "TC_RAW", "TC_CONSTANT", "TC_NUMBER", "TC_STRING", "TC_WHITESPACE", "TC_LABEL", "TC_OFFSET", "TC_DOLLAR_CURLY", "TC_VARNAME", "TC_QUOTED_STRING", "BOOL_TRUE", "BOOL_FALSE", "NULL_NULL", "END_OF_LINE", "'='", "':'", "','", "'.'", "'\"'", "'\\''", "'^'", "'+'", "'-'", "'/'", "'*'", "'%'", "'$'", "'~'", "'<'", "'>'", "'?'", "'@'", "'{'", "'}'", "'|'", "'&'", "'!'", "']'", "'('", "')'", "$accept", "statement_list", "statement", "section_string_or_value", "string_or_value", "option_offset", "encapsed_list", "var_string_list_section", "var_string_list", "expr", "cfg_var_ref", "constant_literal", "constant_string", 0}

// #define YYPACT_NINF       - 25

// #define yypact_value_is_default(Yystate) ( ! ! ( ( Yystate ) == ( - 25 ) ) )

// #define YYTABLE_NINF       - 1

// #define yytable_value_is_error(Yytable_value) 0

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */

var IniYypact []yytype_int8 = []yytype_int8{-25, 9, -25, 73, -17, 81, -25, -25, -25, -25, -25, -25, -25, 15, -25, -20, 93, -25, -25, 0, -25, -25, -25, -25, -25, -25, -12, 101, -25, -25, -7, 36, -25, -25, -25, -25, -25, -25, -25, -25, 28, 28, 28, -25, 101, -1, 40, 30, -25, -25, -25, -25, -25, -25, -25, 80, -25, -25, 33, 28, 28, 28, -25, 0, 100, -25, -25, -25, -25, -25, -25, -25}

/* YYDEFACT[STATE-NUM] -- Default reduction number in state STATE-NUM.
   Performed when YYTABLE does not specify something else to do.  Zero
   means the default is an error.  */

var IniYydefact []yytype_uint8 = []yytype_uint8{3, 0, 1, 10, 7, 17, 8, 2, 42, 41, 43, 44, 45, 0, 20, 0, 9, 21, 22, 0, 47, 46, 48, 49, 50, 20, 0, 16, 27, 28, 0, 0, 4, 20, 24, 25, 12, 13, 14, 15, 0, 0, 0, 5, 33, 11, 0, 0, 20, 30, 31, 40, 19, 23, 18, 0, 37, 38, 0, 0, 0, 0, 29, 0, 0, 26, 39, 36, 34, 35, 6, 32}

/* YYPGOTO[NTERM-NUM].  */

var IniYypgoto []yytype_int8 = []yytype_int8{-25, -25, -25, -25, -9, -25, -23, -25, 50, 4, -3, 44, -24}

/* YYDEFGOTO[NTERM-NUM].  */

var IniYydefgoto []yytype_int8 = []yytype_int8{-1, 1, 7, 15, 43, 26, 31, 16, 44, 45, 28, 18, 29}

/* YYTABLE[YYPACT[STATE-NUM]] -- What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule whose
   number is the opposite.  If YYTABLE_NINF, syntax error.  */

var IniYytable []yytype_uint8 = []yytype_uint8{17, 19, 46, 50, 20, 21, 22, 23, 24, 2, 55, 13, 3, 34, 36, 37, 38, 39, 4, 5, 50, 32, 25, 59, 49, 64, 6, 30, 54, 47, 51, 40, 20, 21, 22, 23, 24, 60, 61, 13, 41, 49, 42, 54, 56, 57, 58, 13, 63, 52, 25, 13, 54, 52, 70, 27, 0, 59, 53, 40, 35, 54, 62, 67, 68, 69, 0, 0, 41, 0, 42, 60, 61, 0, 0, 0, 66, 8, 9, 10, 11, 12, 0, 0, 13, 20, 21, 22, 23, 24, 0, 13, 13, 52, 0, 14, 0, 8, 9, 10, 11, 12, 65, 25, 13, 20, 21, 22, 23, 24, 0, 13, 13, 52, 0, 33, 0, 0, 0, 0, 0, 0, 71, 48}
var IniYycheck []yytype_int8 = []yytype_int8{3, 18, 25, 27, 4, 5, 6, 7, 8, 0, 33, 11, 3, 16, 14, 15, 16, 17, 9, 10, 44, 41, 22, 24, 27, 48, 17, 12, 31, 41, 37, 31, 4, 5, 6, 7, 8, 38, 39, 11, 40, 44, 42, 46, 40, 41, 42, 11, 18, 13, 22, 11, 55, 13, 63, 5, -1, 24, 22, 31, 16, 64, 22, 59, 60, 61, -1, -1, 40, -1, 42, 38, 39, -1, -1, -1, 43, 4, 5, 6, 7, 8, -1, -1, 11, 4, 5, 6, 7, 8, -1, 11, 11, 13, -1, 22, -1, 4, 5, 6, 7, 8, 22, 22, 11, 4, 5, 6, 7, 8, -1, 11, 11, 13, -1, 22, -1, -1, -1, -1, -1, -1, 22, 22}

/* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
   symbol of state STATE-NUM.  */

var IniYystos []yytype_uint8 = []yytype_uint8{0, 45, 0, 3, 9, 10, 17, 46, 4, 5, 6, 7, 8, 11, 22, 47, 51, 54, 55, 18, 4, 5, 6, 7, 8, 22, 49, 52, 54, 56, 12, 50, 41, 22, 54, 55, 14, 15, 16, 17, 31, 40, 42, 48, 52, 53, 50, 41, 22, 54, 56, 37, 13, 22, 54, 50, 53, 53, 53, 24, 38, 39, 22, 18, 50, 22, 43, 53, 53, 53, 48, 22}

/* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */

var IniYyr1 []yytype_uint8 = []yytype_uint8{0, 44, 45, 45, 46, 46, 46, 46, 46, 47, 47, 48, 48, 48, 48, 48, 49, 49, 50, 50, 50, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 53, 53, 53, 53, 53, 53, 53, 54, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56}

/* YYR2[YYN] -- Number of symbols on the right hand side of rule YYN.  */

var IniYyr2 []yytype_uint8 = []yytype_uint8{0, 2, 2, 0, 3, 3, 5, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 2, 2, 0, 1, 1, 3, 2, 2, 4, 1, 1, 3, 2, 2, 4, 1, 3, 3, 3, 2, 2, 3, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

// #define yyerrok       ( yyerrstatus = 0 )

// #define yyclearin       ( yychar = YYEMPTY )

// #define YYEMPTY       ( - 2 )

// #define YYEOF       0

// #define YYACCEPT       goto yyacceptlab

// #define YYABORT       goto yyabortlab

// #define YYERROR       goto yyerrorlab

// #define YYRECOVERING() ( ! ! yyerrstatus )

// #define YYBACKUP(Token,Value) do if ( yychar == YYEMPTY ) { yychar = ( Token ) ; yylval = ( Value ) ; YYPOPSTACK ( yylen ) ; yystate = * yyssp ; goto yybackup ; } else { yyerror ( YY_ ( "syntax error: cannot back up" ) ) ; YYERROR ; } while ( 0 )

/* Error token number */

// #define YYTERROR       1

// #define YYERRCODE       256

/* Enable debugging if requested.  */

// #define YYDPRINTF(Args)

// #define YY_SYMBOL_PRINT(Title,Type,Value,Location)

// #define YY_STACK_PRINT(Bottom,Top)

// #define YY_REDUCE_PRINT(Rule)

/* YYINITDEPTH -- initial size of the parser's stacks.  */

// #define YYINITDEPTH       200

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

// #define YYMAXDEPTH       10000

/* Return the length of YYSTR.  */

func IniYystrlen(yystr *byte) int {
	var yylen int
	for yylen = 0; yystr[yylen]; yylen++ {
		continue
	}
	return yylen
}

/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */

func IniYystpcpy(yydest *byte, yysrc *byte) *byte {
	var yyd *byte = yydest
	var yys *byte = yysrc
	for g.Assign(&(g.PostInc(&(*yyd))), g.PostInc(&(*yys))) != '0' {
		continue
	}
	return yyd - 1
}

/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for yyerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from yytname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */

func IniYytnamerr(yyres *byte, yystr *byte) int {
	if (*yystr) == '"' {
		var yyn int = 0
		var yyp *byte = yystr
		for {
			switch *(g.PreInc(&yyp)) {
			case '\'':

			case ',':
				goto do_not_strip_quotes
			case '\\':
				if (*(g.PreInc(&yyp))) != '\\' {
					goto do_not_strip_quotes
				}
			default:
				if yyres != nil {
					yyres[yyn] = *yyp
				}
				yyn++
				break
			case '"':
				if yyres != nil {
					yyres[yyn] = '0'
				}
				return yyn
			}
		}
	do_not_strip_quotes:
	}
	if yyres == nil {
		return yystrlen(yystr)
	}
	return yystpcpy(yyres, yystr) - yyres
}

/* Copy into *YYMSG, which is of size *YYMSG_ALLOC, an error message
   about the unexpected token YYTOKEN for the state stack whose top is
   YYSSP.

   Return 0 if *YYMSG was successfully written.  Return 1 if *YYMSG is
   not large enough to hold the message.  In that case, also set
   *YYMSG_ALLOC to the required number of bytes.  Return 2 if the
   required number of bytes is too large to store.  */

func IniYysyntaxError(yymsg_alloc *int, yymsg **byte, yyssp *yytype_int16, yytoken int) int {
	var yysize0 int = yytnamerr(0, yytname[yytoken])
	var yysize int = yysize0
	const YYERROR_VERBOSE_ARGS_MAXIMUM = 5

	/* Internationalized format string. */

	var yyformat *byte = 0

	/* Arguments of yyformat. */

	var yyarg *[]byte

	/* Number of reported tokens (one for the "unexpected", one per
	   "expected"). */

	var yycount int = 0

	/* There are many possibilities here to consider:
	   - If this state is a consistent state with a default action, then
	     the only way this function was invoked is if the default action
	     is an error action.  In that case, don't check for expected
	     tokens because there are none.
	   - The only way there can be no lookahead present (in yychar) is if
	     this state is a consistent state with a default action.  Thus,
	     detecting the absence of a lookahead is sufficient to determine
	     that there is no unexpected or expected token to report.  In that
	     case, just report a simple "syntax error".
	   - Don't assume there isn't a lookahead just because this state is a
	     consistent state with a default action.  There might have been a
	     previous inconsistent state, consistent state with a non-default
	     action, or user semantic action that manipulated yychar.
	   - Of course, the expected token list depends on states to have
	     correct lookahead information, and it depends on the parser not
	     to perform extra reductions after fetching a lookahead from the
	     scanner and before detecting a syntax error.  Thus, state merging
	     (from LALR or IELR) and default reductions corrupt the expected
	     token list.  However, the list is correct for canonical LR with
	     one exception: it will still contain any token that will not be
	     accepted due to an error action in a later state.
	*/

	if yytoken != -2 {
		var yyn int = yypact[*yyssp]
		yyarg[g.PostInc(&yycount)] = yytname[yytoken]
		if yyn != -25 {

			/* Start YYX at -YYN if negative to avoid negative indexes in
			   YYCHECK.  In other words, skip the first -YYN actions for
			   this state because they are default actions.  */

			var yyxbegin int = g.Cond(yyn < 0, -yyn, 0)

			/* Stay within bounds of both yycheck and yytname.  */

			var yychecklim int = 123 - yyn + 1
			var yyxend int = g.Cond(yychecklim < 44, yychecklim, 44)
			var yyx int
			for yyx = yyxbegin; yyx < yyxend; yyx++ {
				if yycheck[yyx+yyn] == yyx && yyx != 1 {
					if yycount == YYERROR_VERBOSE_ARGS_MAXIMUM {
						yycount = 1
						yysize = yysize0
						break
					}
					yyarg[g.PostInc(&yycount)] = yytname[yyx]
					var yysize1 int = yysize + yytnamerr(0, yytname[yyx])
					if !(yysize <= yysize1 && yysize1 <= size_t-1) {
						return 2
					}
					yysize = yysize1
				}
			}
		}
	}
	switch yycount {
	case 0:
		yyformat = "syntax error"
		break
	case 1:
		yyformat = "syntax error, unexpected %s"
		break
	case 2:
		yyformat = "syntax error, unexpected %s, expecting %s"
		break
	case 3:
		yyformat = "syntax error, unexpected %s, expecting %s or %s"
		break
	case 4:
		yyformat = "syntax error, unexpected %s, expecting %s or %s or %s"
		break
	case 5:
		yyformat = "syntax error, unexpected %s, expecting %s or %s or %s or %s"
		break
	}
	var yysize1 int = yysize + yystrlen(yyformat)
	if !(yysize <= yysize1 && yysize1 <= size_t-1) {
		return 2
	}
	yysize = yysize1
	if (*yymsg_alloc) < yysize {
		*yymsg_alloc = 2 * yysize
		if !(yysize <= (*yymsg_alloc) && (*yymsg_alloc) <= size_t-1) {
			*yymsg_alloc = size_t - 1
		}
		return 1
	}

	/* Avoid sprintf, as that infringes on the user's name space.
	   Don't have undefined behavior even if the translation
	   produced a string with the wrong number of "%s"s.  */

	var yyp *byte = *yymsg
	var yyi int = 0
	for g.Assign(&(*yyp), *yyformat) != '0' {
		if (*yyp) == '%' && yyformat[1] == 's' && yyi < yycount {
			yyp += yytnamerr(yyp, yyarg[g.PostInc(&yyi)])
			yyformat += 2
		} else {
			yyp++
			yyformat++
		}
	}
	return 0
}

/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

func IniYydestruct(yymsg *byte, yytype int, yyvaluep *Zval) {
	void(yyvaluep)
	if yymsg == nil {
		yymsg = "Deleting"
	}
	switch yytype {
	case 4:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 5:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 6:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 7:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 8:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 9:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 10:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 12:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 14:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 15:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 16:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 47:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 48:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 49:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 50:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 51:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 52:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 53:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 54:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 55:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 56:
		ZvalIniDtor(&(*yyvaluep))
		break
	default:
		break
	}
}

/*----------.
| yyparse.  |
`----------*/

func IniParse() int {
	/* The lookahead symbol.  */

	var yychar int

	/* The semantic value of the lookahead symbol.  */

	var yyval_default Zval
	var yylval Zval = yyval_default

	/* Number of syntax errors so far.  */

	var ini_nerrs int
	var yystate int

	/* Number of tokens to shift before error messages enabled.  */

	var yyerrstatus int

	/* The stacks and their tools:
	   'yyss': related to states.
	   'yyvs': related to semantic values.

	   Refer to the stacks through separate pointers, to allow yyoverflow
	   to reallocate them elsewhere.  */

	var yyssa []yytype_int16
	var yyss *yytype_int16
	var yyssp *yytype_int16

	/* The semantic value stack.  */

	var yyvsa []Zval
	var yyvs *Zval
	var yyvsp *Zval
	var yystacksize int
	var yyn int
	var yyresult int

	/* Lookahead token as an internal (translated) token number.  */

	var yytoken int = 0

	/* The variables used to return semantic value and location from the
	   action routines.  */

	var yyval Zval

	/* Buffer for error messages, and its allocated size.  */

	var yymsgbuf []byte
	var yymsg *byte = yymsgbuf
	var yymsg_alloc int = g.SizeOf(yymsgbuf)

	// #define YYPOPSTACK(N) ( yyvsp -= ( N ) , yyssp -= ( N ) )

	/* The number of symbols on the RHS of the reduced rule.
	   Keep to zero when no symbol should be popped.  */

	var yylen int = 0
	yyss = yyssa
	yyssp = yyss
	yyvs = yyvsa
	yyvsp = yyvs
	yystacksize = 200
	yystate = 0
	yyerrstatus = 0
	ini_nerrs = 0
	yychar = -2
	goto yysetstate

	/*------------------------------------------------------------.
	  | yynewstate -- Push a new state, which is found in yystate.  |
	  `------------------------------------------------------------*/

yynewstate:

	/* In all cases, when you get here, the value and location stacks
	   have just been pushed.  So pushing a state here evens the stacks.  */

	yyssp++
yysetstate:
	*yyssp = yystate
	if yyss+yystacksize-1 <= yyssp {

		/* Get the current used size of the three stacks, in elements.  */

		var yysize int = yyssp - yyss + 1

		/* Extend the stack our own way.  */

		if 10000 <= yystacksize {
			goto yyexhaustedlab
		}
		yystacksize *= 2
		if 10000 < yystacksize {
			yystacksize = 10000
		}
		var yyss1 *yytype_int16 = yyss
		var yyptr *__union__yyalloc = (*__union__yyalloc)(Malloc(yystacksize*(g.SizeOf("yytype_int16")+g.SizeOf("YYSTYPE")) + (g.SizeOf("union yyalloc") - 1)))
		if yyptr == nil {
			goto yyexhaustedlab
		}
		var yynewbytes int
		var yyi int
		for yyi = 0; yyi < yysize; yyi++ {
			&yyptr.yyss_alloc[yyi] = yyss[yyi]
		}
		yyss = &yyptr.yyss_alloc
		yynewbytes = yystacksize*g.SizeOf("* yyss") + (g.SizeOf("union yyalloc") - 1)
		yyptr += yynewbytes / g.SizeOf("* yyptr")
		var yynewbytes int
		var yyi int
		for yyi = 0; yyi < yysize; yyi++ {
			&yyptr.yyvs_alloc[yyi] = yyvs[yyi]
		}
		yyvs = &yyptr.yyvs_alloc
		yynewbytes = yystacksize*g.SizeOf("* yyvs") + (g.SizeOf("union yyalloc") - 1)
		yyptr += yynewbytes / g.SizeOf("* yyptr")
		if yyss1 != yyssa {
			Free(yyss1)
		}
		yyssp = yyss + yysize - 1
		yyvsp = yyvs + yysize - 1
		if yyss+yystacksize-1 <= yyssp {
			goto yyabortlab
		}
	}
	if yystate == 2 {
		goto yyacceptlab
	}
	goto yybackup

	/*-----------.
	  | yybackup.  |
	  `-----------*/

yybackup:

	/* Do appropriate processing given the current state.  Read a
	   lookahead token if we need one and don't already have one.  */

	yyn = yypact[yystate]
	if yyn == -25 {
		goto yydefault
	}

	/* Not known => get a lookahead token if don't already have one.  */

	if yychar == -2 {
		yychar = IniLex(&yylval)
	}
	if yychar <= 0 {
		yytoken = 0
		yychar = yytoken
	} else {
		yytoken = uint(g.CondF1(yychar <= 272, func() __auto__ { return yytranslate[yychar] }, 2))
	}

	/* If the proper action on seeing token YYTOKEN is to reduce or to
	   detect an error, take that action.  */

	yyn += yytoken
	if yyn < 0 || 123 < yyn || yycheck[yyn] != yytoken {
		goto yydefault
	}
	yyn = yytable[yyn]
	if yyn <= 0 {

		yyn = -yyn
		goto yyreduce
	}

	/* Count tokens shifted since error; after three, turn off error
	   status.  */

	if yyerrstatus != 0 {
		yyerrstatus--
	}

	/* Shift the lookahead token.  */

	/* Discard the shifted token.  */

	yychar = -2
	yystate = yyn
	*(g.PreInc(&yyvsp)) = yylval
	goto yynewstate

	/*-----------------------------------------------------------.
	  | yydefault -- do the default action for the current state.  |
	  `-----------------------------------------------------------*/

yydefault:
	yyn = yydefact[yystate]
	if yyn == 0 {
		goto yyerrlab
	}
	goto yyreduce

	/*-----------------------------.
	  | yyreduce -- Do a reduction.  |
	  `-----------------------------*/

yyreduce:

	/* yyn is the number of a rule to reduce with.  */

	yylen = yyr2[yyn]

	/* If YYLEN is nonzero, implement the default value of the action:
	   '$$ = $1'.

	   Otherwise, the following line sets YYVAL to garbage.
	   This behavior is undocumented and Bison
	   users should not rely upon it.  Assigning to YYVAL
	   unconditionally makes the parser a bit smaller, and it avoids a
	   GCC warning that YYVAL may be used uninitialized.  */

	yyval = yyvsp[1-yylen]
	switch yyn {
	case 4:
		CG.GetIniParserParam().GetIniParserCb()(&yyvsp[-1], nil, nil, 2, CG.GetIniParserParam().GetArg())
		ZendStringRelease(yyvsp[-1].GetValue().GetStr())
		break
	case 5:
		CG.GetIniParserParam().GetIniParserCb()(&yyvsp[-2], &yyvsp[0], nil, 1, CG.GetIniParserParam().GetArg())
		ZendStringRelease(yyvsp[-2].GetValue().GetStr())
		ZvalIniDtor(&yyvsp[0])
		break
	case 6:
		CG.GetIniParserParam().GetIniParserCb()(&yyvsp[-4], &yyvsp[0], &yyvsp[-3], 3, CG.GetIniParserParam().GetArg())
		ZendStringRelease(yyvsp[-4].GetValue().GetStr())
		ZvalIniDtor(&yyvsp[-3])
		ZvalIniDtor(&yyvsp[0])
		break
	case 7:
		CG.GetIniParserParam().GetIniParserCb()(&yyvsp[0], nil, nil, 1, CG.GetIniParserParam().GetArg())
		ZendStringRelease(yyvsp[0].GetValue().GetStr())
		break
	case 9:
		yyval = yyvsp[0]
		break
	case 10:
		ZendIniInitString(&yyval)
		break
	case 11:
		yyval = yyvsp[0]
		break
	case 12:
		yyval = yyvsp[0]
		break
	case 13:
		yyval = yyvsp[0]
		break
	case 14:
		yyval = yyvsp[0]
		break
	case 15:
		ZendIniInitString(&yyval)
		break
	case 16:
		yyval = yyvsp[0]
		break
	case 17:
		ZendIniInitString(&yyval)
		break
	case 18:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 19:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 20:
		ZendIniInitString(&yyval)
		break
	case 21:
		yyval = yyvsp[0]
		break
	case 22:
		yyval = yyvsp[0]
		break
	case 23:
		yyval = yyvsp[-1]
		break
	case 24:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 25:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 26:
		ZendIniAddString(&yyval, &yyvsp[-3], &yyvsp[-1])
		ZendStringFree(yyvsp[-1].GetValue().GetStr())
		break
	case 27:
		yyval = yyvsp[0]
		break
	case 28:
		yyval = yyvsp[0]
		break
	case 29:
		yyval = yyvsp[-1]
		break
	case 30:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 31:
		ZendIniAddString(&yyval, &yyvsp[-1], &yyvsp[0])
		ZendStringFree(yyvsp[0].GetValue().GetStr())
		break
	case 32:
		ZendIniAddString(&yyval, &yyvsp[-3], &yyvsp[-1])
		ZendStringFree(yyvsp[-1].GetValue().GetStr())
		break
	case 33:
		yyval = yyvsp[0]
		break
	case 34:
		ZendIniDoOp('|', &yyval, &yyvsp[-2], &yyvsp[0])
		break
	case 35:
		ZendIniDoOp('&', &yyval, &yyvsp[-2], &yyvsp[0])
		break
	case 36:
		ZendIniDoOp('^', &yyval, &yyvsp[-2], &yyvsp[0])
		break
	case 37:
		ZendIniDoOp('~', &yyval, &yyvsp[0], nil)
		break
	case 38:
		ZendIniDoOp('!', &yyval, &yyvsp[0], nil)
		break
	case 39:
		yyval = yyvsp[-1]
		break
	case 40:
		ZendIniGetVar(&yyval, &yyvsp[-1])
		ZendStringFree(yyvsp[-1].GetValue().GetStr())
		break
	case 41:
		yyval = yyvsp[0]
		break
	case 42:
		yyval = yyvsp[0]
		break
	case 43:
		yyval = yyvsp[0]
		break
	case 44:
		yyval = yyvsp[0]
		break
	case 45:
		yyval = yyvsp[0]
		break
	case 46:
		ZendIniGetConstant(&yyval, &yyvsp[0])
		break
	case 47:
		yyval = yyvsp[0]
		break
	case 48:
		yyval = yyvsp[0]
		break
	case 49:
		yyval = yyvsp[0]
		break
	case 50:
		yyval = yyvsp[0]
		break
	default:
		break
	}

	/* User semantic actions sometimes alter yychar, and that requires
	   that yytoken be updated with the new translation.  We take the
	   approach of translating immediately before every use of yytoken.
	   One alternative is translating here after every semantic action,
	   but that translation would be missed if the semantic action invokes
	   YYABORT, YYACCEPT, or YYERROR immediately after altering yychar or
	   if it invokes YYBACKUP.  In the case of YYABORT or YYACCEPT, an
	   incorrect destructor might then be invoked immediately.  In the
	   case of YYERROR or YYBACKUP, subsequent parser actions might lead
	   to an incorrect destructor call or verbose syntax error message
	   before the lookahead is translated.  */

	yyvsp -= yylen
	yyssp -= yylen
	yylen = 0
	*(g.PreInc(&yyvsp)) = yyval

	/* Now 'shift' the result of the reduction.  Determine what state
	   that goes to, based on the state we popped back to and the rule
	   number reduced by.  */

	yyn = yyr1[yyn]
	yystate = yypgoto[yyn-44] + (*yyssp)
	if 0 <= yystate && yystate <= 123 && yycheck[yystate] == (*yyssp) {
		yystate = yytable[yystate]
	} else {
		yystate = yydefgoto[yyn-44]
	}
	goto yynewstate

	/*--------------------------------------.
	  | yyerrlab -- here on detecting error.  |
	  `--------------------------------------*/

yyerrlab:

	/* Make sure we have latest lookahead translation.  See comments at
	   user semantic actions for why this is necessary.  */

	if yychar == -2 {
		yytoken = -2
	} else {
		yytoken = uint(g.CondF1(yychar <= 272, func() __auto__ { return yytranslate[yychar] }, 2))
	}

	/* If not already recovering from an error, report this error.  */

	if yyerrstatus == 0 {
		ini_nerrs++

		// #define YYSYNTAX_ERROR       yysyntax_error ( & yymsg_alloc , & yymsg , yyssp , yytoken )

		var yymsgp *byte = "syntax error"
		var yysyntax_error_status int
		yysyntax_error_status = yysyntax_error(&yymsg_alloc, &yymsg, yyssp, yytoken)
		if yysyntax_error_status == 0 {
			yymsgp = yymsg
		} else if yysyntax_error_status == 1 {
			if yymsg != yymsgbuf {
				Free(yymsg)
			}
			yymsg = (*byte)(Malloc(yymsg_alloc))
			if yymsg == nil {
				yymsg = yymsgbuf
				yymsg_alloc = g.SizeOf(yymsgbuf)
				yysyntax_error_status = 2
			} else {
				yysyntax_error_status = yysyntax_error(&yymsg_alloc, &yymsg, yyssp, yytoken)
				yymsgp = yymsg
			}
		}
		IniError(yymsgp)
		if yysyntax_error_status == 2 {
			goto yyexhaustedlab
		}

		// #define YYSYNTAX_ERROR       yysyntax_error ( & yymsg_alloc , & yymsg , yyssp , yytoken )

	}
	if yyerrstatus == 3 {

		/* If just tried and failed to reuse lookahead token after an
		   error, discard it.  */

		if yychar <= 0 {

			/* Return failure if at end of input.  */

			if yychar == 0 {
				goto yyabortlab
			}

			/* Return failure if at end of input.  */

		} else {
			yydestruct("Error: discarding", yytoken, &yylval)
			yychar = -2
		}

		/* If just tried and failed to reuse lookahead token after an
		   error, discard it.  */

	}

	/* Else will try to reuse lookahead token after shifting the error
	   token.  */

	goto yyerrlab1

	/*---------------------------------------------------.
	  | yyerrorlab -- error raised explicitly by YYERROR.  |
	  `---------------------------------------------------*/

yyerrorlab:

	/* Pacify compilers like GCC when the user code never invokes
	   YYERROR and the label yyerrorlab therefore never appears in user
	   code.  */

	/* Do not reclaim the symbols of the rule whose action triggered
	   this YYERROR.  */

	yyvsp -= yylen
	yyssp -= yylen
	yylen = 0
	yystate = *yyssp
	goto yyerrlab1

	/*-------------------------------------------------------------.
	  | yyerrlab1 -- common code for both syntax error and YYERROR.  |
	  `-------------------------------------------------------------*/

yyerrlab1:
	yyerrstatus = 3
	for {
		yyn = yypact[yystate]
		if yyn != -25 {
			yyn += 1
			if 0 <= yyn && yyn <= 123 && yycheck[yyn] == 1 {
				yyn = yytable[yyn]
				if 0 < yyn {
					break
				}
			}
		}

		/* Pop the current state because it cannot handle the error token.  */

		if yyssp == yyss {
			goto yyabortlab
		}
		yydestruct("Error: popping", yystos[yystate], yyvsp)
		yyvsp -= 1
		yyssp -= 1
		yystate = *yyssp
	}
	*(g.PreInc(&yyvsp)) = yylval

	/* Shift the error token.  */

	yystate = yyn
	goto yynewstate

	/*-------------------------------------.
	  | yyacceptlab -- YYACCEPT comes here.  |
	  `-------------------------------------*/

yyacceptlab:
	yyresult = 0
	goto yyreturn

	/*-----------------------------------.
	  | yyabortlab -- YYABORT comes here.  |
	  `-----------------------------------*/

yyabortlab:
	yyresult = 1
	goto yyreturn

	/*-------------------------------------------------.
	  | yyexhaustedlab -- memory exhaustion comes here.  |
	  `-------------------------------------------------*/

yyexhaustedlab:
	IniError("memory exhausted")
	yyresult = 2

	/* Fall through.  */

yyreturn:
	if yychar != -2 {

		/* Make sure we have latest lookahead translation.  See comments at
		   user semantic actions for why this is necessary.  */

		yytoken = uint(g.CondF1(yychar <= 272, func() __auto__ { return yytranslate[yychar] }, 2))
		yydestruct("Cleanup: discarding lookahead", yytoken, &yylval)
	}

	/* Do not reclaim the symbols of the rule whose action triggered
	   this YYABORT or YYACCEPT.  */

	yyvsp -= yylen
	yyssp -= yylen
	for yyssp != yyss {
		yydestruct("Cleanup: popping", yystos[*yyssp], yyvsp)
		yyvsp -= 1
		yyssp -= 1
	}
	if yyss != yyssa {
		Free(yyss)
	}
	if yymsg != yymsgbuf {
		Free(yymsg)
	}
	return yyresult
}
