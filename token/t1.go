package main

import (
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/codegangsta/negroni"
	)

var (
	key = "key"
)

func main() {

	http.HandleFunc("/", test)
	http.Handle("/login", negroni.New(
		negroni.HandlerFunc(logihandler),
		negroni.Wrap(http.HandlerFunc(result)),
	))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	//pwd := r.PostFormValue("pwd")

	//tp := r.Header.Get("type")

	fmt.Println("name is ",name)
		jwttest(w)

	fmt.Fprint(w)
}

func jwttest(w http.ResponseWriter) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["name"]="zjx"
	token.Claims = claims

	s, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}


	fmt.Fprint(w,s)
}

func logihandler(w http.ResponseWriter, r *http.Request,next http.HandlerFunc){
	fmt.Println("key",string(key))
	fmt.Printf("AuthorizationHeaderExtractor= %+v\n", request.AuthorizationHeaderExtractor.Filter)
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	fmt.Printf("token=%+v \n",token)
	if err == nil {
		if token.Valid {
			fmt.Println(r.PostFormValue("name"))
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource ",err)
	}

}

func result(w http.ResponseWriter, r *http.Request){
	name:=r.PostFormValue("name")
	fmt.Println("name",name)
	fmt.Fprint(w,"aaaaaaaaaaaaaaaa ", name)
}