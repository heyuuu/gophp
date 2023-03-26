package standard

import (
	"sik/zend"
)

const PhpMySetlocale = setlocale

/* For str_getcsv() support */

/* For php_next_utf8_char() */

const STR_PAD_LEFT = 0
const STR_PAD_RIGHT = 1
const STR_PAD_BOTH = 2
const PHP_PATHINFO_DIRNAME = 1
const PHP_PATHINFO_BASENAME = 2
const PHP_PATHINFO_EXTENSION = 4
const PHP_PATHINFO_FILENAME = 8
const PHP_PATHINFO_ALL zend.ZendLong = PHP_PATHINFO_DIRNAME | PHP_PATHINFO_BASENAME | PHP_PATHINFO_EXTENSION | PHP_PATHINFO_FILENAME
const STR_STRSPN = 0
const STR_STRCSPN = 1

/* this is read-only, so it's ok */

var Hexconvtab = "0123456789abcdef"

/* localeconv mutex */

const _HEB_BLOCK_TYPE_ENG = 1
const _HEB_BLOCK_TYPE_HEB = 2

const PHP_TAG_BUF_SIZE = 1023

/* {{{ php_tag_find
 *
 * Check if tag is in a set of tags
 *
 * states:
 *
 * 0 start tag
 * 1 first non-whitespace char seen
 */
