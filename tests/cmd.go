package tests

import "strings"

type command struct {
	bin  string
	args []string
}

func newCommand(bin string) *command {
	return &command{bin: bin}
}

func (c *command) String() string {
	return c.bin + " " + strings.Join(c.args, " ")
}

func (c *command) arg(args ...string) *command {
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}
		c.args = append(c.args, arg)
	}
	return c
}

func (c *command) wrapArg(arg string) *command {
	c.args = append(c.args, `"`+arg+`"`)
	return c
}
