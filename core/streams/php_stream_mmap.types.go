package streams

/**
 * PhpStreamMmapRange
 */
type PhpStreamMmapRange struct {
	offset int
	length int
	mode   PhpStreamMmapAccessT
	mapped *byte
}

// func MakePhpStreamMmapRange(offset int, length int, mode PhpStreamMmapAccessT, mapped *byte) PhpStreamMmapRange {
//     return PhpStreamMmapRange{
//         offset:offset,
//         length:length,
//         mode:mode,
//         mapped:mapped,
//     }
// }
func (this *PhpStreamMmapRange) GetOffset() int                     { return this.offset }
func (this *PhpStreamMmapRange) SetOffset(value int)                { this.offset = value }
func (this *PhpStreamMmapRange) GetLength() int                     { return this.length }
func (this *PhpStreamMmapRange) SetLength(value int)                { this.length = value }
func (this *PhpStreamMmapRange) GetMode() PhpStreamMmapAccessT      { return this.mode }
func (this *PhpStreamMmapRange) SetMode(value PhpStreamMmapAccessT) { this.mode = value }
func (this *PhpStreamMmapRange) GetMapped() *byte                   { return this.mapped }
func (this *PhpStreamMmapRange) SetMapped(value *byte)              { this.mapped = value }
