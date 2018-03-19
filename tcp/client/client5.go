package main

import (
	"net"
	"libs/leaf/log"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	si:=make(chan os.Signal,1)
	signal.Notify(si,os.Interrupt)
	conn,err:=net.Dial("tcp",":8080")

	if	err!=nil{
		log.Error("conn err %s",err)
	}

	//testsend(conn)

	for{
		heart(conn)
		time.Sleep(time.Second*3)
	}


	<-si
	 fmt.Println("断开连接")
	 conn.Write([]byte("close"))
	signal.Stop(si)
	}

	func testsend(conn net.Conn){
		defer conn.Close()
		bf:=bytes.NewBuffer(nil)
		var a int32
		a=1
		bf.Write([]byte(""))
		binary.Write(bf,binary.BigEndian,a)
		binary.Write(bf,binary.BigEndian,[]byte(""))

		data:=bf.Bytes()


		_,err:=conn.Write(data)
		if err!=nil {
			log.Release("conn write err %s",err)
		}
		result:=make([]byte,10)


		_,err=conn.Read(result)

		if err!=nil {
			log.Release("conn Read err %s",err)
		}
		fmt.Printf("result %s",result)
	}

	func heart(conn net.Conn){
		defer conn.Close()

		buf:=bytes.NewBuffer(nil)
		var a int32
		a=2
		binary.Write(buf,binary.BigEndian,a)
		binary.Write(buf,binary.BigEndian,[]byte("heart"))

		conn.Write(buf.Bytes())
		fmt.Println("客户端发送",buf.Bytes())
		resu:=make([]byte,100)
		conn.Read(resu)
		fmt.Println("服务器返回",string(resu))


	}