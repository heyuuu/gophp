package types

type ZendRefcounted struct {
	gc struct {
		refcount uint32
	}
}

// Refcount
func (this *ZendRefcounted) GetRefcount() uint32 {
	return this.gc.refcount
}
