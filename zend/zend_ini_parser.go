package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"log"
)

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

const YYSTYPE_IS_TRIVIAL = 1
const YYSTYPE_IS_DECLARED = 1

var IniParse func() int

/* A Bison parser, made by GNU Bison 3.0.2.  */

const YYBISON = 1

/* Bison version.  */

const YYBISON_VERSION = "3.0.2"

/* Skeleton name.  */

const YYSKELETON_NAME = "yacc.c"

/* Pure parsers.  */

const YYPURE = 2

/* Push parsers.  */

const YYPUSH = 0

/* Pull parsers.  */

const YYPULL = 1

/* Substitute the variable and function names.  */

const IniYyparse = IniParse
const IniYylex = IniLex
const IniYyerror = IniError
const IniYydebug = ini_debug
const IniYynerrs = ini_nerrs

/* Copy the first part of user declarations.  */

const DEBUG_CFG_PARSER = 0

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_ini.h"

// # include "zend_constants.h"

// # include "zend_ini_scanner.h"

// # include "zend_extensions.h"

// #define YYSTYPE       zval

var ZEND_INI_PARSER_CB = CG__().GetIniParserParam().GetIniParserCb()
var ZEND_INI_PARSER_ARG = CG__().GetIniParserParam().GetArg()
var ZEND_SYSTEM_INI = CG__().GetIniParserUnbufferedErrors()

func GetIntVal(op *types.Zval) int {
	switch op.GetType() {
	case types.IS_LONG:
		return op.Long()
	case types.IS_DOUBLE:
		return int(op.Double())
	case types.IS_STRING:
		var val int = atoi(op.String().GetVal())
		//types.ZendStringFree(op.String())
		return val
	default:

	}
}

/* {{{ zend_ini_do_op()
 */

func ZendIniDoOp(type_ byte, result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
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
	case '&':
		i_result = i_op1 & i_op2
	case '^':
		i_result = i_op1 ^ i_op2
	case '~':
		i_result = ^i_op1
	case '!':
		i_result = !i_op1
	default:
		i_result = 0
	}
	str_len = sprintf(str_result, "%d", i_result)
	result.SetString(types.NewString(b.CastStr(str_result, str_len)))
}

/* }}} */

func ZendIniInitString(result *types.Zval) {
	result.SetStringVal("")
}

/* }}} */

func ZendIniAddString(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	var length int
	var op1_len int
	if op1.GetType() != types.IS_STRING {

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

		if ZEND_SYSTEM_INI != 0 {
			var tmp_str *types.String
			var str *types.String = ZvalGetTmpString(op1, &tmp_str)
			op1.SetStringVal(str.GetStr())
			ZendTmpStringRelease(tmp_str)
		} else {
			op1.SetString(ZvalGetStringFunc(op1))
		}

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

	}
	op1_len = int(op1.String().GetLen())
	if op2.GetType() != types.IS_STRING {
		ConvertToString(op2)
	}
	length = op1_len + int(op2.String().GetLen())
	result.SetString(types.ZendStringExtend(op1.String(), length))
	memcpy(result.String().GetVal()+op1_len, op2.String().GetVal(), op2.String().GetLen()+1)
}

func ZendIniGetConstant(result *types.Zval, name *types.Zval) {
	var c *types.Zval
	var tmp types.Zval

	/* If name contains ':' it is not a constant. Bug #26893. */

	if !(memchr(name.String().GetVal(), ':', name.String().GetLen())) && b.Assign(&c, ZendGetConstant(name.GetStrVal())) != 0 {
		if c.GetType() != types.IS_STRING {
			types.ZVAL_COPY_OR_DUP(&tmp, c)
			if tmp.IsConstantAst() {
				ZvalUpdateConstantEx(&tmp, nil)
			}
			ConvertToString(&tmp)
			c = &tmp
		}
		result.SetString(types.NewString(c.String().GetStr()))
		if c == &tmp {
			// types.ZendStringRelease(tmp.String())
		}
		//types.ZendStringFree(name.String())
	} else {
		*result = *name
	}

	/* If name contains ':' it is not a constant. Bug #26893. */
}

func IniError(msg *byte) {
	var error_buf *byte
	var error_buf_len int
	var currently_parsed_filename *byte
	currently_parsed_filename = ZendIniScannerGetFilename()
	if currently_parsed_filename != nil {
		error_buf_len = 128 + int(strlen(msg)+int(strlen(currently_parsed_filename)))
		error_buf = (*byte)(Emalloc(error_buf_len))
		sprintf(error_buf, "%s in %s on line %d\n", msg, currently_parsed_filename, ZendIniScannerGetLineno())
	} else {
		error_buf = Estrdup("Invalid configuration directive\n")
	}
	if CG__().GetIniParserUnbufferedErrors() != 0 {
		log.Printf("PHP:  %s", error_buf)
	} else {
		faults.Error(faults.E_WARNING, "%s", error_buf)
	}
	Efree(error_buf)
}

/* }}} */

func ZendParseIniFile(fh *ZendFileHandle, unbuffered_errors types.ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CG__().SetIniParserParam(&ini_parser_param)
	if ZendIniOpenFileForScanning(fh, scanner_mode) == types.FAILURE {
		return types.FAILURE
	}
	CG__().SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	fh.Destroy()
	ShutdownIniScanner()
	if retval == 0 {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}

/* }}} */

func ZendParseIniString(str *byte, unbuffered_errors types.ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CG__().SetIniParserParam(&ini_parser_param)
	if ZendIniPrepareStringForScanning(str, scanner_mode) == types.FAILURE {
		return types.FAILURE
	}
	CG__().SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	ShutdownIniScanner()
	if retval == 0 {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}

/* }}} */

func ZvalIniDtor(zv *types.Zval) {
	if zv.IsString() {
		// types.ZendStringRelease(zv.String())
	}
}

/* }}} */

const YY_NULLPTR = 0

/* Enabling verbose error messages.  */

const YYERROR_VERBOSE = 1

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

const YYSIZE_T = size_t
const YYSIZE_MAXIMUM = YYSIZE_T - 1

func YY_(Msgid string) string { return Msgid }

// #define YY_ATTRIBUTE(Spec)

// #define YY_ATTRIBUTE_PURE       YY_ATTRIBUTE ( ( __pure__ ) )

// #define YY_ATTRIBUTE_UNUSED       YY_ATTRIBUTE ( ( __unused__ ) )

// #define _Noreturn       YY_ATTRIBUTE ( ( __noreturn__ ) )

/* Suppress unused-variable warnings by "using" E.  */

func YYUSE(E __auto__) { void(E) }

// #define YY_INITIAL_VALUE(Value) Value

// #define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN

// #define YY_IGNORE_MAYBE_UNINITIALIZED_END

/* The parser invokes alloca or malloc; define the __special__  necessary symbols.  */

const YYSTACK_ALLOC = YYMALLOC
const YYSTACK_FREE = YYFREE
const YYSTACK_ALLOC_MAXIMUM YYSIZE_T = YYSIZE_MAXIMUM
const YYMALLOC = Malloc

var Malloc func(YYSIZE_T) any

const YYFREE = Free

var Free func(any)

/* A type that is properly aligned for any stack member.  */

/**
 * IniUnionYyalloc
 */
type IniUnionYyalloc struct /* union */ {
	yyss_alloc yytype_int16
	yyvs_alloc types.Zval
}

// func MakeIniUnionYyalloc(yyss_alloc yytype_int16, yyvs_alloc Zval) IniUnionYyalloc {
//     return IniUnionYyalloc{
//         yyss_alloc:yyss_alloc,
//         yyvs_alloc:yyvs_alloc,
//     }
// }
// func (this *IniUnionYyalloc)  GetYyssAlloc() yytype_int16      { return this.yyss_alloc }
// func (this *IniUnionYyalloc) SetYyssAlloc(value yytype_int16) { this.yyss_alloc = value }
// func (this *IniUnionYyalloc)  GetYyvsAlloc() Zval      { return this.yyvs_alloc }
// func (this *IniUnionYyalloc) SetYyvsAlloc(value Zval) { this.yyvs_alloc = value }

/* The size of the maximum gap between one aligned stack and the next.  */

const YYSTACK_GAP_MAXIMUM = b.SizeOf("union yyalloc") - 1

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */

func YYSTACK_BYTES(N __auto__) int {
	return N*(b.SizeOf("yytype_int16")+b.SizeOf("YYSTYPE")) + YYSTACK_GAP_MAXIMUM
}

const YYCOPY_NEEDED = 1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */

// #define YYSTACK_RELOCATE(Stack_alloc,Stack) do { YYSIZE_T yynewbytes ; YYCOPY ( & yyptr -> Stack_alloc , Stack , yysize ) ; Stack = & yyptr -> Stack_alloc ; yynewbytes = yystacksize * sizeof ( * Stack ) + YYSTACK_GAP_MAXIMUM ; yyptr += yynewbytes / sizeof ( * yyptr ) ; } while ( 0 )

/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */

func YYCOPY(Dst __auto__, Src __auto__, Count __auto__) {
	var yyi YYSIZE_T
	for yyi = 0; yyi < Count; yyi++ {
		Dst[yyi] = Src[yyi]
	}
}

/* YYFINAL -- State number of the termination state.  */

const YYFINAL = 2

/* YYLAST -- Last index in YYTABLE.  */

const YYLAST = 123

/* YYNTOKENS -- Number of terminals.  */

const YYNTOKENS = 44

/* YYNNTS -- Number of nonterminals.  */

const YYNNTS = 13

/* YYNRULES -- Number of rules.  */

const YYNRULES = 50

/* YYNSTATES -- Number of states.  */

const YYNSTATES = 72

/* YYTRANSLATE[YYX] -- Symbol number corresponding to YYX as returned
   by yylex, with out-of-bounds checking.  */

const YYUNDEFTOK = 2
const YYMAXUTOK = 272

func YYTRANSLATE(YYX __auto__) uint {
	return uint(b.CondF1(YYX <= YYMAXUTOK, func() __auto__ { return yytranslate[YYX] }, YYUNDEFTOK))
}

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, without out-of-bounds checking.  */

var IniYytranslate []yytype_uint8 = []yytype_uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 40, 22, 2, 30, 29, 39, 23, 42, 43, 28, 25, 20, 26, 21, 27, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 19, 2, 32, 18, 33, 34, 35, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 41, 24, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 36, 38, 37, 31, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */

var IniYytname []*byte = []*byte{"\"end of file\"", "error", "$undefined", "TC_SECTION", "TC_RAW", "TC_CONSTANT", "TC_NUMBER", "TC_STRING", "TC_WHITESPACE", "TC_LABEL", "TC_OFFSET", "TC_DOLLAR_CURLY", "TC_VARNAME", "TC_QUOTED_STRING", "BOOL_TRUE", "BOOL_FALSE", "NULL_NULL", "END_OF_LINE", "'='", "':'", "','", "'.'", "'\"'", "'\\''", "'^'", "'+'", "'-'", "'/'", "'*'", "'%'", "'$'", "'~'", "'<'", "'>'", "'?'", "'@'", "'{'", "'}'", "'|'", "'&'", "'!'", "']'", "'('", "')'", "$accept", "statement_list", "statement", "section_string_or_value", "string_or_value", "option_offset", "encapsed_list", "var_string_list_section", "var_string_list", "expr", "cfg_var_ref", "constant_literal", "constant_string", YY_NULLPTR}

const YYPACT_NINF = -25

func IniYypactValueIsDefault(Yystate __auto__) bool { return !!(Yystate == -25) }

const YYTABLE_NINF = -1

func IniYytableValueIsError(Yytable_value __auto__) int { return 0 }

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

const IniYyerrok = b.Assign(&yyerrstatus, 0)
const IniYyclearin = b.Assign(&yychar, YYEMPTY)
const YYEMPTY = -2
const YYEOF = 0

// #define YYACCEPT       goto yyacceptlab

// #define YYABORT       goto yyabortlab

// #define YYERROR       goto yyerrorlab

func YYRECOVERING() bool { return !!yyerrstatus }
func YYBACKUP(Token __auto__, Value __auto__) {
	if yychar == YYEMPTY {
		yychar = Token
		yylval = Value
		YYPOPSTACK(yylen)
		yystate = *yyssp
		goto yybackup
	} else {
		yyerror(YY_("syntax error: cannot back up"))
		goto yyerrorlab
	}
}

/* Error token number */

const YYTERROR = 1
const YYERRCODE = 256

/* Enable debugging if requested.  */

// #define YYDPRINTF(Args)

// #define YY_SYMBOL_PRINT(Title,Type,Value,Location)

// #define YY_STACK_PRINT(Bottom,Top)

// #define YY_REDUCE_PRINT(Rule)

/* YYINITDEPTH -- initial size of the parser's stacks.  */

const YYINITDEPTH = 200

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

const YYMAXDEPTH = 10000

/* Return the length of YYSTR.  */

func IniYystrlen(yystr *byte) YYSIZE_T {
	var yylen YYSIZE_T
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
	for b.Assign(&(b.PostInc(&(*yyd))), b.PostInc(&(*yys))) != '0' {
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

func IniYytnamerr(yyres *byte, yystr *byte) YYSIZE_T {
	if (*yystr) == '"' {
		var yyn YYSIZE_T = 0
		var yyp *byte = yystr
		for {
			switch *(b.PreInc(&yyp)) {
			case '\'':
				fallthrough
			case ',':
				goto do_not_strip_quotes
			case '\\':
				if (*(b.PreInc(&yyp))) != '\\' {
					goto do_not_strip_quotes
				}
				fallthrough
			default:
				if yyres != nil {
					yyres[yyn] = *yyp
				}
				yyn++
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

func IniYysyntaxError(yymsg_alloc *YYSIZE_T, yymsg **byte, yyssp *yytype_int16, yytoken int) int {
	var yysize0 YYSIZE_T = yytnamerr(YY_NULLPTR, yytname[yytoken])
	var yysize YYSIZE_T = yysize0
	const YYERROR_VERBOSE_ARGS_MAXIMUM = 5

	/* Internationalized format string. */

	var yyformat *byte = YY_NULLPTR

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

	if yytoken != YYEMPTY {
		var yyn int = yypact[*yyssp]
		yyarg[b.PostInc(&yycount)] = yytname[yytoken]
		if !(yypact_value_is_default(yyn)) {

			/* Start YYX at -YYN if negative to avoid negative indexes in
			   YYCHECK.  In other words, skip the first -YYN actions for
			   this state because they are default actions.  */

			var yyxbegin int = b.Cond(yyn < 0, -yyn, 0)

			/* Stay within bounds of both yycheck and yytname.  */

			var yychecklim int = YYLAST - yyn + 1
			var yyxend int = b.Cond(yychecklim < YYNTOKENS, yychecklim, YYNTOKENS)
			var yyx int
			for yyx = yyxbegin; yyx < yyxend; yyx++ {
				if yycheck[yyx+yyn] == yyx && yyx != YYTERROR && !(yytable_value_is_error(yytable[yyx+yyn])) {
					if yycount == YYERROR_VERBOSE_ARGS_MAXIMUM {
						yycount = 1
						yysize = yysize0
						break
					}
					yyarg[b.PostInc(&yycount)] = yytname[yyx]
					var yysize1 YYSIZE_T = yysize + yytnamerr(YY_NULLPTR, yytname[yyx])
					if !(yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM) {
						return 2
					}
					yysize = yysize1
				}
			}
		}
	}
	switch yycount {
	case 0:
		yyformat = YY_("syntax error")
	case 1:
		yyformat = YY_("syntax error, unexpected %s")
	case 2:
		yyformat = YY_("syntax error, unexpected %s, expecting %s")
	case 3:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s")
	case 4:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s or %s")
	case 5:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s or %s or %s")
	}
	var yysize1 YYSIZE_T = yysize + yystrlen(yyformat)
	if !(yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM) {
		return 2
	}
	yysize = yysize1
	if (*yymsg_alloc) < yysize {
		*yymsg_alloc = 2 * yysize
		if !(yysize <= (*yymsg_alloc) && (*yymsg_alloc) <= YYSTACK_ALLOC_MAXIMUM) {
			*yymsg_alloc = YYSTACK_ALLOC_MAXIMUM
		}
		return 1
	}

	/* Avoid sprintf, as that infringes on the user's name space.
	   Don't have undefined behavior even if the translation
	   produced a string with the wrong number of "%s"s.  */

	var yyp *byte = *yymsg
	var yyi int = 0
	for b.Assign(&(*yyp), *yyformat) != '0' {
		if (*yyp) == '%' && yyformat[1] == 's' && yyi < yycount {
			yyp += yytnamerr(yyp, yyarg[b.PostInc(&yyi)])
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

func IniYydestruct(yymsg *byte, yytype int, yyvaluep *types.Zval) {
	YYUSE(yyvaluep)
	if yymsg == nil {
		yymsg = "Deleting"
	}
	switch yytype {
	case 4:
		ZvalIniDtor(&(*yyvaluep))
	case 5:
		ZvalIniDtor(&(*yyvaluep))
	case 6:
		ZvalIniDtor(&(*yyvaluep))
	case 7:
		ZvalIniDtor(&(*yyvaluep))
	case 8:
		ZvalIniDtor(&(*yyvaluep))
	case 9:
		ZvalIniDtor(&(*yyvaluep))
	case 10:
		ZvalIniDtor(&(*yyvaluep))
	case 12:
		ZvalIniDtor(&(*yyvaluep))
	case 14:
		ZvalIniDtor(&(*yyvaluep))
	case 15:
		ZvalIniDtor(&(*yyvaluep))
	case 16:
		ZvalIniDtor(&(*yyvaluep))
	case 47:
		ZvalIniDtor(&(*yyvaluep))
	case 48:
		ZvalIniDtor(&(*yyvaluep))
	case 49:
		ZvalIniDtor(&(*yyvaluep))
	case 50:
		ZvalIniDtor(&(*yyvaluep))
	case 51:
		ZvalIniDtor(&(*yyvaluep))
	case 52:
		ZvalIniDtor(&(*yyvaluep))
	case 53:
		ZvalIniDtor(&(*yyvaluep))
	case 54:
		ZvalIniDtor(&(*yyvaluep))
	case 55:
		ZvalIniDtor(&(*yyvaluep))
	case 56:
		ZvalIniDtor(&(*yyvaluep))
	default:

	}
}

/*----------.
| yyparse.  |
`----------*/
