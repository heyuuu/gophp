// <<generate>>

package core

import (
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/php_ini.h>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// #define PHP_INI_H

// # include "zend_ini.h"

// #define PHP_INI_USER       ZEND_INI_USER

// #define PHP_INI_PERDIR       ZEND_INI_PERDIR

// #define PHP_INI_SYSTEM       ZEND_INI_SYSTEM

// #define PHP_INI_ALL       ZEND_INI_ALL

// #define php_ini_entry       zend_ini_entry

// #define PHP_INI_MH       ZEND_INI_MH

// #define PHP_INI_DISP       ZEND_INI_DISP

// #define PHP_INI_BEGIN       ZEND_INI_BEGIN

// #define PHP_INI_END       ZEND_INI_END

// #define PHP_INI_ENTRY3_EX       ZEND_INI_ENTRY3_EX

// #define PHP_INI_ENTRY3       ZEND_INI_ENTRY3

// #define PHP_INI_ENTRY2_EX       ZEND_INI_ENTRY2_EX

// #define PHP_INI_ENTRY2       ZEND_INI_ENTRY2

// #define PHP_INI_ENTRY1_EX       ZEND_INI_ENTRY1_EX

// #define PHP_INI_ENTRY1       ZEND_INI_ENTRY1

// #define PHP_INI_ENTRY_EX       ZEND_INI_ENTRY_EX

// #define PHP_INI_ENTRY       ZEND_INI_ENTRY

// #define STD_PHP_INI_ENTRY       STD_ZEND_INI_ENTRY

// #define STD_PHP_INI_ENTRY_EX       STD_ZEND_INI_ENTRY_EX

// #define STD_PHP_INI_BOOLEAN       STD_ZEND_INI_BOOLEAN

// #define PHP_INI_DISPLAY_ORIG       ZEND_INI_DISPLAY_ORIG

// #define PHP_INI_DISPLAY_ACTIVE       ZEND_INI_DISPLAY_ACTIVE

// #define PHP_INI_STAGE_STARTUP       ZEND_INI_STAGE_STARTUP

// #define PHP_INI_STAGE_SHUTDOWN       ZEND_INI_STAGE_SHUTDOWN

// #define PHP_INI_STAGE_ACTIVATE       ZEND_INI_STAGE_ACTIVATE

// #define PHP_INI_STAGE_DEACTIVATE       ZEND_INI_STAGE_DEACTIVATE

// #define PHP_INI_STAGE_RUNTIME       ZEND_INI_STAGE_RUNTIME

// #define PHP_INI_STAGE_HTACCESS       ZEND_INI_STAGE_HTACCESS

// #define php_ini_boolean_displayer_cb       zend_ini_boolean_displayer_cb

// #define php_ini_color_displayer_cb       zend_ini_color_displayer_cb

// #define php_alter_ini_entry       zend_alter_ini_entry

// #define php_ini_long       zend_ini_long

// #define php_ini_double       zend_ini_double

// #define php_ini_string       zend_ini_string

// Source: <main/php_ini.c>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "ext/standard/info.h"

// # include "zend_ini.h"

// # include "zend_ini_scanner.h"

// # include "php_ini.h"

// # include "ext/standard/dl.h"

// # include "zend_extensions.h"

// # include "zend_highlight.h"

// # include "SAPI.h"

// # include "php_main.h"

// # include "php_scandir.h"

// # include < dirent . h >

// #define TRANSLATE_SLASHES_LOWER(path)

// @type PhpExtensionLists struct

/* True globals */

var IsSpecialSection int = 0
var ActiveIniHash *zend.HashTable
var ConfigurationHash zend.HashTable
var HasPerDirConfig int = 0
var HasPerHostConfig int = 0
var PhpIniOpenedPath *byte = nil
var ExtensionLists PhpExtensionLists
var PhpIniScannedPath *byte = nil
var PhpIniScannedFiles *byte = nil

/* {{{ php_ini_displayer_cb
 */

func PhpIniDisplayerCb(ini_entry *zend.ZendIniEntry, type_ int) {
	if ini_entry.displayer != nil {
		ini_entry.displayer(ini_entry, type_)
	} else {
		var display_string *byte
		var display_string_length int
		var esc_html int = 0
		if type_ == 1 && ini_entry.modified != 0 {
			if ini_entry.orig_value != nil && ini_entry.orig_value.val[0] {
				display_string = ini_entry.orig_value.val
				display_string_length = ini_entry.orig_value.len_
				esc_html = !(sapi_module.GetPhpinfoAsText())
			} else {
				if sapi_module.GetPhpinfoAsText() == 0 {
					display_string = "<i>no value</i>"
					display_string_length = g.SizeOf("\"<i>no value</i>\"") - 1
				} else {
					display_string = "no value"
					display_string_length = g.SizeOf("\"no value\"") - 1
				}
			}
		} else if ini_entry.value != nil && ini_entry.value.val[0] {
			display_string = ini_entry.value.val
			display_string_length = ini_entry.value.len_
			esc_html = !(sapi_module.GetPhpinfoAsText())
		} else {
			if sapi_module.GetPhpinfoAsText() == 0 {
				display_string = "<i>no value</i>"
				display_string_length = g.SizeOf("\"<i>no value</i>\"") - 1
			} else {
				display_string = "no value"
				display_string_length = g.SizeOf("\"no value\"") - 1
			}
		}
		if esc_html != 0 {
			PhpHtmlPuts(display_string, display_string_length)
		} else {
			PhpOutputWrite(display_string, display_string_length)
		}
	}
}

/* }}} */

func DisplayIniEntries(module *zend.ZendModuleEntry) {
	var module_number int
	var ini_entry *zend.ZendIniEntry
	var first zend.ZendBool = 1
	if module != nil {
		module_number = module.module_number
	} else {
		module_number = 0
	}
	for {
		var __ht *zend.HashTable = zend.EG.ini_directives
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			ini_entry = _z.value.ptr
			if ini_entry.module_number != module_number {
				continue
			}
			if first != 0 {
				standard.PhpInfoPrintTableStart()
				standard.PhpInfoPrintTableHeader(3, "Directive", "Local Value", "Master Value")
				first = 0
			}
			if sapi_module.GetPhpinfoAsText() == 0 {
				var __str *byte = "<tr>"
				PhpOutputWrite(__str, strlen(__str))
				var __str *byte = "<td class=\"e\">"
				PhpOutputWrite(__str, strlen(__str))
				PhpOutputWrite(ini_entry.name.val, ini_entry.name.len_)
				var __str *byte = "</td><td class=\"v\">"
				PhpOutputWrite(__str, strlen(__str))
				PhpIniDisplayerCb(ini_entry, 2)
				var __str *byte = "</td><td class=\"v\">"
				PhpOutputWrite(__str, strlen(__str))
				PhpIniDisplayerCb(ini_entry, 1)
				var __str *byte = "</td></tr>\n"
				PhpOutputWrite(__str, strlen(__str))
			} else {
				PhpOutputWrite(ini_entry.name.val, ini_entry.name.len_)
				var __str *byte = " => "
				PhpOutputWrite(__str, strlen(__str))
				PhpIniDisplayerCb(ini_entry, 2)
				var __str *byte = " => "
				PhpOutputWrite(__str, strlen(__str))
				PhpIniDisplayerCb(ini_entry, 1)
				var __str *byte = "\n"
				PhpOutputWrite(__str, strlen(__str))
			}
		}
		break
	}
	if first == 0 {
		standard.PhpInfoPrintTableEnd()
	}
}

/* }}} */

// #define PHP_EXTENSION_TOKEN       "extension"

// #define ZEND_EXTENSION_TOKEN       "zend_extension"

/* {{{ config_zval_dtor
 */

func ConfigZvalDtor(zvalue *zend.Zval) {
	if zvalue.u1.v.type_ == 7 {
		zend.ZendHashDestroy(zvalue.value.arr)
		zend.Free(zvalue.value.arr)
	} else if zvalue.u1.v.type_ == 6 {
		zend.ZendStringReleaseEx(zvalue.value.str, 1)
	}
}

/* Reset / free active_ini_sectin global */

// #define RESET_ACTIVE_INI_HASH() do { active_ini_hash = NULL ; is_special_section = 0 ; } while ( 0 )

/* }}} */

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
	case 1:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}

		/* PHP and Zend extensions are not added into configuration hash! */

		if IsSpecialSection == 0 && !(strcasecmp(arg1.value.str.val, "extension")) {
			extension_name = zend._estrndup(arg2.value.str.val, arg2.value.str.len_)
			zend.ZendLlistAddElement(&ExtensionLists.functions, &extension_name)
		} else if IsSpecialSection == 0 && !(strcasecmp(arg1.value.str.val, "zend_extension")) {
			extension_name = zend._estrndup(arg2.value.str.val, arg2.value.str.len_)
			zend.ZendLlistAddElement(&ExtensionLists.engine, &extension_name)
		} else {

			/* Store in active hash */

			entry = zend.ZendHashUpdate(active_hash, arg1.value.str, arg2)
			entry.value.str = zend.ZendStringDup(entry.value.str, 1)
		}

		/* PHP and Zend extensions are not added into configuration hash! */

		break
	case 3:
		var option_arr zend.Zval
		var find_arr *zend.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}

		/* fprintf(stdout, "ZEND_INI_PARSER_POP_ENTRY: %s[%s] = %s\n",Z_STRVAL_P(arg1), Z_STRVAL_P(arg3), Z_STRVAL_P(arg2)); */

		if g.Assign(&find_arr, zend.ZendHashFind(active_hash, arg1.value.str)) == nil || find_arr.u1.v.type_ != 7 {
			var __z *zend.Zval = &option_arr
			var _arr *zend.ZendArray = (*zend.ZendArray)(zend.Malloc(g.SizeOf("zend_array")))
			__z.value.arr = _arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend._zendHashInit(option_arr.value.arr, 8, ConfigZvalDtor, 1)
			find_arr = zend.ZendHashUpdate(active_hash, arg1.value.str, &option_arr)
		}

		/* arg3 is possible option offset name */

		if arg3 != nil && arg3.value.str.len_ > 0 {
			entry = zend.ZendSymtableUpdate(find_arr.value.arr, arg3.value.str, arg2)
		} else {
			entry = zend.ZendHashNextIndexInsert(find_arr.value.arr, arg2)
		}
		entry.value.str = zend.ZendStringDup(entry.value.str, 1)
		break
	case 2:

		/* fprintf(stdout, "ZEND_INI_PARSER_SECTION: %s\n",Z_STRVAL_P(arg1)); */

		var key *byte = nil
		var key_len int

		/* PATH sections */

		if zend.ZendBinaryStrncasecmp(arg1.value.str.val, arg1.value.str.len_, "PATH", g.SizeOf("\"PATH\"")-1, g.SizeOf("\"PATH\"")-1) == 0 {
			key = arg1.value.str.val
			key = key + g.SizeOf("\"PATH\"") - 1
			key_len = arg1.value.str.len_ - g.SizeOf("\"PATH\"") + 1
			IsSpecialSection = 1
			HasPerDirConfig = 1

			/* make the path lowercase on Windows, for case insensitivity. Does nothing for other platforms */

			/* make the path lowercase on Windows, for case insensitivity. Does nothing for other platforms */

		} else if zend.ZendBinaryStrncasecmp(arg1.value.str.val, arg1.value.str.len_, "HOST", g.SizeOf("\"HOST\"")-1, g.SizeOf("\"HOST\"")-1) == 0 {
			key = arg1.value.str.val
			key = key + g.SizeOf("\"HOST\"") - 1
			key_len = arg1.value.str.len_ - g.SizeOf("\"HOST\"") + 1
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

			if g.Assign(&entry, zend.ZendHashStrFind(target_hash, key, key_len)) == nil {
				var section_arr zend.Zval
				var __z *zend.Zval = &section_arr
				var _arr *zend.ZendArray = (*zend.ZendArray)(zend.Malloc(g.SizeOf("zend_array")))
				__z.value.arr = _arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				zend._zendHashInit(section_arr.value.arr, 8, zend.DtorFuncT(ConfigZvalDtor), 1)
				entry = zend.ZendHashStrUpdate(target_hash, key, key_len, &section_arr)
			}
			if entry.u1.v.type_ == 7 {
				ActiveIniHash = entry.value.arr
			}
		}
		break
	}
}

/* }}} */

func PhpLoadPhpExtensionCb(arg any) {
	standard.PhpLoadExtension(*((**byte)(arg)), 1, 0)
}

/* }}} */

func PhpLoadZendExtensionCb(arg any) {
	var filename *byte = *((**byte)(arg))
	var length int = strlen(filename)
	void(length)
	if filename[0] == '/' {
		zend.ZendLoadExtension(filename)
	} else {
		var handle any
		var libpath *byte
		var extension_dir *byte = zend.ZendIniStringEx("extension_dir", g.SizeOf("\"extension_dir\"")-1, 0, nil)
		var slash_suffix int = 0
		var err1 *byte
		var err2 *byte
		if extension_dir != nil && extension_dir[0] {
			slash_suffix = extension_dir[strlen(extension_dir)-1] == '/'
		}

		/* Try as filename first */

		if slash_suffix != 0 {
			zend.ZendSpprintf(&libpath, 0, "%s%s", extension_dir, filename)
		} else {
			zend.ZendSpprintf(&libpath, 0, "%s%c%s", extension_dir, '/', filename)
		}
		handle = any(standard.PhpLoadShlib(libpath, &err1))
		if !handle {

			/* If file does not exist, consider as extension name and build file name */

			var orig_libpath *byte = libpath
			if slash_suffix != 0 {
				zend.ZendSpprintf(&libpath, 0, "%s"+""+"%s."+"so", extension_dir, filename)
			} else {
				zend.ZendSpprintf(&libpath, 0, "%s%c"+""+"%s."+"so", extension_dir, '/', filename)
			}
			handle = any(standard.PhpLoadShlib(libpath, &err2))
			if !handle {
				zend.ZendError(1<<5, "Failed loading Zend extension '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
				zend._efree(orig_libpath)
				zend._efree(err1)
				zend._efree(libpath)
				zend._efree(err2)
				return
			}
			zend._efree(orig_libpath)
			zend._efree(err1)
		}
		zend.ZendLoadExtensionHandle(handle, libpath)
		zend._efree(libpath)
	}
}

/* }}} */

func PhpInitConfig() int {
	var php_ini_file_name *byte = nil
	var php_ini_search_path *byte = nil
	var php_ini_scanned_path_len int
	var open_basedir *byte
	var free_ini_search_path int = 0
	var opened_path *zend.ZendString = nil
	var fp *r.FILE
	var filename *byte
	zend._zendHashInit(&ConfigurationHash, 8, ConfigZvalDtor, 1)
	if sapi_module.GetIniDefaults() != nil {
		sapi_module.GetIniDefaults()(&ConfigurationHash)
	}
	zend.ZendLlistInit(&ExtensionLists.engine, g.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	zend.ZendLlistInit(&ExtensionLists.functions, g.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
	open_basedir = CoreGlobals.GetOpenBasedir()
	if sapi_module.GetPhpIniPathOverride() != nil {
		php_ini_file_name = sapi_module.GetPhpIniPathOverride()
		php_ini_search_path = sapi_module.GetPhpIniPathOverride()
		free_ini_search_path = 0
	} else if sapi_module.GetPhpIniIgnore() == 0 {
		var search_path_size int
		var default_location *byte
		var env_location *byte
		var paths_separator []byte = []byte{':', 0}
		env_location = getenv("PHPRC")
		if env_location == nil {
			env_location = ""
		}

		/*
		 * Prepare search path
		 */

		search_path_size = 256*4 + int(strlen(env_location)+3+1)
		php_ini_search_path = (*byte)(zend._emalloc(search_path_size))
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
		if CoreGlobals.GetPhpBinary() != nil {
			var separator_location *byte
			var binary_location *byte
			binary_location = zend._estrdup(CoreGlobals.GetPhpBinary())
			separator_location = strrchr(binary_location, '/')
			if separator_location != nil && separator_location != binary_location {
				*separator_location = 0
			}
			if *php_ini_search_path {
				strlcat(php_ini_search_path, paths_separator, search_path_size)
			}
			strlcat(php_ini_search_path, binary_location, search_path_size)
			zend._efree(binary_location)
		}

		/* Add default location */

		default_location = "/usr/local/lib"
		if *php_ini_search_path {
			strlcat(php_ini_search_path, paths_separator, search_path_size)
		}
		strlcat(php_ini_search_path, default_location, search_path_size)
	}
	CoreGlobals.SetOpenBasedir(nil)

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
			if !(stat(php_ini_file_name, &statbuf)) {
				if (statbuf.st_mode & S_IFMT) != S_IFDIR {
					fp = r.Fopen(php_ini_file_name, "r")
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
			zend.ZendSpprintf(&ini_fname, 0, fmt, sapi_module.GetName())
			fp = PhpFopenWithPath(ini_fname, "r", php_ini_search_path, &opened_path)
			zend._efree(ini_fname)
			if fp != nil {
				filename = opened_path.val
			}
		}

		/* If still no ini file found, search for php.ini file in search path */

		if fp == nil {
			fp = PhpFopenWithPath("php.ini", "r", php_ini_search_path, &opened_path)
			if fp != nil {
				filename = opened_path.val
			}
		}

		/* If still no ini file found, search for php.ini file in search path */

	}
	if free_ini_search_path != 0 {
		zend._efree(php_ini_search_path)
	}
	CoreGlobals.SetOpenBasedir(open_basedir)
	if fp != nil {
		var fh zend.ZendFileHandle
		zend.ZendStreamInitFp(&fh, fp, filename)
		ActiveIniHash = nil
		IsSpecialSection = 0
		zend.ZendParseIniFile(&fh, 1, 0, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash)
		var tmp zend.Zval
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(fh.filename, strlen(fh.filename), 1)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendHashStrUpdate(&ConfigurationHash, "cfg_file_path", g.SizeOf("\"cfg_file_path\"")-1, &tmp)
		if opened_path != nil {
			zend.ZendStringReleaseEx(opened_path, 0)
		} else {
			zend._efree((*byte)(fh.filename))
		}
		PhpIniOpenedPath = zend.ZendStrndup(tmp.value.str.val, tmp.value.str.len_)
	}

	/* Check for PHP_INI_SCAN_DIR environment variable to override/set config file scan directory */

	PhpIniScannedPath = getenv("PHP_INI_SCAN_DIR")
	if PhpIniScannedPath == nil {

		/* Or fall back using possible --with-config-file-scan-dir setting (defaults to empty string!) */

		PhpIniScannedPath = ""

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
		zend.ZendLlistInit(&scanned_ini_list, g.SizeOf("char *"), zend.LlistDtorFuncT(zend.FreeEstring), 1)
		bufpath = zend._estrdup(PhpIniScannedPath)
		for debpath = bufpath; debpath != nil; debpath = endpath {
			endpath = strchr(debpath, ':')
			if endpath != nil {
				*(g.PostInc(&endpath)) = 0
			}
			if !(debpath[0]) {

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

				debpath = ""

				/* empty string means default builtin value
				   to allow "/foo/php.d:" or ":/foo/php.d" */

			}
			lenpath = int(strlen(debpath))
			if lenpath > 0 && g.Assign(&ndir, scandir(debpath, &namelist, 0, alphasort)) > 0 {
				for i = 0; i < ndir; i++ {

					/* check for any file with .ini extension */

					if !(g.Assign(&p, strrchr(namelist[i].d_name, '.'))) || p != nil && strcmp(p, ".ini") {
						zend.Free(namelist[i])
						continue
					}

					/* Reset active ini section */

					ActiveIniHash = nil
					IsSpecialSection = 0
					if debpath[lenpath-1] == '/' {
						ApPhpSnprintf(ini_file, 256, "%s%s", debpath, namelist[i].d_name)
					} else {
						ApPhpSnprintf(ini_file, 256, "%s%c%s", debpath, '/', namelist[i].d_name)
					}
					if stat(ini_file, &sb) == 0 {
						if (sb.st_mode & S_IFMT) == S_IFREG {
							var fh zend.ZendFileHandle
							zend.ZendStreamInitFp(&fh, r.Fopen(ini_file, "r"), ini_file)
							if fh.handle.fp != nil {
								if zend.ZendParseIniFile(&fh, 1, 0, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash) == zend.SUCCESS {

									/* Here, add it to the list of ini files read */

									l = int(strlen(ini_file))
									total_l += l + 2
									p = zend._estrndup(ini_file, l)
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
		zend._efree(bufpath)
		if total_l != 0 {
			var php_ini_scanned_files_len int = g.CondF1(PhpIniScannedFiles != nil, func() int { return int(strlen(PhpIniScannedFiles) + 1) }, 0)
			PhpIniScannedFiles = (*byte)(realloc(PhpIniScannedFiles, php_ini_scanned_files_len+total_l+1))
			if php_ini_scanned_files_len == 0 {
				*PhpIniScannedFiles = '0'
			}
			total_l += php_ini_scanned_files_len
			for element = scanned_ini_list.head; element != nil; element = element.next {
				if php_ini_scanned_files_len != 0 {
					strlcat(PhpIniScannedFiles, ",\n", total_l)
				}
				strlcat(PhpIniScannedFiles, *((**byte)(element.data)), total_l)
				strlcat(PhpIniScannedFiles, g.Cond(element.next != nil, ",\n", "\n"), total_l)
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

		ActiveIniHash = nil
		IsSpecialSection = 0
		zend.ZendParseIniString(sapi_module.GetIniEntries(), 1, 0, zend.ZendIniParserCbT(PhpIniParserCb), &ConfigurationHash)
	}
	return zend.SUCCESS
}

/* }}} */

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

/* }}} */

func PhpIniRegisterExtensions() {
	zend.ZendLlistApply(&ExtensionLists.engine, PhpLoadZendExtensionCb)
	zend.ZendLlistApply(&ExtensionLists.functions, PhpLoadPhpExtensionCb)
	zend.ZendLlistDestroy(&ExtensionLists.engine)
	zend.ZendLlistDestroy(&ExtensionLists.functions)
}

/* }}} */

func PhpParseUserIniFile(dirname *byte, ini_filename *byte, target_hash *zend.HashTable) int {
	var sb zend.ZendStatT
	var ini_file []byte
	ApPhpSnprintf(ini_file, 256, "%s%c%s", dirname, '/', ini_filename)
	if stat(ini_file, &sb) == 0 {
		if (sb.st_mode & S_IFMT) == S_IFREG {
			var fh zend.ZendFileHandle
			zend.ZendStreamInitFp(&fh, r.Fopen(ini_file, "r"), ini_file)
			if fh.handle.fp != nil {

				/* Reset active ini section */

				ActiveIniHash = nil
				IsSpecialSection = 0
				if zend.ZendParseIniFile(&fh, 1, 0, zend.ZendIniParserCbT(PhpIniParserCb), target_hash) == zend.SUCCESS {

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

/* }}} */

func PhpIniActivateConfig(source_hash *zend.HashTable, modify_type int, stage int) {
	var str *zend.ZendString
	var data *zend.Zval

	/* Walk through config hash and alter matching ini entries using the values found in the hash */

	for {
		var __ht *zend.HashTable = source_hash
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			str = _p.key
			data = _z
			zend.ZendAlterIniEntryEx(str, data.value.str, modify_type, stage, 0)
		}
		break
	}

	/* Walk through config hash and alter matching ini entries using the values found in the hash */
}

/* }}} */

func PhpIniHasPerDirConfig() int { return HasPerDirConfig }

/* }}} */

func PhpIniActivatePerDirConfig(path *byte, path_len int) {
	var tmp2 *zend.Zval
	var ptr *byte
	if path_len > 256 {
		return
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */

	if HasPerDirConfig != 0 && path != nil && path_len != 0 {
		ptr = path + 1
		for g.Assign(&ptr, strchr(ptr, '/')) != nil {
			*ptr = 0

			/* Search for source array matching the path from configuration_hash */

			if g.Assign(&tmp2, zend.ZendHashStrFind(&ConfigurationHash, path, strlen(path))) != nil {
				PhpIniActivateConfig(tmp2.value.arr, 1<<2, 1<<2)
			}
			*ptr = '/'
			ptr++
		}
	}

	/* Walk through each directory in path and apply any found per-dir-system-configuration from configuration_hash */
}

/* }}} */

func PhpIniHasPerHostConfig() int { return HasPerHostConfig }

/* }}} */

func PhpIniActivatePerHostConfig(host *byte, host_len int) {
	var tmp *zend.Zval
	if HasPerHostConfig != 0 && host != nil && host_len != 0 {

		/* Search for source array matching the host from configuration_hash */

		if g.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, host, host_len)) != nil {
			PhpIniActivateConfig(tmp.value.arr, 1<<2, 1<<2)
		}

		/* Search for source array matching the host from configuration_hash */

	}
}

/* }}} */

func CfgGetEntryEx(name *zend.ZendString) *zend.Zval {
	return zend.ZendHashFind(&ConfigurationHash, name)
}

/* }}} */

func CfgGetEntry(name *byte, name_length int) *zend.Zval {
	return zend.ZendHashStrFind(&ConfigurationHash, name, name_length)
}

/* }}} */

func CfgGetLong(varname *byte, result *zend.ZendLong) int {
	var tmp *zend.Zval
	if g.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = 0
		return zend.FAILURE
	}
	*result = zend.ZvalGetLong(tmp)
	return zend.SUCCESS
}

/* }}} */

func CfgGetDouble(varname *byte, result *float64) int {
	var tmp *zend.Zval
	if g.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = float64(0)
		return zend.FAILURE
	}
	*result = zend.ZvalGetDouble(tmp)
	return zend.SUCCESS
}

/* }}} */

func CfgGetString(varname *byte, result **byte) int {
	var tmp *zend.Zval
	if g.Assign(&tmp, zend.ZendHashStrFind(&ConfigurationHash, varname, strlen(varname))) == nil {
		*result = nil
		return zend.FAILURE
	}
	*result = tmp.value.str.val
	return zend.SUCCESS
}

/* }}} */

func PhpIniGetConfigurationHash() *zend.HashTable { return &ConfigurationHash }
