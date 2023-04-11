# 不支持的 Ini 配置

- `zend.enable_gc`
- `zend.multibyte=Off`
- `zend.script_encoding`
- `zend.detect_unicode`

# 不支持的方法

- `dl`
- `cli_set_process_title`
- `cli_get_process_title`
- `apache_child_terminate`
- `phpinfo`

进程控制函数
- `proc_open`
- `proc_close`
- `proc_terminate`
- `proc_get_status`
- `proc_nice`

字符串函数
- `localeconv`
- `setlocale`
- `money_format`


# 实现细节不同的方法

- 默认随机数直接使用了 golang 的 rand。影响 `lcg_value`