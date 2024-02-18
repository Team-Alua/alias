package cecie

import (
	"io"
)

// {"RequestType": "rtUploadFile", "upload": {"target": "/data/dump2/abc.txt", "size": 1}}
// a

type UploadDetails struct {
	TargetFile string `json:"target"`
	Size int64 `json:"size"`
}

type UploadRequest struct {
	Type string `json:"RequestType"`
	Details UploadDetails `json:"upload"`
}

func NewUploadRequest(targetFile string, size int64) *UploadRequest {
	var req UploadRequest
	req.Type = "rtUploadFile"
	
	var det UploadDetails
	det.TargetFile = targetFile
	det.Size = size
	req.Details = det
	
	return &req
}

func (c *Client) UploadFile(remote string, r io.Reader, size int64) error {
	req := NewUploadRequest(remote, size)
	if err := c.sendCmd(req); err != nil {
		return err
	}

	if err := c.recvOkay(); err != nil {
		return err
	}

	if _, err := io.CopyN(c.conn, r, size); err != nil {
		return err
	}

	return c.recvOkay()
}

