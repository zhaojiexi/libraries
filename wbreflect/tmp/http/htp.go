package http

import (
	"strings"
	"fmt"
	"reflect"
	"net/http"
	"libraries/wbreflect/tmp/router"
)

type S struct{}

func Start() error {
	s := S{}

	err := http.ListenAndServe(":8080", &s)

	return err
}

func (s *S) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	u := r.URL.String()

	url := strings.Replace(u, "/", "", 1)
	fmt.Println("url", url)
	if url == "favicon.ico" {
		return
	}

	vs := []reflect.Value{
		reflect.ValueOf(w),
		reflect.ValueOf(r),
	}

	V := router.RT[url].Call(vs)
	fmt.Println(V)
}
