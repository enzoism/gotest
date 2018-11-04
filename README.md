# gotest
Docker打包go项目的时候传入外部参数


##1、参考网址：
- go工程启动时带入变量参数：https://blog.csdn.net/niyuelin1990/article/details/79035728

##2、编程过程：
#####1、创建一个工程，使用原生的http进行创建
> ![](https://upload-images.jianshu.io/upload_images/1096351-d05be2129d1dbda8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
#####2、创建一个main.go
```
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
```
#####3、创建Dockerfile文件
```
FROM scratch
MAINTAINER enzo "https://github.com/enzoism"
ENV GOPATH /apps/enzogo/
WORKDIR /apps/enzogo/src/enzoism/gotest/main
COPY . /apps/enzogo/src/enzoism/gotest/main
ADD main /
ENTRYPOINT ["/main"]
```
#####4、进行静态编译+检验项目是否可运行
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
> ![](https://upload-images.jianshu.io/upload_images/1096351-ab3c230017523115.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#####5、进行镜像编译，并运行镜像
```
docker build -t gotest:1.0.0 . 
docker run -i -d --name=gotest  -p 9090:9090 gotest:1.0.0
docker run -i -d --name gotest3 -p 9092:9090 -e TASKID=abc  gotest:1.0.0 
```
#####6、访问网址
> ![](https://upload-images.jianshu.io/upload_images/1096351-86a787cc8a3f034d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

>![image.png](https://upload-images.jianshu.io/upload_images/1096351-aa5e9987bc97af74.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#####7、总结
- 1、使用syscall可以将参数值传入docker打包镜像中
- 2、使用-e TASKID=abc进行镜像编译的可以传入参数到工程中
