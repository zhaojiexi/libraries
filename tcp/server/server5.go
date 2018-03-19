package main

import (
	"net"
	"libs/leaf/log"
	"fmt"
	"io"
	"encoding/binary"
)

func main() {


	li,err:=net.Listen("tcp",":8080")
	if	err!=nil{
		log.Fatal("tcp listen fatail: %s",err)
	}

	for{
		conn,err:=li.Accept()
		if	err!=nil{
			log.Release("conn:%+v Accept fatail: %s",conn,err)
		}
		go func(conn net.Conn){
			defer conn.Close()
			data:=make([]byte,4)
			_,err=io.ReadFull(conn,data)
			//_,err:=conn.Read(data)

			dataid:=binary.BigEndian.Uint32(data)
			fmt.Println("id ",dataid)

/*			if dataid==1 {
				conn.Write([]byte("success"))
				a:=make([]byte,1024)
				conn.Read(a)
				fmt.Printf("ssssssss%s",a)
			}else if dataid==2{
					fmt.Println("dataid=====2")
				a:=make([]byte,5)

					conn.Read(a)
					fmt.Printf("收到客户端发来的心跳%s",a)
					conn.Write([]byte("result heart"))
			}*/
			if dataid==2 {
				for{

					a:=make([]byte,10)
					_,err:=io.ReadFull(conn,a)
					if	err!=nil{
						fmt.Println(err)
					}
					fmt.Println(a)

				}
			}
			if err!=nil {
				log.Release("read data fatail: %s",err)
			}

		}(conn)



	}

	}

