package kgp

import (
	"strings"
	"testing"
)

// TestNewCommand tests command creation
func TestNewCommand(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	if cmd == nil {
		t.Fatal("NewCommand returned nil")
	}
	if cmd.controlData["a"] != string(ActionTransmit) {
		t.Errorf("Expected action %s, got %s", ActionTransmit, cmd.controlData["a"])
	}
}

// TestCommandSetKey tests setting string key-value pairs
func TestCommandSetKey(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	cmd.SetKey("test", "value")
	if cmd.controlData["test"] != "value" {
		t.Errorf("Expected 'value', got %s", cmd.controlData["test"])
	}
}

// TestCommandSetKeyInt tests setting integer values
func TestCommandSetKeyInt(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	cmd.SetKeyInt("width", 800)
	if cmd.controlData["width"] != "800" {
		t.Errorf("Expected '800', got %s", cmd.controlData["width"])
	}
}

// TestCommandSetKeyUint32 tests setting uint32 values
func TestCommandSetKeyUint32(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	cmd.SetKeyUint32("id", 12345)
	if cmd.controlData["id"] != "12345" {
		t.Errorf("Expected '12345', got %s", cmd.controlData["id"])
	}
}

// TestCommandSetPayload tests setting payload data
func TestCommandSetPayload(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	testData := []byte("test payload")
	cmd.SetPayload(testData)
	if string(cmd.payload) != "test payload" {
		t.Errorf("Expected 'test payload', got %s", string(cmd.payload))
	}
}

// TestCommandEncode tests basic command encoding
func TestCommandEncode(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	encoded := cmd.Encode()

	// Should start with ESC_G
	if !strings.HasPrefix(encoded, "\x1b_G") {
		t.Error("Encoded command should start with ESC_G")
	}

	// Should end with ESC\
	if !strings.HasSuffix(encoded, "\x1b\\") {
		t.Error("Encoded command should end with ESC\\")
	}

	// Should contain action
	if !strings.Contains(encoded, "a=t") {
		t.Error("Encoded command should contain action")
	}
}

// TestCommandEncodeWithPayload tests encoding with payload
func TestCommandEncodeWithPayload(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	cmd.SetPayload([]byte("test"))
	encoded := cmd.Encode()

	// Should contain semicolon separator
	if !strings.Contains(encoded, ";") {
		t.Error("Encoded command with payload should contain semicolon")
	}

	// Should contain base64 encoded data
	if !strings.Contains(encoded, "dGVzdA==") { // "test" in base64
		t.Error("Encoded command should contain base64 payload")
	}
}

// TestCommandEncodeChunked tests chunked encoding
func TestCommandEncodeChunked(t *testing.T) {
	cmd := NewCommand(ActionTransmit)
	// Create payload larger than chunk size
	payload := make([]byte, 100)
	for i := range payload {
		payload[i] = byte(i % 256)
	}
	cmd.SetPayload(payload)

	chunks := cmd.EncodeChunked(64)

	if len(chunks) < 2 {
		t.Error("Expected multiple chunks for large payload")
	}

	// First chunk should have m=1
	if !strings.Contains(chunks[0], "m=1") {
		t.Error("First chunk should have m=1")
	}

	// Last chunk should NOT have m=1 (implicitly m=0)
	lastChunk := chunks[len(chunks)-1]
	if strings.Contains(lastChunk, "m=1") {
		t.Error("Last chunk should not have m=1")
	}
}

// TestCommandEncodeChunkedPanic tests that invalid chunk size panics
func TestCommandEncodeChunkedPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid chunk size")
		}
	}()

	cmd := NewCommand(ActionTransmit)
	cmd.SetPayload([]byte("test"))
	cmd.EncodeChunked(5000) // Too large
}

// TestCommandEncodeChunkedNotDivisibleBy4 tests chunk size validation
func TestCommandEncodeChunkedNotDivisibleBy4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for chunk size not divisible by 4")
		}
	}()

	cmd := NewCommand(ActionTransmit)
	cmd.SetPayload([]byte("test"))
	cmd.EncodeChunked(102) // Not divisible by 4
}

// TestParseResponse tests response parsing
func TestParseResponse(t *testing.T) {
	tests := []struct {
		name        string
		response    string
		wantSuccess bool
		wantImageID uint32
		wantPlaceID uint32
		wantError   string
		wantMessage string
	}{
		{
			name:        "Success response",
			response:    "\x1b_Gi=10,p=1;OK\x1b\\",
			wantSuccess: true,
			wantImageID: 10,
			wantPlaceID: 1,
		},
		{
			name:        "Error response",
			response:    "\x1b_Gi=5;ENOSPC:Storage quota exceeded\x1b\\",
			wantSuccess: false,
			wantImageID: 5,
			wantError:   "ENOSPC",
			wantMessage: "Storage quota exceeded",
		},
		{
			name:        "Simple success",
			response:    "\x1b_G;OK\x1b\\",
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ParseResponse(tt.response)
			if err != nil {
				t.Fatalf("ParseResponse error: %v", err)
			}

			if resp.Success != tt.wantSuccess {
				t.Errorf("Success = %v, want %v", resp.Success, tt.wantSuccess)
			}
			if resp.ImageID != tt.wantImageID {
				t.Errorf("ImageID = %d, want %d", resp.ImageID, tt.wantImageID)
			}
			if resp.PlacementID != tt.wantPlaceID {
				t.Errorf("PlacementID = %d, want %d", resp.PlacementID, tt.wantPlaceID)
			}
			if resp.ErrorCode != tt.wantError {
				t.Errorf("ErrorCode = %s, want %s", resp.ErrorCode, tt.wantError)
			}
			if resp.Message != tt.wantMessage {
				t.Errorf("Message = %s, want %s", resp.Message, tt.wantMessage)
			}
		})
	}
}

// TestParseResponseInvalid tests invalid response handling
func TestParseResponseInvalid(t *testing.T) {
	_, err := ParseResponse("invalid")
	if err == nil {
		t.Error("Expected error for invalid response")
	}
}

// TestActions tests action constants
func TestActions(t *testing.T) {
	actions := map[Action]string{
		ActionTransmit:        "t",
		ActionTransmitDisplay: "T",
		ActionPut:             "p",
		ActionDelete:          "d",
		ActionFrame:           "f",
		ActionAnimate:         "a",
		ActionCompose:         "c",
		ActionQuery:           "q",
	}

	for action, expected := range actions {
		if string(action) != expected {
			t.Errorf("Action %v = %s, want %s", action, string(action), expected)
		}
	}
}

// TestFormats tests format constants
func TestFormats(t *testing.T) {
	if FormatRGB != 24 {
		t.Errorf("FormatRGB = %d, want 24", FormatRGB)
	}
	if FormatRGBA != 32 {
		t.Errorf("FormatRGBA = %d, want 32", FormatRGBA)
	}
	if FormatPNG != 100 {
		t.Errorf("FormatPNG = %d, want 100", FormatPNG)
	}
}

// TestAnimationStates tests animation state constants
func TestAnimationStates(t *testing.T) {
	if AnimationStop != 1 {
		t.Errorf("AnimationStop = %d, want 1", AnimationStop)
	}
	if AnimationLoading != 2 {
		t.Errorf("AnimationLoading = %d, want 2", AnimationLoading)
	}
	if AnimationLoop != 3 {
		t.Errorf("AnimationLoop = %d, want 3", AnimationLoop)
	}
}

// TestCompositionModes tests composition mode constants
func TestCompositionModes(t *testing.T) {
	if CompositionBlend != 0 {
		t.Errorf("CompositionBlend = %d, want 0", CompositionBlend)
	}
	if CompositionReplace != 1 {
		t.Errorf("CompositionReplace = %d, want 1", CompositionReplace)
	}
}
