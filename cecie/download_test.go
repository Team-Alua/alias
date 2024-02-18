package cecie

import (
	"testing"
	"net"
	"bufio"
	"bytes"
)

func mockDownload(server net.Conn, data []byte) {
	go func() {
		defer server.Close()
		// This can be ignored
		// since we aren't actually
		// looking it up
		b := make([]byte, 256)
		// receive cmd
		_, err := server.Read(b)
		if err != nil {
			return
		}

		buf := bufio.NewWriter(server)

		var r DownloadResponse
		r.Size = int64(len(data))
		buf.Write(newResponseJson(&r))
		buf.Write(data)
		buf.Flush()
		return
	}()
}

func TestSmallDownload(t *testing.T) {
	server, client := net.Pipe()

	expected := []byte{1}
	mockDownload(server, expected)


	c := New(client)
    var b bytes.Buffer
	w := bufio.NewWriter(&b)

	if err := c.DownloadFile("", w); err != nil {
		t.Fatal(err)
	}

	data := b.Bytes()
	if !testEq(data, expected) {
		t.Fatalf("expected: %v got: %v", expected, data)
	}
}

func TestLargeDownload(t *testing.T) {
	server, client := net.Pipe()

	expected := make([]byte, 65536)
	for i := range 65535 {
	    expected[i] = byte('a')
	}
	mockDownload(server, expected)


	c := New(client)
    var b bytes.Buffer
	w := bufio.NewWriter(&b)

	if err := c.DownloadFile("", w); err != nil {
		t.Fatal(err)
	}

	data := b.Bytes()
	if !testEq(data, expected) {
		t.Fatalf("expected length: %d got: %d", len(expected), len(data))
	}
}

