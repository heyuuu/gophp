package core

import (
	"sik/ext/spl"
	"sik/ext/standard"
	"sik/zend"
)

var PhpBuiltinExtensions []*zend.ZendModuleEntry = []*zend.ZendModuleEntry{
	//phpext_date_ptr,
	//phpext_libxml_ptr,
	//phpext_pcre_ptr,
	//phpext_sqlite3_ptr,
	//phpext_ctype_ptr,
	//phpext_dom_ptr,
	//phpext_fileinfo_ptr,
	//phpext_filter_ptr,
	//phpext_hash_ptr,
	//phpext_iconv_ptr,
	//phpext_json_ptr,
	spl.PhpextSplPtr,
	//phpext_pdo_ptr,
	//phpext_pdo_sqlite_ptr,
	//phpext_phar_ptr,
	//phpext_posix_ptr,
	//phpext_reflection_ptr,
	//phpext_session_ptr,
	//phpext_simplexml_ptr,
	standard.PhpextStandardPtr,
	//phpext_tokenizer_ptr,
	//phpext_xml_ptr,
	//phpext_xmlreader_ptr,
	//phpext_xmlwriter_ptr,
}

func PhpRegisterInternalExtensions() bool {
	return PhpRegisterExtensions(PhpBuiltinExtensions)
}
