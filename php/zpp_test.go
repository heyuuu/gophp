package php

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
	"testing"
)

type zppTester struct {
	// options
	minNumArgs int
	maxNumArgs int
	flags      int
	args       []types.Zval
	// runtime
	hasInit bool
	eh      *CollectErrorHandler
	ctx     *Context
	fpp     *FastParamParser
}

func (tester *zppTester) init() {
	if tester.hasInit {
		return
	}
	tester.hasInit = true

	tester.eh = NewCollectErrorHandler()
	ctx := MockContext()
	ctx.eh = tester.eh

	fn := types.NewInternalFunction("mockFunc", nil, 0)
	ex := NewExecuteData(ctx, tester.args, nil)
	ex.fn = fn

	tester.fpp = NewFastParamParser(ex, tester.minNumArgs, tester.maxNumArgs, tester.flags)
}

func (tester *zppTester) parser() *FastParamParser {
	tester.init()
	return tester.fpp
}

func (tester *zppTester) errorLog() string {
	tester.init()
	return tester.eh.String()
}

func (tester *zppTester) checkError(t *testing.T, expectErrLog string) {
	expectErr := expectErrLog != ""
	if got := tester.parser().HasError(); got != expectErr {
		t.Errorf("CheckNumArgs() error = %v, want = %v", got, expectErr)
	}
	if log := tester.errorLog(); strings.TrimSpace(log) != strings.TrimSpace(expectErrLog) {
		t.Errorf("CheckNumArgs() log = %v, want = %v", log, expectErrLog)
	}
}

func TestFastParamParser_CheckNumArgs(t *testing.T) {
	tests := []struct {
		name   string
		tester zppTester
		log    string
	}{
		{
			"success-1",
			zppTester{minNumArgs: 0, maxNumArgs: 0, args: nil},
			"",
		},
		{
			"success-2",
			zppTester{minNumArgs: 1, maxNumArgs: 1, args: make([]types.Zval, 1)},
			"",
		},
		{
			"exactly fail-1",
			zppTester{minNumArgs: 0, maxNumArgs: 0, args: make([]types.Zval, 1)},
			`E_WARNING: mockFunc() expects exactly 0 parameters, 1 given`,
		},
		{
			"exactly fail-2",
			zppTester{minNumArgs: 1, maxNumArgs: 1, args: make([]types.Zval, 2)},
			`E_WARNING: mockFunc() expects exactly 1 parameter, 2 given`,
		},
		{
			"at least-1",
			zppTester{minNumArgs: 1, maxNumArgs: 10, args: make([]types.Zval, 0)},
			`E_WARNING: mockFunc() expects at least 1 parameter, 0 given`,
		},
		{
			"at least-2",
			zppTester{minNumArgs: 2, maxNumArgs: 10, args: make([]types.Zval, 1)},
			`E_WARNING: mockFunc() expects at least 2 parameters, 1 given`,
		},
		{
			"at most-1",
			zppTester{minNumArgs: 0, maxNumArgs: 1, args: make([]types.Zval, 11)},
			`E_WARNING: mockFunc() expects at most 1 parameter, 11 given`,
		},
		{
			"at most-2",
			zppTester{minNumArgs: 2, maxNumArgs: 10, args: make([]types.Zval, 12)},
			`E_WARNING: mockFunc() expects at most 10 parameters, 12 given`,
		},
		{
			"zero-limit",
			zppTester{minNumArgs: 0, maxNumArgs: 0, args: make([]types.Zval, 10)},
			`E_WARNING: mockFunc() expects exactly 0 parameters, 10 given`,
		},
		{
			"zero-no-limit",
			zppTester{minNumArgs: 0, maxNumArgs: -1, args: make([]types.Zval, 10)},
			``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			tester.parser().CheckNumArgs()
			tester.checkError(t, tt.log)
		})
	}
}
