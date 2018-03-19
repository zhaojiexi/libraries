package main

import (
	"net"
	"io"
	"encoding/binary"
	"fmt"
	"encoding/json"
)
type H struct {
	Hello struct{
		Name string `json:"Name"`
	} `json:"Hello"`
	
}

var byteOrder = binary.BigEndian
func main(){

	li,err:=net.Listen("tcp",":8989")
	if err!=nil {
		panic(err)
	}
	for {

		con,err:=li.Accept()
		if err!=nil {
			panic(err)
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()

			msglen:=make([]byte,4)

			_,err:=io.ReadFull(conn,msglen)

			if	err!=nil{
				fmt.Printf("err read msg len %s \n",err)
			}

			datalen:=byteOrder.Uint32(msglen)
			data:=make([]byte,datalen)

			_,err=io.ReadFull(conn,data)

			if	err!=nil{
				fmt.Printf("err read data %s \n",err)
			}

			a:=H{}
			//m:=make(map[string]interface{})

			fmt.Println(string(data))
			json.Unmarshal(data,&a)

			fmt.Printf("%+v\n",a)
			
			//fmt.Println(m["Hello"].(map[string]interface{})["Name"])

		}(con)




	}


}

