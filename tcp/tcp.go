package main

import (
	"net"
	"fmt"
	"io"
	"encoding/binary"
	"time"
	"bytes"
	"encoding/json"
)

type TmpMsg struct {
	Id int32
	Name string
	Score int32
}

func (m *TmpMsg) GetId() int32 {
	return 1
}

type Tmp2Msg struct {
	Id int32
	Name string

}


func (m *Tmp2Msg) GetId() int32 {
	return 2
}
type ServerResult struct {
	Result int32
	Name string
}

var byteOrder = binary.BigEndian

func main() {
	go func() {
		ln, err := net.Listen("tcp", ":9988")
		if err != nil {
			panic(err)
		}

		for {
			c, err := ln.Accept()
			if err != nil {
				fmt.Println("accept error", err)
				return
			}

			go func(conn net.Conn){
				defer conn.Close()

				/*
				msgLenData:是发送消息的长度
				读取之后就不会存在tcp发送消息的通道里 通道消息的长度就会减去读取的长度
				*/
				msgLenData := make([]byte, 4)
				for {

					_, err := io.ReadFull(conn, msgLenData)
					if err != nil {
						fmt.Printf("conn:%v read msg len err:%v\n", conn, err)
						return
					}

					msgLen := byteOrder.Uint32(msgLenData)
					if msgLen < 4 {
						fmt.Printf("conn:%v msgLen:%d < 4 data:%v\n", conn, msgLen, msgLenData)
						return
					}
					fmt.Println("recv msg Len:", msgLen)

					msgData := make([]byte, msgLen)
					fmt.Println("len:",msgLen)
					_, err = io.ReadFull(conn, msgData)
					if err != nil {
						fmt.Printf("conn:%v read msg data err:%v\n", conn, err)
						return
					}
					result:=0
					msgId := int32(byteOrder.Uint32(msgData[:4]))
					if msgId == 1 {
						msg := &TmpMsg{}
						msg.Id = int32(byteOrder.Uint32(msgData[4:8]))
						nameLen := byteOrder.Uint32(msgData[8:12])
						msg.Name = string(msgData[12:12+nameLen])
						msg.Score = int32(byteOrder.Uint32(msgData[12+nameLen:]))
						fmt.Printf("recv msg:%+v\n", msg)
						result=1
					} else if msgId == 2 {

					}

					if result==1 {
						r:=ServerResult{1,"ok"}
						b,err:=json.Marshal(r)
						if	err!=nil{
							panic(err)
						}
						a:=make([]byte,4+len(b))

						relen:=uint32(len(b))
						//PutUint32 uint32添加进数组
						byteOrder.PutUint32(a[:4], relen)
						copy(a[4:],b)
						fmt.Println("aaa",a)
						conn.Write(a)
					}


				}
			}(c)
		}
	}()

	time.Sleep(time.Second * 3)

	cli, err := net.Dial("tcp", "127.0.0.1:9988")
	if err != nil {
		panic(err)
	}

	msg := &TmpMsg{
		Id:100,
		Name:"xia",
		Score:59,
	}

	buf := bytes.NewBuffer(nil)

	err = binary.Write(buf, byteOrder, msg.GetId())
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, byteOrder, msg.Id)
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, byteOrder, int32(len(msg.Name)))
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, byteOrder, []byte(msg.Name))
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, byteOrder, msg.Score)
	if err != nil {
		panic(err)
	}

	//golang操作slice的中间库binary bytes io
	msgData := buf.Bytes()
	//binary.Write 添加的数据 len=19
	msgLen := uint32(len(msgData))
	// 4 ： 加上发送总长度 0 0 0 19
	msgSendData := make([]byte, 4 + msgLen)


	byteOrder.PutUint32(msgSendData[:4], msgLen)

	copy(msgSendData[4:], msgData)

	fmt.Printf("send data len:%d, data:%v\n", msgLen, msgSendData)
	cli.Write(msgSendData)



	t:=make([]byte,4)
	cliread,err:=cli.Read(t)
	resultlen:=byteOrder.Uint32(t)
	if err!=nil {
		panic(err)
	}
	 fmt.Println(resultlen,cliread)

	t2:=make([]byte,resultlen)

	_,err=cli.Read(t2)

	s:=ServerResult{}

	json.Unmarshal(t2,&s)

	if s.Result==1 {
		fmt.Println("发送成功并接受服务器反馈 ",s)
	}


	if err!=nil {
		panic(err)
	}


	time.Sleep(time.Second * 2)
}