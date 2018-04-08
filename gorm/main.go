package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int       `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	Product   []Product `gorm:"foreignKey:UserID;AssociationForeignKey:ID"`
	HuiYuan   HuiYuan	`gorm:"foreignKey:HuiYuanID;AssociationForeignKey:ID"`
	HuiYuanID int
}
type Product struct {
	ID          int    `gorm:"primary_key;column:id"`
	ProductName string `gorm:"column:name"`
	UserID      int
}
type HuiYuan struct {
	ID  int    `gorm:"column:id;primary_key"`
	Pwd string `gorm:"column:pwd"`
}

func (User) TableName() string {
	return "user"
}
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return nil
}
func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		"root",
		"123qWE",
		"localhost",
		3306,
		"test")
	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)
	//db.Set("gorm:table_options","charset=utf-8")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Product{}, &HuiYuan{})

	u := User{Name: "zjx222222222222222222"}
	p := Product{ProductName: "zzzz"}
	h := HuiYuan{Pwd: "123"}
	u.HuiYuan = h

	u.Product = append(u.Product, p)
	u.Product = append(u.Product, p)
	u.Product = append(u.Product, p)
	db.Create(&u)

	u2 := User{ID: 1}

	db.Find(&u2)
	db.Model(&u2).Related(&u2.HuiYuan)
	fmt.Printf("%+v\n", u2)

	/*db.Model(&u).Update(&User{Name: "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"})
	if err != nil {
		panic(err)
	}
*/

}
