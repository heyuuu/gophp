// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZendGeneratorGetCurrent(generator *ZendGenerator) *ZendGenerator {
	var leaf *ZendGenerator
	var root *ZendGenerator
	if generator.GetNode().GetParent() == nil {

		/* we're not in yield from mode */

		return generator

		/* we're not in yield from mode */

	}
	if generator.GetNode().GetChildren() != 0 {
		leaf = generator.GetNode().GetPtrLeaf()
	} else {
		leaf = generator
	}
	root = leaf.GetNode().GetRoot()
	if root.GetExecuteData() != nil && root.GetNode().GetParent() == nil {

		/* generator still running */

		return root

		/* generator still running */

	}
	return ZendGeneratorUpdateCurrent(generator, leaf)
}
func ZendGeneratorRestoreCallStack(generator *ZendGenerator) {
	var call *ZendExecuteData
	var new_call *ZendExecuteData
	var prev_call *ZendExecuteData = nil
	call = generator.GetFrozenCallStack()
	for {
		new_call = ZendVmStackPushCallFrame(ZEND_CALL_INFO(call) & ^ZEND_CALL_ALLOCATED, call.GetFunc(), call.NumArgs(), call.GetThis().GetPtr())
		memcpy((*types.Zval)(new_call)+ZEND_CALL_FRAME_SLOT, (*types.Zval)(call)+ZEND_CALL_FRAME_SLOT, call.NumArgs()*b.SizeOf("zval"))
		new_call.SetPrevExecuteData(prev_call)
		prev_call = new_call
		call = call.GetPrevExecuteData()
		if call == nil {
			break
		}
	}
	generator.GetExecuteData().SetCall(prev_call)
	Efree(generator.GetFrozenCallStack())
	generator.SetFrozenCallStack(nil)
}
func ZendGeneratorFreezeCallStack(executeData *ZendExecuteData) *ZendExecuteData {
	var used_stack int
	var call *ZendExecuteData
	var new_call *ZendExecuteData
	var prev_call *ZendExecuteData = nil
	var stack *types.Zval

	/* calculate required stack size */

	used_stack = 0
	call = executeData.GetCall()
	for {
		used_stack += ZEND_CALL_FRAME_SLOT + call.NumArgs()
		call = call.GetPrevExecuteData()
		if call == nil {
			break
		}
	}
	stack = Emalloc(used_stack * b.SizeOf("zval"))

	/* save stack, linking frames in reverse order */

	call = executeData.GetCall()
	for {
		var frame_size int = ZEND_CALL_FRAME_SLOT + call.NumArgs()
		new_call = (*ZendExecuteData)(stack + used_stack - frame_size)
		memcpy(new_call, call, frame_size*b.SizeOf("zval"))
		used_stack -= frame_size
		new_call.SetPrevExecuteData(prev_call)
		prev_call = new_call
		new_call = call.GetPrevExecuteData()
		ZendVmStackFreeCallFrame(call)
		call = new_call
		if call == nil {
			break
		}
	}
	executeData.SetCall(nil)
	b.Assert(prev_call == (*ZendExecuteData)(stack))
	return prev_call
}
func ZendGeneratorCleanupUnfinishedExecution(generator *ZendGenerator, executeData *ZendExecuteData, catch_op_num uint32) {
	var op_array *ZendOpArray = executeData.GetFunc().GetOpArray()
	if executeData.GetOpline() != op_array.GetOpcodes() {

		/* -1 required because we want the last run opcode, not the next to-be-run one. */

		var op_num uint32 = executeData.GetOpline() - op_array.GetOpcodes() - 1
		if generator.GetFrozenCallStack() != nil {

			/* Temporarily restore generator->executeData if it has been NULLed out already. */

			var save_ex *ZendExecuteData = generator.GetExecuteData()
			generator.SetExecuteData(executeData)
			ZendGeneratorRestoreCallStack(generator)
			generator.SetExecuteData(save_ex)
		}
		ZendCleanupUnfinishedExecution(executeData, op_num, catch_op_num)
	}
}
func ZendGeneratorClose(generator *ZendGenerator, finished_execution types.ZendBool) {
	if generator.GetExecuteData() != nil {
		var executeData *ZendExecuteData = generator.GetExecuteData()

		/* Null out executeData early, to prevent double frees if GC runs while we're
		 * already cleaning up executeData. */

		generator.SetExecuteData(nil)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			ZendCleanAndCacheSymbolTable(executeData.GetSymbolTable())
		}

		/* always free the CV's, in the symtable are only not-free'd IS_INDIRECT's */

		ZendFreeCompiledVariables(executeData)
		if (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0 {
			OBJ_RELEASE(executeData.GetThis().GetObj())
		}

		/* A fatal error / die occurred during the generator execution.
		 * Trying to clean up the stack may not be safe in this case. */

		if CG__().GetUncleanShutdown() != 0 {
			generator.SetExecuteData(nil)
			return
		}
		ZendVmStackFreeExtraArgs(executeData)

		/* Some cleanups are only necessary if the generator was closed
		 * before it could finish execution (reach a return statement). */

		if finished_execution == 0 {
			ZendGeneratorCleanupUnfinishedExecution(generator, executeData, 0)
		}

		/* Free closure object */

		if (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0 {
			OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc(

				/* Free GC buffer. GC for closed generators doesn't need an allocated buffer */)))
		}

		if generator.GetGcBuffer() != nil {
			Efree(generator.GetGcBuffer())
			generator.SetGcBuffer(nil)
		}
		Efree(executeData)
	}
}
func ZendGeneratorDtorStorage(object *types.ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	var ex *ZendExecuteData = generator.GetExecuteData()
	var op_num uint32
	var try_catch_offset uint32
	var i int

	/* leave yield from mode to properly allow finally execution */

	if generator.GetValues().GetType() != types.IS_UNDEF {
		ZvalPtrDtor(generator.GetValues())
		generator.GetValues().SetUndef()
	}
	if generator.GetNode().GetChildren() == 0 {
		var root *ZendGenerator = generator.GetNode().GetRoot()
		var next *ZendGenerator
		for root != generator {
			next = ZendGeneratorGetChild(root.GetNode(), generator)
			generator.GetNode().SetRoot(next)
			next.GetNode().SetParent(nil)
			OBJ_RELEASE(root.GetStd())
			root = next
		}
	}
	if ex == nil || !ex.GetFunc().GetOpArray().IsHasFinallyBlock() || CG__().GetUncleanShutdown() != 0 {
		return
	}

	/* -1 required because we want the last run opcode, not the
	 * next to-be-run one. */

	op_num = ex.GetOpline() - ex.GetFunc().GetOpArray().GetOpcodes() - 1
	try_catch_offset = -1

	/* Find the innermost try/catch that we are inside of. */

	for i = 0; i < ex.GetFunc().GetOpArray().GetLastTryCatch(); i++ {
		var try_catch *ZendTryCatchElement = ex.GetFunc().GetOpArray().GetTryCatchArray()[i]
		if op_num < try_catch.GetTryOp() {
			break
		}
		if op_num < try_catch.GetCatchOp() || op_num < try_catch.GetFinallyEnd() {
			try_catch_offset = i
		}
	}

	/* Walk try/catch/finally structures upwards, performing the necessary actions. */

	for try_catch_offset != uint32-1 {
		var try_catch *ZendTryCatchElement = ex.GetFunc().GetOpArray().GetTryCatchArray()[try_catch_offset]
		if op_num < try_catch.GetFinallyOp() {

			/* Go to finally block */

			var fast_call *types.Zval = ZEND_CALL_VAR(ex, ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar())
			ZendGeneratorCleanupUnfinishedExecution(generator, ex, try_catch.GetFinallyOp())
			fast_call.SetObj(EG__().GetException())
			EG__().SetException(nil)
			fast_call.SetOplineNum(uint32 - 1)
			ex.SetOpline(ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyOp()])
			generator.SetIsForcedClose(true)
			ZendGeneratorResume(generator)

			/* TODO: If we hit another yield inside try/finally,
			 * should we also jump to the next finally block? */

			return

			/* TODO: If we hit another yield inside try/finally,
			 * should we also jump to the next finally block? */

		} else if op_num < try_catch.GetFinallyEnd() {
			var fast_call *types.Zval = ZEND_CALL_VAR(ex, ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar())

			/* Clean up incomplete return statement */

			if fast_call.GetOplineNum() != uint32-1 {
				var retval_op *ZendOp = ex.GetFunc().GetOpArray().GetOpcodes()[fast_call.GetOplineNum()]
				if (retval_op.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
					ZvalPtrDtor(ZEND_CALL_VAR(ex, retval_op.GetOp2().GetVar()))
				}
			}

			/* Clean up backed-up exception */

			if fast_call.GetObj() != nil {
				OBJ_RELEASE(fast_call.GetObj())
			}

			/* Clean up backed-up exception */

		}
		try_catch_offset--
	}

	/* Walk try/catch/finally structures upwards, performing the necessary actions. */
}
func ZendGeneratorFreeStorage(object *types.ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	ZendGeneratorClose(generator, 0)

	/* we can't immediately free them in zend_generator_close() else yield from won't be able to fetch it */

	ZvalPtrDtor(generator.GetValue())
	ZvalPtrDtor(generator.GetKey())
	if !(generator.GetRetval().IsUndef()) {
		ZvalPtrDtor(generator.GetRetval())
	}
	if generator.GetNode().GetChildren() > 1 {
		generator.GetNode().GetHt().Destroy()
		Efree(generator.GetNode().GetHt())
	}
	ZendObjectStdDtor(generator.GetStd())
}
func CalcGcBufferSize(generator *ZendGenerator) uint32 {
	var size uint32 = 4
	if generator.GetExecuteData() != nil {
		var executeData *ZendExecuteData = generator.GetExecuteData()
		var op_array *ZendOpArray = executeData.GetFunc().op_array

		/* Compiled variables */

		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			size += op_array.GetLastVar()
		}

		/* Extra args */

		if (EX_CALL_INFO() & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
			size += executeData.NumArgs() - op_array.GetNumArgs()
		}
		size += (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0
		size += (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0

		/* Live vars */

		if executeData.GetOpline() != op_array.GetOpcodes() {

			/* -1 required because we want the last run opcode, not the next to-be-run one. */

			var i uint32
			var op_num uint32 = executeData.GetOpline() - op_array.GetOpcodes() - 1
			for i = 0; i < op_array.GetLastLiveRange(); i++ {
				var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
				if range_.GetStart() > op_num {

					/* Further ranges will not be relevant... */

					break

					/* Further ranges will not be relevant... */

				} else if op_num < range_.GetEnd() {

					/* LIVE_ROPE and LIVE_SILENCE not relevant for GC */

					var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
					if kind == ZEND_LIVE_TMPVAR || kind == ZEND_LIVE_LOOP {
						size++
					}
				}
			}
		}

		/* Yield from root references */

		if generator.GetNode().GetChildren() == 0 {
			var root *ZendGenerator = generator.GetNode().GetRoot()
			for root != generator {
				root = ZendGeneratorGetChild(root.GetNode(), generator)
				size++
			}
		}

		/* Yield from root references */

	}
	return size
}
func ZendGeneratorGetGc(object *types.Zval, table **types.Zval, n *int) *types.Array {
	var generator *ZendGenerator = (*ZendGenerator)(object.GetObj())
	var executeData *ZendExecuteData = generator.GetExecuteData()
	var op_array *ZendOpArray
	var gc_buffer *types.Zval
	var gc_buffer_size uint32
	if executeData == nil {

		/* If the generator has been closed, it can only hold on to three values: The value, key
		 * and retval. These three zvals are stored sequentially starting at &generator->value. */

		*table = generator.GetValue()
		*n = 3
		return nil
	}
	if generator.IsCurrentlyRunning() {

		/* If the generator is currently running, we certainly won't be able to GC any values it
		 * holds on to. The executeData state might be inconsistent during execution (e.g. because
		 * GC has been triggered in the middle of a variable reassignment), so we should not try
		 * to inspect it here. */

		*table = nil
		*n = 0
		return nil
	}
	op_array = executeData.GetFunc().op_array
	gc_buffer_size = CalcGcBufferSize(generator)
	if generator.GetGcBufferSize() < gc_buffer_size {
		generator.SetGcBuffer(SafeErealloc(generator.GetGcBuffer(), b.SizeOf("zval"), gc_buffer_size, 0))
		generator.SetGcBufferSize(gc_buffer_size)
	}
	*n = gc_buffer_size
	gc_buffer = generator.GetGcBuffer()
	*table = gc_buffer
	types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetValue())
	types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetKey())
	types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetRetval())
	types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetValues())
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
		var i uint32
		var num_cvs uint32 = executeData.GetFunc().op_array.last_var
		for i = 0; i < num_cvs; i++ {
			types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), executeData.VarNum(i))
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
		var zv *types.Zval = executeData.VarNum(op_array.GetLastVar() + op_array.GetT())
		var end *types.Zval = zv + (executeData.NumArgs() - op_array.GetNumArgs())
		for zv != end {
			types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), b.PostInc(&zv))
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0 {
		b.PostInc(&gc_buffer).SetObject(executeData.GetThis().GetObj())
	}
	if (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0 {
		b.PostInc(&gc_buffer).SetObject(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
	}
	if executeData.GetOpline() != op_array.GetOpcodes() {
		var i uint32
		var op_num uint32 = executeData.GetOpline() - op_array.GetOpcodes() - 1
		for i = 0; i < op_array.GetLastLiveRange(); i++ {
			var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
			if range_.GetStart() > op_num {
				break
			} else if op_num < range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *types.Zval = EX_VAR(var_num)
				if kind == ZEND_LIVE_TMPVAR || kind == ZEND_LIVE_LOOP {
					types.ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), var_)
				}
			}
		}
	}
	if generator.GetNode().GetChildren() == 0 {
		var root *ZendGenerator = generator.GetNode().GetRoot()
		for root != generator {
			b.PostInc(&gc_buffer).SetObject(root.GetStd())
			root = ZendGeneratorGetChild(root.GetNode(), generator)
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		return executeData.GetSymbolTable()
	} else {
		return nil
	}
}
func ZendGeneratorCreate(class_type *types.ClassEntry) *types.ZendObject {
	var generator *ZendGenerator
	generator = Emalloc(b.SizeOf("zend_generator"))
	memset(generator, 0, b.SizeOf("zend_generator"))

	/* The key will be incremented on first use, so it'll start at 0 */

	generator.SetLargestUsedIntegerKey(-1)
	generator.GetRetval().SetUndef()
	generator.GetValues().SetUndef()

	/* By default we have a tree of only one node */

	generator.GetNode().SetParent(nil)
	generator.GetNode().SetChildren(0)
	generator.GetNode().SetRoot(generator)
	ZendObjectStdInit(generator.GetStd(), class_type)
	generator.GetStd().SetHandlers(&ZendGeneratorHandlers)
	return (*types.ZendObject)(generator)
}
func ZendGeneratorGetConstructor(object *types.ZendObject) *ZendFunction {
	faults.ThrowError(nil, "The \"Generator\" class is reserved for internal use and cannot be manually instantiated")
	return nil
}
func ZendGeneratorCheckPlaceholderFrame(ptr *ZendExecuteData) *ZendExecuteData {
	if ptr.GetFunc() == nil && ptr.GetThis().IsObject() {
		if types.Z_OBJCE(ptr.GetThis()) == ZendCeGenerator {
			var generator *ZendGenerator = (*ZendGenerator)(ptr.GetThis().GetObj())
			var root *ZendGenerator = b.CondF2(generator.GetNode().GetChildren() < 1, generator, func() *ZendGenerator { return generator.GetNode().GetPtrLeaf() }).node.ptr.root
			var prev *ZendExecuteData = ptr.GetPrevExecuteData()
			if generator.GetNode().GetParent() != root {
				for {
					generator.GetExecuteData().SetPrevExecuteData(prev)
					prev = generator.GetExecuteData()
					generator = generator.GetNode().GetParent()
					if generator.GetNode().GetParent() == root {
						break
					}
				}
			}
			generator.GetExecuteData().SetPrevExecuteData(prev)
			ptr = generator.GetExecuteData()
		}
	}
	return ptr
}
func ZendGeneratorThrowException(generator *ZendGenerator, exception *types.Zval) {
	var original_execute_data *ZendExecuteData = CurrEX()

	/* if we don't stop an array/iterator yield from, the exception will only reach the generator after the values were all iterated over */

	if generator.GetValues().GetType() != types.IS_UNDEF {
		ZvalPtrDtor(generator.GetValues())
		generator.GetValues().SetUndef()
	}

	/* Throw the exception in the context of the generator. Decrementing the opline
	 * to pretend the exception happened during the YIELD opcode. */

	EG__().SetCurrentExecuteData(generator.GetExecuteData())
	generator.GetExecuteData().GetOpline()--
	if exception != nil {
		faults.ThrowExceptionObject(exception)
	} else {
		faults.RethrowException(CurrEX())
	}
	generator.GetExecuteData().GetOpline()++
	EG__().SetCurrentExecuteData(original_execute_data)
}
func ZendGeneratorGetChild(node *ZendGeneratorNode, leaf *ZendGenerator) *ZendGenerator {
	if node.GetChildren() == 0 {
		return nil
	} else if node.GetChildren() == 1 {
		return node.GetChildSingleChild()
	} else {
		return types.ZendHashIndexFindPtr(node.GetHt(), ZendUlong(leaf))
	}
}
func ZendGeneratorSearchMultiChildrenNode(node *ZendGeneratorNode) *ZendGeneratorNode {
	for node.GetChildren() == 1 {
		node = node.GetChildSingleChild().GetNode()
	}
	if node.GetChildren() > 1 {
		return node
	} else {
		return nil
	}
}
func ZendGeneratorAddSingleChild(node *ZendGeneratorNode, child *ZendGenerator, leaf *ZendGenerator) {
	if node.GetChildren() == 0 {
		node.SetChildSingleLeaf(leaf)
		node.SetChildSingleChild(child)
	} else {
		if node.GetChildren() == 1 {
			var ht *types.Array = Emalloc(b.SizeOf("HashTable"))
			types.ZendHashInit(ht, 0, nil, nil, 0)
			types.ZendHashIndexAddPtr(ht, ZendUlong(node.GetChildSingleLeaf()), node.GetChildSingleChild())
			node.SetHt(ht)
		}
		types.ZendHashIndexAddPtr(node.GetHt(), ZendUlong(leaf), child)
	}
	node.GetChildren()++
}
func ZendGeneratorMergeChildNodes(dest *ZendGeneratorNode, src *ZendGeneratorNode, child *ZendGenerator) {
	var leaf ZendUlong
	b.Assert(src.GetChildren() > 1)
	var __ht *types.Array = src.GetHt()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		leaf = _p.GetH()
		ZendGeneratorAddSingleChild(dest, child, (*ZendGenerator)(leaf))
	}
}
func ZendGeneratorAddChild(generator *ZendGenerator, child *ZendGenerator) {
	var leaf *ZendGenerator = b.CondF1(child.GetNode().GetChildren() != 0, func() *ZendGenerator { return child.GetNode().GetPtrLeaf() }, child)
	var multi_children_node *ZendGeneratorNode
	var was_leaf types.ZendBool = generator.GetNode().GetChildren() == 0
	if was_leaf != 0 {
		var next *ZendGenerator = generator.GetNode().GetParent()
		leaf.GetNode().SetRoot(generator.GetNode().GetRoot())
		generator.GetStd().AddRefcount()
		generator.GetNode().SetPtrLeaf(leaf)
		for next != nil {
			if next.GetNode().GetChildren() > 1 {
				var child *ZendGenerator = types.ZendHashIndexFindPtr(next.GetNode().GetHt(), ZendUlong(generator))
				types.ZendHashIndexDel(next.GetNode().GetHt(), ZendUlong(generator))
				types.ZendHashIndexAddPtr(next.GetNode().GetHt(), ZendUlong(leaf), child)
			}
			next.GetNode().SetPtrLeaf(leaf)
			next = next.GetNode().GetParent()
		}
	} else if generator.GetNode().GetChildren() == 1 {
		multi_children_node = ZendGeneratorSearchMultiChildrenNode(generator.GetNode())
		if multi_children_node != nil {
			generator.GetNode().SetChildren(0)
			ZendGeneratorMergeChildNodes(generator.GetNode(), multi_children_node, generator.GetNode().GetChildSingleChild())
		}
	}
	if was_leaf == 0 {
		multi_children_node = ZendGeneratorSearchMultiChildrenNode(child.GetNode())
	} else {
		multi_children_node = (*ZendGeneratorNode)(0x1)
	}
	var parent *ZendGenerator = generator.GetNode().GetParent()
	var cur *ZendGenerator = generator
	if multi_children_node > (*ZendGeneratorNode)(0x1) {
		ZendGeneratorMergeChildNodes(generator.GetNode(), multi_children_node, child)
	} else {
		ZendGeneratorAddSingleChild(generator.GetNode(), child, leaf)
	}
	for parent != nil {
		if parent.GetNode().GetChildren() > 1 {
			if multi_children_node == (*ZendGeneratorNode)(0x1) {
				multi_children_node = ZendGeneratorSearchMultiChildrenNode(child.GetNode())
			}
			if multi_children_node != nil {
				ZendGeneratorMergeChildNodes(parent.GetNode(), multi_children_node, cur)
			} else {
				ZendGeneratorAddSingleChild(parent.GetNode(), cur, leaf)
			}
		}
		cur = parent
		parent = parent.GetNode().GetParent()
	}
}
func ZendGeneratorYieldFrom(generator *ZendGenerator, from *ZendGenerator) {
	ZendGeneratorAddChild(from, generator)
	generator.GetNode().SetParent(from)
	ZendGeneratorGetCurrent(generator)
	from.GetStd().DelRefcount()
	generator.SetIsDoInit(true)
}
func ZendGeneratorUpdateCurrent(generator *ZendGenerator, leaf *ZendGenerator) *ZendGenerator {
	var old_root *ZendGenerator
	var root *ZendGenerator = leaf.GetNode().GetRoot()

	/* generator at the root had stopped */

	if root != generator {
		old_root = root
		root = ZendGeneratorGetChild(root.GetNode(), leaf)
	} else {
		old_root = nil
	}
	for root.GetExecuteData() == nil && root != generator {
		OBJ_RELEASE(old_root.GetStd())
		old_root = root
		root = ZendGeneratorGetChild(root.GetNode(), leaf)
	}
	if root.GetNode().GetParent() != nil {
		if root.GetNode().GetParent().GetExecuteData() == nil {
			if EG__().GetException() == nil {
				var yield_from *ZendOp = (*ZendOp)(root.GetExecuteData().GetOpline() - 1)
				if yield_from.GetOpcode() == ZEND_YIELD_FROM {
					if root.GetNode().GetParent().GetRetval().IsUndef() {

						/* Throw the exception in the context of the generator */

						var original_execute_data *ZendExecuteData = CurrEX()
						EG__().SetCurrentExecuteData(root.GetExecuteData())
						if root == generator {
							root.GetExecuteData().SetPrevExecuteData(original_execute_data)
						} else {
							root.GetExecuteData().SetPrevExecuteData(generator.GetExecuteFake())
							generator.GetExecuteFake().SetPrevExecuteData(original_execute_data)
						}
						root.GetExecuteData().GetOpline()--
						faults.ThrowException(zend_ce_ClosedGeneratorException, "Generator yielded from aborted, no return value available", 0)
						EG__().SetCurrentExecuteData(original_execute_data)
						if (b.Cond(old_root != nil, old_root, generator).flags & ZEND_GENERATOR_CURRENTLY_RUNNING) == 0 {
							leaf.GetNode().SetRoot(root)
							root.GetNode().SetParent(nil)
							if old_root != nil {
								OBJ_RELEASE(old_root.GetStd())
							}
							ZendGeneratorResume(leaf)
							return leaf.GetNode().GetRoot()
						}
					} else {
						ZvalPtrDtor(root.GetValue())
						types.ZVAL_COPY(root.GetValue(), root.GetNode().GetParent().GetValue())
						types.ZVAL_COPY(ZEND_CALL_VAR(root.GetExecuteData(), yield_from.GetResult().GetVar()), root.GetNode().GetParent().GetRetval())
					}
				}
			}
			root.GetNode().SetParent(nil)
		} else {
			for {
				root = root.GetNode().GetParent()
				root.GetStd().AddRefcount()
				if root.GetNode().GetParent() == nil {
					break
				}
			}
		}
	}
	leaf.GetNode().SetRoot(root)
	if old_root != nil {
		OBJ_RELEASE(old_root.GetStd())
	}
	return root
}
func ZendGeneratorGetNextDelegatedValue(generator *ZendGenerator) int {
	var value *types.Zval
	if generator.GetValues().IsArray() {
		var ht *types.Array = generator.GetValues().GetArr()
		var pos types.HashPosition = generator.GetValues().GetFePos()
		var p *types.Bucket
		for {
			if pos >= ht.GetNNumUsed() {

				/* Reached end of array */

				goto failure

				/* Reached end of array */

			}
			p = ht.GetArData()[pos]
			value = p.GetVal()
			if value.IsIndirect() {
				value = value.GetZv()
			}
			pos++
			if !(value.IsUndef()) {
				break
			}
		}
		ZvalPtrDtor(generator.GetValue())
		types.ZVAL_COPY(generator.GetValue(), value)
		ZvalPtrDtor(generator.GetKey())
		if p.GetKey() != nil {
			generator.GetKey().SetStringCopy(p.GetKey())
		} else {
			generator.GetKey().SetLong(p.GetH())
		}
		generator.GetValues().GetFePos() = pos
	} else {
		var iter *ZendObjectIterator = (*ZendObjectIterator)(generator.GetValues().GetObj())
		if b.PostInc(&(iter.GetIndex())) > 0 {
			iter.GetFuncs().GetMoveForward()(iter)
			if EG__().GetException() != nil {
				goto exception
			}
		}
		if iter.GetFuncs().GetValid()(iter) == types.FAILURE {
			if EG__().GetException() != nil {
				goto exception
			}

			/* reached end of iteration */

			goto failure

			/* reached end of iteration */

		}
		value = iter.GetFuncs().GetGetCurrentData()(iter)
		if EG__().GetException() != nil {
			goto exception
		} else if value == nil {
			goto failure
		}
		ZvalPtrDtor(generator.GetValue())
		types.ZVAL_COPY(generator.GetValue(), value)
		ZvalPtrDtor(generator.GetKey())
		if iter.GetFuncs().GetGetCurrentKey() != nil {
			iter.GetFuncs().GetGetCurrentKey()(iter, generator.GetKey())
			if EG__().GetException() != nil {
				generator.GetKey().SetUndef()
				goto exception
			}
		} else {
			generator.GetKey().SetLong(iter.GetIndex())
		}
	}
	return types.SUCCESS
exception:
	ZendGeneratorThrowException(generator, nil)
failure:
	ZvalPtrDtor(generator.GetValues())
	generator.GetValues().SetUndef()
	return types.FAILURE
}
func ZendGeneratorResume(orig_generator *ZendGenerator) {
	var generator *ZendGenerator = ZendGeneratorGetCurrent(orig_generator)

	/* The generator is already closed, thus can't resume */

	if generator.GetExecuteData() == nil {
		return
	}
try_again:
	if generator.IsCurrentlyRunning() {
		faults.ThrowError(nil, "Cannot resume an already running generator")
		return
	}
	if orig_generator.IsDoInit() && !(generator.GetValue().IsUndef()) {

		/* We must not advance Generator if we yield from a Generator being currently run */

		orig_generator.SetIsDoInit(false)
		return
	}
	if !(generator.GetValues().IsUndef()) {
		if ZendGeneratorGetNextDelegatedValue(generator) == types.SUCCESS {
			orig_generator.SetIsDoInit(false)
			return
		}
	}

	/* Drop the AT_FIRST_YIELD flag */

	orig_generator.SetIsAtFirstYield(false)

	/* Backup executor globals */

	var original_execute_data *ZendExecuteData = CurrEX()

	/* Set executor globals */

	EG__().SetCurrentExecuteData(generator.GetExecuteData())

	/* We want the backtrace to look as if the generator function was
	 * called from whatever method we are current running (e.g. next()).
	 * So we have to link generator call frame with caller call frame. */

	if generator == orig_generator {
		generator.GetExecuteData().SetPrevExecuteData(original_execute_data)
	} else {

		/* We need some executeData placeholder in stacktrace to be replaced
		 * by the real stack trace when needed */

		generator.GetExecuteData().SetPrevExecuteData(orig_generator.GetExecuteFake())
		orig_generator.GetExecuteFake().SetPrevExecuteData(original_execute_data)
	}
	if generator.GetFrozenCallStack() != nil {

		/* Restore frozen call-stack */

		ZendGeneratorRestoreCallStack(generator)

		/* Restore frozen call-stack */

	}

	/* Resume execution */

	generator.SetIsCurrentlyRunning(true)
	ZendExecuteEx(generator.GetExecuteData())
	generator.SetIsCurrentlyRunning(false)
	generator.SetFrozenCallStack(nil)
	if generator.GetExecuteData() != nil && generator.GetExecuteData().GetCall() != nil {

		/* Frize call-stack */

		generator.SetFrozenCallStack(ZendGeneratorFreezeCallStack(generator.GetExecuteData()))

		/* Frize call-stack */

	}

	/* Restore executor globals */

	EG__().SetCurrentExecuteData(original_execute_data)

	/* If an exception was thrown in the generator we have to internally
	 * rethrow it in the parent scope.
	 * In case we did yield from, the Exception must be rethrown into
	 * its calling frame (see above in if (check_yield_from). */

	if EG__().GetException() != nil {
		if generator == orig_generator {
			ZendGeneratorClose(generator, 0)
			if CurrEX() == nil {
				faults.ThrowExceptionInternal(nil)
			} else if CurrEX().GetFunc() != nil && ZEND_USER_CODE(CurrEX().GetFunc().GetCommonType()) {
				faults.RethrowException(CurrEX())
			}
		} else {
			generator = ZendGeneratorGetCurrent(orig_generator)
			ZendGeneratorThrowException(generator, nil)
			orig_generator.SetIsDoInit(false)
			goto try_again
		}
	}

	/* yield from was used, try another resume. */

	if generator != orig_generator && !(generator.GetRetval().IsUndef()) || generator.GetExecuteData() != nil && (generator.GetExecuteData().GetOpline()-1).opcode == ZEND_YIELD_FROM {
		generator = ZendGeneratorGetCurrent(orig_generator)
		goto try_again
	}

	/* yield from was used, try another resume. */

	orig_generator.SetIsDoInit(false)
}
func ZendGeneratorEnsureInitialized(generator *ZendGenerator) {
	if generator.GetValue().IsUndef() && generator.GetExecuteData() != nil && generator.GetNode().GetParent() == nil {
		ZendGeneratorResume(generator)
		generator.SetIsAtFirstYield(true)
	}
}
func ZendGeneratorRewind(generator *ZendGenerator) {
	ZendGeneratorEnsureInitialized(generator)
	if !generator.IsAtFirstYield() {
		faults.ThrowException(nil, "Cannot rewind a generator that was already run", 0)
	}
}
func zim_Generator_rewind(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorRewind(generator)
}
func zim_Generator_valid(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorGetCurrent(generator)
	types.ZVAL_BOOL(return_value, generator.GetExecuteData() != nil)
	return
}
func zim_Generator_current(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetValue().GetType() != types.IS_UNDEF {
		var value *types.Zval = root.GetValue()
		types.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_Generator_key(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetKey().GetType() != types.IS_UNDEF {
		var key *types.Zval = root.GetKey()
		types.ZVAL_COPY_DEREF(return_value, key)
	}
}
func zim_Generator_next(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorResume(generator)
}
func zim_Generator_send(executeData *ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	var generator *ZendGenerator
	var root *ZendGenerator
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			value = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)

	/* The generator is already closed, thus can't send anything */

	if generator.GetExecuteData() == nil {
		return
	}
	root = ZendGeneratorGetCurrent(generator)

	/* Put sent value in the target VAR slot, if it is used */

	if root.GetSendTarget() != nil {
		types.ZVAL_COPY(root.GetSendTarget(), value)
	}
	ZendGeneratorResume(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		var value *types.Zval = root.GetValue()
		types.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_Generator_throw(executeData *ZendExecuteData, return_value *types.Zval) {
	var exception *types.Zval
	var generator *ZendGenerator
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			exception = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	exception.TryAddRefcount()
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if generator.GetExecuteData() != nil {
		var root *ZendGenerator = ZendGeneratorGetCurrent(generator)
		ZendGeneratorThrowException(root, exception)
		ZendGeneratorResume(generator)
		root = ZendGeneratorGetCurrent(generator)
		if generator.GetExecuteData() != nil {
			var value *types.Zval = root.GetValue()
			types.ZVAL_COPY_DEREF(return_value, value)
		}
	} else {

		/* If the generator is already closed throw the exception in the
		 * current context */

		faults.ThrowExceptionObject(exception)

		/* If the generator is already closed throw the exception in the
		 * current context */

	}
}
func zim_Generator_getReturn(executeData *ZendExecuteData, return_value *types.Zval) {
	var generator *ZendGenerator
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS(executeData).GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if EG__().GetException() != nil {
		return
	}
	if generator.GetRetval().IsUndef() {

		/* Generator hasn't returned yet -> error! */

		faults.ThrowException(nil, "Cannot get return value of a generator that hasn't returned", 0)
		return
	}
	types.ZVAL_COPY(return_value, generator.GetRetval())
}
func ZendGeneratorIteratorDtor(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	generator.SetIterator(nil)
	ZvalPtrDtor(iterator.GetData())
}
func ZendGeneratorIteratorValid(iterator *ZendObjectIterator) int {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendGeneratorIteratorGetData(iterator *ZendObjectIterator) *types.Zval {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	return root.GetValue()
}
func ZendGeneratorIteratorGetKey(iterator *ZendObjectIterator, key *types.Zval) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if root.GetKey().GetType() != types.IS_UNDEF {
		var zv *types.Zval = root.GetKey()
		types.ZVAL_COPY_DEREF(key, zv)
	} else {
		key.SetNull()
	}
}
func ZendGeneratorIteratorMoveForward(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorResume(generator)
}
func ZendGeneratorIteratorRewind(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	ZendGeneratorRewind(generator)
}
func ZendGeneratorGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendObjectIterator
	var generator *ZendGenerator = (*ZendGenerator)(object.GetObj())
	if generator.GetExecuteData() == nil {
		faults.ThrowException(nil, "Cannot traverse an already closed generator", 0)
		return nil
	}
	if by_ref != 0 && !generator.GetExecuteData().GetFunc().GetOpArray().IsReturnReference() {
		faults.ThrowException(nil, "You can only iterate a generator by-reference if it declared that it yields by-reference", 0)
		return nil
	}
	generator.SetIterator(Emalloc(b.SizeOf("zend_object_iterator")))
	iterator = generator.GetIterator()
	ZendIteratorInit(iterator)
	iterator.SetFuncs(&ZendGeneratorIteratorFunctions)
	object.AddRefcount()
	iterator.GetData().SetObject(object.GetObj())
	return iterator
}
func ZendRegisterGeneratorCe() {
	var ce types.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Generator", b.SizeOf("\"Generator\"")-1, 1))
	ce.SetBuiltinFunctions(GeneratorFunctions)
	ZendCeGenerator = ZendRegisterInternalClass(&ce)
	ZendCeGenerator.SetIsFinal(true)
	ZendCeGenerator.SetCreateObject(ZendGeneratorCreate)
	ZendCeGenerator.SetSerialize(ZendClassSerializeDeny)
	ZendCeGenerator.SetUnserialize(ZendClassUnserializeDeny)

	/* get_iterator has to be assigned *after* implementing the inferface */

	ZendClassImplements(ZendCeGenerator, 1, ZendCeIterator)
	ZendCeGenerator.SetGetIterator(ZendGeneratorGetIterator)
	memcpy(&ZendGeneratorHandlers, &StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	ZendGeneratorHandlers.SetFreeObj(ZendGeneratorFreeStorage)
	ZendGeneratorHandlers.SetDtorObj(ZendGeneratorDtorStorage)
	ZendGeneratorHandlers.SetGetGc(ZendGeneratorGetGc)
	ZendGeneratorHandlers.SetCloneObj(nil)
	ZendGeneratorHandlers.SetGetConstructor(ZendGeneratorGetConstructor)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ClosedGeneratorException", b.SizeOf("\"ClosedGeneratorException\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	zend_ce_ClosedGeneratorException = ZendRegisterInternalClassEx(&ce, faults.ZendCeException)
}
