// <<generate>>

package streams

import (
	"sik/core"
)

func PhpStreamBucketAddref(bucket __auto__) int {
	bucket.refcount++
	return bucket.refcount - 1
}
func PhpStreamFilterAlloc(fops *PhpStreamFilterOps, thisptr any, persistent uint8) *core.PhpStreamFilter {
	return _phpStreamFilterAlloc(fops, thisptr, persistent)
}
func PhpStreamFilterAllocRel(fops *PhpStreamFilterOps, thisptr any, persistent uint8) *core.PhpStreamFilter {
	return _phpStreamFilterAlloc(fops, thisptr, persistent)
}
func PhpStreamFilterPrepend(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	_phpStreamFilterPrepend(chain, filter)
}
func PhpStreamFilterAppend(chain PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	_phpStreamFilterAppend(chain, filter)
}
func PhpStreamFilterFlush(filter *core.PhpStreamFilter, finish int) int {
	return _phpStreamFilterFlush(filter, finish)
}
func PhpStreamIsFiltered(stream *core.PhpStream) bool {
	return stream.readfilters.GetHead() != nil || stream.writefilters.GetHead() != nil
}
