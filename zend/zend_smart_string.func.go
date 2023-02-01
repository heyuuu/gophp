// <<generate>>

package zend

func SmartStringAlloc(str *SmartString, len_ int, persistent ZendBool) int { return str.Alloc(len_) }
func SmartStringAppends(str *SmartString, src string)                      { str.AppendString(src) }
func SmartStringAppendl(str *SmartString, src string)                      { str.AppendString(src) }
func SmartStringAppendc(str *SmartString, c byte)                          { str.AppendByte(c) }
func SmartStringFree(s *SmartString)                                       { s.Free() }
func SmartString0(str *SmartString)                                        { str.ZeroTail() }
func SmartStringReset(str *SmartString)                                    { str.Reset() }
