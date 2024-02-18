package cecie

import (
	"encoding/json"
)

func (c *Client) sendCmd(payload interface{}) error {
	b, _ := json.Marshal(payload)
	msg := append(b, []byte("\r\n")...)

	if _, err := c.w.Write(msg); err != nil {
		return err
	}
	return c.w.Flush()
}

