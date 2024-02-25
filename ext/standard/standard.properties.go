package standard

import (
	"github.com/heyuuu/gophp/php/streams"
	"time"
)

// properties for EntityStage3Row
func (t *EntityStage3Row) Ambiguous() bool {
	return t.ambiguous
}
func (t *EntityStage3Row) Entity() string {
	return t.entity
}
func (t *EntityStage3Row) MultiCodepointTable() []EntityMulticodepointRow {
	return t.multiCodepointTable
}

// properties for EntityTableOpt
func (t *EntityTableOpt) MsTable() []EntityStage1Row {
	return t.msTable
}
func (t *EntityTableOpt) Table() []*EntityStage3Row {
	return t.table
}

// properties for FileGlobals
func (t *FileGlobals) PcloseRet() int {
	return t.pcloseRet
}
func (t *FileGlobals) SetPcloseRet(v int) {
	t.pcloseRet = v
}
func (t *FileGlobals) DefChunkSize() int {
	return t.defChunkSize
}
func (t *FileGlobals) SetDefChunkSize(v int) {
	t.defChunkSize = v
}
func (t *FileGlobals) AutoDetectLineEndings() bool {
	return t.autoDetectLineEndings
}
func (t *FileGlobals) SetAutoDetectLineEndings(v bool) {
	t.autoDetectLineEndings = v
}
func (t *FileGlobals) DefaultSocketTimeout() time.Duration {
	return t.defaultSocketTimeout
}
func (t *FileGlobals) SetDefaultSocketTimeout(v time.Duration) {
	t.defaultSocketTimeout = v
}
func (t *FileGlobals) UserAgent() string {
	return t.userAgent
}
func (t *FileGlobals) SetUserAgent(v string) {
	t.userAgent = v
}
func (t *FileGlobals) FromAddress() string {
	return t.fromAddress
}
func (t *FileGlobals) SetFromAddress(v string) {
	t.fromAddress = v
}
func (t *FileGlobals) UserStreamCurrentFilename() string {
	return t.userStreamCurrentFilename
}
func (t *FileGlobals) SetUserStreamCurrentFilename(v string) {
	t.userStreamCurrentFilename = v
}
func (t *FileGlobals) DefaultContext() *streams.StreamContext {
	return t.defaultContext
}
func (t *FileGlobals) SetDefaultContext(v *streams.StreamContext) {
	t.defaultContext = v
}

// properties for UniToEnc
func (t *UniToEnc) UnCodePoint() uint16 {
	return t.unCodePoint
}
func (t *UniToEnc) CsCode() uint8 {
	return t.csCode
}
