package db

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "root"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "test"
)

// 数据库连接池
var DB *sql.DB

// 初始化
func InitDB() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	// root:123456@tcp(127.0.0.1:3306)/test
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	// 打开数据库
	DB, _ = sql.Open("mysql", path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}

	fmt.Println("connect success")
}

// 插入
func Insert() bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("transaction fail")
		return false
	}

	stmt, err := tx.Prepare("INSERT INTO test VALUES (0, ?, ?)")
	//defer stmt.Close()
	if err != nil {
		fmt.Println("statement fail: ", err)
		return false
	}

	// 执行 sql 语句
	res, err := stmt.Exec("pear", "1234")
	if err != nil {
		fmt.Println("exec fail")
		return false
	}

	// 提交事务
	tx.Commit()

	// 获得插入后的自增 id
	fmt.Println(res.LastInsertId())

	return true
}
