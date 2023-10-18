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
- `get_resources` (注: 在 go 无弱引用情况下，记录全局resource影响自动 GC。在有 Weak Reference 方案后再支持)
- `convert_cyr_string` (注: `本函数已自 PHP 7.4.0 起弃用，自 PHP 8.0.0 起移除。强烈建议不要依赖本函数。`)

# 部分功能不支持的函数

- `assert(mixed $assertion, Throwable $exception = ?)` 不再支持 `$assertion` 为 string 类型的情况 (此功能在 PHP 中本已是废弃功能: `自 PHP 7.2.0 起弃用 string 作为 assertion，自 PHP 8.0.0 起删除。`)

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
- `WeakReference`类

# 实现细节不同的方法

- 默认随机数直接使用了 golang 的 rand。影响 `lcg_value`、`mt_srand`、`srand`、`mt_rand`、`rand`
- 因为并发性考虑，`putenv` 不再影响进程环境变量
- `debug_zval_dump` 输出结果不再有引用计数(`refcount`)相关内容
- Object.handle 使用指针地址代替全局发号器作为对应标识。影响函数 `spl_object_id`/`spl_object_hash` 的值，不影响值的唯一性。