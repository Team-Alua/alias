package cecie

type RequestType string

const (
	KeySet RequestType = "rtKeySet"
	ListSaveFiles = "rtListSaveFiles"
	CreateSave = "rtCreateSave"
	DumpSave = "rtDumpSave"
	UpdateSave = "rtUpdateSave"
	ResignSave = "rtResignSave"
	Clean = "rtClean"
	UploadFile = "rtUploadFile"
	DownloadFile = "rtDownloadFile"
	ListFiles = "rtListFiles"
	Invalid = "rtInvalid"
)


