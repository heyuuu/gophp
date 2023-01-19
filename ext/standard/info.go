// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/sapi/cli"
	"sik/zend"
)

// Source: <ext/standard/info.h>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Colin Viebrock <colin@viebrock.ca>                          |
   +----------------------------------------------------------------------+
*/

// #define INFO_H

// #define PHP_ENTRY_NAME_COLOR       "#ccf"

// #define PHP_CONTENTS_COLOR       "#ccc"

// #define PHP_HEADER_COLOR       "#99c"

// #define PHP_INFO_GENERAL       ( 1 << 0 )

// #define PHP_INFO_CREDITS       ( 1 << 1 )

// #define PHP_INFO_CONFIGURATION       ( 1 << 2 )

// #define PHP_INFO_MODULES       ( 1 << 3 )

// #define PHP_INFO_ENVIRONMENT       ( 1 << 4 )

// #define PHP_INFO_VARIABLES       ( 1 << 5 )

// #define PHP_INFO_LICENSE       ( 1 << 6 )

// #define PHP_INFO_ALL       0xFFFFFFFF

// #define HAVE_CREDITS_DEFS

// #define PHP_CREDITS_GROUP       ( 1 << 0 )

// #define PHP_CREDITS_GENERAL       ( 1 << 1 )

// #define PHP_CREDITS_SAPI       ( 1 << 2 )

// #define PHP_CREDITS_MODULES       ( 1 << 3 )

// #define PHP_CREDITS_DOCS       ( 1 << 4 )

// #define PHP_CREDITS_FULLPAGE       ( 1 << 5 )

// #define PHP_CREDITS_QA       ( 1 << 6 )

// #define PHP_CREDITS_WEB       ( 1 << 7 )

// #define PHP_CREDITS_ALL       0xFFFFFFFF

// #define PHP_LOGO_DATA_URI       "data:image/png;base64,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"

// #define PHP_EGG_LOGO_DATA_URI       "data:image/png;base64,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"

// #define ZEND_LOGO_DATA_URI       "data:image/png;base64,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"

var PhpInfoHtmlEscWrite func(string *byte, str_len int)
var PhpPrintStyle func()

// Source: <ext/standard/info.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Colin Viebrock <colin@viebrock.ca>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "php_globals.h"

// # include "ext/standard/head.h"

// # include "ext/standard/html.h"

// # include "info.h"

// # include "credits.h"

// # include "css.h"

// # include "SAPI.h"

// # include < time . h >

// # include "php_main.h"

// # include "zend_globals.h"

// # include "zend_extensions.h"

// # include "zend_highlight.h"

// # include < sys / utsname . h >

// # include "url.h"

// # include "php_string.h"

// #define SECTION(name) if ( ! sapi_module . phpinfo_as_text ) { php_info_print ( "<h2>" name "</h2>\n" ) ; } else { php_info_print_table_start ( ) ; php_info_print_table_header ( 1 , name ) ; php_info_print_table_end ( ) ; }

var PhpIniOpenedPath *byte
var PhpIniScannedPath *byte
var PhpIniScannedFiles *byte

func PhpInfoPrintHtmlEsc(str *byte, len_ int) int {
	var written int
	var new_str *zend.ZendString
	new_str = PhpEscapeHtmlEntities((*uint8)(str), len_, 0, 2|1, "utf-8")
	written = core.PhpOutputWrite(new_str.val, new_str.len_)
	zend.ZendStringFree(new_str)
	return written
}

/* }}} */

func PhpInfoPrintf(fmt string, _ ...any) int {
	var buf *byte
	var len_ int
	var written int
	var argv va_list
	va_start(argv, fmt)
	len_ = zend.ZendVspprintf(&buf, 0, fmt, argv)
	va_end(argv)
	written = core.PhpOutputWrite(buf, len_)
	zend._efree(buf)
	return written
}

/* }}} */

func PhpInfoPrint(str *byte) int {
	return core.PhpOutputWrite(str, strlen(str))
}

/* }}} */

func PhpInfoPrintStreamHash(name string, ht *zend.HashTable) {
	var key *zend.ZendString
	if ht != nil {
		if ht.nNumOfElements != 0 {
			var first int = 1
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrintf("<tr><td class=\"e\">Registered %s</td><td class=\"v\">", name)
			} else {
				PhpInfoPrintf("\nRegistered %s => ", name)
			}
			for {
				var __ht *zend.HashTable = ht
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					key = _p.key
					if key != nil {
						if first != 0 {
							first = 0
						} else {
							PhpInfoPrint(", ")
						}
						if core.sapi_module.phpinfo_as_text == 0 {
							PhpInfoPrintHtmlEsc(key.val, key.len_)
						} else {
							PhpInfoPrint(key.val)
						}
					}
				}
				break
			}
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrint("</td></tr>\n")
			}
		} else {
			var reg_name []byte
			core.ApPhpSnprintf(reg_name, g.SizeOf("reg_name"), "Registered %s", name)
			PhpInfoPrintTableRow(2, reg_name, "none registered")
		}
	} else {
		PhpInfoPrintTableRow(2, name, "disabled")
	}
}

/* }}} */

func PhpInfoPrintModule(zend_module *zend.ZendModuleEntry) {
	if zend_module.info_func != nil || zend_module.version != nil {
		if core.sapi_module.phpinfo_as_text == 0 {
			var url_name *zend.ZendString = PhpUrlEncode(zend_module.name, strlen(zend_module.name))
			PhpStrtolower(url_name.val, url_name.len_)
			PhpInfoPrintf("<h2><a name=\"module_%s\">%s</a></h2>\n", url_name.val, zend_module.name)
			zend._efree(url_name)
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableHeader(1, zend_module.name)
			PhpInfoPrintTableEnd()
		}
		if zend_module.info_func != nil {
			zend_module.info_func(zend_module)
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableRow(2, "Version", zend_module.version)
			PhpInfoPrintTableEnd()
			core.DisplayIniEntries(zend_module)
		}
	} else {
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrintf("<tr><td class=\"v\">%s</td></tr>\n", zend_module.name)
		} else {
			PhpInfoPrintf("%s\n", zend_module.name)
		}
	}
}

/* }}} */

func PhpPrintGpcseArray(name string, name_length uint32) {
	var data *zend.Zval
	var tmp *zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var key *zend.ZendString
	key = zend.ZendStringInit(name, name_length, 0)
	zend.ZendIsAutoGlobal(key)
	if g.Assign(&data, zend.ZendHashFindDeref(&zend.EG.symbol_table, key)) != nil && data.u1.v.type_ == 7 {
		for {
			var __ht *zend.HashTable = data.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				tmp = _z
				if core.sapi_module.phpinfo_as_text == 0 {
					PhpInfoPrint("<tr>")
					PhpInfoPrint("<td class=\"e\">")
				}
				PhpInfoPrint("$")
				PhpInfoPrint(name)
				PhpInfoPrint("['")
				if string_key != nil {
					if core.sapi_module.phpinfo_as_text == 0 {
						PhpInfoPrintHtmlEsc(string_key.val, string_key.len_)
					} else {
						PhpInfoPrint(string_key.val)
					}
				} else {
					PhpInfoPrintf("%"+"llu", num_key)
				}
				PhpInfoPrint("']")
				if core.sapi_module.phpinfo_as_text == 0 {
					PhpInfoPrint("</td><td class=\"v\">")
				} else {
					PhpInfoPrint(" => ")
				}
				if tmp.u1.v.type_ == 10 {
					tmp = &(*tmp).value.ref.val
				}
				if tmp.u1.v.type_ == 7 {
					if core.sapi_module.phpinfo_as_text == 0 {
						var str *zend.ZendString = zend.ZendPrintZvalRToStr(tmp, 0)
						PhpInfoPrint("<pre>")
						PhpInfoPrintHtmlEsc(str.val, str.len_)
						PhpInfoPrint("</pre>")
						zend.ZendStringReleaseEx(str, 0)
					} else {
						zend.ZendPrintZvalR(tmp, 0)
					}
				} else {
					var tmp2 *zend.ZendString
					var str *zend.ZendString = zend.ZvalGetTmpString(tmp, &tmp2)
					if core.sapi_module.phpinfo_as_text == 0 {
						if str.len_ == 0 {
							PhpInfoPrint("<i>no value</i>")
						} else {
							PhpInfoPrintHtmlEsc(str.val, str.len_)
						}
					} else {
						PhpInfoPrint(str.val)
					}
					zend.ZendTmpStringRelease(tmp2)
				}
				if core.sapi_module.phpinfo_as_text == 0 {
					PhpInfoPrint("</td></tr>\n")
				} else {
					PhpInfoPrint("\n")
				}
			}
			break
		}
	}
	zend.ZendStringEfree(key)
}

/* }}} */

func PhpInfoPrintStyle() {
	PhpInfoPrintf("<style type=\"text/css\">\n")
	PhpInfoPrintCss()
	PhpInfoPrintf("</style>\n")
}

/* }}} */

func PhpInfoHtmlEsc(string *byte) *zend.ZendString {
	return PhpEscapeHtmlEntities((*uint8)(string), strlen(string), 0, 2|1, nil)
}

/* }}} */

/* {{{ php_get_uname
 */

func PhpGetUname(mode byte) *zend.ZendString {
	var php_uname *byte
	var tmp_uname []byte
	var buf __struct__utsname
	if uname((*__struct__utsname)(&buf)) == -1 {
		php_uname = "Darwin bogon 21.6.0 Darwin Kernel Version 21.6.0: Wed Aug 10 14:28:23 PDT 2022; root:xnu-8020.141.5~2/RELEASE_ARM64_T6000 arm64"
	} else {
		if mode == 's' {
			php_uname = buf.sysname
		} else if mode == 'r' {
			php_uname = buf.release
		} else if mode == 'n' {
			php_uname = buf.nodename
		} else if mode == 'v' {
			php_uname = buf.version
		} else if mode == 'm' {
			php_uname = buf.machine
		} else {
			core.ApPhpSnprintf(tmp_uname, g.SizeOf("tmp_uname"), "%s %s %s %s %s", buf.sysname, buf.nodename, buf.release, buf.version, buf.machine)
			php_uname = tmp_uname
		}
	}
	return zend.ZendStringInit(php_uname, strlen(php_uname), 0)
}

/* }}} */

func PhpPrintInfoHtmlhead() {
	PhpInfoPrint("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"DTD/xhtml1-transitional.dtd\">\n")
	PhpInfoPrint("<html xmlns=\"http://www.w3.org/1999/xhtml\">")
	PhpInfoPrint("<head>\n")
	PhpInfoPrintStyle()
	PhpInfoPrintf("<title>PHP %s - phpinfo()</title>", "7.4.33")
	PhpInfoPrint("<meta name=\"ROBOTS\" content=\"NOINDEX,NOFOLLOW,NOARCHIVE\" />")
	PhpInfoPrint("</head>\n")
	PhpInfoPrint("<body><div class=\"center\">\n")
}

/* }}} */

func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(f.val.value.ptr).name, (*zend.ZendModuleEntry)(s.val.value.ptr).name)
}

/* }}} */

func PhpPrintInfo(flag int) {
	var env **byte
	var tmp1 **byte
	var tmp2 **byte
	var php_uname *zend.ZendString
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpPrintInfoHtmlhead()
	} else {
		PhpInfoPrint("phpinfo()\n")
	}
	if (flag & 1 << 0) != 0 {
		var zend_version *byte = zend.GetZendVersion()
		var temp_api []byte
		php_uname = PhpGetUname('a')
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrintBoxStart(1)
		}
		if core.sapi_module.phpinfo_as_text == 0 {
			var the_time int64
			var ta *__struct__tm
			var tmbuf __struct__tm
			the_time = time(nil)
			ta = localtime_r(&the_time, &tmbuf)
			PhpInfoPrint("<a href=\"http://www.php.net/\"><img border=\"0\" src=\"")
			if ta != nil && ta.tm_mon == 3 && ta.tm_mday == 1 {
				PhpInfoPrint("data:image/png;base64,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" + "\" alt=\"PHP logo\" /></a>")
			} else {
				PhpInfoPrint("data:image/png;base64,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" + "\" alt=\"PHP logo\" /></a>")
			}
		}
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrintf("<h1 class=\"p\">PHP Version %s</h1>\n", "7.4.33")
		} else {
			PhpInfoPrintTableRow(2, "PHP Version", "7.4.33")
		}
		PhpInfoPrintBoxEnd()
		PhpInfoPrintTableStart()
		PhpInfoPrintTableRow(2, "System", php_uname.val)
		PhpInfoPrintTableRow(2, "Build Date", __DATE__+" "+__TIME__)
		PhpInfoPrintTableRow(2, "Configure Command", " './configure'  '--with-iconv=/opt/homebrew/Cellar/libiconv/1.16'")
		if core.sapi_module.pretty_name != nil {
			PhpInfoPrintTableRow(2, "Server API", core.sapi_module.pretty_name)
		}
		PhpInfoPrintTableRow(2, "Virtual Directory Support", "disabled")
		PhpInfoPrintTableRow(2, "Configuration File (php.ini) Path", "/usr/local/lib")
		PhpInfoPrintTableRow(2, "Loaded Configuration File", g.Cond(PhpIniOpenedPath != nil, PhpIniOpenedPath, "(none)"))
		PhpInfoPrintTableRow(2, "Scan this dir for additional .ini files", g.Cond(PhpIniScannedPath != nil, PhpIniScannedPath, "(none)"))
		PhpInfoPrintTableRow(2, "Additional .ini files parsed", g.Cond(PhpIniScannedFiles != nil, PhpIniScannedFiles, "(none)"))
		core.ApPhpSnprintf(temp_api, g.SizeOf("temp_api"), "%d", 20190902)
		PhpInfoPrintTableRow(2, "PHP API", temp_api)
		core.ApPhpSnprintf(temp_api, g.SizeOf("temp_api"), "%d", 20190902)
		PhpInfoPrintTableRow(2, "PHP Extension", temp_api)
		core.ApPhpSnprintf(temp_api, g.SizeOf("temp_api"), "%d", 320190902)
		PhpInfoPrintTableRow(2, "Zend Extension", temp_api)
		PhpInfoPrintTableRow(2, "Zend Extension Build", "API"+"320190902"+",NTS")
		PhpInfoPrintTableRow(2, "PHP Extension Build", "API"+"20190902"+",NTS")
		PhpInfoPrintTableRow(2, "Debug Build", "no")
		PhpInfoPrintTableRow(2, "Thread Safety", "disabled")
		PhpInfoPrintTableRow(2, "Zend Signal Handling", "enabled")
		PhpInfoPrintTableRow(2, "Zend Memory Manager", g.Cond(zend.IsZendMm() != 0, "enabled", "disabled"))
		var functions *zend.ZendMultibyteFunctions = zend.ZendMultibyteGetFunctions()
		var descr *byte
		if functions != nil {
			zend.ZendSpprintf(&descr, 0, "provided by %s", functions.provider_name)
		} else {
			descr = zend._estrdup("disabled")
		}
		PhpInfoPrintTableRow(2, "Zend Multibyte Support", descr)
		zend._efree(descr)
		PhpInfoPrintTableRow(2, "IPv6 Support", "enabled")
		PhpInfoPrintTableRow(2, "DTrace Support", "disabled")
		PhpInfoPrintStreamHash("PHP Streams", streams._phpStreamGetUrlStreamWrappersHash())
		PhpInfoPrintStreamHash("Stream Socket Transports", streams.PhpStreamXportGetHash())
		PhpInfoPrintStreamHash("Stream Filters", streams._phpGetStreamFiltersHash())
		PhpInfoPrintTableEnd()

		/* Zend Engine */

		PhpInfoPrintBoxStart(0)
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<a href=\"http://www.zend.com/\"><img border=\"0\" src=\"")
			PhpInfoPrint("data:image/png;base64,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" + "\" alt=\"Zend logo\" /></a>\n")
		}
		PhpInfoPrint("This program makes use of the Zend Scripting Language Engine:")
		PhpInfoPrint(g.Cond(core.sapi_module.phpinfo_as_text == 0, "<br />", "\n"))
		if core.sapi_module.phpinfo_as_text != 0 {
			PhpInfoPrint(zend_version)
		} else {
			zend.ZendHtmlPuts(zend_version, strlen(zend_version))
		}
		PhpInfoPrintBoxEnd()
		zend.ZendStringFree(php_uname)
	}
	zend.ZendIniSortEntries()
	if (flag & 1 << 2) != 0 {
		PhpInfoPrintHr()
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<h1>Configuration</h1>\n")
		} else {
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrint("<h2>" + "Configuration" + "</h2>\n")
			} else {
				PhpInfoPrintTableStart()
				PhpInfoPrintTableHeader(1, "Configuration")
				PhpInfoPrintTableEnd()
			}
		}
		if (flag & 1 << 3) == 0 {
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrint("<h2>" + "PHP Core" + "</h2>\n")
			} else {
				PhpInfoPrintTableStart()
				PhpInfoPrintTableHeader(1, "PHP Core")
				PhpInfoPrintTableEnd()
			}
			core.DisplayIniEntries(nil)
		}
	}
	if (flag & 1 << 3) != 0 {
		var sorted_registry zend.HashTable
		var module *zend.ZendModuleEntry
		zend._zendHashInit(&sorted_registry, &zend.ModuleRegistry.nNumOfElements, nil, 1)
		zend.ZendHashCopy(&sorted_registry, &zend.ModuleRegistry, nil)
		zend.ZendHashSortEx(&sorted_registry, zend.ZendSort, ModuleNameCmp, 0)
		for {
			var __ht *zend.HashTable = &sorted_registry
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				module = _z.value.ptr
				if module.info_func != nil || module.version != nil {
					PhpInfoPrintModule(module)
				}
			}
			break
		}
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<h2>" + "Additional Modules" + "</h2>\n")
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableHeader(1, "Additional Modules")
			PhpInfoPrintTableEnd()
		}
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "Module Name")
		for {
			var __ht *zend.HashTable = &sorted_registry
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				module = _z.value.ptr
				if module.info_func == nil && module.version == nil {
					PhpInfoPrintModule(module)
				}
			}
			break
		}
		PhpInfoPrintTableEnd()
		zend.ZendHashDestroy(&sorted_registry)
	}
	if (flag & 1 << 4) != 0 {
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<h2>" + "Environment" + "</h2>\n")
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableHeader(1, "Environment")
			PhpInfoPrintTableEnd()
		}
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(2, "Variable", "Value")
		tsrm_env_lock()
		for env = cli.Environ; env != nil && (*env) != nil; env++ {
			tmp1 = zend._estrdup(*env)
			if !(g.Assign(&tmp2, strchr(tmp1, '='))) {
				zend._efree(tmp1)
				continue
			}
			*tmp2 = 0
			tmp2++
			PhpInfoPrintTableRow(2, tmp1, tmp2)
			zend._efree(tmp1)
		}
		tsrm_env_unlock()
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 5) != 0 {
		var data *zend.Zval
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<h2>" + "PHP Variables" + "</h2>\n")
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableHeader(1, "PHP Variables")
			PhpInfoPrintTableEnd()
		}
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(2, "Variable", "Value")
		if g.Assign(&data, zend.ZendHashStrFind(&zend.EG.symbol_table, "PHP_SELF", g.SizeOf("\"PHP_SELF\"")-1)) != nil && data.u1.v.type_ == 6 {
			PhpInfoPrintTableRow(2, "PHP_SELF", data.value.str.val)
		}
		if g.Assign(&data, zend.ZendHashStrFind(&zend.EG.symbol_table, "PHP_AUTH_TYPE", g.SizeOf("\"PHP_AUTH_TYPE\"")-1)) != nil && data.u1.v.type_ == 6 {
			PhpInfoPrintTableRow(2, "PHP_AUTH_TYPE", data.value.str.val)
		}
		if g.Assign(&data, zend.ZendHashStrFind(&zend.EG.symbol_table, "PHP_AUTH_USER", g.SizeOf("\"PHP_AUTH_USER\"")-1)) != nil && data.u1.v.type_ == 6 {
			PhpInfoPrintTableRow(2, "PHP_AUTH_USER", data.value.str.val)
		}
		if g.Assign(&data, zend.ZendHashStrFind(&zend.EG.symbol_table, "PHP_AUTH_PW", g.SizeOf("\"PHP_AUTH_PW\"")-1)) != nil && data.u1.v.type_ == 6 {
			PhpInfoPrintTableRow(2, "PHP_AUTH_PW", data.value.str.val)
		}
		PhpPrintGpcseArray("_REQUEST", g.SizeOf("\"_REQUEST\"")-1)
		PhpPrintGpcseArray("_GET", g.SizeOf("\"_GET\"")-1)
		PhpPrintGpcseArray("_POST", g.SizeOf("\"_POST\"")-1)
		PhpPrintGpcseArray("_FILES", g.SizeOf("\"_FILES\"")-1)
		PhpPrintGpcseArray("_COOKIE", g.SizeOf("\"_COOKIE\"")-1)
		PhpPrintGpcseArray("_SERVER", g.SizeOf("\"_SERVER\"")-1)
		PhpPrintGpcseArray("_ENV", g.SizeOf("\"_ENV\"")-1)
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 1) != 0 {
		PhpInfoPrintHr()
		PhpPrintCredits(0xffffffff & ^(1 << 5))
	}
	if (flag & 1 << 6) != 0 {
		if core.sapi_module.phpinfo_as_text == 0 {
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrint("<h2>" + "PHP License" + "</h2>\n")
			} else {
				PhpInfoPrintTableStart()
				PhpInfoPrintTableHeader(1, "PHP License")
				PhpInfoPrintTableEnd()
			}
			PhpInfoPrintBoxStart(0)
			PhpInfoPrint("<p>\n")
			PhpInfoPrint("This program is free software; you can redistribute it and/or modify ")
			PhpInfoPrint("it under the terms of the PHP License as published by the PHP Group ")
			PhpInfoPrint("and included in the distribution in the file:  LICENSE\n")
			PhpInfoPrint("</p>\n")
			PhpInfoPrint("<p>")
			PhpInfoPrint("This program is distributed in the hope that it will be useful, ")
			PhpInfoPrint("but WITHOUT ANY WARRANTY; without even the implied warranty of ")
			PhpInfoPrint("MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.\n")
			PhpInfoPrint("</p>\n")
			PhpInfoPrint("<p>")
			PhpInfoPrint("If you did not receive a copy of the PHP license, or have any questions about ")
			PhpInfoPrint("PHP licensing, please contact license@php.net.\n")
			PhpInfoPrint("</p>\n")
			PhpInfoPrintBoxEnd()
		} else {
			PhpInfoPrint("\nPHP License\n")
			PhpInfoPrint("This program is free software; you can redistribute it and/or modify\n")
			PhpInfoPrint("it under the terms of the PHP License as published by the PHP Group\n")
			PhpInfoPrint("and included in the distribution in the file:  LICENSE\n")
			PhpInfoPrint("\n")
			PhpInfoPrint("This program is distributed in the hope that it will be useful,\n")
			PhpInfoPrint("but WITHOUT ANY WARRANTY; without even the implied warranty of\n")
			PhpInfoPrint("MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.\n")
			PhpInfoPrint("\n")
			PhpInfoPrint("If you did not receive a copy of the PHP license, or have any\n")
			PhpInfoPrint("questions about PHP licensing, please contact license@php.net.\n")
		}
	}
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("</div></body></html>")
	}
}

/* }}} */

func PhpInfoPrintTableStart() {
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("<table>\n")
	} else {
		PhpInfoPrint("\n")
	}
}

/* }}} */

func PhpInfoPrintTableEnd() {
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("</table>\n")
	}
}

/* }}} */

func PhpInfoPrintBoxStart(flag int) {
	PhpInfoPrintTableStart()
	if flag != 0 {
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<tr class=\"h\"><td>\n")
		}
	} else {
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<tr class=\"v\"><td>\n")
		} else {
			PhpInfoPrint("\n")
		}
	}
}

/* }}} */

func PhpInfoPrintBoxEnd() {
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("</td></tr>\n")
	}
	PhpInfoPrintTableEnd()
}

/* }}} */

func PhpInfoPrintHr() {
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("<hr />\n")
	} else {
		PhpInfoPrint("\n\n _______________________________________________________________________\n\n")
	}
}

/* }}} */

func PhpInfoPrintTableColspanHeader(num_cols int, header string) {
	var spaces int
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrintf("<tr class=\"h\"><th colspan=\"%d\">%s</th></tr>\n", num_cols, header)
	} else {
		spaces = int(74 - strlen(header))
		PhpInfoPrintf("%*s%s%*s\n", int(spaces/2), " ", header, int(spaces/2), " ")
	}
}

/* }}} */

func PhpInfoPrintTableHeader(num_cols int, _ ...any) {
	var i int
	var row_elements va_list
	var row_element *byte
	va_start(row_elements, num_cols)
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("<tr class=\"h\">")
	}
	for i = 0; i < num_cols; i++ {
		row_element = __va_arg(row_elements, (*byte)(_))
		if row_element == nil || !(*row_element) {
			row_element = " "
		}
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint("<th>")
			PhpInfoPrint(row_element)
			PhpInfoPrint("</th>")
		} else {
			PhpInfoPrint(row_element)
			if i < num_cols-1 {
				PhpInfoPrint(" => ")
			} else {
				PhpInfoPrint("\n")
			}
		}
	}
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("</tr>\n")
	}
	va_end(row_elements)
}

/* }}} */

func PhpInfoPrintTableRowInternal(num_cols int, value_class *byte, row_elements va_list) {
	var i int
	var row_element *byte
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("<tr>")
	}
	for i = 0; i < num_cols; i++ {
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrintf("<td class=\"%s\">", g.Cond(i == 0, "e", value_class))
		}
		row_element = __va_arg(row_elements, (*byte)(_))
		if row_element == nil || !(*row_element) {
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrint("<i>no value</i>")
			} else {
				PhpInfoPrint(" ")
			}
		} else {
			if core.sapi_module.phpinfo_as_text == 0 {
				PhpInfoPrintHtmlEsc(row_element, strlen(row_element))
			} else {
				PhpInfoPrint(row_element)
				if i < num_cols-1 {
					PhpInfoPrint(" => ")
				}
			}
		}
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrint(" </td>")
		} else if i == num_cols-1 {
			PhpInfoPrint("\n")
		}
	}
	if core.sapi_module.phpinfo_as_text == 0 {
		PhpInfoPrint("</tr>\n")
	}
}

/* }}} */

func PhpInfoPrintTableRow(num_cols int, _ ...any) {
	var row_elements va_list
	va_start(row_elements, num_cols)
	PhpInfoPrintTableRowInternal(num_cols, "v", row_elements)
	va_end(row_elements)
}

/* }}} */

func PhpInfoPrintTableRowEx(num_cols int, value_class *byte, _ ...any) {
	var row_elements va_list
	va_start(row_elements, value_class)
	PhpInfoPrintTableRowInternal(num_cols, value_class, row_elements)
	va_end(row_elements)
}

/* }}} */

func RegisterPhpinfoConstants(type_ int, module_number int) {
	zend.ZendRegisterLongConstant("INFO_GENERAL", g.SizeOf("\"INFO_GENERAL\"")-1, 1<<0, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_CREDITS", g.SizeOf("\"INFO_CREDITS\"")-1, 1<<1, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_CONFIGURATION", g.SizeOf("\"INFO_CONFIGURATION\"")-1, 1<<2, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_MODULES", g.SizeOf("\"INFO_MODULES\"")-1, 1<<3, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_ENVIRONMENT", g.SizeOf("\"INFO_ENVIRONMENT\"")-1, 1<<4, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_VARIABLES", g.SizeOf("\"INFO_VARIABLES\"")-1, 1<<5, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_LICENSE", g.SizeOf("\"INFO_LICENSE\"")-1, 1<<6, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("INFO_ALL", g.SizeOf("\"INFO_ALL\"")-1, 0xffffffff, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_GROUP", g.SizeOf("\"CREDITS_GROUP\"")-1, 1<<0, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_GENERAL", g.SizeOf("\"CREDITS_GENERAL\"")-1, 1<<1, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_SAPI", g.SizeOf("\"CREDITS_SAPI\"")-1, 1<<2, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_MODULES", g.SizeOf("\"CREDITS_MODULES\"")-1, 1<<3, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_DOCS", g.SizeOf("\"CREDITS_DOCS\"")-1, 1<<4, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_FULLPAGE", g.SizeOf("\"CREDITS_FULLPAGE\"")-1, 1<<5, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_QA", g.SizeOf("\"CREDITS_QA\"")-1, 1<<6, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("CREDITS_ALL", g.SizeOf("\"CREDITS_ALL\"")-1, 0xffffffff, 1<<1|1<<0, module_number)
}

/* }}} */

func ZifPhpinfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var flag zend.ZendLong = 0xffffffff
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &flag, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Andale!  Andale!  Yee-Hah! */

	core.PhpOutputStartDefault()
	PhpPrintInfo(int(flag))
	core.PhpOutputEnd()
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifPhpversion(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ext_name *byte = nil
	var ext_name_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &ext_name, &ext_name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if ext_name == nil {
		var _s *byte = "7.4.33"
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var version *byte
		version = zend.ZendGetModuleVersion(ext_name)
		if version == nil {
			return_value.u1.type_info = 2
			return
		}
		var _s *byte = version
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func ZifPhpcredits(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var flag zend.ZendLong = 0xffffffff
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &flag, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	PhpPrintCredits(int(flag))
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifPhpSapiName(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if core.sapi_module.name != nil {
		var _s *byte = core.sapi_module.name
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifPhpUname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var mode *byte = "a"
	var modelen int = g.SizeOf("\"a\"") - 1
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &mode, &modelen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpGetUname(*mode)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifPhpIniScannedFiles(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if PhpIniScannedFiles != nil {
		var _s *byte = PhpIniScannedFiles
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifPhpIniLoadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if PhpIniOpenedPath != nil {
		var _s *byte = PhpIniOpenedPath
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */
