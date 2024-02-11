package ftp

import (
	"github.com/jlaffaye/ftp"
	"time"
)

func createConnection(ip, port string) (*ftp.ServerConn, error) {
	conn, err := ftp.Dial(ip + ":" + port, ftp.DialWithTimeout(1*time.Second))
	if err != nil {
		return nil, err
	}

	err = conn.Login("anonymous", "anonymous")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

