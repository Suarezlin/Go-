# 4 复合数据类型
主要介绍四种复合数据类型 —— 数组、slice、map 和 结构体。
## 数组
数组是一个由固定长度的特定类型元素组成的序列，在 Go 语言中由于数组长度固定，所以一般不直接使用数组而是使用 `Slice` 

Go 语言数组定义格式：

```go
var 数组名 [数组长度]元素数据类型
```

可以在定义时就给数组赋值：

```go
var a [3]int = [3]int{1, 2, 3}
```

也可以使用简化变量定义来赋值：

```go
b := [3]int{1, 2, 3}
```

若在定义时直接赋值，则可以用 `...` 来代替数组长度，让编译器自己推断数组长度。

```go
c := [...]int{1, 2, 3}
```

访问数组元素可以通过元素下标来实现，Go 语言中数组下表从 `0` 开始一直到 `数组长度 - 1` 。

```go
e := c[0]
e = c[1]
e = c[2]
```
若数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时我们可以可以通过 `==` 来比较两个数组。只有当两个数组元素都是相等的时候数组才能相等。

数组作为函数参数时，一般使用指向数组的指针。尽管可以使用指针，数组类型依然是僵化的类型，数组信息包含了僵化的长度信息，所以一般使用 `Slice` 来替代数组。

## Slice

`Slice` 代表变长的序列。

`Slice` 定义格式：

```go
Slice 变量名 := []元素类型{元素值}
```

Go 语言中，`Slice` 和数组定义非常相似，只是不需要填入数组长度。

数组和 `Slice` 之间有着很密切的关系，一个 `Slice` 可以看做一个数组的映射，而且 `Slice` 底层确实引用了一个数组对象。`Slice` 由三个部分构成：指针、长度和容量。指针指向第一个元素对应底层数组的地址，长度对应 `Slice` 中元素个数，容量是从 `Slice` 开始到底层数组的结尾位置。

```go
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

```

运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fprgslcl6tj31js0960u9.jpg)

`Slice` 可以通过对数组进行切片得到。

```go
a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
s:= a[2 : 5] // 3, 4, 5
```

被切片的数组就是创建出的 `Slice` 的底层数组。

和数组不同， `Slice` 之间无法直接比较，标准库提供了 `bytes.Equal()` 函数来判断字节型 `Slice` 是否相等，对于其他类型的 `Slice` 我们只能遍历判断元素是否相等。

```go
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
```

Go 语言中数组和 `Slice` 的遍历方式比较接近。

```go
func traverse1(s []string)  {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
```

这种方式与其他语言类似，直接定义循环变量，然后通过循环变量来访问 `Slice` 元素。

```go
func traverse2(s []string) {
	for i, e := range s {
		fmt.Printf("(%d %s)\n", i, e)
	}
}
```

通过 `range` 可以获取每一次循环的下表与元素值。

运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fprh5j7tm9j31js09ejt2.jpg)

直接通过 `var s []int` 创建的 `Slice` 默认值为 `nil` ，要想创建一颗空 `Slice` 应使用 `make()` 函数。

```go
s := make([]int, 0, 10)
```

`make()` 函数第一个参数为创建什么类型的 `Slice`，第二个参数为 `Slice` 长度，第三个参数为底层数组的容量。

向 `Slice` 中添加元素可以使用内置的 `append()` 函数。

```go
var runes []rune
for _, e := range "Hello, 世界" {
	runes = append(runes, e)
}
fmt.Println(runes)
```

运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fprhd13ofxj31ju09cwg5.jpg)

使用 `append()` 函数时，若 `Slice` 在添加的元素之前，则元素被添加到 `Slice` 的最后，反之被添加到 `Slice` 的最前方。

`append()` 函数在调用后，若原 `Slice` 底层数组容量不够，则会自动分配新容量，这会导致底层数组的改变，为了保险起见我们一般在 `append()` 结束后再将结果赋给原 `Slice` 。

```go
package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 3)
	for i := 1; i <= 20; i++ {
		s = append(s, i)
		fmt.Printf("%d %d\n", len(s), cap(s))
	}
}

```

运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fprhjw47aqj31jw0okq5p.jpg)

可以看到，当容量不够时，会分配两倍于当前容量的新容量。

## Map

在 Go 语言中一个 `Map` 就是对一个哈希表的引用，`Map` 类型可以写为 `map[K]V` 其中 `K` 为 `key` 的类型，`V` 为 `value` 的类型。

我们可以使用内置的 `make()` 函数来创建一个 `Map` ：

```go
ages := make(map[string]int)
```

也可以通过语法直接创建 `Map` 并赋初值。

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
```

`Map` 中的元素通过 `key` 来返回。

```go
fmt.Println(ages["alice"])
```
若元素不在 `Map` 中，上述操作将返回零值。为了确定是否操作成功，我们可以：

```go
age, ok := ages["alice"]
if !ok {
    return 0
} else {
    //operation
}
```

通过 `key` 获取元素会返回两个值，若 `ok` 为 `true` 则返回正常，若为 `false` 则反回失败。

遍历 `Map` 的方式与遍历数组和 `Slice` 的方式相同，都可以使用 `range` 来遍历。

```go
for key, value := range ages {
	fmt.Printf("$s\t%d\n", key, value)
}
```

`Map` 是无序的，每次迭代的顺序可能都不一样。

## 结构体

结构体是一种聚合数据类型，是由零个或多个任意类型的值聚合成的实体。

结构体定义：

```go
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
var dilbert Employee
```

结构体成员变量可以使用 `.` 来访问。

```go
dilbert.ID = 1
```

与 C 语言不同， Go 语言中指向结构体的指针访问结构体变量时也是用 `.` 操作符。

结构体的零值为每个成员变量均为零值。如果结构体没有任何成员那么就是空结构体，写作 `struct{}` 。

结构体成员变量的值可以通过字面值来赋值。

```go
type Point struct {
	x, y int
}

p := Point{1, 2}
```

更为常用的是以下写法

```go
type Point struct {
	x, y int
}

p := Point{x: 1, y: 2}
```

Go 语言结构体可以作为函数的参数或返回值，若结构体较大则推荐使用指向结构体的指针作为参数。若需要修改传入的结构体则必须使用指向结构体的指针。因为 Go 语言参数传递方式为值传递。

如果结构体的所有成员变量是可以比较的则结构体也可以比较的，可以使用 `==` 和 `!=` 来比较两个结构体。

```go
type Point struct {
	X, Y int
}

func main() {
	// var dilbert Employee
	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p == q)
}
```

运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpsd25rx20j31jw08875q.jpg)

Go 语言拥有特殊的结构体嵌入机制让一个命名的结构体包含另一个结构体类型的匿名成员。

```go
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}
```

我们可以像上面那样简化定义结构体，但这样也会使结构体的修改成员变量变得繁琐。Go 语言提供了特殊的结构体嵌套机制，我们可以只声明一个成员对应的数据类型二部指明其名字，这类成员称为匿名成员。

匿名成员的数据类型必须是命名的类型或是一个指向一个命名的类型的指针。

```go
type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
```

上面这种定义方式中，我们可以认为 `Point` 嵌入了 `Circle` 中，`Circle` 嵌入了 `Wheel` 中，这样我们可以直接访问嵌套结构体的成员。

```go
var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```

但是，匿名成员声明时不能采用简短声明的方式。

## 文本和 HTML 模板

有时候我们可能需要打印复杂的文本格式，这时候一般要将格式化代码分离出来以便安全的修改，这些功能由 `text/template` 和 `html/template` 等模板包提供的。

```go
const temp1 = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Numbers: {{.Numbers}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}
`
```

以上为一个简单的模板字符串。模板中 `{{range .Items}}` 和 `{{end}}` 对应一个循环 `action` 。在 `action` 中 `|` 表示将前一个表达式的结果作为后一个函数的输入。

生成模板需要两个处理步骤。第一步是要分析模板并转为内部表示，然后基于指定的输入执行模板。

```go
report, err := template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1)
if err != nil {
	log.Fatal(err)
}
```

`template.New()` 先创建并返回一个模板，`Funcs()` 方法将自定义函数注册到模板中，最后调用 `Parse()` 分析模板。

我们可以使用 `template.Must()` 来简化错误处理。

```go
var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func main() {
	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
```