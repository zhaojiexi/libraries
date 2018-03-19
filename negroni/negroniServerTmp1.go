package main

import (
	"net/http"
	"fmt"
	"github.com/codegangsta/negroni"
	"libs/leaf/log"
)

func main() {

	serverMux:=http.NewServeMux()
	serverMux.Handle("/",negroni.New(
		negroni.HandlerFunc(middle),
		negroni.Wrap(http.HandlerFunc(t1)),
	))

	n:=negroni.Classic()
	//n.UseHandler(serverMux)
	n.Use(negroni.Wrap(serverMux))
	n.Use(negroni.NewLogger())
	//n.UseHandler(serverMux)
	//http.ListenAndServe(":8080",nil)
	n.Run(":8080")
}

func t1(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"welcome this t1")
}
func middle(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){

	if r.PostFormValue("name")=="zjx" {
		log.Release("this middle ready next")
		next(rw,r)
	}


}