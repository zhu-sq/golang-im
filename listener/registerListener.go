package listener

import (
	"fmt"
	"golang-im/common"
	"golang-im/db"
)

type RegisterListener struct {
	Src *common.Usr
}

func (this RegisterListener) OnProcess(m *db.Msg) {

	if m.MsgType != db.MSG_TYPE_REGISTER {
		return
	}
	
	reMsg := &db.Msg{}
	reMsg.MsgType = db.MSG_TYPE_REGISTER
	uInfo := &m.UserInfo


	if uInfo.Account == "" || uInfo.Pwd == "" {
		reMsg.Code = 201
		reMsg.ReError = "账户或者密码为空"
		fmt.Printf("%v\n", "账户或者密码为空")
		err := this.Src.Write(reMsg)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		return
	}
	
	err := db.SetUserInfo(uInfo)
	if err != nil {
		reMsg.Code = 204
		reMsg.ReError = "数据库操作错误,注册失败：" + err.Error()
		fmt.Printf("%v\n", err.Error())
		err := this.Src.Write(reMsg)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		return
	}

	reMsg.Code = 0
	fmt.Printf("%s 注册成功 -- %s\n", this.Src.Conn.RemoteAddr().String(), uInfo.Account)
	// existsSrc := common.GetOnlineUsrByAccount(uInfo.Account)
	// if existsSrc != nil {
	// 	fmt.Printf("旧连接 %v 被关闭: %v\n", uInfo.Account, existsSrc.Conn.RemoteAddr().String())
	// 	existsSrc.Free() //释放用户的资源
	// }
	err = this.Src.Write(reMsg)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
}
