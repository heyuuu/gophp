package streams

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpStreamBucket
 */
type PhpStreamBucket struct {
	next          *PhpStreamBucket
	prev          **PhpStreamBucket
	brigade       *PhpStreamBucketBrigade
	buf           *byte
	buflen        int
	own_buf       uint8
	is_persistent uint8
	refcount      int
}

func (this *PhpStreamBucket) GetNext() *PhpStreamBucket                { return this.next }
func (this *PhpStreamBucket) SetNext(value *PhpStreamBucket)           { this.next = value }
func (this *PhpStreamBucket) GetPrev() **PhpStreamBucket               { return this.prev }
func (this *PhpStreamBucket) SetPrev(value **PhpStreamBucket)          { this.prev = value }
func (this *PhpStreamBucket) GetBrigade() *PhpStreamBucketBrigade      { return this.brigade }
func (this *PhpStreamBucket) SetBrigade(value *PhpStreamBucketBrigade) { this.brigade = value }
func (this *PhpStreamBucket) GetBuf() *byte                            { return this.buf }
func (this *PhpStreamBucket) SetBuf(value *byte)                       { this.buf = value }
func (this *PhpStreamBucket) GetBuflen() int                           { return this.buflen }
func (this *PhpStreamBucket) SetBuflen(value int)                      { this.buflen = value }
func (this *PhpStreamBucket) GetOwnBuf() uint8                         { return this.own_buf }
func (this *PhpStreamBucket) SetOwnBuf(value uint8)                    { this.own_buf = value }
func (this *PhpStreamBucket) GetIsPersistent() uint8                   { return this.is_persistent }
func (this *PhpStreamBucket) SetIsPersistent(value uint8)              { this.is_persistent = value }
func (this *PhpStreamBucket) GetRefcount() int                         { return this.refcount }
func (this *PhpStreamBucket) SetRefcount(value int)                    { this.refcount = value }

/**
 * PhpStreamBucketBrigade
 */
type PhpStreamBucketBrigade struct {
	head *PhpStreamBucket
	tail **PhpStreamBucket
}

func MakePhpStreamBucketBrigade(head *PhpStreamBucket, tail **PhpStreamBucket) PhpStreamBucketBrigade {
	return PhpStreamBucketBrigade{
		head: head,
		tail: tail,
	}
}
func (this *PhpStreamBucketBrigade) GetHead() *PhpStreamBucket       { return this.head }
func (this *PhpStreamBucketBrigade) SetHead(value *PhpStreamBucket)  { this.head = value }
func (this *PhpStreamBucketBrigade) GetTail() **PhpStreamBucket      { return this.tail }
func (this *PhpStreamBucketBrigade) SetTail(value **PhpStreamBucket) { this.tail = value }

/**
 * PhpStreamFilterOps
 */
type PhpStreamFilterOps struct {
	filter func(
		stream *core.PhpStream,
		thisfilter *core.PhpStreamFilter,
		buckets_in *PhpStreamBucketBrigade,
		buckets_out *PhpStreamBucketBrigade,
		bytes_consumed *int,
		flags int,
	) PhpStreamFilterStatusT
	dtor  func(thisfilter *core.PhpStreamFilter)
	label *byte
}

func MakePhpStreamFilterOps(filter func(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *PhpStreamBucketBrigade,
	buckets_out *PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) PhpStreamFilterStatusT, dtor func(thisfilter *core.PhpStreamFilter), label *byte) PhpStreamFilterOps {
	return PhpStreamFilterOps{
		filter: filter,
		dtor:   dtor,
		label:  label,
	}
}
func (this *PhpStreamFilterOps) GetFilter() func(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *PhpStreamBucketBrigade,
	buckets_out *PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) PhpStreamFilterStatusT {
	return this.filter
}
func (this *PhpStreamFilterOps) SetFilter(value func(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *PhpStreamBucketBrigade,
	buckets_out *PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) PhpStreamFilterStatusT) {
	this.filter = value
}
func (this *PhpStreamFilterOps) GetDtor() func(thisfilter *core.PhpStreamFilter) { return this.dtor }

func (this *PhpStreamFilterOps) GetLabel() *byte { return this.label }

/**
 * PhpStreamFilterChain
 */
type PhpStreamFilterChain struct {
	head   *core.PhpStreamFilter
	tail   **core.PhpStreamFilter
	stream *core.PhpStream
}

func (this *PhpStreamFilterChain) GetHead() *core.PhpStreamFilter       { return this.head }
func (this *PhpStreamFilterChain) SetHead(value *core.PhpStreamFilter)  { this.head = value }
func (this *PhpStreamFilterChain) GetTail() **core.PhpStreamFilter      { return this.tail }
func (this *PhpStreamFilterChain) SetTail(value **core.PhpStreamFilter) { this.tail = value }
func (this *PhpStreamFilterChain) GetStream() *core.PhpStream           { return this.stream }
func (this *PhpStreamFilterChain) SetStream(value *core.PhpStream)      { this.stream = value }

/**
 * PhpStreamFilter
 */
type PhpStreamFilter struct {
	fops          *PhpStreamFilterOps
	abstract      types.Zval
	next          *core.PhpStreamFilter
	prev          *core.PhpStreamFilter
	is_persistent int
	chain         *PhpStreamFilterChain
	buffer        PhpStreamBucketBrigade
	res           *types.ZendResource
}

func (this *PhpStreamFilter) GetFops() *PhpStreamFilterOps         { return this.fops }
func (this *PhpStreamFilter) SetFops(value *PhpStreamFilterOps)    { this.fops = value }
func (this *PhpStreamFilter) GetAbstract() types.Zval              { return this.abstract }
func (this *PhpStreamFilter) GetNext() *core.PhpStreamFilter       { return this.next }
func (this *PhpStreamFilter) SetNext(value *core.PhpStreamFilter)  { this.next = value }
func (this *PhpStreamFilter) GetPrev() *core.PhpStreamFilter       { return this.prev }
func (this *PhpStreamFilter) SetPrev(value *core.PhpStreamFilter)  { this.prev = value }
func (this *PhpStreamFilter) GetIsPersistent() int                 { return this.is_persistent }
func (this *PhpStreamFilter) SetIsPersistent(value int)            { this.is_persistent = value }
func (this *PhpStreamFilter) GetChain() *PhpStreamFilterChain      { return this.chain }
func (this *PhpStreamFilter) SetChain(value *PhpStreamFilterChain) { this.chain = value }

func (this *PhpStreamFilter) GetRes() *types.ZendResource      { return this.res }
func (this *PhpStreamFilter) SetRes(value *types.ZendResource) { this.res = value }

/**
 * PhpStreamFilterFactory
 */
type PhpStreamFilterFactory struct {
	create_filter func(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter
}

func MakePhpStreamFilterFactory(create_filter func(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter) PhpStreamFilterFactory {
	return PhpStreamFilterFactory{
		create_filter: create_filter,
	}
}
func (this *PhpStreamFilterFactory) GetCreateFilter() func(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	return this.create_filter
}
