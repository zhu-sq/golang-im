package protocol

import (
	"encoding/json"
	"strconv"
)

const (
	HEAD    = "142857"
	HEADLEN = 6
	MSGLEN  = 8
)

/*//封包
func Packet(message []byte) []byte {
	//return append(append([]byte(HEAD), IntToBytes(len(message))...), message...)
	return
}
*/
//解包
func Unpack(buffer []byte, msgChan chan *ClnMsg) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+HEADLEN+MSGLEN {
			break
		}
		if string(buffer[i:i+HEADLEN]) == HEAD {
			msglen, err := strconv.Atoi(string(buffer[i+HEADLEN : i+HEADLEN+MSGLEN]))
			if err != nil {
				break
			}
			if length < i+HEADLEN+MSGLEN+msglen {
				break
			}
			dataBuf := buffer[i+HEADLEN+MSGLEN : i+HEADLEN+MSGLEN+msglen]
			msg := &ClnMsg{}
			err = json.Unmarshal(dataBuf, msg)
			if err != nil {

				msgChan <- nil
				break
			}

			msgChan <- msg
			i += HEADLEN + msglen + MSGLEN - 1
			break
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
