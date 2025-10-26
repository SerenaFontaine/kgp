package kgp

import (
	"strings"
	"testing"
)

// TestNewFrame tests creating a frame builder
func TestNewFrame(t *testing.T) {
	fb := NewFrame(42)
	if fb == nil {
		t.Fatal("NewFrame returned nil")
	}
	if fb.cmd.controlData["a"] != string(ActionFrame) {
		t.Errorf("Action should be %s, got %s", ActionFrame, fb.cmd.controlData["a"])
	}
	if fb.cmd.controlData["i"] != "42" {
		t.Errorf("Image ID should be '42', got %s", fb.cmd.controlData["i"])
	}
}

// TestFrameBuilder_FrameData tests setting frame data
func TestFrameBuilder_FrameData(t *testing.T) {
	fb := NewFrame(10)
	data := []byte("frame data")
	fb.FrameData(data)
	if string(fb.frameData) != "frame data" {
		t.Errorf("Expected 'frame data', got %s", string(fb.frameData))
	}
}

// TestFrameBuilder_Format tests setting format
func TestFrameBuilder_Format(t *testing.T) {
	fb := NewFrame(10)
	fb.Format(FormatRGBA)
	if fb.cmd.controlData["f"] != "32" {
		t.Errorf("Expected format '32', got %s", fb.cmd.controlData["f"])
	}
}

// TestFrameBuilder_Dimensions tests setting dimensions
func TestFrameBuilder_Dimensions(t *testing.T) {
	fb := NewFrame(10)
	fb.Dimensions(800, 600)
	if fb.cmd.controlData["s"] != "800" {
		t.Errorf("Expected width '800', got %s", fb.cmd.controlData["s"])
	}
	if fb.cmd.controlData["v"] != "600" {
		t.Errorf("Expected height '600', got %s", fb.cmd.controlData["v"])
	}
}

// TestFrameBuilder_BackgroundFrame tests setting background frame
func TestFrameBuilder_BackgroundFrame(t *testing.T) {
	fb := NewFrame(10)
	fb.BackgroundFrame(0)
	if fb.cmd.controlData["c"] != "0" {
		t.Errorf("Expected background frame '0', got %s", fb.cmd.controlData["c"])
	}
}

// TestFrameBuilder_Gap tests setting gap
func TestFrameBuilder_Gap(t *testing.T) {
	fb := NewFrame(10)
	fb.Gap(100)
	if fb.cmd.controlData["z"] != "100" {
		t.Errorf("Expected gap '100', got %s", fb.cmd.controlData["z"])
	}
}

// TestFrameBuilder_Composition tests setting composition mode
func TestFrameBuilder_Composition(t *testing.T) {
	fb := NewFrame(10)
	fb.Composition(CompositionReplace)
	if fb.cmd.controlData["X"] != "1" {
		t.Errorf("Expected composition '1', got %s", fb.cmd.controlData["X"])
	}
}

// TestFrameBuilder_BackgroundColor tests setting background color
func TestFrameBuilder_BackgroundColor(t *testing.T) {
	fb := NewFrame(10)
	rgba := CreateRGBAColor(255, 128, 64, 255)
	fb.BackgroundColor(rgba)
	// Just verify the value was set, actual value may vary by implementation
	if fb.cmd.controlData["Y"] == "" {
		t.Error("Background color should be set")
	}
}

// TestFrameBuilder_ResponseSuppression tests response suppression
func TestFrameBuilder_ResponseSuppression(t *testing.T) {
	fb := NewFrame(10)
	fb.ResponseSuppression(ResponseErrorsOnly)
	if fb.cmd.controlData["q"] != "1" {
		t.Errorf("Expected response suppression '1', got %s", fb.cmd.controlData["q"])
	}
}

// TestFrameBuilder_Build tests building the command
func TestFrameBuilder_Build(t *testing.T) {
	fb := NewFrame(10)
	fb.FrameData([]byte("test"))

	cmd := fb.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}

	// Check payload was set
	if string(cmd.payload) != "test" {
		t.Errorf("Expected payload 'test', got %s", string(cmd.payload))
	}
}

// TestNewAnimate tests creating an animate builder
func TestNewAnimate(t *testing.T) {
	ab := NewAnimate(42)
	if ab == nil {
		t.Fatal("NewAnimate returned nil")
	}
	if ab.cmd.controlData["a"] != string(ActionAnimate) {
		t.Errorf("Action should be %s, got %s", ActionAnimate, ab.cmd.controlData["a"])
	}
	if ab.cmd.controlData["i"] != "42" {
		t.Errorf("Image ID should be '42', got %s", ab.cmd.controlData["i"])
	}
}

// TestAnimateBuilder_State tests setting animation state
func TestAnimateBuilder_State(t *testing.T) {
	ab := NewAnimate(10)
	ab.State(AnimationLoop)
	if ab.cmd.controlData["s"] != "3" {
		t.Errorf("Expected state '3', got %s", ab.cmd.controlData["s"])
	}
}

// TestAnimateBuilder_LoopCount tests setting loop count
func TestAnimateBuilder_LoopCount(t *testing.T) {
	ab := NewAnimate(10)
	ab.LoopCount(5)
	if ab.cmd.controlData["v"] != "5" {
		t.Errorf("Expected loop count '5', got %s", ab.cmd.controlData["v"])
	}
}

// TestAnimateBuilder_GapOverride tests setting gap override
func TestAnimateBuilder_GapOverride(t *testing.T) {
	ab := NewAnimate(10)
	ab.GapOverride(50)
	if ab.cmd.controlData["z"] != "50" {
		t.Errorf("Expected gap override '50', got %s", ab.cmd.controlData["z"])
	}
}

// TestAnimateBuilder_Frame tests setting frame number
func TestAnimateBuilder_Frame(t *testing.T) {
	ab := NewAnimate(10)
	ab.Frame(3)
	if ab.cmd.controlData["c"] != "3" {
		t.Errorf("Expected frame '3', got %s", ab.cmd.controlData["c"])
	}
}

// TestAnimateBuilder_ResponseSuppression tests response suppression
func TestAnimateBuilder_ResponseSuppression(t *testing.T) {
	ab := NewAnimate(10)
	ab.ResponseSuppression(ResponseOKOnly)
	if ab.cmd.controlData["q"] != "2" {
		t.Errorf("Expected response suppression '2', got %s", ab.cmd.controlData["q"])
	}
}

// TestAnimateBuilder_Build tests building the command
func TestAnimateBuilder_Build(t *testing.T) {
	ab := NewAnimate(10)
	ab.State(AnimationLoop)

	cmd := ab.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}
}

// TestNewCompose tests creating a compose builder
func TestNewCompose(t *testing.T) {
	cb := NewCompose(42)
	if cb == nil {
		t.Fatal("NewCompose returned nil")
	}
	if cb.cmd.controlData["a"] != string(ActionCompose) {
		t.Errorf("Action should be %s, got %s", ActionCompose, cb.cmd.controlData["a"])
	}
	if cb.cmd.controlData["i"] != "42" {
		t.Errorf("Image ID should be '42', got %s", cb.cmd.controlData["i"])
	}
}

// TestComposeBuilder_BackgroundFrame tests setting background frame
func TestComposeBuilder_BackgroundFrame(t *testing.T) {
	cb := NewCompose(10)
	cb.BackgroundFrame(2)
	if cb.cmd.controlData["c"] != "2" {
		t.Errorf("Expected background frame '2', got %s", cb.cmd.controlData["c"])
	}
}

// TestComposeBuilder_Composition tests setting composition mode
func TestComposeBuilder_Composition(t *testing.T) {
	cb := NewCompose(10)
	cb.Composition(CompositionBlend)
	if cb.cmd.controlData["X"] != "0" {
		t.Errorf("Expected composition '0', got %s", cb.cmd.controlData["X"])
	}
}

// TestComposeBuilder_BackgroundColor tests setting background color
func TestComposeBuilder_BackgroundColor(t *testing.T) {
	cb := NewCompose(10)
	rgba := CreateRGBAColor(100, 150, 200, 255)
	cb.BackgroundColor(rgba)
	// Just verify the value was set, actual value may vary by implementation
	if cb.cmd.controlData["Y"] == "" {
		t.Error("Background color should be set")
	}
}

// TestComposeBuilder_ResponseSuppression tests response suppression
func TestComposeBuilder_ResponseSuppression(t *testing.T) {
	cb := NewCompose(10)
	cb.ResponseSuppression(ResponseAll)
	if cb.cmd.controlData["q"] != "0" {
		t.Errorf("Expected response suppression '0', got %s", cb.cmd.controlData["q"])
	}
}

// TestComposeBuilder_Build tests building the command
func TestComposeBuilder_Build(t *testing.T) {
	cb := NewCompose(10)
	cb.BackgroundFrame(0)

	cmd := cb.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}
}

// TestPlayAnimation tests helper function
func TestPlayAnimation(t *testing.T) {
	cmd := PlayAnimation(10)
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=a") {
		t.Error("Should contain action=a")
	}
	if !strings.Contains(encoded, "i=10") {
		t.Error("Should contain image ID 10")
	}
	if !strings.Contains(encoded, "s=3") {
		t.Error("Should contain state=3 (loop)")
	}
	if !strings.Contains(encoded, "v=2") {
		t.Error("Should contain loop count=2")
	}
}

// TestPlayAnimationLoop tests helper function
func TestPlayAnimationLoop(t *testing.T) {
	cmd := PlayAnimationLoop(20)
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=a") {
		t.Error("Should contain action=a")
	}
	if !strings.Contains(encoded, "i=20") {
		t.Error("Should contain image ID 20")
	}
	if !strings.Contains(encoded, "s=3") {
		t.Error("Should contain state=3 (loop)")
	}
	if !strings.Contains(encoded, "v=1") {
		t.Error("Should contain loop count=1 (infinite)")
	}
}

// TestStopAnimation tests helper function
func TestStopAnimation(t *testing.T) {
	cmd := StopAnimation(30)
	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=a") {
		t.Error("Should contain action=a")
	}
	if !strings.Contains(encoded, "i=30") {
		t.Error("Should contain image ID 30")
	}
	if !strings.Contains(encoded, "s=1") {
		t.Error("Should contain state=1 (stop)")
	}
}

// TestFrameBuilder_CompleteFlow tests a complete frame flow
func TestFrameBuilder_CompleteFlow(t *testing.T) {
	data := []byte("rgba data")
	cmd := NewFrame(300).
		Format(FormatRGBA).
		Dimensions(80, 80).
		FrameData(data).
		Gap(100).
		Composition(CompositionReplace).
		Build()

	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=f") {
		t.Error("Should contain action=f")
	}
	if !strings.Contains(encoded, "i=300") {
		t.Error("Should contain image ID 300")
	}
	if !strings.Contains(encoded, "f=32") {
		t.Error("Should contain format 32")
	}
	if !strings.Contains(encoded, "s=80") {
		t.Error("Should contain width 80")
	}
	if !strings.Contains(encoded, "v=80") {
		t.Error("Should contain height 80")
	}
	if !strings.Contains(encoded, "z=100") {
		t.Error("Should contain gap 100")
	}
	if !strings.Contains(encoded, "X=1") {
		t.Error("Should contain composition 1")
	}
}
