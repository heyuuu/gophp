package ir

func nullsafe[T any, R any](arg *T, handler func(*T) *R) *R {
	if arg == nil {
		return nil
	}
	return handler(arg)
}
