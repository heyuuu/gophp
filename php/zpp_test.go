package php

import (
	"github.com/heyuuu/gophp/php/lang"
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
	strict     bool
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
	if tester.strict {
		tester.fpp.strictType = true
	}
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

func TestFastParamParser_ParseBool_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   bool
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), false, ""},
		{"type-null", weakTester(types.Null), false, ""},
		{"type-false", weakTester(types.False), false, ""},
		{"type-true", weakTester(types.True), true, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), false, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), true, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), false, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), true, ""},
		{"type-string-0", weakTester(types.ZvalString("")), false, ""},
		{"type-string-1", weakTester(types.ZvalString("0")), false, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), true, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), false, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), false, strictError("unknown")},
		{"strict-type-null", strictTester(types.Null), false, strictError("null")},
		{"strict-type-false", strictTester(types.False), false, ""},
		{"strict-type-true", strictTester(types.True), true, ""},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), false, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), false, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), false, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), false, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), false, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), false, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), false, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), false, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseBool()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseBool_reftype(t *testing.T) {
	refArgs := func(args ...types.Zval) []types.Zval {
		refArgs := make([]types.Zval, len(args))
		for i, arg := range args {
			refArgs[i] = types.ZvalRef(types.NewReference(arg))
		}
		return refArgs
	}
	weakTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: false}
	}
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   bool
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), false, ""},
		{"type-null", weakTester(types.Null), false, ""},
		{"type-false", weakTester(types.False), false, ""},
		{"type-true", weakTester(types.True), true, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), false, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), true, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), false, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), true, ""},
		{"type-string-0", weakTester(types.ZvalString("")), false, ""},
		{"type-string-1", weakTester(types.ZvalString("0")), false, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), true, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), false, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), false, strictError("unknown")},
		{"strict-type-null", strictTester(types.Null), false, strictError("null")},
		{"strict-type-false", strictTester(types.False), false, ""},
		{"strict-type-true", strictTester(types.True), true, ""},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), false, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), false, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), false, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), false, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), false, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), false, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), false, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), false, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseBool()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseBoolNullable_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), false, ""},
		{"type-null", weakTester(types.Null), false, ""},
		{"type-false", weakTester(types.False), false, ""},
		{"type-true", weakTester(types.True), true, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), false, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), true, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), false, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), true, ""},
		{"type-string-0", weakTester(types.ZvalString("")), false, ""},
		{"type-string-1", weakTester(types.ZvalString("0")), false, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), true, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), false, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), false, strictError("unknown")},
		{"strict-type-null", strictTester(types.Null), false, strictError("null")},
		{"strict-type-false", strictTester(types.False), false, ""},
		{"strict-type-true", strictTester(types.True), true, ""},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), false, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), false, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), false, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), false, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), false, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), false, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), false, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), false, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseBoolNullable()
			tester.checkError(t, tt.log)
			var simpleGot = lang.Cond(got != nil, any(*got), any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseBool() = %v, want %v", simpleGot, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseBoolNullable_reftype(t *testing.T) {
	refArgs := func(args ...types.Zval) []types.Zval {
		refArgs := make([]types.Zval, len(args))
		for i, arg := range args {
			refArgs[i] = types.ZvalRef(types.NewReference(arg))
		}
		return refArgs
	}
	weakTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: false}
	}
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be bool, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), false, ""},
		{"type-null", weakTester(types.Null), false, ""},
		{"type-false", weakTester(types.False), false, ""},
		{"type-true", weakTester(types.True), true, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), false, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), true, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), false, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), true, ""},
		{"type-string-0", weakTester(types.ZvalString("")), false, ""},
		{"type-string-1", weakTester(types.ZvalString("0")), false, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), true, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), false, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), false, strictError("unknown")},
		{"strict-type-null", strictTester(types.Null), false, strictError("null")},
		{"strict-type-false", strictTester(types.False), false, ""},
		{"strict-type-true", strictTester(types.True), true, ""},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), false, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), false, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), false, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), false, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), false, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), false, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), false, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), false, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), false, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseBoolNullable()
			tester.checkError(t, tt.log)
			var simpleGot = lang.Cond(got != nil, any(*got), any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
