package ast

//go:generate stringer -linecomment -type=AssignOpKind,BinaryOpKind,UnaryOpKind,MagicConstKind,CastKind,IncludeKind -output=enum_string.go

// AssignOpKind
type AssignOpKind uint8

const (
	AssignOpBitwiseAnd AssignOpKind = iota + 1 // &=
	AssignOpBitwiseOr                          // |=
	AssignOpBitwiseXor                         // ^=
	AssignOpCoalesce                           // ??=
	AssignOpConcat                             // .=
	AssignOpDiv                                // /=
	AssignOpMinus                              // -=
	AssignOpMod                                // %=
	AssignOpMul                                // *=
	AssignOpPlus                               // +=
	AssignOpPow                                // **=
	AssignOpShiftLeft                          // <<=
	AssignOpShiftRight                         // >>=
)

// BinaryOpKind
type BinaryOpKind uint8

const (
	BinaryOpBitwiseAnd     BinaryOpKind = iota + 1 // &
	BinaryOpBitwiseOr                              // |
	BinaryOpBitwiseXor                             // ^
	BinaryOpBooleanAnd                             // &&
	BinaryOpBooleanOr                              // ||
	BinaryOpBooleanXor                             // xor
	BinaryOpCoalesce                               // ??
	BinaryOpConcat                                 // .
	BinaryOpDiv                                    // /
	BinaryOpEqual                                  // ==
	BinaryOpGreater                                // >
	BinaryOpGreaterOrEqual                         // >=
	BinaryOpIdentical                              // ===
	BinaryOpMinus                                  // -
	BinaryOpMod                                    // %
	BinaryOpMul                                    // *
	BinaryOpNotEqual                               // !=
	BinaryOpNotIdentical                           // !==
	BinaryOpPlus                                   // +
	BinaryOpPow                                    // **
	BinaryOpShiftLeft                              // <<
	BinaryOpShiftRight                             // >>
	BinaryOpSmaller                                // <
	BinaryOpSmallerOrEqual                         // <=
	BinaryOpSpaceship                              // <=>
)

// UnaryOpKind
type UnaryOpKind uint8

const (
	UnaryOpPlus       UnaryOpKind = iota + 1 // +
	UnaryOpMinus                             // -
	UnaryOpBooleanNot                        // !
	UnaryOpBitwiseNot                        // ~
	UnaryOpPreInc                            // ++
	UnaryOpPreDec                            // --
	UnaryOpPostInc                           // ++
	UnaryOpPostDec                           // --
)

// MagicConstKind
type MagicConstKind uint8

const (
	MagicConstClass     MagicConstKind = iota + 1 // __CLASS__
	MagicConstDir                                 // __DIR__
	MagicConstFile                                // __FILE__
	MagicConstFunction                            // __FUNCTION__
	MagicConstLine                                // __LINE__
	MagicConstMethod                              // __METHOD__
	MagicConstNamespace                           // __NAMESPACE__
	MagicConstTrait                               // __TRAIT__
)

// CastKind
type CastKind uint8

const (
	CastArray  CastKind = iota + 1 // (array)
	CastBool                       // (bool)
	CastDouble                     // (float)
	CastInt                        // (int)
	CastObject                     // (object)
	CastString                     // (string)
	CastUnset                      // (unset)
)

// IncludeKind
type IncludeKind uint8

const (
	KindInclude     IncludeKind = iota + 1 // include
	KindIncludeOnce                        // include_once
	KindRequire                            // require
	KindRequireOnce                        // require_once
)

// NameKind
type NameKind uint8

const (
	NameNormal NameKind = iota + 1
	NameFullyQualified
	NameRelative
)

// UseType for UseStmt
type UseType int

const (
	UseNormal   UseType = 1 // Class or namespace import
	UseFunction UseType = 2 // Function import
	UseConstant UseType = 3 // Constant import
)

// flags
type Flags int

func (f Flags) Is(flags Flags) bool { return f&flags != 0 }

const (
	// 此处不写成 1 << iota 形式，为了表示与 PHP Parser 对齐
	FlagPublic    Flags = 1
	FlagProtected Flags = 2
	FlagPrivate   Flags = 4
	FlagStatic    Flags = 8
	FlagAbstract  Flags = 16
	FlagFinal     Flags = 32
	FlagReadonly  Flags = 64

	VisibilityModifierMask = FlagPublic | FlagProtected | FlagPrivate
)

type CommentType int

const (
	CommentLine CommentType = iota
	CommentDoc
)
