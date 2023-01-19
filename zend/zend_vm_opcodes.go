// <<generate>>

package zend

// Source: <Zend/zend_vm_opcodes.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_VM_OPCODES_H

// #define ZEND_VM_SPEC       1

// #define ZEND_VM_LINES       0

// #define ZEND_VM_KIND_CALL       1

// #define ZEND_VM_KIND_SWITCH       2

// #define ZEND_VM_KIND_GOTO       3

// #define ZEND_VM_KIND_HYBRID       4

/* HYBRID requires support for computed GOTO and global register variables*/

// #define ZEND_VM_KIND       ZEND_VM_KIND_CALL

// #define ZEND_VM_OP_SPEC       0x00000001

// #define ZEND_VM_OP_CONST       0x00000002

// #define ZEND_VM_OP_TMPVAR       0x00000004

// #define ZEND_VM_OP_TMPVARCV       0x00000008

// #define ZEND_VM_OP_MASK       0x000000f0

// #define ZEND_VM_OP_NUM       0x00000010

// #define ZEND_VM_OP_JMP_ADDR       0x00000020

// #define ZEND_VM_OP_TRY_CATCH       0x00000030

// #define ZEND_VM_OP_THIS       0x00000050

// #define ZEND_VM_OP_NEXT       0x00000060

// #define ZEND_VM_OP_CLASS_FETCH       0x00000070

// #define ZEND_VM_OP_CONSTRUCTOR       0x00000080

// #define ZEND_VM_OP_CONST_FETCH       0x00000090

// #define ZEND_VM_OP_CACHE_SLOT       0x000000a0

// #define ZEND_VM_EXT_VAR_FETCH       0x00010000

// #define ZEND_VM_EXT_ISSET       0x00020000

// #define ZEND_VM_EXT_CACHE_SLOT       0x00040000

// #define ZEND_VM_EXT_ARRAY_INIT       0x00080000

// #define ZEND_VM_EXT_REF       0x00100000

// #define ZEND_VM_EXT_FETCH_REF       0x00200000

// #define ZEND_VM_EXT_DIM_OBJ_WRITE       0x00400000

// #define ZEND_VM_EXT_MASK       0x0f000000

// #define ZEND_VM_EXT_NUM       0x01000000

// #define ZEND_VM_EXT_LAST_CATCH       0x02000000

// #define ZEND_VM_EXT_JMP_ADDR       0x03000000

// #define ZEND_VM_EXT_OP       0x04000000

// #define ZEND_VM_EXT_TYPE       0x07000000

// #define ZEND_VM_EXT_EVAL       0x08000000

// #define ZEND_VM_EXT_TYPE_MASK       0x09000000

// #define ZEND_VM_EXT_SRC       0x0b000000

// #define ZEND_VM_NO_CONST_CONST       0x40000000

// #define ZEND_VM_COMMUTATIVE       0x80000000

// #define ZEND_VM_OP1_FLAGS(flags) ( flags & 0xff )

// #define ZEND_VM_OP2_FLAGS(flags) ( ( flags >> 8 ) & 0xff )

// #define ZEND_NOP       0

// #define ZEND_ADD       1

// #define ZEND_SUB       2

// #define ZEND_MUL       3

// #define ZEND_DIV       4

// #define ZEND_MOD       5

// #define ZEND_SL       6

// #define ZEND_SR       7

// #define ZEND_CONCAT       8

// #define ZEND_BW_OR       9

// #define ZEND_BW_AND       10

// #define ZEND_BW_XOR       11

// #define ZEND_POW       12

// #define ZEND_BW_NOT       13

// #define ZEND_BOOL_NOT       14

// #define ZEND_BOOL_XOR       15

// #define ZEND_IS_IDENTICAL       16

// #define ZEND_IS_NOT_IDENTICAL       17

// #define ZEND_IS_EQUAL       18

// #define ZEND_IS_NOT_EQUAL       19

// #define ZEND_IS_SMALLER       20

// #define ZEND_IS_SMALLER_OR_EQUAL       21

// #define ZEND_ASSIGN       22

// #define ZEND_ASSIGN_DIM       23

// #define ZEND_ASSIGN_OBJ       24

// #define ZEND_ASSIGN_STATIC_PROP       25

// #define ZEND_ASSIGN_OP       26

// #define ZEND_ASSIGN_DIM_OP       27

// #define ZEND_ASSIGN_OBJ_OP       28

// #define ZEND_ASSIGN_STATIC_PROP_OP       29

// #define ZEND_ASSIGN_REF       30

// #define ZEND_QM_ASSIGN       31

// #define ZEND_ASSIGN_OBJ_REF       32

// #define ZEND_ASSIGN_STATIC_PROP_REF       33

// #define ZEND_PRE_INC       34

// #define ZEND_PRE_DEC       35

// #define ZEND_POST_INC       36

// #define ZEND_POST_DEC       37

// #define ZEND_PRE_INC_STATIC_PROP       38

// #define ZEND_PRE_DEC_STATIC_PROP       39

// #define ZEND_POST_INC_STATIC_PROP       40

// #define ZEND_POST_DEC_STATIC_PROP       41

// #define ZEND_JMP       42

// #define ZEND_JMPZ       43

// #define ZEND_JMPNZ       44

// #define ZEND_JMPZNZ       45

// #define ZEND_JMPZ_EX       46

// #define ZEND_JMPNZ_EX       47

// #define ZEND_CASE       48

// #define ZEND_CHECK_VAR       49

// #define ZEND_SEND_VAR_NO_REF_EX       50

// #define ZEND_CAST       51

// #define ZEND_BOOL       52

// #define ZEND_FAST_CONCAT       53

// #define ZEND_ROPE_INIT       54

// #define ZEND_ROPE_ADD       55

// #define ZEND_ROPE_END       56

// #define ZEND_BEGIN_SILENCE       57

// #define ZEND_END_SILENCE       58

// #define ZEND_INIT_FCALL_BY_NAME       59

// #define ZEND_DO_FCALL       60

// #define ZEND_INIT_FCALL       61

// #define ZEND_RETURN       62

// #define ZEND_RECV       63

// #define ZEND_RECV_INIT       64

// #define ZEND_SEND_VAL       65

// #define ZEND_SEND_VAR_EX       66

// #define ZEND_SEND_REF       67

// #define ZEND_NEW       68

// #define ZEND_INIT_NS_FCALL_BY_NAME       69

// #define ZEND_FREE       70

// #define ZEND_INIT_ARRAY       71

// #define ZEND_ADD_ARRAY_ELEMENT       72

// #define ZEND_INCLUDE_OR_EVAL       73

// #define ZEND_UNSET_VAR       74

// #define ZEND_UNSET_DIM       75

// #define ZEND_UNSET_OBJ       76

// #define ZEND_FE_RESET_R       77

// #define ZEND_FE_FETCH_R       78

// #define ZEND_EXIT       79

// #define ZEND_FETCH_R       80

// #define ZEND_FETCH_DIM_R       81

// #define ZEND_FETCH_OBJ_R       82

// #define ZEND_FETCH_W       83

// #define ZEND_FETCH_DIM_W       84

// #define ZEND_FETCH_OBJ_W       85

// #define ZEND_FETCH_RW       86

// #define ZEND_FETCH_DIM_RW       87

// #define ZEND_FETCH_OBJ_RW       88

// #define ZEND_FETCH_IS       89

// #define ZEND_FETCH_DIM_IS       90

// #define ZEND_FETCH_OBJ_IS       91

// #define ZEND_FETCH_FUNC_ARG       92

// #define ZEND_FETCH_DIM_FUNC_ARG       93

// #define ZEND_FETCH_OBJ_FUNC_ARG       94

// #define ZEND_FETCH_UNSET       95

// #define ZEND_FETCH_DIM_UNSET       96

// #define ZEND_FETCH_OBJ_UNSET       97

// #define ZEND_FETCH_LIST_R       98

// #define ZEND_FETCH_CONSTANT       99

// #define ZEND_CHECK_FUNC_ARG       100

// #define ZEND_EXT_STMT       101

// #define ZEND_EXT_FCALL_BEGIN       102

// #define ZEND_EXT_FCALL_END       103

// #define ZEND_EXT_NOP       104

// #define ZEND_TICKS       105

// #define ZEND_SEND_VAR_NO_REF       106

// #define ZEND_CATCH       107

// #define ZEND_THROW       108

// #define ZEND_FETCH_CLASS       109

// #define ZEND_CLONE       110

// #define ZEND_RETURN_BY_REF       111

// #define ZEND_INIT_METHOD_CALL       112

// #define ZEND_INIT_STATIC_METHOD_CALL       113

// #define ZEND_ISSET_ISEMPTY_VAR       114

// #define ZEND_ISSET_ISEMPTY_DIM_OBJ       115

// #define ZEND_SEND_VAL_EX       116

// #define ZEND_SEND_VAR       117

// #define ZEND_INIT_USER_CALL       118

// #define ZEND_SEND_ARRAY       119

// #define ZEND_SEND_USER       120

// #define ZEND_STRLEN       121

// #define ZEND_DEFINED       122

// #define ZEND_TYPE_CHECK       123

// #define ZEND_VERIFY_RETURN_TYPE       124

// #define ZEND_FE_RESET_RW       125

// #define ZEND_FE_FETCH_RW       126

// #define ZEND_FE_FREE       127

// #define ZEND_INIT_DYNAMIC_CALL       128

// #define ZEND_DO_ICALL       129

// #define ZEND_DO_UCALL       130

// #define ZEND_DO_FCALL_BY_NAME       131

// #define ZEND_PRE_INC_OBJ       132

// #define ZEND_PRE_DEC_OBJ       133

// #define ZEND_POST_INC_OBJ       134

// #define ZEND_POST_DEC_OBJ       135

// #define ZEND_ECHO       136

// #define ZEND_OP_DATA       137

// #define ZEND_INSTANCEOF       138

// #define ZEND_GENERATOR_CREATE       139

// #define ZEND_MAKE_REF       140

// #define ZEND_DECLARE_FUNCTION       141

// #define ZEND_DECLARE_LAMBDA_FUNCTION       142

// #define ZEND_DECLARE_CONST       143

// #define ZEND_DECLARE_CLASS       144

// #define ZEND_DECLARE_CLASS_DELAYED       145

// #define ZEND_DECLARE_ANON_CLASS       146

// #define ZEND_ADD_ARRAY_UNPACK       147

// #define ZEND_ISSET_ISEMPTY_PROP_OBJ       148

// #define ZEND_HANDLE_EXCEPTION       149

// #define ZEND_USER_OPCODE       150

// #define ZEND_ASSERT_CHECK       151

// #define ZEND_JMP_SET       152

// #define ZEND_UNSET_CV       153

// #define ZEND_ISSET_ISEMPTY_CV       154

// #define ZEND_FETCH_LIST_W       155

// #define ZEND_SEPARATE       156

// #define ZEND_FETCH_CLASS_NAME       157

// #define ZEND_CALL_TRAMPOLINE       158

// #define ZEND_DISCARD_EXCEPTION       159

// #define ZEND_YIELD       160

// #define ZEND_GENERATOR_RETURN       161

// #define ZEND_FAST_CALL       162

// #define ZEND_FAST_RET       163

// #define ZEND_RECV_VARIADIC       164

// #define ZEND_SEND_UNPACK       165

// #define ZEND_YIELD_FROM       166

// #define ZEND_COPY_TMP       167

// #define ZEND_BIND_GLOBAL       168

// #define ZEND_COALESCE       169

// #define ZEND_SPACESHIP       170

// #define ZEND_FUNC_NUM_ARGS       171

// #define ZEND_FUNC_GET_ARGS       172

// #define ZEND_FETCH_STATIC_PROP_R       173

// #define ZEND_FETCH_STATIC_PROP_W       174

// #define ZEND_FETCH_STATIC_PROP_RW       175

// #define ZEND_FETCH_STATIC_PROP_IS       176

// #define ZEND_FETCH_STATIC_PROP_FUNC_ARG       177

// #define ZEND_FETCH_STATIC_PROP_UNSET       178

// #define ZEND_UNSET_STATIC_PROP       179

// #define ZEND_ISSET_ISEMPTY_STATIC_PROP       180

// #define ZEND_FETCH_CLASS_CONSTANT       181

// #define ZEND_BIND_LEXICAL       182

// #define ZEND_BIND_STATIC       183

// #define ZEND_FETCH_THIS       184

// #define ZEND_SEND_FUNC_ARG       185

// #define ZEND_ISSET_ISEMPTY_THIS       186

// #define ZEND_SWITCH_LONG       187

// #define ZEND_SWITCH_STRING       188

// #define ZEND_IN_ARRAY       189

// #define ZEND_COUNT       190

// #define ZEND_GET_CLASS       191

// #define ZEND_GET_CALLED_CLASS       192

// #define ZEND_GET_TYPE       193

// #define ZEND_ARRAY_KEY_EXISTS       194

// #define ZEND_VM_LAST_OPCODE       194

// Source: <Zend/zend_vm_opcodes.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < zend . h >

// # include < zend_vm_opcodes . h >

var ZendVmOpcodesNames []*byte = []*byte{"ZEND_NOP", "ZEND_ADD", "ZEND_SUB", "ZEND_MUL", "ZEND_DIV", "ZEND_MOD", "ZEND_SL", "ZEND_SR", "ZEND_CONCAT", "ZEND_BW_OR", "ZEND_BW_AND", "ZEND_BW_XOR", "ZEND_POW", "ZEND_BW_NOT", "ZEND_BOOL_NOT", "ZEND_BOOL_XOR", "ZEND_IS_IDENTICAL", "ZEND_IS_NOT_IDENTICAL", "ZEND_IS_EQUAL", "ZEND_IS_NOT_EQUAL", "ZEND_IS_SMALLER", "ZEND_IS_SMALLER_OR_EQUAL", "ZEND_ASSIGN", "ZEND_ASSIGN_DIM", "ZEND_ASSIGN_OBJ", "ZEND_ASSIGN_STATIC_PROP", "ZEND_ASSIGN_OP", "ZEND_ASSIGN_DIM_OP", "ZEND_ASSIGN_OBJ_OP", "ZEND_ASSIGN_STATIC_PROP_OP", "ZEND_ASSIGN_REF", "ZEND_QM_ASSIGN", "ZEND_ASSIGN_OBJ_REF", "ZEND_ASSIGN_STATIC_PROP_REF", "ZEND_PRE_INC", "ZEND_PRE_DEC", "ZEND_POST_INC", "ZEND_POST_DEC", "ZEND_PRE_INC_STATIC_PROP", "ZEND_PRE_DEC_STATIC_PROP", "ZEND_POST_INC_STATIC_PROP", "ZEND_POST_DEC_STATIC_PROP", "ZEND_JMP", "ZEND_JMPZ", "ZEND_JMPNZ", "ZEND_JMPZNZ", "ZEND_JMPZ_EX", "ZEND_JMPNZ_EX", "ZEND_CASE", "ZEND_CHECK_VAR", "ZEND_SEND_VAR_NO_REF_EX", "ZEND_CAST", "ZEND_BOOL", "ZEND_FAST_CONCAT", "ZEND_ROPE_INIT", "ZEND_ROPE_ADD", "ZEND_ROPE_END", "ZEND_BEGIN_SILENCE", "ZEND_END_SILENCE", "ZEND_INIT_FCALL_BY_NAME", "ZEND_DO_FCALL", "ZEND_INIT_FCALL", "ZEND_RETURN", "ZEND_RECV", "ZEND_RECV_INIT", "ZEND_SEND_VAL", "ZEND_SEND_VAR_EX", "ZEND_SEND_REF", "ZEND_NEW", "ZEND_INIT_NS_FCALL_BY_NAME", "ZEND_FREE", "ZEND_INIT_ARRAY", "ZEND_ADD_ARRAY_ELEMENT", "ZEND_INCLUDE_OR_EVAL", "ZEND_UNSET_VAR", "ZEND_UNSET_DIM", "ZEND_UNSET_OBJ", "ZEND_FE_RESET_R", "ZEND_FE_FETCH_R", "ZEND_EXIT", "ZEND_FETCH_R", "ZEND_FETCH_DIM_R", "ZEND_FETCH_OBJ_R", "ZEND_FETCH_W", "ZEND_FETCH_DIM_W", "ZEND_FETCH_OBJ_W", "ZEND_FETCH_RW", "ZEND_FETCH_DIM_RW", "ZEND_FETCH_OBJ_RW", "ZEND_FETCH_IS", "ZEND_FETCH_DIM_IS", "ZEND_FETCH_OBJ_IS", "ZEND_FETCH_FUNC_ARG", "ZEND_FETCH_DIM_FUNC_ARG", "ZEND_FETCH_OBJ_FUNC_ARG", "ZEND_FETCH_UNSET", "ZEND_FETCH_DIM_UNSET", "ZEND_FETCH_OBJ_UNSET", "ZEND_FETCH_LIST_R", "ZEND_FETCH_CONSTANT", "ZEND_CHECK_FUNC_ARG", "ZEND_EXT_STMT", "ZEND_EXT_FCALL_BEGIN", "ZEND_EXT_FCALL_END", "ZEND_EXT_NOP", "ZEND_TICKS", "ZEND_SEND_VAR_NO_REF", "ZEND_CATCH", "ZEND_THROW", "ZEND_FETCH_CLASS", "ZEND_CLONE", "ZEND_RETURN_BY_REF", "ZEND_INIT_METHOD_CALL", "ZEND_INIT_STATIC_METHOD_CALL", "ZEND_ISSET_ISEMPTY_VAR", "ZEND_ISSET_ISEMPTY_DIM_OBJ", "ZEND_SEND_VAL_EX", "ZEND_SEND_VAR", "ZEND_INIT_USER_CALL", "ZEND_SEND_ARRAY", "ZEND_SEND_USER", "ZEND_STRLEN", "ZEND_DEFINED", "ZEND_TYPE_CHECK", "ZEND_VERIFY_RETURN_TYPE", "ZEND_FE_RESET_RW", "ZEND_FE_FETCH_RW", "ZEND_FE_FREE", "ZEND_INIT_DYNAMIC_CALL", "ZEND_DO_ICALL", "ZEND_DO_UCALL", "ZEND_DO_FCALL_BY_NAME", "ZEND_PRE_INC_OBJ", "ZEND_PRE_DEC_OBJ", "ZEND_POST_INC_OBJ", "ZEND_POST_DEC_OBJ", "ZEND_ECHO", "ZEND_OP_DATA", "ZEND_INSTANCEOF", "ZEND_GENERATOR_CREATE", "ZEND_MAKE_REF", "ZEND_DECLARE_FUNCTION", "ZEND_DECLARE_LAMBDA_FUNCTION", "ZEND_DECLARE_CONST", "ZEND_DECLARE_CLASS", "ZEND_DECLARE_CLASS_DELAYED", "ZEND_DECLARE_ANON_CLASS", "ZEND_ADD_ARRAY_UNPACK", "ZEND_ISSET_ISEMPTY_PROP_OBJ", "ZEND_HANDLE_EXCEPTION", "ZEND_USER_OPCODE", "ZEND_ASSERT_CHECK", "ZEND_JMP_SET", "ZEND_UNSET_CV", "ZEND_ISSET_ISEMPTY_CV", "ZEND_FETCH_LIST_W", "ZEND_SEPARATE", "ZEND_FETCH_CLASS_NAME", "ZEND_CALL_TRAMPOLINE", "ZEND_DISCARD_EXCEPTION", "ZEND_YIELD", "ZEND_GENERATOR_RETURN", "ZEND_FAST_CALL", "ZEND_FAST_RET", "ZEND_RECV_VARIADIC", "ZEND_SEND_UNPACK", "ZEND_YIELD_FROM", "ZEND_COPY_TMP", "ZEND_BIND_GLOBAL", "ZEND_COALESCE", "ZEND_SPACESHIP", "ZEND_FUNC_NUM_ARGS", "ZEND_FUNC_GET_ARGS", "ZEND_FETCH_STATIC_PROP_R", "ZEND_FETCH_STATIC_PROP_W", "ZEND_FETCH_STATIC_PROP_RW", "ZEND_FETCH_STATIC_PROP_IS", "ZEND_FETCH_STATIC_PROP_FUNC_ARG", "ZEND_FETCH_STATIC_PROP_UNSET", "ZEND_UNSET_STATIC_PROP", "ZEND_ISSET_ISEMPTY_STATIC_PROP", "ZEND_FETCH_CLASS_CONSTANT", "ZEND_BIND_LEXICAL", "ZEND_BIND_STATIC", "ZEND_FETCH_THIS", "ZEND_SEND_FUNC_ARG", "ZEND_ISSET_ISEMPTY_THIS", "ZEND_SWITCH_LONG", "ZEND_SWITCH_STRING", "ZEND_IN_ARRAY", "ZEND_COUNT", "ZEND_GET_CLASS", "ZEND_GET_CALLED_CLASS", "ZEND_GET_TYPE", "ZEND_ARRAY_KEY_EXISTS"}
var ZendVmOpcodesFlags []uint32 = []uint32{0x0, 0xb0b, 0xb0b, 0x80000b0b, 0x707, 0xb0b, 0xb0b, 0xb0b, 0x40000707, 0x80000b0b, 0x80000b0b, 0x80000b0b, 0x707, 0x7, 0x7, 0x80000707, 0x80000303, 0x80000303, 0x80000707, 0x80000707, 0xb0b, 0xb0b, 0x301, 0x6701, 0x40751, 0x40000, 0x4000701, 0x4006701, 0x4000751, 0x4000000, 0xb000101, 0x3, 0xb040751, 0xb040000, 0x1, 0x1, 0x1, 0x1, 0x40000, 0x40000, 0x40000, 0x40000, 0x20, 0x2007, 0x2007, 0x3002007, 0x2007, 0x2007, 0x705, 0x101, 0x1001, 0x7000003, 0x7, 0x707, 0x1000701, 0x1000701, 0x1000701, 0x0, 0x1, 0x1040300, 0x0, 0x1040310, 0x3, 0xa110, 0x40310, 0x1007, 0x1001, 0x1001, 0x100a173, 0x1040300, 0x5, 0x186703, 0x106703, 0x8000007, 0x10107, 0x701, 0x40751, 0x2003, 0x3000001, 0x0, 0x10107, 0x707, 0x40757, 0x10107, 0x6701, 0x640751, 0x10107, 0x6701, 0x40751, 0x10107, 0x707, 0x40757, 0x10107, 0x6703, 0x240753, 0x10107, 0x701, 0x40751, 0x70b, 0x40391, 0x1001, 0x0, 0x0, 0x0, 0x0, 0x1000000, 0x1001, 0x2042003, 0x3, 0x40771, 0x57, 0xb000003, 0x1040757, 0x1048773, 0x30107, 0x20707, 0x1003, 0x1001, 0x1000703, 0x1000000, 0x1003, 0x7, 0x40003, 0x9000007, 0xa103, 0x2003, 0x3000001, 0x5, 0x1000700, 0x0, 0x0, 0x0, 0x40751, 0x40751, 0x40751, 0x40751, 0x7, 0x0, 0x47305, 0x0, 0x101, 0x0, 0x40103, 0x303, 0x3, 0x303, 0x40000, 0x0, 0x60757, 0x0, 0x0, 0x2000, 0x2003, 0x101, 0x20101, 0x701, 0x101, 0x71, 0x0, 0x0, 0xb000303, 0x3, 0x20, 0x3000, 0xa110, 0x0, 0x3, 0x105, 0x40301, 0x2003, 0x707, 0x101, 0x103, 0x47000, 0x647000, 0x47000, 0x47000, 0x247000, 0x47000, 0x40000, 0x67000, 0x40373, 0x100101, 0x100101, 0x101, 0x1001, 0x101, 0x300030b, 0x300030b, 0x1000303, 0x107, 0x107, 0x101, 0x103, 0x707}

func ZendGetOpcodeName(opcode ZendUchar) *byte {
	if opcode > 194 {
		return nil
	}
	return ZendVmOpcodesNames[opcode]
}
func ZendGetOpcodeFlags(opcode ZendUchar) uint32 {
	if opcode > 194 {
		opcode = 0
	}
	return ZendVmOpcodesFlags[opcode]
}
