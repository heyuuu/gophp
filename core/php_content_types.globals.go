// <<generate>>

package core

import (
	b "sik/builtin"
)

const DEFAULT_POST_CONTENT_TYPE = "application/x-www-form-urlencoded"

var PhpPostEntries []SapiPostEntry = []SapiPostEntry{
	{
		DEFAULT_POST_CONTENT_TYPE,
		b.SizeOf("DEFAULT_POST_CONTENT_TYPE") - 1,
		SapiReadStandardFormData,
		PhpStdPostHandler,
	},
	{MULTIPART_CONTENT_TYPE, b.SizeOf("MULTIPART_CONTENT_TYPE") - 1, nil, Rfc1867PostHandler},
	{nil, 0, nil, nil},
}
