package kgp

import (
	"image"
	"image/color"
	"testing"
)

// TestImageToRGBA tests converting image to RGBA bytes
func TestImageToRGBA(t *testing.T) {
	// Create a simple 2x2 image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	img.Set(1, 0, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img.Set(0, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	img.Set(1, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})

	rgba := ImageToRGBA(img)

	// Should be 2x2 * 4 bytes = 16 bytes
	if len(rgba) != 16 {
		t.Errorf("Expected 16 bytes, got %d", len(rgba))
	}

	// Check first pixel (red)
	if rgba[0] != 255 || rgba[1] != 0 || rgba[2] != 0 || rgba[3] != 255 {
		t.Errorf("First pixel should be red, got R=%d G=%d B=%d A=%d", rgba[0], rgba[1], rgba[2], rgba[3])
	}

	// Check second pixel (green)
	if rgba[4] != 0 || rgba[5] != 255 || rgba[6] != 0 || rgba[7] != 255 {
		t.Errorf("Second pixel should be green, got R=%d G=%d B=%d A=%d", rgba[4], rgba[5], rgba[6], rgba[7])
	}
}

// TestImageToRGB tests converting image to RGB bytes
func TestImageToRGB(t *testing.T) {
	// Create a simple 2x2 image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	img.Set(1, 0, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img.Set(0, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	img.Set(1, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})

	rgb := ImageToRGB(img)

	// Should be 2x2 * 3 bytes = 12 bytes
	if len(rgb) != 12 {
		t.Errorf("Expected 12 bytes, got %d", len(rgb))
	}

	// Check first pixel (red, no alpha)
	if rgb[0] != 255 || rgb[1] != 0 || rgb[2] != 0 {
		t.Errorf("First pixel should be red, got R=%d G=%d B=%d", rgb[0], rgb[1], rgb[2])
	}

	// Check second pixel (green, no alpha)
	if rgb[3] != 0 || rgb[4] != 255 || rgb[5] != 0 {
		t.Errorf("Second pixel should be green, got R=%d G=%d B=%d", rgb[3], rgb[4], rgb[5])
	}
}

// TestImageToPNG tests converting image to PNG bytes
func TestImageToPNG(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			img.Set(x, y, color.RGBA{R: 128, G: 128, B: 128, A: 255})
		}
	}

	pngData, err := ImageToPNG(img)
	if err != nil {
		t.Fatalf("ImageToPNG failed: %v", err)
	}

	if len(pngData) == 0 {
		t.Error("PNG data should not be empty")
	}

	// Check PNG signature (first 8 bytes)
	pngSignature := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	if len(pngData) < 8 {
		t.Fatal("PNG data too short")
	}
	for i := 0; i < 8; i++ {
		if pngData[i] != pngSignature[i] {
			t.Error("Invalid PNG signature")
			break
		}
	}
}

// TestCompressZlib tests zlib compression
func TestCompressZlib(t *testing.T) {
	data := []byte("test data to compress")
	compressed, err := CompressZlib(data)
	if err != nil {
		t.Fatalf("CompressZlib failed: %v", err)
	}

	if len(compressed) == 0 {
		t.Error("Compressed data should not be empty")
	}

	// Compressed data should be different from original
	if string(compressed) == string(data) {
		t.Error("Compressed data should differ from original")
	}
}

// TestTransmitImage tests the convenience function
func TestTransmitImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 5, 5))
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			img.Set(x, y, color.RGBA{R: 200, G: 100, B: 50, A: 255})
		}
	}

	cmd, err := TransmitImage(img)
	if err != nil {
		t.Fatalf("TransmitImage failed: %v", err)
	}

	if cmd == nil {
		t.Fatal("TransmitImage returned nil command")
	}

	// Check it's a transmit+display action
	if cmd.controlData["a"] != string(ActionTransmitDisplay) {
		t.Errorf("Expected action %s, got %s", ActionTransmitDisplay, cmd.controlData["a"])
	}

	// Check format is PNG
	if cmd.controlData["f"] != "100" {
		t.Errorf("Expected format 100 (PNG), got %s", cmd.controlData["f"])
	}

	// Check payload exists
	if len(cmd.payload) == 0 {
		t.Error("Command payload should not be empty")
	}
}

// TestTransmitImageWithID tests the convenience function with ID
func TestTransmitImageWithID(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 5, 5))
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			img.Set(x, y, color.RGBA{R: 200, G: 100, B: 50, A: 255})
		}
	}

	cmd, err := TransmitImageWithID(img, 42)
	if err != nil {
		t.Fatalf("TransmitImageWithID failed: %v", err)
	}

	if cmd == nil {
		t.Fatal("TransmitImageWithID returned nil command")
	}

	// Check image ID
	if cmd.controlData["i"] != "42" {
		t.Errorf("Expected image ID 42, got %s", cmd.controlData["i"])
	}
}

// TestTransmitImageRGBA tests RGBA transmission
func TestTransmitImageRGBA(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			img.Set(x, y, color.RGBA{R: 128, G: 128, B: 128, A: 255})
		}
	}

	// Test without compression
	cmd, err := TransmitImageRGBA(img, false)
	if err != nil {
		t.Fatalf("TransmitImageRGBA failed: %v", err)
	}

	if cmd == nil {
		t.Fatal("TransmitImageRGBA returned nil command")
	}

	// Check format is RGBA
	if cmd.controlData["f"] != "32" {
		t.Errorf("Expected format 32 (RGBA), got %s", cmd.controlData["f"])
	}

	// Check dimensions
	if cmd.controlData["s"] != "10" || cmd.controlData["v"] != "10" {
		t.Errorf("Expected dimensions 10x10, got %sx%s", cmd.controlData["s"], cmd.controlData["v"])
	}

	// Test with compression
	cmdCompressed, err := TransmitImageRGBA(img, true)
	if err != nil {
		t.Fatalf("TransmitImageRGBA with compression failed: %v", err)
	}

	if cmdCompressed.controlData["o"] != "z" {
		t.Error("Expected compression to be set")
	}
}

// TestCreateRGBAColor tests RGBA color creation
func TestCreateRGBAColor(t *testing.T) {
	tests := []struct {
		r, g, b, a uint8
		want       uint32
	}{
		{255, 0, 0, 255, 0xFF0000FF},
		{0, 255, 0, 255, 0x00FF00FF},
		{0, 0, 255, 255, 0x0000FFFF},
		{128, 128, 128, 255, 0x808080FF},
		{255, 255, 255, 0, 0xFFFFFF00},
	}

	for _, tt := range tests {
		got := CreateRGBAColor(tt.r, tt.g, tt.b, tt.a)
		if got != tt.want {
			t.Errorf("CreateRGBAColor(%d,%d,%d,%d) = 0x%X, want 0x%X",
				tt.r, tt.g, tt.b, tt.a, got, tt.want)
		}
	}
}

// TestSolidColorImage tests solid color image creation
func TestSolidColorImage(t *testing.T) {
	rgba := SolidColorImage(10, 10, 255, 128, 64, 255)

	// Should be 10x10 * 4 bytes = 400 bytes
	if len(rgba) != 400 {
		t.Errorf("Expected 400 bytes, got %d", len(rgba))
	}

	// Check first pixel
	if rgba[0] != 255 || rgba[1] != 128 || rgba[2] != 64 || rgba[3] != 255 {
		t.Errorf("First pixel incorrect: R=%d G=%d B=%d A=%d", rgba[0], rgba[1], rgba[2], rgba[3])
	}

	// Check last pixel (should be same as first)
	lastIdx := len(rgba) - 4
	if rgba[lastIdx] != 255 || rgba[lastIdx+1] != 128 || rgba[lastIdx+2] != 64 || rgba[lastIdx+3] != 255 {
		t.Errorf("Last pixel incorrect: R=%d G=%d B=%d A=%d",
			rgba[lastIdx], rgba[lastIdx+1], rgba[lastIdx+2], rgba[lastIdx+3])
	}

	// Verify all pixels are the same
	for i := 0; i < len(rgba); i += 4 {
		if rgba[i] != 255 || rgba[i+1] != 128 || rgba[i+2] != 64 || rgba[i+3] != 255 {
			t.Errorf("Pixel at index %d is incorrect", i/4)
			break
		}
	}
}

// TestSolidColorImageDimensions tests various dimensions
func TestSolidColorImageDimensions(t *testing.T) {
	tests := []struct {
		width, height int
		expectedLen   int
	}{
		{1, 1, 4},
		{10, 10, 400},
		{100, 100, 40000},
		{5, 3, 60},
	}

	for _, tt := range tests {
		rgba := SolidColorImage(tt.width, tt.height, 0, 0, 0, 255)
		if len(rgba) != tt.expectedLen {
			t.Errorf("SolidColorImage(%dx%d) = %d bytes, want %d bytes",
				tt.width, tt.height, len(rgba), tt.expectedLen)
		}
	}
}
