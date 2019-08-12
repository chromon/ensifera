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
		fmt.Println("open database fail:", err)
		return
	}

	fmt.Println("connect success")
}

// 插入
func Insert() bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("transaction open fail")
		return false
	}
	// 预处理插入数据
	stmt, err := tx.Prepare("INSERT INTO test VALUES (0, ?, ?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("prepare statement fail: ", err)
		return false
	}
	// 执行 sql 语句
	res, err := stmt.Exec("pear", "1234")
	if err != nil {
		fmt.Println("exec fail:", err)
		return false
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("transaction roll back error:", err.Error())
			return false
		}
	}
	// 获得插入后的自增 id
	lastInsertId, err := res.LastInsertId()
	if err == nil {
		fmt.Println("last inserted id:", lastInsertId)
	}
	// 影响的行数
	rowsAffected, err := res.RowsAffected()
	if err == nil {
		fmt.Println("rows affected:", rowsAffected)
	}

	return true
}

// 删除数据
func Delete() bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("transaction open fail:", err)
		return false
	}
	//预处理数据
	stmt, err := tx.Prepare("DELETE FROM test WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Println("prepare statement fail:", err)
		return false
	}
	// 执行 sql 语句
	res, err := stmt.Exec(3)
	if err != nil {
		fmt.Println("exec fail:", err)
		return false
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("transaction roll back error:", err.Error())
			return false
		}
	}
	// 影响的行数
	rowsAffected, err := res.RowsAffected()
	if err == nil {
		fmt.Println("delete rows affected:",rowsAffected)
	}

	return true
}

// 更新
func Update() bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("transaction open fail:", err)
		return false
	}
	// 预处理数据
	stmt, err := tx.Prepare("UPDATE test SET name = ? WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Println("prepare statement fail:", err)
		return false
	}
	// 执行 sql 语句
	res, err := stmt.Exec("apple", 2)
	if err != nil {
		fmt.Println("exec fail:", err)
		return false
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("transaction roll back error:", err.Error())
			return false
		}
	}
	// 影响的行数
	rowsAffected, err := res.RowsAffected()
	if err == nil {
		fmt.Println("delete rows affected:", rowsAffected)
	}

	return true
}

// 查询
func Query() {
	var id int
	var name, password string

	// 执行查询语句
	rows, err := DB.Query("SELECT * FROM test")
	if err != nil {
		fmt.Println("query fail:", err)
	}
	// 循环遍历结果
	for rows.Next() {
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			fmt.Println("scan fail:", err)
		}
		fmt.Println("id=", id, ", name=", name, ", password=", password)
	}
	defer rows.Close()
}

// 查询一行
func QueryRow() {
	var id int
	var name, password string
	// 执行查询语句
	row := DB.QueryRow("SELECT * FROM test WHERE id = ?", 2)
	err := row.Scan(&id, &name, &password)
	if err != nil {
		fmt.Println("scan fail:", err)
	}
	fmt.Println("id=", id, ", name=", name, ", password=", password)
}