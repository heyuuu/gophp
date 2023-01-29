// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func LangYypactValueIsDefault(Yystate __auto__) bool      { return !!(Yystate == -753) }
func LangYytableValueIsError(Yytable_value __auto__) bool { return !!(Yytable_value == -477) }
func LangYystrlen(yystr *byte) YYSIZE_T {
	var yylen YYSIZE_T
	for yylen = 0; yystr[yylen]; yylen++ {
		continue
	}
	return yylen
}
func LangYystpcpy(yydest *byte, yysrc *byte) *byte {
	var yyd *byte = yydest
	var yys *byte = yysrc
	for b.Assign(&(b.PostInc(&(*yyd))), b.PostInc(&(*yys))) != '0' {
		continue
	}
	return yyd - 1
}
func LangYysyntaxError(yymsg_alloc *YYSIZE_T, yymsg **byte, yyssp *yytype_int16, yytoken int) int {
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
func LangYydestruct(yymsg *byte, yytype int, yyvaluep *ZendParserStackElem) {
	YYUSE(yyvaluep)
	if yymsg == nil {
		yymsg = "Deleting"
	}
	switch yytype {
	case 71:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 72:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 73:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 74:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 75:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 76:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 77:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 78:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 79:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 174:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 175:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 176:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 177:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 178:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 182:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 183:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 185:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 186:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 187:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 188:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 189:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 190:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 191:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 192:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 193:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 194:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 196:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 197:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 198:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 199:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 200:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 201:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 204:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 209:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 211:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 213:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 214:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 215:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 216:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 217:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 218:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 219:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 220:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 221:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 223:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 224:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 225:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 226:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 227:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 228:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 229:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 230:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 231:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 232:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 233:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 234:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 235:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 236:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 237:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 238:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 239:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 240:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 241:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 242:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 243:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 244:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 245:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 246:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 247:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 248:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 249:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 250:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 251:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 252:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 257:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 258:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 259:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 260:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 261:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 262:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 263:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 264:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 265:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 266:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 268:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 269:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 270:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 273:
		if yyvaluep.GetStr() != nil {
			ZendStringReleaseEx(yyvaluep.GetStr(), 0)
		}
		break
	case 277:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 278:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 279:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 280:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 281:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 282:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 283:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 284:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 285:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 286:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 287:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 288:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 289:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 290:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 291:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 292:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 293:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 294:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 295:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 296:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 297:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 298:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 299:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 300:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 301:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 302:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 303:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 304:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 305:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 306:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 307:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 308:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	case 309:
		ZendAstDestroy(yyvaluep.GetAst())
		break
	default:
		break
	}
}
func ZendYytnamerr(yyres *byte, yystr *byte) YYSIZE_T {
	/* CG(parse_error) states:
	 * 0 => yyres = NULL, yystr is the unexpected token
	 * 1 => yyres = NULL, yystr is one of the expected tokens
	 * 2 => yyres != NULL, yystr is the unexpected token
	 * 3 => yyres != NULL, yystr is one of the expected tokens
	 */

	if yyres != nil && __CG().GetParseError() < 2 {
		__CG().SetParseError(2)
	}
	if __CG().GetParseError()%2 == 0 {

		/* The unexpected token */

		var buffer []byte
		var end *uint8
		var str *uint8
		var tok1 *uint8 = nil
		var tok2 *uint8 = nil
		var len_ uint = 0
		var toklen uint = 0
		var yystr_len uint
		__CG().GetParseError()++
		if __INI_SCNG().GetYyText()[0] == 0 && __INI_SCNG().GetYyLeng() == 1 && strcmp(yystr, "\"end of file\"") == 0 {
			if yyres != nil {
				yystpcpy(yyres, "end of file")
			}
			return b.SizeOf("\"end of file\"") - 1
		}
		str = __INI_SCNG().GetYyText()
		end = memchr(str, '\n', __INI_SCNG().GetYyLeng())
		yystr_len = uint(yystrlen(yystr))
		if b.Assign(&tok1, memchr(yystr, '(', yystr_len)) != nil && b.Assign(&tok2, ZendMemrchr(yystr, ')', yystr_len)) != nil {
			toklen = tok2 - tok1 + 1
		} else {
			tok2 = nil
			tok1 = tok2
			toklen = 0
		}
		if end == nil {
			if __INI_SCNG().GetYyLeng() > 30 {
				len_ = 30
			} else {
				len_ = __INI_SCNG().GetYyLeng()
			}
		} else {
			if end-str > 30 {
				len_ = 30
			} else {
				len_ = end - str
			}
		}
		if yyres != nil {
			if toklen != 0 {
				core.Snprintf(buffer, b.SizeOf("buffer"), "'%.*s' %.*s", len_, str, toklen, tok1)
			} else {
				core.Snprintf(buffer, b.SizeOf("buffer"), "'%.*s'", len_, str)
			}
			yystpcpy(yyres, buffer)
		}
		return len_ + b.Cond(toklen != 0, toklen+1, 0) + 2
	}

	/* One of the expected tokens */

	if yyres == nil {
		return yystrlen(yystr) - b.Cond((*yystr) == '"', 2, 0)
	}
	if (*yystr) == '"' {
		var yyn YYSIZE_T = 0
		var yyp *byte = yystr
		for ; (*(b.PreInc(&yyp))) != '"'; yyn++ {
			yyres[yyn] = *yyp
		}
		yyres[yyn] = '0'
		return yyn
	}
	yystpcpy(yyres, yystr)
	return strlen(yystr)
}
