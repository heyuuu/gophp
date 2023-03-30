package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

func FCGI_HASH_FUNC(var_ __auto__, var_len int) __auto__ {
	if var_len < 3 {
		return uint(var_len)
	} else {
		return (uint(var_[3]) << 2) + (uint(var_[var_len-2]) << 4) + (uint(var_[var_len-1]) << 2) + var_len
	}
}
func FCGI_GETENV(request *FcgiRequest, name string) *byte {
	return FcgiQuickGetenv(request, name, b.SizeOf("name")-1, FCGI_HASH_FUNC(name, b.SizeOf("name")-1))
}
func FCGI_PUTENV(request *FcgiRequest, name string, value *byte) *byte {
	return FcgiQuickPutenv(request, name, b.SizeOf("name")-1, FCGI_HASH_FUNC(name, b.SizeOf("name")-1), value)
}
func FcgiHashInit(h *FcgiHash) {
	memset(h.GetHashTable(), 0, b.SizeOf("h -> hash_table"))
	h.SetList(nil)
	h.SetBuckets((*FcgiHashBuckets)(zend.Malloc(b.SizeOf("fcgi_hash_buckets"))))
	h.GetBuckets().SetIdx(0)
	h.GetBuckets().SetNext(nil)
	h.SetData((*FcgiDataSeg)(zend.Malloc(b.SizeOf("fcgi_data_seg") - 1 + FCGI_HASH_SEG_SIZE)))
	h.GetData().SetPos(h.GetData().GetData())
	h.GetData().SetEnd(h.GetData().GetPos() + FCGI_HASH_SEG_SIZE)
	h.GetData().SetNext(nil)
}
func FcgiHashDestroy(h *FcgiHash) {
	var b *FcgiHashBuckets
	var p *FcgiDataSeg
	b = h.GetBuckets()
	for b != nil {
		var q *FcgiHashBuckets = b
		b = b.GetNext()
		zend.Free(q)
	}
	p = h.GetData()
	for p != nil {
		var q *FcgiDataSeg = p
		p = p.GetNext()
		zend.Free(q)
	}
}
func FcgiHashClean(h *FcgiHash) {
	memset(h.GetHashTable(), 0, b.SizeOf("h -> hash_table"))
	h.SetList(nil)

	/* delete all bucket blocks except the first one */

	for h.GetBuckets().GetNext() != nil {
		var q *FcgiHashBuckets = h.GetBuckets()
		h.SetBuckets(h.GetBuckets().GetNext())
		zend.Free(q)
	}
	h.GetBuckets().SetIdx(0)

	/* delete all data segments except the first one */

	for h.GetData().GetNext() != nil {
		var q *FcgiDataSeg = h.GetData()
		h.SetData(h.GetData().GetNext())
		zend.Free(q)
	}
	h.GetData().SetPos(h.GetData().GetData())
}
func FcgiHashStrndup(h *FcgiHash, str *byte, str_len uint) *byte {
	var ret *byte
	if h.GetData().GetPos()+str_len+1 >= h.GetData().GetEnd() {
		var seg_size uint = b.Cond(str_len+1 > FCGI_HASH_SEG_SIZE, str_len+1, FCGI_HASH_SEG_SIZE)
		var p *FcgiDataSeg = (*FcgiDataSeg)(zend.Malloc(b.SizeOf("fcgi_data_seg") - 1 + seg_size))
		p.SetPos(p.GetData())
		p.SetEnd(p.GetPos() + seg_size)
		p.SetNext(h.GetData())
		h.SetData(p)
	}
	ret = h.GetData().GetPos()
	memcpy(ret, str, str_len)
	ret[str_len] = 0
	h.GetData().SetPos(h.GetData().GetPos() + str_len + 1)
	return ret
}
func FcgiHashSet(
	h *FcgiHash,
	hash_value uint,
	var_ *byte,
	var_len uint,
	val *byte,
	val_len uint,
) *byte {
	var idx uint = hash_value & FCGI_HASH_TABLE_MASK
	var p *FcgiHashBucket = h.GetHashTable()[idx]
	for p != nil {
		if p.GetHashValue() == hash_value && p.GetVarLen() == var_len && memcmp(p.GetVar(), var_, var_len) == 0 {
			p.SetValLen(val_len)
			p.SetVal(FcgiHashStrndup(h, val, val_len))
			return p.GetVal()
		}
		p = p.GetNext()
	}
	if h.GetBuckets().GetIdx() >= FCGI_HASH_TABLE_SIZE {
		var b *FcgiHashBuckets = (*FcgiHashBuckets)(zend.Malloc(b.SizeOf("fcgi_hash_buckets")))
		b.SetIdx(0)
		b.SetNext(h.GetBuckets())
		h.SetBuckets(b)
	}
	p = h.GetBuckets().GetData() + h.GetBuckets().GetIdx()
	h.GetBuckets().GetIdx()++
	p.SetNext(h.GetHashTable()[idx])
	h.GetHashTable()[idx] = p
	p.SetListNext(h.GetList())
	h.SetList(p)
	p.SetHashValue(hash_value)
	p.SetVarLen(var_len)
	p.SetVar(FcgiHashStrndup(h, var_, var_len))
	p.SetValLen(val_len)
	p.SetVal(FcgiHashStrndup(h, val, val_len))
	return p.GetVal()
}
func FcgiHashDel(h *FcgiHash, hash_value uint, var_ *byte, var_len uint) {
	var idx uint = hash_value & FCGI_HASH_TABLE_MASK
	var p **FcgiHashBucket = h.GetHashTable()[idx]
	for (*p) != nil {
		if p.GetHashValue() == hash_value && p.GetVarLen() == var_len && memcmp(p.GetVar(), var_, var_len) == 0 {
			p.SetVal(nil)
			p.SetValLen(0)
			*p = p.GetNext()
			return
		}
		p = p.GetNext()
	}
}
func FcgiHashGet(h *FcgiHash, hash_value uint, var_ *byte, var_len uint, val_len *uint) *byte {
	var idx uint = hash_value & FCGI_HASH_TABLE_MASK
	var p *FcgiHashBucket = h.GetHashTable()[idx]
	for p != nil {
		if p.GetHashValue() == hash_value && p.GetVarLen() == var_len && memcmp(p.GetVar(), var_, var_len) == 0 {
			*val_len = p.GetValLen()
			return p.GetVal()
		}
		p = p.GetNext()
	}
	return nil
}
func FcgiHashApply(h *FcgiHash, func_ FcgiApplyFunc, arg any) {
	var p *FcgiHashBucket = h.GetList()
	for p != nil {
		if p.GetVal() != nil {
			func_(p.GetVar(), p.GetVarLen(), p.GetVal(), p.GetValLen(), arg)
		}
		p = p.GetListNext()
	}
}
func FcgiSignalHandler(signo int) {
	if signo == SIGUSR1 || signo == SIGTERM {
		InShutdown = 1
	}
}
func FcgiSetupSignals() {
	var new_sa __struct__sigaction
	var old_sa __struct__sigaction
	sigemptyset(new_sa.sa_mask)
	new_sa.sa_flags = 0
	new_sa.sa_handler = FcgiSignalHandler
	sigaction(SIGUSR1, &new_sa, nil)
	sigaction(SIGTERM, &new_sa, nil)
	sigaction(SIGPIPE, nil, &old_sa)
	if old_sa.sa_handler == SIG_DFL {
		sigaction(SIGPIPE, &new_sa, nil)
	}
}
func FcgiInShutdown() int         { return InShutdown }
func FcgiSetLogger(lg FcgiLogger) { FcgiLog = lg }
func FcgiInit() int {
	if IsInitialized == 0 {
		var sa SaT
		var len_ socklen_t = b.SizeOf("sa")
		&FcgiMgmtVars = types.MakeArrayEx(8, FcgiFreeMgmtVarCb, 1)
		FcgiSetMgmtVar("FCGI_MPXS_CONNS", b.SizeOf("\"FCGI_MPXS_CONNS\"")-1, "0", b.SizeOf("\"0\"")-1)
		IsInitialized = 1
		errno = 0
		if getpeername(0, (*__struct__sockaddr)(&sa), &len_) != 0 && errno == ENOTCONN {
			FcgiSetupSignals()
			IsFastcgi = 1
			return IsFastcgi
		} else {
			IsFastcgi = 0
			return IsFastcgi
		}
	}
	return IsFastcgi
}
func FcgiIsFastcgi() int {
	if IsInitialized == 0 {
		return FcgiInit()
	} else {
		return IsFastcgi
	}
}
func FcgiShutdown() {
	if IsInitialized != 0 {
		FcgiMgmtVars.Destroy()
	}
	IsFastcgi = 0
	if AllowedClients != nil {
		zend.Free(AllowedClients)
	}
}
func IsPortNumber(bindpath *byte) int {
	for *bindpath {
		if (*bindpath) < '0' || (*bindpath) > '9' {
			return 0
		}
		bindpath++
	}
	return 1
}
func FcgiListen(path *byte, backlog int) int {
	var s *byte
	var tcp int = 0
	var host []byte
	var port short = 0
	var listen_socket int
	var sa SaT
	var sock_len socklen_t
	if b.Assign(&s, strchr(path, ':')) {
		port = atoi(s + 1)
		if port != 0 && s-path < MAXPATHLEN {
			strncpy(host, path, s-path)
			host[s-path] = '0'
			tcp = 1
		}
	} else if IsPortNumber(path) != 0 {
		port = atoi(path)
		if port != 0 {
			host[0] = '0'
			tcp = 1
		}
	}

	/* Prepare socket address */

	if tcp != 0 {
		memset(sa.GetSaInet(), 0, b.SizeOf("sa . sa_inet"))
		sa.GetSaInet().sin_family = AF_INET
		sa.GetSaInet().sin_port = htons(port)
		sock_len = b.SizeOf("sa . sa_inet")
		if !(*host) || !(strncmp(host, "*", b.SizeOf("\"*\"")-1)) {
			sa.GetSaInet().sin_addr.s_addr = htonl(INADDR_ANY)
		} else {
			sa.GetSaInet().sin_addr.s_addr = inet_addr(host)
			if sa.GetSaInet().sin_addr.s_addr == INADDR_NONE {
				var hep *__struct__hostent
				if strlen(host) > MAXFQDNLEN {
					hep = nil
				} else {
					hep = PhpNetworkGethostbyname(host)
				}
				if hep == nil || hep.h_addrtype != AF_INET || !(hep.h_addr_list[0]) {
					FcgiLog(FCGI_ERROR, "Cannot resolve host name '%s'!\n", host)
					return -1
				} else if hep.h_addr_list[1] {
					FcgiLog(FCGI_ERROR, "Host '%s' has multiple addresses. You must choose one explicitly!\n", host)
					return -1
				}
				sa.GetSaInet().sin_addr.s_addr = (*__struct__in_addr)(hep.h_addr_list[0]).s_addr
			}
		}
	} else {
		var path_len int = strlen(path)
		if path_len >= b.SizeOf("sa . sa_unix . sun_path") {
			FcgiLog(FCGI_ERROR, "Listening socket's path name is too long.\n")
			return -1
		}
		memset(sa.GetSaUnix(), 0, b.SizeOf("sa . sa_unix"))
		sa.GetSaUnix().sun_family = AF_UNIX
		memcpy(sa.GetSaUnix().sun_path, path, path_len+1)
		sock_len = size_t((*__struct__sockaddr_un)(0).sun_path) + path_len
		sa.GetSaUnix().sun_len = sock_len
		unlink(path)
	}

	/* Create, bind socket and start listen on it */

	if b.Assign(&listen_socket, socket(sa.GetSa().sa_family, SOCK_STREAM, 0)) < 0 || bind(listen_socket, (*__struct__sockaddr)(&sa), sock_len) < 0 || listen(listen_socket, backlog) < 0 {
		close(listen_socket)
		FcgiLog(FCGI_ERROR, "Cannot bind/listen socket - [%d] %s.\n", errno, strerror(errno))
		return -1
	}
	if tcp == 0 {
		chmod(path, 0777)
	} else {
		var ip *byte = getenv("FCGI_WEB_SERVER_ADDRS")
		var cur *byte
		var end *byte
		var n int
		if ip != nil {
			ip = strdup(ip)
			cur = ip
			n = 0
			for *cur {
				if (*cur) == ',' {
					n++
				}
				cur++
			}
			AllowedClients = zend.Malloc(b.SizeOf("sa_t") * (n + 2))
			n = 0
			cur = ip
			for cur != nil {
				end = strchr(cur, ',')
				if end != nil {
					*end = 0
					end++
				}
				if inet_pton(AF_INET, cur, AllowedClients[n].GetSaInet().sin_addr) > 0 {
					AllowedClients[n].GetSa().sa_family = AF_INET
					n++
				} else if inet_pton(AF_INET6, cur, AllowedClients[n].GetSaInet6().sin6_addr) > 0 {
					AllowedClients[n].GetSa().sa_family = AF_INET6
					n++
				} else {
					FcgiLog(FCGI_ERROR, "Wrong IP address '%s' in listen.allowed_clients", cur)
				}
				cur = end
			}
			AllowedClients[n].GetSa().sa_family = 0
			zend.Free(ip)
			if n == 0 {
				FcgiLog(FCGI_ERROR, "There are no allowed addresses")
			}
		}
	}
	if IsInitialized == 0 {
		FcgiInit()
	}
	IsFastcgi = 1
	FcgiSetupSignals()
	return listen_socket
}
func FcgiSetAllowedClients(ip *byte) {
	var cur *byte
	var end *byte
	var n int
	if ip != nil {
		ip = strdup(ip)
		cur = ip
		n = 0
		for *cur {
			if (*cur) == ',' {
				n++
			}
			cur++
		}
		if AllowedClients != nil {
			zend.Free(AllowedClients)
		}
		AllowedClients = zend.Malloc(b.SizeOf("sa_t") * (n + 2))
		n = 0
		cur = ip
		for cur != nil {
			end = strchr(cur, ',')
			if end != nil {
				*end = 0
				end++
			}
			if inet_pton(AF_INET, cur, AllowedClients[n].GetSaInet().sin_addr) > 0 {
				AllowedClients[n].GetSa().sa_family = AF_INET
				n++
			} else if inet_pton(AF_INET6, cur, AllowedClients[n].GetSaInet6().sin6_addr) > 0 {
				AllowedClients[n].GetSa().sa_family = AF_INET6
				n++
			} else {
				FcgiLog(FCGI_ERROR, "Wrong IP address '%s' in listen.allowed_clients", cur)
			}
			cur = end
		}
		AllowedClients[n].GetSa().sa_family = 0
		zend.Free(ip)
		if n == 0 {
			FcgiLog(FCGI_ERROR, "There are no allowed addresses")
		}
	}
}
func FcgiHookDummy() { return }
func FcgiInitRequest(listen_socket int, on_accept func(), on_read func(), on_close func()) *FcgiRequest {
	var req *FcgiRequest = calloc(1, b.SizeOf("fcgi_request"))
	req.SetListenSocket(listen_socket)
	req.SetFd(-1)
	req.SetId(-1)

	/*
	       req->in_len = 0;
	       req->in_pad = 0;

	       req->out_hdr = NULL;

	   #ifdef TCP_NODELAY
	       req->nodelay = 0;
	   #endif

	       req->env = NULL;
	       req->has_env = 0;

	*/

	req.SetOutPos(req.GetOutBuf())
	if on_accept != nil {
		req.GetHook().SetOnAccept(on_accept)
	} else {
		req.GetHook().SetOnAccept(FcgiHookDummy)
	}
	if on_read != nil {
		req.GetHook().SetOnRead(on_read)
	} else {
		req.GetHook().SetOnRead(FcgiHookDummy)
	}
	if on_close != nil {
		req.GetHook().SetOnClose(on_close)
	} else {
		req.GetHook().SetOnClose(FcgiHookDummy)
	}
	FcgiHashInit(req.GetEnv())
	return req
}
func FcgiDestroyRequest(req *FcgiRequest) {
	FcgiHashDestroy(req.GetEnv())
	zend.Free(req)
}
func SafeWrite(req *FcgiRequest, buf any, count int) ssize_t {
	var ret int
	var n int = 0
	for {
		errno = 0
		ret = write(req.GetFd(), (*byte)(buf)+n, count-n)
		if ret > 0 {
			n += ret
		} else if ret <= 0 && errno != 0 && errno != EINTR {
			return ret
		}
		if n == count {
			break
		}
	}
	return n
}
func SafeRead(req *FcgiRequest, buf any, count int) ssize_t {
	var ret int
	var n int = 0
	for {
		errno = 0
		ret = read(req.GetFd(), (*byte)(buf)+n, count-n)
		if ret > 0 {
			n += ret
		} else if ret == 0 && errno == 0 {
			return n
		} else if ret <= 0 && errno != 0 && errno != EINTR {
			return ret
		}
		if n == count {
			break
		}
	}
	return n
}
func FcgiMakeHeader(hdr *FcgiHeader, type_ FcgiRequestType, req_id int, len_ int) int {
	var pad int = (len_ + 7 & ^7) - len_
	hdr.SetContentLengthB0(uint8(len_ & 0xff))
	hdr.SetContentLengthB1(uint8(len_ >> 8 & 0xff))
	hdr.SetPaddingLength(uint8(pad))
	hdr.SetRequestIdB0(uint8(req_id & 0xff))
	hdr.SetRequestIdB1(uint8(req_id >> 8 & 0xff))
	hdr.SetReserved(0)
	hdr.SetType(type_)
	hdr.SetVersion(FCGI_VERSION_1)
	if pad != 0 {
		memset((*uint8)(hdr)+b.SizeOf("fcgi_header")+len_, 0, pad)
	}
	return pad
}
func FcgiGetParams(req *FcgiRequest, p *uint8, end *uint8) int {
	var name_len uint
	var val_len uint
	for p < end {
		*p++
		name_len = (*p) - 1
		if name_len >= 128 {
			if p+3 >= end {
				return 0
			}
			name_len = (name_len & 0x7f) << 24
			name_len |= b.PostInc(&(*p)) << 16
			name_len |= b.PostInc(&(*p)) << 8
			*p++
			name_len |= (*p) - 1
		}
		if p >= end {
			return 0
		}
		*p++
		val_len = (*p) - 1
		if val_len >= 128 {
			if p+3 >= end {
				return 0
			}
			val_len = (val_len & 0x7f) << 24
			val_len |= b.PostInc(&(*p)) << 16
			val_len |= b.PostInc(&(*p)) << 8
			*p++
			val_len |= (*p) - 1
		}
		if name_len+val_len > uint(end-p) {

			/* Malformated request */

			return 0

			/* Malformated request */

		}
		FcgiHashSet(req.GetEnv(), FCGI_HASH_FUNC(p, name_len), (*byte)(p), name_len, (*byte)(p+name_len), val_len)
		p += name_len + val_len
	}
	return 1
}
func FcgiReadRequest(req *FcgiRequest) int {
	var hdr FcgiHeader
	var len_ int
	var padding int
	var buf []uint8
	req.SetKeep(0)
	req.SetEnded(0)
	req.SetInLen(0)
	req.SetOutHdr(nil)
	req.SetOutPos(req.GetOutBuf())
	if req.GetHasEnv() != 0 {
		FcgiHashClean(req.GetEnv())
	} else {
		req.SetHasEnv(1)
	}
	if SafeRead(req, &hdr, b.SizeOf("fcgi_header")) != b.SizeOf("fcgi_header") || hdr.GetVersion() < FCGI_VERSION_1 {
		return 0
	}
	len_ = hdr.GetContentLengthB1()<<8 | hdr.GetContentLengthB0()
	padding = hdr.GetPaddingLength()
	for hdr.GetType() == FCGI_STDIN && len_ == 0 {
		if SafeRead(req, &hdr, b.SizeOf("fcgi_header")) != b.SizeOf("fcgi_header") || hdr.GetVersion() < FCGI_VERSION_1 {
			return 0
		}
		len_ = hdr.GetContentLengthB1()<<8 | hdr.GetContentLengthB0()
		padding = hdr.GetPaddingLength()
	}
	if len_+padding > FCGI_MAX_LENGTH {
		return 0
	}
	req.SetId((hdr.GetRequestIdB1() << 8) + hdr.GetRequestIdB0())
	if hdr.GetType() == FCGI_BEGIN_REQUEST && len_ == b.SizeOf("fcgi_begin_request") {
		var b *FcgiBeginRequest
		if SafeRead(req, buf, len_+padding) != len_+padding {
			return 0
		}
		b = (*FcgiBeginRequest)(buf)
		req.SetKeep(b.GetFlags() & FCGI_KEEP_CONN)
		switch (b.GetRoleB1() << 8) + b.GetRoleB0() {
		case FCGI_RESPONDER:
			FcgiHashSet(req.GetEnv(), FCGI_HASH_FUNC("FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1), "FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1, "RESPONDER", b.SizeOf("\"RESPONDER\"")-1)
		case FCGI_AUTHORIZER:
			FcgiHashSet(req.GetEnv(), FCGI_HASH_FUNC("FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1), "FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1, "AUTHORIZER", b.SizeOf("\"AUTHORIZER\"")-1)
		case FCGI_FILTER:
			FcgiHashSet(req.GetEnv(), FCGI_HASH_FUNC("FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1), "FCGI_ROLE", b.SizeOf("\"FCGI_ROLE\"")-1, "FILTER", b.SizeOf("\"FILTER\"")-1)
		default:
			return 0
		}
		if SafeRead(req, &hdr, b.SizeOf("fcgi_header")) != b.SizeOf("fcgi_header") || hdr.GetVersion() < FCGI_VERSION_1 {
			return 0
		}
		len_ = hdr.GetContentLengthB1()<<8 | hdr.GetContentLengthB0()
		padding = hdr.GetPaddingLength()
		for hdr.GetType() == FCGI_PARAMS && len_ > 0 {
			if len_+padding > FCGI_MAX_LENGTH {
				return 0
			}
			if SafeRead(req, buf, len_+padding) != len_+padding {
				req.SetKeep(0)
				return 0
			}
			if FcgiGetParams(req, buf, buf+len_) == 0 {
				req.SetKeep(0)
				return 0
			}
			if SafeRead(req, &hdr, b.SizeOf("fcgi_header")) != b.SizeOf("fcgi_header") || hdr.GetVersion() < FCGI_VERSION_1 {
				req.SetKeep(0)
				return 0
			}
			len_ = hdr.GetContentLengthB1()<<8 | hdr.GetContentLengthB0()
			padding = hdr.GetPaddingLength()
		}
	} else if hdr.GetType() == FCGI_GET_VALUES {
		var p *uint8 = buf + b.SizeOf("fcgi_header")
		var value *types.Zval
		var zlen uint
		var q *FcgiHashBucket
		if SafeRead(req, buf, len_+padding) != len_+padding {
			req.SetKeep(0)
			return 0
		}
		if FcgiGetParams(req, buf, buf+len_) == 0 {
			req.SetKeep(0)
			return 0
		}
		q = req.GetEnv().GetList()
		for q != nil {
			if b.Assign(&value, FcgiMgmtVars.KeyFind(b.CastStr(q.GetVar(), q.GetVarLen()))) == nil {
				q = q.GetListNext()
				continue
			}
			zlen = uint(value.GetStr().GetLen())
			if p+4+4+q.GetVarLen()+zlen >= buf+b.SizeOf("buf") {
				break
			}
			if q.GetVarLen() < 0x80 {
				b.PostInc(&(*p)) = q.GetVarLen()
			} else {
				b.PostInc(&(*p)) = q.GetVarLen()>>24&0xff | 0x80
				b.PostInc(&(*p)) = q.GetVarLen() >> 16 & 0xff
				b.PostInc(&(*p)) = q.GetVarLen() >> 8 & 0xff
				b.PostInc(&(*p)) = q.GetVarLen() & 0xff
			}
			if zlen < 0x80 {
				b.PostInc(&(*p)) = zlen
			} else {
				b.PostInc(&(*p)) = zlen>>24&0xff | 0x80
				b.PostInc(&(*p)) = zlen >> 16 & 0xff
				b.PostInc(&(*p)) = zlen >> 8 & 0xff
				b.PostInc(&(*p)) = zlen & 0xff
			}
			memcpy(p, q.GetVar(), q.GetVarLen())
			p += q.GetVarLen()
			memcpy(p, value.GetStr().GetVal(), zlen)
			p += zlen
			q = q.GetListNext()
		}
		len_ = int(p - buf - b.SizeOf("fcgi_header"))
		len_ += FcgiMakeHeader((*FcgiHeader)(buf), FCGI_GET_VALUES_RESULT, 0, len_)
		if SafeWrite(req, buf, b.SizeOf("fcgi_header")+len_) != ssize_t(b.SizeOf("fcgi_header")+len_) {
			req.SetKeep(0)
			return 0
		}
		return 0
	} else {
		return 0
	}
	return 1
}
func FcgiRead(req *FcgiRequest, str *byte, len_ int) int {
	var ret int
	var n int
	var rest int
	var hdr FcgiHeader
	var buf []uint8
	n = 0
	rest = len_
	for rest > 0 {
		if req.GetInLen() == 0 {
			if SafeRead(req, &hdr, b.SizeOf("fcgi_header")) != b.SizeOf("fcgi_header") || hdr.GetVersion() < FCGI_VERSION_1 || hdr.GetType() != FCGI_STDIN {
				req.SetKeep(0)
				return 0
			}
			req.SetInLen(hdr.GetContentLengthB1()<<8 | hdr.GetContentLengthB0())
			req.SetInPad(hdr.GetPaddingLength())
			if req.GetInLen() == 0 {
				return n
			}
		}
		if req.GetInLen() >= rest {
			ret = int(SafeRead(req, str, rest))
		} else {
			ret = int(SafeRead(req, str, req.GetInLen()))
		}
		if ret < 0 {
			req.SetKeep(0)
			return ret
		} else if ret > 0 {
			req.SetInLen(req.GetInLen() - ret)
			rest -= ret
			n += ret
			str += ret
			if req.GetInLen() == 0 {
				if req.GetInPad() != 0 {
					if SafeRead(req, buf, req.GetInPad()) != req.GetInPad() {
						req.SetKeep(0)
						return ret
					}
				}
			} else {
				return n
			}
		} else {
			return n
		}
	}
	return n
}
func FcgiClose(req *FcgiRequest, force int, destroy int) {
	if destroy != 0 && req.GetHasEnv() != 0 {
		FcgiHashClean(req.GetEnv())
		req.SetHasEnv(0)
	}
	if (force != 0 || req.GetKeep() == 0) && req.GetFd() >= 0 {
		if force == 0 {
			var buf []byte
			shutdown(req.GetFd(), 1)

			/* read any remaining data, it may be omitted */

			for recv(req.GetFd(), buf, b.SizeOf("buf"), 0) > 0 {

			}

			/* read any remaining data, it may be omitted */

		}
		close(req.GetFd())
		req.SetFd(-1)
		req.GetHook().GetOnClose()()
	}
}
func FcgiIsClosed(req *FcgiRequest) int { return req.GetFd() < 0 }
func FcgiIsAllowed() int {
	var i int
	if ClientSa.GetSa().sa_family == AF_UNIX {
		return 1
	}
	if AllowedClients == nil {
		return 1
	}
	if ClientSa.GetSa().sa_family == AF_INET {
		for i = 0; AllowedClients[i].GetSa().sa_family; i++ {
			if AllowedClients[i].GetSa().sa_family == AF_INET && !(memcmp(ClientSa.GetSaInet().sin_addr, AllowedClients[i].GetSaInet().sin_addr, 4)) {
				return 1
			}
		}
	}
	if ClientSa.GetSa().sa_family == AF_INET6 {
		for i = 0; AllowedClients[i].GetSa().sa_family; i++ {
			if AllowedClients[i].GetSa().sa_family == AF_INET6 && !(memcmp(ClientSa.GetSaInet6().sin6_addr, AllowedClients[i].GetSaInet6().sin6_addr, 12)) {
				return 1
			}
		}
	}
	return 0
}
func FcgiAcceptRequest(req *FcgiRequest) int {
	for true {
		if req.GetFd() < 0 {
			for true {
				if InShutdown != 0 {
					return -1
				}
				req.GetHook().GetOnAccept()()
				var listen_socket int = req.GetListenSocket()
				var sa SaT
				var len_ socklen_t = b.SizeOf("sa")
				req.SetFd(accept(listen_socket, (*__struct__sockaddr)(&sa), &len_))
				ClientSa = sa
				if req.GetFd() >= 0 && FcgiIsAllowed() == 0 {
					FcgiLog(FCGI_ERROR, "Connection disallowed: IP address '%s' has been dropped.", FcgiGetLastClientIp())
					Closesocket(req.GetFd())
					req.SetFd(-1)
					continue
				}
				if req.GetFd() < 0 && (InShutdown != 0 || errno != EINTR && errno != ECONNABORTED) {
					return -1
				}
				if req.GetFd() >= 0 {
					var fds __struct__pollfd
					var ret int
					fds.fd = req.GetFd()
					fds.events = POLLIN
					fds.revents = 0
					for {
						errno = 0
						ret = poll(&fds, 1, 5000)
						if !(ret < 0 && errno == EINTR) {
							break
						}
					}
					if ret > 0 && (fds.revents&POLLIN) != 0 {
						break
					}
					FcgiClose(req, 1, 0)
				}
			}
		} else if InShutdown != 0 {
			return -1
		}
		req.GetHook().GetOnRead()()
		if FcgiReadRequest(req) != 0 {
			return req.GetFd()
		} else {
			FcgiClose(req, 1, 1)
		}
	}
}
func OpenPacket(req *FcgiRequest, type_ FcgiRequestType) *FcgiHeader {
	req.SetOutHdr((*FcgiHeader)(req.GetOutPos()))
	req.GetOutHdr().SetType(type_)
	req.SetOutPos(req.GetOutPos() + b.SizeOf("fcgi_header"))
	return req.GetOutHdr()
}
func ClosePacket(req *FcgiRequest) {
	if req.GetOutHdr() != nil {
		var len_ int = int(req.GetOutPos() - (*uint8)(req.GetOutHdr()+b.SizeOf("fcgi_header")))
		req.SetOutPos(req.GetOutPos() + FcgiMakeHeader(req.GetOutHdr(), FcgiRequestType(req.GetOutHdr().GetType()), req.GetId(), len_))
		req.SetOutHdr(nil)
	}
}
func FcgiFlush(req *FcgiRequest, end int) int {
	var len_ int
	ClosePacket(req)
	len_ = int(req.GetOutPos() - req.GetOutBuf())
	if end != 0 {
		var rec *FcgiEndRequestRec = (*FcgiEndRequestRec)(req.GetOutPos())
		FcgiMakeHeader(rec.GetHdr(), FCGI_END_REQUEST, req.GetId(), b.SizeOf("fcgi_end_request"))
		rec.GetBody().SetAppStatusB3(0)
		rec.GetBody().SetAppStatusB2(0)
		rec.GetBody().SetAppStatusB1(0)
		rec.GetBody().SetAppStatusB0(0)
		rec.GetBody().SetProtocolStatus(FCGI_REQUEST_COMPLETE)
		len_ += b.SizeOf("fcgi_end_request_rec")
	}
	if SafeWrite(req, req.GetOutBuf(), len_) != len_ {
		req.SetKeep(0)
		req.SetOutPos(req.GetOutBuf())
		return 0
	}
	req.SetOutPos(req.GetOutBuf())
	return 1
}
func FcgiWrite(req *FcgiRequest, type_ FcgiRequestType, str *byte, len_ int) int {
	var limit int
	var rest int
	if len_ <= 0 {
		return 0
	}
	if req.GetOutHdr() != nil && req.GetOutHdr().GetType() != type_ {
		ClosePacket(req)
	}

	/* Optimized version */

	limit = int(b.SizeOf("req -> out_buf") - (req.GetOutPos() - req.GetOutBuf()))
	if req.GetOutHdr() == nil {
		limit -= b.SizeOf("fcgi_header")
		if limit < 0 {
			limit = 0
		}
	}
	if len_ < limit {
		if req.GetOutHdr() == nil {
			OpenPacket(req, type_)
		}
		memcpy(req.GetOutPos(), str, len_)
		req.SetOutPos(req.GetOutPos() + len_)
	} else if len_-limit < int(b.SizeOf("req -> out_buf")-b.SizeOf("fcgi_header")) {
		if req.GetOutHdr() == nil {
			OpenPacket(req, type_)
		}
		if limit > 0 {
			memcpy(req.GetOutPos(), str, limit)
			req.SetOutPos(req.GetOutPos() + limit)
		}
		if FcgiFlush(req, 0) == 0 {
			return -1
		}
		if len_ > limit {
			OpenPacket(req, type_)
			memcpy(req.GetOutPos(), str+limit, len_-limit)
			req.SetOutPos(req.GetOutPos() + len_ - limit)
		}
	} else {
		var pos int = 0
		var pad int
		ClosePacket(req)
		for len_-pos > 0xffff {
			OpenPacket(req, type_)
			FcgiMakeHeader(req.GetOutHdr(), type_, req.GetId(), 0xfff8)
			req.SetOutHdr(nil)
			if FcgiFlush(req, 0) == 0 {
				return -1
			}
			if SafeWrite(req, str+pos, 0xfff8) != 0xfff8 {
				req.SetKeep(0)
				return -1
			}
			pos += 0xfff8
		}
		pad = (len_ - pos + 7 & ^7) - (len_ - pos)
		if pad != 0 {
			rest = 8 - pad
		} else {
			rest = 0
		}
		OpenPacket(req, type_)
		FcgiMakeHeader(req.GetOutHdr(), type_, req.GetId(), len_-pos-rest)
		req.SetOutHdr(nil)
		if FcgiFlush(req, 0) == 0 {
			return -1
		}
		if SafeWrite(req, str+pos, len_-pos-rest) != len_-pos-rest {
			req.SetKeep(0)
			return -1
		}
		if pad != 0 {
			OpenPacket(req, type_)
			memcpy(req.GetOutPos(), str+len_-rest, rest)
			req.SetOutPos(req.GetOutPos() + rest)
		}
	}
	return len_
}
func FcgiEnd(req *FcgiRequest) int {
	var ret int = 1
	if req.GetEnded() == 0 {
		ret = FcgiFlush(req, 1)
		req.SetEnded(1)
	}
	return ret
}
func FcgiFinishRequest(req *FcgiRequest, force_close int) int {
	var ret int = 1
	if req.GetFd() >= 0 {
		ret = FcgiEnd(req)
		FcgiClose(req, force_close, 1)
	}
	return ret
}
func FcgiHasEnv(req *FcgiRequest) int {
	return req != nil && req.GetHasEnv() != 0
}
func FcgiGetenv(req *FcgiRequest, var_ *byte, var_len int) *byte {
	var val_len uint
	if req == nil {
		return nil
	}
	return FcgiHashGet(req.GetEnv(), FCGI_HASH_FUNC(var_, var_len), (*byte)(var_), var_len, &val_len)
}
func FcgiQuickGetenv(req *FcgiRequest, var_ *byte, var_len int, hash_value uint) *byte {
	var val_len uint
	return FcgiHashGet(req.GetEnv(), hash_value, (*byte)(var_), var_len, &val_len)
}
func FcgiPutenv(req *FcgiRequest, var_ *byte, var_len int, val *byte) *byte {
	if req == nil {
		return nil
	}
	if val == nil {
		FcgiHashDel(req.GetEnv(), FCGI_HASH_FUNC(var_, var_len), var_, var_len)
		return nil
	} else {
		return FcgiHashSet(req.GetEnv(), FCGI_HASH_FUNC(var_, var_len), var_, var_len, val, uint(strlen(val)))
	}
}
func FcgiQuickPutenv(req *FcgiRequest, var_ *byte, var_len int, hash_value uint, val *byte) *byte {
	if val == nil {
		FcgiHashDel(req.GetEnv(), hash_value, var_, var_len)
		return nil
	} else {
		return FcgiHashSet(req.GetEnv(), hash_value, var_, var_len, val, uint(strlen(val)))
	}
}
func FcgiLoadenv(req *FcgiRequest, func_ FcgiApplyFunc, array *types.Zval) {
	FcgiHashApply(req.GetEnv(), func_, array)
}
func FcgiSetMgmtVar(name string, name_len int, value string, value_len int) {
	var zvalue types.Zval
	var key *types.String = types.NewString(b.CastStr(name, name_len))
	zvalue.SetString(types.NewString(b.CastStr(value, value_len)))
	types.GC_MAKE_PERSISTENT_LOCAL(key)
	types.GC_MAKE_PERSISTENT_LOCAL(zvalue.GetStr())
	FcgiMgmtVars.KeyAdd(key.GetStr(), &zvalue)
	types.ZendStringReleaseEx(key, 1)
}
func FcgiFreeMgmtVarCb(zv *types.Zval) { zend.Pefree(zv.GetStr(), 1) }
func FcgiGetLastClientIp() *byte {
	var str []byte

	/* Ipv4 */

	if ClientSa.GetSa().sa_family == AF_INET {
		return inet_ntop(ClientSa.GetSa().sa_family, ClientSa.GetSaInet().sin_addr, str, INET6_ADDRSTRLEN)
	}

	/* Ipv6 */

	if ClientSa.GetSa().sa_family == AF_INET6 {
		return inet_ntop(ClientSa.GetSa().sa_family, ClientSa.GetSaInet6().sin6_addr, str, INET6_ADDRSTRLEN)
	}

	/* Unix socket */

	return nil

	/* Unix socket */
}
