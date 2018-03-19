package main

import (
	"cx/server/tmpmysql/tmp"
	"libs/leaf/log"
	in"cx/server/tmpmysql/init"
	syslog"log"
)

func main() {
	logger,err:=log.New(in.MqlConfig.LogPath,1024*1024*200,in.MqlConfig.LogLevel,syslog.Lshortfile|syslog.LstdFlags)
	if err != nil {
		panic(err)
	}

	log.Export(logger)
	log.Release("begin")

	tmp.MRZ.Test()
}
