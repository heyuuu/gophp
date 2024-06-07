# gophp

使用 go 编写的 php 代码解释器+转译器，主要两方面功能:

- php 代码解释器: 执行 php 代码
- php 代码转译器: 转译 php 代码为对应的 go 代码，在保证功能不变的情况下提高运行效率和利用 go 生态

## 版本更新

### 0.2.x 版本: [v0.2.x](https://github.com/heyuuu/gophp/tree/v0.2.x) 

当前最新版本。

纯 go 实现版本，包含 PHP 转译功能和执行功能，包含 PHP 核心标准库(`core`、`stardard`等)，支持通过 go 扩展 PHP。
版本基于 0.1.x 版本的大范围重构，很多设计不兼容。

### 0.1.x 版本: [v0.1.x](https://github.com/heyuuu/gophp/tree/v0.1.x)

纯 go 实现版本，包含 PHP 代码解释和转译功能，包含 PHP 核心标准库(`core`、`stardard`等)。

### 0.0.x 版本(MVP版本): [mvp分支](https://github.com/heyuuu/gophp/tree/mvp)

最小可运行版本，初步实现 PHP 代码解释和转译功能。编译阶段依赖 C-PHP 运行时。