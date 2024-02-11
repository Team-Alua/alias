package ftp

import (
	"os"
	"io"
)

func upload(fj *FtpJob) {
	// Source is on the file system
	// Target is on ftp
	resp := fj.Response

	conn, err := createConnection(fj.Ip, fj.Port)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	defer conn.Quit()

	f, err := os.Open(fj.Source)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	defer f.Close()

	r,w := io.Pipe()
	go func() {
		err := conn.Stor(fj.Target,r)
		if err != nil {
			w.CloseWithError(err)
		} else {
			w.Close()
		}
	}()

	if _, err := io.Copy(w, f); err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}

	resp <- &FtpResponse{Error: nil}
}
