// <<generate>>

package spl

// Source: <ext/spl/spl_iterators.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/pcre/php_pcre.h"

// Source: <ext/spl/spl_iterators.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* }}} */

/* {{{ proto RecursiveIteratorIterator::__construct(RecursiveIterator|IteratorAggregate it [, int mode = RIT_LEAVES_ONLY [, int flags = 0]]) throws InvalidArgumentException
   Creates a RecursiveIteratorIterator from a RecursiveIterator. */

/* {{{ proto void RecursiveIteratorIterator::rewind()
   Rewind the iterator to the first element of the top level inner iterator. */

/* {{{ proto bool RecursiveIteratorIterator::valid()
   Check whether the current position is valid */

/* {{{ proto mixed RecursiveIteratorIterator::key()
   Access the current key */

/* {{{ proto mixed RecursiveIteratorIterator::current()
   Access the current element value */

/* {{{ proto void RecursiveIteratorIterator::next()
   Move forward to the next element */

/* {{{ proto int RecursiveIteratorIterator::getDepth()
   Get the current depth of the recursive iteration */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getSubIterator([int level])
   The current active sub iterator or the iterator at specified level */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getInnerIterator()
   The current active sub iterator */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::beginIteration()
   Called when iteration begins (after first rewind() call) */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::endIteration()
   Called when iteration ends (when valid() first returns false */

/* {{{ proto bool RecursiveIteratorIterator::callHasChildren()
   Called for each element to test whether it has children */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::callGetChildren()
   Return children of current element */

/* {{{ proto void RecursiveIteratorIterator::beginChildren()
   Called when recursing one level down */

/* {{{ proto void RecursiveIteratorIterator::endChildren()
   Called when end recursing one level */

/* {{{ proto void RecursiveIteratorIterator::nextElement()
   Called when the next element is available */

/* {{{ proto void RecursiveIteratorIterator::setMaxDepth([$max_depth = -1])
   Set the maximum allowed depth (or any depth if pmax_depth = -1] */

/* {{{ proto int|false RecursiveIteratorIterator::getMaxDepth()
   Return the maximum accepted depth or false if any depth is allowed */

/* {{{ spl_RecursiveIteratorIterator_dtor */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto RecursiveTreeIterator::__construct(RecursiveIterator|IteratorAggregate it [, int flags = RTIT_BYPASS_KEY [, int cit_flags = CIT_CATCH_GET_CHILD [, mode = RIT_SELF_FIRST ]]]) throws InvalidArgumentException
   RecursiveIteratorIterator to generate ASCII graphic trees for the entries in a RecursiveIterator */

/* {{{ proto void RecursiveTreeIterator::setPrefixPart(int part, string prefix) throws OutOfRangeException
   Sets prefix parts as used in getPrefix() */

/* {{{ proto string RecursiveTreeIterator::getPrefix()
   Returns the string to place in front of current element */

/* {{{ proto void RecursiveTreeIterator::setPostfix(string prefix)
   Sets postfix as used in getPostfix() */

/* {{{ proto string RecursiveTreeIterator::getEntry()
   Returns the string presentation built for current element */

/* {{{ proto string RecursiveTreeIterator::getPostfix()
   Returns the string to place after the current element */

/* {{{ proto mixed RecursiveTreeIterator::current()
   Returns the current element prefixed and postfixed */

/* {{{ proto mixed RecursiveTreeIterator::key()
   Returns the current key prefixed and postfixed */

/* {{{ proto FilterIterator::__construct(Iterator it)
   Create an Iterator from another iterator */

/* {{{ proto CallbackFilterIterator::__construct(Iterator it, callback func)
   Create an Iterator from another iterator */

/* {{{ proto Iterator FilterIterator::getInnerIterator()
    proto Iterator CachingIterator::getInnerIterator()
    proto Iterator LimitIterator::getInnerIterator()
    proto Iterator ParentIterator::getInnerIterator()
Get the inner iterator */

/* {{{ proto void ParentIterator::rewind()
       proto void IteratorIterator::rewind()
   Rewind the iterator
*/

/* {{{ proto bool FilterIterator::valid()
    proto bool ParentIterator::valid()
    proto bool IteratorIterator::valid()
    proto bool NoRewindIterator::valid()
Check whether the current element is valid */

/* {{{ proto mixed FilterIterator::key()
    proto mixed CachingIterator::key()
    proto mixed LimitIterator::key()
    proto mixed ParentIterator::key()
    proto mixed IteratorIterator::key()
    proto mixed NoRewindIterator::key()
    proto mixed AppendIterator::key()
Get the current key */

/* {{{ proto mixed FilterIterator::current()
    proto mixed CachingIterator::current()
    proto mixed LimitIterator::current()
    proto mixed ParentIterator::current()
    proto mixed IteratorIterator::current()
    proto mixed NoRewindIterator::current()
Get the current element value */

/* {{{ proto void ParentIterator::next()
    proto void IteratorIterator::next()
    proto void NoRewindIterator::next()
Move the iterator forward */

/* {{{ proto void FilterIterator::rewind()
   Rewind the iterator */

/* {{{ proto void FilterIterator::next()
   Move the iterator forward */

/* {{{ proto RecursiveCallbackFilterIterator::__construct(RecursiveIterator it, callback func)
   Create a RecursiveCallbackFilterIterator from a RecursiveIterator */

/* {{{ proto RecursiveFilterIterator::__construct(RecursiveIterator it)
   Create a RecursiveFilterIterator from a RecursiveIterator */

/* {{{ proto bool RecursiveFilterIterator::hasChildren()
   Check whether the inner iterator's current element has children */

/* {{{ proto RecursiveFilterIterator RecursiveFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveFilterIterator */

/* {{{ proto RecursiveCallbackFilterIterator RecursiveCallbackFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveCallbackFilterIterator */

/* {{{ proto ParentIterator::__construct(RecursiveIterator it)
   Create a ParentIterator from a RecursiveIterator */

/* {{{ proto RegexIterator::__construct(Iterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RegexIterator from another iterator and a regular expression */

/* {{{ proto bool CallbackFilterIterator::accept()
   Calls the callback with the current value, the current key and the inner iterator as arguments */

/* }}} */

/* {{{ proto string RegexIterator::getRegex()
   Returns current regular expression */

/* {{{ proto bool RegexIterator::getMode()
   Returns current operation mode */

/* {{{ proto bool RegexIterator::setMode(int new_mode)
   Set new operation mode */

/* {{{ proto bool RegexIterator::getFlags()
   Returns current operation flags */

/* {{{ proto bool RegexIterator::setFlags(int new_flags)
   Set operation flags */

/* {{{ proto bool RegexIterator::getFlags()
   Returns current PREG flags (if in use or NULL) */

/* {{{ proto bool RegexIterator::setPregFlags(int new_flags)
   Set PREG flags */

/* {{{ proto RecursiveRegexIterator::__construct(RecursiveIterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RecursiveRegexIterator from another recursive iterator and a regular expression */

/* {{{ proto RecursiveRegexIterator RecursiveRegexIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveRegexIterator */

/* {{{ spl_dual_it_dtor */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto LimitIterator::__construct(Iterator it [, int offset, int count])
   Construct a LimitIterator from an Iterator with a given starting offset and optionally a maximum count */

/* {{{ proto void LimitIterator::rewind()
   Rewind the iterator to the specified starting offset */

/* {{{ proto bool LimitIterator::valid()
   Check whether the current element is valid */

/* {{{ proto void LimitIterator::next()
   Move the iterator forward */

/* {{{ proto void LimitIterator::seek(int position)
   Seek to the given position */

/* {{{ proto int LimitIterator::getPosition()
   Return the current position */

/* {{{ proto CachingIterator::__construct(Iterator it [, flags = CIT_CALL_TOSTRING])
   Construct a CachingIterator from an Iterator */

/* {{{ proto void CachingIterator::rewind()
   Rewind the iterator */

/* {{{ proto bool CachingIterator::valid()
   Check whether the current element is valid */

/* {{{ proto void CachingIterator::next()
   Move the iterator forward */

/* {{{ proto bool CachingIterator::hasNext()
   Check whether the inner iterator has a valid next element */

/* {{{ proto string CachingIterator::__toString()
   Return the string representation of the current element */

/* {{{ proto void CachingIterator::offsetSet(mixed index, mixed newval)
   Set given index in cache */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto RecursiveCachingIterator::__construct(RecursiveIterator it [, flags = CIT_CALL_TOSTRING])
   Create an iterator from a RecursiveIterator */

/* {{{ proto bool RecursiveCachingIterator::hasChildren()
   Check whether the current element of the inner iterator has children */

/* {{{ proto RecursiveCachingIterator RecursiveCachingIterator::getChildren()
Return the inner iterator's children as a RecursiveCachingIterator */

/* {{{ proto IteratorIterator::__construct(Traversable it)
   Create an iterator from anything that is traversable */

/* {{{ proto NoRewindIterator::__construct(Iterator it)
   Create an iterator from another iterator */

/* {{{ proto void NoRewindIterator::rewind()
   Prevent a call to inner iterators rewind() */

/* {{{ proto bool NoRewindIterator::valid()
   Return inner iterators valid() */

/* {{{ proto mixed NoRewindIterator::key()
   Return inner iterators key() */

/* {{{ proto mixed NoRewindIterator::current()
   Return inner iterators current() */

/* {{{ proto void NoRewindIterator::next()
   Return inner iterators next() */

/* {{{ proto InfiniteIterator::__construct(Iterator it)
   Create an iterator from another iterator */

/* {{{ proto void InfiniteIterator::next()
   Prevent a call to inner iterators rewind() (internally the current data will be fetched if valid()) */

/* {{{ proto void EmptyIterator::rewind()
   Does nothing  */

/* {{{ proto false EmptyIterator::valid()
   Return false */

/* {{{ proto void EmptyIterator::key()
   Throws exception BadMethodCallException */

/* {{{ proto void EmptyIterator::current()
   Throws exception BadMethodCallException */

/* {{{ proto void EmptyIterator::next()
   Does nothing */

/* {{{ proto AppendIterator::__construct()
   Create an AppendIterator */

/* {{{ proto void AppendIterator::append(Iterator it)
   Append an iterator */

/* {{{ proto mixed AppendIterator::current()
   Get the current element value */

/* {{{ proto void AppendIterator::rewind()
   Rewind to the first iterator and rewind the first iterator, too */

/* {{{ proto bool AppendIterator::valid()
   Check if the current state is valid */

/* {{{ proto void AppendIterator::next()
   Forward to next element */

/* {{{ proto int AppendIterator::getIteratorIndex()
   Get index of iterator */

/* {{{ proto ArrayIterator AppendIterator::getArrayIterator()
   Get access to inner ArrayIterator */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ PHP_MINIT_FUNCTION(spl_iterators)
 */

/* }}} */
