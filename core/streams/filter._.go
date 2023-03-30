package streams

import (
	"github.com/heyuuu/gophp/zend/types"
)

/* Global filter hash, copied to FG(stream_filters) on registration of volatile filter */

var StreamFiltersHash types.Array

/* Should only be used during core initialization */

/* Normal hash selection/retrieval call */

/* API for registering GLOBAL filters */

/* API for registering VOLATILE wrappers */

/* Buckets */

/* Given a bucket, returns a version of that bucket with a writeable buffer.
 * If the original bucket has a refcount of 1 and owns its buffer, then it
 * is returned unchanged.
 * Otherwise, a copy of the buffer is made.
 * In both cases, the original bucket is unlinked from its brigade.
 * If a copy is made, the original bucket is delref'd.
 * */

/* We allow very simple pattern matching for filter factories:
 * if "convert.charset.utf-8/sjis" is requested, we search first for an exact
 * match. If that fails, we try "convert.charset.*", then "convert.*"
 * This means that we don't need to clog up the hashtable with a zillion
 * charsets (for example) but still be able to provide them all as filters */
