# From
* 配合文档： https://books.studygolang.com/gopl-zh/

* 文档代码： https://github.com/adonovan/gopl.io/


### 2.4 赋值
```go
// 交换两个变量的值
x, y := 0, 1
x, y = y, x+y


v = m[key]      // map 查找，失败时返回零值
v = x.(T)       // type 断言，失败时panic异常
v = <-ch        // 管道接收，失败时返回零值（阻塞不算是失败）


_, err = io.Copy(dst, src) // 丢弃字节数
_, ok = x.(T)              // 只检测类型，忽略具体值
```



# The Go Programming Language

This repository provides the downloadable example programs
for the book, "The Go Programming Language"; see http://www.gopl.io.

These example programs are licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.<br/>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png"/></a>

You can download, build, and run the programs with the following commands:

	$ export GOPATH=$HOME/gobook            # choose workspace directory
	$ go get gopl.io/ch1/helloworld         # fetch, build, install
	$ $GOPATH/bin/helloworld                # run
	Hello, 世界

Many of the programs contain comments of the form `//!+` and `//!-`.
These comments bracket the parts of the programs that are excerpted in the
book; you can safely ignore them.  In a few cases, programs
have been reformatted in an unnatural way so that they can be presented
in stages in the book.

