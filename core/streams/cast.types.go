// <<generate>>

package streams

/**
 * COOKIE_IO_FUNCTIONS_T
 */
type COOKIE_IO_FUNCTIONS_T struct {
	reader func(any, *byte, int) int
	writer func(any, *byte, int) int
	seeker func(any, PHP_FPOS_T, int) PHP_FPOS_T
	closer func(any) int
}

func (this *COOKIE_IO_FUNCTIONS_T) GetReader() func(any, *byte, int) int      { return this.reader }
func (this *COOKIE_IO_FUNCTIONS_T) SetReader(value func(any, *byte, int) int) { this.reader = value }
func (this *COOKIE_IO_FUNCTIONS_T) GetWriter() func(any, *byte, int) int      { return this.writer }
func (this *COOKIE_IO_FUNCTIONS_T) SetWriter(value func(any, *byte, int) int) { this.writer = value }
func (this *COOKIE_IO_FUNCTIONS_T) GetSeeker() func(any, PHP_FPOS_T, int) PHP_FPOS_T {
	return this.seeker
}
func (this *COOKIE_IO_FUNCTIONS_T) SetSeeker(value func(any, PHP_FPOS_T, int) PHP_FPOS_T) {
	this.seeker = value
}
func (this *COOKIE_IO_FUNCTIONS_T) GetCloser() func(any) int      { return this.closer }
func (this *COOKIE_IO_FUNCTIONS_T) SetCloser(value func(any) int) { this.closer = value }
