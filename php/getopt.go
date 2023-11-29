package php

import (
	"fmt"
	"strings"
)

// Opt
type Opt struct {
	char      byte
	needParam int
	name      string
}

func MakeOpt(char byte, needParam int, optName string) Opt {
	return Opt{char: char, needParam: needParam, name: optName}
}
func (opt *Opt) Char() byte          { return opt.char }
func (opt *Opt) Name() string        { return opt.name }
func (opt *Opt) NeedParam() int      { return opt.needParam }
func (opt *Opt) MustNeedParam() bool { return opt.needParam == 1 }

func (opt *Opt) GetOptName() string      { return opt.name }
func (opt *Opt) SetOptChar(value byte)   { opt.char = value }
func (opt *Opt) SetOptName(value string) { opt.name = value }
func (opt *Opt) SetNeedParam(value int)  { opt.needParam = value }
func (opt *Opt) IncNeedParam()           { opt.needParam++ }

// OptError
type OptError string

func (err OptError) Error() string { return string(err) }

func errorColon(argIdx int, charIdx int) error {
	return OptError(fmt.Sprintf("Error in argument %d, char %d: : in flags", argIdx, charIdx))
}
func errorNotFound(argIdx int, charIdx int, char byte) error {
	return OptError(fmt.Sprintf("Error in argument %d, char %d: option not found %c", argIdx, charIdx, char))
}
func errorArg(argIdx int, charIdx int, char byte) error {
	return OptError(fmt.Sprintf("Error in argument %d, char %d: no argument for option %c", argIdx, charIdx, char))
}

// OptsParser
type OptsParser struct {
	args     []string
	opts     []Opt
	idx      int
	startIdx int
}

func (p *OptsParser) Index() int { return p.idx }
func (p *OptsParser) IncIndex()  { p.idx++ }

func NewOptsParser(args []string, opts []Opt, startIdx int) *OptsParser {
	p := &OptsParser{args: args, opts: opts, startIdx: startIdx}
	p.Reset()
	return p
}
func (p *OptsParser) Reset() { p.idx = p.startIdx }
func (p *OptsParser) Next() (*Opt, string, error) {
	return p.getopt()
}
func (p *OptsParser) Each(handler func(opt *Opt, optArg string)) {
	_ = p.EachEx(false, func(opt *Opt, optArg string) error {
		handler(opt, optArg)
		return nil
	})
}
func (p *OptsParser) EachEx(stopOnError bool, handler func(opt *Opt, optArg string) error) error {
	p.Reset()
	for {
		opt, optArg, err := p.getopt()
		if stopOnError && err != nil {
			return err
		} else if err != nil {
			continue
		} else if opt == nil {
			break
		}
		err = handler(opt, optArg)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *OptsParser) RemainArgs() []string {
	if p.idx < len(p.args) {
		return p.args[p.idx:]
	}
	return nil
}

func (p *OptsParser) eof() (*Opt, string, error) { return nil, "", nil }
func (p *OptsParser) getopt() (*Opt, string, error) {
	for p.idx < len(p.args) && p.args[p.idx] == "" {
		p.idx++
	}
	if p.idx >= len(p.args) {
		return p.eof()
	}

	arg := p.args[p.idx]
	if arg[0] != '-' || len(arg) == 1 {
		return p.eof()
	}

	/* '--' indicates end of p.args if not followed by a known long option name */
	if arg == "--" {
		p.idx++
		return p.eof()
	}

	/* Check if the guy tries to do a -: kind of flag */
	if strings.HasPrefix(arg, "-:") {
		p.idx++
		return nil, "", errorColon(p.idx-1, 2)
	}

	var matchOpt *Opt
	var optName, optArg string
	var optchr int

	// parse first arg for long-opt or short-opt
	if strings.HasPrefix(arg, "--") {
		/* Check for <arg>=<val> */
		optchr = 2
		optName, optArg, _ = strings.Cut(arg[2:], "=")
		matchOpt = getOptByName(p.opts, optName)
		p.idx++
		if matchOpt == nil {
			return nil, "", errorArg(p.idx-1, optchr, arg[optchr])
		}
	} else {
		optchr = 1
		optName, optArg = arg[1:2], arg[2:]
		matchOpt = getOptByChar(p.opts, arg[1])
		p.idx++
		if matchOpt == nil {
			return nil, "", errorNotFound(p.idx-1, optchr, arg[optchr])
		}
	}

	// check needParam
	needParam := matchOpt.NeedParam()
	if optArg != "" && needParam != 0 {
		if p.idx < len(p.args) {
			optArg = p.args[p.idx]
			p.idx++
		} else if needParam == 1 { // must required param
			return nil, "", errorArg(p.idx, optchr, arg[optchr])
		}
	}
	if optArg != "" && optArg[0] == '=' {
		optArg = optArg[1:]
	}
	return matchOpt, optArg, nil
}

func getOptByName(opts []Opt, name string) *Opt {
	for i, opt := range opts {
		if opt.Name() == name {
			return &opts[i]
		}
	}
	return nil
}

func getOptByChar(opts []Opt, char byte) *Opt {
	for i, opt := range opts {
		if opt.Char() == char {
			return &opts[i]
		}
	}
	return nil
}
