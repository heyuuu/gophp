package standard

func PHP_VAR_UNSERIALIZE_INIT(d PhpUnserializeDataT) PhpUnserializeDataT {
	d = PhpVarUnserializeInit()
	return d
}
func PHP_VAR_UNSERIALIZE_DESTROY(d PhpUnserializeDataT) { PhpVarUnserializeDestroy(d) }
