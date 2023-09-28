package ast

// AssignOpKind
type AssignOpKind uint8

const (
	AssignOpAssign AssignOpKind = iota + 1
	AssignOpBitwiseAnd
	AssignOpBitwiseOr
	AssignOpBitwiseXor
	AssignOpCoalesce
	AssignOpConcat
	AssignOpDiv
	AssignOpMinus
	AssignOpMod
	AssignOpMul
	AssignOpPlus
	AssignOpPow
	AssignOpShiftLeft
	AssignOpShiftRight
)

// BinaryOpKind
type BinaryOpKind uint8

const (
	BinaryOpBitwiseAnd BinaryOpKind = iota + 1
	BinaryOpBitwiseOr
	BinaryOpBitwiseXor
	BinaryOpBooleanAnd
	BinaryOpBooleanOr
	BinaryOpCoalesce
	BinaryOpConcat
	BinaryOpDiv
	BinaryOpEqual
	BinaryOpGreater
	BinaryOpGreaterOrEqual
	BinaryOpIdentical
	BinaryOpLogicalAnd
	BinaryOpLogicalOr
	BinaryOpLogicalXor
	BinaryOpMinus
	BinaryOpMod
	BinaryOpMul
	BinaryOpNotEqual
	BinaryOpNotIdentical
	BinaryOpPlus
	BinaryOpPow
	BinaryOpShiftLeft
	BinaryOpShiftRight
	BinaryOpSmaller
	BinaryOpSmallerOrEqual
	BinaryOpSpaceship
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
	MagicConstClass MagicConstKind = iota + 1
	MagicConstDir
	MagicConstFile
	MagicConstFunction
	MagicConstLine
	MagicConstMethod
	MagicConstNamespace
	MagicConstTrait
)

// CastKind
type CastKind uint8

const (
	CastArray CastKind = iota + 1
	CastBool
	CastDouble
	CastInt
	CastObject
	CastString
	CastUnset
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
