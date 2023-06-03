package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"strconv"
	"strings"
)

const smartStrStartSize = 256

// VK_ESCAPE, Ascii(27 | 0x1B)，在 PHP 中的转义符为 '\e'
const VkEscape = '\x1b'

/**
 * SmartStr
 */
type SmartString = SmartStr
type SmartStr struct {
	buffer strings.Builder
}

func (s *SmartStr) Write(p []byte) (n int, err error) {
	return s.buffer.Write(p)
}

func (s *SmartStr) GetLen() int    { return s.buffer.Len() }
func (s *SmartStr) GetStr() string { return s.buffer.String() }

// 分配内存，确认至少有 len_ 的未使用内存
func (s *SmartStr) Alloc(len_ int) int {
	if s.buffer.Cap() == 0 && len_ < smartStrStartSize {
		// 初始化时，最小尺寸为 SMART_STR_START_SIZE，避免小尺寸重复扩展
		s.buffer.Grow(smartStrStartSize)
	} else {
		s.buffer.Grow(len_)
	}
	return s.buffer.Len() + len_
}

func (s *SmartStr) WriteString(str string) {
	s.buffer.WriteString(str)
}

func (s *SmartStr) WriteByte(c byte) {
	s.buffer.WriteByte(c)
}

func (s *SmartStr) AppendSmartStr(str *SmartStr) {
	s.buffer.WriteString(str.GetStr())
}

func (s *SmartStr) SetString(str string) {
	s.Reset()
	s.buffer.WriteString(str)
}

func (s *SmartStr) Reset() {
	s.buffer.Reset()
}

func (s *SmartStr) Free() {
	s.Reset()
}

func (s *SmartStr) ZeroTail() {
	// c 字符串尾部设置0
}

/**
 * 快捷方法
 */
func (s *SmartStr) AppendLong(num ZendLong) {
	var str = strconv.FormatInt(int64(num), 10)
	s.WriteString(str)
}

func (s *SmartStr) AppendUlong(num ZendUlong) {
	var str = strconv.FormatUint(uint64(num), 10)
	s.WriteString(str)
}

func (s *SmartStr) AppendEscaped(str string) {
	for _, c := range []byte(str) {
		if c < 32 || c == '\\' || c > 126 {
			s.WriteByte('\\')
			switch c {
			case '\n':
				s.WriteByte('n')
			case '\r':
				s.WriteByte('r')
			case '\t':
				s.WriteByte('t')
			case '\f':
				s.WriteByte('f')
			case '\v':
				s.WriteByte('v')
			case '\\':
				s.WriteByte('\\')
			case VkEscape:
				s.WriteByte('e')
			default:
				s.WriteByte('x')
				if c>>4 < 10 {
					s.WriteByte(c>>4 + '0')
				} else {
					s.WriteByte(c>>4 + 'A' - 10)
				}
				if (c & 0xf) < 10 {
					s.WriteByte(c&0xf + '0')
				} else {
					s.WriteByte(c&0xf + 'A' - 10)
				}
			}
		} else {
			s.WriteByte(c)
		}
	}
}

/**
 * todo 待移除方法
 */

func (s *SmartStr) GetS() *types.String {
	// todo 需要确认是否兼容 ZendStringAlloc() 但未使用时的空 []byte
	return types.NewString(s.GetStr())
}
func (s *SmartStr) GetC() *byte {
	// todo 仅占位，实际使用需替换
	var str = s.GetStr()
	var char = str[0]
	return &char
}
