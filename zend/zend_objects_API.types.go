// <<generate>>

package zend

import "sik/zend/types"

/**
 * ZendObjectsStore
 */
type ZendObjectsStore struct {
	object_buckets **types.ZendObject
	top            uint32
	size           uint32
	free_list_head int
}

// func MakeZendObjectsStore(object_buckets **ZendObject, top uint32, size uint32, free_list_head int) ZendObjectsStore {
//     return ZendObjectsStore{
//         object_buckets:object_buckets,
//         top:top,
//         size:size,
//         free_list_head:free_list_head,
//     }
// }
func (this *ZendObjectsStore) GetObjectBuckets() **types.ZendObject      { return this.object_buckets }
func (this *ZendObjectsStore) SetObjectBuckets(value **types.ZendObject) { this.object_buckets = value }
func (this *ZendObjectsStore) GetTop() uint32                            { return this.top }
func (this *ZendObjectsStore) SetTop(value uint32)                       { this.top = value }
func (this *ZendObjectsStore) GetSize() uint32                           { return this.size }
func (this *ZendObjectsStore) SetSize(value uint32)                      { this.size = value }
func (this *ZendObjectsStore) GetFreeListHead() int                      { return this.free_list_head }
func (this *ZendObjectsStore) SetFreeListHead(value int)                 { this.free_list_head = value }
