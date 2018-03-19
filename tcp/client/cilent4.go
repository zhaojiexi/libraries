package main

import (
	"net"
	"libs/leaf/log"
	"encoding/binary"
	"bytes"
)

func main(){

	cn,err:=net.Dial("tcp",":8080")
	if err!=nil {
		log.Fatal("dial err",err)
	}

		buf:=bytes.NewBuffer(nil)
		binary.Write(buf,binary.BigEndian,[]byte("byte"))

		buf.Write([]byte("string"))
	b:=buf.Bytes()
		cn.Write(b)

	}

