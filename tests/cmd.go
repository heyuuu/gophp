package tests

import (
	"github.com/heyuuu/gophp/sapi"
	"io"
	"os"
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
	useCgi        bool
	args          []commandArg
	stdin         string `set:""`
	stdinFile     string `set:""`
	captureStdIn  bool
	captureStdOut bool
	captureStdErr bool
}

func commandBuild(bin string, useCgi bool) *command {
	cmd := &command{
		bin:           bin,
		useCgi:        useCgi,
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
	c.args = slices.Grow(c.args, len(args))
	for _, cmdArg := range args {
		if cmdArg.value != "" {
			c.args = append(c.args, cmdArg)
		}
	}
	return c
}

func (c *command) option(opt string, val string) *command {
	return c.add(arg(opt), quoteArg(val))
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

func (c *command) CliString() string {
	var buf strings.Builder
	buf.Grow(1024)

	buf.WriteString(c.bin)
	if c.useCgi {
		buf.WriteString(" -C")
	}
	for _, cmdArg := range c.args {
		if cmdArg.value != "" {
			buf.WriteByte(' ')
			buf.WriteString(cmdArg.String())
		}
	}

	if c.captureStdOut && c.captureStdErr {
		buf.WriteString(" 2>&1")
	}
	if c.stdinFile != "" {
		buf.WriteString(` < "` + c.stdinFile + `"`)
	}

	return buf.String()
}

func (c *command) cmdArgs() []string {
	args := make([]string, 0, len(c.args)+1)
	if c.useCgi {
		args = append(args, "-C")
	}
	for _, cmdArg := range c.args {
		args = append(args, cmdArg.value)
	}
	return args
}

func (c *command) Run() (output string, err error) {
	if c.bin == "" {
		return c.runBuiltin()
	} else {
		return c.runExec()
	}
}

func (c *command) prepareStdin() (stdin io.Reader, err error) {
	if c.stdin != "" {
		return strings.NewReader(c.stdin), nil
	} else if c.stdinFile != "" {
		return os.Open(c.stdinFile)
	}
	return nil, nil
}

func (c *command) runBuiltin() (output string, err error) {
	cmd := sapi.Command(c.cmdArgs()...)

	cmd.Stdin, err = c.prepareStdin()
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	if c.captureStdOut {
		cmd.Stdout = &buf
	}
	if c.captureStdErr {
		cmd.Stderr = &buf
	}

	err = cmd.RunSafe()
	return buf.String(), err
}

func (c *command) runExec() (output string, err error) {
	cmd := exec.Command("bash", "-c", c.CliString())

	if c.stdin != "" {
		cmd.Stdin = strings.NewReader(c.stdin)
	}

	var buf strings.Builder
	cmd.Stdout = &buf

	err = cmd.Run()
	return buf.String(), err
}
