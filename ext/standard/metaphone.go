// <<generate>>

package standard

import (
	b "sik/builtin"
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
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &phones, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if Metaphone((*uint8)(zend.ZSTR_VAL(str)), zend.ZSTR_LEN(str), phones, &result, 1) == 0 {
		zend.RETVAL_STR(result)
	} else {
		if result != nil {
			zend.ZendStringFree(result)
		}
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

const SH = 'X'
const TH = '0'

/*-----------------------------  */

var _codes []byte = []byte{1, 16, 4, 16, 9, 2, 4, 16, 9, 2, 0, 2, 2, 2, 1, 4, 0, 2, 4, 4, 1, 0, 0, 0, 8, 0}

func ENCODE(c __auto__) __auto__ {
	if isalpha(c) {
		return _codes[toupper(c)-'A']
	} else {
		return 0
	}
}
func Isvowel(c byte) int { return ENCODE(c) & 1 }

/* These letters are passed through unchanged */

func NOCHANGE(c __auto__) int { return ENCODE(c) & 2 }

/* These form diphthongs when preceding H */

func AFFECTH(c __auto__) int { return ENCODE(c) & 4 }

/* These make C and G soft */

func MAKESOFT(c byte) int { return ENCODE(c) & 8 }

/* These prevent GH from becoming F */

func NOGHTOF(c char) int { return ENCODE(c) & 16 }

/*----------------------------- */

const Next_Letter byte = toupper(word[w_idx+1])

/* Look at the current letter in the word */

const Curr_Letter byte = toupper(word[w_idx])

/* Go N letters back. */

func Look_Back_Letter(n int) char {
	if w_idx >= n {
		return toupper(word[w_idx-n])
	} else {
		return '0'
	}
}

/* Previous letter.  I dunno, should this return null on failure? */

const Prev_Letter = Look_Back_Letter(1)

/* Look two letters down.  It makes sure you don't walk off the string. */

const After_Next_Letter byte = b.CondF1(Next_Letter != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0')

func Look_Ahead_Letter(n int) __auto__ {
	return toupper(Lookahead((*byte)(word+w_idx), n))
}

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

func Phonize(c byte) {
	if p_idx >= max_buffer_len {
		*phoned_word = zend.ZendStringExtend(*phoned_word, 2*b.SizeOf("char")+max_buffer_len, 0)
		max_buffer_len += 2
	}
	zend.ZSTR_VAL(*phoned_word)[b.PostInc(&p_idx)] = c
	zend.ZSTR_LEN(*phoned_word) = p_idx
}

/* Slap a null character on the end of the phoned word */

// #define End_Phoned_Word       { if ( p_idx == max_buffer_len ) { * phoned_word = zend_string_extend ( * phoned_word , 1 * sizeof ( char ) + max_buffer_len , 0 ) ; max_buffer_len += 1 ; } ZSTR_VAL ( * phoned_word ) [ p_idx ] = '\0' ; ZSTR_LEN ( * phoned_word ) = p_idx ; }

/* How long is the phoned word? */

const Phone_Len = p_idx

/* Note is a letter is a 'break' in the word */

func Isbreak(c byte) bool { return !(isalpha(c)) }

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
		*phoned_word = zend.ZendStringAlloc(b.SizeOf("char")*word_len+1, 0)
	} else {
		max_buffer_len = max_phonemes
		*phoned_word = zend.ZendStringAlloc(b.SizeOf("char")*max_phonemes+1, 0)
	}

	/*-- The first phoneme has to be processed specially. --*/

	for ; !(isalpha(Curr_Letter)); w_idx++ {

		/* On the off chance we were given nothing but crap... */

		if Curr_Letter == '0' {
			if p_idx == max_buffer_len {
				*phoned_word = zend.ZendStringExtend(*phoned_word, 1*b.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 1
			}
			zend.ZSTR_VAL(*phoned_word)[p_idx] = '0'
			zend.ZSTR_LEN(*phoned_word) = p_idx
			return zend.SUCCESS
		}

		/* On the off chance we were given nothing but crap... */

	}
	switch Curr_Letter {
	case 'A':
		if Next_Letter == 'E' {
			Phonize('E')
			w_idx += 2
		} else {
			Phonize('A')
			w_idx++
		}
		break
	case 'G':

	case 'K':

	case 'P':
		if Next_Letter == 'N' {
			Phonize('N')
			w_idx += 2
		}
		break
	case 'W':
		if Next_Letter == 'R' {
			Phonize(Next_Letter)
			w_idx += 2
		} else if Next_Letter == 'H' || Isvowel(Next_Letter) != 0 {
			Phonize('W')
			w_idx += 2
		}

		/* else ignore */

		break
	case 'X':
		Phonize('S')
		w_idx++
		break
	case 'E':

	case 'I':

	case 'O':

	case 'U':
		Phonize(Curr_Letter)
		w_idx++
		break
	default:

		/* do nothing */

		break
	}

	/* On to the metaphoning */

	for ; Curr_Letter != '0' && (max_phonemes == 0 || Phone_Len < int(max_phonemes)); w_idx++ {

		/* How many letters to skip because an eariler encoding handled
		 * multiple letters */

		var skip_letter unsigned__short__int = 0

		/* THOUGHT:  It would be nice if, rather than having things like...
		 * well, SCI.  For SCI you encode the S, then have to remember
		 * to skip the C.  So the phonome SCI invades both S and C.  It would
		 * be better, IMHO, to skip the C from the S part of the encoding.
		 * Hell, I'm trying it.
		 */

		if !(isalpha(Curr_Letter)) {
			continue
		}

		/* Drop duplicates, except CC */

		if Curr_Letter == Prev_Letter && Curr_Letter != 'C' {
			continue
		}
		switch Curr_Letter {
		case 'B':
			if Prev_Letter != 'M' {
				Phonize('B')
			}
			break
		case 'C':
			if MAKESOFT(Next_Letter) != 0 {
				if After_Next_Letter == 'A' && Next_Letter == 'I' {
					Phonize(SH)
				} else if Prev_Letter == 'S' {

				} else {
					Phonize('S')
				}
			} else if Next_Letter == 'H' {
				if traditional == 0 && (After_Next_Letter == 'R' || Prev_Letter == 'S') {
					Phonize('K')
				} else {
					Phonize(SH)
				}
				skip_letter++
			} else {
				Phonize('K')
			}
			break
		case 'D':
			if Next_Letter == 'G' && MAKESOFT(After_Next_Letter) != 0 {
				Phonize('J')
				skip_letter++
			} else {
				Phonize('T')
			}
			break
		case 'G':
			if Next_Letter == 'H' {
				if !(NOGHTOF(Look_Back_Letter(3)) != 0 || Look_Back_Letter(4) == 'H') {
					Phonize('F')
					skip_letter++
				}
			} else if Next_Letter == 'N' {
				if Isbreak(After_Next_Letter) || After_Next_Letter == 'E' && Look_Ahead_Letter(3) == 'D' {

				} else {
					Phonize('K')
				}
			} else if MAKESOFT(Next_Letter) != 0 && Prev_Letter != 'G' {
				Phonize('J')
			} else {
				Phonize('K')
			}
			break
		case 'H':
			if Isvowel(Next_Letter) != 0 && AFFECTH(Prev_Letter) == 0 {
				Phonize('H')
			}
			break
		case 'K':
			if Prev_Letter != 'C' {
				Phonize('K')
			}
			break
		case 'P':
			if Next_Letter == 'H' {
				Phonize('F')
			} else {
				Phonize('P')
			}
			break
		case 'Q':
			Phonize('K')
			break
		case 'S':
			if Next_Letter == 'I' && (After_Next_Letter == 'O' || After_Next_Letter == 'A') {
				Phonize(SH)
			} else if Next_Letter == 'H' {
				Phonize(SH)
				skip_letter++
			} else if traditional == 0 && (Next_Letter == 'C' && Look_Ahead_Letter(2) == 'H' && Look_Ahead_Letter(3) == 'W') {
				Phonize(SH)
				skip_letter += 2
			} else {
				Phonize('S')
			}
			break
		case 'T':
			if Next_Letter == 'I' && (After_Next_Letter == 'O' || After_Next_Letter == 'A') {
				Phonize(SH)
			} else if Next_Letter == 'H' {
				Phonize(TH)
				skip_letter++
			} else if !(Next_Letter == 'C' && After_Next_Letter == 'H') {
				Phonize('T')
			}
			break
		case 'V':
			Phonize('F')
			break
		case 'W':
			if Isvowel(Next_Letter) != 0 {
				Phonize('W')
			}
			break
		case 'X':
			Phonize('K')
			Phonize('S')
			break
		case 'Y':
			if Isvowel(Next_Letter) != 0 {
				Phonize('Y')
			}
			break
		case 'Z':
			Phonize('S')
			break
		case 'F':

		case 'J':

		case 'L':

		case 'M':

		case 'N':

		case 'R':
			Phonize(Curr_Letter)
			break
		default:

			/* nothing */

			break
		}
		w_idx += skip_letter
	}
	if p_idx == max_buffer_len {
		*phoned_word = zend.ZendStringExtend(*phoned_word, 1*b.SizeOf("char")+max_buffer_len, 0)
		max_buffer_len += 1
	}
	zend.ZSTR_VAL(*phoned_word)[p_idx] = '0'
	zend.ZSTR_LEN(*phoned_word) = p_idx
	return 0
}

/* }}} */
