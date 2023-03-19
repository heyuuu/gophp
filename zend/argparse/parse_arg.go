package argparse

import (
	"fmt"
)

type parseArgError struct {
	message  string
	severity int
}

func (p parseArgError) Error() string { return p.message }

func (p *TypeSpecParser) parseError(severity int, format string, args ...any) {
	p.err = &parseArgError{severity: severity, message: fmt.Sprintf(format, args...)}
}

func (p *TypeSpecParser) parseLong(checkNull bool, cap bool) {
	if val, isNull, ok := ParseLong(p.arg, checkNull, cap); ok {
		p.vaReceiver.Long(val)
		if checkNull {
			p.vaReceiver.Bool(isNull)
		}
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}
}
func (p *TypeSpecParser) parseDouble(checkNull bool) {
	if val, isNull, ok := ParseDouble(p.arg, checkNull); ok {
		p.vaReceiver.Double(val)
		if checkNull {
			p.vaReceiver.Bool(isNull)
		}
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
	}
}
