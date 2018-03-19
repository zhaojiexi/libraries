package main

import (
	jwt"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

func main() {
		//parse  token and key
	c,err:=parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjExNzE2OTEsImlhdCI6MTUyMTE2ODA5MSwibmFtZSI6InpqeCJ9.rb7MIe1kLSB9281ezHKtV_JYUUA6Xf_dqxwv7afAjcg","key")
	if !err {
		panic(err)
	}
	fmt.Println(c)
	fmt.Println(c.(jwt.MapClaims)["exp"])
	a:=int64(c.(jwt.MapClaims)["exp"].(float64))

	t:=time.Unix(a,0)
	fmt.Println(t)

	now:=time.Now()

	sub:=t.Unix()-now.Unix()
	fmt.Println(sub/60)

}

func parseToken(tokenString string, key string) (interface{}, bool){
		//token and func
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("token claim",token.Claims)

		//token method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	fmt.Println(token)
	//判断类型是否正确
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
