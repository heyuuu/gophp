package core

import (
	b "github.com/heyuuu/gophp/builtin"
)

const DEFAULT_POST_CONTENT_TYPE = "application/x-www-form-urlencoded"

var PhpPostEntries []SapiPostEntry = []SapiPostEntry{
	MakeSapiPostEntry(DEFAULT_POST_CONTENT_TYPE, b.SizeOf("DEFAULT_POST_CONTENT_TYPE")-1, SapiReadStandardFormData, PhpStdPostHandler),
	MakeSapiPostEntry(MULTIPART_CONTENT_TYPE, b.SizeOf("MULTIPART_CONTENT_TYPE")-1, nil, Rfc1867PostHandler),
	MakeSapiPostEntry(nil, 0, nil, nil),
}
