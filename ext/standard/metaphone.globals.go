// <<generate>>

package standard

import (
	b "sik/builtin"
)

const SH = 'X'
const TH = '0'

var _codes []byte = []byte{1, 16, 4, 16, 9, 2, 4, 16, 9, 2, 0, 2, 2, 2, 1, 4, 0, 2, 4, 4, 1, 0, 0, 0, 8, 0}

const Next_Letter byte = toupper(word[w_idx+1])
const Curr_Letter byte = toupper(word[w_idx])
const Prev_Letter = Look_Back_Letter(1)
const After_Next_Letter byte = b.CondF1(Next_Letter != '0', func() __auto__ { return toupper(word[w_idx+2]) }, '0')
const Phone_Len = p_idx
