// <<generate>>

package standard

/**
 * PhpCryptExtendedData
 */
type PhpCryptExtendedData struct {
	initialized int
	saltbits    uint32
	old_salt    uint32
	en_keysl    []uint32
	en_keysr    []uint32
	de_keysl    []uint32
	de_keysr    []uint32
	old_rawkey0 uint32
	old_rawkey1 uint32
	output      []byte
}

//             func MakePhpCryptExtendedData(
// initialized int,
// saltbits uint32,
// old_salt uint32,
// en_keysl []uint32,
// en_keysr []uint32,
// de_keysl []uint32,
// de_keysr []uint32,
// old_rawkey0 uint32,
// old_rawkey1 uint32,
// output []byte,
// ) PhpCryptExtendedData {
//                 return PhpCryptExtendedData{
//                     initialized:initialized,
//                     saltbits:saltbits,
//                     old_salt:old_salt,
//                     en_keysl:en_keysl,
//                     en_keysr:en_keysr,
//                     de_keysl:de_keysl,
//                     de_keysr:de_keysr,
//                     old_rawkey0:old_rawkey0,
//                     old_rawkey1:old_rawkey1,
//                     output:output,
//                 }
//             }
func (this *PhpCryptExtendedData) GetInitialized() int      { return this.initialized }
func (this *PhpCryptExtendedData) SetInitialized(value int) { this.initialized = value }
func (this *PhpCryptExtendedData) GetSaltbits() uint32      { return this.saltbits }
func (this *PhpCryptExtendedData) SetSaltbits(value uint32) { this.saltbits = value }
func (this *PhpCryptExtendedData) GetOldSalt() uint32       { return this.old_salt }
func (this *PhpCryptExtendedData) SetOldSalt(value uint32)  { this.old_salt = value }
func (this *PhpCryptExtendedData) GetEnKeysl() []uint32     { return this.en_keysl }

// func (this *PhpCryptExtendedData) SetEnKeysl(value []uint32) { this.en_keysl = value }
func (this *PhpCryptExtendedData) GetEnKeysr() []uint32 { return this.en_keysr }

// func (this *PhpCryptExtendedData) SetEnKeysr(value []uint32) { this.en_keysr = value }
func (this *PhpCryptExtendedData) GetDeKeysl() []uint32 { return this.de_keysl }

// func (this *PhpCryptExtendedData) SetDeKeysl(value []uint32) { this.de_keysl = value }
func (this *PhpCryptExtendedData) GetDeKeysr() []uint32 { return this.de_keysr }

// func (this *PhpCryptExtendedData) SetDeKeysr(value []uint32) { this.de_keysr = value }
func (this *PhpCryptExtendedData) GetOldRawkey0() uint32      { return this.old_rawkey0 }
func (this *PhpCryptExtendedData) SetOldRawkey0(value uint32) { this.old_rawkey0 = value }
func (this *PhpCryptExtendedData) GetOldRawkey1() uint32      { return this.old_rawkey1 }
func (this *PhpCryptExtendedData) SetOldRawkey1(value uint32) { this.old_rawkey1 = value }
func (this *PhpCryptExtendedData) GetOutput() []byte          { return this.output }

// func (this *PhpCryptExtendedData) SetOutput(value []byte) { this.output = value }
