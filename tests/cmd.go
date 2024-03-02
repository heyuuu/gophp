package tests

import (
	"context"
	"github.com/heyuuu/gophp/kits/slicekit"
	"os/exec"
	"slices"
	"strings"
)

type commandArg struct {
	value string
	quote bool
}

func arg(s string) commandArg      { return commandArg{s, false} }
func quoteArg(s string) commandArg { return commandArg{s, true} }

func (c commandArg) String() string {
	if c.quote {
		return `"` + c.value + `"`
	}
	return c.value
}

func CommandArgsToString(args []commandArg) string {
	var buf strings.Builder
	for i, cmdArg := range args {
		if i != 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(cmdArg.String())
	}
	return buf.String()
}

type command struct {
	bin           string
	args          []commandArg
	captureStdIn  bool
	captureStdOut bool
	captureStdErr bool
}

func newCommand(bin string, args ...commandArg) *command {
	cmd := &command{
		bin:           bin,
		args:          slices.Clone(args),
		captureStdIn:  true,
		captureStdOut: true,
		captureStdErr: true,
	}
	return cmd
}

func (c *command) clone() *command {
	dup := *c
	dup.args = slices.Clone(dup.args)
	return &dup
}

func (c *command) add(args ...commandArg) *command {
	c.args = append(c.args, args...)
	return c
}

func (c *command) option(opt string, val string) *command {
	c.args = append(c.args, arg(opt), quoteArg(val))
	return c
}

func (c *command) addIniSettings(ini *IniSettings) *command {
	c.args = slices.Grow(c.args, ini.Len()*2)
	ini.Each(func(key, val string) {
		c.option("-d", key+"="+addslashes(val))
	})
	return c
}

func (c *command) capture(captureStdIn bool, captureStdOut bool, captureStdErr bool) {
	c.captureStdIn = captureStdIn
	c.captureStdOut = captureStdOut
	c.captureStdErr = captureStdErr
}

func (c *command) String() string {
	var buf strings.Builder
	buf.Grow(1024)

	buf.WriteString(c.bin)
	for _, cmdArg := range c.args {
		if cmdArg.value != "" {
			buf.WriteByte(' ')
			buf.WriteString(cmdArg.String())
		}
	}

	if c.captureStdOut && c.captureStdErr {
		buf.WriteString(" 2>&1")
	}

	return buf.String()
}

func (c *command) Run() (string, error) {
	return c.RunEx(nil)
}

func (c *command) RunEx(ctx context.Context) (string, error) {
	args := slicekit.Map(c.args, func(t commandArg) string {
		return t.value
	})

	var cmd *exec.Cmd
	if ctx != nil {
		cmd = exec.CommandContext(ctx, c.bin, args...)
	} else {
		cmd = exec.Command(c.bin, args...)
	}

	var buf strings.Builder
	if c.captureStdOut {
		cmd.Stdout = &buf
	}
	if c.captureStdErr {
		cmd.Stderr = &buf
	}

	err := cmd.Run()
	return buf.String(), err
}
