// Code generated by "stringer -type=Token"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Comment-1]
	_ = x[DocComment-2]
	_ = x[Ident-3]
	_ = x[Int-4]
	_ = x[Float-5]
	_ = x[String-6]
	_ = x[Not-7]
	_ = x[Tilde-8]
	_ = x[Inc-9]
	_ = x[Dec-10]
	_ = x[PreInc-11]
	_ = x[PreDec-12]
	_ = x[PostInc-13]
	_ = x[PostDec-14]
	_ = x[binary_begin-15]
	_ = x[Add-16]
	_ = x[Sub-17]
	_ = x[Mul-18]
	_ = x[Div-19]
	_ = x[Mod-20]
	_ = x[Pow-21]
	_ = x[Concat-22]
	_ = x[Coalesce-23]
	_ = x[And-24]
	_ = x[Or-25]
	_ = x[Xor-26]
	_ = x[ShiftLeft-27]
	_ = x[ShiftRight-28]
	_ = x[BooleanAnd-29]
	_ = x[BooleanOr-30]
	_ = x[LogicalAnd-31]
	_ = x[LogicalOr-32]
	_ = x[LogicalXor-33]
	_ = x[Equal-34]
	_ = x[NotEqual-35]
	_ = x[Identical-36]
	_ = x[NotIdentical-37]
	_ = x[Greater-38]
	_ = x[GreaterOrEqual-39]
	_ = x[Smaller-40]
	_ = x[SmallerOrEqual-41]
	_ = x[Spaceship-42]
	_ = x[binary_end-43]
	_ = x[assign_begin-44]
	_ = x[Assign-45]
	_ = x[AddAssign-46]
	_ = x[SubAssign-47]
	_ = x[MulAssign-48]
	_ = x[DivAssign-49]
	_ = x[ModAssign-50]
	_ = x[PowAssign-51]
	_ = x[ConcatAssign-52]
	_ = x[CoalesceAssign-53]
	_ = x[AndAssign-54]
	_ = x[OrAssign-55]
	_ = x[XorAssign-56]
	_ = x[ShiftLeftAssign-57]
	_ = x[ShiftRightAssign-58]
	_ = x[assign_end-59]
	_ = x[cast_begin-60]
	_ = x[BoolCast-61]
	_ = x[IntCast-62]
	_ = x[DoubleCast-63]
	_ = x[StringCast-64]
	_ = x[ArrayCast-65]
	_ = x[ObjectCast-66]
	_ = x[UnsetCast-67]
	_ = x[cast_end-68]
	_ = x[magic_const_begin-69]
	_ = x[DirConst-70]
	_ = x[FileConst-71]
	_ = x[LineConst-72]
	_ = x[NamespaceConst-73]
	_ = x[FunctionConst-74]
	_ = x[ClassConst-75]
	_ = x[MethodConst-76]
	_ = x[TraitConst-77]
	_ = x[magic_const_end-78]
	_ = x[internal_call_begin-79]
	_ = x[Isset-80]
	_ = x[Empty-81]
	_ = x[Include-82]
	_ = x[IncludeOnce-83]
	_ = x[Require-84]
	_ = x[RequireOnce-85]
	_ = x[Eval-86]
	_ = x[internal_call_end-87]
	_ = x[NAME_FULLY_QUALIFIED-263]
	_ = x[NAME_RELATIVE-264]
	_ = x[NAME_QUALIFIED-265]
	_ = x[VARIABLE-266]
	_ = x[INLINE_HTML-267]
	_ = x[ENCAPSED_AND_WHITESPACE-268]
	_ = x[STRING_VARNAME-270]
	_ = x[NUM_STRING-271]
	_ = x[PRINT-280]
	_ = x[YIELD-281]
	_ = x[YIELD_FROM-282]
	_ = x[INSTANCEOF-283]
	_ = x[NEW-284]
	_ = x[CLONE-285]
	_ = x[EXIT-286]
	_ = x[IF-287]
	_ = x[ELSEIF-288]
	_ = x[ELSE-289]
	_ = x[ENDIF-290]
	_ = x[ECHO-291]
	_ = x[DO-292]
	_ = x[WHILE-293]
	_ = x[ENDWHILE-294]
	_ = x[FOR-295]
	_ = x[ENDFOR-296]
	_ = x[FOREACH-297]
	_ = x[ENDFOREACH-298]
	_ = x[DECLARE-299]
	_ = x[ENDDECLARE-300]
	_ = x[AS-301]
	_ = x[SWITCH-302]
	_ = x[ENDSWITCH-303]
	_ = x[CASE-304]
	_ = x[DEFAULT-305]
	_ = x[MATCH-306]
	_ = x[BREAK-307]
	_ = x[CONTINUE-308]
	_ = x[GOTO-309]
	_ = x[FUNCTION-310]
	_ = x[FN-311]
	_ = x[CONST-312]
	_ = x[RETURN-313]
	_ = x[TRY-314]
	_ = x[CATCH-315]
	_ = x[FINALLY-316]
	_ = x[THROW-317]
	_ = x[USE-318]
	_ = x[INSTEADOF-319]
	_ = x[GLOBAL-320]
	_ = x[STATIC-321]
	_ = x[ABSTRACT-322]
	_ = x[FINAL-323]
	_ = x[PRIVATE-324]
	_ = x[PROTECTED-325]
	_ = x[PUBLIC-326]
	_ = x[READONLY-327]
	_ = x[VAR-328]
	_ = x[UNSET-329]
	_ = x[HALT_COMPILER-332]
	_ = x[CLASS-333]
	_ = x[TRAIT-334]
	_ = x[INTERFACE-335]
	_ = x[ENUM-336]
	_ = x[EXTENDS-337]
	_ = x[IMPLEMENTS-338]
	_ = x[NAMESPACE-339]
	_ = x[LIST-340]
	_ = x[ARRAY-341]
	_ = x[CALLABLE-342]
	_ = x[ATTRIBUTE-351]
	_ = x[OBJECT_OPERATOR-384]
	_ = x[NULLSAFE_OBJECT_OPERATOR-385]
	_ = x[DOUBLE_ARROW-386]
	_ = x[COMMENT-387]
	_ = x[DOC_COMMENT-388]
	_ = x[OPEN_TAG-389]
	_ = x[OPEN_TAG_WITH_ECHO-390]
	_ = x[CLOSE_TAG-391]
	_ = x[WHITESPACE-392]
	_ = x[START_HEREDOC-393]
	_ = x[END_HEREDOC-394]
	_ = x[DOLLAR_OPEN_CURLY_BRACES-395]
	_ = x[CURLY_OPEN-396]
	_ = x[PAAMAYIM_NEKUDOTAYIM-397]
	_ = x[DOUBLE_COLON-397]
	_ = x[NS_SEPARATOR-398]
	_ = x[ELLIPSIS-399]
	_ = x[AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG-403]
	_ = x[AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG-404]
	_ = x[BAD_CHARACTER-405]
}

const (
	_Token_name_0 = "CommentDocCommentIdentIntFloatStringNotTildeIncDecPreIncPreDecPostIncPostDecbinary_beginAddSubMulDivModPowConcatCoalesceAndOrXorShiftLeftShiftRightBooleanAndBooleanOrLogicalAndLogicalOrLogicalXorEqualNotEqualIdenticalNotIdenticalGreaterGreaterOrEqualSmallerSmallerOrEqualSpaceshipbinary_endassign_beginAssignAddAssignSubAssignMulAssignDivAssignModAssignPowAssignConcatAssignCoalesceAssignAndAssignOrAssignXorAssignShiftLeftAssignShiftRightAssignassign_endcast_beginBoolCastIntCastDoubleCastStringCastArrayCastObjectCastUnsetCastcast_endmagic_const_beginDirConstFileConstLineConstNamespaceConstFunctionConstClassConstMethodConstTraitConstmagic_const_endinternal_call_beginIssetEmptyIncludeIncludeOnceRequireRequireOnceEvalinternal_call_end"
	_Token_name_1 = "NAME_FULLY_QUALIFIEDNAME_RELATIVENAME_QUALIFIEDVARIABLEINLINE_HTMLENCAPSED_AND_WHITESPACE"
	_Token_name_2 = "STRING_VARNAMENUM_STRING"
	_Token_name_3 = "PRINTYIELDYIELD_FROMINSTANCEOFNEWCLONEEXITIFELSEIFELSEENDIFECHODOWHILEENDWHILEFORENDFORFOREACHENDFOREACHDECLAREENDDECLAREASSWITCHENDSWITCHCASEDEFAULTMATCHBREAKCONTINUEGOTOFUNCTIONFNCONSTRETURNTRYCATCHFINALLYTHROWUSEINSTEADOFGLOBALSTATICABSTRACTFINALPRIVATEPROTECTEDPUBLICREADONLYVARUNSET"
	_Token_name_4 = "HALT_COMPILERCLASSTRAITINTERFACEENUMEXTENDSIMPLEMENTSNAMESPACELISTARRAYCALLABLE"
	_Token_name_5 = "ATTRIBUTE"
	_Token_name_6 = "OBJECT_OPERATORNULLSAFE_OBJECT_OPERATORDOUBLE_ARROWCOMMENTDOC_COMMENTOPEN_TAGOPEN_TAG_WITH_ECHOCLOSE_TAGWHITESPACESTART_HEREDOCEND_HEREDOCDOLLAR_OPEN_CURLY_BRACESCURLY_OPENPAAMAYIM_NEKUDOTAYIMNS_SEPARATORELLIPSIS"
	_Token_name_7 = "AMPERSAND_FOLLOWED_BY_VAR_OR_VARARGAMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARGBAD_CHARACTER"
)

var (
	_Token_index_0 = [...]uint16{0, 7, 17, 22, 25, 30, 36, 39, 44, 47, 50, 56, 62, 69, 76, 88, 91, 94, 97, 100, 103, 106, 112, 120, 123, 125, 128, 137, 147, 157, 166, 176, 185, 195, 200, 208, 217, 229, 236, 250, 257, 271, 280, 290, 302, 308, 317, 326, 335, 344, 353, 362, 374, 388, 397, 405, 414, 429, 445, 455, 465, 473, 480, 490, 500, 509, 519, 528, 536, 553, 561, 570, 579, 593, 606, 616, 627, 637, 652, 671, 676, 681, 688, 699, 706, 717, 721, 738}
	_Token_index_1 = [...]uint8{0, 20, 33, 47, 55, 66, 89}
	_Token_index_2 = [...]uint8{0, 14, 24}
	_Token_index_3 = [...]uint16{0, 5, 10, 20, 30, 33, 38, 42, 44, 50, 54, 59, 63, 65, 70, 78, 81, 87, 94, 104, 111, 121, 123, 129, 138, 142, 149, 154, 159, 167, 171, 179, 181, 186, 192, 195, 200, 207, 212, 215, 224, 230, 236, 244, 249, 256, 265, 271, 279, 282, 287}
	_Token_index_4 = [...]uint8{0, 13, 18, 23, 32, 36, 43, 53, 62, 66, 71, 79}
	_Token_index_6 = [...]uint8{0, 15, 39, 51, 58, 69, 77, 95, 104, 114, 127, 138, 162, 172, 192, 204, 212}
	_Token_index_7 = [...]uint8{0, 35, 74, 87}
)

func (i Token) String() string {
	switch {
	case 1 <= i && i <= 87:
		i -= 1
		return _Token_name_0[_Token_index_0[i]:_Token_index_0[i+1]]
	case 263 <= i && i <= 268:
		i -= 263
		return _Token_name_1[_Token_index_1[i]:_Token_index_1[i+1]]
	case 270 <= i && i <= 271:
		i -= 270
		return _Token_name_2[_Token_index_2[i]:_Token_index_2[i+1]]
	case 280 <= i && i <= 329:
		i -= 280
		return _Token_name_3[_Token_index_3[i]:_Token_index_3[i+1]]
	case 332 <= i && i <= 342:
		i -= 332
		return _Token_name_4[_Token_index_4[i]:_Token_index_4[i+1]]
	case i == 351:
		return _Token_name_5
	case 384 <= i && i <= 399:
		i -= 384
		return _Token_name_6[_Token_index_6[i]:_Token_index_6[i+1]]
	case 403 <= i && i <= 405:
		i -= 403
		return _Token_name_7[_Token_index_7[i]:_Token_index_7[i+1]]
	default:
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}