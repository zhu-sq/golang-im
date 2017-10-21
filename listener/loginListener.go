package listener

import (
	"fmt"
	"golang-im/common"
	"golang-im/db"
	"golang-im/utils"
)

type LoginListener struct {
	Src *common.Usr
}

func (this LoginListener) OnProcess(m *db.Msg) {
	if m.MsgType != db.MSG_TYPE_LOGIN {
		return
	}
	reMsg := &db.Msg{}
	reMsg.MsgType = db.MSG_TYPE_LOGIN
	account := m.Account
	pwd := m.Pwd
	if account == "" || pwd == "" {
		reMsg.Code = 101
		reMsg.ReError = "账户或者密码为空"
		err := this.Src.Write(reMsg)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		return
	}

	uInfo, err := db.GetUserInfo(account)
	if err != nil {
		reMsg.Code = 104
		reMsg.ReError = "数据库操作错误：" + err.Error()
		err := this.Src.Write(reMsg)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		return
	}

	if pwd != uInfo.Pwd {
		reMsg.Code = 101
		reMsg.ReError = "账户或者密码不正确"
		err := this.Src.Write(reMsg)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		return
	}

	reMsg.Code = 0
	reMsg.UserInfo = *uInfo

	fmt.Printf("%s 登陆成功 -- %s\n", this.Src.Conn.RemoteAddr().String(), account)
	existsSrc := common.GetOnlineUsrByAccount(account)
	if existsSrc != nil {
		fmt.Printf("旧连接 %v 被关闭: %v\n", account, existsSrc.Conn.RemoteAddr().String())
		existsSrc.Free() //释放用户的资源
	}
	this.Src.Token = utils.NewToken(account, pwd) //保存新的token
	this.Src.Account = account

	err = this.Src.Write(reMsg)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	// //common.UsrMap.Set(account, this.Src) //记录这个账号登陆了
	// // reMsg.Token = &this.Src.Token
	// // reMsg.ResponseState = protobuf.StatusCode_SUCCESS.Enum() //返回成功标志
	// // } else {
	// // 	fmt.Printf("%s 密码错误\n", this.Src.Conn.RemoteAddr().String())
	// // 	// reMsg.ResponseState = protobuf.StatusCode_FAILED.Enum()
	// // 	// errMsg := "用户名或密码错误"
	// // 	// reMsg.ErrMsg = &errMsg
	// // }

	// // 如果android使用Tools，也就是实用Tcp回调方法，以下两行就一定要加上，否则android端无法响应回调
	// // msgUniqueTag := m.GetMsgUniqueTag()
	// // reMsg.MsgUniqueTag = &msgUniqueTag

	// err := this.Src.Write(reMsg)
	// if err != nil {
	// 	fmt.Printf("%v\n", err.Error())
	// }
}
