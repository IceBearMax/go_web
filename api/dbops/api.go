package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCreatential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCreatential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeletUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		log.Printf("DeleteUser error :%s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil

}
