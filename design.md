# 设计

## 目录结构

- `php` 核心目录，定义基础类型、运行时等
  - `token`  : PHP 词法相关定义
  - `ast`    : PHP 语法树(AST)相关定义
  - `ir`     : 中间代码(IR)相关定义
  - `parser` : PHP解析器，将 PHP 源代码转成 AST 语法树
  - `printer`: AST Printer，将 AST 转为人类易读文本

