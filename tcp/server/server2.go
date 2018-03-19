package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"time"
)

func main() {

i:=0
	ls,err:=net.Listen("tcp",":8080")
	defer ls.Close()
	if err != nil {
		log.Print("ls err",err)
	}
	for  {


		i++
		fmt.Printf("wait conn %d \n",i)
		conn,err:=ls.Accept()
		if err != nil {
			log.Print("conn err",err)
		}


		func(conn net.Conn){
		defer conn.Close()
			msglen:=make([]byte,4)

			fmt.Println(conn.RemoteAddr().String())
			_,err:=io.ReadFull(conn,msglen)

			if err != nil {
				log.Print("ls err",err)
			}
		fmt.Println(msglen)
		err=conn.SetReadDeadline(time.Now().Add(time.Duration(1)*time.Second))
		fmt.Printf("err %+v",err)
		}(conn)

	}


}