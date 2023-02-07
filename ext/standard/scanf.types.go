// <<generate>>

package standard

/**
 * CharSet
 */
type CharSet struct {
	exclude int
	nchars  int
	chars   *byte
	nranges int
	ranges  *struct {
		start byte
		end   byte
	}
}

//             func NewCharSet(exclude int, nchars int, chars *byte, nranges int, ranges *struct {
// start byte
// end byte
// }) *CharSet {
//                 return &CharSet{
//                     exclude:exclude,
//                     nchars:nchars,
//                     chars:chars,
//                     nranges:nranges,
//                     ranges:ranges,
//                 }
//             }
//             func MakeCharSet(exclude int, nchars int, chars *byte, nranges int, ranges *struct {
// start byte
// end byte
// }) CharSet {
//                 return CharSet{
//                     exclude:exclude,
//                     nchars:nchars,
//                     chars:chars,
//                     nranges:nranges,
//                     ranges:ranges,
//                 }
//             }
func (this *CharSet) GetExclude() int      { return this.exclude }
func (this *CharSet) SetExclude(value int) { this.exclude = value }
func (this *CharSet) GetNchars() int       { return this.nchars }
func (this *CharSet) SetNchars(value int)  { this.nchars = value }
func (this *CharSet) GetChars() *byte      { return this.chars }
func (this *CharSet) SetChars(value *byte) { this.chars = value }
func (this *CharSet) GetNranges() int      { return this.nranges }
func (this *CharSet) SetNranges(value int) { this.nranges = value }
func (this *CharSet) GetRanges() *struct {
	start byte
	end   byte
} {
	return this.ranges
}
func (this *CharSet) SetRanges(value *struct {
	start byte
	end   byte
}) {
	this.ranges = value
}
