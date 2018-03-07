package tcp

import (
	"encoding/json"
	"net"
)

type Response struct {
	conn net.Conn
}

type Error struct {
	Text string `json:"err"`
}

func (res *Response) SendError(text string) (err error) {
	return res.SendJson(Error{Text: text})
}

func (res *Response) SendJson(object interface{}) (err error) {
	var bytes []byte

	bytes, err = json.Marshal(object)

	err = res.send(bytes)
	if err != nil {
		return
	}

	err = res.send(bytes)

	return
}

func (res *Response) send(bytes []byte) (err error) {
	_, err = res.conn.Write(bytes)

	return
}
