## 来源
* https://helight.info/2018-03-21/golang%E7%9A%84struct%E5%92%8Cinterface/

## struct
* struct 用来自定义复杂数据结构，可以包含多个字段（属性），可以嵌套；go中的struct类型理解为类，可以定义方法，和函数定义有些许区别；struct类型是值类型。

* struct定义:
    ```go
    type User struct {
        Name string
        Age  int32
        mess string
    }
    var user User
    var user1 *User = &User{}
    var user2 *User = new(User)
    ```

* struct的方法
    * 在go语言中，我们可以 __为自定义类型定义类型相关的方法__，比如：

    ```go
    func (p *player) Name() string{
        return p.name
    }
    ```

    * 上面的代码为 player 这个自定义类型声明了一个名为 Name 的方法，该方法返回一个string。值得注意的是 `p *player` 这段代码指定了我们是为 player 创建方法，并将调用该方法的实例指针当作变量p传入该函数，如果没有 `p *player` 这段代码，这个方法就变成了一个普通的全局函数。

* struct的嵌入（Embedding）
    * go语言中的 “_继承_” 和其他语言中的继承有很大区别，比如：

    ```go
    type player struct{
        User
    }
    ```

    * 这是一种 “_继承_” 的写法，在go语言中这种方式叫做 __“嵌入”（embed）__，_此时 player类型 就拥有了 User类型 的 Name 等变量_

* struct的tag
    * 这种方式主要是用在 xml，json 和 struct 间相互转换，非常方便直观，比如接口给的参数一般是 json 传过来，但是内部我们要转为 struct 再进行处理。

    ```go
    import "encoding/json"
    type User struct {
        Name string `json:"userName"`
        Age  int    `json:"userAge"`
    }
    func main() {
        var user User
        user.Name = "nick"
        user.Age = 18    
        conJson, _ := json.Marshal(user)
        fmt.Println(string(conJson))    //{"userName":"nick","userAge":0}
    }
    ```
