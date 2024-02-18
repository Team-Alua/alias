package cecie

import (
	"testing"
)


func TestListSaveFilesRequest(t *testing.T) {
	saveName := "f"
	req := NewListSaveFilesRequest(saveName)

	if req.Type != ListSaveFiles {
		t.Fatalf("request type is not list save files")
	}

	if req.Details.SaveName != saveName {
		t.Fatalf("save name does not match")
	}
}

