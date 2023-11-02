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
