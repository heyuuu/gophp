package array

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"sort"
)

func PhpSplice(in_hash *types.Array, offset zend.ZendLong, length zend.ZendLong, replace *types.Array, removed *types.Array) {
	var out_hash *types.Array
	var num_in zend.ZendLong
	var pos zend.ZendLong
	var p *types.Bucket
	var entry *types.Zval
	var iter_pos = types.ZendHashIteratorsLowerPos(in_hash, 0)

	/* Get number of entries in the input hash */

	num_in = in_hash.Len()

	/* Clamp the offset.. */

	if offset > num_in {
		offset = num_in
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if unsigned(offset+unsigned(length)) > unsigned(num_in) {
		length = num_in - offset
	}

	/* Create and initialize output hash */

	out_hash = types.NewArray(b.Cond(length > 0, num_in-length, 0) + b.CondF1(replace != nil, func() int { return replace.Len() }, 0))

	/* Start at the beginning of the input hash and copy entries to output hash until offset is reached */
	pos = 0
	for ; pos < offset && idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.Bucket(idx)
		if p.GetVal().IsUndef() {
			continue
		}

		/* Get entry and increase reference count */

		entry = p.GetVal()

		/* Update output hash depending on key type */

		if p.GetKey() == nil {
			out_hash.AppendNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* If hash for removed entries exists, go until offset+length and copy the entries to it */

	if removed != nil {
		for ; pos < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.Bucket(idx)
			if p.GetVal().IsUndef() {
				continue
			}
			pos++
			entry = p.GetVal()
			// entry.TryAddRefcount()
			if p.GetKey() == nil {
				removed.AppendNew(entry)
				types.ZendHashDelBucket(in_hash, p)
			} else {
				removed.KeyAddNew(p.GetKey().GetStr(), entry)
				if in_hash == zend.EG__().GetSymbolTable() {
					zend.ZendDeleteGlobalVariable(p.GetKey())
				} else {
					types.ZendHashDelBucket(in_hash, p)
				}
			}
		}
	} else {
		var pos2 = pos
		for ; pos2 < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.Bucket(idx)
			if p.GetVal().IsUndef() {
				continue
			}
			pos2++
			if p.GetKey() != nil && in_hash == zend.EG__().GetSymbolTable() {
				zend.ZendDeleteGlobalVariable(p.GetKey())
			} else {
				types.ZendHashDelBucket(in_hash, p)
			}
		}
	}
	iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos)

	/* If there are entries to insert.. */

	if replace != nil {
		var __ht = replace
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			entry = _z
			// entry.TryAddRefcount()
			out_hash.AppendNew(entry)
			pos++
		}
	}

	/* Copy the remaining input hash entries to the output hash */

	for ; idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.Bucket(idx)
		if p.GetVal().IsUndef() {
			continue
		}
		entry = p.GetVal()
		if p.GetKey() == nil {
			out_hash.AppendNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* replace HashTable data */

	in_hash.Destroy()

	in_hash.SetBy(out_hash)
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
func ZifArrayPop(return_value zpp.Ret, stack zpp.RefArray) *types.Zval {
	if stack.Array().Len() == 0 {
		return types.NewZvalNull()
	}

	var val *types.Zval
	var idx uint32
	var p *types.Bucket

	/* Get the last value and copy it into the return value */
	idx = stack.Array().GetNNumUsed()
	for true {
		if idx == 0 {
			return types.NewZvalNull()
		}
		idx--
		p = stack.Array().Bucket(idx)
		val = p.GetVal()
		if val.IsIndirect() {
			val = val.Indirect()
		}
		if val.IsNotUndef() {
			break
		}
	}
	types.ZVAL_COPY_DEREF(return_value, val)
	if p.GetKey() == nil && stack.Array().GetNNextFreeElement() > 0 && p.GetH() >= zend_ulong(stack.Array().GetNNextFreeElement()-1) {
		stack.Array().SetNNextFreeElement(stack.Array().GetNNextFreeElement() - 1)
	}

	/* Delete the last value */

	if p.GetKey() != nil && stack.Array() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		types.ZendHashDelBucket(stack.Array(), p)
	}
	types.ZendHashInternalPointerReset(stack.Array())
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
	types.ZendHashInternalPointerReset(stack.Array())
}
func ZifArrayUnshift(executeData zpp.Ex, return_value zpp.Ret, stack zpp.RefZval, vars []*types.Zval) {
	var args []*types.Zval = vars
	var new_hash *types.Array
	var argc int
	var i int
	var key *types.String
	var value *types.Zval
	new_hash = types.NewArray(stack.Array().Len() + argc)
	for i = 0; i < argc; i++ {
		args[i].TryAddRefcount()
		new_hash.AppendNew(&args[i])
	}
	var __ht = stack.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			new_hash.KeyAddNew(key.GetStr(), value)
		} else {
			new_hash.AppendNew(value)
		}
	}

	/* replace HashTable data */

	stack.Array().Destroy()

	stack.Array().SetBy(new_hash)

	/* Clean up and return the number of elements in the stack */

	return_value.SetLong(stack.Array().Len())

	/* Clean up and return the number of elements in the stack */
}
func ZifArraySplice(return_value zpp.Ret, arg zpp.RefArray, offset int, _ zpp.Opt, length_ *int, replacement *types.Zval) {
	var array *types.Zval
	var repl_array *types.Zval = replacement
	var rem_hash *types.Array = nil
	var numIn int = arg.Array().Len()

	length := b.Option(length_, numIn)
	if replacement != nil {
		/* Make sure the last argument, if passed, is an array */
		zend.ConvertToArrayEx(repl_array)
	}

	/* Don't create the array of removed elements if it's not going
	 * to be used; e.g. only removing and/or replacing elements */

	if zend.USED_RET() {
		var size = length

		/* Clamp the offset.. */
		if offset < 0 {
			offset = offset + numIn
			if offset < 0 {
				offset = 0
			}
		} else if offset > numIn {
			offset = numIn
		}

		/* ..and the length */
		if length < 0 {
			size = numIn - offset + length
		} else if offset+length > numIn {
			size = numIn - offset
		}

		/* Initialize return value */
		zend.ArrayInitSize(return_value, b.CondF1(size > 0, func() uint32 { return uint32(size) }, 0))
		rem_hash = return_value.Array()
	}

	/* Perform splice */

	PhpSplice(array.Array(), offset, length, b.CondF1(repl_array != nil, func() *types.Array { return repl_array.Array() }, nil), rem_hash)
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
func PhpArrayMerge(dest *types.Array, src *types.Array) int {
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
func ZifArrayKeys(array zpp.ArrayHt, _ zpp.Opt, searchValue *types.Zval, strict bool) *types.Array {
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
func ZifArrayReverse(executeData zpp.Ex, return_value zpp.Ret, input *types.Zval, _ zpp.Opt, preserveKeys *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var preserve_keys = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			input = fp.ParseArray()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, input.Array().Len())
	if input.Array().IsPacked() && preserve_keys == 0 {

		fillScope := types.PackedFillStart(return_value.Array())
		var __ht = input.Array()
		for _, _p := range __ht.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			entry = _z
			if entry.IsReference() && entry.GetRefcount() == 1 {
				entry = types.Z_REFVAL_P(entry)
			}
			// entry.TryAddRefcount()
			fillScope.FillSet(entry)
			fillScope.FillNext()
		}
		fillScope.FillEnd()
	} else {
		var __ht = input.Array()
		for _, _p := range __ht.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			num_key = _p.GetH()
			string_key = _p.GetKey()
			entry = _z
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
}
func ZifArrayPad(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, padSize *types.Zval, padValue *types.Zval) {
	var input *types.Zval
	var pad_value *types.Zval
	var pad_size zend.ZendLong
	var pad_size_abs zend.ZendLong
	var input_size zend.ZendLong
	var num_pads zend.ZendLong
	var i zend.ZendLong
	var key *types.String
	var value *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			input = fp.ParseArray()
			pad_size = fp.ParseLong()
			pad_value = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Do some initial calculations */

	input_size = input.Array().Len()
	pad_size_abs = zend.ZEND_ABS(pad_size)
	if pad_size_abs < 0 || pad_size_abs-input_size > int64(1048576) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return_value.SetFalse()
		return
	}
	if input_size >= pad_size_abs {

		/* Copy the original array */

		types.ZVAL_COPY(return_value, input)
		return
	}
	num_pads = pad_size_abs - input_size
	//if pad_value.IsRefcounted() {
	//	pad_value.RefCounted().AddRefcountEx(num_pads)
	//}
	zend.ArrayInitSize(return_value, pad_size_abs)
	if input.Array().IsPacked() {

		if pad_size < 0 {
			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < num_pads; i++ {
				fillScope.FillSet(pad_value)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		}
		fillScope := types.PackedFillStart(return_value.Array())
		var __ht = input.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()

			value = _z
			// value.TryAddRefcount()
			fillScope.FillSet(value)
			fillScope.FillNext()
		}
		fillScope.FillEnd()
		if pad_size > 0 {
			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < num_pads; i++ {
				fillScope.FillSet(pad_value)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		}
	} else {
		if pad_size < 0 {
			for i = 0; i < num_pads; i++ {
				return_value.Array().AppendNew(pad_value)
			}
		}
		var __ht = input.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			key = _p.GetKey()
			value = _z
			// value.TryAddRefcount()
			if key != nil {
				return_value.Array().KeyAddNew(key.GetStr(), value)
			} else {
				return_value.Array().AppendNew(value)
			}
		}
		if pad_size > 0 {
			for i = 0; i < num_pads; i++ {
				return_value.Array().AppendNew(pad_value)
			}
		}
	}
}
func ZifArrayFlip(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var data types.Zval
	var num_idx zend.ZendUlong
	var str_idx *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			array = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, array.Array().Len())
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_idx = _p.GetH()
		str_idx = _p.GetKey()
		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.Array().IndexUpdate(entry.Long(), &data)
		} else if entry.IsString() {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.Array().SymtableUpdate(entry.String().GetStr(), &data)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	}
}
func ZifArrayChangeKeyCase(executeData zpp.Ex, return_value zpp.Ret, input *types.Zval, _ zpp.Opt, case_ *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var string_key *types.String
	var new_key *types.String
	var num_key zend.ZendUlong
	var change_to_upper = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			array = fp.ParseArray()
			fp.StartOptional()
			change_to_upper = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, array.Array().Len())
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		if string_key == nil {
			entry = return_value.Array().IndexUpdate(num_key, entry)
		} else {
			if change_to_upper != 0 {
				new_key = str.PhpStringToupper(string_key)
			} else {
				new_key = str.PhpStringTolower(string_key)
			}
			entry = return_value.Array().KeyUpdate(new_key.GetStr(), entry)
			// types.ZendStringReleaseEx(new_key, 0)
		}
		zend.ZvalAddRef(entry)
	}
}
func ZifArrayUnique(arg *types.Array, _ zpp.Opt, flags *int) *types.Array {
	var sortType = b.Option(flags, PHP_SORT_STRING)

	if sortType == PHP_SORT_STRING {
		existValues := make(map[string]bool, arg.Len())
		retArr := types.NewArray(arg.Len())
		arg.ForeachIndirect(func(key types.ArrayKey, val *types.Zval) {
			var strVal string = zend.ZvalGetStrVal(val)
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
func ZvalUserCompare(a *types.Zval, b *types.Zval) int {
	var args []types.Zval
	var retval types.Zval
	types.ZVAL_COPY_VALUE(&args[0], a)
	types.ZVAL_COPY_VALUE(&args[1], b)
	standard.BG__().user_compare_fci.param_count = 2
	standard.BG__().user_compare_fci.params = args
	standard.BG__().user_compare_fci.retval = &retval
	standard.BG__().user_compare_fci.no_separation = 0
	if zend.ZendCallFunction(&(standard.BG__().user_compare_fci), &(standard.BG__().user_compare_fci_cache)) == types.SUCCESS && retval.IsNotUndef() {
		var ret = zend.ZvalGetLong(&retval)
		// zend.ZvalPtrDtor(&retval)
		return zend.ZEND_NORMALIZE_BOOL(ret)
	} else {
		return 0
	}
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
func ZifArrayDiff(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, arrays []*types.Zval) {
	var args []*types.Zval
	var argc int
	var i int
	var num uint32
	var exclude *types.Array
	var value *types.Zval
	var str *types.String
	var tmp_str *types.String
	var key *types.String
	var idx zend.ZendLong
	var dummy types.Zval
	if executeData.NumArgs() < 2 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "at least 2 parameters are required, %d given", executeData.NumArgs())
		return
	}
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
	if args[0].GetType() != types.IS_ARRAY {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter 1 to be an array, %s given", types.ZendZvalTypeName(&args[0]))
		return_value.SetNull()
		return
	}
	num = args[0].Array().Len()
	if num == 0 {
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
		}
		return_value.SetEmptyArray()
		return
	} else if num == 1 {
		var found = 0
		var search_str *types.String
		var tmp_search_str *types.String
		value = nil
		var __ht *types.Array = args[0].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value = _z
			break
		}
		if value == nil {
			for i = 1; i < argc; i++ {
				if args[i].GetType() != types.IS_ARRAY {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
					return_value.SetNull()
					return
				}
			}
			return_value.SetEmptyArray()
			return
		}
		search_str = zend.ZvalGetTmpString(value, &tmp_search_str)
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
			if found == 0 {
				var __ht *types.Array = args[i].Array()
				for _, _p := range __ht.ForeachData() {
					var _z = _p.GetVal()
					if _z.IsIndirect() {
						_z = _z.Indirect()
						if _z.IsUndef() {
							continue
						}
					}
					value = _z
					str = zend.ZvalGetTmpString(value, &tmp_str)
					if search_str.GetStr() == str.GetStr() {
						// zend.ZendTmpStringRelease(tmp_str)
						found = 1
						break
					}
					// zend.ZendTmpStringRelease(tmp_str)
				}
			}
		}
		// zend.ZendTmpStringRelease(tmp_search_str)
		if found != 0 {
			return_value.SetEmptyArray()
		} else {
			types.ZVAL_COPY(return_value, &args[0])
		}
		return
	}

	/* count number of elements */

	num = 0
	for i = 1; i < argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
			return_value.SetNull()
			return
		}
		num += args[i].Array().Len()
	}
	if num == 0 {
		types.ZVAL_COPY(return_value, &args[0])
		return
	}
	dummy.SetNull()

	/* create exclude map */

	exclude = types.NewArray(num)
	for i = 1; i < argc; i++ {
		var __ht *types.Array = args[i].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value = _z
			str = zend.ZvalGetTmpString(value, &tmp_str)
			exclude.KeyAdd(str.GetStr(), &dummy)
			// zend.ZendTmpStringRelease(tmp_str)
		}
	}

	/* copy all elements of first array that are not in exclude set */

	zend.ArrayInitSize(return_value, args[0].Array().Len())
	var __ht *types.Array = args[0].Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		idx = _p.GetH()
		key = _p.GetKey()
		value = _z
		str = zend.ZvalGetTmpString(value, &tmp_str)
		if !exclude.KeyExists(str.GetStr()) {
			if key != nil {
				value = return_value.Array().KeyAddNew(key.GetStr(), value)
			} else {
				value = return_value.Array().IndexAddNew(idx, value)
			}
			zend.ZvalAddRef(value)
		}
		// zend.ZendTmpStringRelease(tmp_str)
	}
	exclude.Destroy()
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
	var parseState [2]int = [...]int{0, 0}

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
func ZifArraySum(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
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
	return_value.SetLong(0)
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		zend.FastAddFunction(return_value, return_value, &entry_n)
	}
}
func ZifArrayProduct(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
	var dval float64
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
	return_value.SetLong(1)
	if !(input.Array().Len()) {
		return
	}
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		if entry_n.IsType(types.IS_LONG) && return_value.IsType(types.IS_LONG) {
			dval = float64(return_value.Long() * float64(entry_n.Long()))
			if float64(zend.ZEND_LONG_MIN <= dval && dval <= float64(zend.ZEND_LONG_MAX)) {
				return_value.SetLong(return_value.Long() * entry_n.Long())
				continue
			}
		}
		zend.ConvertToDouble(return_value)
		zend.ConvertToDouble(&entry_n)
		return_value.SetDouble(return_value.Double() * entry_n.Double())
	}
}
func ZifArrayReduce(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, callback *types.Zval, _ zpp.Opt, initial *types.Zval) {
	var input *types.Zval
	var args []types.Zval
	var operand *types.Zval
	var result types.Zval
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var initial *types.Zval = nil
	var htbl *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			input = fp.ParseArray()
			fp.ParseFunc(&fci, &fci_cache)
			fp.StartOptional()
			initial = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() > 2 {
		types.ZVAL_COPY(&result, initial)
	} else {
		result.SetNull()
	}

	/* (zval **)input points to an element of argument stack
	 * the base pointer of which is subject to change.
	 * thus we need to keep the pointer to the hashtable for safety */

	htbl = input.Array()
	if htbl.Len() == 0 {
		types.ZVAL_COPY_VALUE(return_value, &result)
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	fci.SetRetval(&retval)
	fci.SetParamCount(2)
	fci.SetNoSeparation(0)
	var __ht = htbl
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		operand = _z
		types.ZVAL_COPY_VALUE(&args[0], &result)
		types.ZVAL_COPY(&args[1], operand)
		fci.SetParams(args)
		if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
			// zend.ZvalPtrDtor(&args[1])
			// zend.ZvalPtrDtor(&args[0])
			types.ZVAL_COPY_VALUE(&result, &retval)
		} else {
			// zend.ZvalPtrDtor(&args[1])
			// zend.ZvalPtrDtor(&args[0])
			return
		}
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
	zend.ZVAL_ZVAL(return_value, &result, 1, 1)
}
func ZifArrayFilter(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, _ zpp.Opt, callback *types.Zval, useKeys *types.Zval) {
	var array *types.Zval
	var operand *types.Zval
	var key *types.Zval
	var args []types.Zval
	var retval types.Zval
	var have_callback = 0
	var use_type = 0
	var string_key *types.String
	var fci = zend.EmptyFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var num_key zend.ZendUlong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			array = fp.ParseArray()
			fp.StartOptional()
			fp.ParseFunc(&fci, &fci_cache)
			use_type = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInit(return_value)
	if array.Array().Len() == 0 {
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	if executeData.NumArgs() > 1 {
		have_callback = 1
		fci.SetNoSeparation(0)
		fci.SetRetval(&retval)
		if use_type == ARRAY_FILTER_USE_BOTH {
			fci.SetParamCount(2)
			key = &args[1]
		} else {
			fci.SetParamCount(1)
			key = &args[0]
		}
	}
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		operand = _z
		if have_callback != 0 {
			if use_type != 0 {

				/* Set up the key */

				if string_key == nil {
					key.SetLong(num_key)
				} else {
					key.SetStringCopy(string_key)
				}

				/* Set up the key */

			}
			if use_type != ARRAY_FILTER_USE_KEY {
				types.ZVAL_COPY(&args[0], operand)
			}
			fci.SetParams(args)
			if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS {
				var retval_true int
				// zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					// zend.ZvalPtrDtor(&args[1])
				}
				retval_true = zend.ZendIsTrue(&retval)
				// zend.ZvalPtrDtor(&retval)
				if retval_true == 0 {
					continue
				}
			} else {
				// zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					// zend.ZvalPtrDtor(&args[1])
				}
				return
			}
		} else if zend.ZendIsTrue(operand) == 0 {
			continue
		}
		if string_key != nil {
			operand = return_value.Array().KeyUpdate(string_key.GetStr(), operand)
		} else {
			operand = return_value.Array().IndexUpdate(num_key, operand)
		}
		zend.ZvalAddRef(operand)
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
}
func ZifArrayMap(executeData zpp.Ex, return_value zpp.Ret, callback *types.Zval, arrays []*types.Zval) {
	var arrays *types.Zval = nil
	var n_arrays = 0
	var result types.Zval
	var fci = zend.EmptyFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var i int
	var k uint32
	var maxlen uint32 = 0
	for {
		var _flags = 0
		var _min_num_args = 2
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.ParseFuncEx(&fci, &fci_cache, true, false)
			arrays, n_arrays = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetNull()
	if n_arrays == 1 {
		var num_key zend.ZendUlong
		var str_key *types.String
		var zv *types.Zval
		var arg types.Zval
		var ret int
		if arrays[0].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter 2 to be an array, %s given", types.ZendZvalTypeName(&arrays[0]))
			return
		}
		maxlen = arrays[0].Array().Len()

		/* Short-circuit: if no callback and only one array, just return it. */

		if !(zend.ZEND_FCI_INITIALIZED(fci)) || maxlen == 0 {
			types.ZVAL_COPY(return_value, &arrays[0])
			zend.ZendReleaseFcallInfoCache(&fci_cache)
			return
		}
		zend.ArrayInitSize(return_value, maxlen)
		var __ht = arrays[0].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			fci.SetRetval(&result)
			fci.SetParamCount(1)
			fci.SetParams(&arg)
			fci.SetNoSeparation(0)
			types.ZVAL_COPY(&arg, zv)
			ret = zend.ZendCallFunction(&fci, &fci_cache)
			// zend.IZvalPtrDtor(&arg)
			if ret != types.SUCCESS || result.IsUndef() {
				return_value.Array().DestroyEx()
				return_value.SetNull()
				return
			}
			if str_key != nil {
				zend._zendHashAppend(return_value.Array(), str_key, &result)
			} else {
				return_value.Array().IndexAddNew(num_key, &result)
			}
		}
		zend.ZendReleaseFcallInfoCache(&fci_cache)
	} else {
		var array_pos = (*types.ArrayPosition)(zend.Ecalloc(n_arrays, b.SizeOf("ArrayPosition")))
		for i = 0; i < n_arrays; i++ {
			if arrays[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+2, types.ZendZvalTypeName(&arrays[i]))
				zend.Efree(array_pos)
				return
			}
			if arrays[i].Array().Len() > maxlen {
				maxlen = arrays[i].Array().Len()
			}
		}
		zend.ArrayInitSize(return_value, maxlen)
		if !(zend.ZEND_FCI_INITIALIZED(fci)) {
			var zv types.Zval

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {

				/* If no callback, the result will be an array, consisting of current
				 * entries from all arrays. */

				zend.ArrayInitSize(&result, n_arrays)
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].Array().GetNNumUsed() {
							zv.SetNull()
							break
						} else if arrays[i].Array().GetArData()[pos].GetVal().IsNotUndef() {
							types.ZVAL_COPY(&zv, arrays[i].Array().GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
					result.Array().AppendNew(&zv)
				}
				return_value.Array().AppendNew(&result)
			}

			/* We iterate through all the arrays at once. */

		} else {
			var params = (*types.Zval)(zend.SafeEmalloc(n_arrays, b.SizeOf("zval"), 0))

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].Array().GetNNumUsed() {
							params[i].SetNull()
							break
						} else if arrays[i].Array().GetArData()[pos].GetVal().IsNotUndef() {
							types.ZVAL_COPY(&params[i], arrays[i].Array().GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
				}
				fci.SetRetval(&result)
				fci.SetParamCount(n_arrays)
				fci.SetParams(params)
				fci.SetNoSeparation(0)
				if zend.ZendCallFunction(&fci, &fci_cache) != types.SUCCESS || result.IsUndef() {
					zend.Efree(array_pos)
					return_value.Array().DestroyEx()
					for i = 0; i < n_arrays; i++ {
						// zend.ZvalPtrDtor(&params[i])
					}
					zend.Efree(params)
					return_value.SetNull()
					return
				} else {
					for i = 0; i < n_arrays; i++ {
						// zend.ZvalPtrDtor(&params[i])
					}
				}
				return_value.Array().AppendNew(&result)
			}
			zend.Efree(params)
			zend.ZendReleaseFcallInfoCache(&fci_cache)
		}
		zend.Efree(array_pos)
	}
}

//@zif -alias key_exists
func ZifArrayKeyExists(executeData zpp.Ex, return_value zpp.Ret, key *types.Zval, search *types.Zval) {
	var key *types.Zval
	var array *types.Zval
	var ht *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			key = fp.ParseZval()
			array = fp.ParseArrayOrObject()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if array.IsType(types.IS_ARRAY) {
		ht = array.Array()
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
	}
	switch key.GetType() {
	case types.IS_STRING:
		return_value.SetBool(ht.SymtableExistsInd(key.String().GetStr()))
	case types.IS_LONG:
		return_value.SetBool(ht.IndexExists(key.Long()))
	case types.IS_NULL:
		return_value.SetBool(ht.KeyExistsIndirect(types.NewString("").GetStr()))
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The first argument should be either a string or an integer")
		return_value.SetFalse()
	}
	if array.GetType() != types.IS_ARRAY {
		zend.ZendReleaseProperties(ht)
	}
}
func ZifArrayChunk(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, size *types.Zval, _ zpp.Opt, preserveKeys *types.Zval) {
	var num_in int
	var size zend.ZendLong
	var current = 0
	var str_key *types.String
	var num_key zend.ZendUlong
	var preserve_keys = 0
	var input *types.Zval = nil
	var chunk types.Zval
	var entry *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			input = fp.ParseArray()
			size = fp.ParseLong()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Do bounds checking for size parameter. */

	if size < 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Size parameter expected to be greater than 0")
		return
	}
	num_in = input.Array().Len()
	if size > num_in {
		if num_in > 0 {
			size = num_in
		} else {
			size = 1
		}
	}
	zend.ArrayInitSize(return_value, uint32((num_in-1)/size+1))
	chunk.SetUndef()
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		str_key = _p.GetKey()
		entry = _z

		/* If new chunk, create and initialize it. */

		if chunk.IsUndef() {
			zend.ArrayInitSize(&chunk, uint32(size))
		}

		/* Add entry to the chunk, preserving keys if necessary. */

		if preserve_keys != 0 {
			if str_key != nil {
				entry = chunk.Array().KeyUpdate(str_key.GetStr(), entry)
			} else {
				entry = chunk.Array().IndexUpdate(num_key, entry)
			}
		} else {
			entry = chunk.Array().Append(entry)
		}
		zend.ZvalAddRef(entry)

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

		if b.PreInc(&current)%size == 0 {
			zend.AddNextIndexZval(return_value, &chunk)
			chunk.SetUndef()
		}

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

	}

	/* Add the final chunk if there is one. */

	if chunk.IsNotUndef() {
		zend.AddNextIndexZval(return_value, &chunk)
	}

	/* Add the final chunk if there is one. */
}
func ZifArrayCombine(executeData zpp.Ex, return_value zpp.Ret, keys *types.Zval, values *types.Zval) {
	var values *types.Array
	var keys *types.Array
	var pos_values uint32 = 0
	var entry_keys *types.Zval
	var entry_values *types.Zval
	var num_keys int
	var num_values int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			keys = fp.ParseArrayHt()
			values = fp.ParseArrayHt()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	num_keys = keys.Len()
	num_values = values.Len()
	if num_keys != num_values {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Both parameters should have an equal number of elements")
		return_value.SetFalse()
		return
	}
	if num_keys == 0 {
		return_value.SetEmptyArray()
		return
	}
	zend.ArrayInitSize(return_value, num_keys)
	var __ht = keys
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry_keys = _z
		for true {
			if pos_values >= values.GetNNumUsed() {
				break
			} else if values.GetArData()[pos_values].GetVal().IsNotUndef() {
				entry_values = values.GetArData()[pos_values].GetVal()
				if entry_keys.IsType(types.IS_LONG) {
					entry_values = return_value.Array().IndexUpdate(entry_keys.Long(), entry_values)
				} else {
					var tmp_key *types.String
					var key = zend.ZvalGetTmpString(entry_keys, &tmp_key)
					entry_values = return_value.Array().SymtableUpdate(key.GetStr(), entry_values)
					// zend.ZendTmpStringRelease(tmp_key)
				}
				zend.ZvalAddRef(entry_values)
				pos_values++
				break
			}
			pos_values++
		}
	}
}
