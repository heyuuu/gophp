// <<generate>>

package zend

func DummyEncodingFetcher(encoding_name *byte) *ZendEncoding            { return nil }
func DummyEncodingNameGetter(encoding *ZendEncoding) *byte              { return (*byte)(encoding) }
func DummyEncodingLexerCompatibilityChecker(encoding *ZendEncoding) int { return 0 }
func DummyEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return nil
}
func DummyEncodingConverter(
	to **uint8,
	to_length *int,
	from *uint8,
	from_length int,
	encoding_to *ZendEncoding,
	encoding_from *ZendEncoding,
) int {
	return -1
}
func DummyEncodingListParser(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	*return_list = Pemalloc(0, persistent)
	*return_size = 0
	return SUCCESS
}
func DummyInternalEncodingGetter() *ZendEncoding             { return nil }
func DummyInternalEncodingSetter(encoding *ZendEncoding) int { return FAILURE }
func ZendMultibyteGetFunctions() *ZendMultibyteFunctions {
	if MultibyteFunctions.GetProviderName() != nil {
		return &MultibyteFunctions
	} else {
		return nil
	}
}
func ZendMultibyteFetchEncoding(name *byte) *ZendEncoding {
	return MultibyteFunctions.GetEncodingFetcher()(name)
}
func ZendMultibyteGetEncodingName(encoding *ZendEncoding) *byte {
	return MultibyteFunctions.GetEncodingNameGetter()(encoding)
}
func ZendMultibyteEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return MultibyteFunctions.GetEncodingDetector()(string, length, list, list_size)
}
func ZendMultibyteEncodingConverter(
	to **uint8,
	to_length *int,
	from *uint8,
	from_length int,
	encoding_to *ZendEncoding,
	encoding_from *ZendEncoding,
) int {
	return MultibyteFunctions.GetEncodingConverter()(to, to_length, from, from_length, encoding_to, encoding_from)
}
func ZendMultibyteParseEncodingList(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	return MultibyteFunctions.GetEncodingListParser()(encoding_list, encoding_list_len, return_list, return_size, persistent)
}
func ZendMultibyteGetInternalEncoding() *ZendEncoding {
	return MultibyteFunctions.GetInternalEncodingGetter()()
}
func ZendMultibyteSetScriptEncoding(encoding_list **ZendEncoding, encoding_list_size int) int {
	if CG__().GetScriptEncodingList() != nil {
		Free((*byte)(CG__().GetScriptEncodingList()))
	}
	CG__().SetScriptEncodingList(encoding_list)
	CG__().SetScriptEncodingListSize(encoding_list_size)
	return SUCCESS
}
func ZendMultibyteSetScriptEncodingByString(new_value *byte, new_value_length int) int {
	var list **ZendEncoding = 0
	var size int = 0
	if new_value == nil {
		ZendMultibyteSetScriptEncoding(nil, 0)
		return SUCCESS
	}
	if FAILURE == ZendMultibyteParseEncodingList(new_value, new_value_length, &list, &size, 1) {
		return FAILURE
	}
	if size == 0 {
		Pefree(any(list), 1)
		return FAILURE
	}
	if FAILURE == ZendMultibyteSetScriptEncoding(list, size) {
		return FAILURE
	}
	return SUCCESS
}
