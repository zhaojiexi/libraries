package main

import (
	ht "libraries/wbreflect/tmp/http"
	"libs/leaf/log"
	lg "log"
)

var (
	LogLevel   = "debug"
	LogFlag    = lg.LstdFlags | lg.Lshortfile
	LogMaxSize = 1024 * 1024 * 200 //日志文件最大100M
)

func main() {

	logger, err := log.New("./log/wbreflect.log", LogMaxSize, LogLevel, LogFlag)

	if err != nil {
		panic(err)
	}
	log.Export(logger)
	defer logger.Close()
	log.Release("==========%+v", err)

	err = ht.Start()

	if err != nil {
		log.Error(" http start err:%s", err)
	}

}
