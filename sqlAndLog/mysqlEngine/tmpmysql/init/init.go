package init

import (
	"encoding/json"
	"io/ioutil"
	"cx/server/tmpmysql/global"
)

type MysqlConfig struct {
	LogLevel string
	LogPath string
	Mysql Mysql

}
type Mysql struct{
	Host string
	Port int
	User string
	Passwd string
	DbName string
	MaxIdle int
	MaxActive int
}
var MqlConfig *MysqlConfig

func init(){

	b,err:=ioutil.ReadFile(global.Config_File)
	if err != nil {
		panic(err)
	}

	err=json.Unmarshal(b,&MqlConfig)
	if err != nil {
		panic(err)
	}


}