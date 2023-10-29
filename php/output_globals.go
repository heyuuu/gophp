package php

import (
	"io"
	"os"
)

// todo output 临时实现，待完善

type OutputGlobals struct {
	active io.Writer
}

func (g *OutputGlobals) writer() io.Writer {
	if g.active != nil {
		return g.active
	}
	return os.Stdout
}

func (g *OutputGlobals) Write(data []byte) {
	g.writer().Write(data)
}

func (g *OutputGlobals) WriteString(str string) {
	io.WriteString(g.writer(), str)
}

func (g *OutputGlobals) WriteStringUnbuffered(str string) {
	io.WriteString(g.writer(), str)
}

func (g *OutputGlobals) PushHandler(handler io.Writer) {
	g.active = handler
}
