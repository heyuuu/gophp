// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/metaphone.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_metaphone.h"

/* {{{ proto string metaphone(string text[, int phones])
   Break english phrases down into their phonemes */

func ZifMetaphone(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var result *zend.ZendString = nil
	var phones zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &phones, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if Metaphone((*uint8)(str.val), str.len_, phones, &result, 1) == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = result
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	} else {
		if result != nil {
			zend.ZendStringFree(result)
		}
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

// #define SH       'X'

// #define TH       '0'

/*-----------------------------  */

var _codes []byte = []byte{1, 16, 4, 16, 9, 2, 4, 16, 9, 2, 0, 2, 2, 2, 1, 4, 0, 2, 4, 4, 1, 0, 0, 0, 8, 0}

// #define ENCODE(c) ( isalpha ( c ) ? _codes [ ( ( toupper ( c ) ) - 'A' ) ] : 0 )

// #define isvowel(c) ( ENCODE ( c ) & 1 )

/* These letters are passed through unchanged */

// #define NOCHANGE(c) ( ENCODE ( c ) & 2 )

/* These form diphthongs when preceding H */

// #define AFFECTH(c) ( ENCODE ( c ) & 4 )

/* These make C and G soft */

// #define MAKESOFT(c) ( ENCODE ( c ) & 8 )

/* These prevent GH from becoming F */

// #define NOGHTOF(c) ( ENCODE ( c ) & 16 )

/*----------------------------- */

// #define Next_Letter       ( toupper ( word [ w_idx + 1 ] ) )

/* Look at the current letter in the word */

// #define Curr_Letter       ( toupper ( word [ w_idx ] ) )

/* Go N letters back. */

// #define Look_Back_Letter(n) ( w_idx >= n ? toupper ( word [ w_idx - n ] ) : '\0' )

/* Previous letter.  I dunno, should this return null on failure? */

// #define Prev_Letter       ( Look_Back_Letter ( 1 ) )

/* Look two letters down.  It makes sure you don't walk off the string. */

// #define After_Next_Letter       ( Next_Letter != '\0' ? toupper ( word [ w_idx + 2 ] ) : '\0' )

// #define Look_Ahead_Letter(n) ( toupper ( Lookahead ( ( char * ) word + w_idx , n ) ) )

/* Allows us to safely look ahead an arbitrary # of letters */

func Lookahead(word *byte, how_far int) byte {
	var letter_ahead byte = '0'
	var idx int
	for idx = 0; word[idx] != '0' && idx < how_far; idx++ {

	}

	/* Edge forward in the string... */

	letter_ahead = word[idx]
	return letter_ahead
}

/* phonize one letter
 * We don't know the buffers size in advance. On way to solve this is to just
 * re-allocate the buffer size. We're using an extra of 2 characters (this
 * could be one though; or more too). */

// #define Phonize(c) { if ( p_idx >= max_buffer_len ) { * phoned_word = zend_string_extend ( * phoned_word , 2 * sizeof ( char ) + max_buffer_len , 0 ) ; max_buffer_len += 2 ; } ZSTR_VAL ( * phoned_word ) [ p_idx ++ ] = c ; ZSTR_LEN ( * phoned_word ) = p_idx ; }

/* Slap a null character on the end of the phoned word */

// #define End_Phoned_Word       { if ( p_idx == max_buffer_len ) { * phoned_word = zend_string_extend ( * phoned_word , 1 * sizeof ( char ) + max_buffer_len , 0 ) ; max_buffer_len += 1 ; } ZSTR_VAL ( * phoned_word ) [ p_idx ] = '\0' ; ZSTR_LEN ( * phoned_word ) = p_idx ; }

/* How long is the phoned word? */

// #define Phone_Len       ( p_idx )

/* Note is a letter is a 'break' in the word */

// #define Isbreak(c) ( ! isalpha ( c ) )

/* {{{ metaphone
 */

func Metaphone(word *uint8, word_len int, max_phonemes zend.ZendLong, phoned_word **zend.ZendString, traditional int) int {
	var w_idx int = 0
	var p_idx int = 0
	var max_buffer_len int = 0

	/*-- Parameter checks --*/

	if max_phonemes < 0 {
		return -1
	}

	/* Empty/null string is meaningless */

	if word == nil {
		return -1
	}

	/*-- Allocate memory for our phoned_phrase --*/

	if max_phonemes == 0 {
		max_buffer_len = word_len
		*phoned_word = zend.ZendStringAlloc(g.SizeOf("char")*word_len+1, 0)
	} else {
		max_buffer_len = max_phonemes
		*phoned_word = zend.ZendStringAlloc(g.SizeOf("char")*max_phonemes+1, 0)
	}

	/*-- The first phoneme has to be processed specially. --*/

	for ; !(isalpha(toupper(word[w_idx]))); w_idx++ {

		/* On the off chance we were given nothing but crap... */

		if toupper(word[w_idx]) == '0' {
			if p_idx == max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 1*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 1
			}
			(*phoned_word).val[p_idx] = '0'
			(*phoned_word).len_ = p_idx
			return zend.SUCCESS
		}

		/* On the off chance we were given nothing but crap... */

	}
	switch toupper(word[w_idx]) {
	case 'A':
		if toupper(word[w_idx+1]) == 'E' {
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'E'
			(*phoned_word).len_ = p_idx
			w_idx += 2
		} else {
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'A'
			(*phoned_word).len_ = p_idx
			w_idx++
		}
		break
	case 'G':

	case 'K':

	case 'P':
		if toupper(word[w_idx+1]) == 'N' {
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'N'
			(*phoned_word).len_ = p_idx
			w_idx += 2
		}
		break
	case 'W':
		if toupper(word[w_idx+1]) == 'R' {
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = toupper(word[w_idx+1])
			(*phoned_word).len_ = p_idx
			w_idx += 2
		} else if toupper(word[w_idx+1]) == 'H' || (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0)&1) != 0 {
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'W'
			(*phoned_word).len_ = p_idx
			w_idx += 2
		}

		/* else ignore */

		break
	case 'X':
		if p_idx >= max_buffer_len {
			*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
			max_buffer_len += 2
		}
		(*phoned_word).val[g.PostInc(&p_idx)] = 'S'
		(*phoned_word).len_ = p_idx
		w_idx++
		break
	case 'E':

	case 'I':

	case 'O':

	case 'U':
		if p_idx >= max_buffer_len {
			*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
			max_buffer_len += 2
		}
		(*phoned_word).val[g.PostInc(&p_idx)] = toupper(word[w_idx])
		(*phoned_word).len_ = p_idx
		w_idx++
		break
	default:

		/* do nothing */

		break
	}

	/* On to the metaphoning */

	for ; toupper(word[w_idx]) != '0' && (max_phonemes == 0 || p_idx < int(max_phonemes)); w_idx++ {

		/* How many letters to skip because an eariler encoding handled
		 * multiple letters */

		var skip_letter unsigned__short__int = 0

		/* THOUGHT:  It would be nice if, rather than having things like...
		 * well, SCI.  For SCI you encode the S, then have to remember
		 * to skip the C.  So the phonome SCI invades both S and C.  It would
		 * be better, IMHO, to skip the C from the S part of the encoding.
		 * Hell, I'm trying it.
		 */

		if !(isalpha(toupper(word[w_idx]))) {
			continue
		}

		/* Drop duplicates, except CC */

		if toupper(word[w_idx]) == g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') && toupper(word[w_idx]) != 'C' {
			continue
		}
		switch toupper(word[w_idx]) {
		case 'B':
			if g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') != 'M' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'B'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'C':
			if (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0) & 8) != 0 {
				if g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'A' && toupper(word[w_idx+1]) == 'I' {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
					(*phoned_word).len_ = p_idx
				} else if g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') == 'S' {

				} else {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'S'
					(*phoned_word).len_ = p_idx
				}
			} else if toupper(word[w_idx+1]) == 'H' {
				if traditional == 0 && (g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'R' || g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') == 'S') {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
					(*phoned_word).len_ = p_idx
				} else {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
					(*phoned_word).len_ = p_idx
				}
				skip_letter++
			} else {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'D':
			if toupper(word[w_idx+1]) == 'G' && (g.CondF1(isalpha(g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0')), func() byte {
				return _codes[toupper(g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0'))-'A']
			}, 0)&8) != 0 {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'J'
				(*phoned_word).len_ = p_idx
				skip_letter++
			} else {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'T'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'G':
			if toupper(word[w_idx+1]) == 'H' {
				if !((g.CondF1(isalpha(g.CondF1(w_idx >= 3, func() __auto__ { return toupper(word[w_idx-3]) }, '0')), func() byte {
					return _codes[toupper(g.CondF1(w_idx >= 3, func() __auto__ { return toupper(word[w_idx-3]) }, '0'))-'A']
				}, 0)&16) != 0 || g.CondF1(w_idx >= 4, func() __auto__ { return toupper(word[w_idx-4]) }, '0') == 'H') {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'F'
					(*phoned_word).len_ = p_idx
					skip_letter++
				}
			} else if toupper(word[w_idx+1]) == 'N' {
				if !(isalpha(g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0'))) || g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'E' && toupper(Lookahead((*byte)(word+w_idx), 3)) == 'D' {

				} else {
					if p_idx >= max_buffer_len {
						*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
						max_buffer_len += 2
					}
					(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
					(*phoned_word).len_ = p_idx
				}
			} else if (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0)&8) != 0 && g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') != 'G' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'J'
				(*phoned_word).len_ = p_idx
			} else {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'H':
			if (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0)&1) != 0 && (g.CondF1(isalpha(g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0')), func() byte {
				return _codes[toupper(g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0'))-'A']
			}, 0)&4) == 0 {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'H'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'K':
			if g.CondF1(w_idx >= 1, func() __auto__ { return toupper(word[w_idx-1]) }, '0') != 'C' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'P':
			if toupper(word[w_idx+1]) == 'H' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'F'
				(*phoned_word).len_ = p_idx
			} else {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'P'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'Q':
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
			(*phoned_word).len_ = p_idx
			break
		case 'S':
			if toupper(word[w_idx+1]) == 'I' && (g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'O' || g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'A') {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
				(*phoned_word).len_ = p_idx
			} else if toupper(word[w_idx+1]) == 'H' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
				(*phoned_word).len_ = p_idx
				skip_letter++
			} else if traditional == 0 && (toupper(word[w_idx+1]) == 'C' && toupper(Lookahead((*byte)(word+w_idx), 2)) == 'H' && toupper(Lookahead((*byte)(word+w_idx), 3)) == 'W') {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
				(*phoned_word).len_ = p_idx
				skip_letter += 2
			} else {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'S'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'T':
			if toupper(word[w_idx+1]) == 'I' && (g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'O' || g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'A') {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'X'
				(*phoned_word).len_ = p_idx
			} else if toupper(word[w_idx+1]) == 'H' {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = '0'
				(*phoned_word).len_ = p_idx
				skip_letter++
			} else if !(toupper(word[w_idx+1]) == 'C' && g.CondF1(toupper(word[w_idx+1]) != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0') == 'H') {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'T'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'V':
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'F'
			(*phoned_word).len_ = p_idx
			break
		case 'W':
			if (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0) & 1) != 0 {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'W'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'X':
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'K'
			(*phoned_word).len_ = p_idx
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'S'
			(*phoned_word).len_ = p_idx
			break
		case 'Y':
			if (g.CondF1(isalpha(toupper(word[w_idx+1])), func() byte { return _codes[toupper(toupper(word[w_idx+1]))-'A'] }, 0) & 1) != 0 {
				if p_idx >= max_buffer_len {
					*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
					max_buffer_len += 2
				}
				(*phoned_word).val[g.PostInc(&p_idx)] = 'Y'
				(*phoned_word).len_ = p_idx
			}
			break
		case 'Z':
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = 'S'
			(*phoned_word).len_ = p_idx
			break
		case 'F':

		case 'J':

		case 'L':

		case 'M':

		case 'N':

		case 'R':
			if p_idx >= max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 2*g.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 2
			}
			(*phoned_word).val[g.PostInc(&p_idx)] = toupper(word[w_idx])
			(*phoned_word).len_ = p_idx
			break
		default:

			/* nothing */

			break
		}
		w_idx += skip_letter
	}
	if p_idx == max_buffer_len {
		*phoned_word = zend.ZendStringExtend(*phoned_word, 1*g.SizeOf("char")+max_buffer_len, 0)
		max_buffer_len += 1
	}
	(*phoned_word).val[p_idx] = '0'
	(*phoned_word).len_ = p_idx
	return 0
}

/* }}} */
