module example.com/04-import-local-package

go 1.25

// 使用 replace 指令将远程模块路径映射到本地路径
// 语法：replace 模块名 => 本地路径
replace myproject => ../03-import-local
