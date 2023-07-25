package ir

import "reflect"

func nullsafe[T any, R any](arg *T, handler func(*T) *R) *R {
	if arg == nil {
		return nil
	}
	return handler(arg)
}

func nullsafeOrDefault[T any, R any](arg *T, handler func(*T) R, defaultValue R) R {
	if arg == nil {
		return defaultValue
	}
	return handler(arg)
}

func isNil(n any) bool {
	if n == nil {
		return true
	}

	v := reflect.ValueOf(n)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}
