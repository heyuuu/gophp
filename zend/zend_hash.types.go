// <<generate>>

package zend

/**
 * ZendHashKey
 */
type ZendHashKey struct {
	h   ZendUlong
	key *ZendString
}

func (this ZendHashKey) GetH() ZendUlong           { return this.h }
func (this *ZendHashKey) SetH(value ZendUlong)     { this.h = value }
func (this ZendHashKey) GetKey() *ZendString       { return this.key }
func (this *ZendHashKey) SetKey(value *ZendString) { this.key = value }
