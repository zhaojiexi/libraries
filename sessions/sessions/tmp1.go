package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
	"encoding/gob"
)

const (
	SESSION_KEY = "key"
)

var (
	s_store *sessions.CookieStore
	m       = make(map[string]interface{})
)

func init() {
	s_store = sessions.NewCookieStore([]byte(SESSION_KEY))
	//使用复杂的结构存储 或结构体 必须先注册
	gob.Register(&m)
}
func main() {

	http.HandleFunc("/show", showSessionValues)
	http.HandleFunc("/", sessiontest)
	http.HandleFunc("/del", sessiondel)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
func sessiontest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test")
	name := r.PostFormValue("name")

	s, err := s_store.Get(r, "t1")
	fmt.Println("session is new:", s.IsNew)
	if err != nil {
		panic(err)
	}

	m["s"] = "this M map have string"
	m["i"] = "this M map have int"

	s.Values["name"] = name
	s.Values["m"] = m
	s.Save(r, w)
	fmt.Println("sessiontest", s.Values)

	fmt.Fprint(w, s.Values)
}
func sessiondel(w http.ResponseWriter, r *http.Request) {

	s, err := s_store.Get(r, "t1")
	fmt.Println("session is new:", s.IsNew)
	if err != nil {
		panic(err)
	}
	fmt.Println("ssss:", s)
	delete(s.Values, "name")

	s.Save(r, w)
	fmt.Println("sessiondel", s.Values)
	fmt.Fprint(w, s.Values)
}
func showSessionValues(w http.ResponseWriter, r *http.Request) {
	s, err := s_store.Get(r, "t1")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%+v", s.Values["m"])
}
