package redisEngine

import (
	"github.com/garyburd/redigo/redis"
	"errors"
	"libs/leaf/module"
	"time"
)

type CRedisRet struct {
	PlayerId int
	OprId int
	Reply interface{}
	Err  error
}

type IRedisSink interface {
	OnRet(*CRedisRet)
}

type RedisEngine struct {
	Skeleton *module.Skeleton
	RedisPool *redis.Pool
}

func(re *RedisEngine) Start(sk *module.Skeleton, dbAddr string, requirePass string, maxIdle int, maxActive int, IdleTimeOut time.Duration) error {
	re.Skeleton = sk
	re.RedisPool = &redis.Pool{
		MaxIdle:maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: IdleTimeOut,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			if requirePass != "" {
				dialInfo := redis.DialPassword(requirePass)
				c, err = redis.Dial("tcp", dbAddr, dialInfo)
			} else {
				c, err = redis.Dial("tcp", dbAddr)
			}

			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return nil
}

func(re *RedisEngine) Stop() {
	re.RedisPool.Close()
}

func (re *RedisEngine) GetRedisPool() *redis.Pool {
	return re.RedisPool
}

func(re *RedisEngine) Request(cmd string, args... interface{}) (interface{}, error) {
	if re.RedisPool == nil {
		return nil, errors.New("redis engine not start")
	}

	conn := re.RedisPool.Get()
	defer conn.Close()

	return conn.Do(cmd, args...)
}

func(re *RedisEngine) AsyncRequest(oprId int, playerId int, sink IRedisSink, cmd string, args... interface{}) error {
	ret := &CRedisRet{
		PlayerId:playerId,
		OprId:oprId,
	}

	re.Skeleton.Go(func(){
		ret.Reply, ret.Err = re.Request(cmd, args...)
	}, func() {
		sink.OnRet(ret)
	})

	return nil
}