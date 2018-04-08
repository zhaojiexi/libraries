package router

import (
	"reflect"
	"fmt"
	"net/http"
	"libraries/wbreflect/tmp/db"
	"libs/leaf/log"
	"cx/server/common/mysqlEngine"
)

type M1 struct{}

var (
	RT = make(map[string]reflect.Value)
)

func init() {
	m := M1{}

	vf := reflect.ValueOf(&m)
	vt := vf.Type()

	num := vt.NumMethod()

	fmt.Println(num)
	for i := 0; i < num; i++ {
		name := vt.Method(i).Name
		RT[name] = vf.Method(i)
		fmt.Println(name)
	}
}

func (m *M1) T1(w http.ResponseWriter, r *http.Request) {

	d := db.DB{}

	name := r.FormValue("name")

	log.Release("form value name :%s ", name)

	d.Start()
	d.M.RequestByName("wbtest.selecttest1", nil, mysqlEngine.IDBSinkFunc2(db.SelectAdmin),w,r, name)


	fmt.Fprint(w, "this t1 func ")
}

func (m *M1) T2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "this t2 func ")
}
