// <<generate>>

package zend

import "sik/zend/types"

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

var ZendCeGenerator *types.ClassEntry
var zend_ce_ClosedGeneratorException *types.ClassEntry

/* The concept of `yield from` exposes problems when accessed at different levels of the chain of delegated generators. We need to be able to reference the currently executed Generator in all cases and still being able to access the return values of finished Generators.
 * The solution to this problem is a doubly-linked tree, which all Generators referenced in maintain a reference to. It should be impossible to avoid walking the tree in all cases. This way, we only need tree walks from leaf to root in case where some part of the `yield from` chain is passed to another `yield from`. (Update of leaf node pointer and list of multi-children nodes needed when leaf gets a child in direct path from leaf to root node.) But only in that case, which should be a fairly rare case (which is then possible, but not totally cheap).
 * The root of the tree is then the currently executed Generator. The subnodes of the tree (all except the root node) are all Generators which do `yield from`. Each node of the tree knows a pointer to one leaf descendant node. Each node with multiple children needs a list of all leaf descendant nodes paired with pointers to their respective child node. (The stack is determined by leaf node pointers) Nodes with only one child just don't need a list, there it is enough to just have a pointer to the child node. Further, leaf nodes store a pointer to the root node.
 * That way, when we advance any generator, we just need to look up a leaf node (which all have a reference to a root node). Then we can see at the root node whether current Generator is finished. If it isn't, all is fine and we can just continue. If the Generator finished, there will be two cases. Either it is a simple node with just one child, then go down to child node. Or it has multiple children and we now will remove the current leaf node from the list of nodes (unnecessary, is microoptimization) and go down to the child node whose reference was paired with current leaf node. Child node is then removed its parent reference and becomes new top node. Or the current node references the Generator we're currently executing, then we can continue from the YIELD_FROM opcode. When a node referenced as root node in a leaf node has a parent, then we go the way up until we find a root node without parent.
 * In case we go into a new `yield from` level, a node is created on top of current root and becomes the new root. Leaf node needs to be updated with new root node then.
 * When a Generator referenced by a node of the tree is added to `yield from`, that node now gets a list of children (we need to walk the descendants of that node and nodes of the tree of the other Generator down to the first multi-children node and copy all the leaf node pointers from there). In case there was no multi-children node (linear tree), we just add a pair (pointer to leaf node, pointer to child node), with the child node being in a direct path from leaf to this node.
 */

var ZEND_GENERATOR_CURRENTLY_RUNNING types.ZendUchar = 0x1
var ZEND_GENERATOR_FORCED_CLOSE types.ZendUchar = 0x2
var ZEND_GENERATOR_AT_FIRST_YIELD types.ZendUchar = 0x4
var ZEND_GENERATOR_DO_INIT types.ZendUchar = 0x8

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

var ZendGeneratorHandlers ZendObjectHandlers

/* Pay attention so that the root of each subtree of the Generators tree is referenced
 * once per leaf */

var ZendGeneratorIteratorFunctions ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendGeneratorIteratorDtor, ZendGeneratorIteratorValid, ZendGeneratorIteratorGetData, ZendGeneratorIteratorGetKey, ZendGeneratorIteratorMoveForward, ZendGeneratorIteratorRewind, nil)
var ArginfoGeneratorVoid []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(-1),
}
var ArginfoGeneratorSend []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("value"),
}
var ArginfoGeneratorThrow []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("exception"),
}
var GeneratorFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntryEx("rewind", ZEND_ACC_PUBLIC, zim_Generator_rewind, ArginfoGeneratorVoid),
	MakeZendFunctionEntryEx("valid", ZEND_ACC_PUBLIC, zim_Generator_valid, ArginfoGeneratorVoid),
	MakeZendFunctionEntryEx("current", ZEND_ACC_PUBLIC, zim_Generator_current, ArginfoGeneratorVoid),
	MakeZendFunctionEntryEx("key", ZEND_ACC_PUBLIC, zim_Generator_key, ArginfoGeneratorVoid),
	MakeZendFunctionEntryEx("next", ZEND_ACC_PUBLIC, zim_Generator_next, ArginfoGeneratorVoid),
	MakeZendFunctionEntryEx("send", ZEND_ACC_PUBLIC, zim_Generator_send, ArginfoGeneratorSend),
	MakeZendFunctionEntryEx("throw", ZEND_ACC_PUBLIC, zim_Generator_throw, ArginfoGeneratorThrow),
	MakeZendFunctionEntryEx("getReturn", ZEND_ACC_PUBLIC, zim_Generator_getReturn, ArginfoGeneratorVoid),
}
