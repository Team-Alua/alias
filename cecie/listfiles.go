package cecie


// {"RequestType": "rtListSaveFiles", "list": {"saveName": "data0001"}}

//type ListEntry = object
//  kind*: PathComponent
//  path*: string
//  size*: Off
//  mode*: Mode
//  uid*: Uid
//  gid*: Gid


type ListFilesDetails struct {
	Folder string `json:"folder"`
}

type ListFilesRequest struct {
	Type RequestType `json:"RequestType"`
	Details ListFilesDetails `json:"ls"`
}


func NewListFilesRequest(folder string) *ListFilesRequest {
	var req ListFilesRequest
	req.Type = ListFiles

	var det ListFilesDetails
	det.Folder = folder
	req.Details = det

	return &req
}

type ListFilesEntry struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
	Size int64 `json:"size"`
	Mode uint32 `json:"mode"`
	Uid  uint32 `json:"uid"`
	Gid  uint32 `json:"gid"`
}

type ListFilesResponse []ListFilesEntry

func (c *Client) ListFiles(remoteFolder string) (ListFilesResponse, error) {
	var r ListFilesResponse

	req := NewListFilesRequest(remoteFolder)
	if err := c.sendCmd(req); err != nil {
		return r, err
	}

	if err := c.recvJson(&r); err != nil {
		return r, err
	}
	
	return r, nil
}

