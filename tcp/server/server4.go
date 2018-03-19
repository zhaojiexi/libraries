package main

import (
	"net"
	"libs/leaf/log"
	"fmt"
)

func main() {


	ls,err:=net.Listen("tcp",":8080")
	if	err!=nil{
		log.Fatal("listen err",err)
	}
	fmt.Println("listen begin")
	for  {
		con,err:=ls.Accept()

		go func(con net.Conn) {
			defer con.Close()
			if	err!=nil{
				log.Fatal("accpt err",err)
			}
			a:=make([]byte,1024)
			con.Read(a)
			fmt.Println(a)
			fmt.Println(string(a))
		}(con)

	}




}