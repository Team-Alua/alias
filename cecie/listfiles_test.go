package cecie

import (
	"testing"
)


func TestListFilesRequest(t *testing.T) {
	folder := "f"
	req := NewListFilesRequest(folder)

	if req.Type != ListFiles {
		t.Fatalf("request type is not list files")
	}

	if req.Details.Folder != folder {
		t.Fatalf("folder does not match")
	}
}

