// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/zend"
)

const MULTIPART_CONTENT_TYPE = "multipart/form-data"
const MULTIPART_EVENT_START = 0
const MULTIPART_EVENT_FORMDATA = 1
const MULTIPART_EVENT_FILE_START = 2
const MULTIPART_EVENT_FILE_DATA = 3
const MULTIPART_EVENT_FILE_END = 4
const MULTIPART_EVENT_END = 5

type PhpRfc1867EncodingTranslationT func() int
type PhpRfc1867GetDetectOrderT func(list ***zend.ZendEncoding, list_size *int)
type PhpRfc1867SetInputEncodingT func(encoding *zend.ZendEncoding)
type PhpRfc1867GetwordT func(encoding *zend.ZendEncoding, line **byte, stop byte) *byte
type PhpRfc1867GetwordConfT func(encoding *zend.ZendEncoding, str *byte) *byte
type PhpRfc1867BasenameT func(encoding *zend.ZendEncoding, str *byte) *byte

const DEBUG_FILE_UPLOAD = 0

var PhpRfc1867EncodingTranslation PhpRfc1867EncodingTranslationT = DummyEncodingTranslation
var PhpRfc1867GetDetectOrder PhpRfc1867GetDetectOrderT = nil
var PhpRfc1867SetInputEncoding PhpRfc1867SetInputEncodingT = nil
var PhpRfc1867Getword PhpRfc1867GetwordT = PhpApGetword
var PhpRfc1867GetwordConf PhpRfc1867GetwordConfT = PhpApGetwordConf
var PhpRfc1867Basename PhpRfc1867BasenameT = nil
var PhpRfc1867Callback func(event uint, event_data any, extra *any) int = nil

const MAX_SIZE_OF_INDEX = b.SizeOf("\"[tmp_name]\"")
const MAX_SIZE_ANONNAME = 33
const UPLOAD_ERROR_OK = 0
const UPLOAD_ERROR_A = 1
const UPLOAD_ERROR_B = 2
const UPLOAD_ERROR_C = 3
const UPLOAD_ERROR_D = 4
const UPLOAD_ERROR_E = 6
const UPLOAD_ERROR_F = 7
const UPLOAD_ERROR_X = 8
const FILLUNIT = 1024 * 5
