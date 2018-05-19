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
