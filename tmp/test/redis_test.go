package test

import (
	"testing"
	"libreries/tmp/db"
	"fmt"
	"os"
	"io"
	"sync"
	"time"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
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

func TestRedis1(t *testing.T) {
	r := db.DBC{}
	r.GetDB()
	c := r.Redis.Get()
	defer c.Close()

	val, err := c.Do("Lpush", "testlist", "0", "100")
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
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
