package cecie

import (
	"testing"
)


func TestResignRequest(t *testing.T) {
	saveName := "f"
	accountId := uint64(0xFFFFFFFFF)
	req := NewResignRequest(saveName, accountId)

	if req.Type != ResignSave {
		t.Fatalf("request type is not resign")
	}

	if req.Details.SaveName != saveName {
		t.Fatalf("save name does not match")
	}

	if req.Details.AccountId != accountId {
		t.Fatalf("accountId does not match")
	}
}

