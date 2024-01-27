package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

type ErrorHandling interface {
	OnError(typ perr.ErrorType, message string)
	OnException(ce *types.Class, message string, code int)
}

type CollectErrorHandler struct {
	buf strings.Builder
}

func NewCollectErrorHandler() *CollectErrorHandler {
	return &CollectErrorHandler{}
}

func (t *CollectErrorHandler) OnError(typ perr.ErrorType, message string) {
	_, _ = fmt.Fprintf(&t.buf, "%s: %s\n", typ.String(), message)
}

func (t *CollectErrorHandler) OnException(ce *types.Class, message string, code int) {
	_, _ = fmt.Fprintf(&t.buf, "Exception: (%d) %s\n", code, message)
}

func (t *CollectErrorHandler) Reset() {
	t.buf.Reset()
}

func (t *CollectErrorHandler) String() string {
	return t.buf.String()
}
