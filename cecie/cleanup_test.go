package cecie

import (
	"testing"
)

func TestCleanupRequest(t *testing.T) {
	saveName := "t"
	folder := "f"
	req := NewCleanRequest(saveName, folder)

	if req.Type != Clean {
		t.Fatalf("request type is not cleanup")
	}

	if req.Details.SaveName != saveName {
		t.Fatalf("save name does not match")
	}

	if req.Details.Folder != folder {
		t.Fatalf("folder does not match")
	}
}
