package cecie

import (
	"testing"
)


func TestUpdateRequest(t *testing.T) {
	saveName := "n"
	sourceFolder := "t"
	selectOnly := []string{"a"}

	req := NewUpdateRequest(saveName, sourceFolder, selectOnly)

	if req.Type != UpdateSave {
		t.Fatalf("request type is not update")
	}

	if req.Details.SaveName != saveName {
		t.Fatalf("save name does not match")
	}

	if req.Details.SourceFolder != sourceFolder {
		t.Fatalf("target folder does not match")
	}
	
	if !testEq(req.Details.SelectOnly, selectOnly) {
		t.Fatalf("selection does not match")
	}
}


