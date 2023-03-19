package argparse

import (
	"fmt"
	"sik/zend"
	"sik/zend/types"
)

func (p *OldParser) parseTypeError(arg *types.Zval, expectedType string) {
	p.err = parseError(0, "to be %s, %s given", expectedType, zend.ZendZvalTypeName(arg))
}
func (p *OldParser) parseError(severity int, format string, args ...any) {
	p.err = &parseArgError{severity: severity, message: fmt.Sprintf(format, args...)}
}

func (p *OldParser) parseLong(checkNull bool, cap bool) {
	if val, isNull, ok := ParseLong(p.arg, checkNull, cap); ok {
		p.vaReceiver.Long(val)
		if checkNull {
			p.vaReceiver.Bool(isNull)
		}
	} else {
		p.parseTypeError(p.arg, "int")
	}
}
func (p *OldParser) parseDouble(checkNull bool) {
	if val, isNull, ok := ParseDouble(p.arg, checkNull); ok {
		p.vaReceiver.Double(val)
		if checkNull {
			p.vaReceiver.Bool(isNull)
		}
	} else {
		p.parseTypeError(p.arg, "float")
	}
}
