// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_ini_parser.h>

/* A Bison parser, made by GNU Bison 3.0.2.  */

/* Debug traces.  */

/* Token type.  */

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

// Source: <Zend/zend_ini_parser.c>

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
const ZEND_INI_PARSER_CB = __CG().GetIniParserParam().GetIniParserCb()
const ZEND_INI_PARSER_ARG = __CG().GetIniParserParam().GetArg()
const ZEND_SYSTEM_INI = __CG().GetIniParserUnbufferedErrors()

/* {{{ zend_ini_do_op()
 */

const YY_NULLPTR = 0

/* Enabling verbose error messages.  */

const YYERROR_VERBOSE = 1

/* In a future release of Bison, this section will be replaced
   by #include "zend_ini_parser.h".  */

/* Debug traces.  */

/* Token type.  */

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

const YYSIZE_T = size_t
const YYSIZE_MAXIMUM = YYSIZE_T - 1

/* Suppress unused-variable warnings by "using" E.  */

/* The parser invokes alloca or malloc; define the __special__  necessary symbols.  */

const YYSTACK_ALLOC = YYMALLOC
const YYSTACK_FREE = YYFREE
const YYSTACK_ALLOC_MAXIMUM YYSIZE_T = YYSIZE_MAXIMUM
const YYMALLOC = Malloc

var Malloc func(YYSIZE_T) any

const YYFREE = Free

var Free func(any)

/* A type that is properly aligned for any stack member.  */

/* The size of the maximum gap between one aligned stack and the next.  */

const YYSTACK_GAP_MAXIMUM = b.SizeOf("union yyalloc") - 1

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */

const YYCOPY_NEEDED = 1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */

/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */

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

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, without out-of-bounds checking.  */

var IniYytranslate []yytype_uint8 = []yytype_uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 40, 22, 2, 30, 29, 39, 23, 42, 43, 28, 25, 20, 26, 21, 27, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 19, 2, 32, 18, 33, 34, 35, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 41, 24, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 36, 38, 37, 31, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */

var IniYytname []*byte = []*byte{"\"end of file\"", "error", "$undefined", "TC_SECTION", "TC_RAW", "TC_CONSTANT", "TC_NUMBER", "TC_STRING", "TC_WHITESPACE", "TC_LABEL", "TC_OFFSET", "TC_DOLLAR_CURLY", "TC_VARNAME", "TC_QUOTED_STRING", "BOOL_TRUE", "BOOL_FALSE", "NULL_NULL", "END_OF_LINE", "'='", "':'", "','", "'.'", "'\"'", "'\\''", "'^'", "'+'", "'-'", "'/'", "'*'", "'%'", "'$'", "'~'", "'<'", "'>'", "'?'", "'@'", "'{'", "'}'", "'|'", "'&'", "'!'", "']'", "'('", "')'", "$accept", "statement_list", "statement", "section_string_or_value", "string_or_value", "option_offset", "encapsed_list", "var_string_list_section", "var_string_list", "expr", "cfg_var_ref", "constant_literal", "constant_string", YY_NULLPTR}

const YYPACT_NINF = -25
const YYTABLE_NINF = -1

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

/* Error token number */

const YYTERROR = 1
const YYERRCODE = 256

/* Enable debugging if requested.  */

/* YYINITDEPTH -- initial size of the parser's stacks.  */

const YYINITDEPTH = 200

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

const YYMAXDEPTH = 10000

/* Return the length of YYSTR.  */

/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */

/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for yyerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from yytname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */

/* Copy into *YYMSG, which is of size *YYMSG_ALLOC, an error message
   about the unexpected token YYTOKEN for the state stack whose top is
   YYSSP.

   Return 0 if *YYMSG was successfully written.  Return 1 if *YYMSG is
   not large enough to hold the message.  In that case, also set
   *YYMSG_ALLOC to the required number of bytes.  Return 2 if the
   required number of bytes is too large to store.  */

/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

/*----------.
| yyparse.  |
`----------*/
