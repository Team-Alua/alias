package cecie

// {"RequestType": "rtKeySet"}


type KeySetRequest struct {
	Type RequestType `json:"RequestType"`
}

func (c *Client) KeySet() (int32, error) {
	var req KeySetRequest	
	req.Type = KeySet

	if err := c.sendCmd(req); err != nil {
		return 0, err
	}

	return c.recvKeySet()
}

