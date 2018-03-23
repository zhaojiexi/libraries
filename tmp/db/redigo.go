package db

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"libs/leaf/log"
	"time"
)

type DBC struct {
	Redis *redis.Pool
}

var (
	db             DBC
	Redis_Password = "Zxcasdqwe123!@#"
	Redis_Host     = "127.0.0.1"
	Redis_Port     = 6379
)



func (db *DBC)GetDB() {

	r := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 60 * 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			var con redis.Conn
			var err error

			if len(Redis_Password) > 0 {
				pwd := redis.DialPassword(Redis_Password)
				con, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", Redis_Host, Redis_Port), pwd)
			} else {
				con, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", Redis_Host, Redis_Port))
			}
			if err != nil {
				log.Release("conn err ", err)
				return con, err
			}
			return con, nil
		},
	}

	db.Redis = r

}
