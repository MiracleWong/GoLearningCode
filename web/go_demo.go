package main

import (
	"html/template"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age int
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	// 解析 html 文件，自动创建一个模板，关联到变量 t1 上
	// 备注这个地方，写什么路径？绝对路径 or 相对路径
	// web/test.html 时，在GoLand 中运行 OK，go run 则不行。反之
	t1, err := template.ParseFiles("web/test.html")
	if err != nil {
		panic(err)
	}
	// 执行解析后的模板对象，合并&替换
	t1.Execute(w, "Hello World")
}

func main() {
	//server := http.Server{
	//	Addr: "127.0.0.1:8080",
	//}
	//http.HandleFunc("/tmpl", tmpl)
	//server.ListenAndServe()

	p := Person{"MiracleWong", 29}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}