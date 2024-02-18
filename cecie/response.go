package cecie

import (
	"encoding/json"
)

type ResponseType string

const (
	ResponseOk ResponseType = "srOk"
	ResponseInvalid = "srInvalid"
	ResponseKeySet = "srKeySet"
	ResponseJson = "srJson"
)

type ClientResponse struct {
	Type ResponseType `json:"ResponseType"`
	KeySet int32 `json:"keyset,omitempty"`
	ErrorCode string `json:"code,omitempty"`
	Json string `json:"json,omitempty"`
}


func newResponseJson(v any) []byte {
	var c ClientResponse
	c.Type = ResponseJson

	j, _ := json.Marshal(v)
	c.Json = string(j)

	r, _ := json.Marshal(c)
	return append(r, []byte("\r\n")...)
}

func newResponseOkay() []byte {
	var c ClientResponse
	c.Type = ResponseOk
	r, _ := json.Marshal(c)
	return append(r, []byte("\r\n")...)
}

func newResponseInvalid(code string) []byte {
	var c ClientResponse
	c.Type = ResponseInvalid
	c.ErrorCode = code
	r, _ := json.Marshal(c)
	return append(r, []byte("\r\n")...)
}

