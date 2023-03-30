package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

const SCAN_MAX_ARGS = 0xff

/* passed to (f|s)scanf. This is an artificial   */

const SCAN_SUCCESS = types.SUCCESS
const SCAN_ERROR_EOF = -1

/* can be caused by bad parameters or format*/

const SCAN_ERROR_INVALID_FORMAT = SCAN_ERROR_EOF - 1
const SCAN_ERROR_VAR_PASSED_BYVAL = SCAN_ERROR_INVALID_FORMAT - 1
const SCAN_ERROR_WRONG_PARAM_COUNT = SCAN_ERROR_VAR_PASSED_BYVAL - 1
const SCAN_ERROR_INTERNAL = SCAN_ERROR_WRONG_PARAM_COUNT - 1

/*
 * The following are here solely for the benefit of the scanf type functions
 * e.g. fscanf
 */

/*
 * Flag values used internally by [f|s]canf.
 */

const SCAN_NOSKIP = 0x1
const SCAN_SUPPRESS = 0x2
const SCAN_UNSIGNED = 0x4
const SCAN_WIDTH = 0x8
const SCAN_SIGNOK = 0x10
const SCAN_NODIGITS = 0x20
const SCAN_NOZERO = 0x40
const SCAN_XOK = 0x80
const SCAN_PTOK = 0x100
const SCAN_EXPOK = 0x200

/*
 * The following structure contains the information associated with
 * a character set.
 */

/*
 * Declarations for functions used only in this file.
 */

/* {{{ BuildCharSet
 *----------------------------------------------------------------------
 *
 * BuildCharSet --
 *
 *    This function examines a character set format specification
 *    and builds a CharSet containing the individual characters and
 *    character ranges specified.
 *
 * Results:
 *    Returns the next format position.
 *
 * Side effects:
 *    Initializes the charset.
 *
 *----------------------------------------------------------------------
 */
