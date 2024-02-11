package ftp

func mkdir(fj *FtpJob) {
	// Target is the directory to make on ftp
	resp := fj.Response

	conn, err := createConnection(fj.Ip, fj.Port)
	if err != nil {
		resp <- &FtpResponse{Error: err}
		return
	}
	defer conn.Quit()
	err = conn.MakeDir(fj.Target)
	resp <- &FtpResponse{Error: err}
}
