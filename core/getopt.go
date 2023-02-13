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

type OptValue struct {
	*Opt
	Value string
}
type OptArgs struct {
	OptionValues []OptValue
	Arguments    []string
}

func (o *OptArgs) addOptionValue(opt *Opt, value string) {
	o.OptionValues = append(o.OptionValues, OptValue{Opt: opt, Value: value})
}
func (o *OptArgs) addArgument(argument string) {
	o.Arguments = append(o.Arguments, argument)
}

func GetOpts(args []string, opts []Opt) (optArgs OptArgs, err error) {
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
			var optName, optValue string
			pair := strings.SplitN(arg[2:], "=", 2)
			optName = pair[0]
			if len(pair) == 2 {
				optValue = pair[1]
			}

			// 查找对应 opt
			if opt, ok := getOptByName(opts, optName); ok {
				if len(optValue) == 0 && i+1 < argc && args[i+1][0] != '-' {
					i++
					optValue = args[i]
				}
				optArgs.addOptionValue(opt, optValue)
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

				optArgs.addOptionValue(opt, val)
			} else {
				optArgs.addOptionValue(opt, "")

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
					optArgs.addOptionValue(opt, "")
				}
			}
		}

		// 非 option 情况
		optArgs.addArgument(arg)
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
