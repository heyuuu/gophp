package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func IniParse(ctx *Context, str string, mode IniScanMode, iniParseCb IniScanCallback) bool {
	if mode != IniScanNormal && mode != IniScanRaw && mode != IniScanTyped {
		Error(ctx, perr.E_WARNING, "Invalid scanner mode")
		return false
	}

	err := IniScan(str, iniParseCb)
	if err != nil {
		return false
	}
	return true
}

func IniParseString(ctx *Context, str string, mode IniScanMode, iniParserCb IniScanCallback) bool {
	return IniParse(ctx, str, mode, iniParserCb)
}

func IniParseFile(ctx *Context, fh *FileHandle, scannerMode IniScanMode, iniParserCb IniScanCallback) bool {
	str, ok := fh.ReadAll()
	if !ok {
		return false
	}

	return IniParse(ctx, str, scannerMode, iniParserCb)
}

func TryParseIniOffsetKey(key string) (baseKey string, offset string, ok bool) {
	if idx := strings.IndexByte(key, '['); idx >= 0 && key[len(key)-1] == ']' {
		baseKey = strings.TrimSpace(key[:idx])
		offset = strings.TrimSpace(key[idx+1 : len(key)-1])
		return baseKey, offset, true
	}
	return key, "", false
}

// --- types

// IniGlobalParser 原 php_ini_parser_cb 的替代
type IniGlobalParser struct {
	ig               *IniGlobals
	targetHash       *types.Array
	activeHash       *types.Array
	isSpecialSection bool
}

func NewDefaultIniParserCb(ctx *Context, targetHash *types.Array) *IniGlobalParser {
	ig := ctx.INI()
	return &IniGlobalParser{
		ig:               ig,
		targetHash:       targetHash,
		activeHash:       targetHash,
		isSpecialSection: false,
	}
}

func (cb *IniGlobalParser) markHasPreDirConfig() {
	cb.ig.hasPerDirConfig = true
}
func (cb *IniGlobalParser) markHasPreHostConfig() {
	cb.ig.hasPerHostConfig = true
}
func (cb *IniGlobalParser) addPhpExtension(ext string) {
	cb.ig.phpExtensions = append(cb.ig.phpExtensions, ext)
}
func (cb *IniGlobalParser) addZendExtension(ext string) {
	cb.ig.zendExtensions = append(cb.ig.zendExtensions, ext)
}

func (cb *IniGlobalParser) Comment(comment string) {}

func (cb *IniGlobalParser) SectionStart(section string) {
	lcSection := ascii.StrToLower(section)
	var key string
	if strings.HasPrefix(lcSection, "path") {
		key = section[4:]
		cb.isSpecialSection = true
		cb.markHasPreDirConfig()
	} else if strings.HasPrefix(lcSection, "host") {
		key = section[4:]
		cb.isSpecialSection = true
		cb.markHasPreHostConfig()
	} else {
		cb.isSpecialSection = false
	}
	if key != "" {
		/* Strip any trailing slashes */
		key = strings.TrimRight(key, "/\\")
		/* Strip any leading whitespace and '=' */
		key = strings.TrimLeft(key, "= \t")

		entry := cb.targetHash.KeyFind(key)
		if !entry.IsArray() {
			entry = types.ZvalArrayInit()
			cb.targetHash.KeyUpdate(key, entry)
		}
		cb.activeHash = entry.Array()
	}
}

func (cb *IniGlobalParser) Pair(section string, key string, value string) {
	lcKey := ascii.StrToLower(key)
	if !cb.isSpecialSection && lcKey == "extension" {
		cb.addPhpExtension(value)
	} else if !cb.isSpecialSection && lcKey == "zend_extension" {
		cb.addZendExtension(value)
	} else {
		cb.arrayUpdate(cb.activeHash, key, value)
	}
}

func (cb *IniGlobalParser) arrayUpdate(baseArr *types.Array, key string, value string) {
	// key with offset. e.g. `key[]=value` or `key[name]=value`
	if baseKey, offset, ok := TryParseIniOffsetKey(key); ok {
		arr := baseArr.KeyFind(baseKey)
		if !arr.IsArray() {
			baseArr.KeyUpdate(baseKey, types.ZvalArrayInit())
			arr = baseArr.KeyFind(baseKey)
		}
		if offset == "" {
			arr.Array().Append(String(value))
		} else {
			arr.Array().KeyUpdate(offset, String(value))
		}
		return
	}

	// basic k-v pair
	baseArr.KeyUpdate(key, String(value))
}
