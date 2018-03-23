package main

import (
	"libraries/tmp/file"
	"strings"
	"fmt"
	"time"
	"os"
	"os/signal"
	"libraries/tmp/db"
	"github.com/garyburd/redigo/redis"
)

const (
	dataname = "datalist"
)

var datalen int

type MyUser struct {
	UserName  string
	UserPhone string
}

func main() {
	datalen := Datalen().(int)

	var a []string

	fi := file.File{}
	si := make(chan os.Signal, 1)
	signal.Notify(si, os.Interrupt)
	ch := make(chan string, 1)
	go func() {
		for {
			time.Sleep(4 * time.Second)
			s := fi.Read()
			ch <- s
		}
	}()

	go func() {
		for {

			select {
			case s := <-ch:

				fmt.Println("a")
				slist := strings.Split(s, "\n")

				listlen := len(slist)

				//if datalen == 0 {
				//	adddatalen()
				//}

				if listlen > datalen {
					diff := listlen - datalen
					fmt.Printf("%d,%d发现新增数据 %d条 \n", listlen, datalen, listlen-datalen)
					for i := diff; i > 0; i-- {

						index := strings.Index(slist[listlen-i], ":")

						fmt.Println("新增数据！！！", slist[listlen-i][index+1:])
						data := slist[listlen-i][index+1:]
						addlist(data)

					}
					adddatalen(listlen)
					datalen = listlen
				}

				for _, v := range slist {
					index := strings.Index(v, ":")

					a = append(a, v[index+1:])

				}
				for _, v := range a {
					fmt.Println(v)
				}



			case <-si:
				os.Exit(1)
			}
		}
	}()

	<-si
	fmt.Println("主动关闭")

}

func Datalen() interface{} {
	r := db.DBC{}
	r.GetDB()

	c := r.Redis.Get()
	defer c.Close()

	i, err := redis.Int(c.Do("GET", "datalen"))
	if err != nil {
		panic(err)
	}

	return i
}
func adddatalen(i int) {

	r := db.DBC{}
	r.GetDB()

	c := r.Redis.Get()
	defer c.Close()

	val, err := c.Do("SET", "datalen", i)

	if err != nil {
		panic(err)
	}

	fmt.Println(val)

}

func addlist(s string) {

	r := db.DBC{}
	r.GetDB()

	c := r.Redis.Get()
	defer c.Close()

	v, err := c.Do("lpush", dataname, s)
	if err != nil {
		panic(err)
	}
	fmt.Println("v", v)
}
