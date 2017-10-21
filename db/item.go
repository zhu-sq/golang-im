package db

/*
登录注册 1-99
好友操作 100-199
群操作 200-299
消息操作 300-390
个人信息操作 400-490
*/
const (
	MSG_TYPE_LOGIN        = 1
	MSG_TYPE_REGISTER     = 2
	MSG_TYPE_CHANGE_PWD   = 3
	MSG_TYPE_ADD_FRIEND   = 100
	MSG_TYPE_DEL_FRIEND   = 101
	MSG_TYPE_SEND_MSG     = 300
	MSG_TYPE_MODIFY_UINFO = 400
)

type Msg struct {
	Code     int32   `json:"code,omitempty"`
	MsgType  int32   `json:"msgType,omitempty"`
	OptType  int32   `json:"optType,omitempty"`
	Account  string  `json:"account,omitempty"`
	Pwd      string  `json:"pwd,omitempty"`
	MsgTag   string  `json:"msgTag,omitempty"`
	ReMsg    string  `json:"reMsg,omitempty"`
	ReError  string  `json:"reError,omitempty"`
	UserInfo UInfo   `json:"userInfo,omitempty"`
	TalkMsg  TalkMsg `json:"talkMsg,omitempty"`
}

type UInfo struct {
	UID      int32  `json:"uID,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Account  string `json:"account,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    int64  `json:"phone,omitempty"`
	Pwd      string `json:"pwd,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Addr     string `json:"addr,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Birthday int64  `json:"birthday,omitempty"`
}

type TalkMsg struct {
	MsgType     int32  `json:"msgType,omitempty"`
	FromAccount string `json:"fromAccount,omitempty"`
	ToAccount   string `json:"toAccount,omitempty"`
	Text        string `json:"text,omitempty"`
	Photo       string `json:"photo,omitempty"`
	File        string `json:"file,omitempty"`
}
