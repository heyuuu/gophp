// <<generate>>

package zend

import (
	b "sik/builtin"
)

func DummyEncodingFetcher(encoding_name *byte) *ZendEncoding            { return nil }
func DummyEncodingNameGetter(encoding *ZendEncoding) *byte              { return (*byte)(encoding) }
func DummyEncodingLexerCompatibilityChecker(encoding *ZendEncoding) int { return 0 }
func DummyEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return nil
}
func DummyEncodingConverter(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int {
	return size_t - 1
}
func DummyEncodingListParser(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	*return_list = Pemalloc(0, persistent)
	*return_size = 0
	return SUCCESS
}
func DummyInternalEncodingGetter() *ZendEncoding             { return nil }
func DummyInternalEncodingSetter(encoding *ZendEncoding) int { return FAILURE }
func ZendMultibyteSetFunctions(functions *ZendMultibyteFunctions) int {
	ZendMultibyteEncodingUtf32be = functions.GetEncodingFetcher()("UTF-32BE")
	if ZendMultibyteEncodingUtf32be == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf32le = functions.GetEncodingFetcher()("UTF-32LE")
	if ZendMultibyteEncodingUtf32le == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf16be = functions.GetEncodingFetcher()("UTF-16BE")
	if ZendMultibyteEncodingUtf16be == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf16le = functions.GetEncodingFetcher()("UTF-16LE")
	if ZendMultibyteEncodingUtf16le == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf8 = functions.GetEncodingFetcher()("UTF-8")
	if ZendMultibyteEncodingUtf8 == nil {
		return FAILURE
	}
	MultibyteFunctionsDummy = MultibyteFunctions
	MultibyteFunctions = *functions

	/* As zend_multibyte_set_functions() gets called after ini settings were
	 * populated, we need to reinitialize script_encoding here.
	 */

	var value *byte = ZendIniString("zend.script_encoding", b.SizeOf("\"zend.script_encoding\"")-1, 0)
	ZendMultibyteSetScriptEncodingByString(value, strlen(value))
	return SUCCESS
}
func ZendMultibyteRestoreFunctions() {
	MultibyteFunctions = MultibyteFunctionsDummy
}
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
func ZendMultibyteCheckLexerCompatibility(encoding *ZendEncoding) int {
	return MultibyteFunctions.GetLexerCompatibilityChecker()(encoding)
}
func ZendMultibyteEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return MultibyteFunctions.GetEncodingDetector()(string, length, list, list_size)
}
func ZendMultibyteEncodingConverter(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int {
	return MultibyteFunctions.GetEncodingConverter()(to, to_length, from, from_length, encoding_to, encoding_from)
}
func ZendMultibyteParseEncodingList(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	return MultibyteFunctions.GetEncodingListParser()(encoding_list, encoding_list_len, return_list, return_size, persistent)
}
func ZendMultibyteGetInternalEncoding() *ZendEncoding {
	return MultibyteFunctions.GetInternalEncodingGetter()()
}
func ZendMultibyteGetScriptEncoding() *ZendEncoding { return __INI_SCNG().script_encoding }
func ZendMultibyteSetScriptEncoding(encoding_list **ZendEncoding, encoding_list_size int) int {
	if __CG().GetScriptEncodingList() != nil {
		Free((*byte)(__CG().GetScriptEncodingList()))
	}
	__CG().SetScriptEncodingList(encoding_list)
	__CG().SetScriptEncodingListSize(encoding_list_size)
	return SUCCESS
}
func ZendMultibyteSetInternalEncoding(encoding *ZendEncoding) int {
	return MultibyteFunctions.GetInternalEncodingSetter()(encoding)
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
