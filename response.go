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

type Status struct {
	Status string `json:"status"`
}

func (res *Response) SendError(text string) (err error) {
	return res.SendJson(Error{Text: text})
}

func (res *Response) SendStatus(isWork bool) (err error) {
	if isWork {
		err = res.SendJson(Status{Status: "ok"})
	}

	return
}

func (res *Response) SendJson(object interface{}) (err error) {
	var bytes []byte

	bytes, err = json.Marshal(object)

	err = res.send(bytes)

	return
}

func (res *Response) send(bytes []byte) (err error) {
	_, err = res.conn.Write(append(bytes, '\n'))

	return
}
