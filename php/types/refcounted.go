package types

type ZendRefcounted struct {
	gc struct {
		refcount uint32
	}
}
