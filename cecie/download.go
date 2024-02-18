package cecie

import (
	"io"
)

type DownloadDetails struct {
	Source string `json:"source"`
}

type DownloadRequest struct {
	Type RequestType `json:"ResponseType"`
	Details DownloadDetails `json:"download"`
}

type DownloadResponse struct {
	Size int64 `json:"size"`
}

func NewDownloadRequest(source string) *DownloadRequest {
	var req DownloadRequest
	req.Type = DownloadFile

	var det DownloadDetails
	det.Source = source
	req.Details = det

	return &req
}

func (c *Client) DownloadFile(remote string, w io.Writer) error {
	req := NewDownloadRequest(remote)
	if err := c.sendCmd(req); err != nil {
		return err
	}

	var dr DownloadResponse
	if err := c.recvJson(&dr); err != nil {
		return err
	}

	
	buffered := c.r.Buffered()
	if buffered > 0 {
		if _, err := io.CopyN(w, c.r, int64(buffered)); err != nil {
			return err	
		}
	}

	rem := dr.Size - int64(buffered)
	if _, err := io.CopyN(w, c.conn, rem); err != nil {
		return err
	}

	return nil
}

