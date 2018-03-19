package mysqlEngine

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"fmt"
	"libs/leaf/go"
	"io/ioutil"
	"encoding/json"
	"context"
	"libs/leaf/log"
)

const (
	CONF_FILE = "db_engine.json"
)

type MysqlEngine struct {
	ConfPath string
	Host string
	Port int
	User string
	Passwd string
	DbName string
	MaxIdle int
	MaxActive int
	WorkerNum int
	GoSer *g.Go

	db *sql.DB
	engineConf DBConf
	smtsById map[int]*stmtItem
	smtsByName map[string]*stmtItem
	ctx context.Context
	ctx_cb context.CancelFunc
}

type stmtItem struct {
	*sql.Stmt
	*DBConfItem
}

type DBConfItem struct {
	Id int
	Name string
	Sql string
	Select bool
}

type DBConf struct {
	Sqls []*DBConfItem
}

func(me *MysqlEngine) Start() (err error) {
	if me.GoSer == nil {
		err = fmt.Errorf("mysqlEngine must has GoSer")
		return
	}

	if len(me.ConfPath) == 0 || me.ConfPath[len(me.ConfPath) -1] == '/' {
		me.ConfPath += CONF_FILE
	}


	data, err := ioutil.ReadFile(me.ConfPath )
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &me.engineConf)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		me.User,
		me.Passwd,
		me.Host,
		me.Port,
		me.DbName);

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(me.MaxIdle)
	db.SetMaxOpenConns(me.MaxActive)
	me.db = db

	me.smtsById = make(map[int]*stmtItem)
	me.smtsByName = make(map[string]*stmtItem)

	for _, si := range me.engineConf.Sqls {
		smt, err := me.db.Prepare(si.Sql)
		if err != nil {
			return err
		}

		if _, ok := me.smtsById[si.Id];ok {
			err = fmt.Errorf("dbengin id:%d repeat", si.Id)
			return err
		}

		if _, ok := me.smtsByName[si.Name];ok {
			err = fmt.Errorf("dbengin name:%s repeat", si.Name)
			return err
		}

		sti := &stmtItem{smt, si}
		me.smtsById[si.Id] = sti
		me.smtsByName[si.Name] = sti
	}
	me.ctx, me.ctx_cb = context.WithCancel(context.Background())

	return
}

func (me *MysqlEngine) Stop() {
	me.ctx_cb()

	for _, smt := range me.smtsById {
		smt.Close()
	}

	me.db.Close()
}

type DBRet struct {
	Err error
	Rows *sql.Rows
	Result sql.Result
}

type IDBSink interface {
	OnRet(key interface{}, ret *DBRet)
}

type IDBSinkFunc func(key interface{}, ret *DBRet)
func(s IDBSinkFunc) OnRet(key interface{}, ret *DBRet) {
	s(key, ret)
}

func (me *MysqlEngine) Request(reqId int, key interface{}, sink IDBSink, args ...interface{}) (err error) {
	sti, ok := me.smtsById[reqId]
	if !ok {
		return fmt.Errorf("invalid req id:%d", reqId)
	}

	dbRet := &DBRet{
	}


	gf := func(){
		if sti.Select {
			dbRet.Rows, dbRet.Err = sti.Stmt.QueryContext(me.ctx, args...)
		} else {
			dbRet.Result, dbRet.Err = sti.Stmt.ExecContext(me.ctx, args...)
		}
	}

	cb := func(){
		defer func() {
			if dbRet.Err != nil {
				if dbRet.Rows != nil {
					dbRet.Rows.Close()
				}
			}
		}()

		if sink != nil {
			sink.OnRet(key, dbRet)
		} else {
			if dbRet.Err != nil {
				log.Error("mysqlEngine Request err:%v", dbRet.Err)
			}
		}
	}

	me.GoSer.Go(gf, cb)
	return
}