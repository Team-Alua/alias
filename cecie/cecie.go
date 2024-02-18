package cecie

import (
	"bufio"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
	r *bufio.Reader
	w *bufio.Writer
}

func New(conn net.Conn) *Client {
	var c Client
	c.conn = conn
	c.r = bufio.NewReader(conn)
	c.w = bufio.NewWriter(conn)
	return &c
}


func Connect(ip, port string) (res *Client, err error) {
	host := ip + ":" + port
	conn, err := net.DialTimeout("tcp", host, 1 * time.Second)

	if err == nil {
		res = New(conn)
	}
	return
}

