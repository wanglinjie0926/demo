package kunpengBattle

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type battleConnection struct {
	conn         net.Conn
	numBytesSent int64
	numMsgSent   int64
	numBytesRecv int64
	numMsgRecv   int64
}

func (c *battleConnection) Write(msg []byte) (int, error) {
	n, err := c.conn.Write(msg)

	atomic.AddInt64(&c.numBytesSent, int64(n))
	atomic.AddInt64(&c.numMsgSent, 1)

	return n, err
}

func (c *battleConnection) Read(msg []byte) (int, error) {
	n, err := c.conn.Read(msg)

	atomic.AddInt64(&c.numBytesRecv, int64(n))
	atomic.AddInt64(&c.numMsgSent, 1)

	return n, err
}

func (c *battleConnection) reset() {
	atomic.StoreInt64(&c.numBytesRecv, 0)
	atomic.StoreInt64(&c.numBytesSent, 0)
	atomic.StoreInt64(&c.numMsgRecv, 0)
	atomic.StoreInt64(&c.numMsgSent, 0)
}

func (c *battleConnection) connect(address string) error {
	var err error
	// var tcpAddr *net.TCPAddr
	c.reset()
	// tcpAddr, err = net.ResolveTCPAddr("tcp4", address)
	// if err != nil {
	// 	log.Printf("ResolveTCPAddr Error: %v", err)
	// 	return err
	// }

	// c.conn, err = net.DialTCP("tcp4", nil, tcpAddr)
	c.conn, err = net.DialTimeout("tcp", address, time.Second*1)

	if err != nil {
		fmt.Printf("DialTCP Error: %v\n", err)
		return err
	}

	fmt.Println("TCP Socket Connected to:", c.conn.RemoteAddr())

	return err
}
