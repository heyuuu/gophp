// <<generate>>

package standard

func PHP_VAR_SERIALIZE_INIT(d PhpSerializeDataT) PhpSerializeDataT {
	d = PhpVarSerializeInit()
	return d
}
func PHP_VAR_SERIALIZE_DESTROY(d PhpSerializeDataT) { PhpVarSerializeDestroy(d) }
func PHP_VAR_UNSERIALIZE_INIT(d PhpUnserializeDataT) PhpUnserializeDataT {
	d = PhpVarUnserializeInit()
	return d
}
func PHP_VAR_UNSERIALIZE_DESTROY(d PhpUnserializeDataT) { PhpVarUnserializeDestroy(d) }
