package php

import (
	"github.com/heyuuu/gophp/php/types"
	"testing"
)

func testInitZppParser(args []types.Zval, minNumArgs int, maxNumArgs int, flags int) *FastParamParser {
	ctx := MockContext()
	ex := NewExecuteData(ctx, args, nil)
	fp := NewFastParamParser(ex, minNumArgs, maxNumArgs, flags)
	return fp
}

func TestFastParamParser_CheckNumArgs(t *testing.T) {
	type fields struct {
		minNumArgs int
		maxNumArgs int
		args       []types.Zval
	}
	tests := []struct {
		name     string
		fields   fields
		hasError bool
	}{
		{"-", fields{0, 0, nil}, false},
		{"-", fields{0, 0, []types.Zval{types.Null}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := testInitZppParser(tt.fields.args, tt.fields.minNumArgs, tt.fields.maxNumArgs, 0)
			p.CheckNumArgs()
			if got := p.HasError(); got != tt.hasError {
				t.Errorf("CheckNumArgs() error = %v, want = %v", got, tt.hasError)
			}
		})
	}
}
