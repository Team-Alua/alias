package ftp

type ProcessType int

const (
	Upload ProcessType = iota
	Download
	MakeDirectory
)

type FtpResponse struct {
	Error error
}

type FtpJob struct {
	Type ProcessType
	Ip string
	Port string
	Source string
	Target string
	Response chan *FtpResponse
}

func JobLoop(jrChan chan *FtpJob) {
	for {
		jr := <-jrChan
		if jr.Type == Upload {
			go upload(jr)
		} else if jr.Type == Download {
			go download(jr)
		} else if jr.Type == MakeDirectory {
			go mkdir(jr)
		}
	}
}

