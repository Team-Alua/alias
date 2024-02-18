package cecie

import (
	"testing"
)


func TestDumpRequest(t *testing.T) {
	saveName := "n"
	targetFolder := "t"
	selectOnly := []string{"a"}

	req := NewDumpRequest(saveName, targetFolder, selectOnly)

	if req.Type != DumpSave {
		t.Fatalf("request type is not dump")
	}

	if req.Details.SaveName != saveName {
		t.Fatalf("save name does not match")
	}

	if req.Details.TargetFolder != targetFolder {
		t.Fatalf("target folder does not match")
	}
	
	if !testEq(req.Details.SelectOnly, selectOnly) {
		t.Fatalf("selection does not match")
	}
}

