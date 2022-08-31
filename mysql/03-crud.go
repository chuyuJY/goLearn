package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库
var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意此处不用：:=
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 1. 查询数据
// 先定义一个结构体，用于接收数据
type User struct {
	id       int
	username string
	password string
}

// 1.1 查询单行数据
func queryOneRow() {
	s := "select * from user_tb1 where id = ?"
	user := User{}
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(s, 1).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		fmt.Println("queryOneRow failed, err:", err)
		return
	}
	fmt.Println("query result:", user)
}

// 1.2 多行查询
func queryRows() {
	s := "select * from user_tb1 where id > ?"
	rows, err := db.Query(s, 0) // 取出 id > 0 的 row
	if err != nil {
		fmt.Println("queryRows failed, err:", err)
		return
	}
	defer rows.Close()
	// 循环读取
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.id, &user.username, &user.password)
		if err != nil {
			fmt.Println("scan failed, err:", err)
			return
		}
		fmt.Println("query result:", user)
	}
}

// 2. 插入数据
func insert() {
	s := "insert into user_tb1 (username, password) values(?, ?)"
	r, err := db.Exec(s, "jerry", "789")
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return
	}
	theId, _ := r.LastInsertId() // 新插入数据的id
	fmt.Printf("insert %v success...", theId)
}

// 3. 更新数据
func updateRow() {
	s := "update user_tb1 set username=?, password=? where id=?"
	ret, err := db.Exec(s, "chuu", "0404", 1)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get rowAffected failed, err:", err)
		return
	}
	fmt.Println("update success, affected rows:", n)
}

// 4. 删除数据
func deleteRow() {
	s := "delete from user_tb1 where id=?"
	ret, err := db.Exec(s, 2)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get rowAffected failed, err:", err)
		return
	}
	fmt.Println("delete success, affected rows:", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init db failed, err:", err)
	} else {
		fmt.Println("连接数据库成功...")
	}
	//insert()
	//queryOneRow()
	//queryRows()
	//updateRow()
	deleteRow()
}
