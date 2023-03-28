package token

type Token int

const (
	LNUMBER                                 Token = 260 + iota
	DNUMBER                                 Token = 261
	STRING                                  Token = 262
	NAME_FULLY_QUALIFIED                    Token = 263
	NAME_RELATIVE                           Token = 264
	NAME_QUALIFIED                          Token = 265
	VARIABLE                                Token = 266
	INLINE_HTML                             Token = 267
	ENCAPSED_AND_WHITESPACE                 Token = 268
	CONSTANT_ENCAPSED_STRING                Token = 269
	STRING_VARNAME                          Token = 270
	NUM_STRING                              Token = 271
	INCLUDE                                 Token = 272
	INCLUDE_ONCE                            Token = 273
	EVAL                                    Token = 274
	REQUIRE                                 Token = 275
	REQUIRE_ONCE                            Token = 276
	LOGICAL_OR                              Token = 277
	LOGICAL_XOR                             Token = 278
	LOGICAL_AND                             Token = 279
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
	ISSET                                   Token = 330
	EMPTY                                   Token = 331
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
	LINE                                    Token = 343
	FILE                                    Token = 344
	DIR                                     Token = 345
	CLASS_C                                 Token = 346
	TRAIT_C                                 Token = 347
	METHOD_C                                Token = 348
	FUNC_C                                  Token = 349
	NS_C                                    Token = 350
	ATTRIBUTE                               Token = 351
	PLUS_EQUAL                              Token = 352
	MINUS_EQUAL                             Token = 353
	MUL_EQUAL                               Token = 354
	DIV_EQUAL                               Token = 355
	CONCAT_EQUAL                            Token = 356
	MOD_EQUAL                               Token = 357
	AND_EQUAL                               Token = 358
	OR_EQUAL                                Token = 359
	XOR_EQUAL                               Token = 360
	SL_EQUAL                                Token = 361
	SR_EQUAL                                Token = 362
	COALESCE_EQUAL                          Token = 363
	BOOLEAN_OR                              Token = 364
	BOOLEAN_AND                             Token = 365
	IS_EQUAL                                Token = 366
	IS_NOT_EQUAL                            Token = 367
	IS_IDENTICAL                            Token = 368
	IS_NOT_IDENTICAL                        Token = 369
	IS_SMALLER_OR_EQUAL                     Token = 370
	IS_GREATER_OR_EQUAL                     Token = 371
	SPACESHIP                               Token = 372
	SL                                      Token = 373
	SR                                      Token = 374
	INC                                     Token = 375
	DEC                                     Token = 376
	INT_CAST                                Token = 377
	DOUBLE_CAST                             Token = 378
	STRING_CAST                             Token = 379
	ARRAY_CAST                              Token = 380
	OBJECT_CAST                             Token = 381
	BOOL_CAST                               Token = 382
	UNSET_CAST                              Token = 383
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
	COALESCE                                Token = 400
	POW                                     Token = 401
	POW_EQUAL                               Token = 402
	AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG     Token = 403
	AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG Token = 404
	BAD_CHARACTER                           Token = 405
)
