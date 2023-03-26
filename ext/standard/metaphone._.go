package standard

import (
	b "sik/builtin"
)

const SH = 'X'
const TH = '0'

/*-----------------------------  */

var _codes []byte = []byte{1, 16, 4, 16, 9, 2, 4, 16, 9, 2, 0, 2, 2, 2, 1, 4, 0, 2, 4, 4, 1, 0, 0, 0, 8, 0}

/* These letters are passed through unchanged */

/* These form diphthongs when preceding H */

/* These make C and G soft */

/* These prevent GH from becoming F */

/*----------------------------- */

const Next_Letter byte = toupper(word[w_idx+1])

/* Look at the current letter in the word */

const Curr_Letter byte = toupper(word[w_idx])

/* Go N letters back. */

/* Previous letter.  I dunno, should this return null on failure? */

const Prev_Letter = Look_Back_Letter(1)

/* Look two letters down.  It makes sure you don't walk off the string. */

const After_Next_Letter byte = b.CondF1(Next_Letter != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0')

/* Allows us to safely look ahead an arbitrary # of letters */

/* phonize one letter
 * We don't know the buffers size in advance. On way to solve this is to just
 * re-allocate the buffer size. We're using an extra of 2 characters (this
 * could be one though; or more too). */

/* Slap a null character on the end of the phoned word */

/* How long is the phoned word? */

const Phone_Len = p_idx

/* Note is a letter is a 'break' in the word */

/* {{{ metaphone
 */
