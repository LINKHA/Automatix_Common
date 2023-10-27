package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
CREATE TABLE IF NOT EXISTS person (

	user_id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	sex VARCHAR(10) NOT NULL,
	email VARCHAR(255) NOT NULL,
	PRIMARY KEY (user_id)

);
*/
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

/*
CREATE TABLE IF NOT EXISTS account_config (

	account VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL

);
*/
type AccountConfig struct {
	Account  string `db:"account"`  //账号
	Password string `db:"password"` //密码
}

var Db *sqlx.DB

func Init(mongoUrl string, setModelMap map[string]IBaseModel) {
	//sqlx.Open(数据库类型, "账号:密码@tcp(服务器IP:端口)/数据库")
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3307)/orm_test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close() // 注意这行代码要写在上面err判断的下面

	for k, v := range setModelMap {
	}
}

func Stop() {
	defer func() {
		Db.Close()
	}()
}

func Get(model_name string) {
	return ModelMap[model_name]
}

// func main() {
// 	// findDemo()
// 	selectDemo()
// 	// insertDemo()
// 	// updateDemo()
// 	// deleteDemo()
// }

/**
 * [findDemo 查询单条]
 * Get方法-结构体存储
 * Db语句: Db.Get(结构体指针, sql语句, 占位符值)
 * @return {[type]} [description]
 */
func findDemo() {
	var person Person
	err := Db.Get(&person, "select uid,username,nickname,sex from chat_members where uid=?", 6)
	if err != nil {
		fmt.Println("查询失败, ", err)
		return
	}

	fmt.Println("Get succ:", person)
}

/**
 * [selectDemo 查询多条数据]
 * Select方法--切片存储
 * Db语句: Db.Select(结构体指针, sql语句, 占位符值)
 * @return {[type]} [description]
 */
func selectDemo() {
	var account_config []AccountConfig
	err := Db.Select(&account_config, "select * from account_config where account=?", "account_test")
	if err != nil {
		fmt.Println("查询失败, ", err)
	}
	fmt.Println("Select Data:", account_config)
}

/**
 * [insertDemo 新增]
 * Db语句: Db.Exec(sql添加语句, 占位符值)
 *         LastInsertId() -- 获取新增ID
 * @return {[type]} [description]
 */
func insertDemo() {
	res, err := Db.Exec("insert into account_config(account, password) value(?, ?)", "account_test", "password_test")
	if err != nil {
		fmt.Println("新增失败, ", err)
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("获取新增ID失败, ", err)
	}

	fmt.Println("InsertID: ", insertID)
}

/**
 * [updateDemo 修改]
 * Db语句: Db.Exec(sql修改语句, 占位符值)
 *         RowsAffected() -- 获取受影响行数
 * @return {[type]} [description]
 */
func updateDemo() {
	res, err := Db.Exec("update chat_test set nickname=? where id=?", "昵称修改", 1)
	if err != nil {
		fmt.Println("修改失败, ", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响行数失败, ", rowsAffected)
	}
	fmt.Println("rowsAffected: ", rowsAffected)
}

/**
 * [deleteDemo 删除]
 * Db语句: Db.Exec(sql删除语句, 占位符值)
 *         RowsAffected() -- 获取受影响行数
 * @return {[type]} [description]
 */
func deleteDemo() {
	res, err := Db.Exec("delete from chat_test where id = ?", 2)
	if err != nil {
		fmt.Println("删除失败, ", err)
	}
	deleteNumber, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响行数失败", err)
	}
	fmt.Println("deleteNumber: ", deleteNumber)
}
