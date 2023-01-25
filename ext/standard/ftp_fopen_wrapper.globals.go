// <<generate>>

package standard

import (
	"sik/core"
)

const FTPS_ENCRYPT_DATA = 1

var PhpFtpDirstreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpFtpDirstreamRead, PhpFtpDirstreamClose, nil, "ftpdir", nil, nil, nil, nil}
var FtpStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapFtp, PhpStreamFtpStreamClose, PhpStreamFtpStreamStat, PhpStreamFtpUrlStat, PhpStreamFtpOpendir, "ftp", PhpStreamFtpUnlink, PhpStreamFtpRename, PhpStreamFtpMkdir, PhpStreamFtpRmdir, nil}
var PhpStreamFtpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&FtpStreamWops, nil, 1}
