// <<generate>>

package standard

/**
 * PhpLcgGlobals
 */
type PhpLcgGlobals struct {
	s1     int32
	s2     int32
	seeded int
}

func (this PhpLcgGlobals) GetS1() int32         { return this.s1 }
func (this *PhpLcgGlobals) SetS1(value int32)   { this.s1 = value }
func (this PhpLcgGlobals) GetS2() int32         { return this.s2 }
func (this *PhpLcgGlobals) SetS2(value int32)   { this.s2 = value }
func (this PhpLcgGlobals) GetSeeded() int       { return this.seeded }
func (this *PhpLcgGlobals) SetSeeded(value int) { this.seeded = value }
