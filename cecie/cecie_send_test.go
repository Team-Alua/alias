package cecie

import (
	"testing"
	"time"
	"errors"
	"net"
)

func TestSendCmd(t *testing.T) {
	server, client := net.Pipe()
	test := struct {
		Type int `json:"t"`
	} {Type: 0}
	cn := make(chan error)
	expected := []byte("{\"t\":0}\r\n")

	go func() {
		buf := make([]byte, 256)
		n, err := server.Read(buf)
		if err != nil {
			cn <- err
		}
		
		if testEq(expected, buf[:n]) {
			cn <- nil
		} else {
			cn <- errors.New("Unexpected response")
		}
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		cn <- errors.New("Timeout")	
	}()

	c := New(client)
	if err := c.sendCmd(test); err != nil {
		t.Fatal(err)
	}
	if err := <-cn; err != nil {
		t.Fatal(err)
	}
}

