# 2 程序结构

这一部分是与 Go 语言的程序结构相关的内容。

## 变量声明

Go 语言主要有四种类型的声明语句: `var` , `const` , `type` 和 `func` ，分别对应变量、常量、类型和函数的声明。

一个 Go 语言源程序以包的声明语句开始，说明源文件属于哪一个包，然后是 `import` 语句导入其他的包，然后是包一级的类型、变量、常量、函数的声明。

```go
package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

```

常量 `boilingF` 是包一级的常量，`f` 和 `c` 是在函数中定义的变量。包一级的变量所有的函数均可访问， `main` 中定义的变量为局部变量。

Go 语言变量声明格式：

```go
var 变量名 变量类型 = 表达式
```

变量类型和表达式可以省略一个，若省略类型，则可以由编译器由表达式的类型来推断变量类型。若省略表达式，则可以由编译器自动为变量初始化零值。

数值类型零值为 0 ，字符串零值为空字符串，布尔型变量零值为 `false` ，接口或引用类型(slice, pointer, map, chan, func) 零值为 `nil` 。所以想要创建没有元素的 `slice` 需要使用 `make()` 函数。