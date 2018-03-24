# 1 入门
这一部分是 Go 语言的入门内容，介绍了 Go 语言的基础组件。
## Hello, world!
Go 语言 Hello, world 程序：
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

```
Go 语言是一门编译语言，运行程序：
```bash
go run helloWorld.go
```
运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnx2p5b4mj31js0743zu.jpg)

以上为直接运行 Go 语言程序。我们还可以将 Go 语言程序编译为可执行文件。
```bash
go build helloWorld.go
./helloWorld
```
执行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnx5nvf7bj31jw06yq45.jpg)

Go 语言的代码通过包来组织。此程序开头就有 `package main` ，用来声明这个文件属于 `main` 包。

`main` 包内可以有一个叫做 `main` 的函数，它是这个程序的入口。在编译时，一个 `main` 包内只能有一个 `main` 函数。否则无法编译。
## 命令行参数
Go 语言获取输入数据的途径有许多。输入可以来自程序外部：文件、网络连接、其他程序的输出、用户、命令行参数等。首先是命令行参数。

要获取程序的命令行参数，我们需要导入 `os` 包。使用 `os.Args` 即可访问命令行参数。

Go 语言命令行参数为 `slice` 类型。`slice` 类型与 `Python` 中的数组类似。

`os.Args` 的第一个元素即 `os.Args[0]` 为命令本身的名字，其他元素为命令行的参数。所以通常使用 `os.Args[1:]` 来获取所有的命令行参数。
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	args := os.Args[1:] // os.Args 是一个 string 类型的切片
	for _, arg := range args {
		s += " " + arg
	}
	fmt.Println(s)
}

```
运行：
```bash
go run args.go a b c d e f g
```
运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnxo4fc6rj31jw07eq3x.jpg)

程序成功获取到了通过命令行输入的数据。

这个程序使用了 Go 原因的 `for` 循环，Go 语言中只有 `for` 一种循环语句。之后又使用了 `range` 关键字，它可以返回切片或数组的 `index` 和对应的 `value`。由于 Go 语言定义变量必须要使用，这里我们有不需要切片的 `index`，所以我们使用 `_`  也就是`空标识符`来表示我们不需要 `index` 。

在 Go 语言中变量的声明格式为：`var 变量名 变量类型` 。还有一种简写 `变量名 := 值` ，后一种声明方式只能用在函数内部。且声明变量且赋初值时，可以省略变量类型，由编译器自行推断。

Go 语言在声明变量后会自动将变量的值初始化，数字会初始化为 0 ，字符串会初始化为空字符串。`map` 和 `slice` 会被初始化为 `nil`，要想定义一个没有内容的 `map` 或 `slice` 要使用 `make()` 来创建，如：
```go
var m = make(map[string]string)
var s = make([]int, 0, 10) //第一个参数为创建一个 int slice，第二个参数为 slice 的长度为零表示空 slice，第三个参数为 slice 的物理内存大小，向 slice 中添加元素时，若此空间不够会自动分配更大的空间。
```
## 查找重复的行
这是一个统计文件中各行重复的次数的 Go 语言程序。
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d    %s\n", n, line)
		}
	}
}
```
运行后从键盘输入，结束后使用 `ctrl+z` (Windows) `ctrl+d` (linux/Mac OS) 来结束输入。

结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnyfl12boj31jy0dotad.jpg)

这里首先用了之前说的方式创建了一个 `Map` ，它的键的类型为 `string` ，值的类型为 `int` 。Go 语言中的 `Map` 与 `Java` 中的 `HashMap` ，`Python` 中的 `dict` 类似，使用 `hash` 实现。

程序中使用了 `bufio` 包来处理输入。程序创建了一个 `Scanner` 来读取标准输入中的数据。每次调用 `input.Scan()` 时就会读入下一行，最后一行读取完后，会返回 `false` 结束循环。
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			if f, err := os.Open(file); err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
				f.Close()
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

```
我们修改一下程序，使程序即可以通过标准输入来输入数据也可以通过外部文件获取数据。

Go 语言中，我们通过 `os.Open()` 函数来打开程序。

`func countLines(file *os.File, counts map[string]int)` 

为什么我们要使用一个指向 `os.File` 的指针来作为函数的参数？因为 `os.Open()` 函数返回一个指向 `os.File` 的指针。

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnz0ck6nqj30gm038t90.jpg)

```go
if f, err := os.Open(file); err != nil {
	fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	continue
} else {
	countLines(f, counts)
	f.Close()
}
```
上述错误处理方式在 Go 语言中非常常用。Go 语言中，我们可以在 `if` 内定义变量，它们的作用域为整个 `if block`。

使用以下命令来运行：
```bash
go run dup2.go test1.txt test2.txt
```
运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnz3mqg4aj31jw08awfu.jpg)

第三种写法：
```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, lines := range strings.Split(string(data), "\n") {
			counts[lines]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

```
第三种写法使用了 `ioutil.ReadFile()` 函数，一次性将文件内容读入，再做计数处理。

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnzfaxvy1j30iy01awej.jpg)

可以看出 `ioutil.ReadFile()` 返回的文件内容为一个 `byte` 类型的切片。我们在使用时要先将其转为 `string` 类型。

```go
for _, lines := range strings.Split(string(data), "\n")
```

这里使用了 `strings.Split()` 函数将文件的每行分割开，进行计数统计。

使用以下命令运行：
```bash
go run dup3.go test1.txt test2.txt
```
运行结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpnzjy66rwj31jy098wfw.jpg)

## GIF 动画
Go 语言标准库 `image` 用法演示。

生成李沙育图形(lissajous) (貌似做电路实验的时候在示波器上看过)

```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	lissajous(os.Stdout)
}

```
```bash
go run lissajous.go > out.gif
```
结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpo03hage7g305l05lmzx.jpg)

这个程序中定义了常量，在 Go 语言中常量使用关键字 `const` 来定义。
## 获取 URL
这里简单介绍了 Go 语言的 `http` 包。

这个程序与 `curl` 命令类似，获取一个 `url` 的源文本。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}

```

`http.Get()` 是发送 `GET` 请求的函数，如果没有出错可以在 `resp` 这个结构体中获取请求结果。`ioutil.ReadAll()` 函数读取 `resp.Body` 的内容。

```bash
go run fetch.go http://ip.cn
```

结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpo0l61w9gj31jy09c0v3.jpg)

## 并发获取多个 URL

Go 语言对并发编程的支持非常良好，我们来试一下并发获取多个 `url` 的内容。

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

```

`goroutine` 是一种函数的并发执行方式，`channel` 用来在 `goroutine` 之间进行参数传递。

`go 函数()` 操作表示创建一个 `goroutine` ，并在这个 `goroutine` 中运行函数。

这个程序中 `io.Copy()` 把响应内容输出到 `ioutil.Discard` 中，这个输出流可以看做是一个垃圾桶，可以向其中写入不需要的数据。

当函数需要返回时，会向 `ch` 中写入数据。程序最后有一个循环用来打印 `ch` 中的数据。

当一个 `goroutine` 尝试读取或写入时会阻塞在调用处，知道另一个 `goroutine` 写入或读取值时才会继续进行。这个例子中，`fetch` 函数负责写入数据，主函数读取数据，这样只有所有的 `goroutine` 都向 `ch` 中写入数据后，主程序才会结束。

```bash
go run fetchall.go http://baidu.com http://sina.com http://qq.com http://zhihu.com
```

结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpo2a1lb2kj31jy096gni.jpg)

### Web 服务

Go 语言的内置库可以让我们很方便的编写 Web 服务器。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listen in http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

```

`main` 函数将所有发送到 `/` 下的请求和 `handler` 函数关联，监听 `8000` 端口。请求是 `http.Request` 对象，到达服务器时，这个请求会被交给 `handler` 处理，解析出路径并发送给响应。

结果：

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpo2jsjn6uj31jw096dh5.jpg)

![](https://ws1.sinaimg.cn/large/006m1vi8ly1fpo2k2bj50j309601edfo.jpg)

