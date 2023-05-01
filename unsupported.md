# 不支持的 Ini 配置

- `zend.enable_gc`
- `zend.multibyte=Off`
- `zend.script_encoding`
- `zend.detect_unicode`

# 不支持的函数

- `dl`
- `cli_set_process_title`
- `cli_get_process_title`
- `apache_child_terminate`
- `phpinfo`
- `eval`
- `create_function`

# 部分功能不支持的函数

- `assert(mixed $assertion, Throwable $exception = ?)` 不再支持 `$assertion` 为 string 类型的情况

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

# 不支持的功能

- `ticks` 相关功能，包括 declare(ticks=1) 指令、register_tick_function() 方法、unregister_tick_function() 方法

# 实现细节不同的方法

- 默认随机数直接使用了 golang 的 rand。影响 `lcg_value`、`mt_srand`、`srand`、`mt_rand`、`rand`
- 因为并发性考虑，`putenv` 不再影响进程环境变量