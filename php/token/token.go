package token

type Token int

//go:generate stringer -type=Token
const (
	_          Token = iota
	Comment          // [T_COMMENT]
	DocComment       // [T_DOC_COMMENT]

	// literal
	Ident  // main 		[T_STRING]
	Int    // 12345 	[T_LNUMBER]
	Float  // 123.45 	[T_DNUMBER]
	String // "abc" 	[T_CONSTANT_ENCAPSED_STRING]

	// unary operators
	Not   // !
	Tilde // ~
	Inc   // ++ [T_INC]
	Dec   // -- [T_DEC]

	PreInc  // 并非真实 token，用于 UnaryExpr.Kind 与 PostInc 区分
	PreDec  // 并非真实 token，用于 UnaryExpr.Kind 与 PostDec 区分
	PostInc // 并非真实 token，用于 UnaryExpr.Kind 与 PreInc 区分
	PostDec // 并非真实 token，用于 UnaryExpr.Kind 与 PreDec 区分

	binary_begin
	// binary operators
	Add      // +
	Sub      // -
	Mul      // *
	Div      // /
	Mod      // %
	Pow      // ** [T_POW]
	Concat   // .
	Coalesce // ??	[T_COALESCE]

	And        // &
	Or         // |
	Xor        // ^
	ShiftLeft  // << [T_SL]
	ShiftRight // >> [T_SR]

	BooleanAnd // &&	[T_BOOLEAN_AND]
	BooleanOr  // ||	[T_BOOLEAN_OR]
	LogicalAnd // and 	[T_LOGICAL_AND]
	LogicalOr  // or 	[T_LOGICAL_OR]
	LogicalXor // xor 	[T_LOGICAL_XOR]

	Equal          // ==  [T_IS_EQUAL]
	NotEqual       // !=  [T_IS_NOT_EQUAL]
	Identical      // === [T_IS_IDENTICAL]
	NotIdentical   // !== [T_IS_NOT_IDENTICAL]
	Greater        // >
	GreaterOrEqual // >=  [T_IS_GREATER_OR_EQUAL]
	Smaller        // <
	SmallerOrEqual // <=  [T_IS_SMALLER_OR_EQUAL]
	Spaceship      // <=> [T_SPACESHIP]

	binary_end

	assign_begin
	// assign operators
	Assign         // =
	AddAssign      // +=  [T_PLUS_EQUAL]
	SubAssign      // -=  [T_MINUS_EQUAL]
	MulAssign      // *=  [T_MUL_EQUAL]
	DivAssign      // /=  [T_DIV_EQUAL]
	ModAssign      // %=  [T_MOD_EQUAL]
	PowAssign      // **= [T_POW_EQUAL]
	ConcatAssign   // .=  [T_CONCAT_EQUAL]
	CoalesceAssign // ??= [T_COALESCE_EQUAL]

	AndAssign        // &=  [T_AND_EQUAL]
	OrAssign         // |=  [T_OR_EQUAL]
	XorAssign        // ^=  [T_XOR_EQUAL]
	ShiftLeftAssign  // <<= [T_SL_EQUAL]
	ShiftRightAssign // >>= [T_SR_EQUAL]
	assign_end

	cast_begin
	// cast op tokens
	BoolCast   // (bool)              [T_BOOL_CAST]
	IntCast    // (int) or (integer)  [T_INT_CAST]
	DoubleCast // (float) or (double) [T_DOUBLE_CAST]
	StringCast // (string)            [T_STRING_CAST]
	ArrayCast  // (array)             [T_ARRAY_CAST]
	ObjectCast // (object)            [T_OBJECT_CAST]
	UnsetCast  // (unset)             [T_UNSET_CAST]
	cast_end

	magic_const_begin
	// magic const op tokens
	DirConst       // __DIR__       [T_DIR]
	FileConst      // __FILE__      [T_FILE]
	LineConst      // __LINE__      [T_LINE]
	NamespaceConst // __NAMESPACE__ [T_NS_C]
	FunctionConst  // __FUNCTION__  [T_FUNC_C]
	ClassConst     // __CLASS__     [T_CLASS_C]
	MethodConst    // __METHOD__    [T_METHOD_C]
	TraitConst     // __TRAIT__     [T_TRAIT_C]
	magic_const_end

	internal_call_begin
	// internal call as expr
	Isset       // isset        [T_ISSET]
	Empty       // empty        [T_EMPTY]
	Include     // include      [T_INCLUDE]
	IncludeOnce // include_once [T_INCLUDE_ONCE]
	Require     // require      [T_REQUIRE]
	RequireOnce // require_once [T_REQUIRE_ONCE]
	Eval        // eval         [T_EVAL]
	internal_call_end

	// others
	NAME_FULLY_QUALIFIED                    Token = 263
	NAME_RELATIVE                           Token = 264
	NAME_QUALIFIED                          Token = 265
	VARIABLE                                Token = 266
	INLINE_HTML                             Token = 267
	ENCAPSED_AND_WHITESPACE                 Token = 268
	STRING_VARNAME                          Token = 270
	NUM_STRING                              Token = 271
	PRINT                                   Token = 280
	YIELD                                   Token = 281
	YIELD_FROM                              Token = 282
	INSTANCEOF                              Token = 283
	NEW                                     Token = 284
	CLONE                                   Token = 285
	EXIT                                    Token = 286
	IF                                      Token = 287
	ELSEIF                                  Token = 288
	ELSE                                    Token = 289
	ENDIF                                   Token = 290
	ECHO                                    Token = 291
	DO                                      Token = 292
	WHILE                                   Token = 293
	ENDWHILE                                Token = 294
	FOR                                     Token = 295
	ENDFOR                                  Token = 296
	FOREACH                                 Token = 297
	ENDFOREACH                              Token = 298
	DECLARE                                 Token = 299
	ENDDECLARE                              Token = 300
	AS                                      Token = 301
	SWITCH                                  Token = 302
	ENDSWITCH                               Token = 303
	CASE                                    Token = 304
	DEFAULT                                 Token = 305
	MATCH                                   Token = 306
	BREAK                                   Token = 307
	CONTINUE                                Token = 308
	GOTO                                    Token = 309
	FUNCTION                                Token = 310
	FN                                      Token = 311
	CONST                                   Token = 312
	RETURN                                  Token = 313
	TRY                                     Token = 314
	CATCH                                   Token = 315
	FINALLY                                 Token = 316
	THROW                                   Token = 317
	USE                                     Token = 318
	INSTEADOF                               Token = 319
	GLOBAL                                  Token = 320
	STATIC                                  Token = 321
	ABSTRACT                                Token = 322
	FINAL                                   Token = 323
	PRIVATE                                 Token = 324
	PROTECTED                               Token = 325
	PUBLIC                                  Token = 326
	READONLY                                Token = 327
	VAR                                     Token = 328
	UNSET                                   Token = 329
	HALT_COMPILER                           Token = 332
	CLASS                                   Token = 333
	TRAIT                                   Token = 334
	INTERFACE                               Token = 335
	ENUM                                    Token = 336
	EXTENDS                                 Token = 337
	IMPLEMENTS                              Token = 338
	NAMESPACE                               Token = 339
	LIST                                    Token = 340
	ARRAY                                   Token = 341
	CALLABLE                                Token = 342
	ATTRIBUTE                               Token = 351
	OBJECT_OPERATOR                         Token = 384
	NULLSAFE_OBJECT_OPERATOR                Token = 385
	DOUBLE_ARROW                            Token = 386
	COMMENT                                 Token = 387
	DOC_COMMENT                             Token = 388
	OPEN_TAG                                Token = 389
	OPEN_TAG_WITH_ECHO                      Token = 390
	CLOSE_TAG                               Token = 391
	WHITESPACE                              Token = 392
	START_HEREDOC                           Token = 393
	END_HEREDOC                             Token = 394
	DOLLAR_OPEN_CURLY_BRACES                Token = 395
	CURLY_OPEN                              Token = 396
	PAAMAYIM_NEKUDOTAYIM                    Token = 397
	DOUBLE_COLON                            Token = 397
	NS_SEPARATOR                            Token = 398
	ELLIPSIS                                Token = 399
	AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG     Token = 403
	AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG Token = 404
	BAD_CHARACTER                           Token = 405
)

func IsBinaryOp(token Token) bool {
	return binary_begin < token && token < binary_end
}

func IsAssignOp(token Token) bool {
	return assign_begin < token && token < assign_end
}

func IsCastKind(token Token) bool {
	return cast_begin < token && token < cast_end
}

func IsMagicConstKind(token Token) bool {
	return magic_const_begin < token && token < magic_const_end
}

func IsInternalCall(token Token) bool {
	return internal_call_begin < token && token < internal_call_end
}

var specialTokenNames = map[Token]string{
	// unary operators
	Not:   "!",
	Tilde: "~",
	Inc:   "++",
	Dec:   "--",

	PreInc:  "++",
	PreDec:  "--",
	PostInc: "++",
	PostDec: "--",

	// binary operators
	Add:      "+",
	Sub:      "-",
	Mul:      "*",
	Div:      "/",
	Mod:      "%",
	Pow:      "**",
	Concat:   ".",
	Coalesce: "??",

	And:        "&",
	Or:         "|",
	Xor:        "^",
	ShiftLeft:  "<<",
	ShiftRight: ">>",

	BooleanAnd: "&&",
	BooleanOr:  "||",
	LogicalAnd: "and",
	LogicalOr:  "or",
	LogicalXor: "xor",

	Equal:          "==",
	NotEqual:       "!=",
	Identical:      "===",
	NotIdentical:   "!==",
	Greater:        ">",
	GreaterOrEqual: ">=",
	Smaller:        "<",
	SmallerOrEqual: "<=",
	Spaceship:      "<=>",

	// assign operators
	Assign:         "=",
	AddAssign:      "+=",
	SubAssign:      "-=",
	MulAssign:      "*=",
	DivAssign:      "/=",
	ModAssign:      "%=",
	PowAssign:      "**=",
	ConcatAssign:   ".=",
	CoalesceAssign: "??=",

	AndAssign:        "&=",
	OrAssign:         "|=",
	XorAssign:        "^=",
	ShiftLeftAssign:  "<<=",
	ShiftRightAssign: ">>=",

	// cast op tokens
	BoolCast:   "(bool)",
	IntCast:    "(int))",
	DoubleCast: "(float)",
	StringCast: "(string)",
	ArrayCast:  "(array)",
	ObjectCast: "(object)",
	UnsetCast:  "(unset)",

	// magic const op tokens
	DirConst:       "__DIR__",
	FileConst:      "__FILE__",
	LineConst:      "__LINE__",
	NamespaceConst: "__NAMESPACE__",
	FunctionConst:  "__FUNCTION__",
	ClassConst:     "__CLASS__",
	MethodConst:    "__METHOD__",
	TraitConst:     "__TRAIT__",

	// internal call as expr
	Isset:       "isset",        // isset        [T_ISSET]
	Empty:       "empty",        // empty        [T_EMPTY]
	Include:     "include",      // include      [T_INCLUDE]
	IncludeOnce: "include_once", // include_once [T_INCLUDE_ONCE]
	Require:     "require",      // require      [T_REQUIRE]
	RequireOnce: "require_once", // require_once [T_REQUIRE_ONCE]
	Eval:        "eval",         // eval         [T_EVAL]
}

func TokenName(token Token) string {
	if name, ok := specialTokenNames[token]; ok {
		return name
	}
	return token.String()
}
