// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func ZifMetaphone(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *types.ZendString
	var result *types.ZendString = nil
	var phones zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &phones) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			if (_flags & argparse.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == argparse.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_CLASS {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_ARG {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if Metaphone((*uint8)(str.GetVal()), str.GetLen(), phones, &result, 1) == 0 {
		return_value.SetString(result)
	} else {
		if result != nil {
			types.ZendStringFree(result)
		}
		return_value.SetFalse()
		return
	}
}
func ENCODE(c __auto__) __auto__ {
	if isalpha(c) {
		return _codes[toupper(c)-'A']
	} else {
		return 0
	}
}
func Isvowel(c byte) int      { return ENCODE(c) & 1 }
func NOCHANGE(c __auto__) int { return ENCODE(c) & 2 }
func AFFECTH(c __auto__) int  { return ENCODE(c) & 4 }
func MAKESOFT(c byte) int     { return ENCODE(c) & 8 }
func NOGHTOF(c char) int      { return ENCODE(c) & 16 }
func Look_Back_Letter(n int) char {
	if w_idx >= n {
		return toupper(word[w_idx-n])
	} else {
		return '0'
	}
}
func Look_Ahead_Letter(n int) __auto__ {
	return toupper(Lookahead((*byte)(word+w_idx), n))
}
func Lookahead(word *byte, how_far int) byte {
	var letter_ahead byte = '0'
	var idx int
	for idx = 0; word[idx] != '0' && idx < how_far; idx++ {

	}

	/* Edge forward in the string... */

	letter_ahead = word[idx]
	return letter_ahead
}
func Phonize(c byte) {
	if p_idx >= max_buffer_len {
		*phoned_word = types.ZendStringExtend(*phoned_word, 2*b.SizeOf("char")+max_buffer_len, 0)
		max_buffer_len += 2
	}
	phoned_word.GetVal()[b.PostInc(&p_idx)] = c
	phoned_word.GetLen() = p_idx
}
func Isbreak(c byte) bool { return !(isalpha(c)) }
func Metaphone(word *uint8, word_len int, max_phonemes zend.ZendLong, phoned_word **types.ZendString, traditional int) int {
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
		*phoned_word = types.ZendStringAlloc(b.SizeOf("char")*word_len+1, 0)
	} else {
		max_buffer_len = max_phonemes
		*phoned_word = types.ZendStringAlloc(b.SizeOf("char")*max_phonemes+1, 0)
	}

	/*-- The first phoneme has to be processed specially. --*/

	for ; !(isalpha(Curr_Letter)); w_idx++ {

		/* On the off chance we were given nothing but crap... */

		if Curr_Letter == '0' {
			if p_idx == max_buffer_len {
				*phoned_word = types.ZendStringExtend(*phoned_word, 1*b.SizeOf("char")+max_buffer_len, 0)
				max_buffer_len += 1
			}
			phoned_word.GetVal()[p_idx] = '0'
			phoned_word.SetLen(p_idx)
			return types.SUCCESS
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
	case 'G':
		fallthrough
	case 'K':
		fallthrough
	case 'P':
		if Next_Letter == 'N' {
			Phonize('N')
			w_idx += 2
		}
	case 'W':
		if Next_Letter == 'R' {
			Phonize(Next_Letter)
			w_idx += 2
		} else if Next_Letter == 'H' || Isvowel(Next_Letter) != 0 {
			Phonize('W')
			w_idx += 2
		}

		/* else ignore */

	case 'X':
		Phonize('S')
		w_idx++
	case 'E':
		fallthrough
	case 'I':
		fallthrough
	case 'O':
		fallthrough
	case 'U':
		Phonize(Curr_Letter)
		w_idx++
	default:

		/* do nothing */

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
		case 'D':
			if Next_Letter == 'G' && MAKESOFT(After_Next_Letter) != 0 {
				Phonize('J')
				skip_letter++
			} else {
				Phonize('T')
			}
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
		case 'H':
			if Isvowel(Next_Letter) != 0 && AFFECTH(Prev_Letter) == 0 {
				Phonize('H')
			}
		case 'K':
			if Prev_Letter != 'C' {
				Phonize('K')
			}
		case 'P':
			if Next_Letter == 'H' {
				Phonize('F')
			} else {
				Phonize('P')
			}
		case 'Q':
			Phonize('K')
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
		case 'T':
			if Next_Letter == 'I' && (After_Next_Letter == 'O' || After_Next_Letter == 'A') {
				Phonize(SH)
			} else if Next_Letter == 'H' {
				Phonize(TH)
				skip_letter++
			} else if !(Next_Letter == 'C' && After_Next_Letter == 'H') {
				Phonize('T')
			}
		case 'V':
			Phonize('F')
		case 'W':
			if Isvowel(Next_Letter) != 0 {
				Phonize('W')
			}
		case 'X':
			Phonize('K')
			Phonize('S')
		case 'Y':
			if Isvowel(Next_Letter) != 0 {
				Phonize('Y')
			}
		case 'Z':
			Phonize('S')
		case 'F':
			fallthrough
		case 'J':
			fallthrough
		case 'L':
			fallthrough
		case 'M':
			fallthrough
		case 'N':
			fallthrough
		case 'R':
			Phonize(Curr_Letter)
		default:

			/* nothing */

		}
		w_idx += skip_letter
	}
	if p_idx == max_buffer_len {
		*phoned_word = types.ZendStringExtend(*phoned_word, 1*b.SizeOf("char")+max_buffer_len, 0)
		max_buffer_len += 1
	}
	phoned_word.GetVal()[p_idx] = '0'
	phoned_word.SetLen(p_idx)
	return 0
}
