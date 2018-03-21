## Go 基础语法

[TOC]

### Go 语言程序格式

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello World!\n")
}
```

以上为一个最简单的 Go 语言程序。

Go 语言程序由一个一个的包组成，程序第一行就定义了一个包：`package main`。

随后引入模块 `fmt` ，该模块包含输出函数。

最后是程序入口主函数 `func main`。

### 变量定义

Go 语言变量定义格式：`var 变量名 变量类型`。

Go 语言中变量名在前，类型名在后。

在函数中可以简写为：`变量名 变量类型:= 值`。

在直接赋值的情况下可以省略变量类型，由编译器推断变量类型：

``` go
a, b, c, d := 1, 2, true, "def"
```



在函数外定义变量只可以使用 `var` 但可以使用括号同时定义多个变量：

``` go
var (
  a = 1
  b = 2
  c = 3
)
```

Go 语言内置复数类型。

Go 语言不支持隐式类型转换，只支持强制类型转换：`类型(变量)`。

### 定义常量

Go 语言定义常量格式：`const 变量名` 一般在 Go 语言中常量的变量名不大写，因为首字母大写在 Go 语言中具有特殊的作用。

常量若不指定类型，默认保存为字符型，在需要使用时可以自动转换类型。

如：

``` go
const a, b = 3, 4
var c = int(math.Sqrt(a * a + b * b))
```

Go 语言可以用一组常量来实现枚举类型：

```go
const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
```

实现枚举类型时可以使用特殊的自增运算符来简化：

```go
const (
		cpp = iota
		java
		python
		golang
	)
```

`iota` 为自增运算符。运行后结果为：`0 1 2 3`。

### 选择结构

Go 语言主要有两种选择结构 `if` 和 `switch`。

`if` 用法：

```go
contents, err := ioutil.ReadFile(filename)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Printf("%s\n", contents)
}
```

`if` 的条件不需要加括号。

在条件之前可以定义变量，如：

```go
if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
```

定义完变量后加一个分号，即可开始写条件，`if` 中定义的变量作用域仅在 `if` 代码块当中。

Go 语言的 `switch` 语句可以加条件也可以不加，加条件和 C 语言类似，不加条件可以在 `case` 后加条件。

```go
switch {
	case score < 60:
		return "F"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
```

注意：Go 语言的 `switch` 自带 `break` ，不需要再写。

### 循环结构

Go 语言的循环结构只有 `for` 一种，且条件可省略，条件不加括号。

```go
result := 0
for i := 1; i < 101; i++ {
	result += i
}
```

忽略起始条件与自增：

```go
for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

什么条件也不写为死循环：

```go
for {
	fmt.Println("abc")
}
```

### 函数

Go 语言的函数定义格式：

```go
func 函数名(参数 参数类型) 返回值类型 {
    
}
```

函数定义与变量定义相似，名在前，类型在后。

```go
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
```

Go 语言函数允许多返回值：

```go
func div(a, b int) (int, int) {
	return a / b, a % b
}
```

若多返回值函数在调用时仅想使用一个返回值可使用下划线 `_` 代替其他返回值。

```go
q, _ := div(a, b)
```

多返回值函数一般用于错误处理即返回一个正确返回值与错误。正常情况下错误为 `nil` ，出错时课直接返回错误。

```go
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported Operation: %s", op)
	}
}
```

Go 语言函数可以以函数作为参数：

```go
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function: %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}
```

函数在作为参数传入另一个函数时可使用匿名函数：

```go
fmt.Println(apply(func(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}, 2, 3))
```

Go 语言中函数只有可变参数列表等特性，在参数类型请加上 `...` 即可使用可变参数列表。

```go
func sum1(numbers ...int) int {
	s := 0
	for i :=range numbers {
		s += numbers[i]
	}
	return s
}
```

### 指针

Go 语言中指针定义方式与 C 语言相似。

```go
var 变量名 *变量类型 = &其他变量
```

Go 语言中指针无法参加运算。

Go 语言中参数传递方式只有值传递一种。

```go
func swap(a, b int) {
	a, b = b, a
}
```

上述代码无法交换 `a` 与 `b` 的值。

```go
func swap(a, b *int) {
	*a, *b = *b, *a
}
```

上述代码可以交换 `a` 与 `b` 的值。