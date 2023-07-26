package core

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

func PhpIniDisplayerCb(ini_entry *zend.ZendIniEntry, type_ int) {
	if ini_entry.GetDisplayer() != nil {
		ini_entry.GetDisplayer()(ini_entry, type_)
	} else {
		var display_string *byte
		var display_string_length int
		var esc_html int = 0
		if type_ == zend.ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil && ini_entry.GetOrigValue().GetStr()[0] {
				display_string = ini_entry.GetOrigValue().GetVal()
				display_string_length = ini_entry.GetOrigValue().GetLen()
				esc_html = !(SM__().GetPhpinfoAsText())
			} else {
				if SM__().GetPhpinfoAsText() == 0 {
					display_string = "<i>no value</i>"
					display_string_length = b.SizeOf("\"<i>no value</i>\"") - 1
				} else {
					display_string = "no value"
					display_string_length = b.SizeOf("\"no value\"") - 1
				}
			}
		} else if ini_entry.GetValue() != nil && ini_entry.GetValue().GetStr()[0] {
			display_string = ini_entry.GetValue().GetVal()
			display_string_length = ini_entry.GetValue().GetLen()
			esc_html = !(SM__().GetPhpinfoAsText())
		} else {
			if SM__().GetPhpinfoAsText() == 0 {
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
			PUTS(b.CastStr(display_string, display_string_length))
		}
	}
}
func DisplayIniEntries(module *zend.ModuleEntry) {
	var module_number int
	var first bool = true
	if module != nil {
		module_number = module.GetModuleNumber()
	} else {
		module_number = 0
	}
	zend.EG__().IniDirectives().Foreach(func(_ string, ini_entry *zend.ZendIniEntry) {
		if ini_entry.GetModuleNumber() != module_number {
			return
		}
		if first {
			standard.PhpInfoPrintTableStart()
			standard.PhpInfoPrintTableHeader(3, "Directive", "Local Value", "Master Value")
			first = false
		}
		if SM__().GetPhpinfoAsText() == 0 {
			PUTS("<tr>")
			PUTS("<td class=\"e\">")
			PUTS(ini_entry.GetName().GetStr())
			PUTS("</td><td class=\"v\">")
			PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ACTIVE)
			PUTS("</td><td class=\"v\">")
			PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ORIG)
			PUTS("</td></tr>\n")
		} else {
			PUTS(ini_entry.GetName().GetStr())
			PUTS(" => ")
			PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ACTIVE)
			PUTS(" => ")
			PhpIniDisplayerCb(ini_entry, zend.ZEND_INI_DISPLAY_ORIG)
			PUTS("\n")
		}
	})
	if !first {
		standard.PhpInfoPrintTableEnd()
	}
}
func RESET_ACTIVE_INI_HASH() {
	ActiveIniHash = nil
	IsSpecialSection = 0
}
func PhpIniParserCb(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, target_hash *types.Array) {
	var entry *types.Zval
	var active_hash *types.Array
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

		if IsSpecialSection == 0 && !(strcasecmp(arg1.String().GetVal(), PHP_EXTENSION_TOKEN)) {
			extension_name = zend.Estrndup(arg2.String().GetVal(), arg2.String().GetLen())
			ExtensionLists.GetFunctions().AddElement(&extension_name)
		} else if IsSpecialSection == 0 && !(strcasecmp(arg1.String().GetVal(), ZEND_EXTENSION_TOKEN)) {
			extension_name = zend.Estrndup(arg2.String().GetVal(), arg2.String().GetLen())
			ExtensionLists.GetEngine().AddElement(&extension_name)
		} else {

			/* Store in active hash */

			entry = active_hash.KeyUpdate(arg1.String().GetStr(), arg2)
			entry.SetStringVal(entry.StringVal())
		}

		/* PHP and Zend extensions are not added into configuration hash! */

	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var option_arr types.Zval
		var find_arr *types.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}

		/* fprintf(stdout, "ZEND_INI_PARSER_POP_ENTRY: %s[%s] = %s\n",Z_STRVAL_P(arg1), Z_STRVAL_P(arg3), Z_STRVAL_P(arg2)); */

		if lang.Assign(&find_arr, active_hash.KeyFind(arg1.String().GetStr())) == nil || !find_arr.IsArray() {
			option_arr.SetArray(types.NewArray(8))
			find_arr = active_hash.KeyUpdate(arg1.String().GetStr(), &option_arr)
		}

		/* arg3 is possible option offset name */

		if arg3 != nil && arg3.String().GetLen() > 0 {
			entry = find_arr.Array().SymtableUpdate(arg3.String().GetStr(), arg2)
		} else {
			entry = find_arr.Array().Append(arg2)
		}
		entry.SetStringVal(entry.StringVal())
	case zend.ZEND_INI_PARSER_SECTION:
		var key string
		/* PATH sections */
		lcArg1 := ascii.StrToUpper(arg1.StringVal())
		if strings.HasPrefix(lcArg1, "path") {
			key = arg1.StringVal()[4:]
			IsSpecialSection = 1
			HasPerDirConfig = 1
		} else if strings.HasPrefix(lcArg1, "host") {
			key = lcArg1[4:]
			IsSpecialSection = 1
			HasPerHostConfig = 1
		} else {
			IsSpecialSection = 0
		}
		if key != "" {
			/* Strip any trailing slashes */
			key = strings.TrimRight(key, "/\\")

			/* Strip any leading whitespace and '=' */
			key = strings.TrimLeft(key, "= \t")

			/* Search for existing entry and if it does not exist create one */
			entry = target_hash.KeyFind(key)
			if entry == nil {
				var section_arr types.Zval
				section_arr.SetArray(types.NewArray(8))
				entry = target_hash.KeyUpdate(key, &section_arr)
			}
			if entry.IsType(types.IS_ARRAY) {
				ActiveIniHash = entry.Array()
			}
		}
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
				PhpError(faults.E_CORE_WARNING, "Failed loading Zend extension '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
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
	var opened_path *types.String = nil
	var fp *r.File
	var filename *byte
	Config().Init()
	if SM__().GetIniDefaults() != nil {
		SM__().GetIniDefaults()(Config().GetHash())
	}
	ExtensionLists.GetEngine().Init(b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	ExtensionLists.GetFunctions().Init(b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	open_basedir = PG__().open_basedir
	if SM__().GetPhpIniPathOverride() != nil {
		php_ini_file_name = SM__().GetPhpIniPathOverride()
		php_ini_search_path = SM__().GetPhpIniPathOverride()
		free_ini_search_path = 0
	} else if SM__().GetPhpIniIgnore() == 0 {
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

		if SM__().GetPhpIniIgnoreCwd() == 0 {
			if *php_ini_search_path {
				strlcat(php_ini_search_path, paths_separator, search_path_size)
			}
			strlcat(php_ini_search_path, ".", search_path_size)
		}
		if PG__().php_binary {
			var separator_location *byte
			var binary_location *byte
			binary_location = zend.Estrdup(PG__().php_binary)
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
	PG__().open_basedir = nil

	/*
	 * Find and open actual ini file
	 */

	fp = nil
	filename = nil

	/* If SAPI does not want to ignore all ini files OR an overriding file/path is given.
	 * This allows disabling scanning for ini files in the PHP_CONFIG_FILE_SCAN_DIR but still
	 * load an optional ini file. */

	if SM__().GetPhpIniIgnore() == 0 || SM__().GetPhpIniPathOverride() != nil {

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
			Spprintf(&ini_fname, 0, fmt, SM__().Name())
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
	PG__().open_basedir = open_basedir
	if fp != nil {
		var fh *zend.FileHandle = zend.NewFileHandleByFp(filename, fp)
		RESET_ACTIVE_INI_HASH()
		zend.ZendParseIniFile(fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), Config().GetHash())
		var tmp types.Zval
		tmp.SetString(types.NewString(fh.GetFilename()))
		Config().Set("cfg_file_path", fh.GetFilename())
		PhpIniOpenedPath = zend.ZendStrndup(tmp.String().GetVal(), tmp.String().GetLen())
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

	if SM__().GetPhpIniIgnore() == 0 && php_ini_scanned_path_len != 0 {
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
		scanned_ini_list.Init(b.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
		bufpath = zend.Estrdup(PhpIniScannedPath)
		for debpath = bufpath; debpath != nil; debpath = endpath {
			endpath = strchr(debpath, zend.DEFAULT_DIR_SEPARATOR)
			if endpath != nil {
				*(lang.PostInc(&endpath)) = 0
			}
			if !(debpath[0]) {

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

				debpath = PHP_CONFIG_FILE_SCAN_DIR

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

			}
			lenpath = int(strlen(debpath))
			if lenpath > 0 && lang.Assign(&ndir, PhpScandir(debpath, &namelist, 0, PhpAlphasort)) > 0 {
				for i = 0; i < ndir; i++ {

					/* check for any file with .ini extension */

					if !(lang.Assign(&p, strrchr(namelist[i].d_name, '.'))) || p != nil && strcmp(p, ".ini") {
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
							var fh *zend.FileHandle = zend.NewFileHandleByOpenFile(ini_file)
							if fh != nil {
								if zend.ZendParseIniFile(fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), Config().GetHash()) == types.SUCCESS {

									/* Here, add it to the list of ini files read */

									l = int(strlen(ini_file))
									total_l += l + 2
									p = zend.Estrndup(ini_file, l)
									scanned_ini_list.AddElement(&p)
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
			var php_ini_scanned_files_len int = lang.CondF1(PhpIniScannedFiles != nil, func() int { return int(strlen(PhpIniScannedFiles) + 1) }, 0)
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
				strlcat(PhpIniScannedFiles, lang.Cond(element.GetNext() != nil, ",\n", "\n"), total_l)
			}
		}
		scanned_ini_list.Destroy()
	} else {

		/* Make sure an empty php_ini_scanned_path ends up as NULL */

		PhpIniScannedPath = nil

		/* Make sure an empty php_ini_scanned_path ends up as NULL */

	}
	if SM__().GetIniEntries() != nil {

		/* Reset active ini section */

		RESET_ACTIVE_INI_HASH()
		zend.ZendParseIniString(SM__().GetIniEntries(), 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), Config().GetHash())
	}
	return types.SUCCESS
}
func PhpShutdownConfig() int {
	Config().Destroy()
	if PhpIniOpenedPath != nil {
		zend.Free(PhpIniOpenedPath)
		PhpIniOpenedPath = nil
	}
	if PhpIniScannedFiles != nil {
		zend.Free(PhpIniScannedFiles)
		PhpIniScannedFiles = nil
	}
	return types.SUCCESS
}
func PhpIniRegisterExtensions() {
	ExtensionLists.GetEngine().Apply(PhpLoadZendExtensionCb)
	ExtensionLists.GetFunctions().Apply(PhpLoadPhpExtensionCb)
	ExtensionLists.GetEngine().Destroy()
	ExtensionLists.GetFunctions().Destroy()
}
func PhpParseUserIniFile(dirname *byte, ini_filename *byte, target_hash *types.Array) int {
	var sb zend.ZendStatT
	var ini_file []byte
	Snprintf(ini_file, MAXPATHLEN, "%s%c%s", dirname, zend.DEFAULT_SLASH, ini_filename)
	if zend.VCWD_STAT(ini_file, &sb) == 0 {
		if zend.S_ISREG(sb.st_mode) {
			var fh *zend.FileHandle = zend.NewFileHandleByOpenFile(ini_file)
			if fh != nil {

				/* Reset active ini section */

				RESET_ACTIVE_INI_HASH()
				if zend.ZendParseIniFile(fh, 1, zend.ZEND_INI_SCANNER_NORMAL, zend.ZendIniParserCbT(PhpIniParserCb), target_hash) == types.SUCCESS {

					/* FIXME: Add parsed file to the list of user files read? */

					return types.SUCCESS
				}
				return types.FAILURE
			}
		}
	}
	return types.FAILURE
}
func PhpIniActivateConfig(source_hash *types.Array, modify_type int, stage int) {
	/* Walk through config hash and alter matching ini entries using the values found in the hash */
	source_hash.Foreach(func(key types.ArrayKey, data *types.Zval) {
		zend.ZendAlterIniEntryEx(key.StrKey(), data.String(), modify_type, stage, 0)
	})
}
func PhpIniHasPerDirConfig() int { return HasPerDirConfig }
func PhpIniActivatePerDirConfig(path string) {
	var tmp2 *types.Zval
	var ptr *byte
	if len(path) > MAXPATHLEN {
		return
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */

	if HasPerDirConfig != 0 && path != "" {
		ptr = path + 1
		for lang.Assign(&ptr, strchr(ptr, '/')) != nil {
			*ptr = 0

			/* Search for source array matching the path from configuration_hash */

			if lang.Assign(&tmp2, Config().KeyFind(b.CastStrAuto(path))) != nil {
				PhpIniActivateConfig(tmp2.Array(), PHP_INI_SYSTEM, PHP_INI_STAGE_ACTIVATE)
			}
			*ptr = '/'
			ptr++
		}
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */
}
func PhpIniHasPerHostConfig() int { return HasPerHostConfig }
func PhpIniActivatePerHostConfig(host string) {
	if HasPerHostConfig != 0 && host != "" {
		/* Search for source array matching the host from configuration_hash */
		if tmp := Config().KeyFind(host); tmp != nil {
			PhpIniActivateConfig(tmp.Array(), PHP_INI_SYSTEM, PHP_INI_STAGE_ACTIVATE)
		}
	}
}

func CfgGetEntry(name string) *types.Zval {
	return Config().KeyFind(name)
}
func CfgGetLong(varname string, result *zend.ZendLong) int {
	tmp := Config().KeyFind(varname)
	if tmp == nil {
		*result = 0
		return types.FAILURE
	}
	*result = operators.ZvalGetLong(tmp)
	return types.SUCCESS
}
