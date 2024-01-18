# ZPP (zend parameters parsing)

## 参数类型

| Spec | FAST_ZPP Micro                    | Parameters Type                                    | Zif Type            |
|:----:|-----------------------------------|----------------------------------------------------|---------------------|
|      | Z_PARAM_OPTIONAL                  |                                                    | zpp.Opt             |
|  a   | Z_PARAM_ARRAY(dest)               | dest - zval                                        | zpp.Array           |
|  A   | Z_PARAM_ARRAY_OR_OBJECT(dest)     | dest - zval                                        | zpp.ArrayOrObject   |
|  b   | Z_PARAM_BOOL(dest)                | dest - zend_bool                                   | bool (bool类型)       |
|  C   | Z_PARAM_CLASS(dest)               | dest - zend_class_entry                            | zpp.ClassEntry      |
|  d   | Z_PARAM_DOUBLE(dest)              | dest - double                                      | float64             |
|  f   | Z_PARAM_FUNC(fci, fcc)            | fci - zend_fcall_info, fcc - zend_fcall_info_cache |                     |
|  h   | Z_PARAM_ARRAY_HT(dest)            | dest - HashTable                                   | *types.Array        |
|  H   | Z_PARAM_ARRAY_OR_OBJECT_HT(dest)  | dest - HashTable                                   | zpp.ArrayOrObjectHt |
|  l   | Z_PARAM_LONG(dest)                | dest - long                                        | int                 |
|  L   | Z_PARAM_STRICT_LONG(dest)         | dest - long                                        | zpp.StrictLong      |
|  o   | Z_PARAM_OBJECT(dest)              | dest - zval                                        | zpp.Object          |
|  O   | Z_PARAM_OBJECT_OF_CLASS(dest, ce) | dest - zval                                        |                     |
|  p   | Z_PARAM_PATH(dest, dest_len)      | dest - char*, dest_len - int                       | zpp.Path (string类型) |
|  P   | Z_PARAM_PATH_STR(dest)            | dest - zend_string                                 |                     |
|  r   | Z_PARAM_RESOURCE(dest)            | dest - zval                                        |                     |
|  s   | Z_PARAM_STRING(dest, dest_len)    | dest - char*, dest_len - int                       | string (string 类型)  |
|  S   | Z_PARAM_STR(dest)                 | dest - zend_string                                 | *types.String       |
|  z   | Z_PARAM_ZVAL(dest)                | dest - zval                                        | *types.Zval         |
|      | Z_PARAM_ZVAL_DEREF(dest)          | dest - zval                                        | zpp.ZvalDeref       |
|  +   | Z_PARAM_VARIADIC('+', dest, num)  | dest - zval*, num int                              | []*Zval             |
|      | Z_PARAM_VARIADIC('*', dest, num)  | dest - zval*, num int                              | []*Zval             |

## Zif 定义

通过将 Zif 定义转化为 ZifHandler 等相关代码，除了上表中对应的类型外，还有几个特殊类型

- `bool`: 相当于通过 Z_PARAM_BOOL 获取后会转为 bool 类型
- `string`: 相当于通过 Z_PARAM_STRING 获取后将参数转为 string 类型
- `zpp.Path`: 相当于通过 Z_PARAM_PATH 获取后将参数转为 string 类型
- `zpp.Ex`: 当前EX对象，用于直接依赖 executeData 的函数，不直接使用 *zend.ZendExecuteData 是为了和原始写法区分
- `zpp.Ret`: 返回值，用于需要直接操作返回值的情况
