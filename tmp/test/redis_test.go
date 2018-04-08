package test

import (
	"testing"
	"fmt"
	"os"
	"io"
	"sync"
	"time"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"os/signal"
	"bufio"
	"libraries/sqlAndLog/redisEngine"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"encoding/xml"
	"libs/leaf/log"
	"text/template"
	"net/http"
)

type PageStruct struct {
	Id         int
	Name       string
	RoleType   int
	ChName     string
	Apidata    []*Api_Config
	Createtime string
	CreateUser string
	PageId     int
}

type Api_Config struct {
	Name       string
	PageName   string
	ServerType int
}

type Redis struct {
	Addr        string `json:"addr"`
	RequirePass string
	MaxActive   int    `json:"maxActive"`
	MaxIdle     int    `json:"maxIdle"`
	IdleTimeOut int    `json:"idleTimeOut"`
}

type protocolCfg struct {
	SendId  int32   `json:"sendId"`
	RecvIds []int32 `json:"recvIds"`
}

type protocolConfig struct {
	Msgs []*protocolCfg `json:"msgs"`
}

func TestWrite2(t *testing.T) {

	fileName := "d:/test.txt"
	fi := &os.File{}
	var err error

	exit := true
	_, err = os.Stat(fileName)

	if os.IsNotExist(err) {
		exit = false
	}

	if exit {
		fi, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		fi, err = os.Create(fileName)
	}

	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			time.Sleep(500 * time.Millisecond)
			n, err := io.WriteString(fi, fmt.Sprintf("%d:%d\r\n", i, i))
			if err != nil {
				panic(err)
			}

			fmt.Printf("写入了 %d 个字节\n", n)
		}

	}()
	/*go func() {
		b, err := ioutil.ReadAll(fi)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" 文件内容 %s \n", b)
	}()*/
	wg.Wait()
}
func TestScan(t *testing.T) {

	a := "/group/vsda/ad/asd11f"
	index := strings.Index(a, "/")
	fmt.Println(index)
	if index == 0 {
		a = a[index+1:]
		fmt.Println(a)
	}
	arr := strings.Split(a, "/")
	fmt.Println(arr)
	var conAll string
	for i := 0; i < len(arr); i++ {
		con := arr[i][:1]
		conUp := strings.ToUpper(con)
		conNew := conUp + arr[i][1:]
		fmt.Println("con", conNew)
		conAll += conNew
	}

	fmt.Println(conAll)
}
func TestDel(t *testing.T) {
	mps := make(map[string]interface{})
	mps["tt"] = 1311
	t.Log(mps)
	delete(mps, "tt")
	t.Log(mps)
}
func TestDB(t *testing.T) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		"root",
		"123qWE",
		"localhost",
		3306,
		"adminserver");
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	r, err := d.Query("SELECT * from t_page")
	if err != nil {
		panic(err)
	}
	ps := make(map[string]*PageStruct)
	for r.Next() {
		apidata := ""
		p := &PageStruct{}
		err := r.Scan(&p.Id, &p.Name, &p.ChName, &p.RoleType, &apidata, &p.Createtime, &p.CreateUser, &p.PageId)
		if err != nil {
			panic(err)
		}

		p.Apidata = []*Api_Config{}

		err = json.Unmarshal([]byte(apidata), &p.Apidata)
		if err != nil {
			panic(err)
		}

		for _, v := range p.Apidata {
			fmt.Printf("%+v\n", v)
		}

		ps[p.Name] = p

	}

}
func TestClosing(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt)

	closing := make(chan bool)

	go func(closing chan bool) {
		defer wg.Done()
		select {
		case <-closing:
			fmt.Println("closing")
		}

	}(closing)

	<-ch
	close(closing)
	wg.Wait()

}
func TestReadFile(t *testing.T) {

	var seek int64
	seek = 0
LOOP:
	f, err := os.Open("d:/test.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	b := bufio.NewReader(f)

	for {

		f.Seek(seek, 1)

		s, err := b.ReadString('\n')

		if err != nil && err == io.EOF {
			fmt.Println("eof seek:", seek)
			time.Sleep(3 * time.Second)
			goto LOOP
		} else if (err != nil) {
			panic(err)
		}

		seek += 5
		fmt.Println("s:", s)
		fmt.Println("len(s)", len(s))
		fmt.Println("seek", seek)

	}

}
func TestRedis(t *testing.T) {

	redisEngine := &redisEngine.RedisEngine{}
	redisEngine.Start(nil,
		"192.168.1.232:6388",
		"dkkaI#27KlmQ-3k2OPj",
		50,
		300,
		3600)

	_, e := redisEngine.Request("SET", "test111", 2)

	a, e := redis.Int(redisEngine.Request("GET", "test111"))

	if e != nil {
		panic(e)
	}
	fmt.Println(a)

}
func TestJson(t *testing.T) {
	s := ` {
	"RequirePass": "192.168.1.232:6388",
	"addr": "192.168.1.232:6388",
	"maxIdle": 50,
	"maxActive": 300,
	"idleTimeOut": 3600
	}`

	r := Redis{}

	err := json.Unmarshal([]byte(s), &r)

	fmt.Println(err)

	fmt.Printf("%+v \n", r)
}
func TestWb1(t *testing.T) {

	p := protocolConfig{}

	f, err := os.Open("D:/yuwan/yuwan_server/protocol/protocol.xml")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}

	deskClient := make(map[int32]string)
	deskServer := make(map[int32]string)

	deskServer[100201] = "100101aa"

	deskServer[200202] = "100101aa"
	deskServer[218201] = "100101aa"

	deskClient[218101] = "100201aa"

	deskClient[213101] = "100201aa"
	deskClient[100101] = "100201aa"

	for k, _ := range deskServer {
		//flag := false
		for _, b := range p.Msgs {
			if k == b.SendId {
				for _, d := range b.RecvIds {
					for e, _ := range deskClient {
						if e == d {
							//flag = true
							delete(deskClient, e)
							delete(deskServer, k)
							break
						}
					}
				}
			}
		}

	}
	for k, _ := range deskServer {
		log.Release(" %d request time out", k)
	}
	log.Release("deskServer :%+v", deskServer)
	log.Release("deskclient :%+v", deskClient)

}
func Test6(t *testing.T) {

	n := time.Now()
	a := int64(n.Unix() / 60)
	fmt.Println(a)
	tu := time.Unix(a*60, 0).Format("2006-01-02 15:04:05")
	fmt.Println(tu)

}
func TestTempLateDemo(t *testing.T) {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		tp := template.Template{}
		tpnew := tp.New("demo1")
		tp2, err := tpnew.Parse(`
		{{.Name}}
		`)
		if err != nil {
			panic(err)
		}

		a := Api_Config{Name: "zjx"}

		err = tp2.Execute(w, a)

		if err != nil {
			panic(err)
		}

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
func TestMysqlDemo(t *testing.T){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		"root",
		"123qWE",
		"localhost",
		3306,
		"test");
	d, err := sql.Open("mysql", dsn)

	r,err :=d.Query("select  a.i_d,a.account,a.pass_word,article_text  from USER a inner join article ar on a.i_d= ar.i_d ")
	if err != nil {
		panic(err)
	}

	if r.Next() {
		var id int
		var account string
		var pwd sql.NullString
		var age string
		err := r.Scan(&id,&account,&pwd,&age)
		if err != nil {
			panic(err)
		}
		fmt.Println(account,pwd,age)
	}
}

