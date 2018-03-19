package main

import (
	"net"
	"log"
	"bytes"
	"encoding/binary"
	"time"
	"cx/tmp/tcp/client/basic"
	"fmt"
)
type User struct {
	Name string
	Age int32
}


func (User)GetID()int{
	return 1
}

func main() {

		tc,err:=net.Dial("tcp","localhost:8080")

		if err != nil {
			log.Print("ls err",err)
		}
	rqs:=&basic.ReqLoginMessage{"zjxxxxxxxxxxxx","123qwe"}

		ReadWrite(tc,rqs)

		time.Sleep(6*time.Second)

}
func ReadWrite(conn net.Conn,rqs *basic.ReqLoginMessage){

	buf:=bytes.NewBuffer(nil)

	rqs.Write(buf)
	bs:=buf.Bytes()
	msglen:=len(bs)
	fmt.Println(bs)

	msgdata:=make([]byte,msglen+4)

	binary.BigEndian.PutUint32(msgdata[:4],uint32(msglen))

	copy(msgdata[4:],bs)
	fmt.Println(msgdata)
	fmt.Println(len(msgdata))
	conn.Write(msgdata)

}