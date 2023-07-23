package ir

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
