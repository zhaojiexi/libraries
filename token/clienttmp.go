package main

import (
//	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

func main() {

	val := url.Values{
		"name": {"zjx"},
		"pwd":  {"123"},
	}
	c := http.Client{
	}
	req,err:=http.NewRequest("POST","http://localhost:8080", strings.NewReader(val.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("type","jwt")

	rsp,err:=c.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	b,err:=ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))


	loginWithToken(string(b))



}

func loginWithToken(token string){

	c := http.Client{
	}
	val2 := url.Values{
		"name": {"zjx"},
		"pwd":  {"123"},
	}

	res2,err:=http.NewRequest("POST","http://localhost:8080/login",strings.NewReader(val2.Encode()))
	if err != nil {
		panic(err)
	}
	res2.Header.Set("Content-Type","application/x-www-form-urlencoded")
	res2.Header.Add("Authorization",string(token))

	fmt.Println(res2)
	rsp2,err:=c.Do(res2)
	if err != nil {
		panic(err)
	}

	defer rsp2.Body.Close()

	b2,err:=ioutil.ReadAll(rsp2.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("result :=",string(b2))

}