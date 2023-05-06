package array

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"sort"
)

func phpSplice(arr *types.Array, offset int, length int, replace *types.Array) (*types.Array, *types.Array) {
	arrLen := arr.Len()
	// check offset [0,arrLen-1]
	if offset > arrLen {
		offset = arrLen
	} else if offset < 0 {
		offset += arrLen
		if offset < 0 {
			offset = 0
		}
	}
	// check length [0, arrLen - offset]
	if length < 0 {
		length = arrLen - offset + length
	} else if offset+length > arrLen {
		length = arrLen - offset
	}
	if length < 0 {
		length = 0
	}

	outHash := types.NewArray(0)
	removed := types.NewArray(0)
	pairs := arr.Pairs()

	// handle range [0, offset)
	for _, pair := range pairs[:offset] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	// handle range [offset, offset+length)
	for _, pair := range pairs[offset : offset+length] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			removed.Append(val)
		} else {
			removed.KeyAdd(key.StrKey(), val)
		}
	}

	// handle insert
	if replace != nil {
		replace.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			outHash.Append(value)
		})
	}

	// handle range [offset+length, len(array))
	for _, pair := range pairs[offset+length:] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	return outHash, removed
}
func ZifArrayPush(stack zpp.RefArray, _ zpp.Opt, args []*types.Zval) (int, bool) {
	for _, arg := range args {
		if stack.Array().Append(arg) == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return 0, false
		}
	}

	return stack.Array().Len(), true
}
func ZifArrayPop(stack zpp.RefArray) *types.Zval {
	pair := stack.Array().LastPairIndirect()
	if pair == nil {
		return types.NewZvalNull()
	}

	stack.Array().Delete(pair.GetKey())
	stack.Array().ResetInternalPointer()

	return pair.GetVal().DeRef()
}
func ZifArrayShift(stack zpp.RefArray) {
	if stack.Array().Len() == 0 {
		return
	}

	var p *types.Bucket

	/* Get the first value and copy it into the return value */
	key, val := stack.Array().FirstIndirect()
	if val == nil {
		return
	}

	val = val.DeRef()

	/* Delete the first value */
	if key.IsStrKey() && stack.Array() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		stack.Array().Delete(key)
	}

	/* re-index like it did before */
	var k = 0
	stack.Array().MapWithKey(func(key types.ArrayKey, value *types.Zval) (types.ArrayKey, *types.Zval) {
		if !key.IsStrKey() {
			key = types.IdxKey(k)
			k++
		}
		return key, value
	})

	// reset internal pointer
	stack.Array().ResetInternalPointer()
}
func ZifArrayUnshift(stack zpp.RefZval, values []*types.Zval) int {
	newArr := types.NewArray(stack.Array().Len() + len(values))
	for _, value := range values {
		newArr.Append(value)
	}
	stack.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			newArr.KeyAdd(key.StrKey(), value)
		} else {
			newArr.Append(value)
		}
	})

	stack.SetArray(newArr)
	return stack.Array().Len()
}
func ZifArraySplice(array zpp.RefArray, offset int, _ zpp.Opt, length_ *int, replacement *types.Zval) *types.Array {
	var numIn = array.Array().Len()
	var length = b.Option(length_, numIn)
	var replaceArr *types.Array = nil

	if replacement != nil {
		/* Make sure the last argument, if passed, is an array */
		zend.ConvertToArrayEx(replacement)
		replaceArr = replacement.Array()
	}

	newArr, removedArr := phpSplice(array.Array(), offset, length, replaceArr)
	array.SetArray(newArr)
	return removedArr
}
func ZifArraySlice(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, offset *types.Zval, _ zpp.Opt, length *types.Zval, preserveKeys *types.Zval) {
	var input *types.Zval
	var z_length *types.Zval = nil
	var entry *types.Zval
	var offset zend.ZendLong
	var length = 0
	var preserve_keys = 0
	var num_in int
	var pos int
	var string_key *types.String
	var num_key zend.ZendUlong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			input = fp.ParseArray()
			offset = fp.ParseLong()
			fp.StartOptional()
			z_length = fp.ParseZval()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Get number of entries in the input hash */

	num_in = input.Array().Len()

	/* We want all entries from offset to the end if length is not passed or is null */

	if executeData.NumArgs() < 3 || z_length.IsType(types.IS_NULL) {
		length = num_in
	} else {
		length = zend.ZvalGetLong(z_length)
	}

	/* Clamp the offset.. */

	if offset > num_in {
		return_value.SetEmptyArray()
		return
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > unsigned(num_in) {
		length = num_in - offset
	}
	if length <= 0 {
		return_value.SetEmptyArray()
		return
	}

	/* Initialize returned array */

	zend.ArrayInitSize(return_value, uint32(length))

	/* Start at the beginning and go until we hit offset */

	pos = 0
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		pos++
		if pos <= offset {
			continue
		}
		if pos > offset+length {
			break
		}
		if string_key != nil {
			entry = return_value.Array().KeyAddNew(string_key.GetStr(), entry)
		} else {
			if preserve_keys != 0 {
				entry = return_value.Array().IndexAddNew(num_key, entry)
			} else {
				entry = return_value.Array().AppendNew(entry)
			}
		}
		zend.ZvalAddRef(entry)
	}
}
func PhpArrayMergeRecursive(dest *types.Array, src *types.Array) int {
	for iter := src.Iterator(); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Current()
		if !key.IsStrKey() {
			dest.Append(value)
			continue
		}

		strKey := key.StrKey()
		destEntry := dest.KeyFind(strKey)
		if destEntry == nil {
			dest.KeyAddNew(strKey, value)
		}

		var srcZval = value.DeRef()
		var destZval = destEntry.DeRef()
		var thash *types.Array
		var tmp types.Zval
		var ret int
		if destZval.IsType(types.IS_ARRAY) {
			thash = destZval.Array()
		} else {
			thash = nil
		}
		if thash != nil && thash.IsRecursive() || value == destEntry && destEntry.IsReference() && destEntry.GetRefcount()%2 != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		b.Assert(!(destEntry.IsReference()) || destEntry.GetRefcount() > 1)
		types.SeparateZval(destEntry)
		destZval = destEntry
		if destZval.IsType(types.IS_NULL) {
			zend.ConvertToArrayEx(destZval)
			zend.AddNextIndexNull(destZval)
		} else {
			zend.ConvertToArrayEx(destZval)
		}
		tmp.SetUndef()
		if srcZval.IsType(types.IS_OBJECT) {
			types.ZVAL_COPY(&tmp, srcZval)
			zend.ConvertToArray(&tmp)
			srcZval = &tmp
		}
		if srcZval.IsType(types.IS_ARRAY) {
			if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
				thash.ProtectRecursive()
			}
			ret = PhpArrayMergeRecursive(destZval.Array(), srcZval.Array())
			if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
				thash.UnprotectRecursive()
			}
			if ret == 0 {
				return 0
			}
		} else {
			destZval.Array().Append(srcZval)
		}
	}
	return 1
}
func PhpArrayMerge(dest *types.Array, src *types.Array) {
	src.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if value.IsReference() && value.GetRefcount() == 1 {
			value = types.Z_REFVAL_P(value)
		}
		if key.IsStrKey() {
			dest.KeyUpdate(key.StrKey(), value)
		} else {
			dest.Append(value)
		}
	})
}
func PhpArrayReplaceRecursive(dest *types.Array, src *types.Array) int {
	var src_entry *types.Zval
	var dest_entry *types.Zval
	var src_zval *types.Zval
	var dest_zval *types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var ret int
	var __ht = src
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		src_entry = _z
		src_zval = src_entry
		src_zval = types.ZVAL_DEREF(src_zval)
		if string_key != nil {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.KeyFind(string_key.GetStr())) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv = dest.KeyUpdate(string_key.GetStr(), src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		} else {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.IndexFind(num_key)) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv = dest.IndexUpdate(num_key, src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		}
		dest_zval = dest_entry
		dest_zval = types.ZVAL_DEREF(dest_zval)
		if dest_zval.IsRecursive() || src_zval.IsRecursive() || src_entry.IsReference() && dest_entry.IsReference() && src_entry.Reference() == dest_entry.Reference() && dest_entry.GetRefcount()%2 != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		b.Assert(!(dest_entry.IsReference()) || dest_entry.GetRefcount() > 1)
		types.SeparateZval(dest_entry)
		dest_zval = dest_entry
		if dest_zval.IsRefcounted() {
			dest_zval.ProtectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.ProtectRecursive()
		}
		ret = PhpArrayReplaceRecursive(dest_zval.Array(), src_zval.Array())
		if dest_zval.IsRefcounted() {
			dest_zval.UnprotectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.UnprotectRecursive()
		}
		if ret == 0 {
			return 0
		}
	}
	return 1
}
func PhpArrayReplaceWrapper(executeData *zend.ZendExecuteData, return_value *types.Zval, recursive int) {
	var args *types.Zval = nil
	var arg *types.Zval
	var argc int
	var i int
	var dest *types.Array
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	for i = 0; i < argc; i++ {
		var arg *types.Zval = args + i
		if arg.GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(arg))
			return_value.SetNull()
			return
		}
	}

	/* copy first array */

	arg = args
	dest = types.ZendArrayDup(arg.Array())
	return_value.SetArray(dest)
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayReplaceRecursive(dest, arg.Array())
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			types.ZendHashMerge(dest, arg.Array(), zend.ZvalAddRef, 1)
		}
	}
}
func arrayMergeWrapper(args []*types.Zval, recursive bool) *types.Array {
	var dest *types.Array
	if len(args) == 0 {
		return types.NewArray(0)
	}

	count := 0
	for i, arg := range args {
		if !arg.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(arg))
			return nil
		}
		count += arg.Array().Len()
	}

	arr := types.NewArray(count)

	args[0].Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if value.IsReference() && value.GetRefcount() == 1 {
			value = types.Z_REFVAL_P(value)
		}
		if key.IsStrKey() {
			arr.KeyUpdate(key.StrKey(), value)
		} else {
			arr.Append(value)
		}
	})

	if recursive {
		for _, arg := range args[1:] {
			PhpArrayMergeRecursive(dest, arg.Array())
		}
	} else {
		for _, arg := range args[1:] {
			PhpArrayMerge(dest, arg.Array())
		}
	}
	return arr
}
func ZifArrayMerge(returnValue zpp.Ret, _ zpp.Opt, arrays []*types.Zval) {
	arr := arrayMergeWrapper(arrays, false)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}
func ZifArrayMergeRecursive(returnValue zpp.Ret, _ zpp.Opt, arrays []*types.Zval) {
	arr := arrayMergeWrapper(arrays, true)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}
func ZifArrayReplace(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, _ zpp.Opt, arrays []*types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 0)
}
func ZifArrayReplaceRecursive(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, _ zpp.Opt, arrays []*types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 1)
}
func ZifArrayKeys(array *types.Array, _ zpp.Opt, searchValue *types.Zval, strict bool) *types.Array {
	arrLen := array.Len()

	/* Base case: empty input */
	if arrLen == 0 {
		return types.NewArray(0)
	}

	/* Initialize return array */
	if searchValue != nil {
		keys := types.NewArray(0)
		if strict {
			array.ForeachIndirect(func(key types.ArrayKey, entry *types.Zval) {
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(searchValue, entry) != 0 {
					keys.Append(key.ToZval())
				}
			})
		} else {
			array.ForeachIndirect(func(key types.ArrayKey, entry *types.Zval) {
				if zend.FastEqualCheckFunction(searchValue, entry) != 0 {
					keys.Append(key.ToZval())
				}
			})
		}
		return keys
	} else {
		keys := types.NewArray(arrLen)
		array.ForeachIndirect(func(key types.ArrayKey, _ *types.Zval) {
			keys.Append(key.ToZval())
		})
		return keys
	}
}
func ZifArrayKeyFirst(array *types.Array) *types.Zval {
	key, val := array.FirstIndirect()
	if val == nil {
		return types.NewZvalNull()
	}
	return key.ToZval()
}
func ZifArrayKeyLast(array *types.Array) *types.Zval {
	key, val := array.Last()
	if val == nil {
		return types.NewZvalNull()
	}
	return key.ToZval()
}
func ZifArrayValues(array *types.Array) *types.Array {
	arrLen := array.Len()
	if arrLen == 0 {
		return types.NewArray(0)
	}

	/* Initialize return array */
	values := types.NewArray(array.Len())
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		if entry.IsReference() && entry.GetRefcount() == 1 {
			entry = types.Z_REFVAL_P(entry)
		}
		values.Append(entry)
	})
	return values
}
func ZifArrayCountValues(executeData zpp.Ex, return_value zpp.Ret, array *types.Array) *types.Array {
	var input *types.Zval
	var entry *types.Zval
	var tmp *types.Zval
	var myht *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			input = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)

	/* Go through input array and add values to the return array */
	var __ht = myht
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if b.Assign(&tmp, return_value.Array().IndexFind(entry.Long())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.Array().IndexUpdate(entry.Long(), &data)
			} else {
				tmp.Long()++
			}
		} else if entry.IsString() {
			if b.Assign(&tmp, return_value.Array().SymtableFind(entry.String().GetStr())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.Array().SymtableUpdate(entry.String().GetStr(), &data)
			} else {
				tmp.Long()++
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only count STRING and INTEGER values!")
		}
	}
}
func ArrayColumnParamHelper(param *types.Zval, name string) types.ZendBool {
	switch param.GetType() {
	case types.IS_DOUBLE:
		if param.GetType() != types.IS_LONG {
			zend.ConvertToLong(param)
		}
		fallthrough
	case types.IS_LONG:
		return 1
	case types.IS_OBJECT:
		if zend.TryConvertToString(param) == 0 {
			return 0
		}
		fallthrough
	case types.IS_STRING:
		return 1
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The %s key should be either a string or an integer", name)
		return 0
	}
}
func ArrayColumnFetchProp(data *types.Zval, name *types.Zval) *types.Zval {
	var rv types.Zval
	var prop *types.Zval = nil
	if data.IsType(types.IS_OBJECT) {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_EXISTS, nil) != 0 || types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_ISSET, nil) != 0 {
			prop = types.Z_OBJ_HT(*data).GetReadProperty()(data, name, zend.BP_VAR_R, nil, &rv)
			if prop != nil {
				prop = types.ZVAL_DEREF(prop)
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.IsType(types.IS_ARRAY) {
		if name.IsString() {
			prop = data.Array().SymtableFind(name.String().GetStr())
		} else if name.IsType(types.IS_LONG) {
			prop = data.Array().IndexFind(name.Long())
		}
		if prop != nil {
			prop = types.ZVAL_DEREF(prop)
		}
	}
	return prop
}
func ZifArrayColumn(array *types.Array, columnKey zpp.ZvalNullable, _ zpp.Opt, indexKey zpp.ZvalNullable) (*types.Array, bool) {
	if columnKey != nil && ArrayColumnParamHelper(columnKey, "column") == 0 || indexKey != nil && ArrayColumnParamHelper(indexKey, "index") == 0 {
		return nil, false
	}

	retArr := types.NewArray(array.Len())
	if indexKey == nil {
		array.Foreach(func(_ types.ArrayKey, data *types.Zval) {
			var columnVal *types.Zval
			data = types.ZVAL_DEREF(data)
			if columnKey == nil {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(data, columnKey); columnVal == nil {
				return
			}

			retArr.Append(columnVal)
		})
	} else {
		array.Foreach(func(key types.ArrayKey, data *types.Zval) {
			data = types.ZVAL_DEREF(data)

			// col
			var columnVal *types.Zval
			if columnKey == nil {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(data, columnKey); columnVal == nil {
				return
			}

			// key
			var keyVal = ArrayColumnFetchProp(data, indexKey)
			if keyVal != nil {
				switch keyVal.GetType() {
				case types.IS_STRING:
					retArr.SymtableUpdate(keyVal.String().GetStr(), columnVal)
				case types.IS_LONG:
					retArr.IndexUpdate(keyVal.Long(), columnVal)
				case types.IS_OBJECT:
					retArr.SymtableUpdate(zend.ZvalGetStrVal(keyVal), columnVal)
				case types.IS_NULL:
					retArr.KeyUpdate("", columnVal)
				case types.IS_DOUBLE:
					retArr.IndexUpdate(zend.DvalToLval(keyVal.Double()), columnVal)
				case types.IS_TRUE:
					retArr.IndexUpdate(1, columnVal)
				case types.IS_FALSE:
					retArr.IndexUpdate(0, columnVal)
				case types.IS_RESOURCE:
					retArr.IndexUpdate(types.Z_RES_HANDLE_P(keyVal), columnVal)
				default:
					retArr.Append(columnVal)
				}
			} else {
				retArr.Append(columnVal)
			}

		})
	}
	return retArr, true
}
func ZifArrayReverse(array *types.Array, _ zpp.Opt, preserveKeys bool) *types.Array {
	retArr := types.NewArray(array.Len())
	array.ForeachReserve(func(key types.ArrayKey, value *types.Zval) {
		if preserveKeys || key.IsStrKey() {
			retArr.Add(key, value)
		} else { // !preserveKeys && !key.IsStrKey()
			retArr.Append(value)
		}
	})
	return retArr
}

const maxPadInOneTimes = 1048576

func ZifArrayPad(array *types.Array, padSize int, padValue *types.Zval) (*types.Array, bool) {
	/* Do some initial calculations */
	inputSize := array.Len()
	padSizeAbs := padSize
	if padSizeAbs < 0 {
		padSizeAbs = -padSizeAbs
	}
	if padSizeAbs < 0 || padSizeAbs-inputSize > maxPadInOneTimes {
		core.PhpErrorDocref(nil, faults.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return nil, false
	}

	/**
	 * notice: 注意此处的差别:
	 * - 当原数组的长度 >= padSize 时，返回数组保留了原数组包括内部指针等信息；
	 * - 当原数组的长度 < padSize 时，产生的新数组内部指针等信息未设置(等于初始值)；
	 */
	if inputSize >= padSizeAbs {
		return array.Copy(), true
	}

	numPads := padSizeAbs - inputSize
	retArr := types.NewArray(padSizeAbs)
	if padSize < 0 {
		for i := 0; i < numPads; i++ {
			retArr.Append(padValue)
		}
	}
	array.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			retArr.KeyAdd(key.StrKey(), value)
		} else {
			retArr.Append(value)
		}
	})
	if padSize > 0 {
		for i := 0; i < numPads; i++ {
			retArr.Append(padValue)
		}
	}
	return retArr, true
}
func ZifArrayFlip(array *types.Array) *types.Array {
	retArr := types.NewArray(array.Len())
	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
		value = value.DeRef()
		if value.IsLong() {
			retArr.IndexUpdate(value.Long(), key.ToZval())
		} else if value.IsString() {
			retArr.SymtableUpdate(value.StringVal(), key.ToZval())
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	})
	return retArr
}
func ZifArrayChangeKeyCase(array *types.Array, _ zpp.Opt, case_ *int) *types.Array {
	var caseFlag = b.Option(case_, CASE_LOWER)
	retArr := types.NewArray(array.Len())
	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			if caseFlag == CASE_LOWER {
				key = types.StrKey(ascii.StrToLower(key.StrKey()))
			} else {
				key = types.StrKey(ascii.StrToUpper(key.StrKey()))
			}
		}

		retArr.Update(key, value)
	})
	return retArr
}
func ZifArrayUnique(arg *types.Array, _ zpp.Opt, flags *int) *types.Array {
	var sortType = b.Option(flags, PHP_SORT_STRING)

	if sortType == PHP_SORT_STRING {
		existValues := make(map[string]bool, arg.Len())
		retArr := types.NewArray(arg.Len())
		arg.ForeachIndirect(func(key types.ArrayKey, val *types.Zval) {
			var strVal = zend.ZvalGetStrVal(val)
			if _, exists := existValues[strVal]; !exists {
				if val.IsReference() && val.GetRefcount() == 1 {
					val = val.DeRef()
				}
				retArr.Add(key, val)
			}

		})
		return retArr
	}

	cmp := phpGetDataCompareFunc(sortType, false)
	retArr := types.ZendArrayDup(arg)

	// 取出数组所有键值对，先排序，后找出重复元素并删除
	pairs := arg.Pairs()
	sort.SliceStable(pairs, func(i, j int) bool {
		return cmp(pairs[i], pairs[j]) < 0
	})

	last := 0
	for i := 1; i < len(pairs); i++ {
		if cmp(pairs[last], pairs[i]) != 0 {
			last = i
		} else {
			pair := pairs[i]
			retArr.Delete(pair.GetKey())
		}
	}

	return retArr
}
func ZvalCompare(first *types.Zval, second *types.Zval) int {
	return zend.StringCompareFunction(first, second)
}

//@zif -c 2,
func ZifArrayIntersectKey(arrays []*types.Zval) (*types.Array, bool) {
	return arrayIntersectKeyWrapper(arrays, nil)
}

//@zif -c=3,
func ZifArrayIntersectUkey(arrays []*types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(callbackKeyCompareFunc)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayIntersect(arrays []*types.Zval) (*types.Array, bool) {
	cmp := arrayDataComparer(zend.StringCompareFunction)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayUintersect(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserDataComparer(callbackDataCompareFunc)
	return arrayIntersectWrapper(arrays, cmp)
}
func ZifArrayIntersectAssoc(arrays []*types.Zval) (*types.Array, bool) {
	cmp := ZvalCompare
	return arrayIntersectKeyWrapper(arrays, cmp)
}
func ZifArrayIntersectUassoc(arrays []*types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompareFunc),
		arrayDataComparer(zend.StringCompareFunction),
	)
	return arrayIntersectWrapper(arrays, cmp)
}
func ZifArrayUintersectAssoc(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserZvalComparer(callbackDataCompareFunc)
	return arrayIntersectKeyWrapper(arrays, cmp)
}

//@zif -c=4,
func ZifArrayUintersectUassoc(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompareFunc),
		arrayUserDataComparer(callbackDataCompareFunc),
	)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayDiffKey(arrays []*types.Zval) (*types.Array, bool) {
	return arrayDiffKeyWrapper(arrays, nil)
}

//@zif -c=3,
func ZifArrayDiffUkey(arrays []*types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(callbackKeyCompFunc)
	return arrayDiffWrapper(arrays, cmp)
}

func simpleArrayDiff(array *types.Array, arrays []*types.Array) *types.Array {
	if array.Len() == 0 {
		return types.NewArray(0)
	}

	// array.Len() > 1
	exclude := make(map[string]bool)
	for _, diffArray := range arrays {
		diffArray.ForeachIndirect(func(_ types.ArrayKey, value *types.Zval) {
			str := zend.ZvalGetStrVal(value)
			exclude[str] = true
		})
	}
	if len(exclude) == 0 {
		return array
	}

	retArr := types.NewArray(array.Len())
	array.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		str := zend.ZvalGetStrVal(value)
		if _, excluded := exclude[str]; !excluded {
			retArr.Add(key, value)
		}
	})
	return retArr
}

//@zif -c=2,
func ZifArrayDiff(arrays []*types.Zval) *types.Zval {
	arrayHts, ok := checkArrayArgs(arrays, 0)
	if !ok {
		return types.NewZvalNull()
	}
	retArr := simpleArrayDiff(arrayHts[0], arrayHts[1:])
	return types.NewZvalArray(retArr)
}

//@zif -c=3,
func ZifArrayUdiff(arrays []*types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserDataComparer(callbackDataCompFunc)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayDiffAssoc(arrays []*types.Zval) (*types.Array, bool) {
	cmp := ZvalCompare
	return arrayDiffKeyWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayDiffUassoc(arrays []*types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompFunc),
		arrayDataComparer(zend.StringCompareFunction),
	)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayUdiffAssoc(arrays []*types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserZvalComparer(callbackDataCompFunc)
	return arrayDiffKeyWrapper(arrays, cmp)
}

//@zif -c=4,
func ZifArrayUdiffUassoc(arrays []*types.Zval, callbackDataCompFunc zpp.Callable, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompFunc),
		arrayUserDataComparer(callbackDataCompFunc),
	)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=1,
func ZifArrayMultisort(args []*types.Zval) bool {
	var parseState = [...]int{0, 0}

	arrays := make([]*types.Zval, 0, len(args))
	multisortFunc := make([]types.ArrayComparer, 0, len(args))

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */
	var sortOrder = PHP_SORT_ASC
	var sortType = PHP_SORT_REGULAR
	for i, arg := range args {
		arg = types.ZVAL_DEREF(arg)
		if arg.IsArray() {
			types.SeparateArray(arg)

			/* We see the next array, so we update the sort flags of
			 * the previous array and reset the sort flags. */
			if i > 0 {
				multisortFunc = append(multisortFunc, phpGetDataCompareFunc(sortType, sortOrder != PHP_SORT_ASC))
				sortOrder = PHP_SORT_ASC
				sortType = PHP_SORT_REGULAR
			}
			arrays = append(arrays, arg)

			/* Next one may be an array or a list of sort flags. */
			for k := 0; k < MULTISORT_LAST; k++ {
				parseState[k] = 1
			}

			/* Next one may be an array or a list of sort flags. */

		} else if arg.IsType(types.IS_LONG) {
			switch arg.Long() & ^PHP_SORT_FLAG_CASE {
			case PHP_SORT_ASC:
				fallthrough
			case PHP_SORT_DESC:
				/* flag allowed here */
				if parseState[MULTISORT_ORDER] == 1 {
					/* Save the flag and make sure then next arg is not the current flag. */
					if arg.Long() == PHP_SORT_DESC {
						sortOrder = PHP_SORT_DESC
					} else {
						sortOrder = PHP_SORT_ASC
					}
					parseState[MULTISORT_ORDER] = 0
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					return false
				}
			case PHP_SORT_REGULAR:
				fallthrough
			case PHP_SORT_NUMERIC:
				fallthrough
			case PHP_SORT_STRING:
				fallthrough
			case PHP_SORT_NATURAL:
				fallthrough
			case PHP_SORT_LOCALE_STRING:
				/* flag allowed here */
				if parseState[MULTISORT_TYPE] == 1 {
					/* Save the flag and make sure then next arg is not the current flag. */
					sortType = int(arg.Long())
					parseState[MULTISORT_TYPE] = 0
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					return false
				}
			default:
				core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is an unknown sort flag", i+1)
				return false
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or a sort flag", i+1)
			return false
		}
	}

	/* Take care of the last array sort flags. */
	multisortFunc = append(multisortFunc, phpGetDataCompareFunc(sortType, sortOrder != PHP_SORT_ASC))

	/* Make sure the arrays are of the same size. */
	arraySize := arrays[0].Array().Len()
	for _, array := range arrays {
		if array.Array().Len() != arraySize {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Array sizes are inconsistent")
			return false
		}
	}

	/* If all arrays are empty we don't need to do anything. */
	if arraySize == 1 {
		return true
	}

	/**
	 * 将所有数组中的键值对取出组成并组成矩阵。(matrix[idx][i] 表示第 i 个数组的第 idx 个键值对)
	 */
	var matrix = make([][]types.ArrayPair, arraySize)
	for _, array := range arrays {
		idx := 0
		array.Array().ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			matrix[idx] = append(matrix[idx], types.MakeArrayPair(key, value))
		})
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */
	sort.SliceStable(matrix, func(idx1, idx2 int) bool {
		col1 := matrix[idx1]
		col2 := matrix[idx2]
		for i, sortFunc := range multisortFunc {
			result := sortFunc(col1[i], col2[i])
			if result < 0 {
				return true
			} else if result > 0 {
				return false
			}
		}
		return false
	})

	/* Restructure the arrays based on sorted indirect - this is mostly taken from zend_hash_sort() function. */
	for i, array := range arrays {
		newHash := types.NewArray(array.Array().Len())
		//newHash.SetNInternalPointer(0)
		for idx := 0; idx < arraySize; idx++ {
			pair := matrix[idx][i]
			if pair.GetKey().IsStrKey() {
				newHash.KeyAdd(pair.GetKey().StrKey(), pair.GetVal())
			} else {
				newHash.Append(pair.GetVal())
			}
		}
		newHash.SetNNextFreeElement(arraySize)
	}

	return true
}
func ZifArrayRand(return_value zpp.Ret, arg *types.Array, _ zpp.Opt, numReq_ *int) *types.Zval {
	var numReq = b.Option(numReq_, 1)
	var string_key *types.String
	var num_key zend.ZendUlong
	var i int
	var numAvail int
	var bitset zend.ZendBitset
	var negative_bitset = 0
	var bitset_len uint32
	numAvail = arg.Len()
	if numAvail == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Array is empty")
		return nil
	}
	if numReq == 1 {
		if numAvail < arg.Cap()/2 {
			/* If less than 1/2 of elements are used, don't sample. Instead search for a
			 * specific offset using linear scan. */
			var i = 0
			var randval = standard.PhpMtRandRange(0, numAvail-1)
			var __ht = arg
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()

				num_key = _p.GetH()
				string_key = _p.GetKey()
				if i == randval {
					if string_key != nil {
						return_value.SetStringCopy(string_key)
						return
					} else {
						return_value.SetLong(num_key)
						return
					}
				}
				i++
			}
		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

		for {
			var randval = standard.PhpMtRandRange(0, arg.GetNNumUsed()-1)
			var bucket *types.Bucket = arg.GetArData()[randval]
			if !(bucket.GetVal().IsUndef()) {
				if bucket.GetKey() != nil {
					return_value.SetStringCopy(bucket.GetKey())
					return
				} else {
					return_value.SetLong(bucket.GetH())
					return
				}
			}

		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

	}
	if numReq <= 0 || numReq > numAvail {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Second argument has to be between 1 and the number of elements in the array")
		return
	}

	/* Make the return value an array only if we need to pass back more than one result. */

	zend.ArrayInitSize(return_value, uint32(numReq))
	if numReq > numAvail>>1 {
		negative_bitset = 1
		numReq = numAvail - numReq
	}
	bitset_len = zend.ZendBitsetLen(numAvail)
	bitset = zend.ZEND_BITSET_ALLOCA(bitset_len, use_heap)
	zend.ZendBitsetClear(bitset, bitset_len)
	i = numReq
	for i != 0 {
		var randval = standard.PhpMtRandRange(0, numAvail-1)
		if zend.ZendBitsetIn(bitset, randval) == 0 {
			zend.ZendBitsetIncl(bitset, randval)
			i--
		}
	}

	/* i = 0; */

	fillScope := types.PackedFillStart(return_value.Array())

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	var __ht = arg
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		if (zend.ZendBitsetIn(bitset, i) ^ negative_bitset) != 0 {
			if string_key != nil {
				fillScope.FillSetStringCopy(string_key)
			} else {
				fillScope.FillSetLong(num_key)
			}
			fillScope.FillNext()
		}
		i++
	}

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	fillScope.FillEnd()
	zend.FreeAlloca(bitset, use_heap)
}
func ZifArraySum(array *types.Array) *types.Zval {
	var num types.Zval
	ret := types.NewZvalLong(0)
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		types.ZVAL_COPY(&num, entry)
		zend.ConvertScalarToNumber(&num)
		zend.FastAddFunction(ret, ret, &num)
	})
	return ret
}
func ZifArrayProduct(array *types.Array) *types.Zval {
	var num types.Zval
	ret := types.NewZvalLong(1)
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		types.ZVAL_COPY(&num, entry)
		zend.ConvertScalarToNumber(&num)
		if num.IsLong() && ret.IsLong() {
			dval := float64(num.Long()) * float64(ret.Long())
			if float64(zend.ZEND_LONG_MIN) <= dval && dval <= float64(zend.ZEND_LONG_MAX) {
				ret.SetLong(ret.Long() * num.Long())
				return
			}
		}
		zend.ConvertToDouble(ret)
		zend.ConvertToDouble(&num)
		ret.SetDouble(ret.Double() * num.Double())
	})
	return ret
}
func ZifArrayReduce(array *types.Array, callback zpp.Callable, _ zpp.Opt, initial *types.Zval) *types.Zval {
	var result types.Zval
	if initial == nil {
		result.SetNull()
	} else {
		result.CopyFrom(initial)
	}

	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
		newResult, ok := callback.Call(&result, value)
		if ok && newResult.IsNotUndef() {
			result.CopyValueFrom(newResult)
		}
	})

	return zend.ZvalZval(&result, true, true)
}
func ZifArrayFilter(array_ *types.Array, _ zpp.Opt, callback zpp.Callable, mode int) *types.Array {
	retArr := types.NewArray(0)
	if array_.Len() == 0 {
		return retArr
	}

	// parse use filter
	var filter func(key types.ArrayKey, value *types.Zval) (keep bool, succ bool)
	if callback != nil {
		filter = func(key types.ArrayKey, value *types.Zval) (bool, bool) {
			var retVal *types.Zval
			var ok bool
			switch mode {
			case ARRAY_FILTER_USE_KEY:
				retVal, ok = callback.Call(key.ToZval())
			case ARRAY_FILTER_USE_BOTH:
				retVal, ok = callback.Call(value, key.ToZval())
			default: // 0
				retVal, ok = callback.Call(value)
			}
			if !ok {
				return false, false
			}
			return zend.ZendIsTrueEx(retVal), true
		}
	}

	// filter
	array_.ForeachIndirectEx(func(key types.ArrayKey, value *types.Zval) bool {
		var keep bool
		if filter != nil {
			var callSucc bool
			keep, callSucc = filter(key, value)
			if !callSucc {
				// 调用 user callback 失败，中断遍历
				return false
			}
		} else {
			keep = zend.ZendIsTrueEx(value)
		}

		if keep {
			retArr.Update(key, value)
		}
		return true
	})
	return retArr
}

func arrayMapSingle(callback zpp.Callable, array *types.Array) *types.Array {
	retArr := types.NewArray(array.Len())
	ok := array.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
		retVal, ok := callback.Call(value)
		if !ok || retVal.IsUndef() {
			// 调用 callback 失败，中断流程
			return false
		}

		retArr.Add(key, retVal)
		return true
	})
	if !ok {
		return nil
	}
	return retArr
}

func arrayMapMulti(callback zpp.Callable, arrays []*types.Array) *types.Array {
	len_ := arrays[0].Len()
	for _, array := range arrays {
		if array.Len() > len_ {
			len_ = array.Len()
		}
	}

	argMatrix := make([][]*types.Zval, len_)
	for _, array := range arrays {
		count := 0
		array.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
			argMatrix[count] = append(argMatrix[count], value)
			count++
			return count < len_
		})
		for i := len_; i < count; i++ {
			argMatrix[i] = append(argMatrix[i], types.NewZvalNull())
		}
	}

	retArr := types.NewArray(len_)
	for _, argColumns := range argMatrix {
		if callback == nil {
			retArr.Append(types.NewZvalArray(types.NewArrayOfZval(argColumns)))
		} else {
			retVal, ok := callback.Call(argColumns...)
			if !ok || retVal.IsUndef() {
				// 调用 callback 失败，中断流程
				return nil
			}
			retArr.Append(retVal)
		}
	}
	return retArr
}

//@zif -c=2,
func ZifArrayMap(callback zpp.Callable, arrays []*types.Zval) *types.Zval {
	b.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(arrays, 1)
	if !ok {
		return types.NewZvalNull()
	}

	var retArr *types.Array
	if len(arrayHts) == 1 {
		retArr = arrayMapSingle(callback, arrayHts[0])
	} else {
		retArr = arrayMapMulti(callback, arrayHts)
	}

	if retArr == nil {
		return types.NewZvalNull()
	}
	return types.NewZvalArray(retArr)
}

//@zif -alias key_exists
func ZifArrayKeyExists(key *types.Zval, array zpp.ArrayOrObject) bool {
	var ht *types.Array
	if array.IsType(types.IS_ARRAY) {
		ht = array.Array()
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. Use isset() or property_exists() instead")
	}
	switch key.GetType() {
	case types.IS_STRING:
		return ht.SymtableExistsInd(key.StringVal())
	case types.IS_LONG:
		return ht.IndexExists(key.Long())
	case types.IS_NULL:
		return ht.KeyExistsIndirect("")
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The first argument should be either a string or an integer")
		return false
	}
}
func ZifArrayChunk(array *types.Array, length int, _ zpp.Opt, preserveKeys bool) *types.Array {
	if length < 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Size parameter expected to be greater than 0")
		return nil
	}

	chunkCount := (array.Len()-1)/length + 1 // ceil(len/length)
	retArr := types.NewArray(chunkCount)

	var currChunk *types.Array = nil
	itemCount := 0
	array.Foreach(func(key types.ArrayKey, entry *types.Zval) {
		if itemCount%length == 0 {
			currChunk = types.NewArray(length)
			retArr.Append(types.NewZvalArray(currChunk))
		}

		if preserveKeys {
			currChunk.Update(key, entry)
		} else {
			currChunk.Append(entry)
		}
		itemCount++
	})
	return retArr
}
func ZifArrayCombine(keys *types.Array, values *types.Array) (*types.Array, bool) {
	if keys.Len() != values.Len() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Both parameters should have an equal number of elements")
		return nil, false
	}
	if keys.Len() == 0 {
		return types.NewArray(0), true
	}

	zvKeys := keys.Values()
	zvValues := values.Values()
	retArr := types.NewArray(keys.Len())
	for i, key := range zvKeys {
		if i > len(zvValues) {
			break
		}
		value := zvValues[i]

		if key.IsLong() {
			retArr.IndexUpdate(key.Long(), value)
		} else {
			strKey := zend.ZvalGetStrVal(key)
			retArr.KeyUpdate(strKey, value)
		}
	}
	return retArr, true
}
