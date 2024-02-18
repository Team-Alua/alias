package cecie

import (
	"encoding/json"
	"testing"
	"net"
	"bufio"
	"bytes"
	"io"
	"fmt"
)

func mockUpload(server net.Conn, bChan chan []byte, size int64) {
	go func() {
		defer server.Close()
		br := bufio.NewReader(server)
		// receive cmd
		b, err := br.ReadBytes('\n')
		if err != nil {
			return
		}

		b = b[:len(b) - 2]

		buf := bufio.NewWriter(server)

		var ur UploadRequest
		if err := json.Unmarshal(b, &ur); err != nil {
			buf.Write(newResponseInvalid("Could not parse request"))
			buf.Flush()	
			return
		}

		if ur.Details.Size != size {
			msg := fmt.Sprintf("expected size: %d got: %d", size, ur.Details.Size)
			msg += "\n" + string(b)
			buf.Write(newResponseInvalid(msg))
			buf.Flush()	
			return
		}

		buf.Write(newResponseOkay())
		buf.Flush()	

		var fb bytes.Buffer
		w := bufio.NewWriter(&fb)

		io.CopyN(w, br, int64(br.Buffered()))
		size -= int64(br.Buffered())

		if _, err := io.CopyN(w, br, size); err != nil {
			buf.Write(newResponseInvalid(err.Error()))
			buf.Flush()	
			return	
		}

		buf.Write(newResponseOkay())
		buf.Flush()
		bChan <- fb.Bytes()
		return
	}()
}

func TestSmallUpload(t *testing.T) {
	server, client := net.Pipe()

	expected := []byte{1}
	size := int64(len(expected))
	bChan := make(chan []byte, 1)
	mockUpload(server, bChan, size)


	c := New(client)
	r := bytes.NewReader(expected)

	if err := c.UploadFile("", r, size); err != nil {
		t.Fatal(err)
	}

	data, ok := <-bChan
	if !ok {
		t.Fatalf("Server closed channel")
	}

	if !testEq(data, expected) {
		t.Fatalf("expected: %v got: %v", expected, data)
	}
}

func TestLargeUpload(t *testing.T) {
	server, client := net.Pipe()

	expected := make([]byte, 65536)
	for i := range 65535 {
	    expected[i] = byte('a')
	}
	size := int64(len(expected))
	bChan := make(chan []byte)
	mockUpload(server, bChan, size)


	c := New(client)
	r := bytes.NewReader(expected)

	if err := c.UploadFile("", r, size); err != nil {
		t.Fatal(err)
	}

	data, ok := <-bChan
	if !ok {
		t.Fatalf("Server closed channel")
	}

	if !testEq(data, expected) {
		t.Fatalf("expected length: %d got: %d", len(expected), len(data))
	}
}

func TestUploadRequest(t *testing.T) {
	remote := "f"
	size := int64(50)
	req := NewUploadRequest(remote, size)

	if req.Type != UploadFile {
		t.Fatalf("request type is not download")
	}

	if req.Details.TargetFile != remote {
		t.Fatalf("target does not match")
	}

	if req.Details.Size != size {
		t.Fatalf("size does not match")
	}
}

