package tcp

import (
	"net"
	"fmt"
	"time"
	"bufio"
)

func (tcp *TCP) Handle(handler func(req *Request, res *Response)) (err error) {
	tcp.listener, err = net.Listen("tcp", tcp.addr)

	var conn net.Conn

	for {
		conn, err = tcp.listener.Accept()
		if err != nil {
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()

			var (
				buf     []byte
				hasMore bool
				req     = &Request{}
				res     = &Response{conn: conn}
				limit   = tcp.GetRequestLimit()
				timeout = tcp.GetTimeout()
			)

			conn.SetDeadline(time.Now().Add(timeout))

			reader := bufio.NewReaderSize(conn, int(limit))
			buf, hasMore, err = reader.ReadLine()

			if err != nil {
				res.SendError(fmt.Sprintf("Error during request reading: %s", err.Error()))

				return
			} else if hasMore {
				res.SendError(fmt.Sprintf("Request is more that %d bytes", limit))

				return
			}

			req.body = buf

			handler(req, res)
		}(conn)
	}
}
