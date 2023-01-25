// <<generate>>

package spl

// #define PHP_SPL_H

// # include "php.h"

// # include < stdarg . h >

// #define SPL_API

// # include "php.h"

// # include "php_ini.h"

// # include "php_main.h"

// # include "ext/standard/info.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_array.h"

// # include "spl_directory.h"

// # include "spl_iterators.h"

// # include "spl_exceptions.h"

// # include "spl_observer.h"

// # include "spl_dllist.h"

// # include "spl_fixedarray.h"

// # include "spl_heap.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "ext/standard/php_mt_rand.h"

// # include "main/snprintf.h"

// #define SPL_ADD_CLASS(class_name,z_list,sub,allow,ce_flags) spl_add_classes ( spl_ce_ ## class_name , z_list , sub , allow , ce_flags )

// #define SPL_LIST_CLASSES(z_list,sub,allow,ce_flags) SPL_ADD_CLASS ( AppendIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ArrayIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ArrayObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( BadFunctionCallException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( BadMethodCallException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( CachingIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( CallbackFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( DirectoryIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( DomainException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( EmptyIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( FilesystemIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( FilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( GlobIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( InfiniteIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( InvalidArgumentException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( IteratorIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LengthException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LimitIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LogicException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( MultipleIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( NoRewindIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OuterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OutOfBoundsException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OutOfRangeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OverflowException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ParentIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RangeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveArrayIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveCachingIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveCallbackFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveDirectoryIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveIteratorIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveRegexIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveTreeIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RegexIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RuntimeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SeekableIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplDoublyLinkedList , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFileInfo , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFileObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFixedArray , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplMinHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplMaxHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplObjectStorage , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplObserver , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplPriorityQueue , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplQueue , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplStack , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplSubject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplTempFileObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( UnderflowException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( UnexpectedValueException , z_list , sub , allow , ce_flags ) ;
