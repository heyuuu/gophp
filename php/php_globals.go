package php

import (
	"github.com/heyuuu/gophp/php/types"
)

type LastError struct {
	Type    int
	Message string
	File    string
	Lineno  uint32
}

// PhpGlobals
type PhpGlobals struct {
	httpGlobals             [6]types.Zval
	implicitFlush           bool       `prop:""`
	outputBuffering         int        `prop:""`
	outputHandler           string     `prop:""`
	unserializeCallbackFunc string     `prop:""`
	serializePrecision      int        `prop:""`
	memoryLimit             int        `prop:""`
	maxInputTime            int        `prop:""`
	trackErrors             bool       `prop:""`
	displayErrors           int        `prop:""` // todo maybe = PHP_DISPLAY_ERRORS_STDERR(2) ?
	displayStartupErrors    bool       `prop:""`
	logErrors               bool       `prop:""`
	logErrorsMaxLen         int        `prop:""`
	ignoreRepeatedErrors    bool       `prop:""`
	ignoreRepeatedSource    bool       `prop:""`
	reportMemleaks          bool       `prop:""`
	errorLog                string     `prop:""`
	docRoot                 string     `prop:""`
	userDir                 string     `prop:""`
	includePath             string     `prop:""`
	openBasedir             string     `prop:""`
	extensionDir            string     `prop:""`
	phpBinary               string     `prop:""`
	sysTempDir              string     `prop:""`
	uploadTmpDir            string     `prop:""`
	uploadMaxFilesize       int        `prop:""`
	errorAppendString       string     `prop:""`
	errorPrependString      string     `prop:""`
	autoPrependFile         string     `prop:""`
	autoAppendFile          string     `prop:""`
	inputEncoding           string     `prop:""`
	internalEncoding        string     `prop:""`
	outputEncoding          string     `prop:""`
	variablesOrder          string     `prop:""`
	connectionStatus        int16      `prop:""`
	ignoreUserAbort         bool       `prop:""`
	headerIsBeingSent       uint8      `prop:""`
	exposePhp               bool       `prop:""`
	registerArgcArgv        bool       `prop:""`
	autoGlobalsJit          bool       `prop:""`
	docrefRoot              string     `prop:""`
	docrefExt               string     `prop:""`
	htmlErrors              bool       `prop:""`
	xmlrpcErrors            bool       `prop:""`
	xmlrpcErrorNumber       int        `prop:""`
	modulesActivated        bool       `prop:""`
	fileUploads             bool       `prop:""`
	duringRequestStartup    bool       `prop:""`
	allowUrlFopen           bool       `prop:""`
	enablePostDataReading   bool       `prop:""`
	reportZendDebug         bool       `prop:""`
	lastError               *LastError `get:""`
	disableFunctions        string     `prop:""`
	disableClasses          string     `prop:""`
	allowUrlInclude         bool       `prop:""`
	maxInputNestingLevel    int        `prop:""`
	maxInputVars            int        `prop:""`
	inUserInclude           bool       `prop:""`
	userIniFilename         string     `prop:""`
	userIniCacheTtl         int        `prop:""`
	requestOrder            string     `prop:""`
	mailXHeader             bool       `prop:""`
	mailLog                 string     `prop:""`
	inErrorLog              bool       `prop:""`
	syslogFacility          int        `prop:""`
	syslogIdent             string     `prop:""`
	haveCalledOpenlog       bool       `prop:""`
	syslogFilter            int        `prop:""`

	// 迁移自 SG(sapi_globals_struct) 的配置项
	defaultMimetype string `prop:""` // config
	defaultCharset  string `prop:""` // config
	postMaxSize     int    `prop:""` // config
}

func (pg *PhpGlobals) Init(ctx *Context, base *PhpGlobals) {
	if base != nil {
		*pg = *base
	} else {
		*pg = PhpGlobals{}
	}
	pg.httpGlobals = [6]types.Zval{}
}

// last error
func (pg *PhpGlobals) AddLastError(typ int, message string, file string, lineno uint32) {
	pg.lastError = &LastError{Type: typ, Message: message, File: file, Lineno: lineno}
}
func (pg *PhpGlobals) ClearLastError() { pg.lastError = nil }
