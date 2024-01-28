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

// ParseBool

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
		{"strict-type-undef", strictTester(types.Undef), false, strictError("null")},
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
		{"strict-type-undef", strictTester(types.Undef), false, strictError("null")},
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
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
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
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
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
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseBoolNullable() = %v, want %v", simpleGot, tt.want)
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
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
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
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
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
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseBoolNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseLong

func TestFastParamParser_ParseLong_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), 0, ""},
		{"type-null", weakTester(types.Null), 0, ""},
		{"type-false", weakTester(types.False), 0, ""},
		{"type-true", weakTester(types.True), 1, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), 0, strictError("null")},
		{"strict-type-null", strictTester(types.Null), 0, strictError("null")},
		{"strict-type-false", strictTester(types.False), 0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 0, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseLong()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseLong_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), 0, ""},
		{"type-null", weakTester(types.Null), 0, ""},
		{"type-false", weakTester(types.False), 0, ""},
		{"type-true", weakTester(types.True), 1, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), 0, strictError("null")},
		{"strict-type-null", strictTester(types.Null), 0, strictError("null")},
		{"strict-type-false", strictTester(types.False), 0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 0, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseLong()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseLongNullable_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), 0, ""},
		{"type-true", weakTester(types.True), 1, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), 0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 0, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseLongNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseLongNullable() = %v, want %v", simpleGot, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseLongNullable_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be int, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), 0, ""},
		{"type-true", weakTester(types.True), 1, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), 0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 0, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseLongNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseLongNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseDouble

func TestFastParamParser_ParseDouble_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), 0.0, ""},
		{"type-null", weakTester(types.Null), 0.0, ""},
		{"type-false", weakTester(types.False), 0.0, ""},
		{"type-true", weakTester(types.True), 1.0, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0.0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5.0, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0.0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5.0, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0.0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0.0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0.0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0.0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), 0.0, strictError("null")},
		{"strict-type-null", strictTester(types.Null), 0.0, strictError("null")},
		{"strict-type-false", strictTester(types.False), 0.0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0.0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0.0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5.0, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0.0, ""},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 5.0, ""},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0.0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0.0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0.0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0.0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseDouble()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseDouble_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), 0.0, ""},
		{"type-null", weakTester(types.Null), 0.0, ""},
		{"type-false", weakTester(types.False), 0.0, ""},
		{"type-true", weakTester(types.True), 1.0, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0.0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5.0, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0.0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5.0, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0.0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0.0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0.0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0.0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), 0.0, strictError("null")},
		{"strict-type-null", strictTester(types.Null), 0.0, strictError("null")},
		{"strict-type-false", strictTester(types.False), 0.0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0.0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0.0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5.0, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0.0, ""},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 5.0, ""},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0.0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0.0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0.0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0.0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseDouble()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseDoubleNullable_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), 0.0, ""},
		{"type-true", weakTester(types.True), 1.0, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0.0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5.0, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0.0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5.0, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0.0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0.0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0.0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0.0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), 0.0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0.0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0.0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5.0, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0.0, ""},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 5.0, ""},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0.0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0.0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0.0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0.0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseDoubleNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseDoubleNullable() = %v, want %v", simpleGot, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseDoubleNullable_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be float, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), 0.0, ""},
		{"type-true", weakTester(types.True), 1.0, ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), 0.0, ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), 5.0, ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), 0.0, ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), 5.0, ""},
		{"type-string-0", weakTester(types.ZvalString("")), 0.0, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), 0.0, ""},
		{"type-string-2", weakTester(types.ZvalString("00")), 0.0, ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), 0.0, weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), 0.0, strictError("bool")},
		{"strict-type-true", strictTester(types.True), 0.0, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), 0.0, ""},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), 5.0, ""},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), 0.0, ""},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), 5.0, ""},
		{"strict-type-string-0", strictTester(types.ZvalString("")), 0.0, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), 0.0, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), 0.0, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), 0.0, strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), 0.0, strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseDoubleNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseDoubleNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseString

func TestFastParamParser_ParseString_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), "", ""},
		{"type-null", weakTester(types.Null), "", ""},
		{"type-false", weakTester(types.False), "", ""},
		{"type-true", weakTester(types.True), "1", ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), "0", ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), "5", ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), "0", ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), "5", ""},
		{"type-string-0", weakTester(types.ZvalString("")), "", ""},
		{"type-string-1", weakTester(types.ZvalString("0")), "0", ""},
		{"type-string-2", weakTester(types.ZvalString("00")), "00", ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), "", weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), "", strictError("null")},
		{"strict-type-null", strictTester(types.Null), "", strictError("null")},
		{"strict-type-false", strictTester(types.False), "", strictError("bool")},
		{"strict-type-true", strictTester(types.True), "", strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), "", strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), "", strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), "", strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), "", strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), "", ""},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), "0", ""},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), "00", ""},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), "", strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseString()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseString_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), "", ""},
		{"type-null", weakTester(types.Null), "", ""},
		{"type-false", weakTester(types.False), "", ""},
		{"type-true", weakTester(types.True), "1", ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), "0", ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), "5", ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), "0", ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), "5", ""},
		{"type-string-0", weakTester(types.ZvalString("")), "", ""},
		{"type-string-1", weakTester(types.ZvalString("0")), "0", ""},
		{"type-string-2", weakTester(types.ZvalString("00")), "00", ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), "", weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), "", strictError("null")},
		{"strict-type-null", strictTester(types.Null), "", strictError("null")},
		{"strict-type-false", strictTester(types.False), "", strictError("bool")},
		{"strict-type-true", strictTester(types.True), "", strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), "", strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), "", strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), "", strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), "", strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), "", ""},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), "0", ""},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), "00", ""},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), "", strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseString()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseStringNullable_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), "", ""},
		{"type-true", weakTester(types.True), "1", ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), "0", ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), "5", ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), "0", ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), "5", ""},
		{"type-string-0", weakTester(types.ZvalString("")), "", ""},
		{"type-string-1", weakTester(types.ZvalString("0")), "0", ""},
		{"type-string-2", weakTester(types.ZvalString("00")), "00", ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), "", weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), "", strictError("bool")},
		{"strict-type-true", strictTester(types.True), "", strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), "", strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), "", strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), "", strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), "", strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), "", ""},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), "0", ""},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), "00", ""},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), "", strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseStringNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseStringNullable() = %v, want %v", simpleGot, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseStringNullable_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be string, " + typ + " given"
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, ""},
		{"type-null", weakTester(types.Null), nil, ""},
		{"type-false", weakTester(types.False), "", ""},
		{"type-true", weakTester(types.True), "1", ""},
		{"type-long-0", weakTester(types.ZvalLong(0)), "0", ""},
		{"type-long-1", weakTester(types.ZvalLong(5)), "5", ""},
		{"type-double-0", weakTester(types.ZvalDouble(0)), "0", ""},
		{"type-double-1", weakTester(types.ZvalDouble(5)), "5", ""},
		{"type-string-0", weakTester(types.ZvalString("")), "", ""},
		{"type-string-1", weakTester(types.ZvalString("0")), "0", ""},
		{"type-string-2", weakTester(types.ZvalString("00")), "00", ""},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), "", weakError("array")},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", weakError("array")},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), "", strictError("bool")},
		{"strict-type-true", strictTester(types.True), "", strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), "", strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), "", strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), "", strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), "", strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), "", ""},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), "0", ""},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), "00", ""},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), "", strictError("array")},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), "", strictError("array")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseStringNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParseStringNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}
