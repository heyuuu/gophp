package core

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * SapiHeader
 */
type SapiHeader struct {
	header     *byte
	header_len int
}

func (this *SapiHeader) GetHeader() *byte       { return this.header }
func (this *SapiHeader) SetHeader(value *byte)  { this.header = value }
func (this *SapiHeader) GetHeaderLen() int      { return this.header_len }
func (this *SapiHeader) SetHeaderLen(value int) { this.header_len = value }

/**
 * SapiHeaders
 */
type SapiHeaders struct {
	headers                zend.ZendLlist[*SapiHeader]
	httpResponseCode       int
	sendDefaultContentType uint8
	mimetype               *byte
	httpStatusLine         *byte
}

func (sh *SapiHeaders) Init() {
	sh.headers.Init()
	sh.sendDefaultContentType = 1
	//sh.httpResponseCode = 200
	sh.httpStatusLine = nil
	sh.mimetype = nil
}

func (sh *SapiHeaders) GetHeaders() zend.ZendLlist { return sh.headers }

/**
 * SapiRequestInfo
 */
type SapiRequestInfo struct {
	requestMethod  string
	queryString    string
	cookieData     string
	contentLength  int
	pathTranslated string
	requestUri     string
	requestBody    *PhpStream
	contentType    string
	headersOnly    bool
	noHeaders      bool
	headersRead    bool
	postEntry      *SapiPostEntry
	contentTypeDup string
	authUser       string
	authPassword   string
	authDigest     string
	argv0          string
	protoNum       int
	args           []string
}

func (info *SapiRequestInfo) InitEmpty() {
	info.requestMethod = ""
	info.authPassword = ""
	info.authUser = ""
	info.authDigest = ""
	info.contentTypeDup = ""
}

func (info *SapiRequestInfo) IsRequestMethod(method string) bool {
	return ascii.StrCaseEquals(info.requestMethod, method)
}

func (info *SapiRequestInfo) RequestMethod() string  { return info.requestMethod }
func (info *SapiRequestInfo) QueryString() string    { return info.queryString }
func (info *SapiRequestInfo) CookieData() string     { return info.cookieData }
func (info *SapiRequestInfo) ContentLength() int     { return info.contentLength }
func (info *SapiRequestInfo) PathTranslated() string { return info.pathTranslated }
func (info *SapiRequestInfo) RequestUri() string     { return info.requestUri }
func (info *SapiRequestInfo) ContentType() string    { return info.contentType }
func (info *SapiRequestInfo) ContentTypeDup() string { return info.contentTypeDup }
func (info *SapiRequestInfo) AuthUser() string       { return info.authUser }
func (info *SapiRequestInfo) AuthPassword() string   { return info.authPassword }
func (info *SapiRequestInfo) AuthDigest() string     { return info.authDigest }
func (info *SapiRequestInfo) Argv0() string          { return info.argv0 }
func (info *SapiRequestInfo) Args() []string         { return info.args }
func (info *SapiRequestInfo) Argc() int              { return len(info.args) }

func (info *SapiRequestInfo) SetRequestMethod(value string)  { info.requestMethod = value }
func (info *SapiRequestInfo) SetQueryString(value string)    { info.queryString = value }
func (info *SapiRequestInfo) SetCookieData(value string)     { info.cookieData = value }
func (info *SapiRequestInfo) SetContentLength(value int)     { info.contentLength = value }
func (info *SapiRequestInfo) SetPathTranslated(value string) { info.pathTranslated = value }
func (info *SapiRequestInfo) SetRequestUri(value string)     { info.requestUri = value }
func (info *SapiRequestInfo) SetContentType(value string)    { info.contentType = value }
func (info *SapiRequestInfo) SetContentTypeDup(value string) { info.contentType = value }
func (info *SapiRequestInfo) SetAuthUser(value string)       { info.authUser = value }
func (info *SapiRequestInfo) GetAuthPassword() string        { return info.authPassword }
func (info *SapiRequestInfo) SetAuthPassword(value string)   { info.authPassword = value }
func (info *SapiRequestInfo) SetAuthDigest(value string)     { info.authDigest = value }
func (info *SapiRequestInfo) SetProtoNum(value int)          { info.protoNum = value }

func (info *SapiRequestInfo) SetArgs(args []string) { info.args = args }

/**
 * SapiGlobals
 */
type SapiGlobals struct {
	serverContext         any
	RequestInfo           SapiRequestInfo
	sapiHeaders           SapiHeaders
	readPostBytes         int64
	postRead              uint8
	headersSent           bool
	globalStat            zend.ZendStatT
	defaultMimetype       *byte
	defaultCharset        *byte
	rfc1867UploadedFiles  map[string]bool
	postMaxSize           zend.ZendLong
	options               int
	sapiStarted           bool
	globalRequestTime     float64
	knownPostContentTypes types.Array
	callbackFunc          types.Zval
	fciCache              types.ZendFcallInfoCache
}

func (sg *SapiGlobals) Activate() {
	sg.sapiHeaders.Init()
	sg.headersSent = false
	sg.callbackFunc.SetUndef()
	sg.readPostBytes = 0
	sg.globalRequestTime = 0
	sg.postRead = 0
	sg.rfc1867UploadedFiles = nil

	sg.RequestInfo.requestBody = nil
	sg.RequestInfo.noHeaders = false
	sg.RequestInfo.postEntry = nil
	sg.RequestInfo.protoNum = 1000

	/* It's possible to override this general case in the activate() callback, if necessary. */
	sg.RequestInfo.headersOnly = ascii.StrCaseEquals(sg.RequestInfo.requestMethod, "HEAD")
}

func (sg *SapiGlobals) Init() {
	sg.knownPostContentTypes = *types.NewArrayCap(8)
	PhpSetupSapiContentTypes()
}

func (sg *SapiGlobals) Destroy() {
	sg.knownPostContentTypes.Destroy()
}

func (sg *SapiGlobals) ResetUploadFiles() {
	sg.rfc1867UploadedFiles = nil
}

func (sg *SapiGlobals) AddUploadFile(path string) {
	if sg.rfc1867UploadedFiles == nil {
		sg.rfc1867UploadedFiles = make(map[string]bool)
	}
	sg.rfc1867UploadedFiles[path] = true
}
func (sg *SapiGlobals) ExistUploadFile(path string) bool {
	return sg.rfc1867UploadedFiles[path]
}

func (sg *SapiGlobals) DeleteUploadFile(path string) {
	delete(sg.rfc1867UploadedFiles, path)
}

/**
 * generate
 */
func (sg *SapiGlobals) GetDefaultMimetype() *byte     { return sg.defaultMimetype }
func (sg *SapiGlobals) GetDefaultCharset() *byte      { return sg.defaultCharset }
func (sg *SapiGlobals) GetPostMaxSize() zend.ZendLong { return sg.postMaxSize }
func (sg *SapiGlobals) GetKnownPostContentTypes() types.Array {
	return sg.knownPostContentTypes
}

/**
 * SapiHeaderLine
 */
type SapiHeaderLine struct {
	line          *byte
	line_len      int
	response_code zend.ZendLong
}

func MakeSapiHeaderLineEx(line string) SapiHeaderLine {
	return SapiHeaderLine{}
}

func MakeSapiHeaderLine(line *byte, line_len int, response_code zend.ZendLong) SapiHeaderLine {
	return SapiHeaderLine{
		line:          line,
		line_len:      line_len,
		response_code: response_code,
	}
}
func (this *SapiHeaderLine) GetLine() *byte                 { return this.line }
func (this *SapiHeaderLine) SetLine(value *byte)            { this.line = value }
func (this *SapiHeaderLine) GetLineLen() int                { return this.line_len }
func (this *SapiHeaderLine) SetLineLen(value int)           { this.line_len = value }
func (this *SapiHeaderLine) GetResponseCode() zend.ZendLong { return this.response_code }

/**
 * SapiPostEntry
 */
type SapiPostEntry struct {
	contentType string
	postReader  func()
	postHandler func(contentTypeDup string, arg *types.Zval)
}

func MakeSapiPostEntry(contentType string, postReader func(), postHandler func(contentTypeDup string, arg *types.Zval)) SapiPostEntry {
	return SapiPostEntry{
		contentType: contentType,
		postReader:  postReader,
		postHandler: postHandler,
	}
}
func (pe *SapiPostEntry) ContentType() string { return pe.contentType }
func (pe *SapiPostEntry) PostReader() {
	if pe.postReader != nil {
		pe.postReader()
	}
}
func (pe *SapiPostEntry) PostHandler(contentType string, arg *types.Zval) {
	pe.postHandler(contentType, arg)
}
