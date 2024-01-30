package php

type ctxWriter struct {
	ctx *Context
}

func (w ctxWriter) WriteString(s string) (n int, err error) {
	w.ctx.WriteString(s)
	return len(s), nil
}

func (w ctxWriter) Write(p []byte) (n int, err error) {
	w.ctx.Write(p)
	return len(p), nil
}
