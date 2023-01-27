// <<generate>>

package cli

/**
 * PhpCliServerExtMimeTypePair
 */
type PhpCliServerExtMimeTypePair struct {
	ext       *byte
	mime_type *byte
}

func (this *PhpCliServerExtMimeTypePair) GetExt() *byte           { return this.ext }
func (this *PhpCliServerExtMimeTypePair) SetExt(value *byte)      { this.ext = value }
func (this *PhpCliServerExtMimeTypePair) GetMimeType() *byte      { return this.mime_type }
func (this *PhpCliServerExtMimeTypePair) SetMimeType(value *byte) { this.mime_type = value }
