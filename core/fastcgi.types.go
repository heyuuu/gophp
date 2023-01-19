// <<generate>>

package core

/**
 * FcgiHeader
 */
type FcgiHeader struct {
	version         uint8
	type_           uint8
	requestIdB1     uint8
	requestIdB0     uint8
	contentLengthB1 uint8
	contentLengthB0 uint8
	paddingLength   uint8
	reserved        uint8
}

func (this FcgiHeader) GetVersion() uint8               { return this.version }
func (this *FcgiHeader) SetVersion(value uint8)         { this.version = value }
func (this FcgiHeader) GetType() uint8                  { return this.type_ }
func (this *FcgiHeader) SetType(value uint8)            { this.type_ = value }
func (this FcgiHeader) GetRequestIdB1() uint8           { return this.requestIdB1 }
func (this *FcgiHeader) SetRequestIdB1(value uint8)     { this.requestIdB1 = value }
func (this FcgiHeader) GetRequestIdB0() uint8           { return this.requestIdB0 }
func (this *FcgiHeader) SetRequestIdB0(value uint8)     { this.requestIdB0 = value }
func (this FcgiHeader) GetContentLengthB1() uint8       { return this.contentLengthB1 }
func (this *FcgiHeader) SetContentLengthB1(value uint8) { this.contentLengthB1 = value }
func (this FcgiHeader) GetContentLengthB0() uint8       { return this.contentLengthB0 }
func (this *FcgiHeader) SetContentLengthB0(value uint8) { this.contentLengthB0 = value }
func (this FcgiHeader) GetPaddingLength() uint8         { return this.paddingLength }
func (this *FcgiHeader) SetPaddingLength(value uint8)   { this.paddingLength = value }
func (this FcgiHeader) GetReserved() uint8              { return this.reserved }
func (this *FcgiHeader) SetReserved(value uint8)        { this.reserved = value }

/**
 * FcgiBeginRequest
 */
type FcgiBeginRequest struct {
	roleB1   uint8
	roleB0   uint8
	flags    uint8
	reserved []uint8
}

func (this FcgiBeginRequest) GetRoleB1() uint8           { return this.roleB1 }
func (this *FcgiBeginRequest) SetRoleB1(value uint8)     { this.roleB1 = value }
func (this FcgiBeginRequest) GetRoleB0() uint8           { return this.roleB0 }
func (this *FcgiBeginRequest) SetRoleB0(value uint8)     { this.roleB0 = value }
func (this FcgiBeginRequest) GetFlags() uint8            { return this.flags }
func (this *FcgiBeginRequest) SetFlags(value uint8)      { this.flags = value }
func (this FcgiBeginRequest) GetReserved() []uint8       { return this.reserved }
func (this *FcgiBeginRequest) SetReserved(value []uint8) { this.reserved = value }

/**
 * FcgiBeginRequestRec
 */
type FcgiBeginRequestRec struct {
	hdr  FcgiHeader
	body FcgiBeginRequest
}

func (this FcgiBeginRequestRec) GetHdr() FcgiHeader              { return this.hdr }
func (this *FcgiBeginRequestRec) SetHdr(value FcgiHeader)        { this.hdr = value }
func (this FcgiBeginRequestRec) GetBody() FcgiBeginRequest       { return this.body }
func (this *FcgiBeginRequestRec) SetBody(value FcgiBeginRequest) { this.body = value }

/**
 * FcgiEndRequest
 */
type FcgiEndRequest struct {
	appStatusB3    uint8
	appStatusB2    uint8
	appStatusB1    uint8
	appStatusB0    uint8
	protocolStatus uint8
	reserved       []uint8
}

func (this FcgiEndRequest) GetAppStatusB3() uint8          { return this.appStatusB3 }
func (this *FcgiEndRequest) SetAppStatusB3(value uint8)    { this.appStatusB3 = value }
func (this FcgiEndRequest) GetAppStatusB2() uint8          { return this.appStatusB2 }
func (this *FcgiEndRequest) SetAppStatusB2(value uint8)    { this.appStatusB2 = value }
func (this FcgiEndRequest) GetAppStatusB1() uint8          { return this.appStatusB1 }
func (this *FcgiEndRequest) SetAppStatusB1(value uint8)    { this.appStatusB1 = value }
func (this FcgiEndRequest) GetAppStatusB0() uint8          { return this.appStatusB0 }
func (this *FcgiEndRequest) SetAppStatusB0(value uint8)    { this.appStatusB0 = value }
func (this FcgiEndRequest) GetProtocolStatus() uint8       { return this.protocolStatus }
func (this *FcgiEndRequest) SetProtocolStatus(value uint8) { this.protocolStatus = value }
func (this FcgiEndRequest) GetReserved() []uint8           { return this.reserved }
func (this *FcgiEndRequest) SetReserved(value []uint8)     { this.reserved = value }

/**
 * FcgiEndRequestRec
 */
type FcgiEndRequestRec struct {
	hdr  FcgiHeader
	body FcgiEndRequest
}

func (this FcgiEndRequestRec) GetHdr() FcgiHeader            { return this.hdr }
func (this *FcgiEndRequestRec) SetHdr(value FcgiHeader)      { this.hdr = value }
func (this FcgiEndRequestRec) GetBody() FcgiEndRequest       { return this.body }
func (this *FcgiEndRequestRec) SetBody(value FcgiEndRequest) { this.body = value }

/**
 * FcgiHashBucket
 */
type FcgiHashBucket struct {
	hash_value uint
	var_len    uint
	var_       *byte
	val_len    uint
	val        *byte
	next       *FcgiHashBucket
	list_next  *FcgiHashBucket
}

func (this FcgiHashBucket) GetHashValue() uint                 { return this.hash_value }
func (this *FcgiHashBucket) SetHashValue(value uint)           { this.hash_value = value }
func (this FcgiHashBucket) GetVarLen() uint                    { return this.var_len }
func (this *FcgiHashBucket) SetVarLen(value uint)              { this.var_len = value }
func (this FcgiHashBucket) GetVar() *byte                      { return this.var_ }
func (this *FcgiHashBucket) SetVar(value *byte)                { this.var_ = value }
func (this FcgiHashBucket) GetValLen() uint                    { return this.val_len }
func (this *FcgiHashBucket) SetValLen(value uint)              { this.val_len = value }
func (this FcgiHashBucket) GetVal() *byte                      { return this.val }
func (this *FcgiHashBucket) SetVal(value *byte)                { this.val = value }
func (this FcgiHashBucket) GetNext() *FcgiHashBucket           { return this.next }
func (this *FcgiHashBucket) SetNext(value *FcgiHashBucket)     { this.next = value }
func (this FcgiHashBucket) GetListNext() *FcgiHashBucket       { return this.list_next }
func (this *FcgiHashBucket) SetListNext(value *FcgiHashBucket) { this.list_next = value }

/**
 * FcgiHashBuckets
 */
type FcgiHashBuckets struct {
	idx  uint
	next *FcgiHashBuckets
	data []FcgiHashBucket
}

func (this FcgiHashBuckets) GetIdx() uint                    { return this.idx }
func (this *FcgiHashBuckets) SetIdx(value uint)              { this.idx = value }
func (this FcgiHashBuckets) GetNext() *FcgiHashBuckets       { return this.next }
func (this *FcgiHashBuckets) SetNext(value *FcgiHashBuckets) { this.next = value }
func (this FcgiHashBuckets) GetData() []FcgiHashBucket       { return this.data }
func (this *FcgiHashBuckets) SetData(value []FcgiHashBucket) { this.data = value }

/**
 * FcgiDataSeg
 */
type FcgiDataSeg struct {
	pos  *byte
	end  *byte
	next *FcgiDataSeg
	data []byte
}

func (this FcgiDataSeg) GetPos() *byte               { return this.pos }
func (this *FcgiDataSeg) SetPos(value *byte)         { this.pos = value }
func (this FcgiDataSeg) GetEnd() *byte               { return this.end }
func (this *FcgiDataSeg) SetEnd(value *byte)         { this.end = value }
func (this FcgiDataSeg) GetNext() *FcgiDataSeg       { return this.next }
func (this *FcgiDataSeg) SetNext(value *FcgiDataSeg) { this.next = value }
func (this FcgiDataSeg) GetData() []byte             { return this.data }
func (this *FcgiDataSeg) SetData(value []byte)       { this.data = value }

/**
 * FcgiHash
 */
type FcgiHash struct {
	hash_table []*FcgiHashBucket
	list       *FcgiHashBucket
	buckets    *FcgiHashBuckets
	data       *FcgiDataSeg
}

func (this FcgiHash) GetHashTable() []*FcgiHashBucket       { return this.hash_table }
func (this *FcgiHash) SetHashTable(value []*FcgiHashBucket) { this.hash_table = value }
func (this FcgiHash) GetList() *FcgiHashBucket              { return this.list }
func (this *FcgiHash) SetList(value *FcgiHashBucket)        { this.list = value }
func (this FcgiHash) GetBuckets() *FcgiHashBuckets          { return this.buckets }
func (this *FcgiHash) SetBuckets(value *FcgiHashBuckets)    { this.buckets = value }
func (this FcgiHash) GetData() *FcgiDataSeg                 { return this.data }
func (this *FcgiHash) SetData(value *FcgiDataSeg)           { this.data = value }

/**
 * FcgiReqHook
 */
type FcgiReqHook struct {
	on_accept func()
	on_read   func()
	on_close  func()
}

func (this FcgiReqHook) GetOnAccept() func()       { return this.on_accept }
func (this *FcgiReqHook) SetOnAccept(value func()) { this.on_accept = value }
func (this FcgiReqHook) GetOnRead() func()         { return this.on_read }
func (this *FcgiReqHook) SetOnRead(value func())   { this.on_read = value }
func (this FcgiReqHook) GetOnClose() func()        { return this.on_close }
func (this *FcgiReqHook) SetOnClose(value func())  { this.on_close = value }

/**
 * FcgiRequest
 */
type FcgiRequest struct {
	listen_socket int
	tcp           int
	fd            int
	id            int
	keep          int
	ended         int
	in_len        int
	in_pad        int
	out_hdr       *FcgiHeader
	out_pos       *uint8
	out_buf       []uint8
	reserved      []uint8
	hook          FcgiReqHook
	has_env       int
	env           FcgiHash
}

func (this FcgiRequest) GetListenSocket() int         { return this.listen_socket }
func (this *FcgiRequest) SetListenSocket(value int)   { this.listen_socket = value }
func (this FcgiRequest) GetTcp() int                  { return this.tcp }
func (this *FcgiRequest) SetTcp(value int)            { this.tcp = value }
func (this FcgiRequest) GetFd() int                   { return this.fd }
func (this *FcgiRequest) SetFd(value int)             { this.fd = value }
func (this FcgiRequest) GetId() int                   { return this.id }
func (this *FcgiRequest) SetId(value int)             { this.id = value }
func (this FcgiRequest) GetKeep() int                 { return this.keep }
func (this *FcgiRequest) SetKeep(value int)           { this.keep = value }
func (this FcgiRequest) GetEnded() int                { return this.ended }
func (this *FcgiRequest) SetEnded(value int)          { this.ended = value }
func (this FcgiRequest) GetInLen() int                { return this.in_len }
func (this *FcgiRequest) SetInLen(value int)          { this.in_len = value }
func (this FcgiRequest) GetInPad() int                { return this.in_pad }
func (this *FcgiRequest) SetInPad(value int)          { this.in_pad = value }
func (this FcgiRequest) GetOutHdr() *FcgiHeader       { return this.out_hdr }
func (this *FcgiRequest) SetOutHdr(value *FcgiHeader) { this.out_hdr = value }
func (this FcgiRequest) GetOutPos() *uint8            { return this.out_pos }
func (this *FcgiRequest) SetOutPos(value *uint8)      { this.out_pos = value }
func (this FcgiRequest) GetOutBuf() []uint8           { return this.out_buf }
func (this *FcgiRequest) SetOutBuf(value []uint8)     { this.out_buf = value }
func (this FcgiRequest) GetReserved() []uint8         { return this.reserved }
func (this *FcgiRequest) SetReserved(value []uint8)   { this.reserved = value }
func (this FcgiRequest) GetHook() FcgiReqHook         { return this.hook }
func (this *FcgiRequest) SetHook(value FcgiReqHook)   { this.hook = value }
func (this FcgiRequest) GetHasEnv() int               { return this.has_env }
func (this *FcgiRequest) SetHasEnv(value int)         { this.has_env = value }
func (this FcgiRequest) GetEnv() FcgiHash             { return this.env }
func (this *FcgiRequest) SetEnv(value FcgiHash)       { this.env = value }

/**
 * SaT
 */
type SaT struct /* union */ {
	sa       __struct__sockaddr
	sa_unix  __struct__sockaddr_un
	sa_inet  __struct__sockaddr_in
	sa_inet6 __struct__sockaddr_in6
}

func (this SaT) GetSa() __struct__sockaddr                { return this.sa }
func (this *SaT) SetSa(value __struct__sockaddr)          { this.sa = value }
func (this SaT) GetSaUnix() __struct__sockaddr_un         { return this.sa_unix }
func (this *SaT) SetSaUnix(value __struct__sockaddr_un)   { this.sa_unix = value }
func (this SaT) GetSaInet() __struct__sockaddr_in         { return this.sa_inet }
func (this *SaT) SetSaInet(value __struct__sockaddr_in)   { this.sa_inet = value }
func (this SaT) GetSaInet6() __struct__sockaddr_in6       { return this.sa_inet6 }
func (this *SaT) SetSaInet6(value __struct__sockaddr_in6) { this.sa_inet6 = value }
