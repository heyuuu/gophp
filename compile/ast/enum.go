package ast

//go:generate stringer -linecomment -type=AssignOpKind,BinaryOpKind,UnaryOpKind,MagicConstKind,CastKind,InternalCallOp -output=enum_string.go

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
	BinaryOpCoalesce                               // ??
	BinaryOpConcat                                 // .
	BinaryOpDiv                                    // /
	BinaryOpEqual                                  // ==
	BinaryOpGreater                                // >
	BinaryOpGreaterOrEqual                         // >=
	BinaryOpIdentical                              // ===
	BinaryOpLogicalAnd                             // and
	BinaryOpLogicalOr                              // or
	BinaryOpLogicalXor                             // xor
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

// InternalCallOp
type InternalCallOp uint8

const (
	ICallIsset       InternalCallOp = iota + 1 // isset
	ICallEmpty                                 // empty
	ICallInclude                               // include
	ICallIncludeOnce                           // include_once
	ICallRequire                               // require
	ICallRequireOnce                           // require_once
	ICallEval                                  // eval
)

// NameKind
type NameKind uint8

const (
	NameNormal NameKind = iota + 1
	NameFullyQualified
	NameRelative
)
