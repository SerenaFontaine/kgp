package kgp

import (
	"errors"
	"strings"
	"testing"
)

// TestNewTransmit tests creating a transmit builder
func TestNewTransmit(t *testing.T) {
	tb := NewTransmit()
	if tb == nil {
		t.Fatal("NewTransmit returned nil")
	}
	if tb.display {
		t.Error("NewTransmit should not set display flag")
	}
	if tb.cmd.controlData["a"] != string(ActionTransmit) {
		t.Errorf("Action should be %s, got %s", ActionTransmit, tb.cmd.controlData["a"])
	}
}

// TestNewTransmitDisplay tests creating a transmit+display builder
func TestNewTransmitDisplay(t *testing.T) {
	tb := NewTransmitDisplay()
	if tb == nil {
		t.Fatal("NewTransmitDisplay returned nil")
	}
	if !tb.display {
		t.Error("NewTransmitDisplay should set display flag")
	}
	if tb.cmd.controlData["a"] != string(ActionTransmitDisplay) {
		t.Errorf("Action should be %s, got %s", ActionTransmitDisplay, tb.cmd.controlData["a"])
	}
}

// TestTransmitBuilder_ImageID tests setting image ID
func TestTransmitBuilder_ImageID(t *testing.T) {
	tb := NewTransmit()
	tb.ImageID(42)
	if tb.cmd.controlData["i"] != "42" {
		t.Errorf("Expected image ID '42', got %s", tb.cmd.controlData["i"])
	}
}

// TestTransmitBuilder_ImageNumber tests setting image number
func TestTransmitBuilder_ImageNumber(t *testing.T) {
	tb := NewTransmit()
	tb.ImageNumber(100)
	if tb.cmd.controlData["I"] != "100" {
		t.Errorf("Expected image number '100', got %s", tb.cmd.controlData["I"])
	}
}

// TestTransmitBuilder_Format tests setting format
func TestTransmitBuilder_Format(t *testing.T) {
	tb := NewTransmit()
	tb.Format(FormatPNG)
	if tb.cmd.controlData["f"] != "100" {
		t.Errorf("Expected format '100', got %s", tb.cmd.controlData["f"])
	}
}

// TestTransmitBuilder_Dimensions tests setting dimensions
func TestTransmitBuilder_Dimensions(t *testing.T) {
	tb := NewTransmit()
	tb.Dimensions(800, 600)
	if tb.cmd.controlData["s"] != "800" {
		t.Errorf("Expected width '800', got %s", tb.cmd.controlData["s"])
	}
	if tb.cmd.controlData["v"] != "600" {
		t.Errorf("Expected height '600', got %s", tb.cmd.controlData["v"])
	}
}

// TestTransmitBuilder_Compress tests enabling compression
func TestTransmitBuilder_Compress(t *testing.T) {
	tb := NewTransmit()
	tb.Compress()
	if tb.cmd.controlData["o"] != string(CompressionZlib) {
		t.Errorf("Expected compression 'z', got %s", tb.cmd.controlData["o"])
	}
	if tb.compression != CompressionZlib {
		t.Error("Compression field should be set")
	}
}

// TestTransmitBuilder_TransmitDirect tests direct transmission
func TestTransmitBuilder_TransmitDirect(t *testing.T) {
	tb := NewTransmit()
	data := []byte("test data")
	tb.TransmitDirect(data)

	if tb.cmd.controlData["t"] != string(TransmitDirect) {
		t.Errorf("Expected transmit mode 'd', got %s", tb.cmd.controlData["t"])
	}
	if string(tb.imageData) != "test data" {
		t.Errorf("Expected 'test data', got %s", string(tb.imageData))
	}
}

// TestTransmitBuilder_TransmitFile tests file transmission
func TestTransmitBuilder_TransmitFile(t *testing.T) {
	tb := NewTransmit()
	tb.TransmitFile("/path/to/image.png")

	if tb.cmd.controlData["t"] != string(TransmitFile) {
		t.Errorf("Expected transmit mode 'f', got %s", tb.cmd.controlData["t"])
	}
	if string(tb.imageData) != "/path/to/image.png" {
		t.Errorf("Expected path, got %s", string(tb.imageData))
	}
}

// TestTransmitBuilder_TransmitFileWithOffset tests file transmission with offset
func TestTransmitBuilder_TransmitFileWithOffset(t *testing.T) {
	tb := NewTransmit()
	tb.TransmitFileWithOffset("/path/to/image.png", 1024, 50000)

	if tb.cmd.controlData["t"] != string(TransmitFile) {
		t.Error("Expected file transmit mode")
	}
	if tb.cmd.controlData["O"] != "1024" {
		t.Errorf("Expected offset '1024', got %s", tb.cmd.controlData["O"])
	}
	if tb.cmd.controlData["S"] != "50000" {
		t.Errorf("Expected size '50000', got %s", tb.cmd.controlData["S"])
	}
}

// TestTransmitBuilder_TransmitTemp tests temporary file transmission
func TestTransmitBuilder_TransmitTemp(t *testing.T) {
	tb := NewTransmit()
	tb.TransmitTemp("/tmp/tty-graphics-protocol-12345.png")

	if tb.cmd.controlData["t"] != string(TransmitTemp) {
		t.Errorf("Expected transmit mode 't', got %s", tb.cmd.controlData["t"])
	}
}

// TestTransmitBuilder_TryTransmitTemp tests non-panicking temp transmission API.
func TestTransmitBuilder_TryTransmitTemp(t *testing.T) {
	tb := NewTransmit()
	got, err := tb.TryTransmitTemp("/tmp/tty-graphics-protocol-12345.png")
	if err != nil {
		t.Fatalf("TryTransmitTemp returned error: %v", err)
	}
	if got != tb {
		t.Fatal("TryTransmitTemp should return the same builder")
	}
	if tb.cmd.controlData["t"] != string(TransmitTemp) {
		t.Errorf("Expected transmit mode 't', got %s", tb.cmd.controlData["t"])
	}
}

// TestTransmitBuilder_TryTransmitTempInvalidPath tests invalid path handling without panic.
func TestTransmitBuilder_TryTransmitTempInvalidPath(t *testing.T) {
	tb := NewTransmit()
	got, err := tb.TryTransmitTemp("/tmp/image.png")
	if !errors.Is(err, ErrInvalidTempPath) {
		t.Fatalf("expected ErrInvalidTempPath, got %v", err)
	}
	if got != nil {
		t.Fatal("TryTransmitTemp should return nil builder on error")
	}
}

// TestTransmitBuilder_TransmitTempInvalidPath tests path validation for temp transmission
func TestTransmitBuilder_TransmitTempInvalidPath(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for invalid temp path")
		}
	}()

	NewTransmit().TransmitTemp("/tmp/image.png")
}

// TestValidateTempPath tests temporary path validation helper.
func TestValidateTempPath(t *testing.T) {
	if err := ValidateTempPath("/tmp/tty-graphics-protocol-1.png"); err != nil {
		t.Fatalf("ValidateTempPath returned unexpected error: %v", err)
	}
	if err := ValidateTempPath("/tmp/other.png"); !errors.Is(err, ErrInvalidTempPath) {
		t.Fatalf("expected ErrInvalidTempPath, got %v", err)
	}
}

// TestTransmitBuilder_TransmitSharedMemory tests shared memory transmission
func TestTransmitBuilder_TransmitSharedMemory(t *testing.T) {
	tb := NewTransmit()
	tb.TransmitSharedMemory("/my-shm-object", 1920000)

	if tb.cmd.controlData["t"] != string(TransmitSharedMem) {
		t.Errorf("Expected transmit mode 's', got %s", tb.cmd.controlData["t"])
	}
	if tb.cmd.controlData["S"] != "1920000" {
		t.Errorf("Expected size '1920000', got %s", tb.cmd.controlData["S"])
	}
	if string(tb.imageData) != "/my-shm-object" {
		t.Errorf("Expected shm name, got %s", string(tb.imageData))
	}
}

// TestTransmitBuilder_PlacementID tests setting placement ID
func TestTransmitBuilder_PlacementID(t *testing.T) {
	tb := NewTransmit()
	tb.PlacementID(5)
	if tb.cmd.controlData["p"] != "5" {
		t.Errorf("Expected placement ID '5', got %s", tb.cmd.controlData["p"])
	}
}

// TestTransmitBuilder_VirtualPlacement tests virtual placement
func TestTransmitBuilder_VirtualPlacement(t *testing.T) {
	tb := NewTransmit()
	tb.VirtualPlacement()
	if tb.cmd.controlData["U"] != "1" {
		t.Errorf("Expected virtual placement '1', got %s", tb.cmd.controlData["U"])
	}
}

// TestTransmitBuilder_ResponseSuppression tests response suppression
func TestTransmitBuilder_ResponseSuppression(t *testing.T) {
	tb := NewTransmit()
	tb.ResponseSuppression(ResponseErrorsOnly)
	if tb.cmd.controlData["q"] != "1" {
		t.Errorf("Expected response suppression '1', got %s", tb.cmd.controlData["q"])
	}
}

// TestTransmitBuilder_CellOffset tests cell offset
func TestTransmitBuilder_CellOffset(t *testing.T) {
	tb := NewTransmit()
	tb.CellOffset(10, 20)
	if tb.cmd.controlData["X"] != "10" {
		t.Errorf("Expected X offset '10', got %s", tb.cmd.controlData["X"])
	}
	if tb.cmd.controlData["Y"] != "20" {
		t.Errorf("Expected Y offset '20', got %s", tb.cmd.controlData["Y"])
	}
}

// TestTransmitBuilder_DisplaySize tests display size
func TestTransmitBuilder_DisplaySize(t *testing.T) {
	tb := NewTransmit()
	tb.DisplaySize(20, 15)
	if tb.cmd.controlData["c"] != "20" {
		t.Errorf("Expected columns '20', got %s", tb.cmd.controlData["c"])
	}
	if tb.cmd.controlData["r"] != "15" {
		t.Errorf("Expected rows '15', got %s", tb.cmd.controlData["r"])
	}
}

// TestTransmitBuilder_SourceRect tests source rectangle
func TestTransmitBuilder_SourceRect(t *testing.T) {
	tb := NewTransmit()
	tb.SourceRect(100, 200, 300, 400)
	if tb.cmd.controlData["x"] != "100" {
		t.Errorf("Expected x '100', got %s", tb.cmd.controlData["x"])
	}
	if tb.cmd.controlData["y"] != "200" {
		t.Errorf("Expected y '200', got %s", tb.cmd.controlData["y"])
	}
	if tb.cmd.controlData["w"] != "300" {
		t.Errorf("Expected width '300', got %s", tb.cmd.controlData["w"])
	}
	if tb.cmd.controlData["h"] != "400" {
		t.Errorf("Expected height '400', got %s", tb.cmd.controlData["h"])
	}
}

// TestTransmitBuilder_ZIndex tests z-index
func TestTransmitBuilder_ZIndex(t *testing.T) {
	tb := NewTransmit()
	tb.ZIndex(-5)
	if tb.cmd.controlData["z"] != "-5" {
		t.Errorf("Expected z-index '-5', got %s", tb.cmd.controlData["z"])
	}
}

// TestTransmitBuilder_CursorMovement tests cursor movement
func TestTransmitBuilder_CursorMovement(t *testing.T) {
	tb := NewTransmit()
	tb.CursorMovement(false)
	if tb.cmd.controlData["C"] != "1" {
		t.Errorf("Expected cursor movement '1' (no move), got %s", tb.cmd.controlData["C"])
	}

	tb2 := NewTransmit()
	tb2.CursorMovement(true)
	if tb2.cmd.controlData["C"] != "0" {
		t.Errorf("Expected cursor movement '0' (move), got %s", tb2.cmd.controlData["C"])
	}
}

// TestTransmitBuilder_RelativeTo tests relative positioning
func TestTransmitBuilder_RelativeTo(t *testing.T) {
	tb := NewTransmit()
	tb.RelativeTo(100, 1, 10, 5)
	if tb.cmd.controlData["P"] != "100" {
		t.Errorf("Expected parent image ID '100', got %s", tb.cmd.controlData["P"])
	}
	if tb.cmd.controlData["Q"] != "1" {
		t.Errorf("Expected parent placement ID '1', got %s", tb.cmd.controlData["Q"])
	}
	if tb.cmd.controlData["H"] != "10" {
		t.Errorf("Expected H offset '10', got %s", tb.cmd.controlData["H"])
	}
	if tb.cmd.controlData["V"] != "5" {
		t.Errorf("Expected V offset '5', got %s", tb.cmd.controlData["V"])
	}
}

// TestTransmitBuilder_Build tests building the command
func TestTransmitBuilder_Build(t *testing.T) {
	tb := NewTransmitDisplay()
	tb.Format(FormatPNG)
	tb.ImageID(10)
	tb.TransmitDirect([]byte("test"))

	cmd := tb.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}

	// Check payload was set
	if string(cmd.payload) != "test" {
		t.Errorf("Expected payload 'test', got %s", string(cmd.payload))
	}
}

// TestTransmitBuilder_CompleteFlow tests a complete transmit flow
func TestTransmitBuilder_CompleteFlow(t *testing.T) {
	cmd := NewTransmitDisplay().
		ImageID(42).
		Format(FormatPNG).
		TransmitDirect([]byte("PNG data here")).
		DisplaySize(20, 15).
		ZIndex(1).
		Build()

	encoded := cmd.Encode()

	// Verify the encoding contains expected keys
	if !strings.Contains(encoded, "a=T") {
		t.Error("Should contain action=T")
	}
	if !strings.Contains(encoded, "i=42") {
		t.Error("Should contain image ID")
	}
	if !strings.Contains(encoded, "f=100") {
		t.Error("Should contain format")
	}
	if !strings.Contains(encoded, "c=20") {
		t.Error("Should contain columns")
	}
	if !strings.Contains(encoded, "r=15") {
		t.Error("Should contain rows")
	}
	if !strings.Contains(encoded, "z=1") {
		t.Error("Should contain z-index")
	}
}
