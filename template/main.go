package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	//tmpl, err := template.ParseFiles("./hello.tmpl")
	htmlByte, err := ioutil.ReadFile("./hello.tmpl")
	if err != nil {
		log.Fatalln("create template err:", err)
		return
	}

	// 自定义一个夸人的模板
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在parse前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
	if err != nil {
		log.Fatalln("create template err:", err)
		return
	}

	// 利用给定数据渲染模板， 将结果写入w
	u := user{
		Name:   "Zjy",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, u)
	log.Println("success")
}

func testDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./test.tmpl", "./ul.tmpl")
	if err != nil {
		log.Fatalln("create template err:", err)
		return
	}
	u := user{
		Name:   "Jy",
		Gender: "男",
		Age:    17,
	}
	tmpl.Execute(w, u)
}

func blockDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		log.Fatalln("create template failed, err:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		log.Fatalln("render template failed, err:", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(a string) template.HTML {
			return template.HTML(a)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		log.Fatalln("create template failed, err:", err)
		return
	}
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/tmpl", testDemo)
	http.HandleFunc("/block", blockDemo)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("HTTP server failed, err:", err)
	}
}
