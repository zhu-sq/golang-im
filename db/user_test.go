package db

import (
	"fmt"
	"testing"
	"time"
)

func TestGetUserInfo(t *testing.T) {
	Testinit()
	var err error
	//数据准备，插入测试账号1,2,3,4
	uInfo1 := &UInfo{}
	uInfo1.Account = "1"
	uInfo1.Pwd = "1"
	uInfo1.Phone = 18826587459
	uInfo1.Addr = "test Addr"
	uInfo1.Email = "test@qq.com"
	uInfo1.Nickname = "test nickname"
	uInfo1.Sex = "男"
	uInfo1.Birthday = time.Now().Unix()

	uInfo2 := &UInfo{}
	uInfo2.Account = "2"
	uInfo2.Pwd = "2"

	uInfo3 := &UInfo{}
	uInfo3.Account = "3"
	uInfo3.Pwd = "3"

	uInfo4 := &UInfo{}
	uInfo4.Account = "4"
	uInfo4.Pwd = "4"

	err = SetUserInfo(uInfo1)
	if err != nil {
		t.Error("test failed 插入uInfo1失败")
		return
	}

	err = SetUserInfo(uInfo2)
	if err != nil {
		t.Error("test failed 插入uInfo2失败")
		return
	}

	err = SetUserInfo(uInfo3)
	if err != nil {
		t.Error("test failed 插入uInfo3失败")
		return
	}

	err = SetUserInfo(uInfo4)
	if err != nil {
		t.Error("test failed 插入uInfo4失败")
		return
	}

	//获取不存在的数据5
	uInf7, err := GetUserInfo("5")
	if err == nil {
		t.Error("test failed 不存在的账号 5 可以获取成功")
		return
	}
	fmt.Println(uInf7.Account)
	//获取空账号
	_, err = GetUserInfo("")
	if err == nil {
		t.Error("test failed 空账号可以获取成功")
		return
	}

	uInfo6, err := GetUserInfo("1")
	if err != nil {
		t.Error("test failed 获取账号信息失败")
		return
	}
	fmt.Printf("accouont=%s \n pwd=%s \n sex=%s \n phone=%d \n birthday=%d \n", uInfo6.Account, uInfo6.Pwd, uInfo6.Sex, uInfo6.Phone, uInfo6.Birthday)

	uInf8, err := GetUserInfo("2")
	if err != nil {
		t.Error("test failed 获取账号信息失败")
		return
	}

	fmt.Println(uInf8.Birthday, uInf8.Account, uInf8.Phone, uInf8.Sex, uInf8.Email, uInf8.UID)
}

func TestSetUserInfo(t *testing.T) {
	Testinit()
	var err error
	//测试空的账号、密码------------1
	uInfo1 := &UInfo{}
	err = SetUserInfo(uInfo1)
	if err == nil {
		t.Error("test fail 可以插入空uInfo")
		return
	}

	uInfo2 := &UInfo{}
	uInfo2.Account = "1"

	err = SetUserInfo(uInfo2)
	if err == nil {
		t.Error("test fail 密码为空可以插入")
		return
	}

	uInfo3 := &UInfo{}
	uInfo3.Pwd = "1"

	err = SetUserInfo(uInfo3)
	if err == nil {
		t.Error("test fail 账号为空可以插入")
		return
	}

	//账号密码有，其他为空------------2
	uInfo4 := &UInfo{}
	uInfo4.Account = "1"
	uInfo4.Pwd = "1"

	err = SetUserInfo(uInfo4)
	if err != nil {
		t.Error("test fail 账号，密码不为空其他为空，插入失败")
		return

	}

	//插入相同的账号------------3
	err = SetUserInfo(uInfo4)
	fmt.Println(err)
	if err == nil {
		t.Error("test fail 相同账号可以插入")
		return
	}
}
