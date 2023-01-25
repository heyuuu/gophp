// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/ext/spl"
	"sik/ext/standard"
	"sik/zend"
)

// Source: <main/internal_functions.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/date/php_date.h"

// failed # include "ext/libxml/php_libxml.h"

// failed # include "ext/pcre/php_pcre.h"

// failed # include "ext/sqlite3/php_sqlite3.h"

// failed # include "ext/ctype/php_ctype.h"

// failed # include "ext/dom/php_dom.h"

// failed # include "ext/fileinfo/php_fileinfo.h"

// failed # include "ext/filter/php_filter.h"

// failed # include "ext/hash/php_hash.h"

// failed # include "ext/iconv/php_iconv.h"

// failed # include "ext/json/php_json.h"

// failed # include "ext/pdo/php_pdo.h"

// failed # include "ext/pdo_sqlite/php_pdo_sqlite.h"

// failed # include "ext/phar/php_phar.h"

// failed # include "ext/posix/php_posix.h"

// failed # include "ext/reflection/php_reflection.h"

// failed # include "ext/session/php_session.h"

// failed # include "ext/simplexml/php_simplexml.h"

// failed # include "ext/tokenizer/php_tokenizer.h"

// failed # include "ext/xml/php_xml.h"

// failed # include "ext/xmlreader/php_xmlreader.h"

// failed # include "ext/xmlwriter/php_xmlwriter.h"

var PhpBuiltinExtensions []*zend.ZendModuleEntry = []*zend.ZendModuleEntry{phpext_date_ptr, phpext_libxml_ptr, phpext_pcre_ptr, phpext_sqlite3_ptr, phpext_ctype_ptr, phpext_dom_ptr, phpext_fileinfo_ptr, phpext_filter_ptr, phpext_hash_ptr, phpext_iconv_ptr, phpext_json_ptr, spl.PhpextSplPtr, phpext_pdo_ptr, phpext_pdo_sqlite_ptr, phpext_phar_ptr, phpext_posix_ptr, phpext_reflection_ptr, phpext_session_ptr, phpext_simplexml_ptr, standard.PhpextStandardPtr, phpext_tokenizer_ptr, phpext_xml_ptr, phpext_xmlreader_ptr, phpext_xmlwriter_ptr}

const EXTCOUNT = b.SizeOf("php_builtin_extensions") / b.SizeOf("zend_module_entry *")
