package core

/* Define structure for one recognized option (both single char and long name).
 * If short_open is '-' this is the last option. */

type _opt = Opt

/* holds the index of the latest fetched element from the opts array */

/* php_getopt will return this value if there is an error in arguments */

const PHP_GETOPT_INVALID_ARG = -2
