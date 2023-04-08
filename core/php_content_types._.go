package core

const DEFAULT_POST_CONTENT_TYPE = "application/x-www-form-urlencoded"

var PhpPostEntries = []SapiPostEntry{
	MakeSapiPostEntry(DEFAULT_POST_CONTENT_TYPE, SapiReadStandardFormData, PhpStdPostHandler),
	MakeSapiPostEntry(MULTIPART_CONTENT_TYPE, nil, Rfc1867PostHandler),
}
