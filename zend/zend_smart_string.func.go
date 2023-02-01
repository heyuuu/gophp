// <<generate>>

package zend

func SmartStringAlloc(str *SmartString, len_ int, persistent ZendBool) int { return str.alloc(len_) }
func SmartStringAppends(str *SmartString, src *byte)                       { str.AppendS(src) }
func SmartStringAppendl(str *SmartString, src *byte, len_ int)             { str.AppendL(src, len_) }
func SmartStringAppendc(str *SmartString, c byte)                          { str.AppendC(c) }
func SmartStringFree(s *SmartString)                                       { s.Free() }
func SmartString0(str *SmartString)                                        { str.ZeroTail() }
func SmartStringReset(str *SmartString)                                    { str.Reset() }
