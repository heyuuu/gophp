package standard

import (
	"sik/core"
)

var PhpFtpDirstreamOps = core.MakePhpStreamOps(nil, PhpFtpDirstreamRead, PhpFtpDirstreamClose, nil, "ftpdir", nil, nil, nil, nil)
var FtpStreamWops = core.MakePhpStreamWrapperOps(PhpStreamUrlWrapFtp, PhpStreamFtpStreamClose, PhpStreamFtpStreamStat, PhpStreamFtpUrlStat, PhpStreamFtpOpendir, "ftp", PhpStreamFtpUnlink, PhpStreamFtpRename, PhpStreamFtpMkdir, PhpStreamFtpRmdir, nil)
var PhpStreamFtpWrapper = core.MakePhpStreamWrapper(&FtpStreamWops, nil, 1)
