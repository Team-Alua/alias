package cecie


// {"RequestType": "rtUpdateSave", "update": {"saveName": "data0001", "sourceFolder": "/data/dump", "selectOnly": []}}

type UpdateDetails struct {
	SaveName string `json:"saveName"`
	SourceFolder string `json:"sourceFolder"`
	SelectOnly []string `json:"selectOnly"`
}

type UpdateRequest struct {
	Type RequestType `json:"RequestType"`
	Details UpdateDetails `json:"update"`
}

func NewUpdateRequest(saveName string, sourceFolder string, selectOnly []string) *UpdateRequest {
	var req UpdateRequest
	req.Type = UpdateSave

	var det UpdateDetails
	det.SaveName = saveName
	det.SourceFolder = sourceFolder
	det.SelectOnly = selectOnly
	req.Details = det

	return &req
}

func (c *Client) UpdateSave(saveName string, sourceFolder string, selectOnly []string) error {
	req := NewUpdateRequest(saveName, sourceFolder, selectOnly)
	if err := c.sendCmd(req); err != nil {
		return err
	}
	
	return c.recvOkay()
}

