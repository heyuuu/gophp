// <<generate>>

package zend

import (
	b "sik/builtin"
	r "sik/runtime"
)

func GetIntVal(op *Zval) int {
	switch op.GetType() {
	case IS_LONG:
		return Z_LVAL_P(op)
	case IS_DOUBLE:
		return int(Z_DVAL_P(op))
	case IS_STRING:
		var val int = atoi(Z_STRVAL_P(op))
		ZendStringFree(Z_STR_P(op))
		return val
	default:
		break
	}
}
func ZendIniDoOp(type_ byte, result *Zval, op1 *Zval, op2 *Zval) {
	var i_result int
	var i_op1 int
	var i_op2 int
	var str_len int
	var str_result []byte
	i_op1 = GetIntVal(op1)
	if op2 != nil {
		i_op2 = GetIntVal(op2)
	} else {
		i_op2 = 0
	}
	switch type_ {
	case '|':
		i_result = i_op1 | i_op2
		break
	case '&':
		i_result = i_op1 & i_op2
		break
	case '^':
		i_result = i_op1 ^ i_op2
		break
	case '~':
		i_result = ^i_op1
		break
	case '!':
		i_result = !i_op1
		break
	default:
		i_result = 0
		break
	}
	str_len = sprintf(str_result, "%d", i_result)
	ZVAL_NEW_STR(result, ZendStringInit(str_result, str_len, ZEND_SYSTEM_INI))
}
func ZendIniInitString(result *Zval) {
	if ZEND_SYSTEM_INI != 0 {
		ZVAL_EMPTY_PSTRING(result)
	} else {
		ZVAL_EMPTY_STRING(result)
	}
}
func ZendIniAddString(result *Zval, op1 *Zval, op2 *Zval) {
	var length int
	var op1_len int
	if op1.GetType() != IS_STRING {

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

		if ZEND_SYSTEM_INI != 0 {
			var tmp_str *ZendString
			var str *ZendString = ZvalGetTmpString(op1, &tmp_str)
			ZVAL_PSTRINGL(op1, str.GetVal(), str.GetLen())
			ZendTmpStringRelease(tmp_str)
		} else {
			ZVAL_STR(op1, ZvalGetStringFunc(op1))
		}

		/* ZEND_ASSERT(!Z_REFCOUNTED_P(op1)); */

	}
	op1_len = int(Z_STRLEN_P(op1))
	if op2.GetType() != IS_STRING {
		ConvertToString(op2)
	}
	length = op1_len + int(Z_STRLEN_P(op2))
	ZVAL_NEW_STR(result, ZendStringExtend(Z_STR_P(op1), length, ZEND_SYSTEM_INI))
	memcpy(Z_STRVAL_P(result)+op1_len, Z_STRVAL_P(op2), Z_STRLEN_P(op2)+1)
}
func ZendIniGetConstant(result *Zval, name *Zval) {
	var c *Zval
	var tmp Zval

	/* If name contains ':' it is not a constant. Bug #26893. */

	if !(memchr(Z_STRVAL_P(name), ':', Z_STRLEN_P(name))) && b.Assign(&c, ZendGetConstant(Z_STR_P(name))) != 0 {
		if c.GetType() != IS_STRING {
			ZVAL_COPY_OR_DUP(&tmp, c)
			if Z_OPT_CONSTANT(tmp) {
				ZvalUpdateConstantEx(&tmp, nil)
			}
			ConvertToString(&tmp)
			c = &tmp
		}
		ZVAL_NEW_STR(result, ZendStringInit(Z_STRVAL_P(c), Z_STRLEN_P(c), ZEND_SYSTEM_INI))
		if c == &tmp {
			ZendStringRelease(tmp.GetStr())
		}
		ZendStringFree(Z_STR_P(name))
	} else {
		*result = *name
	}

	/* If name contains ':' it is not a constant. Bug #26893. */
}
func ZendIniGetVar(result *Zval, name *Zval) {
	var curval *Zval
	var envvar *byte

	/* Fetch configuration option value */

	if b.Assign(&curval, ZendGetConfigurationDirective(Z_STR_P(name))) != nil {
		ZVAL_NEW_STR(result, ZendStringInit(Z_STRVAL_P(curval), Z_STRLEN_P(curval), ZEND_SYSTEM_INI))
	} else if b.Assign(&envvar, ZendGetenv(Z_STRVAL_P(name), Z_STRLEN_P(name))) != nil || b.Assign(&envvar, getenv(Z_STRVAL_P(name))) != nil {
		ZVAL_NEW_STR(result, ZendStringInit(envvar, strlen(envvar), ZEND_SYSTEM_INI))
	} else {
		ZendIniInitString(result)
	}

	/* Fetch configuration option value */
}
func IniError(msg *byte) {
	var error_buf *byte
	var error_buf_len int
	var currently_parsed_filename *byte
	currently_parsed_filename = ZendIniScannerGetFilename()
	if currently_parsed_filename != nil {
		error_buf_len = 128 + int(strlen(msg)+int(strlen(currently_parsed_filename)))
		error_buf = (*byte)(Emalloc(error_buf_len))
		sprintf(error_buf, "%s in %s on line %d\n", msg, currently_parsed_filename, ZendIniScannerGetLineno())
	} else {
		error_buf = Estrdup("Invalid configuration directive\n")
	}
	if CompilerGlobals.GetIniParserUnbufferedErrors() != 0 {
		r.Fprintf(stderr, "PHP:  %s", error_buf)
	} else {
		ZendError(E_WARNING, "%s", error_buf)
	}
	Efree(error_buf)
}
func ZendParseIniFile(fh *ZendFileHandle, unbuffered_errors ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CompilerGlobals.SetIniParserParam(&ini_parser_param)
	if ZendIniOpenFileForScanning(fh, scanner_mode) == FAILURE {
		return FAILURE
	}
	CompilerGlobals.SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	ZendFileHandleDtor(fh)
	ShutdownIniScanner()
	if retval == 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ZendParseIniString(str *byte, unbuffered_errors ZendBool, scanner_mode int, ini_parser_cb ZendIniParserCbT, arg any) int {
	var retval int
	var ini_parser_param ZendIniParserParam
	ini_parser_param.SetIniParserCb(ini_parser_cb)
	ini_parser_param.SetArg(arg)
	CompilerGlobals.SetIniParserParam(&ini_parser_param)
	if ZendIniPrepareStringForScanning(str, scanner_mode) == FAILURE {
		return FAILURE
	}
	CompilerGlobals.SetIniParserUnbufferedErrors(unbuffered_errors)
	retval = IniParse()
	ShutdownIniScanner()
	if retval == 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ZvalIniDtor(zv *Zval) {
	if zv.IsType(IS_STRING) {
		ZendStringRelease(Z_STR_P(zv))
	}
}
func YY_(Msgid string) string { return Msgid }
func YYUSE(E __auto__)        { void(E) }
func YYSTACK_BYTES(N __auto__) int {
	return N*(b.SizeOf("yytype_int16")+b.SizeOf("YYSTYPE")) + YYSTACK_GAP_MAXIMUM
}
func YYCOPY(Dst __auto__, Src __auto__, Count __auto__) {
	var yyi YYSIZE_T
	for yyi = 0; yyi < Count; yyi++ {
		Dst[yyi] = Src[yyi]
	}
}
func YYTRANSLATE(YYX __auto__) uint {
	return uint(b.CondF1(YYX <= YYMAXUTOK, func() __auto__ { return yytranslate[YYX] }, YYUNDEFTOK))
}
func IniYypactValueIsDefault(Yystate __auto__) bool     { return !!(Yystate == -25) }
func IniYytableValueIsError(Yytable_value __auto__) int { return 0 }
func YYRECOVERING() bool                                { return !!yyerrstatus }
func YYBACKUP(Token __auto__, Value __auto__) {
	if yychar == YYEMPTY {
		yychar = Token
		yylval = Value
		YYPOPSTACK(yylen)
		yystate = *yyssp
		goto yybackup
	} else {
		yyerror(YY_("syntax error: cannot back up"))
		goto yyerrorlab
	}
}
func IniYystrlen(yystr *byte) YYSIZE_T {
	var yylen YYSIZE_T
	for yylen = 0; yystr[yylen]; yylen++ {
		continue
	}
	return yylen
}
func IniYystpcpy(yydest *byte, yysrc *byte) *byte {
	var yyd *byte = yydest
	var yys *byte = yysrc
	for b.Assign(&(b.PostInc(&(*yyd))), b.PostInc(&(*yys))) != '0' {
		continue
	}
	return yyd - 1
}
func IniYytnamerr(yyres *byte, yystr *byte) YYSIZE_T {
	if (*yystr) == '"' {
		var yyn YYSIZE_T = 0
		var yyp *byte = yystr
		for {
			switch *(b.PreInc(&yyp)) {
			case '\'':

			case ',':
				goto do_not_strip_quotes
			case '\\':
				if (*(b.PreInc(&yyp))) != '\\' {
					goto do_not_strip_quotes
				}
			default:
				if yyres != nil {
					yyres[yyn] = *yyp
				}
				yyn++
				break
			case '"':
				if yyres != nil {
					yyres[yyn] = '0'
				}
				return yyn
			}
		}
	do_not_strip_quotes:
	}
	if yyres == nil {
		return yystrlen(yystr)
	}
	return yystpcpy(yyres, yystr) - yyres
}
func IniYysyntaxError(yymsg_alloc *YYSIZE_T, yymsg **byte, yyssp *yytype_int16, yytoken int) int {
	var yysize0 YYSIZE_T = yytnamerr(YY_NULLPTR, yytname[yytoken])
	var yysize YYSIZE_T = yysize0
	const YYERROR_VERBOSE_ARGS_MAXIMUM = 5

	/* Internationalized format string. */

	var yyformat *byte = YY_NULLPTR

	/* Arguments of yyformat. */

	var yyarg *[]byte

	/* Number of reported tokens (one for the "unexpected", one per
	   "expected"). */

	var yycount int = 0

	/* There are many possibilities here to consider:
	   - If this state is a consistent state with a default action, then
	     the only way this function was invoked is if the default action
	     is an error action.  In that case, don't check for expected
	     tokens because there are none.
	   - The only way there can be no lookahead present (in yychar) is if
	     this state is a consistent state with a default action.  Thus,
	     detecting the absence of a lookahead is sufficient to determine
	     that there is no unexpected or expected token to report.  In that
	     case, just report a simple "syntax error".
	   - Don't assume there isn't a lookahead just because this state is a
	     consistent state with a default action.  There might have been a
	     previous inconsistent state, consistent state with a non-default
	     action, or user semantic action that manipulated yychar.
	   - Of course, the expected token list depends on states to have
	     correct lookahead information, and it depends on the parser not
	     to perform extra reductions after fetching a lookahead from the
	     scanner and before detecting a syntax error.  Thus, state merging
	     (from LALR or IELR) and default reductions corrupt the expected
	     token list.  However, the list is correct for canonical LR with
	     one exception: it will still contain any token that will not be
	     accepted due to an error action in a later state.
	*/

	if yytoken != YYEMPTY {
		var yyn int = yypact[*yyssp]
		yyarg[b.PostInc(&yycount)] = yytname[yytoken]
		if !(yypact_value_is_default(yyn)) {

			/* Start YYX at -YYN if negative to avoid negative indexes in
			   YYCHECK.  In other words, skip the first -YYN actions for
			   this state because they are default actions.  */

			var yyxbegin int = b.Cond(yyn < 0, -yyn, 0)

			/* Stay within bounds of both yycheck and yytname.  */

			var yychecklim int = YYLAST - yyn + 1
			var yyxend int = b.Cond(yychecklim < YYNTOKENS, yychecklim, YYNTOKENS)
			var yyx int
			for yyx = yyxbegin; yyx < yyxend; yyx++ {
				if yycheck[yyx+yyn] == yyx && yyx != YYTERROR && !(yytable_value_is_error(yytable[yyx+yyn])) {
					if yycount == YYERROR_VERBOSE_ARGS_MAXIMUM {
						yycount = 1
						yysize = yysize0
						break
					}
					yyarg[b.PostInc(&yycount)] = yytname[yyx]
					var yysize1 YYSIZE_T = yysize + yytnamerr(YY_NULLPTR, yytname[yyx])
					if !(yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM) {
						return 2
					}
					yysize = yysize1
				}
			}
		}
	}
	switch yycount {
	case 0:
		yyformat = YY_("syntax error")
		break
	case 1:
		yyformat = YY_("syntax error, unexpected %s")
		break
	case 2:
		yyformat = YY_("syntax error, unexpected %s, expecting %s")
		break
	case 3:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s")
		break
	case 4:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s or %s")
		break
	case 5:
		yyformat = YY_("syntax error, unexpected %s, expecting %s or %s or %s or %s")
		break
	}
	var yysize1 YYSIZE_T = yysize + yystrlen(yyformat)
	if !(yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM) {
		return 2
	}
	yysize = yysize1
	if (*yymsg_alloc) < yysize {
		*yymsg_alloc = 2 * yysize
		if !(yysize <= (*yymsg_alloc) && (*yymsg_alloc) <= YYSTACK_ALLOC_MAXIMUM) {
			*yymsg_alloc = YYSTACK_ALLOC_MAXIMUM
		}
		return 1
	}

	/* Avoid sprintf, as that infringes on the user's name space.
	   Don't have undefined behavior even if the translation
	   produced a string with the wrong number of "%s"s.  */

	var yyp *byte = *yymsg
	var yyi int = 0
	for b.Assign(&(*yyp), *yyformat) != '0' {
		if (*yyp) == '%' && yyformat[1] == 's' && yyi < yycount {
			yyp += yytnamerr(yyp, yyarg[b.PostInc(&yyi)])
			yyformat += 2
		} else {
			yyp++
			yyformat++
		}
	}
	return 0
}
func IniYydestruct(yymsg *byte, yytype int, yyvaluep *Zval) {
	YYUSE(yyvaluep)
	if yymsg == nil {
		yymsg = "Deleting"
	}
	switch yytype {
	case 4:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 5:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 6:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 7:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 8:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 9:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 10:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 12:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 14:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 15:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 16:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 47:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 48:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 49:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 50:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 51:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 52:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 53:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 54:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 55:
		ZvalIniDtor(&(*yyvaluep))
		break
	case 56:
		ZvalIniDtor(&(*yyvaluep))
		break
	default:
		break
	}
}
