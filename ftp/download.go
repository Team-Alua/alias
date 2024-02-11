package ftp

import (
	"io"
	"os"
	
)

func download(fj *FtpJob) {
	// Source is ftp
	// Target is on this filesystem
	resp := fj.Response
	conn, err := createConnection(fj.Ip, fj.Port)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	defer conn.Quit()	
	
	r, err := conn.Retr(fj.Source)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	
	defer r.Close()

	f, err := os.OpenFile(fj.Target, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0777)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	defer f.Close()
	
	if _, err := io.Copy(f, r); err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}

	resp <- &FtpResponse{Error: nil}
}

