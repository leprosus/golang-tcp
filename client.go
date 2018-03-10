package tcp

import (
	"bufio"
	"encoding/json"
	"net"
)

func (tcp *TCP) Send(bytes []byte) (result []byte, err error) {
	tcp.conn, err = net.Dial("tcp", tcp.addr)
	if err != nil {
		return
	}

	_, err = tcp.conn.Write(append(bytes, '\n'))
	if err != nil {
		return
	}

	reader := bufio.NewReader(tcp.conn)
	result, _, err = reader.ReadLine()

	return
}

func (tcp *TCP) SendJson(object interface{}) (result []byte, err error) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return
	}

	result, err = tcp.Send(bytes)

	return
}
