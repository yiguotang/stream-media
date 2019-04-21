package dbops

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 新增用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

// 根据用户名查询用户密码
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("delete user error: %s", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
