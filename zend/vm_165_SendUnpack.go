package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_SEND_UNPACK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var args *types.Zval
	var arg_num int
	args = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
	arg_num = executeData.GetCall().NumArgs() + 1
send_again:
	if args.IsArray() {
		var ht *types.Array = args.GetArr()
		var arg *types.Zval
		var top *types.Zval
		var name *types.String
		ZendVmStackExtendCallFrame(&(executeData.GetCall()), arg_num-1, ht.Len())
		if (opline.GetOp1Type() & (IS_VAR | IS_CV)) != 0 {
			var i uint32
			var separate int = 0

			/* check if any of arguments are going to be passed by reference */

			for i = 0; i < ht.Len(); i++ {
				if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num+i) != 0 {
					separate = 1
					break
				}
			}
			if separate != 0 {
				types.SeparateArray(args)
				ht = args.GetArr()
			}
		}
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			name = _p.GetKey()
			arg = _z
			if name != nil {
				faults.ThrowError(nil, "Cannot unpack array with string keys")
				// 				FREE_OP(free_op1)
				return 0
			}
			top = executeData.GetCall().Arg(arg_num)
			if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
				if arg.IsRef() {
					// 					arg.AddRefcount()
					top.SetReference(arg.Ref())
				} else if (opline.GetOp1Type() & (IS_VAR | IS_CV)) != 0 {

					/* array is already separated above */

					types.ZVAL_MAKE_REF_EX(arg, 2)
					top.SetReference(arg.Ref())
				} else {
					// arg.TryAddRefcount()
					top.SetNewRef(arg)
				}
			} else {
				types.ZVAL_COPY_DEREF(top, arg)
			}
			executeData.GetCall().
				NumArgs()++
			arg_num++
		}
	} else if args.IsObject() {
		var ce *types.ClassEntry = types.Z_OBJCE_P(args)
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			faults.Error(faults.E_WARNING, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, args, 0)
			if iter == nil {
				// 				FREE_OP(free_op1)
				if EG__().NoException() {
					faults.ThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.Name())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for ; iter.GetFuncs().GetValid()(iter) == types.SUCCESS; arg_num++ {
				var arg *types.Zval
				var top *types.Zval
				if EG__().HasException() {
					break
				}
				arg = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG__().HasException() {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key types.Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG__().HasException() {
						break
					}
					if !key.IsLong() {
						faults.ThrowError(nil, b.Cond(key.IsString(), "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						// ZvalPtrDtor(&key)
						break
					}
				}
				if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
					var scopePrefix string
					if executeData.GetCall().GetFunc().GetScope() != nil {
						scopePrefix = executeData.GetCall().GetFunc().GetScope().Name() + "::"
					}

					faults.Error(faults.E_WARNING, fmt.Sprintf("Cannot pass by-reference argument %d of %s%s() by unpacking a Traversable, passing by-value instead", arg_num, scopePrefix, executeData.GetCall().FunctionName()))
				}
				arg = types.ZVAL_DEREF(arg)
				// arg.TryAddRefcount()
				ZendVmStackExtendCallFrame(&(executeData.GetCall()), arg_num-1, 1)
				top = executeData.GetCall().Arg(arg_num)
				top.CopyValueFrom(arg)
				executeData.GetCall().
					NumArgs()++
				iter.GetFuncs().GetMoveForward()(iter)
			}
			//ZendIteratorDtor(iter)
		}
	} else if args.IsRef() {
		args = types.Z_REFVAL_P(args)
		goto send_again
	} else {
		if opline.GetOp1Type() == IS_CV && args.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		faults.Error(faults.E_WARNING, "Only arrays and Traversables can be unpacked")
	}
	// 	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
