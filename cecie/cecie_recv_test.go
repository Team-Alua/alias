package cecie

import (
	"net"
	"testing"
)

func TestRecvMsg(t * testing.T) {
	server, client := net.Pipe()
	expected := []byte("a")
	go func() {
		msg := append(expected, byte('\r'), byte('\n'))
		server.Write(msg)
	}()

	c := New(client)
	b, err := c.recvMsg()
	if err != nil {
		t.Fatal(err)
	}

	if !testEq(expected, b) {
		t.Fatalf("expected: %X got: %X", expected, b)
	}
}

