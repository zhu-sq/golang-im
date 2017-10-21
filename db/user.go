package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserInfo(account string) (*UInfo, error) {
	uInfo := &UInfo{}
	sql := "SELECT pwd FROM user WHERE account = ?"
	rows, err := db.Query(sql, account)
	if err != nil {
		log.Fatal(account, " query pwd failed ", err.Error())
		return uInfo, err
	}

	for rows.Next() {
		if err := rows.Scan(uInfo); err != nil {
			log.Fatal(account, "query pwd failed", err.Error())
			return uInfo, err
		}
	}
	return uInfo, nil
}

func SetUserInfo(uInfo *UInfo) error {
	if uInfo.Sex == "" {
		uInfo.Sex = "保密"
	}

	// if uInfo.Birthday == 0 {
	// 	timestamp := time.Now().Unix()
	// 	uInfo.Birthday = timestamp
	// }
	sql := "insert into user(account,email,phone,pwd,nickname,icon,sex,birthday,addr) values(?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, uInfo.Account, uInfo.Email, uInfo.Phone, uInfo.Pwd, uInfo.Nickname, uInfo.Icon, uInfo.Sex, uInfo.Birthday, uInfo.Addr)
	if err != nil {
		log.Fatal(uInfo.Account, "insert data failed ", err.Error())
		return err
	}
	return nil
}
