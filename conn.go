package tcp

import (
	"net"
	"fmt"
	"time"
	"sync"
	"bufio"
)

const (
	DefaultTimeout           = 30 * time.Second
	DefaultRequestLimit uint = 1024
)

type TCP struct {
	sync.Mutex

	timeout  time.Duration
	limit    uint
	listener net.Listener
}

func Lister(host string, port uint) (tcp *TCP, err error) {
	tcp = &TCP{}

	address := fmt.Sprintf("%s:%d", host, port)
	tcp.listener, err = net.Listen("tcp", address)

	tcp.timeout = DefaultTimeout
	tcp.limit = DefaultRequestLimit

	return
}

func (tcp *TCP) SetTimeout(timeout time.Duration) {
	tcp.Lock()
	defer tcp.Unlock()

	tcp.timeout = timeout
}

func (tcp *TCP) SetRequestLimit(limit uint) {
	tcp.Lock()
	defer tcp.Unlock()

	tcp.limit = limit
}

func (tcp *TCP) Close() {
	tcp.Lock()
	defer tcp.Unlock()

	if tcp.listener != nil {
		tcp.listener.Close()
	}
}

func (tcp *TCP) Handle(handler func(req *Request, res *Response)) (err error) {
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
			)

			reader := bufio.NewReaderSize(conn, int(tcp.limit))
			buf, hasMore, err = reader.ReadLine()

			if err != nil {
				res.SendError(fmt.Sprintf("Errror during request reading: %s", err.Error()))

				return
			} else if hasMore {
				res.SendError(fmt.Sprintf("Request is more that %d bytes", tcp.limit))

				return
			}

			req.body = buf

			conn.SetDeadline(time.Now().Add(tcp.timeout))

			handler(req, res)
		}(conn)
	}
}
