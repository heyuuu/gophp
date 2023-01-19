// <<generate>>

package zend

/**
 * ZendArena
 */
type ZendArena struct {
	ptr  *byte
	end  *byte
	prev *ZendArena
}

func (this ZendArena) GetPtr() *byte             { return this.ptr }
func (this *ZendArena) SetPtr(value *byte)       { this.ptr = value }
func (this ZendArena) GetEnd() *byte             { return this.end }
func (this *ZendArena) SetEnd(value *byte)       { this.end = value }
func (this ZendArena) GetPrev() *ZendArena       { return this.prev }
func (this *ZendArena) SetPrev(value *ZendArena) { this.prev = value }
