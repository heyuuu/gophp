package standard

func CHARSET_UNICODE_COMPAT(cs EntityCharset) bool                      { return cs <= Cs88591 }
func CHARSET_SINGLE_BYTE(cs EntityCharset) bool                         { return cs > CsUtf8 && cs < CsBig5 }
func CHARSET_PARTIAL_SUPPORT(cs EntityCharset) bool                     { return cs >= CsBig5 }
func ENT_ENC_TO_UNI_STAGE1(k unsigned) int                              { return (k & 0xc0) >> 6 }
func ENT_ENC_TO_UNI_STAGE2(k unsigned) int                              { return k & 0x3f }
func ENT_STAGE1_INDEX(k __auto__) int                                   { return (k & 0xfff000) >> 12 }
func ENT_STAGE2_INDEX(k __auto__) int                                   { return (k & 0xfc0) >> 6 }
func ENT_STAGE3_INDEX(k __auto__) int                                   { return k & 0x3f }
func ENT_CODE_POINT_FROM_STAGES(i unsigned, j unsigned, k unsigned) int { return i<<12 | j<<6 | k }
