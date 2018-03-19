package basic

import (
	"encoding/binary"
	"bytes"
	"libs/leaf/log"
)

type ReadWrite interface {
	Write()
	GetID()int
}

type ReqLoginMessage struct {
	Name string
	Pwd string
}
func(w ReqLoginMessage)GetID()int32{
	return 1
}

func(w *ReqLoginMessage)Write(buf *bytes.Buffer){

	err:=binary.Write(buf,binary.BigEndian,w.GetID())
	if	err!=nil{
		log.Release("ReqLoginMessage Write Name err %s",err)
	}

	err=binary.Write(buf,binary.BigEndian,[]byte(w.Name))
	if	err!=nil{
		log.Release("ReqLoginMessage Write Name err %s",err)
	}
	err=binary.Write(buf,binary.BigEndian,[]byte(w.Pwd))
	if	err!=nil{
		log.Release("ReqLoginMessage Write Name err %s",err)
	}


}

