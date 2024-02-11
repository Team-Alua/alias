package main

import (
	"github.com/Team-Alua/alias/ftp"
)

func main() {
	jrChan := make(chan *ftp.FtpJob)
	go ftp.JobLoop(jrChan)

	for {

	}
}
