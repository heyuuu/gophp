package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

func PhpInitCryptR()     {}
func PhpShutdownCryptR() {}
func _cryptExtendedInitR() {
	var initialized sig_atomic_t = 0
	if !initialized {
		__sync_fetch_and_add(&initialized, 1)
		_cryptExtendedInit()
	}
}
func To64(s *byte, v int32, n int) {
	for b.PreDec(&n) >= 0 {
		b.PostInc(&(*s)) = Itoa64[v&0x3f]
		v >>= 6
	}
}
func PhpMd5CryptR(pw *byte, salt *byte, out *byte) *byte {
	var passwd []byte
	var p *byte
	var sp *byte
	var ep *byte
	var final []uint8
	var i uint
	var sl uint
	var pwl uint
	var ctx PHP_MD5_CTX
	var ctx1 PHP_MD5_CTX
	var l uint32
	var pl int
	pwl = strlen(pw)

	/* Refine the salt first */

	sp = salt

	/* If it starts with the magic string, then skip that */

	if strncmp(sp, "$1$", 3) == 0 {
		sp += 3
	}

	/* It stops at the first '$', max 8 chars */

	for ep = sp; (*ep) != '0' && (*ep) != '$' && ep < sp+8; ep++ {

	}

	/* get the length of the true salt */

	sl = ep - sp
	PHP_MD5Init(&ctx)

	/* The password first, since that is what is most unknown */

	PHP_MD5Update(&ctx, (*uint8)(pw), pwl)

	/* Then our magic string */

	PHP_MD5Update(&ctx, (*uint8)("$1$"), 3)

	/* Then the raw salt */

	PHP_MD5Update(&ctx, (*uint8)(sp), sl)

	/* Then just as many characters of the MD5(pw,salt,pw) */

	PHP_MD5Init(&ctx1)
	PHP_MD5Update(&ctx1, (*uint8)(pw), pwl)
	PHP_MD5Update(&ctx1, (*uint8)(sp), sl)
	PHP_MD5Update(&ctx1, (*uint8)(pw), pwl)
	PHP_MD5Final(final, &ctx1)
	for pl = pwl; pl > 0; pl -= 16 {
		PHP_MD5Update(&ctx, final, uint(b.Cond(pl > 16, 16, pl)))
	}

	/* Don't leave anything around in vm they could use. */

	zend.ZEND_SECURE_ZERO(final, b.SizeOf("final"))

	/* Then something really weird... */

	for i = pwl; i != 0; i >>= 1 {
		if (i & 1) != 0 {
			PHP_MD5Update(&ctx, final, 1)
		} else {
			PHP_MD5Update(&ctx, (*uint8)(pw), 1)
		}
	}

	/* Now make the output string */

	memcpy(passwd, "$1$", 3)
	strlcpy(passwd+3, sp, sl+1)
	strcat(passwd, "$")
	PHP_MD5Final(final, &ctx)

	/*
	 * And now, just to make sure things don't run too fast. On a 60 MHz
	 * Pentium this takes 34 msec, so you would need 30 seconds to build
	 * a 1000 entry dictionary...
	 */

	for i = 0; i < 1000; i++ {
		PHP_MD5Init(&ctx1)
		if (i & 1) != 0 {
			PHP_MD5Update(&ctx1, (*uint8)(pw), pwl)
		} else {
			PHP_MD5Update(&ctx1, final, 16)
		}
		if i%3 != 0 {
			PHP_MD5Update(&ctx1, (*uint8)(sp), sl)
		}
		if i%7 != 0 {
			PHP_MD5Update(&ctx1, (*uint8)(pw), pwl)
		}
		if (i & 1) != 0 {
			PHP_MD5Update(&ctx1, final, 16)
		} else {
			PHP_MD5Update(&ctx1, (*uint8)(pw), pwl)
		}
		PHP_MD5Final(final, &ctx1)
	}
	p = passwd + sl + 3 + 1
	l = final[0]<<16 | final[6]<<8 | final[12]
	To64(p, l, 4)
	p += 4
	l = final[1]<<16 | final[7]<<8 | final[13]
	To64(p, l, 4)
	p += 4
	l = final[2]<<16 | final[8]<<8 | final[14]
	To64(p, l, 4)
	p += 4
	l = final[3]<<16 | final[9]<<8 | final[15]
	To64(p, l, 4)
	p += 4
	l = final[4]<<16 | final[10]<<8 | final[5]
	To64(p, l, 4)
	p += 4
	l = final[11]
	To64(p, l, 2)
	p += 2
	*p = '0'

	/* Don't leave anything around in vm they could use. */

	zend.ZEND_SECURE_ZERO(final, b.SizeOf("final"))
	return passwd
}
