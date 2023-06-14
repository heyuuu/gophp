package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/ext/standard/array"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func SplObjectStorageFromObj(obj *types.ZendObject) *SplObjectStorage {
	return (*SplObjectStorage)((*byte)(obj - zend_long((*byte)(&((*SplObjectStorage)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLOBJSTORAGE_P(zv *types.Zval) *SplObjectStorage {
	return SplObjectStorageFromObj(zv.Object())
}
func spl_SplObjectStorage_free_storage(object *types.ZendObject) {
	var intern *SplObjectStorage = SplObjectStorageFromObj(object)
	zend.ZendObjectStdDtor(intern.GetStd())
	intern.GetStorage().Destroy()
	if intern.GetGcdata() != nil {
		zend.Efree(intern.GetGcdata())
	}
}
func SplObjectStorageGetHash(key *types.ArrayKey, intern *SplObjectStorage, this *types.Zval, obj *types.Zval) int {
	if intern.GetFptrGetHash() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith1Params(this, intern.GetStd().GetCe(), intern.GetFptrGetHash(), "getHash", &rv, obj)
		if !(rv.IsUndef()) {
			if rv.IsString() {
				*key = types.StrKey(rv.StringVal())
				return types.SUCCESS
			} else {
				faults.ThrowException(spl_ce_RuntimeException, "Hash needs to be a string", 0)
				return types.FAILURE
			}
		} else {
			return types.FAILURE
		}
	} else {
		*key = types.IdxKey(int(obj.Object().GetHandle()))
		return types.SUCCESS
	}
}
func SplObjectStorageGet(intern *SplObjectStorage, key *types.ArrayKey) *SplObjectStorageElement {
	var ptr any
	if key.IsStrKey() {
		ptr = types.ZendHashFindPtr(intern.GetStorage(), key.StrKey())
	} else {
		ptr = types.ZendHashIndexFindPtr(intern.GetStorage(), key.IdxKey())
	}
	if ptr != nil {
		return ptr.(*SplObjectStorageElement)
	}
	return nil
}
func SplObjectStorageAttach(intern *SplObjectStorage, this *types.Zval, obj *types.Zval, inf *types.Zval) *SplObjectStorageElement {
	var pelement *SplObjectStorageElement
	var element SplObjectStorageElement
	var key types.ArrayKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == types.FAILURE {
		return nil
	}
	pelement = SplObjectStorageGet(intern, &key)
	if pelement != nil {
		// zend.ZvalPtrDtor(pelement.GetInf())
		if inf != nil {
			types.ZVAL_COPY(pelement.GetInf(), inf)
		} else {
			pelement.GetInf().SetNull()
		}
		return pelement
	}
	types.ZVAL_COPY(element.GetObj(), obj)
	if inf != nil {
		types.ZVAL_COPY(element.GetInf(), inf)
	} else {
		element.GetInf().SetNull()
	}
	if key.IsStrKey() {
		pelement = types.ZendHashUpdateMem(intern.GetStorage(), key.StrKey(), &element, b.SizeOf("spl_SplObjectStorageElement"))
	} else {
		pelement = types.ZendHashIndexUpdateMem(intern.GetStorage(), key.IdxKey(), &element, b.SizeOf("spl_SplObjectStorageElement"))
	}
	return pelement
}
func SplObjectStorageDetach(intern *SplObjectStorage, this *types.Zval, obj *types.Zval) int {
	var key types.ArrayKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == types.FAILURE {
		return types.IntBool(false)
	}
	ret := intern.GetStorage().Delete(key)
	return types.IntBool(ret)
}
func SplObjectStorageAddall(intern *SplObjectStorage, this *types.Zval, other *SplObjectStorage) {
	other.GetStorage().Foreach(func(key types.ArrayKey, value *types.Zval) {
		var element *SplObjectStorageElement = value.Ptr()
		SplObjectStorageAttach(intern, this, element.GetObj(), element.GetInf())
	})

	intern.SetIndex(0)
}
func SplObjectStorageNewEx(class_type *types.ClassEntry, orig *types.Zval) *types.ZendObject {
	var intern *SplObjectStorage = NewSplObjectStorage(class_type)
	var parent *types.ClassEntry = class_type
	for parent != nil {
		if parent == spl_ce_SplObjectStorage {
			if class_type != spl_ce_SplObjectStorage {
				intern.SetFptrGetHash(class_type.FunctionTable().Get("gethash"))
				if intern.GetFptrGetHash().GetScope() == spl_ce_SplObjectStorage {
					intern.SetFptrGetHash(nil)
				}
			}
			break
		}
		parent = parent.GetParent()
	}
	if orig != nil {
		var other *SplObjectStorage = Z_SPLOBJSTORAGE_P(orig)
		SplObjectStorageAddall(intern, orig, other)
	}
	return intern.GetStd()
}
func SplObjectStorageClone(zobject *types.Zval) *types.ZendObject {
	var old_object *types.ZendObject
	var new_object *types.ZendObject
	old_object = zobject.Object()
	new_object = SplObjectStorageNewEx(old_object.GetCe(), zobject)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplObjectStorageDebugInfo(obj *types.Zval) *types.Array {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(obj)
	var element *SplObjectStorageElement
	var props *types.Array
	var tmp types.Zval
	var storage types.Zval
	var zname *types.String
	var debug_info *types.Array
	props = types.Z_OBJPROP_P(obj)
	debug_info = types.NewArray(props.Len() + 1)
	types.ZendHashCopy(debug_info, props)
	zend.ArrayInit(&storage)

	intern.GetStorage().Foreach(func(key types.ArrayKey, value *types.Zval) {
		element = value.Ptr()
		md5str := PhpSplObjectHash(element.GetObj())
		zend.ArrayInit(&tmp)

		/* Incrementing the refcount of obj and inf would confuse the garbage collector.
		 * Prefer to null the destructor */
		zend.AddAssocZvalEx(&tmp, "obj", element.GetObj())
		zend.AddAssocZvalEx(&tmp, "inf", element.GetInf())
		storage.Array().KeyUpdate(md5str, &tmp)
	})

	zname = SplGenPrivatePropName(spl_ce_SplObjectStorage, "storage")
	debug_info.SymtableUpdate(zname.GetStr(), &storage)
	// types.ZendStringReleaseEx(zname, 0)
	return debug_info
}
func SplObjectStorageGetGc(obj *types.Zval, table **types.Zval, n *int) *types.Array {
	var i int = 0
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(obj)
	var element *SplObjectStorageElement
	if intern.GetStorage().Len()*2 > intern.GetGcdataNum() {
		intern.SetGcdataNum(intern.GetStorage().Len() * 2)
		intern.SetGcdata((*types.Zval)(zend.Erealloc(intern.GetGcdata(), b.SizeOf("zval")*intern.GetGcdataNum())))
	}
	intern.GetStorage().Foreach(func(key types.ArrayKey, value *types.Zval) {
		element = value.Ptr()
		types.ZVAL_COPY_VALUE(intern.GetGcdata()[b.PostInc(&i)], element.GetObj())
		types.ZVAL_COPY_VALUE(intern.GetGcdata()[b.PostInc(&i)], element.GetInf())
	})
	*table = intern.GetGcdata()
	*n = i
	return zend.ZendStdGetProperties(obj)
}
func SplObjectStorageCompareInfo(e1 *types.Zval, e2 *types.Zval) int {
	var s1 *SplObjectStorageElement = (*SplObjectStorageElement)(e1.Ptr())
	var s2 *SplObjectStorageElement = (*SplObjectStorageElement)(e2.Ptr())
	var result types.Zval
	if operators.CompareFunction(&result, s1.GetInf(), s2.GetInf()) == types.FAILURE {
		return 1
	}
	return zend.ZEND_NORMALIZE_BOOL(result.Long())
}
func SplObjectStorageCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	var zo1 *types.ZendObject = (*types.ZendObject)(o1.Object())
	var zo2 *types.ZendObject = (*types.ZendObject)(o2.Object())
	if zo1.GetCe() != spl_ce_SplObjectStorage || zo2.GetCe() != spl_ce_SplObjectStorage {
		return 1
	}
	return types.ZendHashCompare(Z_SPLOBJSTORAGE_P(o1).GetStorage(), Z_SPLOBJSTORAGE_P(o2).GetStorage(), SplObjectStorageCompareInfo, 0)
}
func spl_SplObjectStorage_new(class_type *types.ClassEntry) *types.ZendObject {
	return SplObjectStorageNewEx(class_type, nil)
}
func SplObjectStorageContains(intern *SplObjectStorage, this *types.Zval, obj *types.Zval) int {
	var found int
	var key types.ArrayKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == types.FAILURE {
		return 0
	}
	if key.IsStrKey() {
		found = types.IntBool(intern.GetStorage().KeyExists(key.StrKey()))
	} else {
		found = types.IntBool(intern.GetStorage().IndexExists(key.IdxKey()))
	}
	return found
}
func zim_spl_SplObjectStorage_attach(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var inf *types.Zval = nil
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if zend.ZendParseParameters(executeData.NumArgs(), "o|z!", &obj, &inf) == types.FAILURE {
		return
	}
	SplObjectStorageAttach(intern, executeData.ThisObjectZval(), obj, inf)
}
func zim_spl_SplObjectStorage_detach(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types.FAILURE {
		return
	}
	SplObjectStorageDetach(intern, executeData.ThisObjectZval(), obj)
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
}
func zim_spl_SplObjectStorage_getHash(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types.FAILURE {
		return
	}
	return_value.SetStringVal(PhpSplObjectHash(obj))
	return
}
func zim_spl_SplObjectStorage_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var element *SplObjectStorageElement
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var key types.ArrayKey
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types.FAILURE {
		return
	}
	if SplObjectStorageGetHash(&key, intern, executeData.ThisObjectZval(), obj) == types.FAILURE {
		return
	}
	element = SplObjectStorageGet(intern, &key)
	if element == nil {
		faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Object not found")
	} else {
		var value *types.Zval = element.GetInf()
		types.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_spl_SplObjectStorage_addAll(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var other *SplObjectStorage
	if zend.ZendParseParameters(executeData.NumArgs(), "O", &obj, spl_ce_SplObjectStorage) == types.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)
	SplObjectStorageAddall(intern, executeData.ThisObjectZval(), other)
	return_value.SetLong(intern.GetStorage().Len())
	return
}
func zim_spl_SplObjectStorage_removeAll(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var other *SplObjectStorage
	var element *SplObjectStorageElement
	if zend.ZendParseParameters(executeData.NumArgs(), "O", &obj, spl_ce_SplObjectStorage) == types.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)

	other.GetStorage().ResetInternalPointer()
	for {
		pair := other.GetStorage().Current()
		if pair == nil {
			break
		}
		var element *SplObjectStorageElement = pair.GetVal().Ptr()
		if SplObjectStorageDetach(intern, executeData.ThisObjectZval(), element.GetObj()) == types.FAILURE {
			other.GetStorage().MoveNext()
		}
	}
	intern.SetPos(other.GetStorage().FirstPos())
	intern.SetIndex(0)
	return_value.SetLong(intern.GetStorage().Len())
	return
}
func zim_spl_SplObjectStorage_removeAllExcept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var other *SplObjectStorage
	var element *SplObjectStorageElement
	if zend.ZendParseParameters(executeData.NumArgs(), "O", &obj, spl_ce_SplObjectStorage) == types.FAILURE {
		return
	}
	other = Z_SPLOBJSTORAGE_P(obj)
	intern.GetStorage().Foreach(func(key types.ArrayKey, value *types.Zval) {
		element = value.Ptr()
		if SplObjectStorageContains(other, executeData.ThisObjectZval(), element.GetObj()) == 0 {
			SplObjectStorageDetach(intern, executeData.ThisObjectZval(), element.GetObj())
		}
	})

	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
	return_value.SetLong(intern.GetStorage().Len())
	return
}
func zim_spl_SplObjectStorage_contains(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types.FAILURE {
		return
	}
	return_value.SetBool(SplObjectStorageContains(intern, executeData.ThisObjectZval(), obj) != 0)
	return
}
func zim_spl_SplObjectStorage_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var mode zend.ZendLong = array.COUNT_NORMAL
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &mode) == types.FAILURE {
		return
	}
	if mode == array.COUNT_RECURSIVE {
		var ret zend.ZendLong
		if mode != array.COUNT_RECURSIVE {
			ret = intern.GetStorage().Len()
		} else {
			ret = array.PhpCountRecursive(intern.GetStorage())
		}
		return_value.SetLong(ret)
		return
		return
	}
	return_value.SetLong(intern.GetStorage().Len())
	return
}
func zim_spl_SplObjectStorage_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	intern.SetIndex(0)
}
func zim_spl_SplObjectStorage_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetBool(types.ZendHashHasMoreElementsEx(intern.GetStorage(), intern.GetPos()))
	return
}
func zim_spl_SplObjectStorage_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(intern.GetIndex())
	return
}
func zim_spl_SplObjectStorage_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var element *SplObjectStorageElement
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	types.ZVAL_COPY(return_value, element.GetObj())
}
func zim_spl_SplObjectStorage_getInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var element *SplObjectStorageElement
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	types.ZVAL_COPY(return_value, element.GetInf())
}
func zim_spl_SplObjectStorage_setInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var element *SplObjectStorageElement
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var inf *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &inf) == types.FAILURE {
		return
	}
	if b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) == nil {
		return
	}
	// zend.ZvalPtrDtor(element.GetInf())
	types.ZVAL_COPY(element.GetInf(), inf)
}
func zim_spl_SplObjectStorage_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	intern.GetIndex()++
}
func zim_spl_SplObjectStorage_serialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var element *SplObjectStorageElement
	var members types.Zval
	var flags types.Zval
	var pos types.ArrayPosition
	if !executeData.CheckNumArgsNone(false) {
		return
	}

	serializer := standard.InitSerializer()

	/* storage */
	serializer.WriteString("x:")
	flags.SetLong(intern.GetStorage().Len())
	serializer.Serialize(&flags)

	types.ZendHashInternalPointerResetEx(intern.GetStorage(), &pos)
	for types.ZendHashHasMoreElementsEx(intern.GetStorage(), &pos) {
		if b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), &pos)) == nil {
			serializer.DestroyData()
			return_value.SetNull()
			return
		}
		serializer.Serialize(element.GetObj())
		serializer.WriteByte(',')
		serializer.Serialize(element.GetInf())
		serializer.WriteByte(';')
		types.ZendHashMoveForwardEx(intern.GetStorage(), &pos)
	}

	/* members */
	serializer.WriteString("m:")
	members.SetArray(types.ZendArrayDup(zend.ZendStdGetProperties(executeData.ThisObjectZval())))
	serializer.Serialize(&members)

	/* done */
	serializer.DestroyData()

	return_value.SetStringVal(serializer.String())
}
func zim_spl_SplObjectStorage_unserialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var entry types.Zval
	var inf types.Zval
	var pcount *types.Zval
	var pmembers *types.Zval
	var element *SplObjectStorageElement
	var count zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "s", &buf, &buf_len) == types.FAILURE {
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
	if standard.PhpVarUnserialize(pcount, &p, s+buf_len, &var_hash) == 0 || !pcount.IsLong() {
		goto outexcept
	}
	p--
	count = pcount.Long()
	if count < 0 {
		goto outexcept
	}
	entry.SetUndef()
	inf.SetUndef()
	for b.PostDec(&count) > 0 {
		var pelement *SplObjectStorageElement
		var key types.ArrayKey
		if (*p) != ';' {
			goto outexcept
		}
		p++
		if (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}

		/* store reference to allow cross-references between different elements */

		if standard.PhpVarUnserialize(&entry, &p, s+buf_len, &var_hash) == 0 {
			// zend.ZvalPtrDtor(&entry)
			goto outexcept
		}
		if (*p) == ',' {
			p++
			if standard.PhpVarUnserialize(&inf, &p, s+buf_len, &var_hash) == 0 {
				// zend.ZvalPtrDtor(&entry)
				// zend.ZvalPtrDtor(&inf)
				goto outexcept
			}
		}
		if !entry.IsObject() {
			// zend.ZvalPtrDtor(&entry)
			// zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		if SplObjectStorageGetHash(&key, intern, executeData.ThisObjectZval(), &entry) == types.FAILURE {
			// zend.ZvalPtrDtor(&entry)
			// zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		pelement = SplObjectStorageGet(intern, &key)
		if pelement != nil {
			if !(pelement.GetInf().IsUndef()) {
				standard.VarPushDtor(&var_hash, pelement.GetInf())
			}
			if !(pelement.GetObj().IsUndef()) {
				standard.VarPushDtor(&var_hash, pelement.GetObj())
			}
		}
		element = SplObjectStorageAttach(intern, executeData.ThisObjectZval(), &entry, b.Cond(inf.IsUndef(), nil, &inf))
		standard.VarReplace(&var_hash, &entry, element.GetObj())
		standard.VarReplace(&var_hash, &inf, element.GetInf())
		// zend.ZvalPtrDtor(&entry)
		entry.SetUndef()
		// zend.ZvalPtrDtor(&inf)
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
	if standard.PhpVarUnserialize(pmembers, &p, s+buf_len, &var_hash) == 0 || !pmembers.IsArray() {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(intern.GetStd(), pmembers.Array())
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	return
outexcept:
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset %zd of %zd bytes", (*byte)(p-buf), buf_len)
	return
}
func zim_spl_SplObjectStorage___serialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var elem *SplObjectStorageElement
	var tmp types.Zval
	if !executeData.CheckNumArgsNone(true) {
		return
	}
	zend.ArrayInit(return_value)

	/* storage */

	zend.ArrayInitSize(&tmp, 2*intern.GetStorage().Len())
	intern.GetStorage().Foreach(func(key types.ArrayKey, value *types.Zval) {
		elem = value.Ptr()
		tmp.Array().Append(elem.GetObj())
		tmp.Array().Append(elem.GetInf())
	})
	return_value.Array().Append(&tmp)

	/* members */

	tmp.SetArray(zend.ZendStdGetProperties(executeData.ThisObjectZval()))
	//tmp.TryAddRefcount()
	return_value.Array().Append(&tmp)
}
func zim_spl_SplObjectStorage___unserialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	var data *types.Array
	var storage_zv *types.Zval
	var members_zv *types.Zval
	var key *types.Zval
	var val *types.Zval
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "h", &data) == types.FAILURE {
		return
	}
	storage_zv = data.IndexFind(0)
	members_zv = data.IndexFind(1)
	if storage_zv == nil || members_zv == nil || !storage_zv.IsArray() || !members_zv.IsArray() {
		faults.ThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	if storage_zv.Array().Len()%2 != 0 {
		faults.ThrowException(spl_ce_UnexpectedValueException, "Odd number of elements", 0)
		return
	}
	key = nil
	storage_zv.Array().Foreach(func(_ types.ArrayKey, value *types.Zval) {
		if key != nil {
			if !key.IsObject() {
				faults.ThrowException(spl_ce_UnexpectedValueException, "Non-object key", 0)
				return
			}
			SplObjectStorageAttach(intern, executeData.ThisObjectZval(), value, val)
			key = nil
		} else {
			key = value
		}
	})

	zend.ObjectPropertiesLoad(intern.GetStd(), members_zv.Array())
}
func zim_spl_SplObjectStorage___debugInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetArray(SplObjectStorageDebugInfo(zend.getThis()))
	return
}
func zim_spl_MultipleIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	var flags zend.ZendLong = MIT_NEED_ALL | MIT_KEYS_NUMERIC
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "|l", &flags) == types.FAILURE {
		return
	}
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	intern.SetFlags(flags)
}
func zim_spl_MultipleIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_MultipleIterator_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if zend.ZendParseParameters(executeData.NumArgs(), "l", intern.GetFlags()) == types.FAILURE {
		return
	}
}
func zim_spl_MultipleIterator_attachIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	var iterator *types.Zval = nil
	var info *types.Zval = nil
	if zend.ZendParseParameters(executeData.NumArgs(), "O|z!", &iterator, zend.ZendCeIterator, &info) == types.FAILURE {
		return
	}
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if info != nil {
		var element *SplObjectStorageElement
		if !info.IsLong() && !info.IsString() {
			faults.ThrowException(spl_ce_InvalidArgumentException, "Info must be NULL, integer or string", 0)
			return
		}
		types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
		for b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil {
			if operators.FastIsIdenticalFunction(info, element.GetInf()) != 0 {
				faults.ThrowException(spl_ce_InvalidArgumentException, "Key duplication error", 0)
				return
			}
			types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
		}
	}
	SplObjectStorageAttach(intern, executeData.ThisObjectZval(), iterator, info)
}
func zim_spl_MultipleIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	var element *SplObjectStorageElement
	var it *types.Zval
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfRewind(), "rewind", nil)
		types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	var element *SplObjectStorageElement
	var it *types.Zval
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfNext(), "next", nil)
		types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	var element *SplObjectStorageElement
	var it *types.Zval
	var retval types.Zval
	var expect zend.ZendLong
	var valid zend.ZendLong
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(intern.GetStorage().Len()) {
		return_value.SetFalse()
		return
	}
	if intern.IsNeedAll() {
		expect = 1
	} else {
		expect = 0
	}
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfValid(), "valid", &retval)
		if !(retval.IsUndef()) {
			valid = retval.IsType(types.IS_TRUE)
			// zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if expect != valid {
			return_value.SetBool(expect == 0)
			return
		}
		types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
	return_value.SetBool(expect != 0)
	return
}
func SplMultipleIteratorGetAll(intern *SplObjectStorage, get_type int, return_value *types.Zval) {
	var element *SplObjectStorageElement
	var it *types.Zval
	var retval types.Zval
	var valid int = 1
	var num_elements int
	num_elements = intern.GetStorage().Len()
	if num_elements < 1 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInitSize(return_value, num_elements)
	types.ZendHashInternalPointerResetEx(intern.GetStorage(), intern.GetPos())
	for b.Assign(&element, types.ZendHashGetCurrentDataPtrEx(intern.GetStorage(), intern.GetPos())) != nil && zend.EG__().GetException() == nil {
		it = element.GetObj()
		zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfValid(), "valid", &retval)
		if !(retval.IsUndef()) {
			valid = retval.IsType(types.IS_TRUE)
			// zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if valid != 0 {
			if SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT == get_type {
				zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfCurrent(), "current", &retval)
			} else {
				zend.ZendCallMethodWith0Params(it, types.Z_OBJCE_P(it), types.Z_OBJCE_P(it).GetIteratorFuncsPtr().GetZfKey(), "key", &retval)
			}
			if retval.IsUndef() {
				faults.ThrowException(spl_ce_RuntimeException, "Failed to call sub iterator method", 0)
				return
			}
		} else if intern.IsNeedAll() {
			if SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT == get_type {
				faults.ThrowException(spl_ce_RuntimeException, "Called current() with non valid sub iterator", 0)
			} else {
				faults.ThrowException(spl_ce_RuntimeException, "Called key() with non valid sub iterator", 0)
			}
			return
		} else {
			retval.SetNull()
		}
		if intern.IsKeysAssoc() {
			switch element.GetInf().GetType() {
			case types.IS_LONG:
				zend.AddIndexZval(return_value, element.GetInf().Long(), &retval)
			case types.IS_STRING:
				return_value.Array().SymtableUpdate(element.GetInf().String().GetStr(), &retval)
			default:
				// zend.ZvalPtrDtor(&retval)
				faults.ThrowException(spl_ce_InvalidArgumentException, "Sub-Iterator is associated with NULL", 0)
				return
			}
		} else {
			zend.AddNextIndexZval(return_value, &retval)
		}
		types.ZendHashMoveForwardEx(intern.GetStorage(), intern.GetPos())
	}
}
func zim_spl_MultipleIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplMultipleIteratorGetAll(intern, SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT, return_value)
}
func zim_spl_MultipleIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplObjectStorage
	intern = Z_SPLOBJSTORAGE_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplMultipleIteratorGetAll(intern, SPL_MULTIPLE_ITERATOR_GET_ALL_KEY, return_value)
}
func ZmStartupSplObserver(type_ int, module_number int) int {
	spl_ce_SplObserver = zend.RegisterInternalInterface("SplObserver", spl_funcs_SplObserver)
	spl_ce_SplSubject = zend.RegisterInternalInterface("SplSubject", spl_funcs_SplSubject)
	spl_ce_SplObjectStorage = zend.RegisterClass("SplObjectStorage", spl_SplObjectStorage_new, spl_funcs_SplObjectStorage)

	spl_handler_SplObjectStorage = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		Offset:         int((*byte)(&((*SplObjectStorage)(nil).GetStd())) - (*byte)(nil)),
		CompareObjects: SplObjectStorageCompareObjects,
		CloneObj:       SplObjectStorageClone,
		DtorObj:        zend.ZendObjectsDestroyObject,
		FreeObj:        spl_SplObjectStorage_free_storage,
	})

	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Countable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, spl_ce_ArrayAccess)
	spl_ce_MultipleIterator = zend.RegisterClass("MultipleIterator", spl_SplObjectStorage_new, spl_funcs_MultipleIterator)
	zend.ZendClassImplements(spl_ce_MultipleIterator, 1, zend.ZendCeIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ANY", MIT_NEED_ANY)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ALL", MIT_NEED_ALL)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_NUMERIC", MIT_KEYS_NUMERIC)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_ASSOC", MIT_KEYS_ASSOC)
	return types.SUCCESS
}
