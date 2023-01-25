// <<generate>>

package zend

import (
	b "sik/builtin"
)

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

type YYSTYPE = int

const YYSTYPE_IS_TRIVIAL = 1
const YYSTYPE_IS_DECLARED = 1

var IniParse func() int

const YYBISON = 1
const YYBISON_VERSION = "3.0.2"
const YYSKELETON_NAME = "yacc.c"
const YYPURE = 2
const YYPUSH = 0
const YYPULL = 1
const IniYyparse = IniParse
const IniYylex = IniLex
const IniYyerror = IniError
const IniYydebug = ini_debug
const IniYynerrs = ini_nerrs
const DEBUG_CFG_PARSER = 0
const ZEND_INI_PARSER_CB = CompilerGlobals.GetIniParserParam().GetIniParserCb()
const ZEND_INI_PARSER_ARG = CompilerGlobals.GetIniParserParam().GetArg()
const ZEND_SYSTEM_INI = CompilerGlobals.GetIniParserUnbufferedErrors()
const YY_NULLPTR = 0
const YYERROR_VERBOSE = 1
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

type IniYytypeUint8 = uint8
type IniYytypeInt8 = signed__char
type IniYytypeUint16 = unsigned__short__int
type IniYytypeInt16 = short__int

const YYSIZE_T = size_t
const YYSIZE_MAXIMUM = YYSIZE_T - 1
const YYSTACK_ALLOC = YYMALLOC
const YYSTACK_FREE = YYFREE
const YYSTACK_ALLOC_MAXIMUM YYSIZE_T = YYSIZE_MAXIMUM
const YYMALLOC = Malloc

var Malloc func(YYSIZE_T) any

const YYFREE = Free

var Free func(any)

const YYSTACK_GAP_MAXIMUM = b.SizeOf("union yyalloc") - 1
const YYCOPY_NEEDED = 1
const YYFINAL = 2
const YYLAST = 123
const YYNTOKENS = 44
const YYNNTS = 13
const YYNRULES = 50
const YYNSTATES = 72
const YYUNDEFTOK = 2
const YYMAXUTOK = 272

var IniYytranslate []yytype_uint8 = []yytype_uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 40, 22, 2, 30, 29, 39, 23, 42, 43, 28, 25, 20, 26, 21, 27, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 19, 2, 32, 18, 33, 34, 35, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 41, 24, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 36, 38, 37, 31, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
var IniYytname []*byte = []*byte{"\"end of file\"", "error", "$undefined", "TC_SECTION", "TC_RAW", "TC_CONSTANT", "TC_NUMBER", "TC_STRING", "TC_WHITESPACE", "TC_LABEL", "TC_OFFSET", "TC_DOLLAR_CURLY", "TC_VARNAME", "TC_QUOTED_STRING", "BOOL_TRUE", "BOOL_FALSE", "NULL_NULL", "END_OF_LINE", "'='", "':'", "','", "'.'", "'\"'", "'\\''", "'^'", "'+'", "'-'", "'/'", "'*'", "'%'", "'$'", "'~'", "'<'", "'>'", "'?'", "'@'", "'{'", "'}'", "'|'", "'&'", "'!'", "']'", "'('", "')'", "$accept", "statement_list", "statement", "section_string_or_value", "string_or_value", "option_offset", "encapsed_list", "var_string_list_section", "var_string_list", "expr", "cfg_var_ref", "constant_literal", "constant_string", YY_NULLPTR}

const YYPACT_NINF = -25
const YYTABLE_NINF = -1

var IniYypact []yytype_int8 = []yytype_int8{-25, 9, -25, 73, -17, 81, -25, -25, -25, -25, -25, -25, -25, 15, -25, -20, 93, -25, -25, 0, -25, -25, -25, -25, -25, -25, -12, 101, -25, -25, -7, 36, -25, -25, -25, -25, -25, -25, -25, -25, 28, 28, 28, -25, 101, -1, 40, 30, -25, -25, -25, -25, -25, -25, -25, 80, -25, -25, 33, 28, 28, 28, -25, 0, 100, -25, -25, -25, -25, -25, -25, -25}
var IniYydefact []yytype_uint8 = []yytype_uint8{3, 0, 1, 10, 7, 17, 8, 2, 42, 41, 43, 44, 45, 0, 20, 0, 9, 21, 22, 0, 47, 46, 48, 49, 50, 20, 0, 16, 27, 28, 0, 0, 4, 20, 24, 25, 12, 13, 14, 15, 0, 0, 0, 5, 33, 11, 0, 0, 20, 30, 31, 40, 19, 23, 18, 0, 37, 38, 0, 0, 0, 0, 29, 0, 0, 26, 39, 36, 34, 35, 6, 32}
var IniYypgoto []yytype_int8 = []yytype_int8{-25, -25, -25, -25, -9, -25, -23, -25, 50, 4, -3, 44, -24}
var IniYydefgoto []yytype_int8 = []yytype_int8{-1, 1, 7, 15, 43, 26, 31, 16, 44, 45, 28, 18, 29}
var IniYytable []yytype_uint8 = []yytype_uint8{17, 19, 46, 50, 20, 21, 22, 23, 24, 2, 55, 13, 3, 34, 36, 37, 38, 39, 4, 5, 50, 32, 25, 59, 49, 64, 6, 30, 54, 47, 51, 40, 20, 21, 22, 23, 24, 60, 61, 13, 41, 49, 42, 54, 56, 57, 58, 13, 63, 52, 25, 13, 54, 52, 70, 27, 0, 59, 53, 40, 35, 54, 62, 67, 68, 69, 0, 0, 41, 0, 42, 60, 61, 0, 0, 0, 66, 8, 9, 10, 11, 12, 0, 0, 13, 20, 21, 22, 23, 24, 0, 13, 13, 52, 0, 14, 0, 8, 9, 10, 11, 12, 65, 25, 13, 20, 21, 22, 23, 24, 0, 13, 13, 52, 0, 33, 0, 0, 0, 0, 0, 0, 71, 48}
var IniYycheck []yytype_int8 = []yytype_int8{3, 18, 25, 27, 4, 5, 6, 7, 8, 0, 33, 11, 3, 16, 14, 15, 16, 17, 9, 10, 44, 41, 22, 24, 27, 48, 17, 12, 31, 41, 37, 31, 4, 5, 6, 7, 8, 38, 39, 11, 40, 44, 42, 46, 40, 41, 42, 11, 18, 13, 22, 11, 55, 13, 63, 5, -1, 24, 22, 31, 16, 64, 22, 59, 60, 61, -1, -1, 40, -1, 42, 38, 39, -1, -1, -1, 43, 4, 5, 6, 7, 8, -1, -1, 11, 4, 5, 6, 7, 8, -1, 11, 11, 13, -1, 22, -1, 4, 5, 6, 7, 8, 22, 22, 11, 4, 5, 6, 7, 8, -1, 11, 11, 13, -1, 22, -1, -1, -1, -1, -1, -1, 22, 22}
var IniYystos []yytype_uint8 = []yytype_uint8{0, 45, 0, 3, 9, 10, 17, 46, 4, 5, 6, 7, 8, 11, 22, 47, 51, 54, 55, 18, 4, 5, 6, 7, 8, 22, 49, 52, 54, 56, 12, 50, 41, 22, 54, 55, 14, 15, 16, 17, 31, 40, 42, 48, 52, 53, 50, 41, 22, 54, 56, 37, 13, 22, 54, 50, 53, 53, 53, 24, 38, 39, 22, 18, 50, 22, 43, 53, 53, 53, 48, 22}
var IniYyr1 []yytype_uint8 = []yytype_uint8{0, 44, 45, 45, 46, 46, 46, 46, 46, 47, 47, 48, 48, 48, 48, 48, 49, 49, 50, 50, 50, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 53, 53, 53, 53, 53, 53, 53, 54, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56}
var IniYyr2 []yytype_uint8 = []yytype_uint8{0, 2, 2, 0, 3, 3, 5, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 2, 2, 0, 1, 1, 3, 2, 2, 4, 1, 1, 3, 2, 2, 4, 1, 3, 3, 3, 2, 2, 3, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

const IniYyerrok = b.Assign(&yyerrstatus, 0)
const IniYyclearin = b.Assign(&yychar, YYEMPTY)
const YYEMPTY = -2
const YYEOF = 0
const YYTERROR = 1
const YYERRCODE = 256
const YYINITDEPTH = 200
const YYMAXDEPTH = 10000
