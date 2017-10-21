package main

import (
	"fmt"
	"golang-im/common"
	"golang-im/db"
	"golang-im/listener"
	"net"
)

const (
	HOST = ":8080"
)

func main() {
	listen, err := net.Listen("tcp", HOST)
	if err != nil {
		panic("监听端口失败" + HOST)
	}
	fmt.Printf("server is running on %v\n", HOST)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("The client %s is connected\n", conn.RemoteAddr().String())
		mySrc := &common.Usr{
			Conn:       conn,
			ListenList: make([]common.IListener, 0, 20),
			Buf:        make([]byte, common.MAX_FRAME_LEN),
			MsgChannel: make(chan *db.Msg, 30),
		}
		go handleConn(mySrc)
	}
}

func handleConn(src *common.Usr) {
	src.AddListener(listener.LoginListener{Src: src})
	src.AddListener(listener.RegisterListener{Src: src})
	go src.HandleRead()
	for msg := range src.MsgChannel {
		for _, listener := range src.ListenList {
			listener.OnProcess(msg)
		}
	}
}
