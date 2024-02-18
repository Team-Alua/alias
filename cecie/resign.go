package cecie

type ResignDetails struct {
	SaveName string `json:"saveName"`
	AccountId uint64 `json:"accountId"`
}

type ResignRequest struct {
	Type RequestType `json:"RequestType"`
	Details ResignDetails `json:"resign"`
}

func NewResignRequest(saveName string, accountId uint64) ResignRequest {
	var req ResignRequest
	req.Type = ResignSave

	var det ResignDetails
	det.SaveName = saveName
	det.AccountId = accountId
	req.Details = det

	return req
}

func (c *Client) Resign(saveName string, accountId uint64) error {
	req := NewResignRequest(saveName, accountId)
	if err := c.sendCmd(req); err != nil {
		return err
	}

	return c.recvOkay()
}

