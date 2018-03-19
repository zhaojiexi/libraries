package tmp

import (
	mql "cx/server/common/mysqlEngine"
	in"cx/server/tmpmysql/init"
	"cx/server/tmpmysql/global"
	gog"libs/leaf/go"
	"libs/leaf/log"
	"time"
	"database/sql"
)

type Mrz interface {
	Test()
}

var(
	 MRZ Mrz
)

func init(){
	MRZ=&mrz{}
}

type mrz struct{
	name string
}

type User2 struct{
	id int
	name string
	birthday time.Time
}



func (m *mrz)Test(){
	servergo:=gog.New(100)

	mysql := &mql.MysqlEngine{
		ConfPath:global.Mysql_Engine_Conf_Path,
		Host:in.MqlConfig.Mysql.Host,
		Port:in.MqlConfig.Mysql.Port,
		User:in.MqlConfig.Mysql.User,
		Passwd:in.MqlConfig.Mysql.Passwd,
		DbName:in.MqlConfig.Mysql.DbName,
		MaxIdle:in.MqlConfig.Mysql.MaxIdle,
		MaxActive:in.MqlConfig.Mysql.MaxActive,
		GoSer:servergo,
	}
	err:=mysql.Start()
	if err != nil {
		panic(err)
	}

	err=mysql.Request(1,nil,mql.IDBSinkFunc(selectt1))
	if err != nil {
		panic(err)
	}
	mysql.Request(2,nil,mql.IDBSinkFunc(func(ket interface{},ret *mql.DBRet) {
		a,_:=ret.Result.RowsAffected()
		log.Release("***********",a)

	}),1,time.Now())


	go func(){for{
		select{
		case a:=<-mysql.GoSer.ChanCb:
		a()
	}
	}}()

	time.Sleep(1*time.Second)
}



func selectt1(key interface{}, dbRet *mql.DBRet) {

	if dbRet.Err != nil {
		log.Error("MYSQL_ReqID_TEST_SELECT1 err:%v", dbRet.Err)
		return
	}


	var id int
	var name sql.NullString
	var birthday time.Time
	cnt := 0
	for dbRet.Rows.Next() {
		err := dbRet.Rows.Scan(&id,&name,&birthday)
		log.Release("MYSQL_ReqID_TEST_SELECT1 scan id:%d name:%s birthday:=%+v, err:%v", id,name.String,birthday,err)
		if err == nil {
			cnt += 1
		}
	}

	if cnt == 0 {
		log.Release("MYSQL_ReqID_TEST_SELECT1 no rows ???")
	}

}







