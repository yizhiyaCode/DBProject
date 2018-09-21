package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"resource"
	"strings"
	"time"
)

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	CreatedAt time.Time
}

var db *gorm.DB

func init() {
	var err error
	path := strings.Join([]string{resource.UserName, ":", resource.Password, "@tcp(", resource.Ip, ":", resource.Port, ")/", resource.DbName, "?charset=utf8"}, "")

	db, err = gorm.Open("mysql", path)

	db.DB().SetConnMaxLifetime(100)
	db.DB().SetMaxIdleConns(10)
	//验证连接
	if err != nil {
		fmt.Println("open database fail")
		return
	}

	create(Like{})
}

//创建表
func create(like Like) bool {
	if !db.HasTable(&like) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&like).
			Error; err != nil {
			return false
		}
	}
	return true
}

//插入
func insert(like *Like) bool {

	//通过db直接创建
	//if err := db.Create(like).Error; err != nil {
	//	return false
	//}

	tx := db.Begin()
	if err := tx.Create(like).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()

	return true
}

//删除
func delete(id int) bool {

	if err := db.Where(&Like{ID: id}).Delete(Like{}).Error; err != nil {
		return false
	}
	return true
}

//更新
func updateTitle(like *Like, name string) bool {

	db.Model(&like).Update("name", name)

	return true
}

//查询一条
func selectById(id int) (like []Like) {

	//db.Where("ID = ?", id).First(&like)

	//查询id相同的所有
	db.Where("ip = ?", "localhost").Find(&like)
	return like
}

//查询所有
func selectAll() (like []Like) {

	db.Find(&like)

	return like
}

func main() {
	//like := &Like{
	//	Ip:        "ip",
	//	Ua:        "ua",
	//	Title:     "hello world ",
	//	CreatedAt: time.Now(),
	//}
	//
	//insert(like)

	//like := selectById(1)
	//j,_ :=json.MarshalIndent(like,"","   ")
	//fmt.Println(string(j))
	//like := selectById(1)
	//for i, value := range like {
	//	j, _ := json.MarshalIndent(value, "", "   ")
	//	fmt.Printf(" like %d  %s \n", i+1, string(j))
	//}

	fmt.Printf("hello world")
}
