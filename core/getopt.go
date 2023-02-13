package core

import (
	"fmt"
	"strings"
)

var _ error = OptError("")

type OptError string

func (err OptError) Error() string { return string(err) }

func optErrorColon() error          { return OptError(": in flags") }
func optErrorNotFound(c byte) error { return OptError(fmt.Sprintf("option not found %c", c)) }
func optErrorArg(c byte) error      { return OptError(fmt.Sprintf("no argument for option %c", c)) }

type OptArg struct {
	*Opt
	Value string
}

func GetOpts(args []string, opts []Opt) (result []OptArg, err error) {
	argc := len(args)
	for i := 0; i < argc; i++ {
		arg := args[i]

		/* '--' indicates end of args if not followed by a known long option name */
		if arg == "--" {
			return
		}

		// longOpt
		if strings.HasPrefix(arg, "--") {
			/* Check for <arg>=<val> */
			var argName, argValue string
			pair := strings.SplitN(arg[2:], "=", 2)
			argName = pair[0]
			if len(pair) == 2 {
				argValue = pair[1]
			}

			// 查找对应 opt
			if opt, ok := getOptByName(opts, argName); ok {
				if len(argValue) == 0 && i+1 < argc && args[i+1][0] != '-' {
					i++
					argValue = args[i]
				}
				result = append(result, OptArg{opt, argValue})
				continue
			} else {
				err = optErrorArg('-')
				return
			}
		}

		/* Check if the guy tries to do a -: kind of flag */
		if strings.HasPrefix(arg, "-:") {
			err = optErrorColon()
			return
		}

		// shortopt
		if strings.HasPrefix(arg, "-") {
			c := arg[1]
			opt, ok := getOptByChar(opts, c)
			if !ok {
				err = optErrorNotFound(c)
				return
			}

			if opt.NeedParam {
				/* Check for cases where the value of the argument
				   is in the form -<arg> <val>, -<arg>=<val> or -<arg><val> */

				var val string
				if len(arg) > 2 {
					val = strings.TrimLeft(arg[2:], "=")
				} else if i+1 < argc && args[i+1][0] != '-' {
					i++
					val = args[i]
				} else {
					err = optErrorArg(c)
					return
				}

				result = append(result, OptArg{opt, val})
			} else {
				result = append(result, OptArg{opt, ""})

				/* multiple options specified as one (exclude long opts) */
				for i := range arg[3:] {
					c := arg[i]
					opt, ok := getOptByChar(opts, c)
					if !ok {
						err = optErrorNotFound(c)
						return
					}
					if opt.NeedParam {
						err = optErrorArg(c)
						return
					}
					result = append(result, OptArg{opt, ""})
				}
			}
		}
	}
	return
}

func getOptByName(opts []Opt, name string) (*Opt, bool) {
	for i, opt := range opts {
		if opt.Name == name {
			return &opts[i], true
		}
	}
	return nil, false
}

func getOptByChar(opts []Opt, char byte) (*Opt, bool) {
	for i, opt := range opts {
		if opt.Char == char {
			return &opts[i], true
		}
	}
	return nil, false
}
