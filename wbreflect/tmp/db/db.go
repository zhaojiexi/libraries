package db

import (
	"cx/server/common/mysqlEngine"
	"libs/leaf/log"
	"libs/leaf/go"
	"net/http"
	"fmt"
)

type DB struct {
	M *mysqlEngine.MysqlEngine
}

func (db *DB) Start() {

	g := g.New(10)

	db.M = &mysqlEngine.MysqlEngine{
		ConfPath:  "./db/mysql_eng.json",
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Passwd:    "123qWE",
		DbName:    "dfh7",
		MaxIdle:   1024,
		MaxActive: 1024,
		GoSer:     g,
	}
	err := db.M.Start()

	if err != nil {
		panic(err)
	}

}
func SelectAdmin(key interface{}, ret *mysqlEngine.DBRet,w http.ResponseWriter, r *http.Request) {
	if ret.Err != nil {
		log.Release(" err %s", ret.Err)
	}
	log.Release("key", key, "ret :", ret.Rows)
	if ret.Rows == nil {
		log.Release(" Rows nil %s")
		return
	}

	for ret.Rows.Next() {
		var id int
		var s, name, pwd string
		err := ret.Rows.Scan(&id, &s, &name, &pwd)
		if err != nil {
			panic(err)
		}
		log.Release("id=%d,s=%s,name=%s,pwd=%s",id,s,name,pwd)
		fmt.Fprint(w,id,s,name,pwd)
	}

}