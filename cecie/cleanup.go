package cecie

type CleanDetails struct {
	SaveName string `json:"saveName"`
	Folder string `json:"folder"`
}

type CleanRequest struct {
	Type RequestType `json:"RequestType"`
	Details CleanDetails `json:"clean"`
}

type CleanFailEntry struct {
	Path string `json:"path"`
	Error string `json:"error"`
}

func NewCleanRequest(saveName string, folder string) *CleanRequest {
	var req CleanRequest
	req.Type = Clean
	
	var det CleanDetails
	det.SaveName = saveName
	det.Folder = folder
	req.Details = det

	return &req
}

func (c *Client) Cleanup(saveName string, folder string) ([]CleanFailEntry, error) {
	var f []CleanFailEntry
	req := NewCleanRequest(saveName, folder)

	if err := c.sendCmd(req); err != nil {
		return f, err
	}

	if err := c.recvJson(&f); err != nil {
		return f, err
	}

	return f, nil
}

