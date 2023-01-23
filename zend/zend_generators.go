// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_generators.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Nikita Popov <nikic@php.net>                                |
   |          Bob Weinand <bobwei9@hotmail.com>                           |
   +----------------------------------------------------------------------+
*/

// #define ZEND_GENERATORS_H

var ZendCeGenerator *ZendClassEntry
var zend_ce_ClosedGeneratorException *ZendClassEntry

/* The concept of `yield from` exposes problems when accessed at different levels of the chain of delegated generators. We need to be able to reference the currently executed Generator in all cases and still being able to access the return values of finished Generators.
 * The solution to this problem is a doubly-linked tree, which all Generators referenced in maintain a reference to. It should be impossible to avoid walking the tree in all cases. This way, we only need tree walks from leaf to root in case where some part of the `yield from` chain is passed to another `yield from`. (Update of leaf node pointer and list of multi-children nodes needed when leaf gets a child in direct path from leaf to root node.) But only in that case, which should be a fairly rare case (which is then possible, but not totally cheap).
 * The root of the tree is then the currently executed Generator. The subnodes of the tree (all except the root node) are all Generators which do `yield from`. Each node of the tree knows a pointer to one leaf descendant node. Each node with multiple children needs a list of all leaf descendant nodes paired with pointers to their respective child node. (The stack is determined by leaf node pointers) Nodes with only one child just don't need a list, there it is enough to just have a pointer to the child node. Further, leaf nodes store a pointer to the root node.
 * That way, when we advance any generator, we just need to look up a leaf node (which all have a reference to a root node). Then we can see at the root node whether current Generator is finished. If it isn't, all is fine and we can just continue. If the Generator finished, there will be two cases. Either it is a simple node with just one child, then go down to child node. Or it has multiple children and we now will remove the current leaf node from the list of nodes (unnecessary, is microoptimization) and go down to the child node whose reference was paired with current leaf node. Child node is then removed its parent reference and becomes new top node. Or the current node references the Generator we're currently executing, then we can continue from the YIELD_FROM opcode. When a node referenced as root node in a leaf node has a parent, then we go the way up until we find a root node without parent.
 * In case we go into a new `yield from` level, a node is created on top of current root and becomes the new root. Leaf node needs to be updated with new root node then.
 * When a Generator referenced by a node of the tree is added to `yield from`, that node now gets a list of children (we need to walk the descendants of that node and nodes of the tree of the other Generator down to the first multi-children node and copy all the leaf node pointers from there). In case there was no multi-children node (linear tree), we just add a pair (pointer to leaf node, pointer to child node), with the child node being in a direct path from leaf to this node.
 */

// @type ZendGeneratorNode struct
// @type ZendGenerator struct
var ZEND_GENERATOR_CURRENTLY_RUNNING ZendUchar = 0x1
var ZEND_GENERATOR_FORCED_CLOSE ZendUchar = 0x2
var ZEND_GENERATOR_AT_FIRST_YIELD ZendUchar = 0x4
var ZEND_GENERATOR_DO_INIT ZendUchar = 0x8

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

// Source: <Zend/zend_generators.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Nikita Popov <nikic@php.net>                                |
   |          Bob Weinand <bobwei9@hotmail.com>                           |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "zend_generators.h"

// # include "zend_closures.h"

var ZendGeneratorHandlers ZendObjectHandlers

func ZendGeneratorRestoreCallStack(generator *ZendGenerator) {
	var call *ZendExecuteData
	var new_call *ZendExecuteData
	var prev_call *ZendExecuteData = nil
	call = generator.GetFrozenCallStack()
	for {
		new_call = ZendVmStackPushCallFrame(call.GetThis().GetTypeInfo() & ^(1<<18), call.GetFunc(), call.GetThis().GetNumArgs(), call.GetThis().GetValue().GetPtr())
		memcpy((*Zval)(new_call)+int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))), (*Zval)(call)+int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))), call.GetThis().GetNumArgs()*g.SizeOf("zval"))
		new_call.SetPrevExecuteData(prev_call)
		prev_call = new_call
		call = call.GetPrevExecuteData()
		if call == nil {
			break
		}
	}
	generator.GetExecuteData().SetCall(prev_call)
	_efree(generator.GetFrozenCallStack())
	generator.SetFrozenCallStack(nil)
}

/* }}} */

func ZendGeneratorFreezeCallStack(execute_data *ZendExecuteData) *ZendExecuteData {
	var used_stack int
	var call *ZendExecuteData
	var new_call *ZendExecuteData
	var prev_call *ZendExecuteData = nil
	var stack *Zval

	/* calculate required stack size */

	used_stack = 0
	call = execute_data.GetCall()
	for {
		used_stack += int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + call.GetThis().GetNumArgs()
		call = call.GetPrevExecuteData()
		if call == nil {
			break
		}
	}
	stack = _emalloc(used_stack * g.SizeOf("zval"))

	/* save stack, linking frames in reverse order */

	call = execute_data.GetCall()
	for {
		var frame_size int = int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + call.GetThis().GetNumArgs()
		new_call = (*ZendExecuteData)(stack + used_stack - frame_size)
		memcpy(new_call, call, frame_size*g.SizeOf("zval"))
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
	r.Assert(prev_call == (*ZendExecuteData)(stack))
	return prev_call
}

/* }}} */

func ZendGeneratorCleanupUnfinishedExecution(generator *ZendGenerator, execute_data *ZendExecuteData, catch_op_num uint32) {
	var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
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

/* }}} */

func ZendGeneratorClose(generator *ZendGenerator, finished_execution ZendBool) {
	if generator.GetExecuteData() != nil {
		var execute_data *ZendExecuteData = generator.GetExecuteData()

		/* Null out execute_data early, to prevent double frees if GC runs while we're
		 * already cleaning up execute_data. */

		generator.SetExecuteData(nil)
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) != 0 {
			ZendCleanAndCacheSymbolTable(execute_data.GetSymbolTable())
		}

		/* always free the CV's, in the symtable are only not-free'd IS_INDIRECT's */

		ZendFreeCompiledVariables(execute_data)
		if (execute_data.GetThis().GetTypeInfo() & 1 << 21) != 0 {
			ZendObjectRelease(execute_data.GetThis().GetValue().GetObj())
		}

		/* A fatal error / die occurred during the generator execution.
		 * Trying to clean up the stack may not be safe in this case. */

		if CG.GetUncleanShutdown() != 0 {
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

		if (execute_data.GetThis().GetTypeInfo() & 1 << 22) != 0 {
			ZendObjectRelease((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
		}

		/* Free GC buffer. GC for closed generators doesn't need an allocated buffer */

		if generator.GetGcBuffer() != nil {
			_efree(generator.GetGcBuffer())
			generator.SetGcBuffer(nil)
		}
		_efree(execute_data)
	}
}

/* }}} */

func ZendGeneratorDtorStorage(object *ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	var ex *ZendExecuteData = generator.GetExecuteData()
	var op_num uint32
	var try_catch_offset uint32
	var i int

	/* leave yield from mode to properly allow finally execution */

	if generator.GetValues().GetType() != 0 {
		ZvalPtrDtor(&generator.values)
		&generator.values.u1.type_info = 0
	}
	if generator.GetNode().GetChildren() == 0 {
		var root *ZendGenerator = generator.GetNode().GetRoot()
		var next *ZendGenerator
		for root != generator {
			next = ZendGeneratorGetChild(&root.node, generator)
			generator.GetNode().SetRoot(next)
			next.GetNode().SetParent(nil)
			ZendObjectRelease(&root.std)
			root = next
		}
	}
	if ex == nil || (ex.GetFunc().GetOpArray().GetFnFlags()&1<<15) == 0 || CG.GetUncleanShutdown() != 0 {
		return
	}

	/* -1 required because we want the last run opcode, not the
	 * next to-be-run one. */

	op_num = ex.GetOpline() - ex.GetFunc().GetOpArray().GetOpcodes() - 1
	try_catch_offset = -1

	/* Find the innermost try/catch that we are inside of. */

	for i = 0; i < ex.GetFunc().GetOpArray().GetLastTryCatch(); i++ {
		var try_catch *ZendTryCatchElement = &ex.func_.GetOpArray().GetTryCatchArray()[i]
		if op_num < try_catch.GetTryOp() {
			break
		}
		if op_num < try_catch.GetCatchOp() || op_num < try_catch.GetFinallyEnd() {
			try_catch_offset = i
		}
	}

	/* Walk try/catch/finally structures upwards, performing the necessary actions. */

	for try_catch_offset != uint32-1 {
		var try_catch *ZendTryCatchElement = &ex.func_.GetOpArray().GetTryCatchArray()[try_catch_offset]
		if op_num < try_catch.GetFinallyOp() {

			/* Go to finally block */

			var fast_call *Zval = (*Zval)((*byte)(ex) + int(ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar()))
			ZendGeneratorCleanupUnfinishedExecution(generator, ex, try_catch.GetFinallyOp())
			fast_call.GetValue().SetObj(EG.GetException())
			EG.SetException(nil)
			fast_call.SetOplineNum(uint32 - 1)
			ex.SetOpline(&ex.func_.GetOpArray().GetOpcodes()[try_catch.GetFinallyOp()])
			generator.SetFlags(generator.GetFlags() | ZEND_GENERATOR_FORCED_CLOSE)
			ZendGeneratorResume(generator)

			/* TODO: If we hit another yield inside try/finally,
			 * should we also jump to the next finally block? */

			return

			/* TODO: If we hit another yield inside try/finally,
			 * should we also jump to the next finally block? */

		} else if op_num < try_catch.GetFinallyEnd() {
			var fast_call *Zval = (*Zval)((*byte)(ex) + int(ex.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar()))

			/* Clean up incomplete return statement */

			if fast_call.GetOplineNum() != uint32-1 {
				var retval_op *ZendOp = &ex.func_.GetOpArray().GetOpcodes()[fast_call.GetOplineNum()]
				if (retval_op.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
					ZvalPtrDtor((*Zval)((*byte)(ex) + int(retval_op.GetOp2().GetVar())))
				}
			}

			/* Clean up backed-up exception */

			if fast_call.GetValue().GetObj() != nil {
				ZendObjectRelease(fast_call.GetValue().GetObj())
			}

			/* Clean up backed-up exception */

		}
		try_catch_offset--
	}

	/* Walk try/catch/finally structures upwards, performing the necessary actions. */
}

/* }}} */

func ZendGeneratorFreeStorage(object *ZendObject) {
	var generator *ZendGenerator = (*ZendGenerator)(object)
	ZendGeneratorClose(generator, 0)

	/* we can't immediately free them in zend_generator_close() else yield from won't be able to fetch it */

	ZvalPtrDtor(&generator.value)
	ZvalPtrDtor(&generator.key)
	if generator.GetRetval().GetType() != 0 {
		ZvalPtrDtor(&generator.retval)
	}
	if generator.GetNode().GetChildren() > 1 {
		ZendHashDestroy(generator.GetNode().GetHt())
		_efree(generator.GetNode().GetHt())
	}
	ZendObjectStdDtor(&generator.std)
}

/* }}} */

func CalcGcBufferSize(generator *ZendGenerator) uint32 {
	var size uint32 = 4
	if generator.GetExecuteData() != nil {
		var execute_data *ZendExecuteData = generator.GetExecuteData()
		var op_array *ZendOpArray = &(execute_data.GetFunc()).op_array

		/* Compiled variables */

		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) == 0 {
			size += op_array.GetLastVar()
		}

		/* Extra args */

		if (execute_data.GetThis().GetTypeInfo() & 1 << 19) != 0 {
			size += execute_data.GetThis().GetNumArgs() - op_array.GetNumArgs()
		}
		size += (execute_data.GetThis().GetTypeInfo() & 1 << 21) != 0
		size += (execute_data.GetThis().GetTypeInfo() & 1 << 22) != 0

		/* Live vars */

		if execute_data.GetOpline() != op_array.GetOpcodes() {

			/* -1 required because we want the last run opcode, not the next to-be-run one. */

			var i uint32
			var op_num uint32 = execute_data.GetOpline() - op_array.GetOpcodes() - 1
			for i = 0; i < op_array.GetLastLiveRange(); i++ {
				var range_ *ZendLiveRange = &op_array.live_range[i]
				if range_.GetStart() > op_num {

					/* Further ranges will not be relevant... */

					break

					/* Further ranges will not be relevant... */

				} else if op_num < range_.GetEnd() {

					/* LIVE_ROPE and LIVE_SILENCE not relevant for GC */

					var kind uint32 = range_.GetVar() & 7
					if kind == 0 || kind == 1 {
						size++
					}
				}
			}
		}

		/* Yield from root references */

		if generator.GetNode().GetChildren() == 0 {
			var root *ZendGenerator = generator.GetNode().GetRoot()
			for root != generator {
				root = ZendGeneratorGetChild(&root.node, generator)
				size++
			}
		}

		/* Yield from root references */

	}
	return size
}

/* }}} */

func ZendGeneratorGetGc(object *Zval, table **Zval, n *int) *HashTable {
	var generator *ZendGenerator = (*ZendGenerator)(object.GetValue().GetObj())
	var execute_data *ZendExecuteData = generator.GetExecuteData()
	var op_array *ZendOpArray
	var gc_buffer *Zval
	var gc_buffer_size uint32
	if execute_data == nil {

		/* If the generator has been closed, it can only hold on to three values: The value, key
		 * and retval. These three zvals are stored sequentially starting at &generator->value. */

		*table = &generator.value
		*n = 3
		return nil
	}
	if (generator.GetFlags() & ZEND_GENERATOR_CURRENTLY_RUNNING) != 0 {

		/* If the generator is currently running, we certainly won't be able to GC any values it
		 * holds on to. The execute_data state might be inconsistent during execution (e.g. because
		 * GC has been triggered in the middle of a variable reassignment), so we should not try
		 * to inspect it here. */

		*table = nil
		*n = 0
		return nil
	}
	op_array = &(execute_data.GetFunc()).op_array
	gc_buffer_size = CalcGcBufferSize(generator)
	if generator.GetGcBufferSize() < gc_buffer_size {
		generator.SetGcBuffer(_safeErealloc(generator.GetGcBuffer(), g.SizeOf("zval"), gc_buffer_size, 0))
		generator.SetGcBufferSize(gc_buffer_size)
	}
	*n = gc_buffer_size
	gc_buffer = generator.GetGcBuffer()
	*table = gc_buffer
	var _z1 *Zval = g.PostInc(&gc_buffer)
	var _z2 *Zval = &generator.value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = g.PostInc(&gc_buffer)
	var _z2 *Zval = &generator.key
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = g.PostInc(&gc_buffer)
	var _z2 *Zval = &generator.retval
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = g.PostInc(&gc_buffer)
	var _z2 *Zval = &generator.values
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (execute_data.GetThis().GetTypeInfo() & 1 << 20) == 0 {
		var i uint32
		var num_cvs uint32 = execute_data.GetFunc().GetOpArray().GetLastVar()
		for i = 0; i < num_cvs; i++ {
			var _z1 *Zval = g.PostInc(&gc_buffer)
			var _z2 *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(i))
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	if (execute_data.GetThis().GetTypeInfo() & 1 << 19) != 0 {
		var zv *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+op_array.GetT()))
		var end *Zval = zv + (execute_data.GetThis().GetNumArgs() - op_array.GetNumArgs())
		for zv != end {
			var _z1 *Zval = g.PostInc(&gc_buffer)
			var _z2 *Zval = g.PostInc(&zv)
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	if (execute_data.GetThis().GetTypeInfo() & 1 << 21) != 0 {
		var __z *Zval = g.PostInc(&gc_buffer)
		__z.GetValue().SetObj(execute_data.GetThis().GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	}
	if (execute_data.GetThis().GetTypeInfo() & 1 << 22) != 0 {
		var __z *Zval = g.PostInc(&gc_buffer)
		__z.GetValue().SetObj((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	}
	if execute_data.GetOpline() != op_array.GetOpcodes() {
		var i uint32
		var op_num uint32 = execute_data.GetOpline() - op_array.GetOpcodes() - 1
		for i = 0; i < op_array.GetLastLiveRange(); i++ {
			var range_ *ZendLiveRange = &op_array.live_range[i]
			if range_.GetStart() > op_num {
				break
			} else if op_num < range_.GetEnd() {
				var kind uint32 = range_.GetVar() & 7
				var var_num uint32 = range_.GetVar() & ^7
				var var_ *Zval = (*Zval)((*byte)(execute_data) + int(var_num))
				if kind == 0 || kind == 1 {
					var _z1 *Zval = g.PostInc(&gc_buffer)
					var _z2 *Zval = var_
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				}
			}
		}
	}
	if generator.GetNode().GetChildren() == 0 {
		var root *ZendGenerator = generator.GetNode().GetRoot()
		for root != generator {
			var __z *Zval = g.PostInc(&gc_buffer)
			__z.GetValue().SetObj(&root.std)
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			root = ZendGeneratorGetChild(&root.node, generator)
		}
	}
	if (execute_data.GetThis().GetTypeInfo() & 1 << 20) != 0 {
		return execute_data.GetSymbolTable()
	} else {
		return nil
	}
}

/* }}} */

func ZendGeneratorCreate(class_type *ZendClassEntry) *ZendObject {
	var generator *ZendGenerator
	generator = _emalloc(g.SizeOf("zend_generator"))
	memset(generator, 0, g.SizeOf("zend_generator"))

	/* The key will be incremented on first use, so it'll start at 0 */

	generator.SetLargestUsedIntegerKey(-1)
	&generator.retval.u1.type_info = 0
	&generator.values.u1.type_info = 0

	/* By default we have a tree of only one node */

	generator.GetNode().SetParent(nil)
	generator.GetNode().SetChildren(0)
	generator.GetNode().SetRoot(generator)
	ZendObjectStdInit(&generator.std, class_type)
	generator.GetStd().SetHandlers(&ZendGeneratorHandlers)
	return (*ZendObject)(generator)
}

/* }}} */

func ZendGeneratorGetConstructor(object *ZendObject) *ZendFunction {
	ZendThrowError(nil, "The \"Generator\" class is reserved for internal use and cannot be manually instantiated")
	return nil
}

/* }}} */

func ZendGeneratorCheckPlaceholderFrame(ptr *ZendExecuteData) *ZendExecuteData {
	if ptr.GetFunc() == nil && ptr.GetThis().GetType() == 8 {
		if ptr.GetThis().GetValue().GetObj().GetCe() == ZendCeGenerator {
			var generator *ZendGenerator = (*ZendGenerator)(ptr.GetThis().GetValue().GetObj())
			var root *ZendGenerator = g.CondF2(generator.GetNode().GetChildren() < 1, generator, func() *ZendGenerator { return generator.GetNode().GetPtrLeaf() }).node.ptr.root
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
	var original_execute_data *ZendExecuteData = EG.GetCurrentExecuteData()

	/* if we don't stop an array/iterator yield from, the exception will only reach the generator after the values were all iterated over */

	if generator.GetValues().GetType() != 0 {
		ZvalPtrDtor(&generator.values)
		&generator.values.u1.type_info = 0
	}

	/* Throw the exception in the context of the generator. Decrementing the opline
	 * to pretend the exception happened during the YIELD opcode. */

	EG.SetCurrentExecuteData(generator.GetExecuteData())
	generator.GetExecuteData().GetOpline()--
	if exception != nil {
		ZendThrowExceptionObject(exception)
	} else {
		ZendRethrowException(EG.GetCurrentExecuteData())
	}
	generator.GetExecuteData().GetOpline()++
	EG.SetCurrentExecuteData(original_execute_data)
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
		node = &node.child.single.child.GetNode()
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
			var ht *HashTable = _emalloc(g.SizeOf("HashTable"))
			_zendHashInit(ht, 0, nil, 0)
			ZendHashIndexAddPtr(ht, ZendUlong(node.GetChildSingleLeaf()), node.GetChildSingleChild())
			node.SetHt(ht)
		}
		ZendHashIndexAddPtr(node.GetHt(), ZendUlong(leaf), child)
	}
	node.GetChildren()++
}
func ZendGeneratorMergeChildNodes(dest *ZendGeneratorNode, src *ZendGeneratorNode, child *ZendGenerator) {
	var leaf ZendUlong
	r.Assert(src.GetChildren() > 1)
	for {
		var __ht *HashTable = src.GetHt()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			leaf = _p.GetH()
			ZendGeneratorAddSingleChild(dest, child, (*ZendGenerator)(leaf))
		}
		break
	}
}

/* Pay attention so that the root of each subtree of the Generators tree is referenced
 * once per leaf */

func ZendGeneratorAddChild(generator *ZendGenerator, child *ZendGenerator) {
	var leaf *ZendGenerator = g.CondF1(child.GetNode().GetChildren() != 0, func() *ZendGenerator { return child.GetNode().GetPtrLeaf() }, child)
	var multi_children_node *ZendGeneratorNode
	var was_leaf ZendBool = generator.GetNode().GetChildren() == 0
	if was_leaf != 0 {
		var next *ZendGenerator = generator.GetNode().GetParent()
		leaf.GetNode().SetRoot(generator.GetNode().GetRoot())
		ZendGcAddref(&(&generator.std).GetGc())
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
		multi_children_node = ZendGeneratorSearchMultiChildrenNode(&generator.node)
		if multi_children_node != nil {
			generator.GetNode().SetChildren(0)
			ZendGeneratorMergeChildNodes(&generator.node, multi_children_node, generator.GetNode().GetChildSingleChild())
		}
	}
	if was_leaf == 0 {
		multi_children_node = ZendGeneratorSearchMultiChildrenNode(&child.node)
	} else {
		multi_children_node = (*ZendGeneratorNode)(0x1)
	}
	var parent *ZendGenerator = generator.GetNode().GetParent()
	var cur *ZendGenerator = generator
	if multi_children_node > (*ZendGeneratorNode)(0x1) {
		ZendGeneratorMergeChildNodes(&generator.node, multi_children_node, child)
	} else {
		ZendGeneratorAddSingleChild(&generator.node, child, leaf)
	}
	for parent != nil {
		if parent.GetNode().GetChildren() > 1 {
			if multi_children_node == (*ZendGeneratorNode)(0x1) {
				multi_children_node = ZendGeneratorSearchMultiChildrenNode(&child.node)
			}
			if multi_children_node != nil {
				ZendGeneratorMergeChildNodes(&parent.node, multi_children_node, cur)
			} else {
				ZendGeneratorAddSingleChild(&parent.node, cur, leaf)
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
	ZendGcDelref(&(&from.std).GetGc())
	generator.SetFlags(generator.GetFlags() | ZEND_GENERATOR_DO_INIT)
}
func ZendGeneratorUpdateCurrent(generator *ZendGenerator, leaf *ZendGenerator) *ZendGenerator {
	var old_root *ZendGenerator
	var root *ZendGenerator = leaf.GetNode().GetRoot()

	/* generator at the root had stopped */

	if root != generator {
		old_root = root
		root = ZendGeneratorGetChild(&root.node, leaf)
	} else {
		old_root = nil
	}
	for root.GetExecuteData() == nil && root != generator {
		ZendObjectRelease(&old_root.std)
		old_root = root
		root = ZendGeneratorGetChild(&root.node, leaf)
	}
	if root.GetNode().GetParent() != nil {
		if root.GetNode().GetParent().GetExecuteData() == nil {
			if EG.GetException() == nil {
				var yield_from *ZendOp = (*ZendOp)(root.GetExecuteData().GetOpline() - 1)
				if yield_from.GetOpcode() == 166 {
					if root.GetNode().GetParent().GetRetval().GetType() == 0 {

						/* Throw the exception in the context of the generator */

						var original_execute_data *ZendExecuteData = EG.GetCurrentExecuteData()
						EG.SetCurrentExecuteData(root.GetExecuteData())
						if root == generator {
							root.GetExecuteData().SetPrevExecuteData(original_execute_data)
						} else {
							root.GetExecuteData().SetPrevExecuteData(&generator.execute_fake)
							generator.GetExecuteFake().SetPrevExecuteData(original_execute_data)
						}
						root.GetExecuteData().GetOpline()--
						ZendThrowException(zend_ce_ClosedGeneratorException, "Generator yielded from aborted, no return value available", 0)
						EG.SetCurrentExecuteData(original_execute_data)
						if (g.Cond(old_root != nil, old_root, generator).flags & ZEND_GENERATOR_CURRENTLY_RUNNING) == 0 {
							leaf.GetNode().SetRoot(root)
							root.GetNode().SetParent(nil)
							if old_root != nil {
								ZendObjectRelease(&old_root.std)
							}
							ZendGeneratorResume(leaf)
							return leaf.GetNode().GetRoot()
						}
					} else {
						ZvalPtrDtor(&root.value)
						var _z1 *Zval = &root.value
						var _z2 *Zval = &root.node.GetParent().GetValue()
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
						var _z1 *Zval = (*Zval)((*byte)(root.GetExecuteData()) + int(yield_from.GetResult().GetVar()))
						var _z2 *Zval = &root.node.GetParent().GetRetval()
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
				}
			}
			root.GetNode().SetParent(nil)
		} else {
			for {
				root = root.GetNode().GetParent()
				ZendGcAddref(&(&root.std).GetGc())
				if root.GetNode().GetParent() == nil {
					break
				}
			}
		}
	}
	leaf.GetNode().SetRoot(root)
	if old_root != nil {
		ZendObjectRelease(&old_root.std)
	}
	return root
}
func ZendGeneratorGetNextDelegatedValue(generator *ZendGenerator) int {
	var value *Zval
	if generator.GetValues().GetType() == 7 {
		var ht *HashTable = generator.GetValues().GetValue().GetArr()
		var pos HashPosition = generator.GetValues().GetFePos()
		var p *Bucket
		for {
			if pos >= ht.GetNNumUsed() {

				/* Reached end of array */

				goto failure

				/* Reached end of array */

			}
			p = &ht.arData[pos]
			value = &p.val
			if value.GetType() == 13 {
				value = value.GetValue().GetZv()
			}
			pos++
			if value.GetType() != 0 {
				break
			}
		}
		ZvalPtrDtor(&generator.value)
		var _z1 *Zval = &generator.value
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		ZvalPtrDtor(&generator.key)
		if p.GetKey() != nil {
			var __z *Zval = &generator.key
			var __s *ZendString = p.GetKey()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = &generator.key
			__z.GetValue().SetLval(p.GetH())
			__z.SetTypeInfo(4)
		}
		generator.GetValues().SetFePos(pos)
	} else {
		var iter *ZendObjectIterator = (*ZendObjectIterator)(generator.GetValues().GetValue().GetObj())
		if g.PostInc(&(iter.GetIndex())) > 0 {
			iter.GetFuncs().GetMoveForward()(iter)
			if EG.GetException() != nil {
				goto exception
			}
		}
		if iter.GetFuncs().GetValid()(iter) == FAILURE {
			if EG.GetException() != nil {
				goto exception
			}

			/* reached end of iteration */

			goto failure

			/* reached end of iteration */

		}
		value = iter.GetFuncs().GetGetCurrentData()(iter)
		if EG.GetException() != nil {
			goto exception
		} else if value == nil {
			goto failure
		}
		ZvalPtrDtor(&generator.value)
		var _z1 *Zval = &generator.value
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		ZvalPtrDtor(&generator.key)
		if iter.GetFuncs().GetGetCurrentKey() != nil {
			iter.GetFuncs().GetGetCurrentKey()(iter, &generator.key)
			if EG.GetException() != nil {
				&generator.key.u1.type_info = 0
				goto exception
			}
		} else {
			var __z *Zval = &generator.key
			__z.GetValue().SetLval(iter.GetIndex())
			__z.SetTypeInfo(4)
		}
	}
	return SUCCESS
exception:
	ZendGeneratorThrowException(generator, nil)
failure:
	ZvalPtrDtor(&generator.values)
	&generator.values.u1.type_info = 0
	return FAILURE
}

/* }}} */

func ZendGeneratorResume(orig_generator *ZendGenerator) {
	var generator *ZendGenerator = ZendGeneratorGetCurrent(orig_generator)

	/* The generator is already closed, thus can't resume */

	if generator.GetExecuteData() == nil {
		return
	}
try_again:
	if (generator.GetFlags() & ZEND_GENERATOR_CURRENTLY_RUNNING) != 0 {
		ZendThrowError(nil, "Cannot resume an already running generator")
		return
	}
	if (orig_generator.GetFlags()&ZEND_GENERATOR_DO_INIT) != 0 && generator.GetValue().GetType() != 0 {

		/* We must not advance Generator if we yield from a Generator being currently run */

		orig_generator.SetFlags(orig_generator.GetFlags() &^ ZEND_GENERATOR_DO_INIT)
		return
	}
	if generator.GetValues().GetType() != 0 {
		if ZendGeneratorGetNextDelegatedValue(generator) == SUCCESS {
			orig_generator.SetFlags(orig_generator.GetFlags() &^ ZEND_GENERATOR_DO_INIT)
			return
		}
	}

	/* Drop the AT_FIRST_YIELD flag */

	orig_generator.SetFlags(orig_generator.GetFlags() &^ ZEND_GENERATOR_AT_FIRST_YIELD)

	/* Backup executor globals */

	var original_execute_data *ZendExecuteData = EG.GetCurrentExecuteData()

	/* Set executor globals */

	EG.SetCurrentExecuteData(generator.GetExecuteData())

	/* We want the backtrace to look as if the generator function was
	 * called from whatever method we are current running (e.g. next()).
	 * So we have to link generator call frame with caller call frame. */

	if generator == orig_generator {
		generator.GetExecuteData().SetPrevExecuteData(original_execute_data)
	} else {

		/* We need some execute_data placeholder in stacktrace to be replaced
		 * by the real stack trace when needed */

		generator.GetExecuteData().SetPrevExecuteData(&orig_generator.execute_fake)
		orig_generator.GetExecuteFake().SetPrevExecuteData(original_execute_data)
	}
	if generator.GetFrozenCallStack() != nil {

		/* Restore frozen call-stack */

		ZendGeneratorRestoreCallStack(generator)

		/* Restore frozen call-stack */

	}

	/* Resume execution */

	generator.SetFlags(generator.GetFlags() | ZEND_GENERATOR_CURRENTLY_RUNNING)
	ZendExecuteEx(generator.GetExecuteData())
	generator.SetFlags(generator.GetFlags() &^ ZEND_GENERATOR_CURRENTLY_RUNNING)
	generator.SetFrozenCallStack(nil)
	if generator.GetExecuteData() != nil && generator.GetExecuteData().GetCall() != nil {

		/* Frize call-stack */

		generator.SetFrozenCallStack(ZendGeneratorFreezeCallStack(generator.GetExecuteData()))

		/* Frize call-stack */

	}

	/* Restore executor globals */

	EG.SetCurrentExecuteData(original_execute_data)

	/* If an exception was thrown in the generator we have to internally
	 * rethrow it in the parent scope.
	 * In case we did yield from, the Exception must be rethrown into
	 * its calling frame (see above in if (check_yield_from). */

	if EG.GetException() != nil {
		if generator == orig_generator {
			ZendGeneratorClose(generator, 0)
			if EG.GetCurrentExecuteData() == nil {
				ZendThrowExceptionInternal(nil)
			} else if EG.GetCurrentExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetFunc().GetCommonType()&1) == 0 {
				ZendRethrowException(EG.GetCurrentExecuteData())
			}
		} else {
			generator = ZendGeneratorGetCurrent(orig_generator)
			ZendGeneratorThrowException(generator, nil)
			orig_generator.SetFlags(orig_generator.GetFlags() &^ ZEND_GENERATOR_DO_INIT)
			goto try_again
		}
	}

	/* yield from was used, try another resume. */

	if generator != orig_generator && generator.GetRetval().GetType() != 0 || generator.GetExecuteData() != nil && (generator.GetExecuteData().GetOpline()-1).opcode == 166 {
		generator = ZendGeneratorGetCurrent(orig_generator)
		goto try_again
	}

	/* yield from was used, try another resume. */

	orig_generator.SetFlags(orig_generator.GetFlags() &^ ZEND_GENERATOR_DO_INIT)
}

/* }}} */

func ZendGeneratorEnsureInitialized(generator *ZendGenerator) {
	if generator.GetValue().GetType() == 0 && generator.GetExecuteData() != nil && generator.GetNode().GetParent() == nil {
		ZendGeneratorResume(generator)
		generator.SetFlags(generator.GetFlags() | ZEND_GENERATOR_AT_FIRST_YIELD)
	}
}

/* }}} */

func ZendGeneratorRewind(generator *ZendGenerator) {
	ZendGeneratorEnsureInitialized(generator)
	if (generator.GetFlags() & ZEND_GENERATOR_AT_FIRST_YIELD) == 0 {
		ZendThrowException(nil, "Cannot rewind a generator that was already run", 0)
	}
}

/* }}} */

func zim_Generator_rewind(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorRewind(generator)
}

/* }}} */

func zim_Generator_valid(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		return_value.SetTypeInfo(3)
	} else {
		return_value.SetTypeInfo(2)
	}
	return
}

/* }}} */

func zim_Generator_current(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetValue().GetType() != 0 {
		var value *Zval = &root.value
		var _z3 *Zval = value
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = return_value
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
}

/* }}} */

func zim_Generator_key(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	var root *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil && root.GetKey().GetType() != 0 {
		var key *Zval = &root.key
		var _z3 *Zval = key
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = return_value
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
}

/* }}} */

func zim_Generator_next(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorResume(generator)
}

/* }}} */

func zim_Generator_send(execute_data *ZendExecuteData, return_value *Zval) {
	var value *Zval
	var generator *ZendGenerator
	var root *ZendGenerator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
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
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &value, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)

	/* The generator is already closed, thus can't send anything */

	if generator.GetExecuteData() == nil {
		return
	}
	root = ZendGeneratorGetCurrent(generator)

	/* Put sent value in the target VAR slot, if it is used */

	if root.GetSendTarget() != nil {
		var _z1 *Zval = root.GetSendTarget()
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZendGeneratorResume(generator)
	root = ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		var value *Zval = &root.value
		var _z3 *Zval = value
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = return_value
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
}

/* }}} */

func zim_Generator_throw(execute_data *ZendExecuteData, return_value *Zval) {
	var exception *Zval
	var generator *ZendGenerator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
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
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &exception, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if exception.GetTypeFlags() != 0 {
		ZvalAddrefP(exception)
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if generator.GetExecuteData() != nil {
		var root *ZendGenerator = ZendGeneratorGetCurrent(generator)
		ZendGeneratorThrowException(root, exception)
		ZendGeneratorResume(generator)
		root = ZendGeneratorGetCurrent(generator)
		if generator.GetExecuteData() != nil {
			var value *Zval = &root.value
			var _z3 *Zval = value
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = return_value
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	} else {

		/* If the generator is already closed throw the exception in the
		 * current context */

		ZendThrowExceptionObject(exception)

		/* If the generator is already closed throw the exception in the
		 * current context */

	}
}

/* }}} */

func zim_Generator_getReturn(execute_data *ZendExecuteData, return_value *Zval) {
	var generator *ZendGenerator
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	generator = (*ZendGenerator)(&(execute_data.GetThis()).GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	if EG.GetException() != nil {
		return
	}
	if generator.GetRetval().GetType() == 0 {

		/* Generator hasn't returned yet -> error! */

		ZendThrowException(nil, "Cannot get return value of a generator that hasn't returned", 0)
		return
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = &generator.retval
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func ZendGeneratorIteratorDtor(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	generator.SetIterator(nil)
	ZvalPtrDtor(&iterator.data)
}

/* }}} */

func ZendGeneratorIteratorValid(iterator *ZendObjectIterator) int {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorGetCurrent(generator)
	if generator.GetExecuteData() != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendGeneratorIteratorGetData(iterator *ZendObjectIterator) *Zval {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	return &root.value
}

/* }}} */

func ZendGeneratorIteratorGetKey(iterator *ZendObjectIterator, key *Zval) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	var root *ZendGenerator
	ZendGeneratorEnsureInitialized(generator)
	root = ZendGeneratorGetCurrent(generator)
	if root.GetKey().GetType() != 0 {
		var zv *Zval = &root.key
		var _z3 *Zval = zv
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = key
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		key.SetTypeInfo(1)
	}
}

/* }}} */

func ZendGeneratorIteratorMoveForward(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	ZendGeneratorEnsureInitialized(generator)
	ZendGeneratorResume(generator)
}

/* }}} */

func ZendGeneratorIteratorRewind(iterator *ZendObjectIterator) {
	var generator *ZendGenerator = (*ZendGenerator)(iterator.GetData().GetValue().GetObj())
	ZendGeneratorRewind(generator)
}

/* }}} */

var ZendGeneratorIteratorFunctions ZendObjectIteratorFuncs = ZendObjectIteratorFuncs{ZendGeneratorIteratorDtor, ZendGeneratorIteratorValid, ZendGeneratorIteratorGetData, ZendGeneratorIteratorGetKey, ZendGeneratorIteratorMoveForward, ZendGeneratorIteratorRewind, nil}

func ZendGeneratorGetIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendObjectIterator
	var generator *ZendGenerator = (*ZendGenerator)(object.GetValue().GetObj())
	if generator.GetExecuteData() == nil {
		ZendThrowException(nil, "Cannot traverse an already closed generator", 0)
		return nil
	}
	if by_ref != 0 && (generator.GetExecuteData().GetFunc().GetOpArray().GetFnFlags()&1<<12) == 0 {
		ZendThrowException(nil, "You can only iterate a generator by-reference if it declared that it yields by-reference", 0)
		return nil
	}
	generator.SetIterator(_emalloc(g.SizeOf("zend_object_iterator")))
	iterator = generator.GetIterator()
	ZendIteratorInit(iterator)
	iterator.SetFuncs(&ZendGeneratorIteratorFunctions)
	ZvalAddrefP(object)
	var __z *Zval = &iterator.data
	__z.GetValue().SetObj(object.GetValue().GetObj())
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	return iterator
}

/* }}} */

var ArginfoGeneratorVoid []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGeneratorSend []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoGeneratorThrow []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"exception", 0, 0, 0}}
var GeneratorFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"rewind",
		zim_Generator_rewind,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_Generator_valid,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_Generator_current,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_Generator_key,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_Generator_next,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"send",
		zim_Generator_send,
		ArginfoGeneratorSend,
		uint32(g.SizeOf("arginfo_generator_send")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"throw",
		zim_Generator_throw,
		ArginfoGeneratorThrow,
		uint32(g.SizeOf("arginfo_generator_throw")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getReturn",
		zim_Generator_getReturn,
		ArginfoGeneratorVoid,
		uint32(g.SizeOf("arginfo_generator_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func ZendRegisterGeneratorCe() {
	var ce ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Generator", g.SizeOf("\"Generator\"")-1, 1))
	ce.SetBuiltinFunctions(GeneratorFunctions)
	ZendCeGenerator = ZendRegisterInternalClass(&ce)
	ZendCeGenerator.SetCeFlags(ZendCeGenerator.GetCeFlags() | 1<<5)
	ZendCeGenerator.create_object = ZendGeneratorCreate
	ZendCeGenerator.SetSerialize(ZendClassSerializeDeny)
	ZendCeGenerator.SetUnserialize(ZendClassUnserializeDeny)

	/* get_iterator has to be assigned *after* implementing the inferface */

	ZendClassImplements(ZendCeGenerator, 1, ZendCeIterator)
	ZendCeGenerator.SetGetIterator(ZendGeneratorGetIterator)
	memcpy(&ZendGeneratorHandlers, &StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	ZendGeneratorHandlers.SetFreeObj(ZendGeneratorFreeStorage)
	ZendGeneratorHandlers.SetDtorObj(ZendGeneratorDtorStorage)
	ZendGeneratorHandlers.SetGetGc(ZendGeneratorGetGc)
	ZendGeneratorHandlers.SetCloneObj(nil)
	ZendGeneratorHandlers.SetGetConstructor(ZendGeneratorGetConstructor)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ClosedGeneratorException", g.SizeOf("\"ClosedGeneratorException\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	zend_ce_ClosedGeneratorException = ZendRegisterInternalClassEx(&ce, ZendCeException)
}

/* }}} */
