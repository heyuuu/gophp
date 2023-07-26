package core

import (
	r "github.com/heyuuu/gophp/builtin/file"
	b "github.com/heyuuu/gophp/php/lang"
	"log"
)

func PhpOptError(
	argc int,
	argv **byte,
	oint int,
	optchr int,
	err int,
	show_err int,
) int {
	if show_err != 0 {
		log.Printf("Error in argument %d, char %d: ", oint, optchr+1)
		switch err {
		case OPTERRCOLON:
			log.Printf(": in flags\n")
		case OPTERRNF:
			log.Printf("option not found %c\n", argv[oint][optchr])
		case OPTERRARG:
			log.Printf("no argument for option %c\n", argv[oint][optchr])
		default:
			log.Printf("unknown\n")
		}
	}
	return PHP_GETOPT_INVALID_ARG
}
func PhpGetopt(
	argc int,
	argv **byte,
	opts []Opt,
	optarg **byte,
	optind *int,
	show_err int,
	arg_start int,
) int {
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
		return r.EOF
	}
	if dash == 0 {
		if argv[*optind][0] != '-' {
			return r.EOF
		} else {
			if !(argv[*optind][1]) {

				/*
				 * use to specify stdin. Need to let pgm process this and
				 * the following args
				 */

				return r.EOF

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
			return r.EOF
		}
		arg_start = 2

		/* Check for <arg>=<val> */

		if b.Assign(&pos, PhpMemnstr(&argv[*optind][arg_start], "=", 1, argv[*optind]+arg_end)) != nil {
			arg_end = pos - &argv[*optind][arg_start]
			arg_start++
		} else {
			arg_end--
		}
		for true {
			PhpOptidx++
			if opts[PhpOptidx].GetOptChar() == '-' {
				*optind++
				return PhpOptError(argc, argv, (*optind)-1, optchr, OPTERRARG, show_err)
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
			return PhpOptError(argc, argv, (*optind)-1, optchr, OPTERRCOLON, show_err)
		}
		arg_start = 1 + optchr
	}
	if PhpOptidx < 0 {
		for {
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
				return PhpOptError(argc, argv, errind, errchr, OPTERRNF, show_err)
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
					return PhpOptError(argc, argv, (*optind)-1, optchr, OPTERRARG, show_err)
				}

				/* Was the value required or is it optional? */

			} else if opts[PhpOptidx].GetNeedParam() == 1 {
				*optarg = argv[b.PostInc(&(*optind))]
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
}
