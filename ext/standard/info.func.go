// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/sapi/cli"
	"sik/zend"
)

func SECTION(name string) {
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("<h2>" + name + "</h2>\n")
	} else {
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, name)
		PhpInfoPrintTableEnd()
	}
}
func PhpInfoPrintHtmlEsc(str *byte, len_ int) int {
	var written int
	var new_str *zend.ZendString
	new_str = PhpEscapeHtmlEntities((*uint8)(str), len_, 0, ENT_QUOTES, "utf-8")
	written = core.PhpOutputWrite(new_str.GetVal(), new_str.GetLen())
	zend.ZendStringFree(new_str)
	return written
}
func PhpInfoPrintf(fmt *byte, _ ...any) int {
	var buf *byte
	var len_ int
	var written int
	var argv va_list
	va_start(argv, fmt)
	len_ = core.Vspprintf(&buf, 0, fmt, argv)
	va_end(argv)
	written = core.PhpOutputWrite(buf, len_)
	zend.Efree(buf)
	return written
}
func PhpInfoPrint(str *byte) int {
	return core.PhpOutputWrite(str, strlen(str))
}
func PhpInfoPrintStreamHash(name string, ht *zend.HashTable) {
	var key *zend.ZendString
	if ht != nil {
		if ht.GetNNumOfElements() {
			var first int = 1
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrintf("<tr><td class=\"e\">Registered %s</td><td class=\"v\">", name)
			} else {
				PhpInfoPrintf("\nRegistered %s => ", name)
			}
			var __ht *zend.HashTable = ht
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()

				key = _p.GetKey()
				if key != nil {
					if first != 0 {
						first = 0
					} else {
						PhpInfoPrint(", ")
					}
					if core.SM__().GetPhpinfoAsText() == 0 {
						PhpInfoPrintHtmlEsc(key.GetVal(), key.GetLen())
					} else {
						PhpInfoPrint(key.GetVal())
					}
				}
			}
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrint("</td></tr>\n")
			}
		} else {
			var reg_name []byte
			core.Snprintf(reg_name, b.SizeOf("reg_name"), "Registered %s", name)
			PhpInfoPrintTableRow(2, reg_name, "none registered")
		}
	} else {
		PhpInfoPrintTableRow(2, name, "disabled")
	}
}
func PhpInfoPrintModule(zend_module *zend.ZendModuleEntry) {
	if zend_module.GetInfoFunc() != nil || zend_module.GetVersion() != nil {
		if core.SM__().GetPhpinfoAsText() == 0 {
			var url_name *zend.ZendString = PhpUrlEncode(zend_module.GetName(), strlen(zend_module.GetName()))
			PhpStrtolower(url_name.GetVal(), url_name.GetLen())
			PhpInfoPrintf("<h2><a name=\"module_%s\">%s</a></h2>\n", url_name.GetVal(), zend_module.GetName())
			zend.Efree(url_name)
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableHeader(1, zend_module.GetName())
			PhpInfoPrintTableEnd()
		}
		if zend_module.GetInfoFunc() != nil {
			zend_module.GetInfoFunc()(zend_module)
		} else {
			PhpInfoPrintTableStart()
			PhpInfoPrintTableRow(2, "Version", zend_module.GetVersion())
			PhpInfoPrintTableEnd()
			zend.DISPLAY_INI_ENTRIES()
		}
	} else {
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrintf("<tr><td class=\"v\">%s</td></tr>\n", zend_module.GetName())
		} else {
			PhpInfoPrintf("%s\n", zend_module.GetName())
		}
	}
}
func PhpPrintGpcseArray(name *byte, name_length uint32) {
	var data *zend.Zval
	var tmp *zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var key *zend.ZendString
	key = zend.ZendStringInit(name, name_length, 0)
	zend.ZendIsAutoGlobal(key)
	if b.Assign(&data, zend.ZendHashFindDeref(zend.EG__().GetSymbolTable(), key)) != nil && data.IsType(zend.IS_ARRAY) {
		var __ht *zend.HashTable = data.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			num_key = _p.GetH()
			string_key = _p.GetKey()
			tmp = _z
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrint("<tr>")
				PhpInfoPrint("<td class=\"e\">")
			}
			PhpInfoPrint("$")
			PhpInfoPrint(name)
			PhpInfoPrint("['")
			if string_key != nil {
				if core.SM__().GetPhpinfoAsText() == 0 {
					PhpInfoPrintHtmlEsc(string_key.GetVal(), string_key.GetLen())
				} else {
					PhpInfoPrint(string_key.GetVal())
				}
			} else {
				PhpInfoPrintf(zend.ZEND_ULONG_FMT, num_key)
			}
			PhpInfoPrint("']")
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrint("</td><td class=\"v\">")
			} else {
				PhpInfoPrint(" => ")
			}
			zend.ZVAL_DEREF(tmp)
			if tmp.IsType(zend.IS_ARRAY) {
				if core.SM__().GetPhpinfoAsText() == 0 {
					var str *zend.ZendString = zend.ZendPrintZvalRToStr(tmp, 0)
					PhpInfoPrint("<pre>")
					PhpInfoPrintHtmlEsc(str.GetVal(), str.GetLen())
					PhpInfoPrint("</pre>")
					zend.ZendStringReleaseEx(str, 0)
				} else {
					zend.ZendPrintZvalR(tmp, 0)
				}
			} else {
				var tmp2 *zend.ZendString
				var str *zend.ZendString = zend.ZvalGetTmpString(tmp, &tmp2)
				if core.SM__().GetPhpinfoAsText() == 0 {
					if str.GetLen() == 0 {
						PhpInfoPrint("<i>no value</i>")
					} else {
						PhpInfoPrintHtmlEsc(str.GetVal(), str.GetLen())
					}
				} else {
					PhpInfoPrint(str.GetVal())
				}
				zend.ZendTmpStringRelease(tmp2)
			}
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrint("</td></tr>\n")
			} else {
				PhpInfoPrint("\n")
			}
		}
	}
	zend.ZendStringEfree(key)
}
func PhpInfoPrintStyle() {
	PhpInfoPrintf("<style type=\"text/css\">\n")
	PhpInfoPrintCss()
	PhpInfoPrintf("</style>\n")
}
func PhpGetUname(mode byte) *zend.ZendString {
	var php_uname *byte
	var tmp_uname []byte
	var buf __struct__utsname
	if uname((*__struct__utsname)(&buf)) == -1 {
		php_uname = core.PHP_UNAME
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
			core.Snprintf(tmp_uname, b.SizeOf("tmp_uname"), "%s %s %s %s %s", buf.sysname, buf.nodename, buf.release, buf.version, buf.machine)
			php_uname = tmp_uname
		}
	}
	return zend.ZendStringInit(php_uname, strlen(php_uname), 0)
}
func PhpPrintInfoHtmlhead() {
	PhpInfoPrint("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"DTD/xhtml1-transitional.dtd\">\n")
	PhpInfoPrint("<html xmlns=\"http://www.w3.org/1999/xhtml\">")
	PhpInfoPrint("<head>\n")
	PhpInfoPrintStyle()
	PhpInfoPrintf("<title>PHP %s - phpinfo()</title>", core.PHP_VERSION)
	PhpInfoPrint("<meta name=\"ROBOTS\" content=\"NOINDEX,NOFOLLOW,NOARCHIVE\" />")
	PhpInfoPrint("</head>\n")
	PhpInfoPrint("<body><div class=\"center\">\n")
}
func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(zend.Z_PTR(f.GetVal())).GetName(), (*zend.ZendModuleEntry)(zend.Z_PTR(s.GetVal())).GetName())
}
func PhpPrintInfo(flag int) {
	var env **byte
	var tmp1 **byte
	var tmp2 **byte
	var php_uname *zend.ZendString
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpPrintInfoHtmlhead()
	} else {
		PhpInfoPrint("phpinfo()\n")
	}
	if (flag & PHP_INFO_GENERAL) != 0 {
		var zend_version *byte = zend.GetZendVersion()
		var temp_api []byte
		php_uname = PhpGetUname('a')
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrintBoxStart(1)
		}
		if core.SM__().GetPhpinfoAsText() == 0 {
			var the_time int64
			var ta *__struct__tm
			var tmbuf __struct__tm
			the_time = time(nil)
			ta = core.PhpLocaltimeR(&the_time, &tmbuf)
			PhpInfoPrint("<a href=\"http://www.php.net/\"><img border=\"0\" src=\"")
			if ta != nil && ta.tm_mon == 3 && ta.tm_mday == 1 {
				PhpInfoPrint(PHP_EGG_LOGO_DATA_URI + "\" alt=\"PHP logo\" /></a>")
			} else {
				PhpInfoPrint(PHP_LOGO_DATA_URI + "\" alt=\"PHP logo\" /></a>")
			}
		}
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrintf("<h1 class=\"p\">PHP Version %s</h1>\n", core.PHP_VERSION)
		} else {
			PhpInfoPrintTableRow(2, "PHP Version", core.PHP_VERSION)
		}
		PhpInfoPrintBoxEnd()
		PhpInfoPrintTableStart()
		PhpInfoPrintTableRow(2, "System", php_uname.GetVal())
		PhpInfoPrintTableRow(2, "Build Date", __DATE__+" "+__TIME__)
		PhpInfoPrintTableRow(2, "Configure Command", core.CONFIGURE_COMMAND)
		PhpInfoPrintTableRow(2, "Server API", core.SM__().PrettyName())
		PhpInfoPrintTableRow(2, "Virtual Directory Support", "disabled")
		PhpInfoPrintTableRow(2, "Configuration File (php.ini) Path", core.PHP_CONFIG_FILE_PATH)
		PhpInfoPrintTableRow(2, "Loaded Configuration File", b.Cond(PhpIniOpenedPath != nil, PhpIniOpenedPath, "(none)"))
		PhpInfoPrintTableRow(2, "Scan this dir for additional .ini files", b.Cond(PhpIniScannedPath != nil, PhpIniScannedPath, "(none)"))
		PhpInfoPrintTableRow(2, "Additional .ini files parsed", b.Cond(PhpIniScannedFiles != nil, PhpIniScannedFiles, "(none)"))
		core.Snprintf(temp_api, b.SizeOf("temp_api"), "%d", core.PHP_API_VERSION)
		PhpInfoPrintTableRow(2, "PHP API", temp_api)
		core.Snprintf(temp_api, b.SizeOf("temp_api"), "%d", zend.ZEND_MODULE_API_NO)
		PhpInfoPrintTableRow(2, "PHP Extension", temp_api)
		core.Snprintf(temp_api, b.SizeOf("temp_api"), "%d", zend.ZEND_EXTENSION_API_NO)
		PhpInfoPrintTableRow(2, "Zend Extension", temp_api)
		PhpInfoPrintTableRow(2, "Zend Extension Build", "API"+"ZEND_EXTENSION_API_NO"+zend.ZEND_BUILD_TS)
		PhpInfoPrintTableRow(2, "PHP Extension Build", "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
		PhpInfoPrintTableRow(2, "Debug Build", "no")
		PhpInfoPrintTableRow(2, "Thread Safety", "disabled")
		PhpInfoPrintTableRow(2, "Zend Signal Handling", "enabled")
		PhpInfoPrintTableRow(2, "Zend Memory Manager", b.Cond(zend.IsZendMm() != 0, "enabled", "disabled"))
		PhpInfoPrintTableRow(2, "Zend Multibyte Support", "disabled")
		zend.Efree(descr)
		PhpInfoPrintTableRow(2, "IPv6 Support", "enabled")
		PhpInfoPrintTableRow(2, "DTrace Support", "disabled")
		PhpInfoPrintStreamHash("PHP Streams", core.PhpStreamGetUrlStreamWrappersHash())
		PhpInfoPrintStreamHash("Stream Socket Transports", streams.PhpStreamXportGetHash())
		PhpInfoPrintStreamHash("Stream Filters", core.PhpGetStreamFiltersHash())
		PhpInfoPrintTableEnd()

		/* Zend Engine */

		PhpInfoPrintBoxStart(0)
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrint("<a href=\"http://www.zend.com/\"><img border=\"0\" src=\"")
			PhpInfoPrint(ZEND_LOGO_DATA_URI + "\" alt=\"Zend logo\" /></a>\n")
		}
		PhpInfoPrint("This program makes use of the Zend Scripting Language Engine:")
		PhpInfoPrint(b.Cond(core.SM__().GetPhpinfoAsText() == 0, "<br />", "\n"))
		if core.SM__().GetPhpinfoAsText() != 0 {
			PhpInfoPrint(zend_version)
		} else {
			zend.ZendHtmlPuts(zend_version, strlen(zend_version))
		}
		PhpInfoPrintBoxEnd()
		zend.ZendStringFree(php_uname)
	}
	zend.ZendIniSortEntries()
	if (flag & PHP_INFO_CONFIGURATION) != 0 {
		PhpInfoPrintHr()
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrint("<h1>Configuration</h1>\n")
		} else {
			SECTION("Configuration")
		}
		if (flag & PHP_INFO_MODULES) == 0 {
			SECTION("PHP Core")
			core.DisplayIniEntries(nil)
		}
	}
	if (flag & PHP_INFO_MODULES) != 0 {
		var sorted_registry zend.HashTable
		var module *zend.ZendModuleEntry
		zend.ZendHashInit(&sorted_registry, zend.ModuleRegistry.GetNNumOfElements(), nil, nil, 1)
		zend.ZendHashCopy(&sorted_registry, &zend.ModuleRegistry, nil)
		sorted_registry.SortCompatible(ModuleNameCmp, 0)
		var __ht *zend.HashTable = &sorted_registry
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			module = _z.GetPtr()
			if module.GetInfoFunc() != nil || module.GetVersion() != nil {
				PhpInfoPrintModule(module)
			}
		}
		SECTION("Additional Modules")
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "Module Name")
		var __ht__1 *zend.HashTable = &sorted_registry
		for _, _p := range __ht__1.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			module = _z.GetPtr()
			if module.GetInfoFunc() == nil && module.GetVersion() == nil {
				PhpInfoPrintModule(module)
			}
		}
		PhpInfoPrintTableEnd()
		sorted_registry.Destroy()
	}
	if (flag & PHP_INFO_ENVIRONMENT) != 0 {
		SECTION("Environment")
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(2, "Variable", "Value")
		tsrm_env_lock()
		for env = cli.Environ; env != nil && (*env) != nil; env++ {
			tmp1 = zend.Estrdup(*env)
			if !(b.Assign(&tmp2, strchr(tmp1, '='))) {
				zend.Efree(tmp1)
				continue
			}
			*tmp2 = 0
			tmp2++
			PhpInfoPrintTableRow(2, tmp1, tmp2)
			zend.Efree(tmp1)
		}
		tsrm_env_unlock()
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_INFO_VARIABLES) != 0 {
		var data *zend.Zval
		SECTION("PHP Variables")
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(2, "Variable", "Value")
		if b.Assign(&data, zend.EG__().GetSymbolTable().KeyFind("PHP_SELF")) != nil && data.IsType(zend.IS_STRING) {
			PhpInfoPrintTableRow(2, "PHP_SELF", zend.Z_STRVAL_P(data))
		}
		if b.Assign(&data, zend.EG__().GetSymbolTable().KeyFind("PHP_AUTH_TYPE")) != nil && data.IsType(zend.IS_STRING) {
			PhpInfoPrintTableRow(2, "PHP_AUTH_TYPE", zend.Z_STRVAL_P(data))
		}
		if b.Assign(&data, zend.EG__().GetSymbolTable().KeyFind("PHP_AUTH_USER")) != nil && data.IsType(zend.IS_STRING) {
			PhpInfoPrintTableRow(2, "PHP_AUTH_USER", zend.Z_STRVAL_P(data))
		}
		if b.Assign(&data, zend.EG__().GetSymbolTable().KeyFind("PHP_AUTH_PW")) != nil && data.IsType(zend.IS_STRING) {
			PhpInfoPrintTableRow(2, "PHP_AUTH_PW", zend.Z_STRVAL_P(data))
		}
		PhpPrintGpcseArray(zend.ZEND_STRL("_REQUEST"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_GET"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_POST"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_FILES"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_COOKIE"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_SERVER"))
		PhpPrintGpcseArray(zend.ZEND_STRL("_ENV"))
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_INFO_CREDITS) != 0 {
		PhpInfoPrintHr()
		PhpPrintCredits(PHP_CREDITS_ALL & ^PHP_CREDITS_FULLPAGE)
	}
	if (flag & PHP_INFO_LICENSE) != 0 {
		if core.SM__().GetPhpinfoAsText() == 0 {
			SECTION("PHP License")
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
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("</div></body></html>")
	}
}
func PhpInfoPrintTableStart() {
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("<table>\n")
	} else {
		PhpInfoPrint("\n")
	}
}
func PhpInfoPrintTableEnd() {
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("</table>\n")
	}
}
func PhpInfoPrintBoxStart(flag int) {
	PhpInfoPrintTableStart()
	if flag != 0 {
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrint("<tr class=\"h\"><td>\n")
		}
	} else {
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrint("<tr class=\"v\"><td>\n")
		} else {
			PhpInfoPrint("\n")
		}
	}
}
func PhpInfoPrintBoxEnd() {
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("</td></tr>\n")
	}
	PhpInfoPrintTableEnd()
}
func PhpInfoPrintHr() {
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("<hr />\n")
	} else {
		PhpInfoPrint("\n\n _______________________________________________________________________\n\n")
	}
}
func PhpInfoPrintTableColspanHeader(num_cols int, header string) {
	var spaces int
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrintf("<tr class=\"h\"><th colspan=\"%d\">%s</th></tr>\n", num_cols, header)
	} else {
		spaces = int(74 - strlen(header))
		PhpInfoPrintf("%*s%s%*s\n", int(spaces/2), " ", header, int(spaces/2), " ")
	}
}
func PhpInfoPrintTableHeader(num_cols int, _ ...any) {
	var i int
	var row_elements va_list
	var row_element *byte
	va_start(row_elements, num_cols)
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("<tr class=\"h\">")
	}
	for i = 0; i < num_cols; i++ {
		row_element = __va_arg(row_elements, (*byte)(_))
		if row_element == nil || !(*row_element) {
			row_element = " "
		}
		if core.SM__().GetPhpinfoAsText() == 0 {
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
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("</tr>\n")
	}
	va_end(row_elements)
}
func PhpInfoPrintTableRowInternal(num_cols int, value_class *byte, row_elements ...any) {
	var i int
	var row_element *byte
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("<tr>")
	}
	for i = 0; i < num_cols; i++ {
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrintf("<td class=\"%s\">", b.Cond(i == 0, "e", value_class))
		}
		row_element = __va_arg(row_elements, (*byte)(_))
		if row_element == nil || !(*row_element) {
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrint("<i>no value</i>")
			} else {
				PhpInfoPrint(" ")
			}
		} else {
			if core.SM__().GetPhpinfoAsText() == 0 {
				PhpInfoPrintHtmlEsc(row_element, strlen(row_element))
			} else {
				PhpInfoPrint(row_element)
				if i < num_cols-1 {
					PhpInfoPrint(" => ")
				}
			}
		}
		if core.SM__().GetPhpinfoAsText() == 0 {
			PhpInfoPrint(" </td>")
		} else if i == num_cols-1 {
			PhpInfoPrint("\n")
		}
	}
	if core.SM__().GetPhpinfoAsText() == 0 {
		PhpInfoPrint("</tr>\n")
	}
}
func PhpInfoPrintTableRow(num_cols int, _ ...any) {
	var row_elements va_list
	va_start(row_elements, num_cols)
	PhpInfoPrintTableRowInternal(num_cols, "v", row_elements)
	va_end(row_elements)
}
func RegisterPhpinfoConstants(type_ int, module_number int) {
	zend.REGISTER_LONG_CONSTANT("INFO_GENERAL", PHP_INFO_GENERAL, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_CREDITS", PHP_INFO_CREDITS, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_CONFIGURATION", PHP_INFO_CONFIGURATION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_MODULES", PHP_INFO_MODULES, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_ENVIRONMENT", PHP_INFO_ENVIRONMENT, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_VARIABLES", PHP_INFO_VARIABLES, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_LICENSE", PHP_INFO_LICENSE, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("INFO_ALL", PHP_INFO_ALL, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_GROUP", PHP_CREDITS_GROUP, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_GENERAL", PHP_CREDITS_GENERAL, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_SAPI", PHP_CREDITS_SAPI, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_MODULES", PHP_CREDITS_MODULES, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_DOCS", PHP_CREDITS_DOCS, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_FULLPAGE", PHP_CREDITS_FULLPAGE, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_QA", PHP_CREDITS_QA, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_LONG_CONSTANT("CREDITS_ALL", PHP_CREDITS_ALL, zend.CONST_PERSISTENT|zend.CONST_CS)
}
func ZifPhpversion(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var ext_name *byte = nil
	var ext_name_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &ext_name, &ext_name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
		return_value.SetRawString(b.CastStrAuto(core.PHP_VERSION))
		return
	} else {
		var version *byte
		version = zend.ZendGetModuleVersion(ext_name)
		if version == nil {
			return_value.SetFalse()
			return
		}
		return_value.SetRawString(b.CastStrAuto(version))
		return
	}
}
func ZifPhpcredits(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var flag zend.ZendLong = PHP_CREDITS_ALL
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flag, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	return_value.SetTrue()
	return
}
func ZifPhpSapiName(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetRawString(core.SM__().Name())
	return
}
func ZifPhpUname(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var mode *byte = "a"
	var modelen int = b.SizeOf("\"a\"") - 1
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &mode, &modelen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	return_value.SetString(PhpGetUname(*mode))
	return
}
func ZifPhpIniScannedFiles(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if PhpIniScannedFiles != nil {
		return_value.SetRawString(b.CastStrAuto(PhpIniScannedFiles))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifPhpIniLoadedFile(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if PhpIniOpenedPath != nil {
		return_value.SetRawString(b.CastStrAuto(PhpIniOpenedPath))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
