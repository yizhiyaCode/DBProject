package main

import (
	"bean"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"resource"
	"strings"
)

type User bean.User

//Db数据库连接池
var DB *sql.DB

//初始化数据库连接池
func init() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{resource.UserName, ":", resource.Password, "@tcp(", resource.Ip, ":", resource.Port, ")/", resource.DbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)

	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("Connect success")
}

func InsertUser(user *User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO nk_user (`name`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.UserName, user.Password)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUser(user *User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM nk_user WHERE name = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}
func UpdateUser(user *User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE nk_user SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName, user.Password, user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func SelectUserById(id int) User {
	var user User
	err := DB.QueryRow("SELECT * FROM nk_user WHERE id = ?", id).Scan(&user.Id, &user.UserName, &user.Password)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return user
}

func SelectAllUser() []User {
	//执行查询语句
	rows, err := DB.Query("SELECT * from nk_user")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var users []User
	//循环读取结果
	for rows.Next() {
		var user User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

func main() {
	defer DB.Close()
	u := &User{Id: 123456, UserName: "yeshiva", Password: "12345678"}

	//InsertUser(u)

	DeleteUser(u)

	for _, value := range SelectAllUser() {
		fmt.Printf("%v\n", value)
	}

	fmt.Println("hello world")
}
