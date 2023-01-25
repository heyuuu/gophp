// <<generate>>

package zend

type ZendEncoding = __struct___zend_encoding
type ZendEncodingFilter func(str **uint8, str_length *int, buf *uint8, length int) int
type ZendEncodingFetcher func(encoding_name *byte) *ZendEncoding
type ZendEncodingNameGetter func(encoding *ZendEncoding) *byte
type ZendEncodingLexerCompatibilityChecker func(encoding *ZendEncoding) int
type ZendEncodingDetector func(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding
type ZendEncodingConverter func(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int
type ZendEncodingListParser func(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int
type ZendEncodingInternalEncodingGetter func() *ZendEncoding
type ZendEncodingInternalEncodingSetter func(encoding *ZendEncoding) int

var MultibyteFunctionsDummy ZendMultibyteFunctions
var MultibyteFunctions ZendMultibyteFunctions = ZendMultibyteFunctions{nil, DummyEncodingFetcher, DummyEncodingNameGetter, DummyEncodingLexerCompatibilityChecker, DummyEncodingDetector, DummyEncodingConverter, DummyEncodingListParser, DummyInternalEncodingGetter, DummyInternalEncodingSetter}
var ZendMultibyteEncodingUtf32be *ZendEncoding = (*ZendEncoding)("UTF-32BE")
var ZendMultibyteEncodingUtf32le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf16be *ZendEncoding = (*ZendEncoding)("UTF-16BE")
var ZendMultibyteEncodingUtf16le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf8 *ZendEncoding = (*ZendEncoding)("UTF-8")
