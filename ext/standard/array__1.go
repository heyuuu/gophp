// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

/* }}} */

func ZifArrayRand(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var num_req zend.ZendLong = 1
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var i int
	var num_avail int
	var bitset zend.ZendBitset
	var negative_bitset int = 0
	var bitset_len uint32
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &num_req, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	num_avail = input.value.arr.nNumOfElements
	if num_avail == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Array is empty")
		return
	}
	if num_req == 1 {
		var ht *zend.HashTable = input.value.arr
		if uint32(num_avail < ht.nNumUsed-(ht.nNumUsed>>1)) != 0 {

			/* If less than 1/2 of elements are used, don't sample. Instead search for a
			 * specific offset using linear scan. */

			var i zend.ZendLong = 0
			var randval zend.ZendLong = PhpMtRandRange(0, num_avail-1)
			for {
				var __ht *zend.HashTable = input.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					num_key = _p.h
					string_key = _p.key
					if i == randval {
						if string_key != nil {
							var __z *zend.Zval = return_value
							var __s *zend.ZendString = string_key
							__z.value.str = __s
							if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
								__z.u1.type_info = 6
							} else {
								zend.ZendGcAddref(&__s.gc)
								__z.u1.type_info = 6 | 1<<0<<8
							}
							return
						} else {
							var __z *zend.Zval = return_value
							__z.value.lval = num_key
							__z.u1.type_info = 4
							return
						}
					}
					i++
				}
				break
			}
		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

		for {
			var randval zend.ZendLong = PhpMtRandRange(0, ht.nNumUsed-1)
			var bucket *zend.Bucket = &ht.arData[randval]
			if bucket.val.u1.v.type_ != 0 {
				if bucket.key != nil {
					var __z *zend.Zval = return_value
					var __s *zend.ZendString = bucket.key
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
					return
				} else {
					var __z *zend.Zval = return_value
					__z.value.lval = bucket.h
					__z.u1.type_info = 4
					return
				}
			}

		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

	}
	if num_req <= 0 || num_req > num_avail {
		core.PhpErrorDocref(nil, 1<<1, "Second argument has to be between 1 and the number of elements in the array")
		return
	}

	/* Make the return value an array only if we need to pass back more than one result. */

	var __arr *zend.ZendArray = zend._zendNewArray(uint32(num_req))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if num_req > num_avail>>1 {
		negative_bitset = 1
		num_req = num_avail - num_req
	}
	bitset_len = zend.ZendBitsetLen(num_avail)
	bitset = zend.ZendBitset(zend._emalloc(bitset_len * g.SizeOf("zend_ulong")))
	zend.ZendBitsetClear(bitset, bitset_len)
	i = num_req
	for i != 0 {
		var randval zend.ZendLong = PhpMtRandRange(0, num_avail-1)
		if zend.ZendBitsetIn(bitset, randval) == 0 {
			zend.ZendBitsetIncl(bitset, randval)
			i--
		}
	}

	/* i = 0; */

	zend.ZendHashRealInitPacked(return_value.value.arr)
	for {
		var __fill_ht *zend.HashTable = return_value.value.arr
		var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
		var __fill_idx uint32 = __fill_ht.nNumUsed
		assert((__fill_ht.u.flags & 1 << 2) != 0)

		/* We can't use zend_hash_index_find()
		 * because the array may have string keys or gaps. */

		for {
			var __ht *zend.HashTable = input.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				if (zend.ZendBitsetIn(bitset, i) ^ negative_bitset) != 0 {
					if string_key != nil {
						var __z *zend.Zval = &__fill_bkt.val
						var __s *zend.ZendString = string_key
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							zend.ZendGcAddref(&__s.gc)
							__z.u1.type_info = 6 | 1<<0<<8
						}
					} else {
						var __z *zend.Zval = &__fill_bkt.val
						__z.value.lval = num_key
						__z.u1.type_info = 4
					}
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
				}
				i++
			}
			break
		}

		/* We can't use zend_hash_index_find()
		 * because the array may have string keys or gaps. */

		__fill_ht.nNumUsed = __fill_idx
		__fill_ht.nNumOfElements = __fill_idx
		__fill_ht.nNextFreeElement = __fill_idx
		__fill_ht.nInternalPointer = 0
		break
	}
	zend._efree(bitset)
}

/* }}} */

func ZifArraySum(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var entry_n zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	__z.value.lval = 0
	__z.u1.type_info = 4
	for {
		var __ht *zend.HashTable = input.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			entry = _z
			if entry.u1.v.type_ == 7 || entry.u1.v.type_ == 8 {
				continue
			}
			var _z1 *zend.Zval = &entry_n
			var _z2 *zend.Zval = entry
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ConvertScalarToNumber(&entry_n)
			zend.FastAddFunction(return_value, return_value, &entry_n)
		}
		break
	}
}

/* }}} */

func ZifArrayProduct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var entry_n zend.Zval
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	__z.value.lval = 1
	__z.u1.type_info = 4
	if input.value.arr.nNumOfElements == 0 {
		return
	}
	for {
		var __ht *zend.HashTable = input.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			entry = _z
			if entry.u1.v.type_ == 7 || entry.u1.v.type_ == 8 {
				continue
			}
			var _z1 *zend.Zval = &entry_n
			var _z2 *zend.Zval = entry
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ConvertScalarToNumber(&entry_n)
			if entry_n.u1.v.type_ == 4 && return_value.u1.v.type_ == 4 {
				dval = float64(return_value.value.lval * float64(entry_n.value.lval))
				if float64(INT64_MIN <= dval && dval <= float64(INT64_MAX)) {
					return_value.value.lval *= entry_n.value.lval
					continue
				}
			}
			zend.ConvertToDouble(return_value)
			zend.ConvertToDouble(&entry_n)
			return_value.value.dval *= entry_n.value.dval
		}
		break
	}
}

/* }}} */

func ZifArrayReduce(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var args []zend.Zval
	var operand *zend.Zval
	var result zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var initial *zend.Zval = nil
	var htbl *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &initial, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if execute_data.This.u2.num_args > 2 {
		var _z1 *zend.Zval = &result
		var _z2 *zend.Zval = initial
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	} else {
		&result.u1.type_info = 1
	}

	/* (zval **)input points to an element of argument stack
	 * the base pointer of which is subject to change.
	 * thus we need to keep the pointer to the hashtable for safety */

	htbl = input.value.arr
	if htbl.nNumOfElements == 0 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &result
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	fci.retval = &retval
	fci.param_count = 2
	fci.no_separation = 0
	for {
		var __ht *zend.HashTable = htbl
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			operand = _z
			var _z1 *zend.Zval = &args[0]
			var _z2 *zend.Zval = &result
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			var _z1 *zend.Zval = &args[1]
			var _z2 *zend.Zval = operand
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			fci.params = args
			if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.u1.v.type_ != 0 {
				zend.ZvalPtrDtor(&args[1])
				zend.ZvalPtrDtor(&args[0])
				var _z1 *zend.Zval = &result
				var _z2 *zend.Zval = &retval
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			} else {
				zend.ZvalPtrDtor(&args[1])
				zend.ZvalPtrDtor(&args[0])
				return
			}
		}
		break
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
	var __z *zend.Zval = return_value
	var __zv *zend.Zval = &result
	if __zv.u1.v.type_ != 10 {
		var _z1 *zend.Zval = __z
		var _z2 *zend.Zval = __zv
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		var _z1 *zend.Zval = __z
		var _z2 *zend.Zval = &(*__zv).value.ref.val
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		zend.ZvalPtrDtor(__zv)
	}
}

/* }}} */

func ZifArrayFilter(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var operand *zend.Zval
	var key *zend.Zval
	var args []zend.Zval
	var retval zend.Zval
	var have_callback zend.ZendBool = 0
	var use_type zend.ZendLong = 0
	var string_key *zend.ZendString
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var num_key zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &use_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if array.value.arr.nNumOfElements == 0 {
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	if execute_data.This.u2.num_args > 1 {
		have_callback = 1
		fci.no_separation = 0
		fci.retval = &retval
		if use_type == 1 {
			fci.param_count = 2
			key = &args[1]
		} else {
			fci.param_count = 1
			key = &args[0]
		}
	}
	for {
		var __ht *zend.HashTable = array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			operand = _z
			if have_callback != 0 {
				if use_type != 0 {

					/* Set up the key */

					if string_key == nil {
						var __z *zend.Zval = key
						__z.value.lval = num_key
						__z.u1.type_info = 4
					} else {
						var __z *zend.Zval = key
						var __s *zend.ZendString = string_key
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							zend.ZendGcAddref(&__s.gc)
							__z.u1.type_info = 6 | 1<<0<<8
						}
					}

					/* Set up the key */

				}
				if use_type != 2 {
					var _z1 *zend.Zval = &args[0]
					var _z2 *zend.Zval = operand
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					if (_t & 0xff00) != 0 {
						zend.ZendGcAddref(&_gc.gc)
					}
				}
				fci.params = args
				if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS {
					var retval_true int
					zend.ZvalPtrDtor(&args[0])
					if use_type == 1 {
						zend.ZvalPtrDtor(&args[1])
					}
					retval_true = zend.ZendIsTrue(&retval)
					zend.ZvalPtrDtor(&retval)
					if retval_true == 0 {
						continue
					}
				} else {
					zend.ZvalPtrDtor(&args[0])
					if use_type == 1 {
						zend.ZvalPtrDtor(&args[1])
					}
					return
				}
			} else if zend.ZendIsTrue(operand) == 0 {
				continue
			}
			if string_key != nil {
				operand = zend.ZendHashUpdate(return_value.value.arr, string_key, operand)
			} else {
				operand = zend.ZendHashIndexUpdate(return_value.value.arr, num_key, operand)
			}
			zend.ZvalAddRef(operand)
		}
		break
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
}

/* }}} */

func ZifArrayMap(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arrays *zend.Zval = nil
	var n_arrays int = 0
	var result zend.Zval
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var i int
	var k uint32
	var maxlen uint32 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 1, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				arrays = _real_arg + 1
				n_arrays = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				arrays = nil
				n_arrays = 0
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	return_value.u1.type_info = 1
	if n_arrays == 1 {
		var num_key zend.ZendUlong
		var str_key *zend.ZendString
		var zv *zend.Zval
		var arg zend.Zval
		var ret int
		if arrays[0].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter 2 to be an array, %s given", zend.ZendZvalTypeName(&arrays[0]))
			return
		}
		maxlen = arrays[0].value.arr.nNumOfElements

		/* Short-circuit: if no callback and only one array, just return it. */

		if fci.size == 0 || maxlen == 0 {
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = &arrays[0]
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ZendReleaseFcallInfoCache(&fci_cache)
			return
		}
		var __arr *zend.ZendArray = zend._zendNewArray(maxlen)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.ZendHashRealInit(return_value.value.arr, arrays[0].value.arr.u.flags&1<<2)
		for {
			var __ht *zend.HashTable = arrays[0].value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				zv = _z
				fci.retval = &result
				fci.param_count = 1
				fci.params = &arg
				fci.no_separation = 0
				var _z1 *zend.Zval = &arg
				var _z2 *zend.Zval = zv
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				if (_t & 0xff00) != 0 {
					zend.ZendGcAddref(&_gc.gc)
				}
				ret = zend.ZendCallFunction(&fci, &fci_cache)
				zend.IZvalPtrDtor(&arg)
				if ret != zend.SUCCESS || result.u1.v.type_ == 0 {
					zend.ZendArrayDestroy(return_value.value.arr)
					return_value.u1.type_info = 1
					return
				}
				if str_key != nil {
					zend._zendHashAppend(return_value.value.arr, str_key, &result)
				} else {
					zend.ZendHashIndexAddNew(return_value.value.arr, num_key, &result)
				}
			}
			break
		}
		zend.ZendReleaseFcallInfoCache(&fci_cache)
	} else {
		var array_pos *uint32 = (*zend.HashPosition)(zend._ecalloc(n_arrays, g.SizeOf("HashPosition")))
		for i = 0; i < n_arrays; i++ {
			if arrays[i].u1.v.type_ != 7 {
				core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+2, zend.ZendZvalTypeName(&arrays[i]))
				zend._efree(array_pos)
				return
			}
			if arrays[i].value.arr.nNumOfElements > maxlen {
				maxlen = arrays[i].value.arr.nNumOfElements
			}
		}
		var __arr *zend.ZendArray = zend._zendNewArray(maxlen)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if fci.size == 0 {
			var zv zend.Zval

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {

				/* If no callback, the result will be an array, consisting of current
				 * entries from all arrays. */

				var __arr *zend.ZendArray = zend._zendNewArray(n_arrays)
				var __z *zend.Zval = &result
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].value.arr.nNumUsed {
							&zv.u1.type_info = 1
							break
						} else if arrays[i].value.arr.arData[pos].val.u1.v.type_ != 0 {
							var _z1 *zend.Zval = &zv
							var _z2 *zend.Zval = &arrays[i].value.arr.arData[pos].val
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							if (_t & 0xff00) != 0 {
								zend.ZendGcAddref(&_gc.gc)
							}
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
					zend.ZendHashNextIndexInsertNew(result.value.arr, &zv)
				}
				zend.ZendHashNextIndexInsertNew(return_value.value.arr, &result)
			}

			/* We iterate through all the arrays at once. */

		} else {
			var params *zend.Zval = (*zend.Zval)(zend._safeEmalloc(n_arrays, g.SizeOf("zval"), 0))

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].value.arr.nNumUsed {
							&params[i].u1.type_info = 1
							break
						} else if arrays[i].value.arr.arData[pos].val.u1.v.type_ != 0 {
							var _z1 *zend.Zval = &params[i]
							var _z2 *zend.Zval = &arrays[i].value.arr.arData[pos].val
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							if (_t & 0xff00) != 0 {
								zend.ZendGcAddref(&_gc.gc)
							}
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
				}
				fci.retval = &result
				fci.param_count = n_arrays
				fci.params = params
				fci.no_separation = 0
				if zend.ZendCallFunction(&fci, &fci_cache) != zend.SUCCESS || result.u1.v.type_ == 0 {
					zend._efree(array_pos)
					zend.ZendArrayDestroy(return_value.value.arr)
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
					zend._efree(params)
					return_value.u1.type_info = 1
					return
				} else {
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
				}
				zend.ZendHashNextIndexInsertNew(return_value.value.arr, &result)
			}
			zend._efree(params)
			zend.ZendReleaseFcallInfoCache(&fci_cache)
		}
		zend._efree(array_pos)
	}
}

/* }}} */

func ZifArrayKeyExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var key *zend.Zval
	var array *zend.Zval
	var ht *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &key, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if array.u1.v.type_ == 7 {
		ht = array.value.arr
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, 1<<13, "Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
	}
	switch key.u1.v.type_ {
	case 6:
		if zend.ZendSymtableExistsInd(ht, key.value.str) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		break
	case 4:
		if zend.ZendHashIndexExists(ht, key.value.lval) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		break
	case 1:
		if zend.ZendHashExistsInd(ht, zend.ZendEmptyString) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "The first argument should be either a string or an integer")
		return_value.u1.type_info = 2
	}
	if array.u1.v.type_ != 7 {
		if ht != nil && (zend.ZvalGcFlags(ht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&ht.gc) == 0 {
			zend.ZendArrayDestroy(ht)
		}
	}
}

/* }}} */

func ZifArrayChunk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num_in int
	var size zend.ZendLong
	var current zend.ZendLong = 0
	var str_key *zend.ZendString
	var num_key zend.ZendUlong
	var preserve_keys zend.ZendBool = 0
	var input *zend.Zval = nil
	var chunk zend.Zval
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Do bounds checking for size parameter. */

	if size < 1 {
		core.PhpErrorDocref(nil, 1<<1, "Size parameter expected to be greater than 0")
		return
	}
	num_in = input.value.arr.nNumOfElements
	if size > num_in {
		if num_in > 0 {
			size = num_in
		} else {
			size = 1
		}
	}
	var __arr *zend.ZendArray = zend._zendNewArray(uint32((num_in-1)/size + 1))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	&chunk.u1.type_info = 0
	for {
		var __ht *zend.HashTable = input.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			str_key = _p.key
			entry = _z

			/* If new chunk, create and initialize it. */

			if chunk.u1.v.type_ == 0 {
				var __arr *zend.ZendArray = zend._zendNewArray(uint32(size))
				var __z *zend.Zval = &chunk
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			}

			/* Add entry to the chunk, preserving keys if necessary. */

			if preserve_keys != 0 {
				if str_key != nil {
					entry = zend.ZendHashUpdate(chunk.value.arr, str_key, entry)
				} else {
					entry = zend.ZendHashIndexUpdate(chunk.value.arr, num_key, entry)
				}
			} else {
				entry = zend.ZendHashNextIndexInsert(chunk.value.arr, entry)
			}
			zend.ZvalAddRef(entry)

			/* If reached the chunk size, add it to the result array, and reset the
			 * pointer. */

			if g.PreInc(&current)%size == 0 {
				zend.AddNextIndexZval(return_value, &chunk)
				&chunk.u1.type_info = 0
			}

			/* If reached the chunk size, add it to the result array, and reset the
			 * pointer. */

		}
		break
	}

	/* Add the final chunk if there is one. */

	if chunk.u1.v.type_ != 0 {
		zend.AddNextIndexZval(return_value, &chunk)
	}

	/* Add the final chunk if there is one. */
}

/* }}} */

func ZifArrayCombine(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var values *zend.HashTable
	var keys *zend.HashTable
	var pos_values uint32 = 0
	var entry_keys *zend.Zval
	var entry_values *zend.Zval
	var num_keys int
	var num_values int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArrayHt(_arg, &keys, 0, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArrayHt(_arg, &values, 0, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	num_keys = keys.nNumOfElements
	num_values = values.nNumOfElements
	if num_keys != num_values {
		core.PhpErrorDocref(nil, 1<<1, "Both parameters should have an equal number of elements")
		return_value.u1.type_info = 2
		return
	}
	if num_keys == 0 {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(num_keys)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = keys
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			entry_keys = _z
			for true {
				if pos_values >= values.nNumUsed {
					break
				} else if values.arData[pos_values].val.u1.v.type_ != 0 {
					entry_values = &values.arData[pos_values].val
					if entry_keys.u1.v.type_ == 4 {
						entry_values = zend.ZendHashIndexUpdate(return_value.value.arr, entry_keys.value.lval, entry_values)
					} else {
						var tmp_key *zend.ZendString
						var key *zend.ZendString = zend.ZvalGetTmpString(entry_keys, &tmp_key)
						entry_values = zend.ZendSymtableUpdate(return_value.value.arr, key, entry_values)
						zend.ZendTmpStringRelease(tmp_key)
					}
					zend.ZvalAddRef(entry_values)
					pos_values++
					break
				}
				pos_values++
			}
		}
		break
	}
}

/* }}} */
