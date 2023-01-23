// <<generate>>

package core

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/getopt.c>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < string . h >

// # include < assert . h >

// # include < stdlib . h >

// # include "php_getopt.h"

// #define OPTERRCOLON       ( 1 )

// #define OPTERRNF       ( 2 )

// #define OPTERRARG       ( 3 )

// Print error message to stderr and return -2 to distinguish it from '?' command line option.

func PhpOptError(argc int, argv **byte, oint int, optchr int, err int, show_err int) int {
	if show_err != 0 {
		r.Fprintf(stderr, "Error in argument %d, char %d: ", oint, optchr+1)
		switch err {
		case 1:
			r.Fprintf(stderr, ": in flags\n")
			break
		case 2:
			r.Fprintf(stderr, "option not found %c\n", argv[oint][optchr])
			break
		case 3:
			r.Fprintf(stderr, "no argument for option %c\n", argv[oint][optchr])
			break
		default:
			r.Fprintf(stderr, "unknown\n")
			break
		}
	}
	return -2
}

/* }}} */

var PhpOptidx int = -1

func PhpGetopt(argc int, argv **byte, opts []Opt, optarg **byte, optind *int, show_err int, arg_start int) int {
	var optchr int = 0
	var dash int = 0
	var prev_optarg **byte = nil
	PhpOptidx = -1
	if prev_optarg != nil && prev_optarg != optarg {

		/* reset the state */

		optchr = 0
		dash = 0
	}
	prev_optarg = optarg
	if (*optind) >= argc {
		return -1
	}
	if dash == 0 {
		if argv[*optind][0] != '-' {
			return -1
		} else {
			if !(argv[*optind][1]) {

				/*
				 * use to specify stdin. Need to let pgm process this and
				 * the following args
				 */

				return -1

				/*
				 * use to specify stdin. Need to let pgm process this and
				 * the following args
				 */

			}
		}
	}
	if argv[*optind][0] == '-' && argv[*optind][1] == '-' {
		var pos *byte
		var arg_end int = strlen(argv[*optind]) - 1

		/* '--' indicates end of args if not followed by a known long option name */

		if argv[*optind][2] == '0' {
			*optind++
			return -1
		}
		arg_start = 2

		/* Check for <arg>=<val> */

		if g.Assign(&pos, zend.ZendMemnstr(&argv[*optind][arg_start], "=", 1, argv[*optind]+arg_end)) != nil {
			arg_end = pos - &argv[*optind][arg_start]
			arg_start++
		} else {
			arg_end--
		}
		for true {
			PhpOptidx++
			if opts[PhpOptidx].GetOptChar() == '-' {
				*optind++
				return PhpOptError(argc, argv, (*optind)-1, optchr, 3, show_err)
			} else if opts[PhpOptidx].GetOptName() != nil && !(strncmp(&argv[*optind][2], opts[PhpOptidx].GetOptName(), arg_end)) && arg_end == strlen(opts[PhpOptidx].GetOptName()) {
				break
			}
		}
		optchr = 0
		dash = 0
		arg_start += int(strlen(opts[PhpOptidx].GetOptName()))
	} else {
		if dash == 0 {
			dash = 1
			optchr = 1
		}

		/* Check if the guy tries to do a -: kind of flag */

		if argv[*optind][optchr] == ':' {
			dash = 0
			*optind++
			return PhpOptError(argc, argv, (*optind)-1, optchr, 1, show_err)
		}
		arg_start = 1 + optchr
	}
	if PhpOptidx < 0 {
		for true {
			PhpOptidx++
			if opts[PhpOptidx].GetOptChar() == '-' {
				var errind int = *optind
				var errchr int = optchr
				if !(argv[*optind][optchr+1]) {
					dash = 0
					*optind++
				} else {
					optchr++
					arg_start++
				}
				return PhpOptError(argc, argv, errind, errchr, 2, show_err)
			} else if argv[*optind][optchr] == opts[PhpOptidx].GetOptChar() {
				break
			}
		}
	}
	if opts[PhpOptidx].GetNeedParam() != 0 {

		/* Check for cases where the value of the argument
		   is in the form -<arg> <val>, -<arg>=<varl> or -<arg><val> */

		dash = 0
		if !(argv[*optind][arg_start]) {
			*optind++
			if (*optind) == argc {

				/* Was the value required or is it optional? */

				if opts[PhpOptidx].GetNeedParam() == 1 {
					return PhpOptError(argc, argv, (*optind)-1, optchr, 3, show_err)
				}

				/* Was the value required or is it optional? */

			} else if opts[PhpOptidx].GetNeedParam() == 1 {
				*optarg = argv[g.PostInc(&(*optind))]
			}
		} else if argv[*optind][arg_start] == '=' {
			arg_start++
			*optarg = &argv[*optind][arg_start]
			*optind++
		} else {
			*optarg = &argv[*optind][arg_start]
			*optind++
		}
		return opts[PhpOptidx].GetOptChar()
	} else {

		/* multiple options specified as one (exclude long opts) */

		if arg_start >= 2 && !(argv[*optind][0] == '-' && argv[*optind][1] == '-') {
			if !(argv[*optind][optchr+1]) {
				dash = 0
				*optind++
			} else {
				optchr++
			}
		} else {
			*optind++
		}
		return opts[PhpOptidx].GetOptChar()
	}
	r.Assert(false)
	return 0
}

/* }}} */
