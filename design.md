# 设计

## 目录结构

- `php` 核心目录，定义基础类型、运行时等
    - `token`  : PHP 词法相关定义
    - `ast`    : PHP 语法树(AST)相关定义
    - `parser` : PHP解析器，将 PHP 源代码转成 AST 语法树
    - `printer`: AST Printer，将 AST 转为人类易读文本
- `compile` PHP 编译相关功能
    - `ir`     : 中间代码(IR)相关定义
    - `render` : 将 IR 生成为 go 代码
- `shim` 模拟新版本 API
    - `builtin` : 模拟新版本内置函数，建议 `import . "gophp/shim/builtin"` 使用
    - `cmp`     : `go1.21` 新包，比较类型
    - `slices`  : `go1.21` 新包，slice 相关泛型方法
    - `maps`    : `go1.21` 新包，map 相关泛型方法