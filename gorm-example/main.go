package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type User struct {
	gorm.Model
	Username string `gorm:unique_index"`
	Email    string
	Age      int
}

var db *gorm.DB

func main() {
	db = initDb()

	//testCreate()
	//testUpdate()
	//testDelete()
	testQuery()

	defer db.Close()
}

// 查询 更多查询 API 查看：https://gorm.io/docs/query.html
func testQuery() {
	user := User{}

	// 查询第一个
	db.First(&user)
	fmt.Println(user)

	// 查询所有数据
	var users []User
	db.Find(&users)
	fmt.Println(users)

	// 模糊查询
	db.Where("username like ?", "%Ma%").Find(&users)
	fmt.Print(users)
}

func testDelete() {
	user := User{}
	user.ID = 1
	db.Delete(&user)
}

func testUpdate() {
	//db.Model(&User{}).Where("id=?", 1).Update("username", "Tony Ma")
	db.Model(&User{}).Where("id=?", 1).Updates(map[string]interface{}{
		"username": "Tony Ma Ma",
		"email":    "2@2.com",
		"age":      20,
	})
}

func testCreate() {
	user := User{Username: "6 Ma", Email: "1@1.com", Age: 18}
	db.Create(&user)
}

func initDb() *gorm.DB {
	if database, err := gorm.Open("mysql", "root:87660543@tcp(mysql.yy:3306)/gorm-example?charset=utf8&parseTime=True"); err == nil {
		database.SingularTable(true) // 取消复数形式的表名
		return database

	} else {
		panic(err)
		return nil
	}
}
