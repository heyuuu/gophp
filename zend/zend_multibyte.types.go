// <<generate>>

package zend

/**
 * ZendMultibyteFunctions
 */
type ZendMultibyteFunctions struct {
	provider_name               *byte
	encoding_fetcher            ZendEncodingFetcher
	encoding_name_getter        ZendEncodingNameGetter
	lexer_compatibility_checker ZendEncodingLexerCompatibilityChecker
	encoding_detector           ZendEncodingDetector
	encoding_converter          ZendEncodingConverter
	encoding_list_parser        ZendEncodingListParser
	internal_encoding_getter    ZendEncodingInternalEncodingGetter
	internal_encoding_setter    ZendEncodingInternalEncodingSetter
}

func MakeZendMultibyteFunctions(
	provider_name *byte,
	encoding_fetcher ZendEncodingFetcher,
	encoding_name_getter ZendEncodingNameGetter,
	lexer_compatibility_checker ZendEncodingLexerCompatibilityChecker,
	encoding_detector ZendEncodingDetector,
	encoding_converter ZendEncodingConverter,
	encoding_list_parser ZendEncodingListParser,
	internal_encoding_getter ZendEncodingInternalEncodingGetter,
	internal_encoding_setter ZendEncodingInternalEncodingSetter,
) ZendMultibyteFunctions {
	return ZendMultibyteFunctions{
		provider_name:               provider_name,
		encoding_fetcher:            encoding_fetcher,
		encoding_name_getter:        encoding_name_getter,
		lexer_compatibility_checker: lexer_compatibility_checker,
		encoding_detector:           encoding_detector,
		encoding_converter:          encoding_converter,
		encoding_list_parser:        encoding_list_parser,
		internal_encoding_getter:    internal_encoding_getter,
		internal_encoding_setter:    internal_encoding_setter,
	}
}
func (this *ZendMultibyteFunctions) GetProviderName() *byte { return this.provider_name }

// func (this *ZendMultibyteFunctions) SetProviderName(value *byte) { this.provider_name = value }
func (this *ZendMultibyteFunctions) GetEncodingFetcher() ZendEncodingFetcher {
	return this.encoding_fetcher
}

// func (this *ZendMultibyteFunctions) SetEncodingFetcher(value ZendEncodingFetcher) { this.encoding_fetcher = value }
func (this *ZendMultibyteFunctions) GetEncodingNameGetter() ZendEncodingNameGetter {
	return this.encoding_name_getter
}

// func (this *ZendMultibyteFunctions) SetEncodingNameGetter(value ZendEncodingNameGetter) { this.encoding_name_getter = value }
func (this *ZendMultibyteFunctions) GetLexerCompatibilityChecker() ZendEncodingLexerCompatibilityChecker {
	return this.lexer_compatibility_checker
}

// func (this *ZendMultibyteFunctions) SetLexerCompatibilityChecker(value ZendEncodingLexerCompatibilityChecker) { this.lexer_compatibility_checker = value }
func (this *ZendMultibyteFunctions) GetEncodingDetector() ZendEncodingDetector {
	return this.encoding_detector
}

// func (this *ZendMultibyteFunctions) SetEncodingDetector(value ZendEncodingDetector) { this.encoding_detector = value }
func (this *ZendMultibyteFunctions) GetEncodingConverter() ZendEncodingConverter {
	return this.encoding_converter
}

// func (this *ZendMultibyteFunctions) SetEncodingConverter(value ZendEncodingConverter) { this.encoding_converter = value }
func (this *ZendMultibyteFunctions) GetEncodingListParser() ZendEncodingListParser {
	return this.encoding_list_parser
}

// func (this *ZendMultibyteFunctions) SetEncodingListParser(value ZendEncodingListParser) { this.encoding_list_parser = value }
func (this *ZendMultibyteFunctions) GetInternalEncodingGetter() ZendEncodingInternalEncodingGetter {
	return this.internal_encoding_getter
}

// func (this *ZendMultibyteFunctions) SetInternalEncodingGetter(value ZendEncodingInternalEncodingGetter) { this.internal_encoding_getter = value }
func (this *ZendMultibyteFunctions) GetInternalEncodingSetter() ZendEncodingInternalEncodingSetter {
	return this.internal_encoding_setter
}

// func (this *ZendMultibyteFunctions) SetInternalEncodingSetter(value ZendEncodingInternalEncodingSetter) { this.internal_encoding_setter = value }
