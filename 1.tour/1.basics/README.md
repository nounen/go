## Go 基础


### 包
* 每个 Go 程序都是由包构成的。

* 程序从 `main` 包开始运行。

* `package` 关键字 __声明包__ (TODO: tour 暂时没有这个)

* `import` 关键字 __导入包__
```go
import "fmt"
import "math"

# 或者: 分组形式导入
import (
	"fmt"
	"math"
)
```

* 导出名
    * 如果一个名字(方法名)以 __大写字母开头，那么它就是已导出的__ 


### 函数
* 注意: __数据类型__ 在 __变量名__ 的 __后面__

* 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
    * eg: `func add(x, y int)` // x, y 都是 int

* 多值返回: 函数可以返回任意数量的返回值
    ```go
    // 返回值是 y 和 x, string 类型
    func swap(x, y string) (string, string) {
        return y, x
    }
    ```

* 命名返回值: Go 的返回值可被命名，它们会 _被视作定义在函数顶部的变量_
    * eg: `func split(sum int) (x, y int)` // 函数结尾直接 `return`, 返回的就是 `x` 和 `y`
