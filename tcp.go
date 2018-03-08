package tcp

import (
	"time"
	"sync"
	"net"
	"fmt"
)

const (
	DefaultTimeout           = 30 * time.Second
	DefaultRequestLimit uint = 1024
)

type TCP struct {
	sync.Mutex

	host string
	port uint
	addr string

	listener net.Listener
	conn     net.Conn

	timeout time.Duration
	limit   uint
}

func Init(host string, port uint) (tcp *TCP) {
	tcp = &TCP{
		host: host,
		port: port,
	}

	tcp.addr = fmt.Sprintf("%s:%d", host, port)

	tcp.SetTimeout(DefaultTimeout)
	tcp.SetRequestLimit(DefaultRequestLimit)

	return
}

func (tcp *TCP) SetTimeout(timeout time.Duration) {
	tcp.Lock()
	defer tcp.Unlock()

	tcp.timeout = timeout
}

func (tcp *TCP) GetTimeout() (timeout time.Duration) {
	tcp.Lock()
	defer tcp.Unlock()

	return tcp.timeout
}

func (tcp *TCP) SetRequestLimit(limit uint) {
	tcp.Lock()
	defer tcp.Unlock()

	tcp.limit = limit
}

func (tcp *TCP) GetRequestLimit() (limit uint) {
	tcp.Lock()
	defer tcp.Unlock()

	return tcp.limit
}

func (tcp *TCP) Close() {
	tcp.Lock()
	defer tcp.Unlock()

	if tcp.listener != nil {
		tcp.listener.Close()
	}

	if tcp.conn != nil {
		tcp.conn.Close()
	}
}
