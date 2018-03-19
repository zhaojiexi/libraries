package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"

	"fmt"

	"time"
)

type User2 struct{

	Id int	`orm:"auto"`
	Name string `orm:"size(100)"`
	Birthday time.Time  `orm:"type(datetime)"`
	ProFile *ProFile `orm:"rel(one)"`
	Post []*Post `orm:"reverse(many)"`
}
type ProFile struct{
	Id int
	Age int
	user *User2`orm:"reverse(one)"`

}

type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User2  `orm:"rel(fk)"`
}

type User3 struct{
	Name string
	Age int
}

const(
	dbUname="root"
	dbPwd="123qWE"
	dbAddr="127.0.0.1:3306"
	dbName="test"
)

func init(){
	orm.RegisterDataBase("default","mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",dbUname,dbPwd,dbAddr,dbName),30)

	orm.RegisterModel(new(ProFile))
	orm.RegisterModel(new(User2))
	orm.RegisterModel(new(Post))

	orm.RunSyncdb("default",false,true)
}
func main() {
	p1:=new(Post)
	p2:=new(Post)
	p1.Title="ttt"
	p2.Title="ttxxxxxxt"
	p:=[]*Post{
		p1,
		p2,
	}

	pr:=ProFile{Age:11}


	u:=User2{Name:"test",Post:p,ProFile:&pr,Birthday:time.Now()}
	orm.Debug=true

	o:=orm.NewOrm()

	o.Begin()

	_,err:=o.Insert(&u)
	_,err=o.Insert(&pr)
	if err!=nil {
		fmt.Println("err",err)
	}

	o.Commit()

	au:=&User2{}
	q:=o.QueryTable("user2")

	q.Filter("ProFile",3).RelatedSel().All(au)
	//fmt.Println(au.ProFile.Age)


	qb,err:=orm.NewQueryBuilder("mysql")
	if	err!=nil {
		fmt.Println("err",err)
	}

	qb.Select("user2.name","pro_file.age").From("user2").InnerJoin("pro_file").On("user2.pro_file_id=pro_file.id")
	sql:=qb.String()

	au1:=[]User3{}
	o.Raw(sql).QueryRows(&au1)

	for i:=0;i<len(au1) ;i++  {
		fmt.Println(au1[i])
	}
}
