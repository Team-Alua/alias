package cecie

import (
	"encoding/json"
	"errors"
)

func (c *Client) recvMsg() ([]byte, error) {
	s, err := c.r.ReadBytes('\n')
	if err != nil {
		return []byte{}, err
	}
	s = s[:len(s) - 1]

	if s[len(s) - 1] == '\r' {
		s = s[:len(s) - 1]
	}
	return s,nil
}

func (c *Client) recvReply() (*ClientResponse, error) {
	s, err := c.recvMsg()
	if err != nil {
		return nil, err
	}

	var res ClientResponse

	if err := json.Unmarshal(s, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) expectResp(resp *ClientResponse, expected ResponseType) error {
	actual := resp.Type
	if actual == expected {
		return nil
	}

	if actual == ResponseInvalid {
		return errors.New("Error " + resp.ErrorCode)
	}

	return errors.New("expected: " + string(expected) + " got: " + string(actual))
}

func (c *Client) recvOkay() error {
	resp, err := c.recvReply()
	if err != nil {
		return err
	}

	return c.expectResp(resp, ResponseOk)
}

func (c *Client) recvJson(v any) error {
	resp, err := c.recvReply()	
	if err != nil {
		return err
	}

	if err := c.expectResp(resp, ResponseJson); err != nil {
		return err
	}

	return json.Unmarshal([]byte(resp.Json), v)
}

func (c *Client) recvKeySet() (int32, error) {
	resp, err := c.recvReply()
	if err != nil {
		return 0, err
	}

	if err := c.expectResp(resp, ResponseKeySet); err != nil {
		return 0, err
	}

	return resp.KeySet, nil
}

