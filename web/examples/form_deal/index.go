package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	r.ParseForm()

	if r.URL.Path == "/favicon.ico" {
		return
	}

	fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println("url_long", r.Form["url_long"])

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println("Form:", r.Form) // //这些信息是输出到服务器端的打印信息

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello lll!") //这个写入到w的是输出到客户端的

	// request.Form是一个 url.Values 类型. TODO: 没搞懂 url.Values 是什么
	//v := url.Values{}
	//fmt.Println("v:", v)
	//
	//v.Set("name", "Ava")
	//v.Add("friend", "Jess")
	//v.Add("friend", "Sarah")
	//v.Add("friend", "Zoe")
	//// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	//fmt.Println(v.Get("name"))
	//fmt.Println(v.Get("friend"))
	//fmt.Println(v["friend"])

	fmt.Println("==================================\n")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	input := r.Form

	// r.Form里面包含了所有请求的参数，比如URL中query-string、POST的数据、PUT的数据，所以当你在URL中的query-string字段和POST冲突时，会保存成一个slice，里面存储了多个值，Go官方文档中说在接下来的版本里面将会把POST、GET这些数据分离开来。
	fmt.Println("Form:", input)

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")

		log.Println(t.Execute(w, nil))
	} else {
		// 校验密码
		if len(input.Get("password"))  == 0 {
			fmt.Println("请输出密码", input.Get("password"))
		}

		// 校验年龄
		getAge, err := strconv.Atoi(input.Get("age"))

		if err != nil {
			fmt.Println("年龄请输入整数", getAge)
		}

		// 校验下拉框
		fruits := []string{
			"apple", "pear", "banana",
		}

		fruit := input.Get("fruit")

		for _, item := range fruits {
			if item == fruit {
				fmt.Println("下拉选项正确")
			}
		}

		// 校验单选框
		genders := []string{
			"1", "2",
		}

		for _, item := range genders {
			if item == input.Get("gender") {
				fmt.Println("单选框正确")
			}
		}

		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])

		// func HTMLEscapeString(s string) string // 转义s之后返回结果字符串
		fmt.Println("username", template.HTMLEscapeString(input.Get("username")))
		fmt.Println("password", template.HTMLEscapeString(input.Get("password")))

		// func HTMLEscape(w io.Writer, b []byte) // 把b进行转义之后写到w
		template.HTMLEscape(w, []byte(input.Get("username")))
		template.HTMLEscape(w, []byte(input.Get("password")))

		fmt.Println("==================================\n")
	}
}

func main() {
	fmt.Println("localhost:9090")

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
