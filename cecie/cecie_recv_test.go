package cecie

import (
	"net"
	"testing"
)

func serverWrite(s net.Conn, msg []byte) {
	go func() {
		s.Write(append(msg, []byte("\r\n")...))
	}()

}

func TestRecvMsg(t *testing.T) {
	server, client := net.Pipe()
	expected := []byte("a")
	serverWrite(server, expected)

	c := New(client)
	b, err := c.recvMsg()
	if err != nil {
		t.Fatal(err)
	}

	if !testEq(expected, b) {
		t.Fatalf("expected: %X got: %X", expected, b)
	}
}


func TestRecvOkay(t *testing.T) {
	server, client := net.Pipe()
	msg := []byte(`{"ResponseType":"srOk"}`)
	serverWrite(server, msg)
	c := New(client)
	if err := c.recvOkay(); err != nil {
		t.Fatal(err)
	}
}

func TestRecvKeySet(t *testing.T) {
	server, client := net.Pipe()
	expected := int32(10)
	msg := []byte(`{"ResponseType": "srKeySet", "keyset": 10}`)
	serverWrite(server, msg)
	c := New(client)
	ks, err := c.recvKeySet()
	if  err != nil {
		t.Fatal(err)
	}
	if ks != expected {
		t.Fatalf("expected: %d got: %d", expected, ks)
	}
}

func TestRecvJson(t *testing.T) {
	server, client := net.Pipe()
	expected := []int{1}
	msg := []byte(`{"ResponseType":"srJson","json":"[1]"}`)
	serverWrite(server, msg)
	var data []int
	c := New(client)
	if err := c.recvJson(&data); err != nil {
		t.Fatal(err)
	}

	if !testEq(expected, data) {
		t.Fatalf("expected: %v got: %v", expected, data)
	}
}

