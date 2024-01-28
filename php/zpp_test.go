package php

import (
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"reflect"
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
		t.Errorf("Parser.HasError() = %v, want = %v", got, expectErr)
	}
	if log := tester.errorLog(); strings.TrimSpace(log) != strings.TrimSpace(expectErrLog) {
		t.Errorf("errorLog = %v, want = %v", log, expectErrLog)
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

// ParsePath

func TestFastParamParser_ParsePath_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
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
			got := tester.parser().ParsePath()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParsePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParsePath_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
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
			got := tester.parser().ParsePath()
			tester.checkError(t, tt.log)
			if got != tt.want {
				t.Errorf("ParsePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParsePathNullable_basetype(t *testing.T) {
	weakTester := func(args ...types.Zval) zppTester { return zppTester{args: args, strict: false} }
	weakError := func(typ string) string {
		return "E_WARNING: mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
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
			got := tester.parser().ParsePathNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParsePathNullable() = %v, want %v", simpleGot, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParsePathNullable_reftype(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be a valid path, " + typ + " given"
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
			got := tester.parser().ParsePathNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return *got }, any(nil))
			if simpleGot != tt.want {
				t.Errorf("ParsePathNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseArray

func TestFastParamParser_ParseArray(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be array, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be array, " + typ + " given"
	}
	values := func(values ...types.Zval) []types.ArrayPair {
		var pairs = make([]types.ArrayPair, len(values))
		for i, value := range values {
			pairs[i] = types.MakeArrayPair(types.IdxKey(i), value)
		}
		return pairs
	}
	tests := []struct {
		name   string
		tester zppTester
		want   any
		log    string
	}{
		// weak
		{"type-undef", weakTester(types.Undef), nil, weakError("null")},
		{"type-null", weakTester(types.Null), nil, weakError("null")},
		{"type-false", weakTester(types.False), nil, weakError("bool")},
		{"type-true", weakTester(types.True), nil, weakError("bool")},
		{"type-long-0", weakTester(types.ZvalLong(0)), nil, weakError("int")},
		{"type-long-1", weakTester(types.ZvalLong(5)), nil, weakError("int")},
		{"type-double-0", weakTester(types.ZvalDouble(0)), nil, weakError("float")},
		{"type-double-1", weakTester(types.ZvalDouble(5)), nil, weakError("float")},
		{"type-string-0", weakTester(types.ZvalString("")), nil, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), nil, weakError("string")},
		{"type-string-2", weakTester(types.ZvalString("00")), nil, weakError("string")},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), values(), ""},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), values(types.Null), ""},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, strictError("null")},
		{"strict-type-null", strictTester(types.Null), nil, strictError("null")},
		{"strict-type-false", strictTester(types.False), nil, strictError("bool")},
		{"strict-type-true", strictTester(types.True), nil, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), nil, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), nil, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), nil, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), nil, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), nil, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), nil, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), nil, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), values(), ""},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), values(types.Null), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseArray()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return got.Pairs() }, any(nil))
			if !reflect.DeepEqual(simpleGot, tt.want) {
				t.Errorf("ParseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseArrayNullable(t *testing.T) {
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
		return "E_WARNING: mockFunc() expects parameter 1 to be array, " + typ + " given"
	}
	strictTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: true}
	}
	strictError := func(typ string) string {
		return "Exception: (0) mockFunc() expects parameter 1 to be array, " + typ + " given"
	}
	values := func(values ...types.Zval) []types.ArrayPair {
		var pairs = make([]types.ArrayPair, len(values))
		for i, value := range values {
			pairs[i] = types.MakeArrayPair(types.IdxKey(i), value)
		}
		return pairs
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
		{"type-false", weakTester(types.False), nil, weakError("bool")},
		{"type-true", weakTester(types.True), nil, weakError("bool")},
		{"type-long-0", weakTester(types.ZvalLong(0)), nil, weakError("int")},
		{"type-long-1", weakTester(types.ZvalLong(5)), nil, weakError("int")},
		{"type-double-0", weakTester(types.ZvalDouble(0)), nil, weakError("float")},
		{"type-double-1", weakTester(types.ZvalDouble(5)), nil, weakError("float")},
		{"type-string-0", weakTester(types.ZvalString("")), nil, weakError("string")},
		{"type-string-1", weakTester(types.ZvalString("0")), nil, weakError("string")},
		{"type-string-2", weakTester(types.ZvalString("00")), nil, weakError("string")},
		{"type-array-0", weakTester(types.ZvalArray(types.NewArray())), values(), ""},
		{"type-array-1", weakTester(types.ZvalArray(types.NewArrayOf(types.Null))), values(types.Null), ""},
		// strict
		{"strict-type-undef", strictTester(types.Undef), nil, ""},
		{"strict-type-null", strictTester(types.Null), nil, ""},
		{"strict-type-false", strictTester(types.False), nil, strictError("bool")},
		{"strict-type-true", strictTester(types.True), nil, strictError("bool")},
		{"strict-type-long-0", strictTester(types.ZvalLong(0)), nil, strictError("int")},
		{"strict-type-long-1", strictTester(types.ZvalLong(5)), nil, strictError("int")},
		{"strict-type-double-0", strictTester(types.ZvalDouble(0)), nil, strictError("float")},
		{"strict-type-double-1", strictTester(types.ZvalDouble(5)), nil, strictError("float")},
		{"strict-type-string-0", strictTester(types.ZvalString("")), nil, strictError("string")},
		{"strict-type-string-1", strictTester(types.ZvalString("0")), nil, strictError("string")},
		{"strict-type-string-2", strictTester(types.ZvalString("00")), nil, strictError("string")},
		{"strict-type-array-0", strictTester(types.ZvalArray(types.NewArray())), values(), ""},
		{"strict-type-array-1", strictTester(types.ZvalArray(types.NewArrayOf(types.Null))), values(types.Null), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseArrayNullable()
			tester.checkError(t, tt.log)
			var simpleGot any = lang.CondF1(got != nil, func() any { return got.Pairs() }, any(nil))
			if !reflect.DeepEqual(simpleGot, tt.want) {
				t.Errorf("ParseArrayNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseZval
func testFastParamParser_deepEqualsZval(v1, v2 types.Zval) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	if v1.IsArray() && v2.IsArray() {
		return reflect.DeepEqual(v1.Array().Pairs(), v2.Array().Pairs())
	}
	return false
}
func testFastParamParser_deepEqualsZvalPtr(v1, v2 *types.Zval) bool {
	if v1 == nil || v2 == nil {
		return v1 == v2
	}
	return testFastParamParser_deepEqualsZval(*v1, *v2)
}

func TestFastParamParser_ParseZval(t *testing.T) {
	refArgs := func(args ...types.Zval) []types.Zval {
		refArgs := make([]types.Zval, len(args))
		for i, arg := range args {
			refArgs[i] = types.ZvalRef(types.NewReference(arg))
		}
		return refArgs
	}
	baseTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: false}
	}
	refTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: false}
	}
	tests := []struct {
		name   string
		tester zppTester
		want   types.Zval
		log    string
	}{
		// base-type
		{"type-undef", baseTester(types.Undef), types.Null, ""},
		{"type-null", baseTester(types.Null), types.Null, ""},
		{"type-false", baseTester(types.False), types.False, ""},
		{"type-true", baseTester(types.True), types.True, ""},
		{"type-long-0", baseTester(types.ZvalLong(0)), types.ZvalLong(0), ""},
		{"type-long-1", baseTester(types.ZvalLong(5)), types.ZvalLong(5), ""},
		{"type-double-0", baseTester(types.ZvalDouble(0)), types.ZvalDouble(0), ""},
		{"type-double-1", baseTester(types.ZvalDouble(5)), types.ZvalDouble(5), ""},
		{"type-string-0", baseTester(types.ZvalString("")), types.ZvalString(""), ""},
		{"type-string-1", baseTester(types.ZvalString("0")), types.ZvalString("0"), ""},
		{"type-string-2", baseTester(types.ZvalString("00")), types.ZvalString("00"), ""},
		{"type-array-0", baseTester(types.ZvalArray(types.NewArray())), types.ZvalArray(types.NewArray()), ""},
		{"type-array-1", baseTester(types.ZvalArray(types.NewArrayOf(types.Null))), types.ZvalArray(types.NewArrayOf(types.Null)), ""},
		// ref-type
		{"type-undef", refTester(types.Undef), types.Null, ""},
		{"type-null", refTester(types.Null), types.Null, ""},
		{"type-false", refTester(types.False), types.False, ""},
		{"type-true", refTester(types.True), types.True, ""},
		{"type-long-0", refTester(types.ZvalLong(0)), types.ZvalLong(0), ""},
		{"type-long-1", refTester(types.ZvalLong(5)), types.ZvalLong(5), ""},
		{"type-double-0", refTester(types.ZvalDouble(0)), types.ZvalDouble(0), ""},
		{"type-double-1", refTester(types.ZvalDouble(5)), types.ZvalDouble(5), ""},
		{"type-string-0", refTester(types.ZvalString("")), types.ZvalString(""), ""},
		{"type-string-1", refTester(types.ZvalString("0")), types.ZvalString("0"), ""},
		{"type-string-2", refTester(types.ZvalString("00")), types.ZvalString("00"), ""},
		{"type-array-0", baseTester(types.ZvalArray(types.NewArray())), types.ZvalArray(types.NewArray()), ""},
		{"type-array-1", baseTester(types.ZvalArray(types.NewArrayOf(types.Null))), types.ZvalArray(types.NewArrayOf(types.Null)), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseZval()
			tester.checkError(t, tt.log)
			if !testFastParamParser_deepEqualsZval(got, tt.want) {
				t.Errorf("ParseZval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastParamParser_ParseZvalNullable(t *testing.T) {
	refArgs := func(args ...types.Zval) []types.Zval {
		refArgs := make([]types.Zval, len(args))
		for i, arg := range args {
			refArgs[i] = types.ZvalRef(types.NewReference(arg))
		}
		return refArgs
	}
	baseTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: false}
	}
	refTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: false}
	}
	tests := []struct {
		name   string
		tester zppTester
		want   *types.Zval
		log    string
	}{
		// base-type
		{"type-undef", baseTester(types.Undef), nil, ""},
		{"type-null", baseTester(types.Null), nil, ""},
		{"type-false", baseTester(types.False), types.NewZvalFalse(), ""},
		{"type-true", baseTester(types.True), types.NewZvalTrue(), ""},
		{"type-long-0", baseTester(types.ZvalLong(0)), types.NewZvalLong(0), ""},
		{"type-long-1", baseTester(types.ZvalLong(5)), types.NewZvalLong(5), ""},
		{"type-double-0", baseTester(types.ZvalDouble(0)), types.NewZvalDouble(0), ""},
		{"type-double-1", baseTester(types.ZvalDouble(5)), types.NewZvalDouble(5), ""},
		{"type-string-0", baseTester(types.ZvalString("")), types.NewZvalString(""), ""},
		{"type-string-1", baseTester(types.ZvalString("0")), types.NewZvalString("0"), ""},
		{"type-string-2", baseTester(types.ZvalString("00")), types.NewZvalString("00"), ""},
		{"type-array-0", baseTester(types.ZvalArray(types.NewArray())), types.NewZvalArray(types.NewArray()), ""},
		{"type-array-1", baseTester(types.ZvalArray(types.NewArrayOf(types.Null))), types.NewZvalArray(types.NewArrayOf(types.Null)), ""},
		// ref-type
		{"type-undef", refTester(types.Undef), nil, ""},
		{"type-null", refTester(types.Null), nil, ""},
		{"type-false", refTester(types.False), types.NewZvalFalse(), ""},
		{"type-true", refTester(types.True), types.NewZvalTrue(), ""},
		{"type-long-0", refTester(types.ZvalLong(0)), types.NewZvalLong(0), ""},
		{"type-long-1", refTester(types.ZvalLong(5)), types.NewZvalLong(5), ""},
		{"type-double-0", refTester(types.ZvalDouble(0)), types.NewZvalDouble(0), ""},
		{"type-double-1", refTester(types.ZvalDouble(5)), types.NewZvalDouble(5), ""},
		{"type-string-0", refTester(types.ZvalString("")), types.NewZvalString(""), ""},
		{"type-string-1", refTester(types.ZvalString("0")), types.NewZvalString("0"), ""},
		{"type-string-2", refTester(types.ZvalString("00")), types.NewZvalString("00"), ""},
		{"type-array-0", baseTester(types.ZvalArray(types.NewArray())), types.NewZvalArray(types.NewArray()), ""},
		{"type-array-1", baseTester(types.ZvalArray(types.NewArrayOf(types.Null))), types.NewZvalArray(types.NewArrayOf(types.Null)), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			got := tester.parser().ParseZvalNullable()
			tester.checkError(t, tt.log)
			if !testFastParamParser_deepEqualsZvalPtr(got, tt.want) {
				t.Errorf("ParseZval() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseVariadic

func TestFastParamParser_ParseVariadic(t *testing.T) {
	refArgs := func(args ...types.Zval) []types.Zval {
		refArgs := make([]types.Zval, len(args))
		for i, arg := range args {
			refArgs[i] = types.ZvalRef(types.NewReference(arg))
		}
		return refArgs
	}
	baseTester := func(args ...types.Zval) zppTester {
		return zppTester{args: args, strict: false}
	}
	refTester := func(args ...types.Zval) zppTester {
		return zppTester{args: refArgs(args...), strict: false}
	}

	tests := []struct {
		name        string
		tester      zppTester
		postVarargs uint
		want        []types.Zval
		log         string
	}{
		{
			"base-0",
			baseTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			0,
			[]types.Zval{Long(2), Long(3), Long(4), Long(5)},
			"",
		},
		{
			"base-1",
			baseTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			1,
			[]types.Zval{Long(2), Long(3), Long(4)},
			"",
		},
		{
			"base-n",
			baseTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			100,
			[]types.Zval{},
			"",
		},
		{
			"ref-0",
			refTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			0,
			[]types.Zval{Long(2), Long(3), Long(4), Long(5)},
			"",
		},
		{
			"ref-1",
			refTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			1,
			[]types.Zval{Long(2), Long(3), Long(4)},
			"",
		},
		{
			"ref-n",
			refTester(Long(1), Long(2), Long(3), Long(4), Long(5)),
			100,
			[]types.Zval{},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := tt.tester
			tt.tester.parser().ParseZval()
			got := tester.parser().ParseVariadic(tt.postVarargs)
			tester.checkError(t, tt.log)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVariadic() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ParseRefZval

func testFastParamParser_checkRefZval(t *testing.T, got types.RefZval, rawVal types.Zval, wantNil bool) {
	mockVal := types.ZvalString("mock-val-123")
	if wantNil {
		if got != nil {
			t.Errorf("ParseRefZval() = %v, want nil", got)
		}
	} else {
		if got == nil {
			t.Errorf("ParseRefZval() = nil, want %v", rawVal)
			return
		}
		if !rawVal.IsRef() {
			t.Errorf("ParseRefZval() rawVal = %v, want reference val", rawVal)
			return
		}
		if gotVal := got.Val(); gotVal != rawVal.RefVal() {
			t.Errorf("ParseRefZval() gotval = %v, want %v", gotVal, rawVal.RefVal())
			return
		}

		got.SetVal(mockVal)
		if gotVal := got.Val(); gotVal != mockVal {
			t.Errorf("ParseRefZval() setval = %v, want %v", gotVal, mockVal)
			return
		}
		if derefRawVal := rawVal.RefVal(); derefRawVal != mockVal {
			t.Errorf("ParseRefZval() rawval after setval = %v, want %v", derefRawVal, mockVal)
			return
		}
	}
}

func TestFastParamParser_ParseRefZval(t *testing.T) {
	ref := func(v types.Zval) types.Zval { return types.ZvalRef(types.NewReference(v)) }

	tests := []struct {
		name    string
		val     types.Zval
		wantNil bool
		log     string
	}{
		// basic
		{"ref-1", ref(types.Undef), false, ""},
		{"ref-2", ref(types.Null), false, ""},
		{"ref-3", ref(types.ZvalLong(1)), false, ""},
		// noRef
		{"noRef-1", types.Undef, true, "E_WARNING: mockFunc() expects parameter 1 to be reference, null given"},
		{"noRef-2", types.Null, true, "E_WARNING: mockFunc() expects parameter 1 to be reference, null given"},
		{"noRef-3", types.ZvalLong(1), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, int given"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := &zppTester{args: []types.Zval{tt.val}}
			got := tester.parser().ParseRefZval()
			tester.checkError(t, tt.log)
			testFastParamParser_checkRefZval(t, got, tt.val, tt.wantNil)
		})
	}
}

func TestFastParamParser_ParseRefZvalNullable(t *testing.T) {
	ref := func(v types.Zval) types.Zval { return types.ZvalRef(types.NewReference(v)) }

	tests := []struct {
		name    string
		val     types.Zval
		wantNil bool
		log     string
	}{
		// basic
		{"ref-1", ref(types.Undef), false, ""},
		{"ref-2", ref(types.Null), false, ""},
		{"ref-3", ref(types.ZvalLong(1)), false, ""},
		// noRef
		{"noRef-1", types.Undef, true, ""},
		{"noRef-2", types.Null, true, ""},
		{"noRef-3", types.ZvalLong(1), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, int given"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := &zppTester{args: []types.Zval{tt.val}}
			got := tester.parser().ParseRefZvalNullable()
			tester.checkError(t, tt.log)
			testFastParamParser_checkRefZval(t, got, tt.val, tt.wantNil)
		})
	}
}

// ParseRefArray

func testFastParamParser_checkRefZvalArray(t *testing.T, got *types.Array, rawVal types.Zval, wantNil bool) {
	mockPairs := []types.ArrayPair{
		types.MakeArrayPair(types.IdxKey(0), Bool(true)),
		types.MakeArrayPair(types.IdxKey(1), Long(1234)),
		types.MakeArrayPair(types.IdxKey(2), String("abc")),
	}
	if wantNil {
		if got != nil {
			t.Errorf("ParseRefZval() = %v, want nil", got)
		}
	} else {
		if got == nil {
			t.Errorf("ParseRefZval() = nil, want %v", rawVal)
			return
		}
		if !rawVal.IsRef() || !rawVal.RefVal().IsArray() {
			t.Errorf("ParseRefZval() rawVal = %v, want reference array val", rawVal)
			return
		}
		rawArr := rawVal.RefVal().Array()

		if gotPairs, rawPairs := got.Pairs(), rawArr.Pairs(); !reflect.DeepEqual(gotPairs, rawPairs) {
			t.Errorf("ParseRefZval() gotPairs = %v, rawPairs = %v", gotPairs, rawPairs)
			return
		}

		got.SetDataByArray(types.NewArrayOfPairs(mockPairs))
		gotPairs, rawPairs := got.Pairs(), rawArr.Pairs()
		if !reflect.DeepEqual(gotPairs, mockPairs) {
			t.Errorf("ParseRefZval() gotPairs after setVal = %v, want = %v", gotPairs, mockPairs)
			return
		}
		if !reflect.DeepEqual(rawPairs, mockPairs) {
			t.Errorf("ParseRefZval() rawPairs after setVal = %v, want = %v", rawPairs, mockPairs)
			return
		}
	}
}

func TestFastParamParser_ParseRefArray(t *testing.T) {
	arr := func(values ...types.Zval) types.Zval { return types.ZvalArray(types.NewArrayOf(values...)) }
	refArr := func(values ...types.Zval) types.Zval {
		v := arr(values...)
		return types.ZvalRef(types.NewReference(v))
	}

	tests := []struct {
		name    string
		val     types.Zval
		wantNil bool
		log     string
	}{
		// basic
		{"ref-1", refArr(types.Undef), false, ""},
		{"ref-2", refArr(types.Null), false, ""},
		{"ref-3", refArr(types.ZvalLong(1)), false, ""},
		// noRef
		{"noRef-1", types.Undef, true, "E_WARNING: mockFunc() expects parameter 1 to be reference, null given"},
		{"noRef-2", types.Null, true, "E_WARNING: mockFunc() expects parameter 1 to be reference, null given"},
		{"noRef-3", types.ZvalLong(1), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, int given"},
		{"noRef-4", arr(types.Undef), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
		{"noRef-5", arr(types.Null), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
		{"noRef-6", arr(types.ZvalLong(1)), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := &zppTester{args: []types.Zval{tt.val}}
			got := tester.parser().ParseRefArray()
			tester.checkError(t, tt.log)
			testFastParamParser_checkRefZvalArray(t, got, tt.val, tt.wantNil)
		})
	}
}

func TestFastParamParser_ParseRefArrayNullable(t *testing.T) {
	arr := func(values ...types.Zval) types.Zval { return types.ZvalArray(types.NewArrayOf(values...)) }
	refArr := func(values ...types.Zval) types.Zval {
		v := arr(values...)
		return types.ZvalRef(types.NewReference(v))
	}

	tests := []struct {
		name    string
		val     types.Zval
		wantNil bool
		log     string
	}{
		// basic
		{"ref-1", refArr(types.Undef), false, ""},
		{"ref-2", refArr(types.Null), false, ""},
		{"ref-3", refArr(types.ZvalLong(1)), false, ""},
		// noRef
		{"noRef-1", types.Undef, true, ""},
		{"noRef-2", types.Null, true, ""},
		{"noRef-3", types.ZvalLong(1), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, int given"},
		{"noRef-4", arr(types.Undef), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
		{"noRef-5", arr(types.Null), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
		{"noRef-6", arr(types.ZvalLong(1)), true, "E_WARNING: mockFunc() expects parameter 1 to be reference, array given"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := &zppTester{args: []types.Zval{tt.val}}
			got := tester.parser().ParseRefArrayNullable()
			tester.checkError(t, tt.log)
			testFastParamParser_checkRefZvalArray(t, got, tt.val, tt.wantNil)
		})
	}
}
