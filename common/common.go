package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-im/db"
	"golang-im/myMap"
	"net"
	"strconv"
	"sync"
)

const (
	HOST             = ":8080"
	SIGN             = "142857"
	FRAME_LEN_FORMAT = "0x12345678"
	MAX_FRAME_LEN    = 10240 //max len of each frame
)

var (
	signLen         = len(SIGN)
	frameLenAttrLen = len(FRAME_LEN_FORMAT)
	headLen         = signLen + frameLenAttrLen
	MAX_DATA_LEN    = MAX_FRAME_LEN - headLen
	UsrMap          *myMap.MyMap
)

type IListener interface {
	OnProcess(msg *db.Msg)
}

// 在包被导入时会自动执行
func init() {
	UsrMap = myMap.CreateMapManager(5)
}

//代表每个用户的资源
type Usr struct {
	Account    string
	Conn       net.Conn
	ListenList []IListener
	Buf        []byte
	bufLen     int
	ConnState  int
	Token      string
	MsgChannel chan *db.Msg
	lock       sync.Mutex
}

// 连接的状态
const (
	CONN_ERR = 1 << iota
	CONN_DIS
)

//调用完成后，缓冲区的长度至少为minLen
func (this *Usr) readConnData(minLen int) {
	var n int
	var err error
	for this.bufLen < minLen {
		n, err = this.Conn.Read(this.Buf[this.bufLen:])
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			this.addState(CONN_ERR)
			return
		}
		this.bufLen += n
	}
}

// 缓冲区中frameLen长度以后的数据会被挪到缓冲区前面
// 如果参数为 -1 代表清空缓冲区
func (this *Usr) finish(frameLen int) {
	if frameLen == -1 {
		this.bufLen = 0
		return
	}
	nextFrameLen := this.bufLen - frameLen
	for i := 0; i < nextFrameLen; i++ {
		this.Buf[i] = this.Buf[frameLen+i]
	}
	this.bufLen = nextFrameLen
}

//为连接添加一个状态
func (this *Usr) addState(state int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ConnState = this.ConnState | state
}

//判断是否有改状态
func (this *Usr) isContainState(state int) bool {
	return (this.ConnState & state) > 0
}

//不停读取数据，读取到的数据被写进channel
func (this *Usr) HandleRead() {
	defer this.Free()
	var err error
	for {
		this.readConnData(headLen)
		if this.isContainState(CONN_ERR) {
			this.Free()
			return
		}
		//签名+数据体长度+数据体
		sign := string(this.Buf[0:signLen])
		if SIGN != sign {
			fmt.Printf("wrong sign ：%v\n", sign)
			return
		}
		dataLen64, err1 := strconv.ParseInt(string(this.Buf[signLen+2:signLen+frameLenAttrLen]), 16, 64)
		if err1 != nil {
			fmt.Printf("%v\n", err1.Error())
			this.finish(-1)
			continue
		}
		dataLen32 := int(dataLen64)
		if dataLen32 > MAX_DATA_LEN {
			fmt.Printf("dataLen more than MAX_DATA_LEN：%v\n", dataLen32)
			this.finish(-1)
			continue
		}
		this.readConnData(headLen + dataLen32)
		if this.isContainState(CONN_DIS) || this.isContainState(CONN_ERR) {
			return
		}
		frame := &db.Msg{}
		err = json.Unmarshal(this.Buf[headLen:headLen+dataLen32], frame)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			this.finish(-1)
			continue
		}
		this.MsgChannel <- frame
		this.finish(headLen + dataLen32)
	}
}

//释放用户的资源，一般在关闭连接时调用
func (this *Usr) Free() {
	if this.isContainState(CONN_DIS) {
		return
	}
	UsrMap.Delete(this.Account)
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ConnState = this.ConnState | CONN_DIS
	close(this.MsgChannel)
	if this.Conn != nil {
		this.Conn.Close()
	}
}

//往用户的连接里写进一个msg
func (this *Usr) Write(msg *db.Msg) error {
	byteData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	DataLen := fmt.Sprintf("%#08X", len(byteData))
	responseData := SIGN + DataLen + string(byteData)
	this.lock.Lock()
	defer this.lock.Unlock()
	if !(this.isContainState(CONN_DIS) || this.isContainState(CONN_ERR)) {
		_, err = this.Conn.Write([]byte(responseData))
		return err
	}
	return errors.New("连接状态异常")
}

// 添加一个监听对象
func (this *Usr) AddListener(l IListener) {
	this.ListenList = append(this.ListenList, l)
}

//根据account获取在线用户
func GetOnlineUsrByAccount(account string) *Usr {
	iUsr, exists := UsrMap.Get(account)
	if !exists {
		return nil
	}
	usr_, ok := iUsr.(*Usr)
	if !ok {
		fmt.Printf("*usr convert error by account:%v\n", account)
		return nil
	}
	return usr_
}

//简单验证登录
func VerifyLogin(account, pwd string) bool {
	return account == pwd
}
