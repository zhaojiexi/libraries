package main

import (
	"libs/ywnet"
	"time"
	cptl "cx/server/common/protocol"
	"libs/leaf/chanrpc"
	"cx/server/protocol"
	"github.com/funny/link"
	"fmt"
	"libs/leaf/log"
)

func main() {
	/*	ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill)*/

	client := &ywnet.Client{
		Network:     "tcp",
		Addr:        "127.0.0.1:8080",
		DialTimeOut: time.Second * 10,

		ConnectKey: "connect",
		ClosedKey:  "closed",
		RecvKey:    "recv",
		ChanServer: chanrpc.NewServer(100),
		ConnectNum: 1,

		Protocol: cptl.Protocol(),
	}
	client.ChanServer.Register("connect", func(args []interface{}) {

		p := protocol.PingServerMessage{}
		p.ClientSessionId = 979797
		p.Updown = 1

		session := args[0].(*link.Session)
		session.Send(p)
		log.Release("send msg %+v",p)

	})
	client.ChanServer.Register("recv", func(args []interface{}) {
		p := args[1].(protocol.PingServerMessage)

		fmt.Printf(" recv %+v \n", p)

	})
	client.ChanServer.Register("closed", func(args []interface{}) {
		log.Release("closed !!!!!!!!!")

	})

	client.Connect()

	for {
		select {
			case a := <-client.ChanServer.ChanCall:
			client.ChanServer.Exec(a)
		/*	case <-ch :
				log.Release("收到中断信号 ")
				os.Exit(100)*/
		}
	}

}
