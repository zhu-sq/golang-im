package db

import (
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserInfo(account string) (*UInfo, error) {
	uInfo := &UInfo{}
	sql := "SELECT * FROM user WHERE account = ?"
	rows, err := db.Query(sql, account)
	if err != nil {
		log.Print(account, " query  failed ", err.Error())
		return uInfo, err
	}
	defer rows.Close()
	var birthdayTime time.Time
	for rows.Next() {
		err := rows.Scan(&uInfo.UID, &uInfo.Account, &uInfo.Pwd, &uInfo.Email, &uInfo.Phone, &uInfo.Nickname, &uInfo.Icon, &uInfo.Sex, &birthdayTime, &uInfo.Addr)
		if err != nil {
			log.Print(account, " query failed ", err.Error())
			return uInfo, err
		}
		uInfo.Birthday = birthdayTime.Unix()
		return uInfo, nil
	}
	return uInfo, errors.New("账号不存在")
}

func SetUserInfo(uInfo *UInfo) error {

	if uInfo.Account == "" || uInfo.Pwd == "" {
		log.Print("空的账号或密码")
		return errors.New("空的账号或密码")
	}

	if uInfo.Sex == "" {
		uInfo.Sex = "保密"
	}

	temBirthTime := time.Unix(uInfo.Birthday, 0)

	sql := "insert into user(account,email,phone,pwd,nickname,icon,sex,birthday,addr) values(?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, uInfo.Account, uInfo.Email, uInfo.Phone, uInfo.Pwd, uInfo.Nickname, uInfo.Icon, uInfo.Sex, temBirthTime, uInfo.Addr)
	if err != nil {
		log.Print(uInfo.Account, " insert data failed ", err.Error())
		return err
	}
	return nil
}
