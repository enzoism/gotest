package main

import (
"fmt"
"log"
"net/http"
"syscall"
)

func main() {
	http.HandleFunc("/", sayhelloName2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

// 1、通过浏览器输出hello go（返回数据是写死的）
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 1.进行请求地址的打印
	fmt.Println("path:", r.URL.Path)
	// 2.将参数返回到页面上
	fmt.Fprintf(w, "hello go")
}

// 2、通过浏览器输出hello go（返回数据是动态的）
func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 1.外部动态的传入参数
	v, ok := syscall.Getenv("TASKID")
	log.Println("Getenv", v, ok)

	// 2.进行请求地址的打印
	fmt.Println("path:", r.URL.Path)
	// 3.将参数返回到页面上
	fmt.Fprintf(w, "hello go")
	fmt.Fprintf(w, v)
}