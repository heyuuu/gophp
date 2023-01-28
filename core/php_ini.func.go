// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

func PhpIniDisplayerCb(ini_entry *zend.ZendIniEntry, type_ int) {
	if ini_entry.GetDisplayer() != nil {
		ini_entry.GetDisplayer()(ini_entry, type_)
	} else {
		var display_string *byte
		var display_string_length int
		var esc_html int = 0
		if type_ == zend.ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil && ini_entry.GetOrigValue().GetVal()[0] {
				display_string = ini_entry.GetOrigValue().GetVal()
				display_string_length = ini_entry.GetOrigValue().GetLen()
				esc_html = !(sapi_module.GetPhpinfoAsText())
			} else {
				if sapi_module.GetPhpinfoAsText() == 0 {
					display_string = "<i>no value</i>"
					display_string_length = b.SizeOf("\"<i>no value</i>\"") - 1
				} else {
					display_string = "no value"
					display_string_length = b.SizeOf("\"no value\"") - 1
				}
			}
		} else if ini_entry.GetValue() != nil && ini_entry.GetValue().GetVal()[0] {
			display_string = ini_entry.GetValue().GetVal()
			display_string_length = ini_entry.GetValue().GetLen()
			esc_html = !(sapi_module.GetPhpinfoAsText())
		} else {
			if sapi_module.GetPhpinfoAsText() == 0 {
				display_string = "<i>no value</i>"
				display_string_length = b.SizeOf("\"<i>no value</i>\"") - 1
			} else {
				display_string = "no value"
				display_string_length = b.SizeOf("\"no value\"") - 1
			}
		}
		if esc_html != 0 {
			PhpHtmlPuts(display_string, display_string_length)
		} else {
			PHPWRITE(display_string, display_string_length)
		}
	}
}
func DisplayIniEntries(module *zend.ZendModuleEntry) {
	var module_number int
	var ini_entry *zend.ZendIniEntry
	var first zend.ZendBool = 1
	if module != nil {
		module_number = module.GetModuleNumber()
	} else {
		module_number = 0
	}
	for {
		var __ht *zend.HashTable = zend.ExecutorGlobals.GetIniDirectives()
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			ini_entry = _z.GetPtr()
			if ini_entry.GetModuleNumber() != module_number {
				continue
			}
			if first != 0 {
				standard.PhpInfoPrintTableStart()
				standard.PhpInfoPrintTableHeader(3, "Directive", "Local Value", "Master Value")
				first = 0
			}
			if sapi_module.GetPhpinfoAsText() == 0 {
				PUTS("<tr>")
				PUTS("<td class=\"e\">")
				PHPWRITE(ini_entry.GetName().GetVal(), ini_entry.GetName().GetLen())
				PUTS("</td><td class=\"v\">")
				PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ACTIVE)
				PUTS("</td><td class=\"v\">")
				PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ORIG)
				PUTS("</td></tr>\n")
			} else {
				PHPWRITE(ini_entry.GetName().GetVal(), ini_entry.GetName().GetLen())
				PUTS(" => ")
				PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ACTIVE)
				PUTS(" => ")
				PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ORIG)
				PUTS("\n")
			}
		}
		break
	}
	if first == 0 {
		standard.PhpInfoPrintTableEnd()
	}
}
func ConfigZvalDtor(zvalue *zend.Zval) {
	if zvalue.IsType(zend.IS_ARRAY) {
		zend.ZendHashDestroy(zvalue.GetArr())
		zend.Free(zvalue.GetArr())
	} else if zvalue.IsType(zend.IS_STRING) {
		zend.ZendStringReleaseEx(zvalue.GetStr(), 1)
	}
}
func RESET_ACTIVE_INI_HASH() {
	ActiveIniHash = nil
	IsSpecialSection = 0
}
func PhpIniParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, target_hash *zend.HashTable) {
	var entry *zend.Zval
	var active_hash *zend.HashTable
	var extension_name *byte
	if ActiveIniHash != nil {
		active_hash = ActiveIniHash
	} else {
		active_hash = target_hash
	}
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}

		/* PHP and Zend extensions are not added into configuration hash! */

		if IsSpecialSection == 0 && !(strcasecmp(zend.Z_STRVAL_P(arg1), PHP_EXTENSION_TOKEN)) {
			extension_name = zend.Estrndup(zend.Z_STRVAL_P(arg2), zend.Z_STRLEN_P(arg2))
			zend.ZendLlistAddElement(ExtensionLists.GetFunctions(), &extension_name)
		} else if IsSpecialSection == 0 && !(strcasecmp(zend.Z_STRVAL_P(arg1), ZEND_EXTENSION_TOKEN)) {
			extension_name = zend.Estrndup(zend.Z_STRVAL_P(arg2), zend.Z_STRLEN_P(arg2))
			zend.ZendLlistAddElement(ExtensionLists.GetEngine(), &extension_name)
		} else {

			/* Store in active hash */

			entry = zend.ZendHashUpdate(active_hash, arg1.GetStr(), arg2)
			entry.SetStr(zend.ZendStringDup(entry.GetStr(), 1))
		}

		/* PHP and Zend extensions are not added into configuration hash! */

		break
	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var option_arr zend.Zval
		var find_arr *zend.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}

		/* fprintf(stdout, "ZEND_INI_PARSER_POP_ENTRY: %s[%s] = %s\n",Z_STRVAL_P(arg1), Z_STRVAL_P(arg3), Z_STRVAL_P(arg2)); */

		if b.Assign(&find_arr, zend.ZendHashFind(active_hash, arg1.GetStr())) == nil || find_arr.GetType() != zend.IS_ARRAY {
			zend.ZVAL_NEW_PERSISTENT_ARR(&option_arr)
			zend.ZendHashInit(option_arr.GetArr(), 8, nil, ConfigZvalDtor, 1)
			find_arr = zend.ZendHashUpdate(active_hash, arg1.GetStr(), &option_arr)
		}

		/* arg3 is possible option offset name */

		if arg3 != nil && zend.Z_STRLEN_P(arg3) > 0 {
			entry = zend.ZendSymtableUpdate(find_arr.GetArr(), arg3.GetStr(), arg2)
		} else {
			entry = zend.ZendHashNextIndexInsert(find_arr.GetArr(), arg2)
		}
		entry.SetStr(zend.ZendStringDup(entry.GetStr(), 1))
		break
	case zend.ZEND_INI_PARSER_SECTION:

		/* fprintf(stdout, "ZEND_INI_PARSER_SECTION: %s\n",Z_STRVAL_P(arg1)); */

		var key *byte = nil
		var key_len int

		/* PATH sections */

		if zend.ZendBinaryStrncasecmp(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1), "PATH", b.SizeOf("\"PATH\"")-1, b.SizeOf("\"PATH\"")-1) == 0 {
			key = zend.Z_STRVAL_P(arg1)
			key = key + b.SizeOf("\"PATH\"") - 1
			key_len = zend.Z_STRLEN_P(arg1) - b.SizeOf("\"PATH\"") + 1
			IsSpecialSection = 1
			HasPerDirConfig = 1

			/* make the path lowercase on Windows, for case insensitivity. Does nothing for other platforms */

			/* make the path lowercase on Windows, for case insensitivity. Does nothing for other platforms */

		} else if zend.ZendBinaryStrncasecmp(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1), "HOST", b.SizeOf("\"HOST\"")-1, b.SizeOf("\"HOST\"")-1) == 0 {
			key = zend.Z_STRVAL_P(arg1)
			key = key + b.SizeOf("\"HOST\"") - 1
			key_len = zend.Z_STRLEN_P(arg1) - b.SizeOf("\"HOST\"") + 1
			IsSpecialSection = 1
			HasPerHostConfig = 1
			zend.ZendStrTolower(key, key_len)
		} else {
			IsSpecialSection = 0
		}
		if key != nil && key_len > 0 {

			/* Strip any trailing slashes */

			for key_len > 0 && (key[key_len-1] == '/' || key[key_len-1] == '\\') {
				key_len--
				key[key_len] = 0
			}

			/* Strip any leading whitespace and '=' */

			for (*key) && ((*key) == '=' || (*key) == ' ' || (*key) == '\t') {
				key++
				key_len--
			}

			/* Search for existing entry and if it does not exist create one */

			if b.Assign(&entry, zend.ZendHashStrFind(target_hash, key, key_len)) == nil {
				var section_arr zend.Zval
				zend.ZVAL_NEW_PERSISTENT_ARR(&section_arr)
				zend.ZendHashInit(section_arr.GetArr(), 8, nil, zend.DtorFuncT(ConfigZvalDtor), 1)
				entry = zend.ZendHashStrUpdate(target_hash, key, key_len, &section_arr)
			}
			if entry.IsType(zend.IS_ARRAY) {
				ActiveIniHash = entry.GetArr()
			}
		}
		break
	}
}
func PhpLoadPhpExtensionCb(arg any) {
	standard.PhpLoadExtension(*((**byte)(arg)), zend.MODULE_PERSISTENT, 0)
}
func PhpLoadZendExtensionCb(arg any) {
	var filename *byte = *((**byte)(arg))
	var length int = strlen(filename)
	void(length)
	if zend.IS_ABSOLUTE_PATH(filename, length) {
		zend.ZendLoadExtension(filename)
	} else {
		var handle any
		var libpath *byte
		var extension_dir *byte = zend.INI_STR("extension_dir")
		var slash_suffix int = 0
		var err1 *byte
		var err2 *byte
		if extension_dir != nil && extension_dir[0] {
			slash_suffix = zend.IS_SLASH(extension_dir[strlen(extension_dir)-1])
		}

		/* Try as filename first */

		if slash_suffix != 0 {
			Spprintf(&libpath, 0, "%s%s", extension_dir, filename)
		} else {
			Spprintf(&libpath, 0, "%s%c%s", extension_dir, zend.DEFAULT_SLASH, filename)
		}
		handle = any(standard.PhpLoadShlib(libpath, &err1))
		if !handle {

			/* If file does not exist, consider as extension name and build file name */

			var orig_libpath *byte = libpath
			if slash_suffix != 0 {
				Spprintf(&libpath, 0, "%s"+PHP_SHLIB_EXT_PREFIX+"%s."+PHP_SHLIB_SUFFIX, extension_dir, filename)
			} else {
				Spprintf(&libpath, 0, "%s%c"+PHP_SHLIB_EXT_PREFIX+"%s."+PHP_SHLIB_SUFFIX, extension_dir, zend.DEFAULT_SLASH, filename)
			}
			handle = any(standard.PhpLoadShlib(libpath, &err2))
			if !handle {
				PhpError(zend.E_CORE_WARNING, "Failed loading Zend extension '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
				zend.Efree(orig_libpath)
				zend.Efree(err1)
				zend.Efree(libpath)
				zend.Efree(err2)
				return
			}
			zend.Efree(orig_libpath)
			zend.Efree(err1)
		}
		zend.ZendLoadExtensionHandle(handle, libpath)
		zend.Efree(libpath)
	}
}
func PhpInitConfig() int {
	var php_ini_file_name *byte = nil
	var php_ini_search_path *byte = nil
	var php_ini_scanned_path_len int
	var open_basedir *byte
	var free_ini_search_path int = 0
	var opened_path *zend.ZendString = nil
	var fp *r.FILE
	var filename *byte
	zend.ZendHashInit(&ConfigurationHash, 8, nil, ConfigZvalDtor, 1)
	if sapi_module.GetIniDefaults() != nil {
		sapi_module.GetIniDefaults()(&ConfigurationHash)
	}
	zend.ZendLlistInit(ExtensionLists.GetEngine(), b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	zend.ZendLlistInit(ExtensionLists.GetFunctions(), b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	open_basedir = PG(open_basedir)
	if sapi_module.GetPhpIniPathOverride() != nil {
		php_ini_file_name = sapi_module.GetPhpIniPathOverride()
		php_ini_search_path = sapi_module.GetPhpIniPathOverride()
		free_ini_search_path = 0
	} else if sapi_module.GetPhpIniIgnore() == 0 {
		var search_path_size int
		var default_location *byte
		var env_location *byte
		var paths_separator []byte = []byte{zend.ZEND_PATHS_SEPARATOR, 0}
		env_location = getenv("PHPRC")
		if env_location == nil {
			env_location = ""
		}

		/*
		 * Prepare search path
		 */

		search_path_size = MAXPATHLEN*4 + int(strlen(env_location)+3+1)
		php_ini_search_path = (*byte)(zend.Emalloc(search_path_size))
		free_ini_search_path = 1
		php_ini_search_path[0] = 0

		/* Add environment location */

		if env_location[0] {
			if *php_ini_search_path {
				strlcat(php_ini_search_path, paths_separator, search_path_size)
			}
			strlcat(php_ini_search_path, env_location, search_path_size)
			php_ini_file_name = env_location
		}

		/* Add cwd (not with CLI) */

		if sapi_module.GetPhpIniIgnoreCwd() == 0 {
			if *php_ini_search_path {
				strlcat(php_ini_search_path, paths_separator, search_path_size)
			}
			strlcat(php_ini_search_path, ".", search_path_size)
		}
		if PG(php_binary) {
			var separator_location *byte
			var binary_location *byte
			binary_location = zend.Estrdup(PG(php_binary))
			separator_location = strrchr(binary_location, zend.DEFAULT_SLASH)
			if separator_location != nil && separator_location != binary_location {
				*separator_location = 0
			}
			if *php_ini_search_path {
				strlcat(php_ini_search_path, paths_separator, search_path_size)
			}
			strlcat(php_ini_search_path, binary_location, search_path_size)
			zend.Efree(binary_location)
		}

		/* Add default location */

		default_location = PHP_CONFIG_FILE_PATH
		if *php_ini_search_path {
			strlcat(php_ini_search_path, paths_separator, search_path_size)
		}
		strlcat(php_ini_search_path, default_location, search_path_size)
	}
	PG(open_basedir) = nil

	/*
	 * Find and open actual ini file
	 */

	fp = nil
	filename = nil

	/* If SAPI does not want to ignore all ini files OR an overriding file/path is given.
	 * This allows disabling scanning for ini files in the PHP_CONFIG_FILE_SCAN_DIR but still
	 * load an optional ini file. */

	if sapi_module.GetPhpIniIgnore() == 0 || sapi_module.GetPhpIniPathOverride() != nil {

		/* Check if php_ini_file_name is a file and can be opened */

		if php_ini_file_name != nil && php_ini_file_name[0] {
			var statbuf zend.ZendStatT
			if !(zend.VCWD_STAT(php_ini_file_name, &statbuf)) {
				if (statbuf.st_mode & S_IFMT) != S_IFDIR {
					fp = zend.VCWD_FOPEN(php_ini_file_name, "r")
					if fp != nil {
						filename = ExpandFilepath(php_ini_file_name, nil)
					}
				}
			}
		}

		/* Otherwise search for php-%sapi-module-name%.ini file in search path */

		if fp == nil {
			var fmt *byte = "php-%s.ini"
			var ini_fname *byte
			Spprintf(&ini_fname, 0, fmt, sapi_module.GetName())
			fp = PhpFopenWithPath(ini_fname, "r", php_ini_search_path, &opened_path)
			zend.Efree(ini_fname)
			if fp != nil {
				filename = opened_path.GetVal()
			}
		}

		/* If still no ini file found, search for php.ini file in search path */

		if fp == nil {
			fp = PhpFopenWithPath("php.ini", "r", php_ini_search_path, &opened_path)
			if fp != nil {
				filename = opened_path.GetVal()
			}
		}

		/* If still no ini file found, search for php.ini file in search path */

	}
	if free_ini_search_path != 0 {
		zend.Efree(php_ini_search_path)
	}
	PG(open_basedir) = open_basedir
	if fp != nil {
		var fh zend.ZendFileHandle
		zend.ZendStreamInitFp(&fh, fp, filename)
		RESET_ACTIVE_INI_HASH()
		zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash)
		var tmp zend.Zval
		zend.ZVAL_NEW_STR(&tmp, zend.ZendStringInit(fh.GetFilename(), strlen(fh.GetFilename()), 1))
		zend.ZendHashStrUpdate(&ConfigurationHash, "cfg_file_path", b.SizeOf("\"cfg_file_path\"")-1, &tmp)
		if opened_path != nil {
			zend.ZendStringReleaseEx(opened_path, 0)
		} else {
			zend.Efree((*byte)(fh.GetFilename()))
		}
		PhpIniOpenedPath = zend.ZendStrndup(zend.Z_STRVAL(tmp), zend.Z_STRLEN(tmp))
	}

	/* Check for PHP_INI_SCAN_DIR environment variable to override/set config file scan directory */

	PhpIniScannedPath = getenv("PHP_INI_SCAN_DIR")
	if PhpIniScannedPath == nil {

		/* Or fall back using possible --with-config-file-scan-dir setting (defaults to empty string!) */

		PhpIniScannedPath = PHP_CONFIG_FILE_SCAN_DIR

		/* Or fall back using possible --with-config-file-scan-dir setting (defaults to empty string!) */

	}
	php_ini_scanned_path_len = int(strlen(PhpIniScannedPath))

	/* Scan and parse any .ini files found in scan path if path not empty. */

	if sapi_module.GetPhpIniIgnore() == 0 && php_ini_scanned_path_len != 0 {
		var namelist **__struct__dirent
		var ndir int
		var i int
		var sb zend.ZendStatT
		var ini_file []byte
		var p *byte
		var scanned_ini_list zend.ZendLlist
		var element *zend.ZendLlistElement
		var l int
		var total_l int = 0
		var bufpath *byte
		var debpath *byte
		var endpath *byte
		var lenpath int
		zend.ZendLlistInit(&scanned_ini_list, b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
		bufpath = zend.Estrdup(PhpIniScannedPath)
		for debpath = bufpath; debpath != nil; debpath = endpath {
			endpath = strchr(debpath, zend.DEFAULT_DIR_SEPARATOR)
			if endpath != nil {
				*(b.PostInc(&endpath)) = 0
			}
			if !(debpath[0]) {

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

				debpath = PHP_CONFIG_FILE_SCAN_DIR

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

			}
			lenpath = int(strlen(debpath))
			if lenpath > 0 && b.Assign(&ndir, PhpScandir(debpath, &namelist, 0, PhpAlphasort)) > 0 {
				for i = 0; i < ndir; i++ {

					/* check for any file with .ini extension */

					if !(b.Assign(&p, strrchr(namelist[i].d_name, '.'))) || p != nil && strcmp(p, ".ini") {
						zend.Free(namelist[i])
						continue
					}

					/* Reset active ini section */

					RESET_ACTIVE_INI_HASH()
					if zend.IS_SLASH(debpath[lenpath-1]) {
						Snprintf(ini_file, MAXPATHLEN, "%s%s", debpath, namelist[i].d_name)
					} else {
						Snprintf(ini_file, MAXPATHLEN, "%s%c%s", debpath, zend.DEFAULT_SLASH, namelist[i].d_name)
					}
					if zend.VCWD_STAT(ini_file, &sb) == 0 {
						if zend.S_ISREG(sb.st_mode) {
							var fh zend.ZendFileHandle
							zend.ZendStreamInitFp(&fh, zend.VCWD_FOPEN(ini_file, "r"), ini_file)
							if fh.GetFp() != nil {
								if zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash) == zend.SUCCESS {

									/* Here, add it to the list of ini files read */

									l = int(strlen(ini_file))
									total_l += l + 2
									p = zend.Estrndup(ini_file, l)
									zend.ZendLlistAddElement(&scanned_ini_list, &p)
								}
							}
						}
					}
					zend.Free(namelist[i])
				}
				zend.Free(namelist)
			}
		}
		zend.Efree(bufpath)
		if total_l != 0 {
			var php_ini_scanned_files_len int = b.CondF1(PhpIniScannedFiles != nil, func() int { return int(strlen(PhpIniScannedFiles) + 1) }, 0)
			PhpIniScannedFiles = (*byte)(realloc(PhpIniScannedFiles, php_ini_scanned_files_len+total_l+1))
			if php_ini_scanned_files_len == 0 {
				*PhpIniScannedFiles = '0'
			}
			total_l += php_ini_scanned_files_len
			for element = scanned_ini_list.GetHead(); element != nil; element = element.GetNext() {
				if php_ini_scanned_files_len != 0 {
					strlcat(PhpIniScannedFiles, ",\n", total_l)
				}
				strlcat(PhpIniScannedFiles, *((**byte)(element.GetData())), total_l)
				strlcat(PhpIniScannedFiles, b.Cond(element.GetNext() != nil, ",\n", "\n"), total_l)
			}
		}
		zend.ZendLlistDestroy(&scanned_ini_list)
	} else {

		/* Make sure an empty php_ini_scanned_path ends up as NULL */

		PhpIniScannedPath = nil

		/* Make sure an empty php_ini_scanned_path ends up as NULL */

	}
	if sapi_module.GetIniEntries() != nil {

		/* Reset active ini section */

		RESET_ACTIVE_INI_HASH()
		zend.ZendParseIniString(sapi_module.GetIniEntries(), 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash)
	}
	return zend.SUCCESS
}
func PhpShutdownConfig() int {
	zend.ZendHashDestroy(&ConfigurationHash)
	if PhpIniOpenedPath != nil {
		zend.Free(PhpIniOpenedPath)
		PhpIniOpenedPath = nil
	}
	if PhpIniScannedFiles != nil {
		zend.Free(PhpIniScannedFiles)
		PhpIniScannedFiles = nil
	}
	return zend.SUCCESS
}
func PhpIniRegisterExtensions() {
	zend.ZendLlistApply(ExtensionLists.GetEngine(), PhpLoadZendExtensionCb)
	zend.ZendLlistApply(ExtensionLists.GetFunctions(), PhpLoadPhpExtensionCb)
	zend.ZendLlistDestroy(ExtensionLists.GetEngine())
	zend.ZendLlistDestroy(ExtensionLists.GetFunctions())
}
func PhpParseUserIniFile(dirname *byte, ini_filename *byte, target_hash *zend.HashTable) int {
	var sb zend.ZendStatT
	var ini_file []byte
	Snprintf(ini_file, MAXPATHLEN, "%s%c%s", dirname, zend.DEFAULT_SLASH, ini_filename)
	if zend.VCWD_STAT(ini_file, &sb) == 0 {
		if zend.S_ISREG(sb.st_mode) {
			var fh zend.ZendFileHandle
			zend.ZendStreamInitFp(&fh, zend.VCWD_FOPEN(ini_file, "r"), ini_file)
			if fh.GetFp() != nil {

				/* Reset active ini section */

				RESET_ACTIVE_INI_HASH()
				if zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), target_hash) == zend.SUCCESS {

					/* FIXME: Add parsed file to the list of user files read? */

					return zend.SUCCESS

					/* FIXME: Add parsed file to the list of user files read? */

				}
				return zend.FAILURE
			}
		}
	}
	return zend.FAILURE
}
func PhpIniActivateConfig(source_hash *zend.HashTable, modify_type int, stage int) {
	var str *zend.ZendString
	var data *zend.Zval

	/* Walk through config hash and alter matching ini entries using the values found in the hash */

	for {
		var __ht *zend.HashTable = source_hash
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			str = _p.GetKey()
			data = _z
			zend.ZendAlterIniEntryEx(str, data.GetStr(), modify_type, stage, 0)
		}
		break
	}

	/* Walk through config hash and alter matching ini entries using the values found in the hash */
}
func PhpIniHasPerDirConfig() int { return HasPerDirConfig }
func PhpIniActivatePerDirConfig(path *byte, path_len int) {
	var tmp2 *zend.Zval
	var ptr *byte
	if path_len > MAXPATHLEN {
		return
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */

	if HasPerDirConfig != 0 && path != nil && path_len != 0 {
		ptr = path + 1
		for b.Assign(&ptr, strchr(ptr, '/')) != nil {
			*ptr = 0

			/* Search for source array matching the path from configuration_hash */

			if b.Assign(&tmp2, zend.ZendHashStrFind(&ConfigurationHash, path, strlen(path))) != nil {
				PhpIniActivateConfig(tmp2.GetArr(), PHP_INI_SYSTEM, PHP_INI_STAGE_ACTIVATE)
			}
			*ptr = '/'
			ptr++
		}
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */
}
func PhpIniHasPerHostConfig() int { return HasPerHostConfig }
func PhpIniActivatePerHostConfig(host *byte, host_len int) {
	var tmp *zend.Zval
	if HasPerHostConfig != 0 && host != nil && host_len != 0 {

		/* Search for source array matching the host from configuration_hash */

		if b.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, host, host_len)) != nil {
			PhpIniActivateConfig(tmp.GetArr(), PHP_INI_SYSTEM, PHP_INI_STAGE_ACTIVATE)
		}

		/* Search for source array matching the host from configuration_hash */

	}
}
func CfgGetEntryEx(name *zend.ZendString) *zend.Zval {
	return zend.ZendHashFind(&ConfigurationHash, name)
}
func CfgGetEntry(name *byte, name_length int) *zend.Zval {
	return zend.ZendHashStrFind(&ConfigurationHash, name, name_length)
}
func CfgGetLong(varname *byte, result *zend.ZendLong) int {
	var tmp *zend.Zval
	if b.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = 0
		return zend.FAILURE
	}
	*result = zend.ZvalGetLong(tmp)
	return zend.SUCCESS
}
func CfgGetDouble(varname *byte, result *float64) int {
	var tmp *zend.Zval
	if b.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = float64(0)
		return zend.FAILURE
	}
	*result = zend.ZvalGetDouble(tmp)
	return zend.SUCCESS
}
func CfgGetString(varname *byte, result **byte) int {
	var tmp *zend.Zval
	if b.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = nil
		return zend.FAILURE
	}
	*result = zend.Z_STRVAL_P(tmp)
	return zend.SUCCESS
}
func PhpIniGetConfigurationHash() *zend.HashTable { return &ConfigurationHash }
