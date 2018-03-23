package main

import (
	"fmt"
	"reflect"
	"net/http"
	"strings"
	"libs/leaf/log"
)

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

//定义路由器结构类型
type Routers struct {
}

func init() {

	ControllerMaps = make(ControllerMapsType, 0)
	rou := Routers{}

	vl := reflect.ValueOf(&rou)

	vt := vl.Type()

	methodNum := vt.NumMethod()

	for i := 0; i < methodNum; i++ {
		name := vt.Method(i).Name
		ControllerMaps[name] = vl.Method(i)
	}

}

//为路由器结构附加功能控制器函数，值传递
func (this *Routers) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login:", w, r)
	fmt.Fprint(w, "aaaaaaaaaaa")

}
func (this *Routers) GroupV1Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login:", w, r)
	fmt.Fprint(w, "aaaaaaaaaaa")

}

//为路由器结构附加功能控制器函数，引用传递
func (this *Routers) ChangeName(msg *string) {
	fmt.Println("ChangeName:", *msg)
	*msg = *msg + " Changed"
}

func main() {
	r := &Routers{}
	err := http.ListenAndServe(
		":8888", r,
	)

	if err != nil {
		panic(err)
	}

}

func (this *Routers) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.String())

	resw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	vals := []reflect.Value{resw, rr}

	url := r.URL.String()

	if url[:1]=="/" {
		log.Release("路由格式错误%s",url)
		return
	}
	//	a := "/group/v1/get"
	//首字母转为大写// GroupV1Get
	index := strings.Index(url, "/")
	fmt.Println(index)
	if index == 0 {
		url = url[index+1:]
		fmt.Println(url)
	}
	arr := strings.Split(url, "/")
	fmt.Println(arr)
	var conAll string
	for i := 0; i < len(arr); i++ {
		con := arr[i][:1]
		conUp := strings.ToUpper(con)
		conNew := conUp + arr[i][1:]
		fmt.Println("con", conNew)
		conAll += conNew
	}

	fmt.Println(conAll)

	ControllerMaps[conAll].Call(vals)

}
