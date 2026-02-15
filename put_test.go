package kgp

import (
	"strings"
	"testing"
)

// TestNewPut tests creating a put builder
func TestNewPut(t *testing.T) {
	pb := NewPut(42)
	if pb == nil {
		t.Fatal("NewPut returned nil")
	}
	if pb.cmd.controlData["a"] != string(ActionPut) {
		t.Errorf("Action should be %s, got %s", ActionPut, pb.cmd.controlData["a"])
	}
	if pb.cmd.controlData["i"] != "42" {
		t.Errorf("Image ID should be '42', got %s", pb.cmd.controlData["i"])
	}
}

// TestPutBuilder_ImageNumber tests setting image number
func TestPutBuilder_ImageNumber(t *testing.T) {
	pb := NewPut(10)
	pb.ImageNumber(100)
	if pb.cmd.controlData["I"] != "100" {
		t.Errorf("Expected image number '100', got %s", pb.cmd.controlData["I"])
	}
}

// TestPutBuilder_PlacementID tests setting placement ID
func TestPutBuilder_PlacementID(t *testing.T) {
	pb := NewPut(10)
	pb.PlacementID(5)
	if pb.cmd.controlData["p"] != "5" {
		t.Errorf("Expected placement ID '5', got %s", pb.cmd.controlData["p"])
	}
}

// TestPutBuilder_CellOffset tests cell offset
func TestPutBuilder_CellOffset(t *testing.T) {
	pb := NewPut(10)
	pb.CellOffset(15, 25)
	if pb.cmd.controlData["X"] != "15" {
		t.Errorf("Expected X offset '15', got %s", pb.cmd.controlData["X"])
	}
	if pb.cmd.controlData["Y"] != "25" {
		t.Errorf("Expected Y offset '25', got %s", pb.cmd.controlData["Y"])
	}
}

// TestPutBuilder_DisplaySize tests display size
func TestPutBuilder_DisplaySize(t *testing.T) {
	pb := NewPut(10)
	pb.DisplaySize(30, 20)
	if pb.cmd.controlData["c"] != "30" {
		t.Errorf("Expected columns '30', got %s", pb.cmd.controlData["c"])
	}
	if pb.cmd.controlData["r"] != "20" {
		t.Errorf("Expected rows '20', got %s", pb.cmd.controlData["r"])
	}
}

// TestPutBuilder_SourceRect tests source rectangle
func TestPutBuilder_SourceRect(t *testing.T) {
	pb := NewPut(10)
	pb.SourceRect(50, 60, 70, 80)
	if pb.cmd.controlData["x"] != "50" {
		t.Errorf("Expected x '50', got %s", pb.cmd.controlData["x"])
	}
	if pb.cmd.controlData["y"] != "60" {
		t.Errorf("Expected y '60', got %s", pb.cmd.controlData["y"])
	}
	if pb.cmd.controlData["w"] != "70" {
		t.Errorf("Expected width '70', got %s", pb.cmd.controlData["w"])
	}
	if pb.cmd.controlData["h"] != "80" {
		t.Errorf("Expected height '80', got %s", pb.cmd.controlData["h"])
	}
}

// TestPutBuilder_ZIndex tests z-index
func TestPutBuilder_ZIndex(t *testing.T) {
	pb := NewPut(10)
	pb.ZIndex(10)
	if pb.cmd.controlData["z"] != "10" {
		t.Errorf("Expected z-index '10', got %s", pb.cmd.controlData["z"])
	}
}

// TestPutBuilder_CursorMovement tests cursor movement
func TestPutBuilder_CursorMovement(t *testing.T) {
	pb := NewPut(10)
	pb.CursorMovement(false)
	if pb.cmd.controlData["C"] != "1" {
		t.Errorf("Expected cursor movement '1' (no move), got %s", pb.cmd.controlData["C"])
	}

	pb2 := NewPut(10)
	pb2.CursorMovement(true)
	if pb2.cmd.controlData["C"] != "0" {
		t.Errorf("Expected cursor movement '0' (move), got %s", pb2.cmd.controlData["C"])
	}
}

// TestPutBuilder_VirtualPlacement tests virtual placement
func TestPutBuilder_VirtualPlacement(t *testing.T) {
	pb := NewPut(10)
	pb.VirtualPlacement()
	if pb.cmd.controlData["U"] != "1" {
		t.Errorf("Expected virtual placement '1', got %s", pb.cmd.controlData["U"])
	}
}

// TestPutBuilder_RelativeTo tests relative positioning
func TestPutBuilder_RelativeTo(t *testing.T) {
	pb := NewPut(10)
	pb.RelativeTo(100, 1, 5, 3)
	if pb.cmd.controlData["P"] != "100" {
		t.Errorf("Expected parent image ID '100', got %s", pb.cmd.controlData["P"])
	}
	if pb.cmd.controlData["Q"] != "1" {
		t.Errorf("Expected parent placement ID '1', got %s", pb.cmd.controlData["Q"])
	}
	if pb.cmd.controlData["H"] != "5" {
		t.Errorf("Expected H offset '5', got %s", pb.cmd.controlData["H"])
	}
	if pb.cmd.controlData["V"] != "3" {
		t.Errorf("Expected V offset '3', got %s", pb.cmd.controlData["V"])
	}
}

// TestPutBuilder_ResponseSuppression tests response suppression
func TestPutBuilder_ResponseSuppression(t *testing.T) {
	pb := NewPut(10)
	pb.ResponseSuppression(ResponseOKOnly)
	if pb.cmd.controlData["q"] != "2" {
		t.Errorf("Expected response suppression '2', got %s", pb.cmd.controlData["q"])
	}
}

// TestPutBuilder_Build tests building the command
func TestPutBuilder_Build(t *testing.T) {
	pb := NewPut(10)
	pb.PlacementID(1)
	pb.DisplaySize(20, 15)

	cmd := pb.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}
}

// TestPutBuilder_CompleteFlow tests a complete put flow
func TestPutBuilder_CompleteFlow(t *testing.T) {
	cmd := NewPut(100).
		PlacementID(2).
		DisplaySize(10, 7).
		CellOffset(5, 5).
		ZIndex(2).
		Build()

	encoded := cmd.Encode()

	// Verify the encoding contains expected keys
	if !strings.Contains(encoded, "a=p") {
		t.Error("Should contain action=p")
	}
	if !strings.Contains(encoded, "i=100") {
		t.Error("Should contain image ID 100")
	}
	if !strings.Contains(encoded, "p=2") {
		t.Error("Should contain placement ID 2")
	}
	if !strings.Contains(encoded, "c=10") {
		t.Error("Should contain columns 10")
	}
	if !strings.Contains(encoded, "r=7") {
		t.Error("Should contain rows 7")
	}
	if !strings.Contains(encoded, "X=5") {
		t.Error("Should contain X offset 5")
	}
	if !strings.Contains(encoded, "Y=5") {
		t.Error("Should contain Y offset 5")
	}
	if !strings.Contains(encoded, "z=2") {
		t.Error("Should contain z-index 2")
	}
}
