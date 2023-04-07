package standard

import (
	"bytes"
	"crypto/md5"
	"strings"
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
func To64(v uint32, n int) []byte {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = Itoa64[v&0x3f]
		v >>= 6
	}
	return buf
}
func PhpMd5CryptR(password string, salt string) string {
	/* If salt starts with the magic string, then skip that */
	if strings.HasPrefix(salt, "$1$") {
		salt = salt[3:]
	}

	/* It stops at the first '$', max 8 chars */
	if pos := strings.IndexByte(salt, '$'); pos >= 0 {
		salt = salt[:pos]
	}
	if len(salt) > 8 {
		salt = salt[:8]
	}

	/* get the length of the true salt */
	var buf bytes.Buffer

	/* The password first, since that is what is most unknown */
	buf.WriteString(password)

	/* Then our magic string */
	buf.WriteString("$1$")

	/* Then the raw salt */
	buf.WriteString(salt)

	/* Then just as many characters of the MD5(pw,salt,pw) */
	final1 := md5.Sum([]byte(password + salt + password))
	for i := len(password) - 1; i >= 0; i -= 16 {
		if i > 16 {
			buf.Write(final1[:])
		} else {
			buf.Write(final1[:i])
		}
	}

	/* Then something really weird... */
	for i := len(password); i != 0; i >>= 1 {
		if (i & 1) != 0 {
			buf.WriteByte(final1[0])
		} else {
			buf.WriteByte(password[0])
		}
	}

	/* Now make the output string */
	final := md5.Sum(buf.Bytes())

	/*
	 * And now, just to make sure things don't run too fast. On a 60 MHz
	 * Pentium this takes 34 msec, so you would need 30 seconds to build
	 * a 1000 entry dictionary...
	 */
	for i := 0; i < 1000; i++ {
		var buf bytes.Buffer
		if (i & 1) != 0 {
			buf.WriteString(password)
		} else {
			buf.Write(final[:])
		}
		if i%3 != 0 {
			buf.WriteString(salt)
		}
		if i%7 != 0 {
			buf.WriteString(password)
		}
		if (i & 1) != 0 {
			buf.Write(final[:])
		} else {
			buf.WriteString(password)
		}
		final = md5.Sum(buf.Bytes())
	}

	//return "$1$" + salt + "$" + base64.NewEncoding(Itoa64).EncodeToString(final[:])
	var result strings.Builder
	var l uint32
	result.WriteString("$1$" + salt + "$")
	l = uint32(final[0])<<16 | uint32(final[6])<<8 | uint32(final[12])
	result.Write(To64(l, 4))
	l = uint32(final[1])<<16 | uint32(final[7])<<8 | uint32(final[13])
	result.Write(To64(l, 4))
	l = uint32(final[2])<<16 | uint32(final[8])<<8 | uint32(final[14])
	result.Write(To64(l, 4))
	l = uint32(final[3])<<16 | uint32(final[9])<<8 | uint32(final[15])
	result.Write(To64(l, 4))
	l = uint32(final[4])<<16 | uint32(final[10])<<8 | uint32(final[5])
	result.Write(To64(l, 4))
	l = uint32(final[11])
	result.Write(To64(l, 42))

	/* Don't leave anything around in vm they could use. */
	return result.String()
}
