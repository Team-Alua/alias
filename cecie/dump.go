package cecie

type DumpDetails struct {
	TargetFolder string `json:"targetFolder"`
	SaveName string `json:"saveName"`
	SelectOnly []string `json:"selectOnly"`
}

type DumpRequest struct {
	Type RequestType `json:"RequestType"`
	Details DumpDetails `json:"dump"`
}

func NewDumpRequest(saveName string, targetFolder string, selectOnly []string) *DumpRequest {
	var req DumpRequest
	req.Type = DumpSave

	var det DumpDetails
	det.SaveName = saveName
	det.TargetFolder = targetFolder
	det.SelectOnly = selectOnly
	req.Details = det

	return &req
}

func (c *Client) DumpSave(saveName string, targetFolder string, selectOnly []string) error {
	req := NewDumpRequest(saveName, targetFolder, selectOnly)
	if err := c.sendCmd(req); err != nil {
		return err
	}
	
	return c.recvOkay()
}

