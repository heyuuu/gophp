// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_language_parser.h>

/* A Bison parser, made by GNU Bison 3.0.2.  */

// #define YY_ZEND_ZEND_ZEND_LANGUAGE_PARSER_H_INCLUDED

/* Debug traces.  */

// #define YYDEBUG       0

/* "%code requires" blocks.  */

// # include "zend.h"

// #define YYTOKENTYPE

const (
	END                        = 0
	PREC_ARROW_FUNCTION        = 258
	T_INCLUDE                  = 259
	T_INCLUDE_ONCE             = 260
	T_REQUIRE                  = 261
	T_REQUIRE_ONCE             = 262
	T_LOGICAL_OR               = 263
	T_LOGICAL_XOR              = 264
	T_LOGICAL_AND              = 265
	T_PRINT                    = 266
	T_YIELD                    = 267
	T_DOUBLE_ARROW             = 268
	T_YIELD_FROM               = 269
	T_PLUS_EQUAL               = 270
	T_MINUS_EQUAL              = 271
	T_MUL_EQUAL                = 272
	T_DIV_EQUAL                = 273
	T_CONCAT_EQUAL             = 274
	T_MOD_EQUAL                = 275
	T_AND_EQUAL                = 276
	T_OR_EQUAL                 = 277
	T_XOR_EQUAL                = 278
	T_SL_EQUAL                 = 279
	T_SR_EQUAL                 = 280
	T_POW_EQUAL                = 281
	T_COALESCE_EQUAL           = 282
	T_COALESCE                 = 283
	T_BOOLEAN_OR               = 284
	T_BOOLEAN_AND              = 285
	T_IS_EQUAL                 = 286
	T_IS_NOT_EQUAL             = 287
	T_IS_IDENTICAL             = 288
	T_IS_NOT_IDENTICAL         = 289
	T_SPACESHIP                = 290
	T_IS_SMALLER_OR_EQUAL      = 291
	T_IS_GREATER_OR_EQUAL      = 292
	T_SL                       = 293
	T_SR                       = 294
	T_INSTANCEOF               = 295
	T_INT_CAST                 = 296
	T_DOUBLE_CAST              = 297
	T_STRING_CAST              = 298
	T_ARRAY_CAST               = 299
	T_OBJECT_CAST              = 300
	T_BOOL_CAST                = 301
	T_UNSET_CAST               = 302
	T_POW                      = 303
	T_NEW                      = 304
	T_CLONE                    = 305
	T_NOELSE                   = 306
	T_ELSEIF                   = 307
	T_ELSE                     = 308
	T_LNUMBER                  = 309
	T_DNUMBER                  = 310
	T_STRING                   = 311
	T_VARIABLE                 = 312
	T_INLINE_HTML              = 313
	T_ENCAPSED_AND_WHITESPACE  = 314
	T_CONSTANT_ENCAPSED_STRING = 315
	T_STRING_VARNAME           = 316
	T_NUM_STRING               = 317
	T_EVAL                     = 318
	T_INC                      = 319
	T_DEC                      = 320
	T_EXIT                     = 321
	T_IF                       = 322
	T_ENDIF                    = 323
	T_ECHO                     = 324
	T_DO                       = 325
	T_WHILE                    = 326
	T_ENDWHILE                 = 327
	T_FOR                      = 328
	T_ENDFOR                   = 329
	T_FOREACH                  = 330
	T_ENDFOREACH               = 331
	T_DECLARE                  = 332
	T_ENDDECLARE               = 333
	T_AS                       = 334
	T_SWITCH                   = 335
	T_ENDSWITCH                = 336
	T_CASE                     = 337
	T_DEFAULT                  = 338
	T_BREAK                    = 339
	T_CONTINUE                 = 340
	T_GOTO                     = 341
	T_FUNCTION                 = 342
	T_FN                       = 343
	T_CONST                    = 344
	T_RETURN                   = 345
	T_TRY                      = 346
	T_CATCH                    = 347
	T_FINALLY                  = 348
	T_THROW                    = 349
	T_USE                      = 350
	T_INSTEADOF                = 351
	T_GLOBAL                   = 352
	T_STATIC                   = 353
	T_ABSTRACT                 = 354
	T_FINAL                    = 355
	T_PRIVATE                  = 356
	T_PROTECTED                = 357
	T_PUBLIC                   = 358
	T_VAR                      = 359
	T_UNSET                    = 360
	T_ISSET                    = 361
	T_EMPTY                    = 362
	T_HALT_COMPILER            = 363
	T_CLASS                    = 364
	T_TRAIT                    = 365
	T_INTERFACE                = 366
	T_EXTENDS                  = 367
	T_IMPLEMENTS               = 368
	T_OBJECT_OPERATOR          = 369
	T_LIST                     = 370
	T_ARRAY                    = 371
	T_CALLABLE                 = 372
	T_LINE                     = 373
	T_FILE                     = 374
	T_DIR                      = 375
	T_CLASS_C                  = 376
	T_TRAIT_C                  = 377
	T_METHOD_C                 = 378
	T_FUNC_C                   = 379
	T_COMMENT                  = 380
	T_DOC_COMMENT              = 381
	T_OPEN_TAG                 = 382
	T_OPEN_TAG_WITH_ECHO       = 383
	T_CLOSE_TAG                = 384
	T_WHITESPACE               = 385
	T_START_HEREDOC            = 386
	T_END_HEREDOC              = 387
	T_DOLLAR_OPEN_CURLY_BRACES = 388
	T_CURLY_OPEN               = 389
	T_PAAMAYIM_NEKUDOTAYIM     = 390
	T_NAMESPACE                = 391
	T_NS_C                     = 392
	T_NS_SEPARATOR             = 393
	T_ELLIPSIS                 = 394
	T_BAD_CHARACTER            = 395
	T_ERROR                    = 396
)

/* Value type.  */

// Source: <Zend/zend_language_parser.c>

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

// #define yyparse       zendparse

// #define yylex       zendlex

// #define yyerror       zenderror

// #define yydebug       zenddebug

// #define yynerrs       zendnerrs

/* Copy the first part of user declarations.  */

// # include "zend_compile.h"

// # include "zend.h"

// # include "zend_list.h"

// # include "zend_globals.h"

// # include "zend_API.h"

// # include "zend_constants.h"

// # include "zend_language_scanner.h"

// # include "zend_exceptions.h"

// #define YYSIZE_T       size_t

// #define yytnamerr       zend_yytnamerr

// #define YYSTYPE       zend_parser_stack_elem

// #define YY_NULLPTR       0

/* Enabling verbose error messages.  */

// #define YYERROR_VERBOSE       1

/* In a future release of Bison, this section will be replaced
   by #include "zend_language_parser.h".  */

// #define YY_ZEND_ZEND_ZEND_LANGUAGE_PARSER_H_INCLUDED

/* Debug traces.  */

// #define YYDEBUG       0

/* "%code requires" blocks.  */

// #define YYTOKENTYPE

const (
	END                        = 0
	PREC_ARROW_FUNCTION        = 258
	T_INCLUDE                  = 259
	T_INCLUDE_ONCE             = 260
	T_REQUIRE                  = 261
	T_REQUIRE_ONCE             = 262
	T_LOGICAL_OR               = 263
	T_LOGICAL_XOR              = 264
	T_LOGICAL_AND              = 265
	T_PRINT                    = 266
	T_YIELD                    = 267
	T_DOUBLE_ARROW             = 268
	T_YIELD_FROM               = 269
	T_PLUS_EQUAL               = 270
	T_MINUS_EQUAL              = 271
	T_MUL_EQUAL                = 272
	T_DIV_EQUAL                = 273
	T_CONCAT_EQUAL             = 274
	T_MOD_EQUAL                = 275
	T_AND_EQUAL                = 276
	T_OR_EQUAL                 = 277
	T_XOR_EQUAL                = 278
	T_SL_EQUAL                 = 279
	T_SR_EQUAL                 = 280
	T_POW_EQUAL                = 281
	T_COALESCE_EQUAL           = 282
	T_COALESCE                 = 283
	T_BOOLEAN_OR               = 284
	T_BOOLEAN_AND              = 285
	T_IS_EQUAL                 = 286
	T_IS_NOT_EQUAL             = 287
	T_IS_IDENTICAL             = 288
	T_IS_NOT_IDENTICAL         = 289
	T_SPACESHIP                = 290
	T_IS_SMALLER_OR_EQUAL      = 291
	T_IS_GREATER_OR_EQUAL      = 292
	T_SL                       = 293
	T_SR                       = 294
	T_INSTANCEOF               = 295
	T_INT_CAST                 = 296
	T_DOUBLE_CAST              = 297
	T_STRING_CAST              = 298
	T_ARRAY_CAST               = 299
	T_OBJECT_CAST              = 300
	T_BOOL_CAST                = 301
	T_UNSET_CAST               = 302
	T_POW                      = 303
	T_NEW                      = 304
	T_CLONE                    = 305
	T_NOELSE                   = 306
	T_ELSEIF                   = 307
	T_ELSE                     = 308
	T_LNUMBER                  = 309
	T_DNUMBER                  = 310
	T_STRING                   = 311
	T_VARIABLE                 = 312
	T_INLINE_HTML              = 313
	T_ENCAPSED_AND_WHITESPACE  = 314
	T_CONSTANT_ENCAPSED_STRING = 315
	T_STRING_VARNAME           = 316
	T_NUM_STRING               = 317
	T_EVAL                     = 318
	T_INC                      = 319
	T_DEC                      = 320
	T_EXIT                     = 321
	T_IF                       = 322
	T_ENDIF                    = 323
	T_ECHO                     = 324
	T_DO                       = 325
	T_WHILE                    = 326
	T_ENDWHILE                 = 327
	T_FOR                      = 328
	T_ENDFOR                   = 329
	T_FOREACH                  = 330
	T_ENDFOREACH               = 331
	T_DECLARE                  = 332
	T_ENDDECLARE               = 333
	T_AS                       = 334
	T_SWITCH                   = 335
	T_ENDSWITCH                = 336
	T_CASE                     = 337
	T_DEFAULT                  = 338
	T_BREAK                    = 339
	T_CONTINUE                 = 340
	T_GOTO                     = 341
	T_FUNCTION                 = 342
	T_FN                       = 343
	T_CONST                    = 344
	T_RETURN                   = 345
	T_TRY                      = 346
	T_CATCH                    = 347
	T_FINALLY                  = 348
	T_THROW                    = 349
	T_USE                      = 350
	T_INSTEADOF                = 351
	T_GLOBAL                   = 352
	T_STATIC                   = 353
	T_ABSTRACT                 = 354
	T_FINAL                    = 355
	T_PRIVATE                  = 356
	T_PROTECTED                = 357
	T_PUBLIC                   = 358
	T_VAR                      = 359
	T_UNSET                    = 360
	T_ISSET                    = 361
	T_EMPTY                    = 362
	T_HALT_COMPILER            = 363
	T_CLASS                    = 364
	T_TRAIT                    = 365
	T_INTERFACE                = 366
	T_EXTENDS                  = 367
	T_IMPLEMENTS               = 368
	T_OBJECT_OPERATOR          = 369
	T_LIST                     = 370
	T_ARRAY                    = 371
	T_CALLABLE                 = 372
	T_LINE                     = 373
	T_FILE                     = 374
	T_DIR                      = 375
	T_CLASS_C                  = 376
	T_TRAIT_C                  = 377
	T_METHOD_C                 = 378
	T_FUNC_C                   = 379
	T_COMMENT                  = 380
	T_DOC_COMMENT              = 381
	T_OPEN_TAG                 = 382
	T_OPEN_TAG_WITH_ECHO       = 383
	T_CLOSE_TAG                = 384
	T_WHITESPACE               = 385
	T_START_HEREDOC            = 386
	T_END_HEREDOC              = 387
	T_DOLLAR_OPEN_CURLY_BRACES = 388
	T_CURLY_OPEN               = 389
	T_PAAMAYIM_NEKUDOTAYIM     = 390
	T_NAMESPACE                = 391
	T_NS_C                     = 392
	T_NS_SEPARATOR             = 393
	T_ELLIPSIS                 = 394
	T_BAD_CHARACTER            = 395
	T_ERROR                    = 396
)

/* Value type.  */

/* Copy the second part of user declarations.  */

type LangYytypeUint8 = uint8
type LangYytypeInt8 = signed__char
type LangYytypeUint16 = unsigned__short__int
type LangYytypeInt16 = short__int

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

// #define YYFREE       free

/* A type that is properly aligned for any stack member.  */

// @type LangUnionYyalloc struct

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

// #define YYFINAL       3

/* YYLAST -- Last index in YYTABLE.  */

// #define YYLAST       7174

/* YYNTOKENS -- Number of terminals.  */

// #define YYNTOKENS       170

/* YYNNTS -- Number of nonterminals.  */

// #define YYNNTS       140

/* YYNRULES -- Number of rules.  */

// #define YYNRULES       504

/* YYNSTATES -- Number of states.  */

// #define YYNSTATES       955

/* YYTRANSLATE[YYX] -- Symbol number corresponding to YYX as returned
   by yylex, with out-of-bounds checking.  */

// #define YYUNDEFTOK       2

// #define YYMAXUTOK       396

// #define YYTRANSLATE(YYX) ( ( unsigned int ) ( YYX ) <= YYMAXUTOK ? yytranslate [ YYX ] : YYUNDEFTOK )

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, without out-of-bounds checking.  */

var LangYytranslate []yytype_uint8 = []yytype_uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 54, 168, 2, 169, 53, 36, 2, 159, 160, 51, 48, 164, 49, 50, 52, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 30, 161, 42, 15, 44, 29, 64, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 165, 2, 166, 35, 2, 167, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 162, 34, 163, 56, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 31, 32, 33, 37, 38, 39, 40, 41, 43, 45, 46, 47, 55, 57, 58, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158}

/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */

var LangYytname []*byte = []*byte{"\"end of file\"", "error", "$undefined", "PREC_ARROW_FUNCTION", "\"include (T_INCLUDE)\"", "\"include_once (T_INCLUDE_ONCE)\"", "\"require (T_REQUIRE)\"", "\"require_once (T_REQUIRE_ONCE)\"", "\"or (T_LOGICAL_OR)\"", "\"xor (T_LOGICAL_XOR)\"", "\"and (T_LOGICAL_AND)\"", "\"print (T_PRINT)\"", "\"yield (T_YIELD)\"", "\"=> (T_DOUBLE_ARROW)\"", "\"yield from (T_YIELD_FROM)\"", "'='", "\"+= (T_PLUS_EQUAL)\"", "\"-= (T_MINUS_EQUAL)\"", "\"*= (T_MUL_EQUAL)\"", "\"/= (T_DIV_EQUAL)\"", "\".= (T_CONCAT_EQUAL)\"", "\"%= (T_MOD_EQUAL)\"", "\"&= (T_AND_EQUAL)\"", "\"|= (T_OR_EQUAL)\"", "\"^= (T_XOR_EQUAL)\"", "\"<<= (T_SL_EQUAL)\"", "\">>= (T_SR_EQUAL)\"", "\"**= (T_POW_EQUAL)\"", "\"?" + "?= (T_COALESCE_EQUAL)\"", "'?'", "':'", "\"?? (T_COALESCE)\"", "\"|| (T_BOOLEAN_OR)\"", "\"&& (T_BOOLEAN_AND)\"", "'|'", "'^'", "'&'", "\"== (T_IS_EQUAL)\"", "\"!= (T_IS_NOT_EQUAL)\"", "\"=== (T_IS_IDENTICAL)\"", "\"!== (T_IS_NOT_IDENTICAL)\"", "\"<=> (T_SPACESHIP)\"", "'<'", "\"<= (T_IS_SMALLER_OR_EQUAL)\"", "'>'", "\">= (T_IS_GREATER_OR_EQUAL)\"", "\"<< (T_SL)\"", "\">> (T_SR)\"", "'+'", "'-'", "'.'", "'*'", "'/'", "'%'", "'!'", "\"instanceof (T_INSTANCEOF)\"", "'~'", "\"(int) (T_INT_CAST)\"", "\"(double) (T_DOUBLE_CAST)\"", "\"(string) (T_STRING_CAST)\"", "\"(array) (T_ARRAY_CAST)\"", "\"(object) (T_OBJECT_CAST)\"", "\"(bool) (T_BOOL_CAST)\"", "\"(unset) (T_UNSET_CAST)\"", "'@'", "\"** (T_POW)\"", "\"new (T_NEW)\"", "\"clone (T_CLONE)\"", "T_NOELSE", "\"elseif (T_ELSEIF)\"", "\"else (T_ELSE)\"", "\"integer number (T_LNUMBER)\"", "\"floating-point number (T_DNUMBER)\"", "\"identifier (T_STRING)\"", "\"variable (T_VARIABLE)\"", "T_INLINE_HTML", "\"quoted-string and whitespace (T_ENCAPSED_AND_WHITESPACE)\"", "\"quoted-string (T_CONSTANT_ENCAPSED_STRING)\"", "\"variable name (T_STRING_VARNAME)\"", "\"number (T_NUM_STRING)\"", "\"eval (T_EVAL)\"", "\"++ (T_INC)\"", "\"-- (T_DEC)\"", "\"exit (T_EXIT)\"", "\"if (T_IF)\"", "\"endif (T_ENDIF)\"", "\"echo (T_ECHO)\"", "\"do (T_DO)\"", "\"while (T_WHILE)\"", "\"endwhile (T_ENDWHILE)\"", "\"for (T_FOR)\"", "\"endfor (T_ENDFOR)\"", "\"foreach (T_FOREACH)\"", "\"endforeach (T_ENDFOREACH)\"", "\"declare (T_DECLARE)\"", "\"enddeclare (T_ENDDECLARE)\"", "\"as (T_AS)\"", "\"switch (T_SWITCH)\"", "\"endswitch (T_ENDSWITCH)\"", "\"case (T_CASE)\"", "\"default (T_DEFAULT)\"", "\"break (T_BREAK)\"", "\"continue (T_CONTINUE)\"", "\"goto (T_GOTO)\"", "\"function (T_FUNCTION)\"", "\"fn (T_FN)\"", "\"const (T_CONST)\"", "\"return (T_RETURN)\"", "\"try (T_TRY)\"", "\"catch (T_CATCH)\"", "\"finally (T_FINALLY)\"", "\"throw (T_THROW)\"", "\"use (T_USE)\"", "\"insteadof (T_INSTEADOF)\"", "\"global (T_GLOBAL)\"", "\"static (T_STATIC)\"", "\"abstract (T_ABSTRACT)\"", "\"final (T_FINAL)\"", "\"private (T_PRIVATE)\"", "\"protected (T_PROTECTED)\"", "\"public (T_PUBLIC)\"", "\"var (T_VAR)\"", "\"unset (T_UNSET)\"", "\"isset (T_ISSET)\"", "\"empty (T_EMPTY)\"", "\"__halt_compiler (T_HALT_COMPILER)\"", "\"class (T_CLASS)\"", "\"trait (T_TRAIT)\"", "\"interface (T_INTERFACE)\"", "\"extends (T_EXTENDS)\"", "\"implements (T_IMPLEMENTS)\"", "\"-> (T_OBJECT_OPERATOR)\"", "\"list (T_LIST)\"", "\"array (T_ARRAY)\"", "\"callable (T_CALLABLE)\"", "\"__LINE__ (T_LINE)\"", "\"__FILE__ (T_FILE)\"", "\"__DIR__ (T_DIR)\"", "\"__CLASS__ (T_CLASS_C)\"", "\"__TRAIT__ (T_TRAIT_C)\"", "\"__METHOD__ (T_METHOD_C)\"", "\"__FUNCTION__ (T_FUNC_C)\"", "\"comment (T_COMMENT)\"", "\"doc comment (T_DOC_COMMENT)\"", "\"open tag (T_OPEN_TAG)\"", "\"open tag with echo (T_OPEN_TAG_WITH_ECHO)\"", "\"close tag (T_CLOSE_TAG)\"", "\"whitespace (T_WHITESPACE)\"", "\"heredoc start (T_START_HEREDOC)\"", "\"heredoc end (T_END_HEREDOC)\"", "\"${ (T_DOLLAR_OPEN_CURLY_BRACES)\"", "\"{$ (T_CURLY_OPEN)\"", "\":: (T_PAAMAYIM_NEKUDOTAYIM)\"", "\"namespace (T_NAMESPACE)\"", "\"__NAMESPACE__ (T_NS_C)\"", "\"\\\\ (T_NS_SEPARATOR)\"", "\"... (T_ELLIPSIS)\"", "\"invalid character (T_BAD_CHARACTER)\"", "T_ERROR", "'('", "')'", "';'", "'{'", "'}'", "','", "'['", "']'", "'`'", "'\"'", "'$'", "$accept", "start", "reserved_non_modifiers", "semi_reserved", "identifier", "top_statement_list", "namespace_name", "name", "top_statement", "$@1", "$@2", "use_type", "group_use_declaration", "mixed_group_use_declaration", "possible_comma", "inline_use_declarations", "unprefixed_use_declarations", "use_declarations", "inline_use_declaration", "unprefixed_use_declaration", "use_declaration", "const_list", "inner_statement_list", "inner_statement", "statement", "$@3", "catch_list", "catch_name_list", "finally_statement", "unset_variables", "unset_variable", "function_declaration_statement", "is_reference", "is_variadic", "class_declaration_statement", "@4", "@5", "class_modifiers", "class_modifier", "trait_declaration_statement", "@6", "interface_declaration_statement", "@7", "extends_from", "interface_extends_list", "implements_list", "foreach_variable", "for_statement", "foreach_statement", "declare_statement", "switch_case_list", "case_list", "case_separator", "while_statement", "if_stmt_without_else", "if_stmt", "alt_if_stmt_without_else", "alt_if_stmt", "parameter_list", "non_empty_parameter_list", "parameter", "optional_type", "type_expr", "type", "return_type", "argument_list", "non_empty_argument_list", "argument", "global_var_list", "global_var", "static_var_list", "static_var", "class_statement_list", "class_statement", "name_list", "trait_adaptations", "trait_adaptation_list", "trait_adaptation", "trait_precedence", "trait_alias", "trait_method_reference", "absolute_trait_method_reference", "method_body", "variable_modifiers", "method_modifiers", "non_empty_member_modifiers", "member_modifier", "property_list", "property", "class_const_list", "class_const_decl", "const_decl", "echo_expr_list", "echo_expr", "for_exprs", "non_empty_for_exprs", "anonymous_class", "@8", "new_expr", "expr", "inline_function", "fn", "function", "backup_doc_comment", "backup_fn_flags", "backup_lex_pos", "returns_ref", "lexical_vars", "lexical_var_list", "lexical_var", "function_call", "class_name", "class_name_reference", "exit_expr", "backticks_expr", "ctor_arguments", "dereferencable_scalar", "scalar", "constant", "optional_expr", "variable_class_name", "dereferencable", "callable_expr", "callable_variable", "variable", "simple_variable", "static_member", "new_variable", "member_name", "property_name", "array_pair_list", "possible_array_pair", "non_empty_array_pair_list", "array_pair", "encaps_list", "encaps_var", "encaps_var_offset", "internal_functions_in_yacc", "isset_variables", "isset_variable", 0}

// #define YYPACT_NINF       - 753

// #define yypact_value_is_default(Yystate) ( ! ! ( ( Yystate ) == ( - 753 ) ) )

// #define YYTABLE_NINF       - 477

// #define yytable_value_is_error(Yytable_value) ( ! ! ( ( Yytable_value ) == ( - 477 ) ) )

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */

var LangYypact []yytype_int16 = []yytype_int16{-753, 100, 1156, -753, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 206, 5216, -753, -753, 90, -753, -753, -753, -29, 351, 351, 23, 37, 5216, 4084, 84, 123, 158, 166, 171, 5216, 5216, 115, -753, -753, 264, 5216, 201, 5216, -5, 5, 209, -753, -753, 220, 239, 252, 255, -753, -753, -753, 263, 268, -753, -753, -753, -753, -753, -753, -753, 234, 25, -753, 298, 5216, -753, -753, 4222, 327, 361, -16, 238, -55, -753, -753, -753, -753, -23, -753, -753, -753, 399, -753, 304, -753, -753, 5916, -753, 397, 397, -753, 321, 205, -753, 316, 333, 60, 330, 338, 4567, -753, -753, -753, -753, 404, 348, 6958, 397, 6958, 6958, 6958, 5440, 7046, 5440, 440, 440, 47, 440, 440, 440, 440, 440, 440, 440, 440, 440, -753, -753, -753, -753, 365, 330, -753, 235, -753, -753, 5216, 5216, 4222, 391, 316, 390, 390, 5216, -753, 5216, -12, -753, 6958, 463, 5216, 5216, 5216, 264, 5216, 6958, 392, 398, 413, 545, 182, -753, 419, -753, 5965, -753, -753, 298, -18, 56, 423, 191, -753, -753, 262, -753, -753, 553, 318, -753, -753, 351, 5216, 5216, 411, 509, 510, 512, 4222, 4222, -39, 231, -753, 4498, 351, 341, -753, 298, 174, 424, 238, 6013, 1294, 351, 429, 5216, 6862, 426, -753, 430, -753, 10, 428, 364, 10, 227, 5216, -753, 520, 4360, -753, -753, -753, 437, 4084, 439, 579, 450, 5216, 5216, 5216, 4664, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 233, 5216, -753, -753, 453, 540, 2978, 5216, 2978, 36, 5216, 5216, -753, 4802, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, 5216, -753, -753, -753, 5216, 330, 5, -753, -753, 36, 5, 5216, 5216, 6062, 6110, 454, 464, 6158, -753, 5216, 455, 6206, 466, 459, 6958, 6796, 70, 6254, -753, -753, -753, 5216, -753, 264, -753, 1432, -753, 76, -753, 552, -20, 298, 119, 467, 322, -753, -753, 59, -753, 5, 5216, -753, 557, 468, -753, 198, 6958, 473, -753, 6302, 478, 514, -753, 515, 485, 486, 574, 421, -753, -753, -74, 5580, 488, -753, -753, -753, 238, -753, 487, -753, 438, 493, -753, -753, -753, -753, -753, -753, -753, 162, 4222, 6958, 4940, 638, 4222, -753, -753, 5628, -753, 5216, -753, 491, -753, 6958, 583, 5216, -753, 5216, -753, -753, 7005, 2962, 5440, 5216, 6910, 6648, 6504, 7078, 7109, 4270, 4407, 4711, 4711, 4711, 4711, 4711, 1008, 1008, 1008, 1008, 733, 733, 62, 62, 62, 47, 47, 47, -753, 440, 66, -753, 499, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, -753, 5216, -753, -753, 500, 501, 330, 504, 500, 501, 330, -753, 5216, -753, 330, 5676, 511, 351, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 5440, 514, -753, -753, -753, 5724, 513, -753, 506, -753, -753, 3282, -753, 5216, 3448, 5216, 5216, 51, -753, 13, 6958, -753, -753, -17, -753, 250, 122, -3, -753, -753, 298, 190, -753, -753, 6958, -753, 351, 507, 5216, 521, -753, -753, 74, 536, 523, 74, -753, 660, -753, -753, 607, -753, -753, -753, 525, -753, 5216, -753, -753, -753, 880, 532, 533, 351, 537, 6958, 5216, -753, -753, 6958, 5078, 535, 514, 6350, 6398, 1570, 6648, 5216, 276, -753, -753, -753, 538, 539, -753, 663, -753, -753, 542, 66, 5772, -753, -753, -753, 5820, -753, -753, -753, 390, 536, -753, -753, -753, -753, 6446, -753, -753, -753, 541, 6958, 351, 546, 4222, -6, 32, 3614, 548, 550, -753, -753, 422, 250, 298, 543, -753, -753, 1, 298, -753, 551, -753, -753, -753, 74, -753, -753, -753, 556, 544, 5216, -753, -753, 410, 1018, -753, 554, 660, 270, 4222, 5440, -753, -753, 536, 4084, 694, 564, 6648, -753, 696, 34, -753, 571, 66, 568, -753, -753, -753, 3116, 572, 1708, 5216, 33, 4222, 581, 51, 3780, -753, -753, -753, -753, 448, -753, -28, 575, 576, -753, 543, -753, 250, 585, 298, 588, -753, -753, 556, 587, 407, 74, -753, 5440, 600, -753, -753, 604, -753, -753, -753, -753, 34, -753, -753, -753, 695, 605, 656, 608, -753, 610, 613, 614, 35, 615, -753, -753, -753, 1846, 479, 617, 5216, 17, 22, -753, 74, -753, 616, -753, -753, 588, 298, 626, -753, 74, -753, -753, -753, -753, -753, -753, -753, -753, -753, 34, 300, 711, -753, -753, 418, -753, 660, 628, 3116, -753, 778, 777, 696, 635, 696, -753, -753, 3946, -753, 3780, 1984, 634, 643, -753, 5868, -753, -753, -753, -753, -753, 42, 2122, -753, 633, -753, -753, 624, -56, 725, 6524, 397, -753, -753, -753, -753, 5216, -753, 45, -753, 639, -753, -753, -753, -753, 644, -753, -753, -753, 3116, 74, 640, -753, -753, -753, -753, 5307, -753, 799, 335, -753, 804, 337, -753, 6524, 691, -753, 6958, 658, 747, -753, 212, -753, 661, -753, 2260, -753, 3116, -753, 671, 726, 348, -753, -753, 682, 5445, -753, 674, 675, 741, 727, 5216, -753, -753, 725, 5216, -753, 6524, -753, -753, 5216, -753, -753, -753, 45, -753, 678, -753, 6524, -753, -753, -753, -753, 6662, 74, 6958, -753, 6958, -753, 683, 6958, 2398, -753, 2536, -753, 2674, -753, -753, -753, 6524, 556, -753, -753, 66, -753, -753, -753, -753, -753, 681, -753, -753, 696, -753, 379, -753, -753, -753, 2812, -753, -753}

/* YYDEFACT[STATE-NUM] -- Default reduction number in state STATE-NUM.
   Performed when YYTABLE does not specify something else to do.  Zero
   means the default is an error.  */

var LangYydefact []yytype_uint16 = []yytype_uint16{81, 0, 2, 1, 0, 0, 0, 0, 0, 375, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 414, 415, 82, 451, 146, 413, 0, 0, 0, 404, 0, 0, 0, 0, 0, 0, 0, 0, 433, 433, 0, 384, 383, 0, 433, 0, 0, 0, 0, 400, 178, 179, 0, 0, 0, 0, 174, 180, 182, 0, 0, 416, 417, 418, 423, 419, 420, 421, 0, 96, 422, 0, 0, 153, 126, 469, 406, 0, 0, 84, 430, 80, 87, 88, 89, 0, 176, 90, 91, 213, 134, 0, 135, 358, 0, 379, 388, 388, 447, 0, 428, 372, 429, 0, 435, 0, 448, 303, 442, 449, 362, 82, 400, 0, 497, 388, 498, 500, 501, 374, 376, 378, 343, 344, 345, 346, 363, 364, 365, 366, 367, 368, 369, 371, 400, 299, 401, 302, 402, 409, 456, 403, 308, 157, 0, 0, 469, 438, 0, 323, 325, 433, 370, 0, 0, 293, 294, 0, 0, 295, 0, 0, 0, 434, 0, 0, 0, 0, 0, 124, 0, 126, 0, 103, 104, 0, 119, 0, 0, 0, 121, 116, 0, 241, 242, 245, 0, 244, 380, 0, 0, 0, 0, 0, 0, 0, 469, 469, 484, 0, 425, 0, 0, 0, 482, 0, 94, 0, 86, 0, 0, 0, 0, 0, 474, 0, 472, 468, 470, 407, 0, 408, 0, 0, 0, 453, 0, 0, 396, 172, 177, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 147, 389, 0, 385, 0, 433, 0, 0, 0, 433, 399, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 322, 324, 385, 0, 409, 0, 410, 301, 0, 0, 0, 433, 0, 0, 0, 0, 0, 145, 0, 0, 0, 0, 296, 298, 0, 0, 0, 140, 141, 156, 0, 102, 0, 142, 0, 155, 119, 122, 0, 0, 0, 119, 0, 0, 98, 100, 0, 143, 0, 0, 144, 0, 109, 164, 436, 504, 109, 502, 0, 0, 184, 385, 186, 0, 0, 0, 0, 424, 483, 0, 0, 436, 481, 427, 480, 85, 93, 0, 81, 357, 0, 133, 125, 127, 128, 129, 130, 131, 436, 469, 477, 0, 412, 469, 373, 426, 0, 83, 0, 234, 109, 236, 238, 0, 0, 214, 0, 126, 217, 328, 330, 329, 0, 0, 361, 326, 327, 331, 333, 332, 349, 350, 347, 348, 355, 351, 352, 353, 354, 341, 342, 335, 336, 334, 337, 339, 340, 356, 338, 225, 385, 0, 3, 4, 6, 7, 8, 9, 10, 46, 47, 11, 12, 13, 16, 17, 78, 5, 14, 15, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 49, 50, 51, 52, 53, 41, 42, 43, 70, 44, 45, 30, 31, 32, 33, 34, 35, 36, 72, 73, 74, 75, 76, 77, 37, 38, 39, 40, 61, 59, 60, 56, 57, 48, 54, 55, 66, 67, 68, 62, 63, 65, 64, 58, 69, 0, 71, 79, 431, 454, 0, 0, 432, 455, 0, 465, 0, 467, 450, 0, 0, 0, 306, 309, 310, 311, 313, 314, 315, 316, 317, 318, 319, 320, 312, 321, 377, 184, 460, 459, 461, 0, 0, 499, 437, 412, 405, 0, 292, 0, 0, 295, 0, 0, 151, 0, 385, 123, 158, 0, 120, 0, 119, 0, 99, 101, 0, 119, 115, 240, 246, 243, 110, 0, 110, 0, 496, 92, 0, 188, 0, 0, 385, 0, 411, 486, 0, 491, 494, 492, 0, 488, 0, 487, 490, 81, 0, 0, 0, 0, 0, 473, 0, 471, 452, 239, 110, 0, 184, 0, 0, 0, 360, 0, 0, 229, 230, 231, 0, 219, 221, 168, 226, 227, 0, 225, 0, 397, 444, 398, 0, 446, 445, 443, 307, 188, 458, 457, 126, 211, 0, 126, 209, 136, 0, 297, 0, 0, 469, 0, 436, 0, 204, 204, 139, 291, 162, 0, 0, 109, 112, 117, 0, 0, 165, 0, 503, 495, 185, 0, 385, 248, 253, 187, 0, 0, 493, 485, 0, 0, 97, 0, 479, 436, 469, 305, 237, 235, 188, 0, 0, 0, 359, 228, 232, 225, 169, 170, 225, 0, 463, 466, 385, 215, 0, 0, 295, 436, 469, 0, 0, 0, 126, 198, 152, 204, 0, 204, 0, 0, 0, 154, 109, 118, 110, 0, 0, 109, 114, 148, 189, 0, 274, 0, 248, 304, 0, 95, 132, 0, 385, 212, 126, 218, 0, 385, 222, 171, 0, 0, 390, 0, 137, 0, 0, 0, 412, 0, 126, 196, 149, 0, 0, 0, 0, 0, 0, 200, 0, 126, 0, 111, 107, 109, 110, 0, 248, 0, 281, 282, 283, 280, 279, 278, 273, 181, 247, 225, 0, 272, 276, 254, 274, 489, 478, 0, 216, 233, 0, 223, 232, 0, 232, 248, 210, 0, 192, 0, 0, 0, 0, 202, 0, 207, 208, 126, 201, 160, 0, 0, 108, 0, 113, 105, 274, 0, 0, 0, 388, 277, 183, 248, 386, 0, 386, 0, 386, 274, 126, 194, 138, 150, 0, 199, 203, 126, 206, 0, 0, 163, 106, 175, 255, 0, 251, 385, 0, 285, 0, 0, 289, 0, 274, 387, 224, 0, 0, 394, 0, 393, 0, 300, 0, 197, 205, 161, 0, 82, 58, 256, 267, 0, 0, 258, 0, 0, 0, 268, 0, 286, 249, 0, 0, 250, 0, 385, 173, 0, 126, 395, 391, 0, 126, 0, 126, 0, 257, 259, 260, 261, 0, 0, 385, 284, 385, 288, 0, 386, 0, 392, 0, 195, 0, 269, 263, 264, 266, 262, 287, 290, 225, 382, 386, 386, 159, 265, 0, 167, 381, 232, 386, 0, 270, 126, 386, 0, 252, 271}

/* YYPGOTO[NTERM-NUM].  */

var LangYypgoto []yytype_int16 = []yytype_int16{-753, -753, -73, -753, -239, -365, -1, -22, -753, -753, -753, 791, -753, -753, -313, 181, 114, 666, 118, -177, 516, 684, -174, -753, 4, -753, -753, -753, -753, -753, 273, 3, -753, -753, 27, -753, -753, -753, 763, 30, -753, 31, -753, -477, -753, -456, 138, -753, 44, -753, -753, -591, 40, -753, -753, -753, -753, -753, -628, -753, 154, 68, 109, 242, -743, -45, -753, 251, -753, 522, -753, 524, -687, -753, -640, -753, -753, -21, -753, -753, -753, -753, -753, -753, -753, -753, -752, -753, -32, -753, -30, 547, -753, 559, -519, -753, -753, -753, -753, 6, 26, -753, -2, -245, -707, -753, -68, -753, -753, -35, -753, 8, 601, -753, -753, 577, 29, -753, 108, 41, -753, -753, -753, -753, 144, 65, -753, -753, 599, 578, -142, 490, -753, -753, 462, 414, -753, -753, -753, 299}

/* YYDEFGOTO[NTERM-NUM].  */

var LangYydefgoto []yytype_int16 = []yytype_int16{-1, 1, 513, 514, 865, 2, 84, 85, 86, 372, 212, 665, 337, 183, 580, 666, 734, 184, 667, 185, 186, 173, 215, 377, 378, 658, 663, 825, 728, 347, 348, 379, 704, 755, 380, 398, 198, 90, 91, 381, 199, 382, 200, 586, 589, 677, 656, 847, 767, 721, 661, 723, 822, 650, 94, 95, 96, 97, 625, 626, 627, 628, 629, 630, 752, 301, 395, 396, 187, 188, 191, 192, 739, 793, 680, 861, 889, 890, 891, 892, 893, 894, 951, 794, 795, 796, 797, 863, 864, 866, 867, 174, 159, 160, 316, 317, 142, 299, 98, 99, 100, 101, 120, 436, 870, 904, 272, 809, 875, 876, 103, 104, 144, 157, 225, 302, 105, 106, 107, 169, 108, 109, 110, 111, 112, 113, 114, 146, 517, 525, 220, 221, 222, 223, 208, 209, 597, 115, 351, 352}

/* YYTABLE[YYPACT[STATE-NUM]] -- What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule whose
   number is the opposite.  If YYTABLE_NINF, syntax error.  */

var LangYytable []yytype_int16 = []yytype_int16{102, 141, 329, 332, 706, 88, 87, 717, 603, 309, 119, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 89, 147, 143, 92, 93, 273, 515, 737, 519, 582, 651, 233, 162, 161, 659, 836, -190, -191, 820, -193, 168, 168, 181, 297, 392, 799, 168, 392, 177, 28, 358, 359, 152, 152, 621, 841, 280, 843, 642, 116, 725, 392, 771, 772, 211, 392, 213, 854, 756, 333, 28, 193, 873, 614, 214, 203, 170, 219, 653, 145, 598, 175, 599, 360, 55, 56, 621, 831, -401, 116, 178, 3, 179, 268, 234, 232, 859, 860, 116, 740, 522, 28, 587, 269, 265, 266, 267, 855, 268, 189, 874, 148, 771, 772, 844, 116, 28, 361, 269, 30, 116, 149, 769, 116, 773, 872, 774, 877, 334, 695, 116, 153, 153, 568, 193, 832, 664, 229, 116, 230, 312, 180, 869, 313, 83, 718, 307, 308, 219, 332, 670, 206, 207, 168, 733, 311, 933, 139, 622, 623, 315, 318, 319, 333, 321, 83, 660, 154, 155, 821, 331, 210, 336, 156, 654, 65, 823, 709, 118, 171, 76, 631, 277, -190, -191, 761, -193, 158, 310, 523, 622, 623, 350, 353, 947, 118, 83, 76, 219, 219, 370, 150, 335, 365, 102, 573, 333, 655, 938, 333, 118, 83, 76, 278, 152, 385, 279, -220, 118, 618, 76, 561, 566, 944, 945, 327, 391, 152, 686, 397, 747, 948, 400, 605, 163, 953, 152, 141, 404, 405, 406, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 570, 433, 143, 669, 934, 116, 28, 168, 164, 190, 526, 168, 333, 529, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 203, 153, 367, 543, 203, 116, 28, 203, 943, 204, 548, 168, 46, 47, 153, 518, 165, 662, 161, 527, 139, -476, 116, 153, 166, -476, 102, -476, 231, 167, 563, 140, 145, 569, 371, -438, 172, 349, 516, 574, 520, 524, 326, 681, 231, 327, 549, 139, 116, 577, 366, 340, 732, 178, 341, 179, -438, -166, 118, 383, 76, -166, 176, -441, 545, 303, -438, 524, 547, -438, 116, 907, 238, 239, 83, 908, 206, 207, 194, 362, 206, 207, 205, 206, 207, 118, 304, 76, 240, 219, 668, 608, 231, 219, 390, 332, 305, 195, 612, 306, 203, 83, 224, 46, 616, 834, 617, 189, 622, 623, 196, 624, 619, 197, 203, 777, 367, 241, 242, 243, 782, 201, 342, 116, 28, 343, 202, 30, 118, -475, 76, 738, 271, -475, 203, -475, 227, 203, 244, 367, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 758, 268, 139, 828, 236, 237, 593, 710, 634, 274, 712, 269, 636, 206, 207, 345, 638, 275, 346, 572, 65, 276, 341, 668, 730, 232, 368, 206, 207, 735, 594, 595, 897, -439, 900, 898, 596, 901, 802, 210, 118, 269, 76, 805, 46, 47, 150, 206, 207, 716, 206, 207, 151, 300, 633, 784, 83, -436, 785, 786, 787, 788, 789, 790, 791, 637, 784, 726, 727, 785, 786, 787, 788, 789, 790, 791, 949, 950, -436, 226, 228, 768, 770, 771, 772, 746, -441, 314, -436, 322, 668, -436, 735, 152, 646, 323, 325, 649, 647, 675, 318, 652, 679, 574, 344, -437, 792, 354, 574, 762, 324, 803, 743, 817, 771, 772, 328, 837, 355, 356, 339, 357, 373, 350, 384, 152, -437, 815, 387, 392, 388, 389, 399, -440, 401, 624, -437, 102, 826, -437, 829, 685, 88, 87, 152, 402, 624, 403, 434, 435, 556, 692, 102, 896, 363, 397, 552, 887, 369, 559, 553, 567, 699, 558, 571, 902, 89, 190, 579, 92, 93, 152, 153, 581, 363, 584, 369, 363, 369, 585, 588, 590, 591, 592, 853, 602, 887, 601, 604, 609, 679, 613, 615, 923, 632, -462, -464, 219, 720, 574, 574, -440, 676, 672, 153, 574, 635, 879, 641, 930, 935, 682, 936, 640, 881, 644, 624, 674, 152, 624, 102, 678, 683, 153, 742, 88, 87, 684, 688, 689, 942, 694, 691, 219, 701, 703, 748, 705, 713, 702, 657, 715, 741, 731, 102, 722, 102, 724, 736, 89, 153, 745, 92, 93, 798, 318, 740, 219, 766, 349, 749, 750, 751, 754, 757, 624, 574, 925, 574, 759, 775, 927, 784, 929, 776, 785, 786, 787, 788, 789, 790, 791, 152, 763, 779, 783, 690, 784, 781, 824, 785, 786, 787, 788, 789, 790, 791, 153, 679, 800, 801, 807, 102, 868, 808, 806, 810, 811, 624, 812, 813, 814, 952, 819, 818, 827, 574, 262, 263, 264, 265, 266, 267, 858, 268, 830, 838, 839, 840, 835, 842, 850, 857, 714, 269, 862, 883, 102, 878, 784, 851, 880, 785, 786, 787, 788, 789, 790, 791, 102, 895, -275, 846, -275, 766, 899, 905, 906, -78, 909, 102, 153, 785, 786, 787, 788, 789, 790, 882, 911, 912, 915, 916, 917, 888, 928, 918, 946, 937, 182, 932, 729, 871, 780, 338, 778, 320, 102, 671, 235, 903, 764, 753, 575, 848, 852, 804, 657, 833, 700, 693, 576, 920, 888, 914, 432, 578, 922, 555, 926, 564, 521, 544, 102, 610, 102, 673, 546, 0, 0, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 679, 0, 0, 0, 0, 919, 0, 0, 0, 921, 0, 0, 0, 0, 924, 0, 0, 0, 0, 624, 0, 0, 0, 0, 0, 0, 0, 102, 0, 102, 0, 102, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 102, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 48, 49, 50, 0, 0, 51, 52, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 60, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 74, 75, 76, 0, 0, 0, 77, 0, 78, 79, 687, 0, 80, 0, 81, 82, 83, -477, -477, -477, -477, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 11, 12, 0, 0, 0, 0, 13, 269, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 48, 49, 50, 0, 0, 51, 52, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 60, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 74, 75, 76, 0, 0, 0, 77, 0, 78, 79, 744, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 48, 49, 50, 0, 0, 51, 52, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 60, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 74, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 376, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 565, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 698, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 760, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 816, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 849, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 856, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 910, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 939, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 940, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 941, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 243, 78, 79, 954, 0, 80, 0, 81, 82, 83, 437, 438, 439, 440, 441, 442, 443, 444, 445, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 446, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 447, 448, 0, 449, 450, 0, 0, 451, 28, 0, 0, 0, 0, 0, 452, 0, 0, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 0, 495, 496, 497, 498, 499, 0, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 510, 511, 0, 0, 0, 0, 0, 0, 0, 512, 0, 0, 0, 0, 0, 0, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 55, 56, 0, 0, 0, 0, 57, 58, 59, 375, 61, 62, 63, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 645, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 648, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 719, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 765, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 845, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 27, 28, 29, 0, 30, 0, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 0, 39, 0, 40, 0, 41, 0, 0, 42, 0, 0, 0, 43, 44, 45, 46, 47, 0, 49, 50, 0, 0, 51, 0, 0, 53, 54, 0, 0, 0, 0, 0, 0, 57, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 78, 79, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 216, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 46, 47, 0, 0, 0, 0, 0, 0, 0, 269, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 217, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 218, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 46, 47, 0, 0, 0, 0, 0, 0, 269, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 393, 0, 0, 77, 394, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 364, 0, 31, 32, 33, 34, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 0, 0, 0, 0, 0, 0, 46, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 0, 0, 73, 0, 295, 296, 0, 118, 75, 76, 0, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 4, 5, 6, 7, 0, 0, 0, 8, 9, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 407, 0, 0, 0, -436, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, -436, 14, 15, 16, 17, 18, 19, 20, 21, 22, -436, 23, 24, -436, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, -477, -477, -477, -477, -477, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 46, 47, 0, 0, 0, 0, 0, 0, 269, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 528, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 606, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 607, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 0, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 4, 5, 6, 7, 0, 0, 73, 8, 9, 0, 10, 118, 75, 76, 393, 0, 0, 77, 0, 0, 0, 0, 0, 80, 0, 81, 82, 83, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 0, 0, 0, 0, 13, 0, 14, 15, 16, 17, 18, 19, 20, 21, 22, 0, 23, 24, 0, 0, 0, 25, 26, 116, 28, 0, 0, 30, 0, 0, 31, 32, 33, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 437, 438, 439, 440, 441, 442, 443, 444, 445, 46, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 0, 0, 0, 0, 0, 0, 0, 58, 59, 0, 0, 0, 0, 0, 0, 0, 64, 65, 0, 66, 67, 68, 69, 70, 71, 72, 0, 0, 0, 0, 446, 0, 73, 0, 0, 0, 0, 118, 75, 76, 0, 447, 448, 77, 449, 450, 0, 0, 884, 80, 0, 81, 82, 83, 0, 452, 0, 0, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 0, 495, 496, 497, 498, 499, 0, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 437, 438, 439, 440, 441, 442, 443, 444, 445, 0, 0, 885, 511, 76, 0, 0, 0, 0, 0, 0, 244, 886, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 446, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 447, 448, 0, 449, 450, 0, 0, 884, 0, 0, 0, 0, 0, 0, 452, 0, 0, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 0, 495, 496, 497, 498, 499, 0, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 0, 241, 242, 243, 0, 0, 0, 0, 0, 0, 0, 885, 511, 76, 0, 0, 0, 0, 0, 0, 0, 913, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 600, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 611, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 639, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 643, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 820, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 707, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 241, 242, 243, 0, 0, 0, 0, 0, 269, 0, 708, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 821, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 241, 242, 243, 0, 0, 0, 0, 270, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 330, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 374, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 550, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 551, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 554, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 557, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 562, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 583, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 696, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 437, 438, 439, 440, 441, 442, 443, 444, 445, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 697, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 446, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 447, 448, 0, 449, 450, 0, 0, 451, 0, 0, 0, 0, 0, 0, 452, 0, 711, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 0, 495, 496, 497, 498, 499, 0, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 437, 438, 439, 440, 441, 442, 443, 444, 445, 0, 0, 510, 511, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 446, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 447, 448, 0, 449, 450, 0, 0, 931, 0, 0, 0, 0, 0, 0, 452, 0, 0, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 785, 786, 787, 788, 789, 790, 491, 492, 493, 494, 0, 495, 496, 497, 498, 499, 0, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 241, 242, 243, 0, 0, 0, 0, 0, 0, 0, 0, 510, 511, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 241, 242, 243, 0, 0, 386, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 560, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 620, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 241, 242, 243, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 242, 243, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 298, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 0, 0, 0, 0, 244, 0, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 0, 268, 0, 0, 0, 0, 0, 0, 0, 0, 0, 269}
var LangYycheck []yytype_int16 = []yytype_int16{2, 23, 176, 180, 632, 2, 2, 13, 373, 151, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 2, 24, 23, 2, 2, 102, 274, 676, 276, 351, 558, 85, 37, 36, 30, 796, 13, 13, 30, 13, 43, 44, 52, 120, 73, 741, 49, 73, 51, 74, 201, 202, 32, 33, 29, 807, 110, 809, 544, 73, 660, 73, 99, 100, 74, 73, 76, 34, 705, 96, 74, 54, 36, 395, 77, 74, 44, 80, 36, 23, 163, 49, 165, 131, 116, 117, 29, 783, 152, 73, 104, 0, 106, 55, 126, 159, 161, 162, 73, 164, 73, 74, 356, 65, 51, 52, 53, 74, 55, 53, 74, 30, 99, 100, 810, 73, 74, 165, 65, 77, 73, 159, 722, 73, 724, 841, 163, 843, 155, 615, 73, 32, 33, 162, 117, 784, 162, 162, 73, 83, 161, 155, 838, 164, 169, 160, 149, 150, 151, 335, 162, 150, 151, 156, 162, 158, 917, 115, 133, 134, 163, 164, 165, 96, 167, 169, 162, 32, 33, 161, 180, 155, 182, 159, 132, 133, 163, 642, 153, 73, 155, 435, 131, 160, 160, 713, 160, 159, 156, 162, 133, 134, 195, 196, 946, 153, 169, 155, 201, 202, 210, 159, 155, 206, 215, 155, 96, 165, 924, 96, 153, 169, 155, 162, 194, 218, 165, 160, 153, 402, 155, 160, 155, 939, 940, 164, 229, 207, 602, 232, 695, 947, 237, 384, 159, 951, 216, 268, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 155, 269, 268, 155, 918, 73, 74, 275, 159, 74, 278, 279, 96, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 74, 194, 76, 298, 74, 73, 74, 74, 937, 76, 305, 306, 104, 105, 207, 275, 159, 563, 313, 279, 115, 160, 73, 216, 159, 164, 329, 166, 155, 159, 325, 126, 268, 335, 161, 131, 73, 194, 274, 341, 276, 277, 161, 589, 155, 164, 306, 115, 73, 344, 207, 161, 666, 104, 164, 106, 152, 160, 153, 216, 155, 164, 162, 159, 300, 131, 162, 303, 304, 165, 73, 160, 69, 70, 169, 164, 150, 151, 159, 149, 150, 151, 149, 150, 151, 153, 152, 155, 85, 384, 568, 386, 155, 388, 168, 573, 162, 159, 393, 165, 74, 169, 76, 104, 399, 106, 401, 343, 133, 134, 159, 434, 407, 159, 74, 729, 76, 8, 9, 10, 734, 159, 161, 73, 74, 164, 159, 77, 153, 160, 155, 677, 36, 164, 74, 166, 76, 74, 29, 76, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 709, 55, 115, 780, 69, 70, 49, 645, 517, 152, 648, 65, 521, 150, 151, 161, 525, 165, 164, 161, 133, 152, 164, 664, 665, 159, 149, 150, 151, 670, 73, 74, 161, 159, 161, 164, 79, 164, 747, 155, 153, 65, 155, 752, 104, 105, 159, 150, 151, 655, 150, 151, 165, 152, 512, 112, 169, 131, 115, 116, 117, 118, 119, 120, 121, 523, 112, 109, 110, 115, 116, 117, 118, 119, 120, 121, 161, 162, 152, 81, 82, 719, 98, 99, 100, 691, 159, 88, 162, 161, 731, 165, 733, 528, 554, 161, 15, 557, 556, 585, 558, 559, 588, 568, 15, 131, 163, 160, 573, 715, 161, 749, 166, 98, 99, 100, 161, 163, 73, 73, 161, 73, 162, 581, 159, 560, 152, 765, 166, 73, 164, 167, 159, 159, 159, 621, 162, 603, 776, 165, 781, 599, 603, 603, 579, 30, 632, 161, 159, 73, 159, 609, 618, 862, 204, 613, 166, 860, 208, 164, 160, 73, 620, 161, 161, 868, 603, 74, 164, 603, 603, 606, 528, 164, 224, 161, 226, 227, 228, 129, 129, 160, 160, 73, 822, 162, 889, 163, 159, 15, 676, 164, 73, 902, 159, 159, 159, 655, 658, 664, 665, 159, 130, 160, 560, 670, 166, 845, 528, 912, 919, 15, 921, 166, 852, 166, 702, 160, 653, 705, 686, 162, 79, 579, 682, 686, 686, 166, 160, 160, 933, 160, 159, 691, 160, 36, 696, 159, 161, 164, 560, 159, 162, 164, 710, 161, 712, 161, 161, 686, 606, 161, 686, 686, 740, 713, 164, 715, 718, 579, 30, 161, 30, 156, 160, 751, 731, 905, 733, 161, 159, 909, 112, 911, 162, 115, 116, 117, 118, 119, 120, 121, 717, 166, 163, 162, 606, 112, 164, 775, 115, 116, 117, 118, 119, 120, 121, 653, 784, 163, 160, 160, 768, 835, 112, 74, 162, 161, 794, 160, 160, 160, 950, 771, 161, 163, 781, 48, 49, 50, 51, 52, 53, 163, 55, 163, 162, 13, 15, 795, 159, 161, 163, 653, 65, 74, 160, 803, 163, 112, 161, 161, 115, 116, 117, 118, 119, 120, 121, 815, 15, 104, 812, 106, 814, 15, 162, 74, 96, 162, 826, 717, 115, 116, 117, 118, 119, 120, 854, 162, 152, 161, 161, 96, 860, 161, 113, 160, 159, 52, 917, 664, 840, 733, 182, 731, 166, 853, 579, 90, 163, 717, 702, 341, 814, 819, 751, 717, 794, 621, 613, 343, 898, 889, 889, 268, 346, 901, 313, 908, 327, 276, 299, 879, 388, 881, 581, 303, -1, -1, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, 918, -1, -1, -1, -1, 895, -1, -1, -1, 899, -1, -1, -1, -1, 904, -1, -1, -1, -1, 937, -1, -1, -1, -1, -1, -1, -1, 925, -1, 927, -1, 929, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, 952, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, 106, 107, 108, -1, -1, 111, 112, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, 48, 49, -1, -1, -1, -1, 54, 65, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, 106, 107, 108, -1, -1, 111, 112, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, 106, 107, 108, -1, -1, 111, 112, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, 85, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, 89, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, 95, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, 93, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, 91, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, 163, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, 10, 161, 162, 163, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, 8, 9, 10, 11, 12, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 66, 67, -1, 69, 70, -1, -1, 73, 74, -1, -1, -1, -1, -1, 80, -1, -1, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, -1, 126, 127, 128, 129, 130, -1, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, 153, 154, -1, -1, -1, -1, -1, -1, -1, 162, -1, -1, -1, -1, -1, -1, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, 116, 117, -1, -1, -1, -1, 122, 123, 124, 125, 126, 127, 128, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, 75, -1, 77, -1, -1, 80, 81, 82, 83, 84, -1, 86, 87, 88, -1, 90, -1, 92, -1, 94, -1, -1, 97, -1, -1, -1, 101, 102, 103, 104, 105, -1, 107, 108, -1, -1, 111, -1, -1, 114, 115, -1, -1, -1, -1, -1, -1, 122, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, 161, 162, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, 36, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 104, 105, -1, -1, -1, -1, -1, -1, -1, 65, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, 156, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, 104, 105, -1, -1, -1, -1, -1, -1, 65, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, 156, -1, -1, 159, 160, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, 78, -1, 80, 81, 82, 83, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, -1, -1, -1, -1, -1, -1, 104, 105, -1, -1, -1, -1, -1, -1, -1, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, -1, -1, 148, -1, 81, 82, -1, 153, 154, 155, -1, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, 4, 5, 6, 7, -1, -1, -1, 11, 12, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 30, -1, -1, -1, 131, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, 152, 56, 57, 58, 59, 60, 61, 62, 63, 64, 162, 66, 67, 165, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, 104, 105, -1, -1, -1, -1, -1, -1, 65, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, 36, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 104, 105, -1, -1, -1, -1, -1, -1, -1, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, 36, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 104, 105, -1, -1, -1, -1, -1, -1, -1, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, -1, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 104, 105, -1, -1, -1, -1, -1, -1, -1, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, -1, -1, 148, 11, 12, -1, 14, 153, 154, 155, 156, -1, -1, 159, -1, -1, -1, -1, -1, 165, -1, 167, 168, 169, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 48, 49, -1, -1, -1, -1, 54, -1, 56, 57, 58, 59, 60, 61, 62, 63, 64, -1, 66, 67, -1, -1, -1, 71, 72, 73, 74, -1, -1, 77, -1, -1, 80, 81, 82, 83, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 104, 105, -1, -1, -1, -1, -1, -1, -1, -1, -1, 115, -1, -1, -1, -1, -1, -1, -1, 123, 124, -1, -1, -1, -1, -1, -1, -1, 132, 133, -1, 135, 136, 137, 138, 139, 140, 141, -1, -1, -1, -1, 55, -1, 148, -1, -1, -1, -1, 153, 154, 155, -1, 66, 67, 159, 69, 70, -1, -1, 73, 165, -1, 167, 168, 169, -1, 80, -1, -1, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, -1, 126, 127, 128, 129, 130, -1, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, 8, 9, 10, 11, 12, -1, -1, 153, 154, 155, -1, -1, -1, -1, -1, -1, 29, 163, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, 55, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, 66, 67, -1, 69, 70, -1, -1, 73, -1, -1, -1, -1, -1, -1, 80, -1, -1, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, -1, 126, 127, 128, 129, 130, -1, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, -1, 8, 9, 10, -1, -1, -1, -1, -1, -1, -1, 153, 154, 155, -1, -1, -1, -1, -1, -1, -1, 163, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, 8, 9, 10, -1, -1, -1, -1, -1, 65, -1, 163, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 161, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, 8, 9, 10, -1, -1, -1, -1, 161, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 161, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, 160, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, 160, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 160, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 66, 67, -1, 69, 70, -1, -1, 73, -1, -1, -1, -1, -1, -1, 80, -1, 160, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, -1, 126, 127, 128, 129, 130, -1, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 4, 5, 6, 7, 8, 9, 10, 11, 12, -1, -1, 153, 154, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 66, 67, -1, 69, 70, -1, -1, 73, -1, -1, -1, -1, -1, -1, 80, -1, -1, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, -1, 126, 127, 128, 129, 130, -1, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 8, 9, 10, -1, -1, -1, -1, -1, -1, -1, -1, 153, 154, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, 8, 9, 10, -1, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, 96, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 8, 9, 10, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, 9, 10, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 13, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, -1, -1, -1, -1, 29, -1, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, -1, 55, -1, -1, -1, -1, -1, -1, -1, -1, -1, 65}

/* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
   symbol of state STATE-NUM.  */

var LangYystos []yytype_uint16 = []yytype_uint16{0, 171, 175, 0, 4, 5, 6, 7, 11, 12, 14, 48, 49, 54, 56, 57, 58, 59, 60, 61, 62, 63, 64, 66, 67, 71, 72, 73, 74, 75, 77, 80, 81, 82, 83, 84, 86, 87, 88, 90, 92, 94, 97, 101, 102, 103, 104, 105, 106, 107, 108, 111, 112, 114, 115, 116, 117, 122, 123, 124, 125, 126, 127, 128, 132, 133, 135, 136, 137, 138, 139, 140, 141, 148, 153, 154, 155, 159, 161, 162, 165, 167, 168, 169, 176, 177, 178, 194, 201, 204, 207, 208, 209, 211, 224, 225, 226, 227, 268, 269, 270, 271, 272, 280, 281, 286, 287, 288, 290, 291, 292, 293, 294, 295, 296, 307, 73, 115, 153, 269, 272, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 115, 126, 177, 266, 281, 282, 295, 297, 269, 30, 159, 159, 165, 286, 288, 294, 294, 159, 283, 159, 262, 263, 269, 194, 159, 159, 159, 159, 159, 269, 289, 289, 73, 73, 191, 261, 289, 162, 269, 104, 106, 155, 176, 181, 183, 187, 189, 190, 238, 239, 295, 74, 240, 241, 270, 159, 159, 159, 159, 206, 210, 212, 159, 159, 74, 76, 149, 150, 151, 304, 305, 155, 176, 180, 176, 269, 192, 36, 132, 156, 269, 300, 301, 302, 303, 76, 284, 304, 76, 304, 162, 295, 155, 159, 235, 126, 208, 69, 70, 69, 70, 85, 8, 9, 10, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 55, 65, 161, 36, 276, 276, 152, 165, 152, 131, 162, 165, 235, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 81, 82, 276, 13, 267, 152, 235, 285, 131, 152, 162, 165, 269, 269, 300, 289, 269, 161, 164, 88, 269, 264, 265, 269, 269, 191, 269, 161, 161, 161, 15, 161, 164, 161, 192, 161, 176, 189, 96, 155, 155, 176, 182, 187, 161, 161, 164, 161, 164, 15, 161, 164, 199, 200, 294, 269, 308, 309, 269, 160, 73, 73, 73, 300, 300, 131, 165, 149, 305, 78, 269, 294, 76, 149, 305, 176, 161, 179, 162, 160, 125, 163, 193, 194, 201, 204, 209, 211, 294, 159, 269, 13, 166, 164, 167, 168, 269, 73, 156, 160, 236, 237, 269, 205, 159, 194, 159, 30, 161, 269, 269, 269, 30, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 282, 269, 159, 73, 273, 4, 5, 6, 7, 8, 9, 10, 11, 12, 55, 66, 67, 69, 70, 73, 80, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 126, 127, 128, 129, 130, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 153, 154, 162, 172, 173, 174, 295, 298, 289, 174, 295, 298, 73, 162, 295, 299, 269, 289, 36, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 285, 295, 299, 295, 269, 289, 160, 160, 166, 160, 160, 263, 159, 160, 161, 164, 96, 160, 160, 269, 261, 163, 155, 73, 162, 176, 155, 161, 161, 155, 176, 190, 239, 269, 241, 164, 184, 164, 184, 160, 161, 129, 213, 273, 129, 214, 160, 160, 73, 49, 73, 74, 79, 306, 163, 165, 163, 163, 162, 175, 159, 300, 36, 132, 269, 15, 301, 163, 269, 164, 184, 73, 269, 269, 192, 269, 30, 29, 133, 134, 177, 228, 229, 230, 231, 232, 233, 273, 159, 269, 235, 166, 235, 269, 235, 163, 166, 294, 213, 163, 166, 30, 194, 269, 30, 194, 223, 264, 269, 36, 132, 165, 216, 294, 195, 30, 162, 220, 273, 196, 162, 181, 185, 188, 189, 155, 162, 200, 160, 309, 160, 177, 130, 215, 162, 177, 244, 273, 15, 79, 166, 269, 175, 163, 160, 160, 294, 159, 269, 237, 160, 213, 160, 160, 85, 269, 233, 160, 164, 36, 202, 159, 228, 163, 163, 215, 192, 160, 192, 161, 294, 159, 300, 13, 160, 30, 194, 219, 161, 221, 161, 221, 109, 110, 198, 185, 189, 164, 184, 162, 186, 189, 161, 244, 273, 242, 164, 162, 269, 166, 163, 161, 300, 215, 194, 30, 161, 30, 234, 230, 156, 203, 228, 160, 273, 161, 89, 264, 300, 166, 216, 30, 194, 218, 192, 221, 98, 99, 100, 221, 163, 159, 162, 184, 188, 163, 186, 164, 184, 162, 112, 115, 116, 117, 118, 119, 120, 121, 163, 243, 253, 254, 255, 256, 177, 242, 163, 160, 273, 192, 232, 273, 74, 160, 112, 277, 162, 161, 160, 160, 160, 192, 95, 98, 161, 269, 30, 161, 222, 163, 177, 197, 192, 163, 184, 189, 163, 242, 244, 231, 106, 272, 256, 163, 162, 13, 15, 234, 159, 234, 242, 30, 194, 217, 218, 93, 161, 161, 222, 192, 34, 74, 163, 163, 163, 161, 162, 245, 74, 257, 258, 174, 259, 260, 276, 242, 274, 269, 274, 36, 74, 278, 279, 274, 163, 192, 161, 192, 177, 160, 73, 153, 163, 174, 177, 246, 247, 248, 249, 250, 251, 15, 273, 161, 164, 15, 161, 164, 174, 163, 275, 162, 74, 160, 164, 162, 91, 162, 152, 163, 247, 161, 161, 96, 113, 269, 258, 269, 260, 273, 269, 192, 279, 192, 161, 192, 174, 73, 172, 256, 244, 273, 273, 159, 274, 163, 163, 163, 174, 228, 274, 274, 160, 234, 274, 161, 162, 252, 192, 274, 163}

/* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */

var LangYyr1 []yytype_uint16 = []yytype_uint16{0, 170, 171, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 173, 173, 173, 173, 173, 173, 173, 174, 174, 175, 175, 176, 176, 177, 177, 177, 178, 178, 178, 178, 178, 178, 178, 179, 178, 180, 178, 178, 178, 178, 178, 178, 181, 181, 182, 182, 183, 183, 184, 184, 185, 185, 186, 186, 187, 187, 188, 188, 189, 189, 190, 190, 191, 191, 192, 192, 193, 193, 193, 193, 193, 193, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 195, 194, 194, 194, 194, 194, 194, 196, 196, 197, 197, 198, 198, 199, 199, 200, 201, 202, 202, 203, 203, 205, 204, 206, 204, 207, 207, 208, 208, 210, 209, 212, 211, 213, 213, 214, 214, 215, 215, 216, 216, 216, 216, 217, 217, 218, 218, 219, 219, 220, 220, 220, 220, 221, 221, 221, 222, 222, 223, 223, 224, 224, 225, 225, 226, 226, 227, 227, 228, 228, 229, 229, 230, 230, 231, 231, 232, 232, 233, 233, 233, 234, 234, 235, 235, 236, 236, 237, 237, 238, 238, 239, 240, 240, 241, 241, 242, 242, 243, 243, 243, 243, 244, 244, 245, 245, 245, 246, 246, 247, 247, 248, 249, 249, 249, 249, 250, 250, 251, 252, 252, 253, 253, 254, 254, 255, 255, 256, 256, 256, 256, 256, 256, 257, 257, 258, 258, 259, 259, 260, 261, 262, 262, 263, 264, 264, 265, 265, 267, 266, 268, 268, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 270, 270, 271, 272, 273, 274, 275, 276, 276, 277, 277, 278, 278, 279, 279, 280, 280, 280, 280, 281, 281, 282, 282, 283, 283, 284, 284, 284, 285, 285, 286, 286, 286, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 288, 288, 288, 289, 289, 290, 291, 291, 291, 292, 292, 292, 293, 293, 293, 293, 293, 293, 294, 294, 294, 295, 295, 295, 296, 296, 297, 297, 297, 297, 297, 297, 298, 298, 298, 299, 299, 299, 300, 301, 301, 302, 302, 303, 303, 303, 303, 303, 303, 303, 304, 304, 304, 304, 305, 305, 305, 305, 305, 305, 305, 306, 306, 306, 306, 307, 307, 307, 307, 307, 307, 307, 308, 308, 309}

/* YYR2[YYN] -- Number of symbols on the right hand side of rule YYN.  */

var LangYyr2 []yytype_uint8 = []yytype_uint8{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 0, 1, 3, 1, 3, 2, 1, 1, 1, 1, 1, 4, 3, 0, 6, 0, 5, 3, 4, 3, 4, 3, 1, 1, 6, 7, 6, 7, 0, 1, 3, 1, 3, 1, 3, 1, 1, 2, 1, 3, 1, 2, 3, 1, 2, 0, 1, 1, 1, 1, 1, 4, 3, 1, 1, 5, 7, 9, 5, 3, 3, 3, 3, 3, 3, 1, 2, 6, 7, 9, 0, 6, 1, 6, 3, 3, 2, 0, 9, 1, 3, 0, 4, 1, 3, 1, 13, 0, 1, 0, 1, 0, 10, 0, 9, 1, 2, 1, 1, 0, 7, 0, 8, 0, 2, 0, 2, 0, 2, 1, 2, 4, 3, 1, 4, 1, 4, 1, 4, 3, 4, 4, 5, 0, 5, 4, 1, 1, 1, 4, 5, 6, 1, 3, 6, 7, 3, 6, 1, 0, 1, 3, 4, 6, 0, 1, 1, 2, 1, 1, 1, 0, 2, 2, 4, 1, 3, 1, 2, 3, 1, 1, 3, 1, 1, 3, 2, 0, 4, 4, 3, 12, 1, 3, 1, 2, 3, 1, 2, 2, 2, 3, 3, 3, 4, 3, 1, 1, 3, 1, 3, 1, 1, 0, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 1, 2, 4, 3, 1, 4, 4, 3, 1, 1, 0, 1, 3, 1, 0, 9, 3, 2, 1, 6, 5, 3, 4, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 5, 4, 3, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 3, 2, 1, 2, 4, 2, 1, 2, 13, 12, 1, 1, 0, 0, 0, 0, 1, 0, 4, 3, 1, 1, 2, 2, 4, 4, 2, 1, 1, 1, 1, 0, 3, 0, 1, 1, 0, 1, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2, 3, 3, 1, 1, 1, 3, 3, 0, 1, 1, 1, 3, 1, 1, 3, 1, 1, 4, 4, 4, 4, 1, 1, 1, 3, 1, 4, 2, 3, 3, 1, 4, 4, 3, 3, 3, 1, 3, 1, 1, 3, 1, 1, 0, 1, 3, 1, 3, 1, 4, 2, 2, 6, 4, 2, 2, 1, 2, 1, 4, 3, 3, 3, 6, 3, 1, 1, 2, 1, 5, 4, 2, 2, 4, 2, 2, 1, 3, 1}

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

func LangYystrlen(yystr *byte) int {
	var yylen int
	for yylen = 0; yystr[yylen]; yylen++ {
		continue
	}
	return yylen
}

/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */

func LangYystpcpy(yydest *byte, yysrc *byte) *byte {
	var yyd *byte = yydest
	var yys *byte = yysrc
	for g.Assign(&(g.PostInc(&(*yyd))), g.PostInc(&(*yys))) != '0' {
		continue
	}
	return yyd - 1
}

/* Copy into *YYMSG, which is of size *YYMSG_ALLOC, an error message
   about the unexpected token YYTOKEN for the state stack whose top is
   YYSSP.

   Return 0 if *YYMSG was successfully written.  Return 1 if *YYMSG is
   not large enough to hold the message.  In that case, also set
   *YYMSG_ALLOC to the required number of bytes.  Return 2 if the
   required number of bytes is too large to store.  */

func LangYysyntaxError(yymsg_alloc *int, yymsg **byte, yyssp *yytype_int16, yytoken int) int {
	var yysize0 int = ZendYytnamerr(0, yytname[yytoken])
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
		if yyn != -753 {

			/* Start YYX at -YYN if negative to avoid negative indexes in
			   YYCHECK.  In other words, skip the first -YYN actions for
			   this state because they are default actions.  */

			var yyxbegin int = g.Cond(yyn < 0, -yyn, 0)

			/* Stay within bounds of both yycheck and yytname.  */

			var yychecklim int = 7174 - yyn + 1
			var yyxend int = g.Cond(yychecklim < 170, yychecklim, 170)
			var yyx int
			for yyx = yyxbegin; yyx < yyxend; yyx++ {
				if yycheck[yyx+yyn] == yyx && yyx != 1 && yytable[yyx+yyn] != -477 {
					if yycount == YYERROR_VERBOSE_ARGS_MAXIMUM {
						yycount = 1
						yysize = yysize0
						break
					}
					yyarg[g.PostInc(&yycount)] = yytname[yyx]
					var yysize1 int = yysize + ZendYytnamerr(0, yytname[yyx])
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
			yyp += ZendYytnamerr(yyp, yyarg[g.PostInc(&yyi)])
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

func LangYydestruct(yymsg *byte, yytype int, yyvaluep *ZendParserStackElem) {
	void(yyvaluep)
	if yymsg == nil {
		yymsg = "Deleting"
	}
	switch yytype {
	case 71:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 72:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 73:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 74:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 75:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 76:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 77:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 78:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 79:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 174:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 175:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 176:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 177:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 178:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 182:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 183:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 185:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 186:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 187:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 188:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 189:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 190:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 191:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 192:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 193:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 194:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 196:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 197:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 198:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 199:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 200:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 201:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 204:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 209:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 211:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 213:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 214:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 215:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 216:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 217:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 218:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 219:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 220:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 221:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 223:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 224:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 225:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 226:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 227:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 228:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 229:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 230:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 231:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 232:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 233:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 234:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 235:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 236:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 237:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 238:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 239:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 240:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 241:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 242:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 243:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 244:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 245:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 246:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 247:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 248:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 249:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 250:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 251:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 252:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 257:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 258:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 259:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 260:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 261:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 262:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 263:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 264:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 265:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 266:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 268:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 269:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 270:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 273:
		if yyvaluep.GetStr() != nil {
			ZendStringReleaseEx(yyvaluep.GetStr(), 0)
		}
		break
	case 277:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 278:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 279:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 280:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 281:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 282:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 283:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 284:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 285:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 286:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 287:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 288:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 289:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 290:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 291:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 292:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 293:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 294:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 295:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 296:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 297:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 298:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 299:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 300:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 301:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 302:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 303:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 304:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 305:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 306:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 307:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 308:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 309:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	default:
		break
	}
}

/*----------.
| yyparse.  |
`----------*/

func Zendparse() int {
	/* The lookahead symbol.  */

	var yychar int

	/* The semantic value of the lookahead symbol.  */

	var yyval_default ZendParserStackElem
	var yylval ZendParserStackElem = yyval_default

	/* Number of syntax errors so far.  */

	var zendnerrs int
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

	var yyvsa []ZendParserStackElem
	var yyvs *ZendParserStackElem
	var yyvsp *ZendParserStackElem
	var yystacksize int
	var yyn int
	var yyresult int

	/* Lookahead token as an internal (translated) token number.  */

	var yytoken int = 0

	/* The variables used to return semantic value and location from the
	   action routines.  */

	var yyval ZendParserStackElem

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
	zendnerrs = 0
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
	if yystate == 3 {
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
	if yyn == -753 {
		goto yydefault
	}

	/* Not known => get a lookahead token if don't already have one.  */

	if yychar == -2 {
		yychar = Zendlex(&yylval)
	}
	if yychar <= 0 {
		yytoken = 0
		yychar = yytoken
	} else {
		yytoken = uint(g.CondF1(yychar <= 396, func() __auto__ { return yytranslate[yychar] }, 2))
	}

	/* If the proper action on seeing token YYTOKEN is to reduce or to
	   detect an error, take that action.  */

	yyn += yytoken
	if yyn < 0 || 7174 < yyn || yycheck[yyn] != yytoken {
		goto yydefault
	}
	yyn = yytable[yyn]
	if yyn <= 0 {
		if yyn == -477 {
			goto yyerrlab
		}
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
	case 2:
		CG.SetAst(yyvsp[0].GetAst())
		break
	case 78:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 79:
		var zv Zval
		ZendLexTstring(&zv)
		yyval.SetAst(ZendAstCreateZval(&zv))
		break
	case 80:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 81:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_STMT_LIST))
		break
	case 82:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 83:
		yyval.SetAst(ZendAstAppendStr(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 84:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(1)
		break
	case 85:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(2)
		break
	case 86:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(0)
		break
	case 87:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 88:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 89:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 90:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 91:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 92:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_HALT_COMPILER, ZendAstCreateZvalFromLong(ZendGetScannedFileOffset())))
		ZendStopLexing()
		break
	case 93:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_NAMESPACE, yyvsp[-1].GetAst(), nil))
		if CG.GetDocComment() != nil {
			ZendStringReleaseEx(CG.GetDocComment(), 0)
			CG.SetDocComment(nil)
		}
		break
	case 94:
		if CG.GetDocComment() != nil {
			ZendStringReleaseEx(CG.GetDocComment(), 0)
			CG.SetDocComment(nil)
		}
		break
	case 95:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_NAMESPACE, yyvsp[-4].GetAst(), yyvsp[-1].GetAst()))
		break
	case 96:
		if CG.GetDocComment() != nil {
			ZendStringReleaseEx(CG.GetDocComment(), 0)
			CG.SetDocComment(nil)
		}
		break
	case 97:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_NAMESPACE, nil, yyvsp[-1].GetAst()))
		break
	case 98:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 99:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(yyvsp[-2].GetNum())
		break
	case 100:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(1 << 0)
		break
	case 101:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(yyvsp[-2].GetNum())
		break
	case 102:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 103:
		yyval.SetNum(1 << 1)
		break
	case 104:
		yyval.SetNum(1 << 2)
		break
	case 105:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GROUP_USE, yyvsp[-5].GetAst(), yyvsp[-2].GetAst()))
		break
	case 106:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GROUP_USE, yyvsp[-5].GetAst(), yyvsp[-2].GetAst()))
		break
	case 107:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GROUP_USE, yyvsp[-5].GetAst(), yyvsp[-2].GetAst()))
		break
	case 108:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GROUP_USE, yyvsp[-5].GetAst(), yyvsp[-2].GetAst()))
		break
	case 111:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 112:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_USE, yyvsp[0].GetAst()))
		break
	case 113:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 114:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_USE, yyvsp[0].GetAst()))
		break
	case 115:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 116:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_USE, yyvsp[0].GetAst()))
		break
	case 117:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(1 << 0)
		break
	case 118:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(yyvsp[-1].GetNum())
		break
	case 119:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_USE_ELEM, yyvsp[0].GetAst(), nil))
		break
	case 120:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_USE_ELEM, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 121:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 122:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 123:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 124:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_CONST_DECL, yyvsp[0].GetAst()))
		break
	case 125:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 126:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_STMT_LIST))
		break
	case 127:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 128:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 129:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 130:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 131:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 132:
		yyval.SetAst(nil)
		ZendThrowException(ZendCeCompileError, "__HALT_COMPILER() can only be used from the outermost scope", 0)
		goto yyerrorlab
		break
	case 133:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 134:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 135:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 136:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_WHILE, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 137:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DO_WHILE, yyvsp[-5].GetAst(), yyvsp[-2].GetAst()))
		break
	case 138:
		yyval.SetAst(ZendAstCreate4(ZEND_AST_FOR, yyvsp[-6].GetAst(), yyvsp[-4].GetAst(), yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 139:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_SWITCH, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 140:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_BREAK, yyvsp[-1].GetAst()))
		break
	case 141:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_CONTINUE, yyvsp[-1].GetAst()))
		break
	case 142:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_RETURN, yyvsp[-1].GetAst()))
		break
	case 143:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 144:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 145:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 146:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_ECHO, yyvsp[0].GetAst()))
		break
	case 147:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 148:
		yyval.SetAst(yyvsp[-3].GetAst())
		break
	case 149:
		yyval.SetAst(ZendAstCreate4(ZEND_AST_FOREACH, yyvsp[-4].GetAst(), yyvsp[-2].GetAst(), nil, yyvsp[0].GetAst()))
		break
	case 150:
		yyval.SetAst(ZendAstCreate4(ZEND_AST_FOREACH, yyvsp[-6].GetAst(), yyvsp[-2].GetAst(), yyvsp[-4].GetAst(), yyvsp[0].GetAst()))
		break
	case 151:
		if ZendHandleEncodingDeclaration(yyvsp[-1].GetAst()) == 0 {
			goto yyerrorlab
		}
		break
	case 152:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DECLARE, yyvsp[-3].GetAst(), yyvsp[0].GetAst()))
		break
	case 153:
		yyval.SetAst(nil)
		break
	case 154:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_TRY, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 155:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_THROW, yyvsp[-1].GetAst()))
		break
	case 156:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_GOTO, yyvsp[-1].GetAst()))
		break
	case 157:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_LABEL, yyvsp[-1].GetAst()))
		break
	case 158:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_CATCH_LIST))
		break
	case 159:
		yyval.SetAst(ZendAstListAdd(yyvsp[-8].GetAst(), ZendAstCreate3(ZEND_AST_CATCH, yyvsp[-5].GetAst(), yyvsp[-4].GetAst(), yyvsp[-1].GetAst())))
		break
	case 160:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_NAME_LIST, yyvsp[0].GetAst()))
		break
	case 161:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 162:
		yyval.SetAst(nil)
		break
	case 163:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 164:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_STMT_LIST, yyvsp[0].GetAst()))
		break
	case 165:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 166:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_UNSET, yyvsp[0].GetAst()))
		break
	case 167:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_FUNC_DECL, yyvsp[-11].GetNum()|yyvsp[0].GetNum(), yyvsp[-12].GetNum(), yyvsp[-9].GetStr(), ZendAstGetStr(yyvsp[-10].GetAst()), yyvsp[-7].GetAst(), nil, yyvsp[-2].GetAst(), yyvsp[-5].GetAst()))
		CG.SetExtraFnFlags(yyvsp[-4].GetNum())
		break
	case 168:
		yyval.SetNum(0)
		break
	case 169:
		yyval.SetNum(1 << 0)
		break
	case 170:
		yyval.SetNum(0)
		break
	case 171:
		yyval.SetNum(1 << 1)
		break
	case 172:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 173:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_CLASS, yyvsp[-9].GetNum(), yyvsp[-7].GetNum(), yyvsp[-3].GetStr(), ZendAstGetStr(yyvsp[-6].GetAst()), yyvsp[-5].GetAst(), yyvsp[-4].GetAst(), yyvsp[-1].GetAst(), nil))
		break
	case 174:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 175:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_CLASS, 0, yyvsp[-7].GetNum(), yyvsp[-3].GetStr(), ZendAstGetStr(yyvsp[-6].GetAst()), yyvsp[-5].GetAst(), yyvsp[-4].GetAst(), yyvsp[-1].GetAst(), nil))
		break
	case 176:
		yyval.SetNum(yyvsp[0].GetNum())
		break
	case 177:
		yyval.SetNum(ZendAddClassModifier(yyvsp[-1].GetNum(), yyvsp[0].GetNum()))
		if yyval.GetNum() == 0 {
			goto yyerrorlab
		}
		break
	case 178:
		yyval.SetNum(1 << 6)
		break
	case 179:
		yyval.SetNum(1 << 5)
		break
	case 180:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 181:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_CLASS, 1<<1, yyvsp[-5].GetNum(), yyvsp[-3].GetStr(), ZendAstGetStr(yyvsp[-4].GetAst()), nil, nil, yyvsp[-1].GetAst(), nil))
		break
	case 182:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 183:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_CLASS, 1<<0, yyvsp[-6].GetNum(), yyvsp[-3].GetStr(), ZendAstGetStr(yyvsp[-5].GetAst()), nil, yyvsp[-4].GetAst(), yyvsp[-1].GetAst(), nil))
		break
	case 184:
		yyval.SetAst(nil)
		break
	case 185:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 186:
		yyval.SetAst(nil)
		break
	case 187:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 188:
		yyval.SetAst(nil)
		break
	case 189:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 190:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 191:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_REF, yyvsp[0].GetAst()))
		break
	case 192:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(1)
		break
	case 193:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(3)
		break
	case 194:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 195:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 196:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 197:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 198:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 199:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 200:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 201:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 202:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 203:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 204:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_SWITCH_LIST))
		break
	case 205:
		yyval.SetAst(ZendAstListAdd(yyvsp[-4].GetAst(), ZendAstCreate2(ZEND_AST_SWITCH_CASE, yyvsp[-2].GetAst(), yyvsp[0].GetAst())))
		break
	case 206:
		yyval.SetAst(ZendAstListAdd(yyvsp[-3].GetAst(), ZendAstCreate2(ZEND_AST_SWITCH_CASE, nil, yyvsp[0].GetAst())))
		break
	case 209:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 210:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 211:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_IF, ZendAstCreate2(ZEND_AST_IF_ELEM, yyvsp[-2].GetAst(), yyvsp[0].GetAst())))
		break
	case 212:
		yyval.SetAst(ZendAstListAdd(yyvsp[-5].GetAst(), ZendAstCreate2(ZEND_AST_IF_ELEM, yyvsp[-2].GetAst(), yyvsp[0].GetAst())))
		break
	case 213:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 214:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), ZendAstCreate2(ZEND_AST_IF_ELEM, nil, yyvsp[0].GetAst())))
		break
	case 215:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_IF, ZendAstCreate2(ZEND_AST_IF_ELEM, yyvsp[-3].GetAst(), yyvsp[0].GetAst())))
		break
	case 216:
		yyval.SetAst(ZendAstListAdd(yyvsp[-6].GetAst(), ZendAstCreate2(ZEND_AST_IF_ELEM, yyvsp[-3].GetAst(), yyvsp[0].GetAst())))
		break
	case 217:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 218:
		yyval.SetAst(ZendAstListAdd(yyvsp[-5].GetAst(), ZendAstCreate2(ZEND_AST_IF_ELEM, nil, yyvsp[-2].GetAst())))
		break
	case 219:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 220:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_PARAM_LIST))
		break
	case 221:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_PARAM_LIST, yyvsp[0].GetAst()))
		break
	case 222:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 223:
		yyval.SetAst(ZendAstCreateEx3(ZEND_AST_PARAM, yyvsp[-2].GetNum()|yyvsp[-1].GetNum(), yyvsp[-3].GetAst(), yyvsp[0].GetAst(), nil))
		break
	case 224:
		yyval.SetAst(ZendAstCreateEx3(ZEND_AST_PARAM, yyvsp[-4].GetNum()|yyvsp[-3].GetNum(), yyvsp[-5].GetAst(), yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 225:
		yyval.SetAst(nil)
		break
	case 226:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 227:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 228:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(yyval.GetAst().GetAttr() | 1<<8)
		break
	case 229:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_TYPE, 7))
		break
	case 230:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_TYPE, 17))
		break
	case 231:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 232:
		yyval.SetAst(nil)
		break
	case 233:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 234:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_ARG_LIST))
		break
	case 235:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 236:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_ARG_LIST, yyvsp[0].GetAst()))
		break
	case 237:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 238:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 239:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_UNPACK, yyvsp[0].GetAst()))
		break
	case 240:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 241:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_STMT_LIST, yyvsp[0].GetAst()))
		break
	case 242:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_GLOBAL, ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst())))
		break
	case 243:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 244:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_STMT_LIST, yyvsp[0].GetAst()))
		break
	case 245:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC, yyvsp[0].GetAst(), nil))
		break
	case 246:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 247:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 248:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_STMT_LIST))
		break
	case 249:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_PROP_GROUP, yyvsp[-2].GetAst(), yyvsp[-1].GetAst()))
		yyval.GetAst().SetAttr(yyvsp[-3].GetNum())
		break
	case 250:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(yyvsp[-3].GetNum())
		break
	case 251:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_USE_TRAIT, yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 252:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_METHOD, yyvsp[-9].GetNum()|yyvsp[-11].GetNum()|yyvsp[0].GetNum(), yyvsp[-10].GetNum(), yyvsp[-7].GetStr(), ZendAstGetStr(yyvsp[-8].GetAst()), yyvsp[-5].GetAst(), nil, yyvsp[-1].GetAst(), yyvsp[-3].GetAst()))
		CG.SetExtraFnFlags(yyvsp[-2].GetNum())
		break
	case 253:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_NAME_LIST, yyvsp[0].GetAst()))
		break
	case 254:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 255:
		yyval.SetAst(nil)
		break
	case 256:
		yyval.SetAst(nil)
		break
	case 257:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 258:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_TRAIT_ADAPTATIONS, yyvsp[0].GetAst()))
		break
	case 259:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 260:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 261:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 262:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_TRAIT_PRECEDENCE, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 263:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_TRAIT_ALIAS, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 264:
		var zv Zval
		ZendLexTstring(&zv)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_TRAIT_ALIAS, yyvsp[-2].GetAst(), ZendAstCreateZval(&zv)))
		break
	case 265:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_TRAIT_ALIAS, yyvsp[-1].GetNum(), yyvsp[-3].GetAst(), yyvsp[0].GetAst()))
		break
	case 266:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_TRAIT_ALIAS, yyvsp[0].GetNum(), yyvsp[-2].GetAst(), nil))
		break
	case 267:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_METHOD_REFERENCE, nil, yyvsp[0].GetAst()))
		break
	case 268:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 269:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_METHOD_REFERENCE, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 270:
		yyval.SetAst(nil)
		break
	case 271:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 272:
		yyval.SetNum(yyvsp[0].GetNum())
		break
	case 273:
		yyval.SetNum(1 << 0)
		break
	case 274:
		yyval.SetNum(1 << 0)
		break
	case 275:
		yyval.SetNum(yyvsp[0].GetNum())
		if (yyval.GetNum() & (1<<0 | 1<<1 | 1<<2)) == 0 {
			yyval.SetNum(yyval.GetNum() | 1<<0)
		}
		break
	case 276:
		yyval.SetNum(yyvsp[0].GetNum())
		break
	case 277:
		yyval.SetNum(ZendAddMemberModifier(yyvsp[-1].GetNum(), yyvsp[0].GetNum()))
		if yyval.GetNum() == 0 {
			goto yyerrorlab
		}
		break
	case 278:
		yyval.SetNum(1 << 0)
		break
	case 279:
		yyval.SetNum(1 << 1)
		break
	case 280:
		yyval.SetNum(1 << 2)
		break
	case 281:
		yyval.SetNum(1 << 4)
		break
	case 282:
		yyval.SetNum(1 << 6)
		break
	case 283:
		yyval.SetNum(1 << 5)
		break
	case 284:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 285:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_PROP_DECL, yyvsp[0].GetAst()))
		break
	case 286:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_PROP_ELEM, yyvsp[-1].GetAst(), nil, g.CondF1(yyvsp[0].GetStr() != nil, func() *ZendAst { return ZendAstCreateZvalFromStr(yyvsp[0].GetStr()) }, nil)))
		break
	case 287:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_PROP_ELEM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), g.CondF1(yyvsp[0].GetStr() != nil, func() *ZendAst { return ZendAstCreateZvalFromStr(yyvsp[0].GetStr()) }, nil)))
		break
	case 288:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 289:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_CLASS_CONST_DECL, yyvsp[0].GetAst()))
		break
	case 290:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_CONST_ELEM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), g.CondF1(yyvsp[0].GetStr() != nil, func() *ZendAst { return ZendAstCreateZvalFromStr(yyvsp[0].GetStr()) }, nil)))
		break
	case 291:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_CONST_ELEM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), g.CondF1(yyvsp[0].GetStr() != nil, func() *ZendAst { return ZendAstCreateZvalFromStr(yyvsp[0].GetStr()) }, nil)))
		break
	case 292:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 293:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_STMT_LIST, yyvsp[0].GetAst()))
		break
	case 294:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_ECHO, yyvsp[0].GetAst()))
		break
	case 295:
		yyval.SetAst(nil)
		break
	case 296:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 297:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 298:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_EXPR_LIST, yyvsp[0].GetAst()))
		break
	case 299:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 300:
		var decl *ZendAst = ZendAstCreateDecl(ZEND_AST_CLASS, 1<<2, yyvsp[-7].GetNum(), yyvsp[-3].GetStr(), nil, yyvsp[-5].GetAst(), yyvsp[-4].GetAst(), yyvsp[-1].GetAst(), nil)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_NEW, decl, yyvsp[-6].GetAst()))
		break
	case 301:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_NEW, yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 302:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 303:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 304:
		yyvsp[-3].GetAst().SetAttr(1)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ASSIGN, yyvsp[-3].GetAst(), yyvsp[0].GetAst()))
		break
	case 305:
		yyvsp[-3].GetAst().SetAttr(3)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ASSIGN, yyvsp[-3].GetAst(), yyvsp[0].GetAst()))
		break
	case 306:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ASSIGN, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 307:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ASSIGN_REF, yyvsp[-3].GetAst(), yyvsp[0].GetAst()))
		break
	case 308:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_CLONE, yyvsp[0].GetAst()))
		break
	case 309:
		yyval.SetAst(ZendAstCreateAssignOp(1, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 310:
		yyval.SetAst(ZendAstCreateAssignOp(2, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 311:
		yyval.SetAst(ZendAstCreateAssignOp(3, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 312:
		yyval.SetAst(ZendAstCreateAssignOp(12, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 313:
		yyval.SetAst(ZendAstCreateAssignOp(4, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 314:
		yyval.SetAst(ZendAstCreateAssignOp(8, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 315:
		yyval.SetAst(ZendAstCreateAssignOp(5, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 316:
		yyval.SetAst(ZendAstCreateAssignOp(10, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 317:
		yyval.SetAst(ZendAstCreateAssignOp(9, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 318:
		yyval.SetAst(ZendAstCreateAssignOp(11, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 319:
		yyval.SetAst(ZendAstCreateAssignOp(6, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 320:
		yyval.SetAst(ZendAstCreateAssignOp(7, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 321:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ASSIGN_COALESCE, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 322:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_POST_INC, yyvsp[-1].GetAst()))
		break
	case 323:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_PRE_INC, yyvsp[0].GetAst()))
		break
	case 324:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_POST_DEC, yyvsp[-1].GetAst()))
		break
	case 325:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_PRE_DEC, yyvsp[0].GetAst()))
		break
	case 326:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_OR, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 327:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_AND, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 328:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_OR, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 329:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_AND, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 330:
		yyval.SetAst(ZendAstCreateBinaryOp(15, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 331:
		yyval.SetAst(ZendAstCreateBinaryOp(9, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 332:
		yyval.SetAst(ZendAstCreateBinaryOp(10, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 333:
		yyval.SetAst(ZendAstCreateBinaryOp(11, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 334:
		yyval.SetAst(ZendAstCreateBinaryOp(8, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 335:
		yyval.SetAst(ZendAstCreateBinaryOp(1, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 336:
		yyval.SetAst(ZendAstCreateBinaryOp(2, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 337:
		yyval.SetAst(ZendAstCreateBinaryOp(3, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 338:
		yyval.SetAst(ZendAstCreateBinaryOp(12, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 339:
		yyval.SetAst(ZendAstCreateBinaryOp(4, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 340:
		yyval.SetAst(ZendAstCreateBinaryOp(5, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 341:
		yyval.SetAst(ZendAstCreateBinaryOp(6, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 342:
		yyval.SetAst(ZendAstCreateBinaryOp(7, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 343:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_UNARY_PLUS, yyvsp[0].GetAst()))
		break
	case 344:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_UNARY_MINUS, yyvsp[0].GetAst()))
		break
	case 345:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_UNARY_OP, 14, yyvsp[0].GetAst()))
		break
	case 346:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_UNARY_OP, 13, yyvsp[0].GetAst()))
		break
	case 347:
		yyval.SetAst(ZendAstCreateBinaryOp(16, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 348:
		yyval.SetAst(ZendAstCreateBinaryOp(17, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 349:
		yyval.SetAst(ZendAstCreateBinaryOp(18, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 350:
		yyval.SetAst(ZendAstCreateBinaryOp(19, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 351:
		yyval.SetAst(ZendAstCreateBinaryOp(20, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 352:
		yyval.SetAst(ZendAstCreateBinaryOp(21, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 353:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GREATER, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 354:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_GREATER_EQUAL, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 355:
		yyval.SetAst(ZendAstCreateBinaryOp(170, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 356:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_INSTANCEOF, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 357:
		yyval.SetAst(yyvsp[-1].GetAst())
		if yyval.GetAst().GetKind() == ZEND_AST_CONDITIONAL {
			yyval.GetAst().SetAttr(1)
		}
		if yyval.GetAst().GetKind() == ZEND_AST_BINARY_OP && yyval.GetAst().GetAttr() == 8 {
			yyval.GetAst().SetAttr(252)
		}
		break
	case 358:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 359:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_CONDITIONAL, yyvsp[-4].GetAst(), yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 360:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_CONDITIONAL, yyvsp[-3].GetAst(), nil, yyvsp[0].GetAst()))
		break
	case 361:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_COALESCE, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 362:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 363:
		yyval.SetAst(ZendAstCreateCast(4, yyvsp[0].GetAst()))
		break
	case 364:
		yyval.SetAst(ZendAstCreateCast(5, yyvsp[0].GetAst()))
		break
	case 365:
		yyval.SetAst(ZendAstCreateCast(6, yyvsp[0].GetAst()))
		break
	case 366:
		yyval.SetAst(ZendAstCreateCast(7, yyvsp[0].GetAst()))
		break
	case 367:
		yyval.SetAst(ZendAstCreateCast(8, yyvsp[0].GetAst()))
		break
	case 368:
		yyval.SetAst(ZendAstCreateCast(16, yyvsp[0].GetAst()))
		break
	case 369:
		yyval.SetAst(ZendAstCreateCast(1, yyvsp[0].GetAst()))
		break
	case 370:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_EXIT, yyvsp[0].GetAst()))
		break
	case 371:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_SILENCE, yyvsp[0].GetAst()))
		break
	case 372:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 373:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_SHELL_EXEC, yyvsp[-1].GetAst()))
		break
	case 374:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_PRINT, yyvsp[0].GetAst()))
		break
	case 375:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_YIELD, nil, nil))
		CG.SetExtraFnFlags(CG.GetExtraFnFlags() | 1<<24)
		break
	case 376:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_YIELD, yyvsp[0].GetAst(), nil))
		CG.SetExtraFnFlags(CG.GetExtraFnFlags() | 1<<24)
		break
	case 377:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_YIELD, yyvsp[0].GetAst(), yyvsp[-2].GetAst()))
		CG.SetExtraFnFlags(CG.GetExtraFnFlags() | 1<<24)
		break
	case 378:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_YIELD_FROM, yyvsp[0].GetAst()))
		CG.SetExtraFnFlags(CG.GetExtraFnFlags() | 1<<24)
		break
	case 379:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 380:
		yyval.SetAst(yyvsp[0].GetAst())
		(*ZendAstDecl)(yyval.GetAst()).SetFlags((*ZendAstDecl)(yyval.GetAst()).GetFlags() | 1<<4)
		break
	case 381:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_CLOSURE, yyvsp[-11].GetNum()|yyvsp[0].GetNum(), yyvsp[-12].GetNum(), yyvsp[-10].GetStr(), ZendStringInit("{closure}", g.SizeOf("\"{closure}\"")-1, 0), yyvsp[-8].GetAst(), yyvsp[-6].GetAst(), yyvsp[-2].GetAst(), yyvsp[-5].GetAst()))
		CG.SetExtraFnFlags(yyvsp[-4].GetNum())
		break
	case 382:
		yyval.SetAst(ZendAstCreateDecl(ZEND_AST_ARROW_FUNC, yyvsp[-10].GetNum()|yyvsp[0].GetNum(), yyvsp[-11].GetNum(), yyvsp[-5].GetStr(), ZendStringInit("{closure}", g.SizeOf("\"{closure}\"")-1, 0), yyvsp[-8].GetAst(), nil, ZendAstCreate1(ZEND_AST_RETURN, yyvsp[-1].GetAst()), yyvsp[-6].GetAst()))
		(*ZendAstDecl)(yyval.GetAst()).SetLexPos(yyvsp[-2].GetPtr())
		CG.SetExtraFnFlags(yyvsp[-3].GetNum())
		break
	case 383:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 384:
		yyval.SetNum(CG.GetZendLineno())
		break
	case 385:
		yyval.SetStr(CG.GetDocComment())
		CG.SetDocComment(nil)
		break
	case 386:
		yyval.SetNum(CG.GetExtraFnFlags())
		CG.SetExtraFnFlags(0)
		break
	case 387:
		yyval.SetPtr(LANG_SCNG.GetYyText())
		break
	case 388:
		yyval.SetNum(0)
		break
	case 389:
		yyval.SetNum(1 << 12)
		break
	case 390:
		yyval.SetAst(nil)
		break
	case 391:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 392:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 393:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_CLOSURE_USES, yyvsp[0].GetAst()))
		break
	case 394:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 395:
		yyval.SetAst(yyvsp[0].GetAst())
		yyval.GetAst().SetAttr(1)
		break
	case 396:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_CALL, yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 397:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_STATIC_CALL, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 398:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_STATIC_CALL, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 399:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_CALL, yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 400:
		var zv Zval
		var __z *Zval = &zv
		var __s *ZendString = ZendKnownStrings[ZEND_STR_STATIC]
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		yyval.SetAst(ZendAstCreateZvalEx(&zv, 1))
		break
	case 401:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 402:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 403:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 404:
		yyval.SetAst(nil)
		break
	case 405:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 406:
		yyval.SetAst(ZendAstCreateZvalFromStr(ZendEmptyString))
		break
	case 407:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 408:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 409:
		yyval.SetAst(ZendAstCreateList0(ZEND_AST_ARG_LIST))
		break
	case 410:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 411:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(2)
		break
	case 412:
		yyval.SetAst(yyvsp[-1].GetAst())
		yyval.GetAst().SetAttr(3)
		break
	case 413:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 414:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 415:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 416:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_LINE))
		break
	case 417:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_FILE))
		break
	case 418:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_DIR))
		break
	case 419:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_TRAIT_C))
		break
	case 420:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_METHOD_C))
		break
	case 421:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_FUNC_C))
		break
	case 422:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_NS_C))
		break
	case 423:
		yyval.SetAst(ZendAstCreateEx0(ZEND_AST_MAGIC_CONST, T_CLASS_C))
		break
	case 424:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 425:
		yyval.SetAst(ZendAstCreateZvalFromStr(ZendEmptyString))
		break
	case 426:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 427:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 428:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 429:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 430:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_CONST, yyvsp[0].GetAst()))
		break
	case 431:
		yyval.SetAst(ZendAstCreateClassConstOrName(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 432:
		yyval.SetAst(ZendAstCreateClassConstOrName(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 433:
		yyval.SetAst(nil)
		break
	case 434:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 435:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 436:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 437:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 438:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 439:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 440:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 441:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 442:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 443:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DIM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst()))
		break
	case 444:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DIM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst()))
		break
	case 445:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_DIM, 1<<1, yyvsp[-3].GetAst(), yyvsp[-1].GetAst()))
		break
	case 446:
		yyval.SetAst(ZendAstCreate3(ZEND_AST_METHOD_CALL, yyvsp[-3].GetAst(), yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 447:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 448:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 449:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 450:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 451:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 452:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 453:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 454:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 455:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 456:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 457:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DIM, yyvsp[-3].GetAst(), yyvsp[-1].GetAst()))
		break
	case 458:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_DIM, 1<<1, yyvsp[-3].GetAst(), yyvsp[-1].GetAst()))
		break
	case 459:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 460:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 461:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_STATIC_PROP, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 462:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 463:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 464:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 465:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 466:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 467:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 468:
		yyval.SetAst(ZendAstListRtrim(yyvsp[0].GetAst()))
		break
	case 469:
		yyval.SetAst(nil)
		break
	case 470:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 471:
		yyval.SetAst(ZendAstListAdd(yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 472:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_ARRAY, yyvsp[0].GetAst()))
		break
	case 473:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ARRAY_ELEM, yyvsp[0].GetAst(), yyvsp[-2].GetAst()))
		break
	case 474:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ARRAY_ELEM, yyvsp[0].GetAst(), nil))
		break
	case 475:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_ARRAY_ELEM, 1, yyvsp[0].GetAst(), yyvsp[-3].GetAst()))
		break
	case 476:
		yyval.SetAst(ZendAstCreateEx2(ZEND_AST_ARRAY_ELEM, 1, yyvsp[0].GetAst(), nil))
		break
	case 477:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_UNPACK, yyvsp[0].GetAst()))
		break
	case 478:
		yyvsp[-1].GetAst().SetAttr(1)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ARRAY_ELEM, yyvsp[-1].GetAst(), yyvsp[-5].GetAst()))
		break
	case 479:
		yyvsp[-1].GetAst().SetAttr(1)
		yyval.SetAst(ZendAstCreate2(ZEND_AST_ARRAY_ELEM, yyvsp[-1].GetAst(), nil))
		break
	case 480:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 481:
		yyval.SetAst(ZendAstListAdd(yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 482:
		yyval.SetAst(ZendAstCreateList1(ZEND_AST_ENCAPS_LIST, yyvsp[0].GetAst()))
		break
	case 483:
		yyval.SetAst(ZendAstCreateList2(ZEND_AST_ENCAPS_LIST, yyvsp[-1].GetAst(), yyvsp[0].GetAst()))
		break
	case 484:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 485:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DIM, ZendAstCreate1(ZEND_AST_VAR, yyvsp[-3].GetAst()), yyvsp[-1].GetAst()))
		break
	case 486:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_PROP, ZendAstCreate1(ZEND_AST_VAR, yyvsp[-2].GetAst()), yyvsp[0].GetAst()))
		break
	case 487:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[-1].GetAst()))
		break
	case 488:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[-1].GetAst()))
		break
	case 489:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_DIM, ZendAstCreate1(ZEND_AST_VAR, yyvsp[-4].GetAst()), yyvsp[-2].GetAst()))
		break
	case 490:
		yyval.SetAst(yyvsp[-1].GetAst())
		break
	case 491:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 492:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 493:
		yyval.SetAst(ZendNegateNumString(yyvsp[0].GetAst()))
		break
	case 494:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_VAR, yyvsp[0].GetAst()))
		break
	case 495:
		yyval.SetAst(yyvsp[-2].GetAst())
		break
	case 496:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_EMPTY, yyvsp[-1].GetAst()))
		break
	case 497:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_INCLUDE_OR_EVAL, 1<<1, yyvsp[0].GetAst()))
		break
	case 498:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_INCLUDE_OR_EVAL, 1<<2, yyvsp[0].GetAst()))
		break
	case 499:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_INCLUDE_OR_EVAL, 1<<0, yyvsp[-1].GetAst()))
		break
	case 500:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_INCLUDE_OR_EVAL, 1<<3, yyvsp[0].GetAst()))
		break
	case 501:
		yyval.SetAst(ZendAstCreateEx1(ZEND_AST_INCLUDE_OR_EVAL, 1<<4, yyvsp[0].GetAst()))
		break
	case 502:
		yyval.SetAst(yyvsp[0].GetAst())
		break
	case 503:
		yyval.SetAst(ZendAstCreate2(ZEND_AST_AND, yyvsp[-2].GetAst(), yyvsp[0].GetAst()))
		break
	case 504:
		yyval.SetAst(ZendAstCreate1(ZEND_AST_ISSET, yyvsp[0].GetAst()))
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
	yystate = yypgoto[yyn-170] + (*yyssp)
	if 0 <= yystate && yystate <= 7174 && yycheck[yystate] == (*yyssp) {
		yystate = yytable[yystate]
	} else {
		yystate = yydefgoto[yyn-170]
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
		yytoken = uint(g.CondF1(yychar <= 396, func() __auto__ { return yytranslate[yychar] }, 2))
	}

	/* If not already recovering from an error, report this error.  */

	if yyerrstatus == 0 {
		zendnerrs++

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
		Zenderror(yymsgp)
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
		if yyn != -753 {
			yyn += 1
			if 0 <= yyn && yyn <= 7174 && yycheck[yyn] == 1 {
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
	Zenderror("memory exhausted")
	yyresult = 2

	/* Fall through.  */

yyreturn:
	if yychar != -2 {

		/* Make sure we have latest lookahead translation.  See comments at
		   user semantic actions for why this is necessary.  */

		yytoken = uint(g.CondF1(yychar <= 396, func() __auto__ { return yytranslate[yychar] }, 2))
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

/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for yyerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from yytname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */

func ZendYytnamerr(yyres *byte, yystr *byte) int {
	/* CG(parse_error) states:
	 * 0 => yyres = NULL, yystr is the unexpected token
	 * 1 => yyres = NULL, yystr is one of the expected tokens
	 * 2 => yyres != NULL, yystr is the unexpected token
	 * 3 => yyres != NULL, yystr is one of the expected tokens
	 */

	if yyres != nil && CG.GetParseError() < 2 {
		CG.SetParseError(2)
	}
	if CG.GetParseError()%2 == 0 {

		/* The unexpected token */

		var buffer []byte
		var end *uint8
		var str *uint8
		var tok1 *uint8 = nil
		var tok2 *uint8 = nil
		var len_ uint = 0
		var toklen uint = 0
		var yystr_len uint
		CG.GetParseError()++
		if LANG_SCNG.GetYyText()[0] == 0 && LANG_SCNG.GetYyLeng() == 1 && strcmp(yystr, "\"end of file\"") == 0 {
			if yyres != nil {
				yystpcpy(yyres, "end of file")
			}
			return g.SizeOf("\"end of file\"") - 1
		}
		str = LANG_SCNG.GetYyText()
		end = memchr(str, '\n', LANG_SCNG.GetYyLeng())
		yystr_len = uint(yystrlen(yystr))
		if g.Assign(&tok1, memchr(yystr, '(', yystr_len)) != nil && g.Assign(&tok2, ZendMemrchr(yystr, ')', yystr_len)) != nil {
			toklen = tok2 - tok1 + 1
		} else {
			tok2 = nil
			tok1 = tok2
			toklen = 0
		}
		if end == nil {
			if LANG_SCNG.GetYyLeng() > 30 {
				len_ = 30
			} else {
				len_ = LANG_SCNG.GetYyLeng()
			}
		} else {
			if end-str > 30 {
				len_ = 30
			} else {
				len_ = end - str
			}
		}
		if yyres != nil {
			if toklen != 0 {
				snprintf(buffer, g.SizeOf("buffer"), "'%.*s' %.*s", len_, str, toklen, tok1)
			} else {
				snprintf(buffer, g.SizeOf("buffer"), "'%.*s'", len_, str)
			}
			yystpcpy(yyres, buffer)
		}
		return len_ + g.Cond(toklen != 0, toklen+1, 0) + 2
	}

	/* One of the expected tokens */

	if yyres == nil {
		return yystrlen(yystr) - g.Cond((*yystr) == '"', 2, 0)
	}
	if (*yystr) == '"' {
		var yyn int = 0
		var yyp *byte = yystr
		for ; (*(g.PreInc(&yyp))) != '"'; yyn++ {
			yyres[yyn] = *yyp
		}
		yyres[yyn] = '0'
		return yyn
	}
	yystpcpy(yyres, yystr)
	return strlen(yystr)
}
