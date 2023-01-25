// <<generate>>

package streams

/**
 * GlobST
 */
type GlobST struct {
	glob        glob_t
	index       int
	flags       int
	path        *byte
	path_len    int
	pattern     *byte
	pattern_len int
}

func (this GlobST) GetGlob() glob_t          { return this.glob }
func (this *GlobST) SetGlob(value glob_t)    { this.glob = value }
func (this GlobST) GetIndex() int            { return this.index }
func (this *GlobST) SetIndex(value int)      { this.index = value }
func (this GlobST) GetFlags() int            { return this.flags }
func (this *GlobST) SetFlags(value int)      { this.flags = value }
func (this GlobST) GetPath() *byte           { return this.path }
func (this *GlobST) SetPath(value *byte)     { this.path = value }
func (this GlobST) GetPathLen() int          { return this.path_len }
func (this *GlobST) SetPathLen(value int)    { this.path_len = value }
func (this GlobST) GetPattern() *byte        { return this.pattern }
func (this *GlobST) SetPattern(value *byte)  { this.pattern = value }
func (this GlobST) GetPatternLen() int       { return this.pattern_len }
func (this *GlobST) SetPatternLen(value int) { this.pattern_len = value }

/* GlobST.flags */
func (this *GlobST) AddFlags(value int)     { this.flags |= value }
func (this *GlobST) SubFlags(value int)     { this.flags &^= value }
func (this GlobST) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *GlobST) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
