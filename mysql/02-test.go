package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
//
//func initDB() (err error) {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// 注意此处不用：:=
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func main() {
//	err := initDB()
//	if err != nil {
//		fmt.Println("init db failed, err:", err)
//		return
//	}
//
//}
