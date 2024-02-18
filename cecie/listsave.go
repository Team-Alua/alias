package cecie


// {"RequestType": "rtListSaveFiles", "list": {"saveName": "data0001"}}

type ListSaveFilesDetails struct {
	SaveName string `json:"saveName"`
}

type ListSaveFilesRequest struct {
	Type RequestType `json:"RequestType"`
	Details ListSaveFilesDetails `json:"list"`
}


func NewListSaveFilesRequest(saveName string) *ListSaveFilesRequest {
	var req ListSaveFilesRequest
	req.Type = ListSaveFiles

	var det ListSaveFilesDetails
	det.SaveName = saveName
	req.Details = det

	return &req
}

type ListSaveFilesEntry struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
	Size int64 `json:"size"`
	Mode uint32 `json:"mode"`
	Uid  uint32 `json:"uid"`
	Gid  uint32 `json:"gid"`
}

type ListSaveFilesResponse []ListSaveFilesEntry

func (c *Client) ListSaveFiles(saveName string) (ListSaveFilesResponse, error) {
	var r ListSaveFilesResponse

	req := NewListSaveFilesRequest(saveName)
	if err := c.sendCmd(req); err != nil {
		return r, err
	}

	if err := c.recvJson(&r); err != nil {
		return r, err
	}
	
	return r, nil
}

