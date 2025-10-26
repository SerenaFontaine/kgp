package kgp

import (
	"strings"
	"testing"
)

// TestNewDelete tests creating a delete builder
func TestNewDelete(t *testing.T) {
	db := NewDelete(DeleteAllPlacements)
	if db == nil {
		t.Fatal("NewDelete returned nil")
	}
	if db.cmd.controlData["a"] != string(ActionDelete) {
		t.Errorf("Action should be %s, got %s", ActionDelete, db.cmd.controlData["a"])
	}
	if db.cmd.controlData["d"] != string(DeleteAllPlacements) {
		t.Errorf("Delete mode should be %s, got %s", DeleteAllPlacements, db.cmd.controlData["d"])
	}
}

// TestDeleteBuilder_ImageID tests setting image ID
func TestDeleteBuilder_ImageID(t *testing.T) {
	db := NewDelete(DeleteByImageID)
	db.ImageID(42)
	if db.cmd.controlData["i"] != "42" {
		t.Errorf("Expected image ID '42', got %s", db.cmd.controlData["i"])
	}
}

// TestDeleteBuilder_ImageNumber tests setting image number
func TestDeleteBuilder_ImageNumber(t *testing.T) {
	db := NewDelete(DeleteByImageNumber)
	db.ImageNumber(100)
	if db.cmd.controlData["I"] != "100" {
		t.Errorf("Expected image number '100', got %s", db.cmd.controlData["I"])
	}
}

// TestDeleteBuilder_PlacementID tests setting placement ID
func TestDeleteBuilder_PlacementID(t *testing.T) {
	db := NewDelete(DeleteByPlacementID)
	db.PlacementID(5)
	if db.cmd.controlData["p"] != "5" {
		t.Errorf("Expected placement ID '5', got %s", db.cmd.controlData["p"])
	}
}

// TestDeleteBuilder_Cell tests setting cell coordinates
func TestDeleteBuilder_Cell(t *testing.T) {
	db := NewDelete(DeleteByCell)
	db.Cell(10, 20)
	if db.cmd.controlData["x"] != "10" {
		t.Errorf("Expected x '10', got %s", db.cmd.controlData["x"])
	}
	if db.cmd.controlData["y"] != "20" {
		t.Errorf("Expected y '20', got %s", db.cmd.controlData["y"])
	}
}

// TestDeleteBuilder_Column tests setting column
func TestDeleteBuilder_Column(t *testing.T) {
	db := NewDelete(DeleteByColumn)
	db.Column(15)
	if db.cmd.controlData["x"] != "15" {
		t.Errorf("Expected column '15', got %s", db.cmd.controlData["x"])
	}
}

// TestDeleteBuilder_Row tests setting row
func TestDeleteBuilder_Row(t *testing.T) {
	db := NewDelete(DeleteByRow)
	db.Row(25)
	if db.cmd.controlData["y"] != "25" {
		t.Errorf("Expected row '25', got %s", db.cmd.controlData["y"])
	}
}

// TestDeleteBuilder_ZIndex tests setting z-index
func TestDeleteBuilder_ZIndex(t *testing.T) {
	db := NewDelete(DeleteByZIndex)
	db.ZIndex(5)
	if db.cmd.controlData["z"] != "5" {
		t.Errorf("Expected z-index '5', got %s", db.cmd.controlData["z"])
	}
}

// TestDeleteBuilder_ResponseSuppression tests response suppression
func TestDeleteBuilder_ResponseSuppression(t *testing.T) {
	db := NewDelete(DeleteAllPlacements)
	db.ResponseSuppression(ResponseErrorsOnly)
	if db.cmd.controlData["q"] != "1" {
		t.Errorf("Expected response suppression '1', got %s", db.cmd.controlData["q"])
	}
}

// TestDeleteBuilder_Build tests building the command
func TestDeleteBuilder_Build(t *testing.T) {
	db := NewDelete(DeleteByImageID)
	db.ImageID(10)

	cmd := db.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}
}

// TestDeleteAll tests helper function
func TestDeleteAll(t *testing.T) {
	cmd := DeleteAll()
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=a") {
		t.Error("Should contain delete mode=a")
	}
}

// TestDeleteAllFree tests helper function
func TestDeleteAllFree(t *testing.T) {
	cmd := DeleteAllFree()
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=A") {
		t.Error("Should contain delete mode=A (free)")
	}
}

// TestDeleteImage tests helper function
func TestDeleteImage(t *testing.T) {
	cmd := DeleteImage(42)
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=i") {
		t.Error("Should contain delete mode=i")
	}
	if !strings.Contains(encoded, "i=42") {
		t.Error("Should contain image ID 42")
	}
}

// TestDeleteImageFree tests helper function
func TestDeleteImageFree(t *testing.T) {
	cmd := DeleteImageFree(42)
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=I") {
		t.Error("Should contain delete mode=I (free)")
	}
	if !strings.Contains(encoded, "i=42") {
		t.Error("Should contain image ID 42")
	}
}

// TestDeleteAtCursor tests helper function
func TestDeleteAtCursor(t *testing.T) {
	cmd := DeleteAtCursor()
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=c") {
		t.Error("Should contain delete mode=c")
	}
}

// TestDeleteAtCursorFree tests helper function
func TestDeleteAtCursorFree(t *testing.T) {
	cmd := DeleteAtCursorFree()
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=C") {
		t.Error("Should contain delete mode=C (free)")
	}
}

// TestDeleteModes tests all delete mode constants
func TestDeleteModes(t *testing.T) {
	modes := map[DeleteMode]string{
		DeleteAllPlacements:     "a",
		DeleteAllPlacementsFree: "A",
		DeleteByImageID:         "i",
		DeleteByImageIDFree:     "I",
		DeleteByImageNumber:     "n",
		DeleteByImageNumberFree: "N",
		DeleteByCursor:          "c",
		DeleteByCursorFree:      "C",
		DeleteByPlacementID:     "p",
		DeleteByPlacementIDFree: "P",
		DeleteByCell:            "q",
		DeleteByCellFree:        "Q",
		DeleteByColumn:          "x",
		DeleteByColumnFree:      "X",
		DeleteByRow:             "y",
		DeleteByRowFree:         "Y",
		DeleteByZIndex:          "z",
		DeleteByZIndexFree:      "Z",
	}

	for mode, expected := range modes {
		if string(mode) != expected {
			t.Errorf("Delete mode %v = %s, want %s", mode, string(mode), expected)
		}
	}
}

// TestDeleteBuilder_CompleteFlow tests a complete delete flow
func TestDeleteBuilder_CompleteFlow(t *testing.T) {
	cmd := NewDelete(DeleteByZIndex).
		ZIndex(5).
		ResponseSuppression(ResponseErrorsOnly).
		Build()

	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=d") {
		t.Error("Should contain action=d")
	}
	if !strings.Contains(encoded, "d=z") {
		t.Error("Should contain delete mode=z")
	}
	if !strings.Contains(encoded, "z=5") {
		t.Error("Should contain z-index 5")
	}
	if !strings.Contains(encoded, "q=1") {
		t.Error("Should contain response suppression 1")
	}
}
