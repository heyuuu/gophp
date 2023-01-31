// <<generate>>

package zend

import (
	b "sik/builtin"
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
		new_call = ZendVmStackPushCallFrame(ZEND_CALL_INFO(call) & ^ZEND_CALL_ALLOCATED, call.GetFunc(), ZEND_CALL_NUM_ARGS(call), call.GetThis().GetPtr())
		memcpy((*Zval)(new_call)+ZEND_CALL_FRAME_SLOT, (*Zval)(call)+ZEND_CALL_FRAME_SLOT, ZEND_CALL_NUM_ARGS(call)*b.SizeOf("zval"))
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
func ZendGeneratorFreezeCallStack(execute_data *ZendExecuteData) *ZendExecuteData {
	var used_stack int
	var call *ZendExecuteData
	var new_call *ZendExecuteData
	var prev_call *ZendExecuteData = nil
	var stack *Zval

	/* calculate required stack size */

	used_stack = 0
	call = EX(call)
	for {
		used_stack += ZEND_CALL_FRAME_SLOT + ZEND_CALL_NUM_ARGS(call)
		call = call.GetPrevExecuteData()
		if call == nil {
			break
		}
	}
	stack = Emalloc(used_stack * b.SizeOf("zval"))

	/* save stack, linking frames in reverse order */

	call = EX(call)
	for {
		var frame_size int = ZEND_CALL_FRAME_SLOT + ZEND_CALL_NUM_ARGS(call)
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
	execute_data.SetCall(nil)
	ZEND_ASSERT(prev_call == (*ZendExecuteData)(stack))
	return prev_call
}
func ZendGeneratorCleanupUnfinishedExecution(generator *ZendGenerator, execute_data *ZendExecuteData, catch_op_num uint32) {
	var op_array *ZendOpArray = execute_data.GetFunc().GetOpArray()
	if execute_data.GetOpline() != op_array.GetOpcodes() {

		/* -1 required because we want the last run opcode, not the next to-be-run one. */

		var op_num uint32 = execute_data.GetOpline() - op_array.GetOpcodes() - 1
		if generator.GetFrozenCallStack() != nil {

			/* Temporarily restore generator->execute_data if it has been NULLed out already. */

			var save_ex *ZendExecuteData = generator.GetExecuteData()
			generator.SetExecuteData(execute_data)
			ZendGeneratorRestoreCallStack(generator)
			generator.SetExecuteData(save_ex)
		}
		ZendCleanupUnfinishedExecution(execute_data, op_num, catch_op_num)
	}
}
func ZendGeneratorClose(generator *ZendGenerator, finished_execution ZendBool) {
	if generator.GetExecuteData() != nil {
		var execute_data *ZendExecuteData = generator.GetExecuteData()

		/* Null out execute_data early, to prevent double frees if GC runs while we're
		 * already cleaning up execute_data. */

		generator.SetExecuteData(nil)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			ZendCleanAndCacheSymbolTable(execute_data.GetSymbolTable())
		}

		/* always free the CV's, in the symtable are only not-free'd IS_INDIRECT's */

		ZendFreeCompiledVariables(execute_data)
		if (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0 {
			OBJ_RELEASE(execute_data.GetThis().GetObj())
		}

		/* A fatal error / die occurred during the generator execution.
		 * Trying to clean up the stack may not be safe in this case. */

		if __CG().GetUncleanShutdown() != 0 {
			generator.SetExecuteData(nil)
			return
		}
		ZendVmStackFreeExtraArgs(execute_data)

		/* Some cleanups are only necessary if the generator was closed
		 * before it could finish execution (reach a return statement). */

		if finished_execution == 0 {
			ZendGeneratorCleanupUnfinishedExecution(generator, execute_data, 0)
		}

		/* Free closure object */

		if (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0 {
			OBJ_RELEASE(ZEND_CLOSURE_OBJECT(EX(func_)))
		}

		/* Free GC buffer. GC for closed generators doesn't need an allocated buffer */

		if generator.GetGcBuffer() != nil {
			Efree(generator.GetGcBuffer())
			generator.SetGcBuffer(nil)
		}
		Efree(execute_data)
	}
}
func ZendGeneratorDtorStorage(object *ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	var ex *ZendExecuteData = generator.GetExecuteData()
	var op_num uint32
	var try_catch_offset uint32
	var i int

	/* leave yield from mode to properly allow finally execution */

	if generator.GetValues().GetType() != IS_UNDEF {
		ZvalPtrDtor(generator.GetValues())
		ZVAL_UNDEF(generator.GetValues())
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
	if ex == nil || !ex.GetFunc().GetOpArray().IsHasFinallyBlock() || __CG().GetUncleanShutdown() != 0 {
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

			var fast_call *Zval = ZEND_CALL_VAR(ex, ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar())
			ZendGeneratorCleanupUnfinishedExecution(generator, ex, try_catch.GetFinallyOp())
			fast_call.SetObj(__EG().GetException())
			__EG().SetException(nil)
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
			var fast_call *Zval = ZEND_CALL_VAR(ex, ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar())

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
func ZendGeneratorFreeStorage(object *ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	ZendGeneratorClose(generator, 0)

	/* we can't immediately free them in zend_generator_close() else yield from won't be able to fetch it */

	ZvalPtrDtor(generator.GetValue())
	ZvalPtrDtor(generator.GetKey())
	if !(Z_ISUNDEF(generator.GetRetval())) {
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
		var execute_data *ZendExecuteData = generator.GetExecuteData()
		var op_array *ZendOpArray = EX(func_).op_array

		/* Compiled variables */

		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			size += op_array.GetLastVar()
		}

		/* Extra args */

		if (EX_CALL_INFO() & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
			size += EX_NUM_ARGS() - op_array.GetNumArgs()
		}
		size += (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0
		size += (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0

		/* Live vars */

		if execute_data.GetOpline() != op_array.GetOpcodes() {

			/* -1 required because we want the last run opcode, not the next to-be-run one. */

			var i uint32
			var op_num uint32 = execute_data.GetOpline() - op_array.GetOpcodes() - 1
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
func ZendGeneratorGetGc(object *Zval, table **Zval, n *int) *HashTable {
	var generator *ZendGenerator = (*ZendGenerator)(object.GetObj())
	var execute_data *ZendExecuteData = generator.GetExecuteData()
	var op_array *ZendOpArray
	var gc_buffer *Zval
	var gc_buffer_size uint32
	if execute_data == nil {

		/* If the generator has been closed, it can only hold on to three values: The value, key
		 * and retval. These three zvals are stored sequentially starting at &generator->value. */

		*table = generator.GetValue()
		*n = 3
		return nil
	}
	if generator.IsCurrentlyRunning() {

		/* If the generator is currently running, we certainly won't be able to GC any values it
		 * holds on to. The execute_data state might be inconsistent during execution (e.g. because
		 * GC has been triggered in the middle of a variable reassignment), so we should not try
		 * to inspect it here. */

		*table = nil
		*n = 0
		return nil
	}
	op_array = EX(func_).op_array
	gc_buffer_size = CalcGcBufferSize(generator)
	if generator.GetGcBufferSize() < gc_buffer_size {
		generator.SetGcBuffer(SafeErealloc(generator.GetGcBuffer(), b.SizeOf("zval"), gc_buffer_size, 0))
		generator.SetGcBufferSize(gc_buffer_size)
	}
	*n = gc_buffer_size
	gc_buffer = generator.GetGcBuffer()
	*table = gc_buffer
	ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetValue())
	ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetKey())
	ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetRetval())
	ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), generator.GetValues())
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
		var i uint32
		var num_cvs uint32 = EX(func_).op_array.last_var
		for i = 0; i < num_cvs; i++ {
			ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), EX_VAR_NUM(i))
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
		var zv *Zval = EX_VAR_NUM(op_array.GetLastVar() + op_array.GetT())
		var end *Zval = zv + (EX_NUM_ARGS() - op_array.GetNumArgs())
		for zv != end {
			ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), b.PostInc(&zv))
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_RELEASE_THIS) != 0 {
		ZVAL_OBJ(b.PostInc(&gc_buffer), execute_data.GetThis().GetObj())
	}
	if (EX_CALL_INFO() & ZEND_CALL_CLOSURE) != 0 {
		ZVAL_OBJ(b.PostInc(&gc_buffer), ZEND_CLOSURE_OBJECT(EX(func_)))
	}
	if execute_data.GetOpline() != op_array.GetOpcodes() {
		var i uint32
		var op_num uint32 = execute_data.GetOpline() - op_array.GetOpcodes() - 1
		for i = 0; i < op_array.GetLastLiveRange(); i++ {
			var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
			if range_.GetStart() > op_num {
				break
			} else if op_num < range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *Zval = EX_VAR(var_num)
				if kind == ZEND_LIVE_TMPVAR || kind == ZEND_LIVE_LOOP {
					ZVAL_COPY_VALUE(b.PostInc(&gc_buffer), var_)
				}
			}
		}
	}
	if generator.GetNode().GetChildren() == 0 {
		var root *ZendGenerator = generator.GetNode().GetRoot()
		for root != generator {
			ZVAL_OBJ(b.PostInc(&gc_buffer), root.GetStd())
			root = ZendGeneratorGetChild(root.GetNode(), generator)
		}
	}
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		return execute_data.GetSymbolTable()
	} else {
		return nil
	}
}
func ZendGeneratorCreate(class_type *ZendClassEntry) *ZendObject {
	var generator *ZendGenerator
	generator = Emalloc(b.SizeOf("zend_generator"))
	memset(generator, 0, b.SizeOf("zend_generator"))

	/* The key will be incremented on first use, so it'll start at 0 */

	generator.SetLargestUsedIntegerKey(-1)
	ZVAL_UNDEF(generator.GetRetval())
	ZVAL_UNDEF(generator.GetValues())

	/* By default we have a tree of only one node */

	generator.GetNode().SetParent(nil)
	generator.GetNode().SetChildren(0)
	generator.GetNode().SetRoot(generator)
	ZendObjectStdInit(generator.GetStd(), class_type)
	generator.GetStd().SetHandlers(&ZendGeneratorHandlers)
	return (*ZendObject)(generator)
}
func ZendGeneratorGetConstructor(object *ZendObject) *ZendFunction {
	ZendThrowError(nil, "The \"Generator\" class is reserved for internal use and cannot be manually instantiated")
	return nil
}
func ZendGeneratorCheckPlaceholderFrame(ptr *ZendExecuteData) *ZendExecuteData {
	if ptr.GetFunc() == nil && ptr.GetThis().IsType(IS_OBJECT) {
		if Z_OBJCE(ptr.GetThis()) == ZendCeGenerator {
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
func ZendGeneratorThrowException(generator *ZendGenerator, exception *Zval) {
	var original_execute_data *ZendExecuteData = __EG().GetCurrentExecuteData()

	/* if we don't stop an array/iterator yield from, the exception will only reach the generator after the values were all iterated over */

	if generator.GetValues().GetType() != IS_UNDEF {
		ZvalPtrDtor(generator.GetValues())
		ZVAL_UNDEF(generator.GetValues())
	}

	/* Throw the exception in the context of the generator. Decrementing the opline
	 * to pretend the exception happened during the YIELD opcode. */

	__EG().SetCurrentExecuteData(generator.GetExecuteData())
	generator.GetExecuteData().GetOpline()--
	if exception != nil {
		ZendThrowExceptionObject(exception)
	} else {
		ZendRethrowException(__EG().GetCurrentExecuteData())
	}
	generator.GetExecuteData().GetOpline()++
	__EG().SetCurrentExecuteData(original_execute_data)
}
func ZendGeneratorGetChild(node *ZendGeneratorNode, leaf *ZendGenerator) *ZendGenerator {
	if node.GetChildren() == 0 {
		return nil
	} else if node.GetChildren() == 1 {
		return node.GetChildSingleChild()
	} else {
		return ZendHashIndexFindPtr(node.GetHt(), ZendUlong(leaf))
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
			var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
			ZendHashInit(ht, 0, nil, nil, 0)
			ZendHashIndexAddPtr(ht, ZendUlong(node.GetChildSingleLeaf()), node.GetChildSingleChild())
			node.SetHt(ht)
		}
		ZendHashIndexAddPtr(node.GetHt(), ZendUlong(leaf), child)
	}
	node.GetChildren()++
}
func ZendGeneratorMergeChildNodes(dest *ZendGeneratorNode, src *ZendGeneratorNode, child *ZendGenerator) {
	var leaf ZendUlong
	ZEND_ASSERT(src.GetChildren() > 1)
	var __ht *HashTable = src.GetHt()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		leaf = _p.GetH()
		ZendGeneratorAddSingleChild(dest, child, (*ZendGenerator)(leaf))
	}
}
func ZendGeneratorAddChild(generator *ZendGenerator, child *ZendGenerator) {
	var leaf *ZendGenerator = b.CondF1(child.GetNode().GetChildren() != 0, func() *ZendGenerator { return child.GetNode().GetPtrLeaf() }, child)
	var multi_children_node *ZendGeneratorNode
	var was_leaf ZendBool = generator.GetNode().GetChildren() == 0
	if was_leaf != 0 {
		var next *ZendGenerator = generator.GetNode().GetParent()
		leaf.GetNode().SetRoot(generator.GetNode().GetRoot())
		generator.GetStd().AddRefcount()
		generator.GetNode().SetPtrLeaf(leaf)
		for next != nil {
			if next.GetNode().GetChildren() > 1 {
				var child *ZendGenerator = ZendHashIndexFindPtr(next.GetNode().GetHt(), ZendUlong(generator))
				ZendHashIndexDel(next.GetNode().GetHt(), ZendUlong(generator))
				ZendHashIndexAddPtr(next.GetNode().GetHt(), ZendUlong(leaf), child)
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
			if __EG().GetException() == nil {
				var yield_from *ZendOp = (*ZendOp)(root.GetExecuteData().GetOpline() - 1)
				if yield_from.GetOpcode() == ZEND_YIELD_FROM {
					if Z_ISUNDEF(root.GetNode().GetParent().GetRetval()) {

						/* Throw the exception in the context of the generator */

						var original_execute_data *ZendExecuteData = __EG().GetCurrentExecuteData()
						__EG().SetCurrentExecuteData(root.GetExecuteData())
						if root == generator {
							root.GetExecuteData().SetPrevExecuteData(original_execute_data)
						} else {
							root.GetExecuteData().SetPrevExecuteData(generator.GetExecuteFake())
							generator.GetExecuteFake().SetPrevExecuteData(original_execute_data)
						}
						root.GetExecuteData().GetOpline()--
						ZendThrowException(zend_ce_ClosedGeneratorException, "Generator yielded from aborted, no return value available", 0)
						__EG().SetCurrentExecuteData(original_execute_data)
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
						ZVAL_COPY(root.GetValue(), root.GetNode().GetParent().GetValue())
						ZVAL_COPY(ZEND_CALL_VAR(root.GetExecuteData(), yield_from.GetResult().GetVar()), root.GetNode().GetParent().GetRetval())
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
	var value *Zval
	if generator.GetValues().IsType(IS_ARRAY) {
		var ht *HashTable = generator.GetValues().GetArr()
		var pos HashPosition = generator.GetValues().GetFePos()
		var p *Bucket
		for {
			if pos >= ht.GetNNumUsed() {

				/* Reached end of array */

				goto failure

				/* Reached end of array */

			}
			p = ht.GetArData()[pos]
			value = p.GetVal()
			if value.IsType(IS_INDIRECT) {
				value = value.GetZv()
			}
			pos++
			if !(Z_ISUNDEF_P(value)) {
				break
			}
		}
		ZvalPtrDtor(generator.GetValue())
		ZVAL_COPY(generator.GetValue(), value)
		ZvalPtrDtor(generator.GetKey())
		if p.GetKey() != nil {
			ZVAL_STR_COPY(generator.GetKey(), p.GetKey())
		} else {
			ZVAL_LONG(generator.GetKey(), p.GetH())
		}
		generator.GetValues().GetFePos() = pos
	} else {
		var iter *ZendObjectIterator = (*ZendObjectIterator)(generator.GetValues().GetObj())
		if b.PostInc(&(iter.GetIndex())) > 0 {
			iter.GetFuncs().GetMoveForward()(iter)
			if __EG().GetException() != nil {
				goto exception
			}
		}
		if iter.GetFuncs().GetValid()(iter) == FAILURE {
			if __EG().GetException() != nil {
				goto exception
			}

			/* reached end of iteration */

			goto failure

			/* reached end of iteration */

		}
		value = iter.GetFuncs().GetGetCurrentData()(iter)
		if __EG().GetException() != nil {
			goto exception
		} else if value == nil {
			goto failure
		}
		ZvalPtrDtor(generator.GetValue())
		ZVAL_COPY(generator.GetValue(), value)
		ZvalPtrDtor(generator.GetKey())
		if iter.GetFuncs().GetGetCurrentKey() != nil {
			iter.GetFuncs().GetGetCurrentKey()(iter, generator.GetKey())
			if __EG().GetException() != nil {
				ZVAL_UNDEF(generator.GetKey())
				goto exception
			}
		} else {
			ZVAL_LONG(generator.GetKey(), iter.GetIndex())
		}
	}
	return SUCCESS
exception:
	ZendGeneratorThrowException(generator, nil)
failure:
	ZvalPtrDtor(generator.GetValues())
	ZVAL_UNDEF(generator.GetValues())
	return FAILURE
}
func ZendGeneratorResume(orig_generator *ZendGenerator) {
	var generator *ZendGenerator = ZendGeneratorGetCurrent(orig_generator)

	/* The generator is already closed, thus can't resume */

	if generator.GetExecuteData() == nil {
		return
	}
try_again:
	if generator.IsCurrentlyRunning() {
		ZendThrowError(nil, "Cannot resume an already running generator")
		return
	}
	if orig_generator.IsDoInit() && !(Z_ISUNDEF(generator.GetValue())) {

		/* We must not advance Generator if we yield from a Generator being currently run */

		orig_generator.SetIsDoInit(false)
		return
	}
	if !(Z_ISUNDEF(generator.GetValues())) {
		if ZendGeneratorGetNextDelegatedValue(generator) == SUCCESS {
			orig_generator.SetIsDoInit(false)
			return
		}
	}

	/* Drop the AT_FIRST_YIELD flag */

	orig_generator.SetIsAtFirstYield(false)

	/* Backup executor globals */

	var original_execute_data *ZendExecuteData = __EG().GetCurrentExecuteData()

	/* Set executor globals */

	__EG().SetCurrentExecuteData(generator.GetExecuteData())

	/* We want the backtrace to look as if the generator function was
	 * called from whatever method we are current running (e.g. next()).
	 * So we have to link generator call frame with caller call frame. */

	if generator == orig_generator {
		generator.GetExecuteData().SetPrevExecuteData(original_execute_data)
	} else {

		/* We need some execute_data placeholder in stacktrace to be replaced
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

	__EG().SetCurrentExecuteData(original_execute_data)

	/* If an exception was thrown in the generator we have to internally
	 * rethrow it in the parent scope.
	 * In case we did yield from, the Exception must be rethrown into
	 * its calling frame (see above in if (check_yield_from). */

	if __EG().GetException() != nil {
		if generator == orig_generator {
			ZendGeneratorClose(generator, 0)
			if __EG().GetCurrentExecuteData() == nil {
				ZendThrowExceptionInternal(nil)
			} else if __EG().GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(__EG().GetCurrentExecuteData().GetFunc().GetCommonType()) {
				ZendRethrowException(__EG().GetCurrentExecuteData())
			}
		} else {
			generator = ZendGeneratorGetCurrent(orig_generator)
			ZendGeneratorThrowException(generator, nil)
			orig_generator.SetIsDoInit(false)
			goto try_again
		}
	}

	/* yield from was used, try another resume. */

	if generator != orig_generator && !(Z_ISUNDEF(generator.GetRetval())) || generator.GetExecuteData() != nil && (generator.GetExecuteData().GetOpline()-1).opcode == ZEND_YIELD_FROM {
		generator = ZendGeneratorGetCurrent(orig_generator)
		goto try_again
	}

	/* yield from was used, try another resume. */

	orig_generator.SetIsDoInit(false)
}
func ZendGeneratorEnsureInitialized(generator *ZendGenerator) {
	if generator.GetValue().IsType(IS_UNDEF) && generator.GetExecuteData() != nil && generator.GetNode().GetParent() == nil {
		ZendGeneratorResume(generator)
		generator.SetIsAtFirstYield(true)
	}
}
func ZendGeneratorRewind(generator *ZendGenerator) {
	ZendGeneratorEnsureInitialized(generator)
	if !generator.IsAtFirstYield() {
		ZendThrowException(nil, "Cannot rewind a generator that was already run", 0)
	}
}
func zim_Generator_rewind(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorRewind(generator)
}
func zim_Generator_valid(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorGetCurrent(generator)
	RETVAL_BOOL(generator.GetExecuteData() != nil)
	return
}
func zim_Generator_current(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetValue().GetType() != IS_UNDEF {
		var value *Zval = root.GetValue()
		ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_Generator_key(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetKey().GetType() != IS_UNDEF {
		var key *Zval = root.GetKey()
		ZVAL_COPY_DEREF(return_value, key)
	}
}
func zim_Generator_next(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorResume(generator)
}
func zim_Generator_send(execute_data *ZendExecuteData, return_value *Zval) {
	var value *Zval
	var generator *ZendGenerator
	var root *ZendGenerator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &value, 0)
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)

	/* The generator is already closed, thus can't send anything */

	if generator.GetExecuteData() == nil {
		return
	}
	root = ZendGeneratorGetCurrent(generator)

	/* Put sent value in the target VAR slot, if it is used */

	if root.GetSendTarget() != nil {
		ZVAL_COPY(root.GetSendTarget(), value)
	}
	ZendGeneratorResume(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		var value *Zval = root.GetValue()
		ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_Generator_throw(execute_data *ZendExecuteData, return_value *Zval) {
	var exception *Zval
	var generator *ZendGenerator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &exception, 0)
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	Z_TRY_ADDREF_P(exception)
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if generator.GetExecuteData() != nil {
		var root *ZendGenerator = ZendGeneratorGetCurrent(generator)
		ZendGeneratorThrowException(root, exception)
		ZendGeneratorResume(generator)
		root = ZendGeneratorGetCurrent(generator)
		if generator.GetExecuteData() != nil {
			var value *Zval = root.GetValue()
			ZVAL_COPY_DEREF(return_value, value)
		}
	} else {

		/* If the generator is already closed throw the exception in the
		 * current context */

		ZendThrowExceptionObject(exception)

		/* If the generator is already closed throw the exception in the
		 * current context */

	}
}
func zim_Generator_getReturn(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if ZendParseParametersNone() == FAILURE {
		return
	}
	generator = (*ZendGenerator)(ZEND_THIS.GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if __EG().GetException() != nil {
		return
	}
	if Z_ISUNDEF(generator.GetRetval()) {

		/* Generator hasn't returned yet -> error! */

		ZendThrowException(nil, "Cannot get return value of a generator that hasn't returned", 0)
		return
	}
	ZVAL_COPY(return_value, generator.GetRetval())
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
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ZendGeneratorIteratorGetData(iterator *ZendObjectIterator) *Zval {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	return root.GetValue()
}
func ZendGeneratorIteratorGetKey(iterator *ZendObjectIterator, key *Zval) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if root.GetKey().GetType() != IS_UNDEF {
		var zv *Zval = root.GetKey()
		ZVAL_COPY_DEREF(key, zv)
	} else {
		ZVAL_NULL(key)
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
func ZendGeneratorGetIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendObjectIterator
	var generator *ZendGenerator = (*ZendGenerator)(object.GetObj())
	if generator.GetExecuteData() == nil {
		ZendThrowException(nil, "Cannot traverse an already closed generator", 0)
		return nil
	}
	if by_ref != 0 && !generator.GetExecuteData().GetFunc().GetOpArray().IsReturnReference() {
		ZendThrowException(nil, "You can only iterate a generator by-reference if it declared that it yields by-reference", 0)
		return nil
	}
	generator.SetIterator(Emalloc(b.SizeOf("zend_object_iterator")))
	iterator = generator.GetIterator()
	ZendIteratorInit(iterator)
	iterator.SetFuncs(&ZendGeneratorIteratorFunctions)
	Z_ADDREF_P(object)
	ZVAL_OBJ(iterator.GetData(), object.GetObj())
	return iterator
}
func ZendRegisterGeneratorCe() {
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Generator", b.SizeOf("\"Generator\"")-1, 1))
	ce.SetBuiltinFunctions(GeneratorFunctions)
	ZendCeGenerator = ZendRegisterInternalClass(&ce)
	ZendCeGenerator.SetIsFinal(true)
	ZendCeGenerator.create_object = ZendGeneratorCreate
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
	ce.SetName(ZendStringInitInterned("ClosedGeneratorException", b.SizeOf("\"ClosedGeneratorException\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	zend_ce_ClosedGeneratorException = ZendRegisterInternalClassEx(&ce, ZendCeException)
}
