// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/ext/standard"
	"sik/zend"
)

func SplObjectStorageFromObj(obj *zend.ZendObject) *spl_SplObjectStorage {
	return (*spl_SplObjectStorage)((*byte)(obj - zend_long((*byte)(&((*spl_SplObjectStorage)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLOBJSTORAGE_P(zv *zend.Zval) *spl_SplObjectStorage {
	return SplObjectStorageFromObj(zv.GetObj())
}
func spl_SplObjectStorage_free_storage(object *zend.ZendObject) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(object)
	zend.ZendObjectStdDtor(intern.GetStd())
	intern.GetStorage().Destroy()
	if intern.GetGcdata() != nil {
		zend.Efree(intern.GetGcdata())
	}
}
func SplObjectStorageGetHash(key *zend.ZendHashKey, intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	if intern.GetFptrGetHash() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith1Params(this, intern.GetStd().GetCe(), intern.GetFptrGetHash(), "getHash", &rv, obj)
		if !(rv.IsUndef()) {
			if rv.IsType(zend.IS_STRING) {
				key.SetKey(rv.GetStr())
				return zend.SUCCESS
			} else {
				zend.ZendThrowException(spl_ce_RuntimeException, "Hash needs to be a string", 0)
				zend.ZvalPtrDtor(&rv)
				return zend.FAILURE
			}
		} else {
			return zend.FAILURE
		}
	} else {
		key.SetKey(nil)
		key.SetH(zend.Z_OBJ_HANDLE_P(obj))
		return zend.SUCCESS
	}
}
func SplObjectStorageFreeHash(intern *spl_SplObjectStorage, key *zend.ZendHashKey) {
	if key.GetKey() != nil {
		zend.ZendStringReleaseEx(key.GetKey(), 0)
	}
}
func SplObjectStorageDtor(element *zend.Zval) {
	var el *spl_SplObjectStorageElement = element.GetPtr()
	zend.ZvalPtrDtor(el.GetObj())
	zend.ZvalPtrDtor(el.GetInf())
	zend.Efree(el)
}
func SplObjectStorageGet(intern *spl_SplObjectStorage, key *zend.ZendHashKey) *spl_SplObjectStorageElement {
	if key.GetKey() != nil {
		return zend.ZendHashFindPtr(intern.GetStorage(), key.GetKey())
	} else {
		return zend.ZendHashIndexFindPtr(intern.GetStorage(), key.GetH())
	}
}
func SplObjectStorageAttach(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval, inf *zend.Zval) *spl_SplObjectStorageElement {
	var pelement *spl_SplObjectStorageElement
	var element spl_SplObjectStorageElement
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return nil
	}
	pelement = SplObjectStorageGet(intern, &key)
	if pelement != nil {
		zend.ZvalPtrDtor(pelement.GetInf())
		if inf != nil {
			zend.ZVAL_COPY(pelement.GetInf(), inf)
		} else {
			pelement.GetInf().SetNull()
		}
		SplObjectStorageFreeHash(intern, &key)
		return pelement
	}
	zend.ZVAL_COPY(element.GetObj(), obj)
	if inf != nil {
		zend.ZVAL_COPY(element.GetInf(), inf)
	} else {
		element.GetInf().SetNull()
	}
	if key.GetKey() != nil {
		pelement = zend.ZendHashUpdateMem(intern.GetStorage(), key.GetKey(), &element, b.SizeOf("spl_SplObjectStorageElement"))
	} else {
		pelement = zend.ZendHashIndexUpdateMem(intern.GetStorage(), key.GetH(), &element, b.SizeOf("spl_SplObjectStorageElement"))
	}
	SplObjectStorageFreeHash(intern, &key)
	return pelement
}
func SplObjectStorageDetach(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	var ret int = zend.FAILURE
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return ret
	}
	if key.GetKey() != nil {
		ret = zend.ZendHashDel(intern.GetStorage(), key.GetKey())
	} else {
		ret = zend.ZendHashIndexDel(intern.GetStorage(), key.GetH())
	}
	SplObjectStorageFreeHash(intern, &key)
	return ret
}
func SplObjectStorageAddall(intern *spl_SplObjectStorage, this *zend.Zval, other *spl_SplObjectStorage) {
	var element *spl_SplObjectStorageElement
	var __ht *zend.HashTable = other.GetStorage()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		element = _z.GetPtr()
		SplObjectStorageAttach(intern, this, element.GetObj(), element.GetInf())
	}
	intern.SetIndex(0)
}
func SplObjectStorageNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval) *zend.ZendObject {
	var intern *spl_SplObjectStorage
	var parent *zend.ZendClassEntry = class_type
	intern = zend.Emalloc(b.SizeOf("spl_SplObjectStorage") + zend.ZendObjectPropertiesSize(parent))
	memset(intern, 0, b.SizeOf("spl_SplObjectStorage")-b.SizeOf("zval"))
	intern.SetPos(0)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	zend.ZendHashInit(intern.GetStorage(), 0, nil, SplObjectStorageDtor, 0)
	intern.GetStd().SetHandlers(&spl_handler_SplObjectStorage)
	for parent != nil {
		if parent == spl_ce_SplObjectStorage {
			if class_type != spl_ce_SplObjectStorage {
				intern.SetFptrGetHash(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "gethash", b.SizeOf("\"gethash\"")-1))
				if intern.GetFptrGetHash().GetScope() == spl_ce_SplObjectStorage {
					intern.SetFptrGetHash(nil)
				}
			}
			break
		}
		parent = parent.GetParent()
	}
	if orig != nil {
		var other *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(orig)
		SplObjectStorageAddall(intern, orig, other)
	}
	return intern.GetStd()
}
func SplObjectStorageClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.GetObj()
	new_object = SplObjectStorageNewEx(old_object.GetCe(), zobject)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplObjectStorageDebugInfo(obj *zend.Zval) *zend.HashTable {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(obj)
	var element *spl_SplObjectStorageElement
	var props *zend.HashTable
	var tmp zend.Zval
	var storage zend.Zval
	var md5str *zend.ZendString
	var zname *zend.ZendString
	var debug_info *zend.HashTable
	props = zend.Z_OBJPROP_P(obj)
	debug_info = zend.ZendNewArray(props.GetNNumOfElements() + 1)
	zend.ZendHashCopy(debug_info, props, zend.CopyCtorFuncT(zend.ZvalAddRef))
	zend.ArrayInit(&storage)
	var __ht *zend.HashTable = intern.GetStorage()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		element = _z.GetPtr()
		md5str = PhpSplObjectHash(element.GetObj())
		zend.ArrayInit(&tmp)

		/* Incrementing the refcount of obj and inf would confuse the garbage collector.
		 * Prefer to null the destructor */

		zend.Z_ARRVAL_P(&tmp).SetPDestructor(nil)
		zend.AddAssocZvalEx(&tmp, "obj", b.SizeOf("\"obj\"")-1, element.GetObj())
		zend.AddAssocZvalEx(&tmp, "inf", b.SizeOf("\"inf\"")-1, element.GetInf())
		storage.GetArr().KeyUpdate(md5str.GetStr(), &tmp)
		zend.ZendStringReleaseEx(md5str, 0)
	}
	zname = SplGenPrivatePropName(spl_ce_SplObjectStorage, "storage")
	debug_info.SymtableUpdate(zname.GetStr(), &storage)
	zend.ZendStringReleaseEx(zname, 0)
	return debug_info
}
func SplObjectStorageGetGc(obj *zend.Zval, table **zend.Zval, n *int) *zend.HashTable {
	var i int = 0
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(obj)
	var element *spl_SplObjectStorageElement
	if intern.GetStorage().GetNNumOfElements()*2 > intern.GetGcdataNum() {
		intern.SetGcdataNum(intern.GetStorage().GetNNumOfElements() * 2)
		intern.SetGcdata((*zend.Zval)(zend.Erealloc(intern.GetGcdata(), b.SizeOf("zval")*intern.GetGcdataNum())))
	}
	var __ht *zend.HashTable = intern.GetStorage()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		element = _z.GetPtr()
		zend.ZVAL_COPY_VALUE(intern.GetGcdata()[b.PostInc(&i)], element.GetObj())
		zend.ZVAL_COPY_VALUE(intern.GetGcdata()[b.PostInc(&i)], element.GetInf())
	}
	*table = intern.GetGcdata()
	*n = i
	return zend.ZendStdGetProperties(obj)
}
func SplObjectStorageCompareInfo(e1 *zend.Zval, e2 *zend.Zval) int {
	var s1 *spl_SplObjectStorageElement = (*spl_SplObjectStorageElement)(e1.GetPtr())
	var s2 *spl_SplObjectStorageElement = (*spl_SplObjectStorageElement)(e2.GetPtr())
	var result zend.Zval
	if zend.CompareFunction(&result, s1.GetInf(), s2.GetInf()) == zend.FAILURE {
		return 1
	}
	return zend.ZEND_NORMALIZE_BOOL(result.GetLval())
}
func SplObjectStorageCompareObjects(o1 *zend.Zval, o2 *zend.Zval) int {
	var zo1 *zend.ZendObject = (*zend.ZendObject)(o1.GetObj())
	var zo2 *zend.ZendObject = (*zend.ZendObject)(o2.GetObj())
	if zo1.GetCe() != spl_ce_SplObjectStorage || zo2.GetCe() != spl_ce_SplObjectStorage {
		return 1
	}
	return zend.ZendHashCompare(Z_SPLOBJSTORAGE_P(o1).GetStorage(), Z_SPLOBJSTORAGE_P(o2).GetStorage(), zend.CompareFuncT(SplObjectStorageCompareInfo), 0)
}
func spl_SplObjectStorage_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplObjectStorageNewEx(class_type, nil)
}
func SplObjectStorageContains(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	var found int
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return 0
	}
	if key.GetKey() != nil {
		found = zend.ZendHashExists(intern.GetStorage(), key.GetKey())
	} else {
		found = zend.ZendHashIndexExists(intern.GetStorage(), key.GetH())
	}
	SplObjectStorageFreeHash(intern, &key)
	return found
}
func zim_spl_SplObjectStorage_attach(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var inf *zend.Zval = nil
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o|z!", &obj, &inf) == zend.FAILURE {
		return
	}
	SplObjectStorageAttach(intern, zend.ZEND_THIS, obj, inf)
}
func zim_spl_SplObjectStorage_detach(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o", &obj) == zend.FAILURE {
		return
	}
	SplObjectStorageDetach(intern, zend.ZEND_THIS, obj)
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
}
func zim_spl_SplObjectStorage_getHash(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o", &obj) == zend.FAILURE {
		return
	}
	return_value.SetString(PhpSplObjectHash(obj))
	return
}
func zim_spl_SplObjectStorage_offsetGet(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var key zend.ZendHashKey
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o", &obj) == zend.FAILURE {
		return
	}
	if SplObjectStorageGetHash(&key, intern, zend.ZEND_THIS, obj) == zend.FAILURE {
		return
	}
	element = SplObjectStorageGet(intern, &key)
	SplObjectStorageFreeHash(intern, &key)
	if element == nil {
		zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Object not found")
	} else {
		var value *zend.Zval = element.GetInf()
		zend.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_spl_SplObjectStorage_addAll(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var other *spl_SplObjectStorage
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)
	SplObjectStorageAddall(intern, zend.ZEND_THIS, other)
	return_value.SetLong(intern.GetStorage().GetNNumOfElements())
	return
}
func zim_spl_SplObjectStorage_removeAll(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var other *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)
	zend.ZendHashInternalPointerReset(other.GetStorage())
	for b.Assign(&element, zend.ZendHashGetCurrentDataPtr(other.GetStorage())) != nil {
		if SplObjectStorageDetach(intern, zend.ZEND_THIS, element.GetObj()) == zend.FAILURE {
			zend.ZendHashMoveForward(other.GetStorage())
		}
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
	return_value.SetLong(intern.GetStorage().GetNNumOfElements())
	return
}
func zim_spl_SplObjectStorage_removeAllExcept(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var other *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)
	var __ht *zend.HashTable = intern.GetStorage()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		element = _z.GetPtr()
		if SplObjectStorageContains(other, zend.ZEND_THIS, element.GetObj()) == 0 {
			SplObjectStorageDetach(intern, zend.ZEND_THIS, element.GetObj())
		}
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
	return_value.SetLong(intern.GetStorage().GetNNumOfElements())
	return
}
func zim_spl_SplObjectStorage_contains(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o", &obj) == zend.FAILURE {
		return
	}
	zend.ZVAL_BOOL(return_value, SplObjectStorageContains(intern, zend.ZEND_THIS, obj) != 0)
	return
}
func zim_spl_SplObjectStorage_count(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var mode zend.ZendLong = standard.COUNT_NORMAL
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|l", &mode) == zend.FAILURE {
		return
	}
	if mode == standard.COUNT_RECURSIVE {
		var ret zend.ZendLong
		if mode != standard.COUNT_RECURSIVE {
			ret = intern.GetStorage().GetNNumOfElements()
		} else {
			ret = standard.PhpCountRecursive(intern.GetStorage())
		}
		return_value.SetLong(ret)
		return
		return
	}
	return_value.SetLong(intern.GetStorage().GetNNumOfElements())
	return
}
func zim_spl_SplObjectStorage_rewind(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
}
func zim_spl_SplObjectStorage_valid(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZVAL_BOOL(return_value, zend.ZendHashHasMoreElementsEx(intern.GetStorage(), intern.GetPos()) == zend.SUCCESS)
	return
}
func zim_spl_SplObjectStorage_key(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetLong(intern.GetIndex())
	return
}
func zim_spl_SplObjectStorage_current(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	zend.ZVAL_COPY(return_value, element.GetObj())
}
func zim_spl_SplObjectStorage_getInfo(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	zend.ZVAL_COPY(return_value, element.GetInf())
}
func zim_spl_SplObjectStorage_setInfo(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var inf *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &inf) == zend.FAILURE {
		return
	}
	if b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	zend.ZvalPtrDtor(element.GetInf())
	zend.ZVAL_COPY(element.GetInf(), inf)
}
func zim_spl_SplObjectStorage_next(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	intern.GetIndex()++
}
func zim_spl_SplObjectStorage_serialize(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var element *spl_SplObjectStorageElement
	var members zend.Zval
	var flags zend.Zval
	var pos zend.HashPosition
	var var_hash standard.PhpSerializeDataT
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	standard.PHP_VAR_SERIALIZE_INIT(var_hash)

	/* storage */

	buf.AppendString("x:")
	flags.SetLong(intern.GetStorage().GetNNumOfElements())
	standard.PhpVarSerialize(&buf, &flags, &var_hash)
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), &pos)
	for zend.ZendHashHasMoreElementsEx(intern.GetStorage(), &pos) == zend.SUCCESS {
		if b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), &pos)) == nil {
			buf.Free()
			standard.PHP_VAR_SERIALIZE_DESTROY(var_hash)
			return_value.SetNull()
			return
		}
		standard.PhpVarSerialize(&buf, element.GetObj(), &var_hash)
		buf.AppendByte(',')
		standard.PhpVarSerialize(&buf, element.GetInf(), &var_hash)
		buf.AppendByte(';')
		zend.ZendHashMoveForwardEx(intern.GetStorage(), &pos)
	}

	/* members */

	buf.AppendString("m:")
	members.SetArray(zend.ZendArrayDup(zend.ZendStdGetProperties(zend.ZEND_THIS)))
	standard.PhpVarSerialize(&buf, &members, &var_hash)
	zend.ZvalPtrDtor(&members)

	/* done */

	standard.PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if buf.GetS() != nil {
		return_value.SetString(buf.GetS())
		return
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_SplObjectStorage_unserialize(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var entry zend.Zval
	var inf zend.Zval
	var pcount *zend.Zval
	var pmembers *zend.Zval
	var element *spl_SplObjectStorageElement
	var count zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}

	/* storage */

	p = (*uint8)(buf)
	s = p
	standard.PHP_VAR_UNSERIALIZE_INIT(var_hash)
	if (*p) != 'x' || (*(b.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	pcount = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(pcount, &p, s+buf_len, &var_hash) == 0 || pcount.GetType() != zend.IS_LONG {
		goto outexcept
	}
	p--
	count = pcount.GetLval()
	if count < 0 {
		goto outexcept
	}
	entry.SetUndef()
	inf.SetUndef()
	for b.PostDec(&count) > 0 {
		var pelement *spl_SplObjectStorageElement
		var key zend.ZendHashKey
		if (*p) != ';' {
			goto outexcept
		}
		p++
		if (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}

		/* store reference to allow cross-references between different elements */

		if standard.PhpVarUnserialize(&entry, &p, s+buf_len, &var_hash) == 0 {
			zend.ZvalPtrDtor(&entry)
			goto outexcept
		}
		if (*p) == ',' {
			p++
			if standard.PhpVarUnserialize(&inf, &p, s+buf_len, &var_hash) == 0 {
				zend.ZvalPtrDtor(&entry)
				zend.ZvalPtrDtor(&inf)
				goto outexcept
			}
		}
		if entry.GetType() != zend.IS_OBJECT {
			zend.ZvalPtrDtor(&entry)
			zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		if SplObjectStorageGetHash(&key, intern, zend.ZEND_THIS, &entry) == zend.FAILURE {
			zend.ZvalPtrDtor(&entry)
			zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		pelement = SplObjectStorageGet(intern, &key)
		SplObjectStorageFreeHash(intern, &key)
		if pelement != nil {
			if !(pelement.GetInf().IsUndef()) {
				standard.VarPushDtor(&var_hash, pelement.GetInf())
			}
			if !(pelement.GetObj().IsUndef()) {
				standard.VarPushDtor(&var_hash, pelement.GetObj())
			}
		}
		element = SplObjectStorageAttach(intern, zend.ZEND_THIS, &entry, b.Cond(inf.IsUndef(), nil, &inf))
		standard.VarReplace(&var_hash, &entry, element.GetObj())
		standard.VarReplace(&var_hash, &inf, element.GetInf())
		zend.ZvalPtrDtor(&entry)
		entry.SetUndef()
		zend.ZvalPtrDtor(&inf)
		inf.SetUndef()
	}
	if (*p) != ';' {
		goto outexcept
	}
	p++

	/* members */

	if (*p) != 'm' || (*(b.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	pmembers = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(pmembers, &p, s+buf_len, &var_hash) == 0 || pmembers.GetType() != zend.IS_ARRAY {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(intern.GetStd(), pmembers.GetArr())
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	return
outexcept:
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset %zd of %zd bytes", (*byte)(p-buf), buf_len)
	return
}
func zim_spl_SplObjectStorage___serialize(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var elem *spl_SplObjectStorageElement
	var tmp zend.Zval
	if zend.ZendParseParametersNoneThrow() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)

	/* storage */

	zend.ArrayInitSize(&tmp, 2*intern.GetStorage().GetNNumOfElements())
	var __ht *zend.HashTable = intern.GetStorage()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		elem = _z.GetPtr()
		elem.GetObj().TryAddRefcount()
		tmp.GetArr().NextIndexInsert(elem.GetObj())
		elem.GetInf().TryAddRefcount()
		tmp.GetArr().NextIndexInsert(elem.GetInf())
	}
	return_value.GetArr().NextIndexInsert(&tmp)

	/* members */

	tmp.SetArray(zend.ZendStdGetProperties(zend.ZEND_THIS))
	tmp.TryAddRefcount()
	return_value.GetArr().NextIndexInsert(&tmp)
}
func zim_spl_SplObjectStorage___unserialize(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	var data *zend.HashTable
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var key *zend.Zval
	var val *zend.Zval
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "h", &data) == zend.FAILURE {
		return
	}
	storage_zv = data.IndexFindH(0)
	members_zv = data.IndexFindH(1)
	if storage_zv == nil || members_zv == nil || storage_zv.GetType() != zend.IS_ARRAY || members_zv.GetType() != zend.IS_ARRAY {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	if zend.Z_ARRVAL_P(storage_zv).GetNNumOfElements()%2 != 0 {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Odd number of elements", 0)
		return
	}
	key = nil
	var __ht *zend.HashTable = storage_zv.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		val = _z
		if key != nil {
			if key.GetType() != zend.IS_OBJECT {
				zend.ZendThrowException(spl_ce_UnexpectedValueException, "Non-object key", 0)
				return
			}
			SplObjectStorageAttach(intern, zend.ZEND_THIS, key, val)
			key = nil
		} else {
			key = val
		}
	}
	zend.ObjectPropertiesLoad(intern.GetStd(), members_zv.GetArr())
}
func zim_spl_SplObjectStorage___debugInfo(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetArray(SplObjectStorageDebugInfo(zend.getThis()))
	return
}
func zim_spl_MultipleIterator___construct(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var flags zend.ZendLong = MIT_NEED_ALL | MIT_KEYS_NUMERIC
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "|l", &flags) == zend.FAILURE {
		return
	}
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	intern.SetFlags(flags)
}
func zim_spl_MultipleIterator_getFlags(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_MultipleIterator_setFlags(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", intern.GetFlags()) == zend.FAILURE {
		return
	}
}
func zim_spl_MultipleIterator_attachIterator(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var iterator *zend.Zval = nil
	var info *zend.Zval = nil
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O|z!", &iterator, zend.ZendCeIterator, &info) == zend.FAILURE {
		return
	}
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if info != nil {
		var element *spl_SplObjectStorageElement
		if info.GetType() != zend.IS_LONG && info.GetType() != zend.IS_STRING {
			zend.ZendThrowException(spl_ce_InvalidArgumentException, "Info must be NULL, integer or string", 0)
			return
		}
		zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
		for b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil {
			if zend.FastIsIdenticalFunction(info, element.GetInf()) != 0 {
				zend.ZendThrowException(spl_ce_InvalidArgumentException, "Key duplication error", 0)
				return
			}
			zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
		}
	}
	SplObjectStorageAttach(intern, zend.ZEND_THIS, iterator, info)
}
func zim_spl_MultipleIterator_rewind(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfRewind(), "rewind", nil)
		zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_next(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfNext(), "next", nil)
		zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_valid(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	var retval zend.Zval
	var expect zend.ZendLong
	var valid zend.ZendLong
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(intern.GetStorage().GetNNumOfElements()) {
		return_value.SetFalse()
		return
	}
	if intern.IsNeedAll() {
		expect = 1
	} else {
		expect = 0
	}
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfValid(), "valid", &retval)
		if !(retval.IsUndef()) {
			valid = retval.IsType(zend.IS_TRUE)
			zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if expect != valid {
			zend.ZVAL_BOOL(return_value, expect == 0)
			return
		}
		zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
	zend.ZVAL_BOOL(return_value, expect != 0)
	return
}
func SplMultipleIteratorGetAll(intern *spl_SplObjectStorage, get_type int, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	var retval zend.Zval
	var valid int = 1
	var num_elements int
	num_elements = intern.GetStorage().GetNNumOfElements()
	if num_elements < 1 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInitSize(return_value, num_elements)
	zend.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfValid(), "valid", &retval)
		if !(retval.IsUndef()) {
			valid = retval.IsType(zend.IS_TRUE)
			zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if valid != 0 {
			if SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT == get_type {
				zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfCurrent(), "current", &retval)
			} else {
				zend.ZendCallMethodWith0Params(it, zend.Z_OBJCE_P(it), zend.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfKey(), "key", &retval)
			}
			if retval.IsUndef() {
				zend.ZendThrowException(spl_ce_RuntimeException, "Failed to call sub iterator method", 0)
				return
			}
		} else if intern.IsNeedAll() {
			if SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT == get_type {
				zend.ZendThrowException(spl_ce_RuntimeException, "Called current() with non valid sub iterator", 0)
			} else {
				zend.ZendThrowException(spl_ce_RuntimeException, "Called key() with non valid sub iterator", 0)
			}
			return
		} else {
			retval.SetNull()
		}
		if intern.IsKeysAssoc() {
			switch element.GetInf().GetType() {
			case zend.IS_LONG:
				zend.AddIndexZval(return_value, element.GetInf().GetLval(), &retval)
			case zend.IS_STRING:
				return_value.GetArr().SymtableUpdate(element.GetInf().GetStr().GetStr(), &retval)
			default:
				zend.ZvalPtrDtor(&retval)
				zend.ZendThrowException(spl_ce_InvalidArgumentException, "Sub-Iterator is associated with NULL", 0)
				return
			}
		} else {
			zend.AddNextIndexZval(return_value, &retval)
		}
		zend.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_current(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplMultipleIteratorGetAll(intern, SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT, return_value)
}
func zim_spl_MultipleIterator_key(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplMultipleIteratorGetAll(intern, SPL_MULTIPLE_ITERATOR_GET_ALL_KEY, return_value)
}
func ZmStartupSplObserver(type_ int, module_number int) int {
	SplRegisterInterface(&spl_ce_SplObserver, "SplObserver", spl_funcs_SplObserver)
	SplRegisterInterface(&spl_ce_SplSubject, "SplSubject", spl_funcs_SplSubject)
	SplRegisterStdClass(&spl_ce_SplObjectStorage, "SplObjectStorage", spl_SplObjectStorage_new, spl_funcs_SplObjectStorage)
	memcpy(&spl_handler_SplObjectStorage, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplObjectStorage.SetOffset(zend_long((*byte)(&((*spl_SplObjectStorage)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_SplObjectStorage.SetCompareObjects(SplObjectStorageCompareObjects)
	spl_handler_SplObjectStorage.SetCloneObj(SplObjectStorageClone)
	spl_handler_SplObjectStorage.SetGetGc(SplObjectStorageGetGc)
	spl_handler_SplObjectStorage.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_SplObjectStorage.SetFreeObj(spl_SplObjectStorage_free_storage)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Countable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_ArrayAccess)
	SplRegisterStdClass(&spl_ce_MultipleIterator, "MultipleIterator", spl_SplObjectStorage_new, spl_funcs_MultipleIterator)
	zend.ZendClassImplements(spl_ce_MultipleIterator, 1, zend.ZendCeIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ANY", b.SizeOf("\"MIT_NEED_ANY\"")-1, zend.ZendLong(MIT_NEED_ANY))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ALL", b.SizeOf("\"MIT_NEED_ALL\"")-1, zend.ZendLong(MIT_NEED_ALL))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_NUMERIC", b.SizeOf("\"MIT_KEYS_NUMERIC\"")-1, zend.ZendLong(MIT_KEYS_NUMERIC))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_ASSOC", b.SizeOf("\"MIT_KEYS_ASSOC\"")-1, zend.ZendLong(MIT_KEYS_ASSOC))
	return zend.SUCCESS
}
