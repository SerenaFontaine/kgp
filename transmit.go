package kgp

import (
	"errors"
	"strings"
)

// ErrInvalidTempPath indicates a temporary transmit path is not protocol-compliant.
var ErrInvalidTempPath = errors.New(`temporary file path must contain "tty-graphics-protocol"`)

// TransmitBuilder builds a transmit action command.
type TransmitBuilder struct {
	cmd         *Command
	display     bool
	imageData   []byte
	compression Compression
}

// NewTransmit creates a new transmit action builder (transmit only, no display).
func NewTransmit() *TransmitBuilder {
	return &TransmitBuilder{
		cmd:     NewCommand(ActionTransmit),
		display: false,
	}
}

// NewTransmitDisplay creates a new transmit action builder (transmit and display).
func NewTransmitDisplay() *TransmitBuilder {
	tb := NewTransmit()
	tb.display = true
	tb.cmd.SetKey("a", string(ActionTransmitDisplay))
	return tb
}

// ImageID sets the image ID (auto-generated if not specified).
func (tb *TransmitBuilder) ImageID(id uint32) *TransmitBuilder {
	tb.cmd.SetKeyUint32("i", id)
	return tb
}

// ImageNumber sets the image number (application-specific, non-unique).
func (tb *TransmitBuilder) ImageNumber(num uint32) *TransmitBuilder {
	tb.cmd.SetKeyUint32("I", num)
	return tb
}

// Format sets the image format.
func (tb *TransmitBuilder) Format(format Format) *TransmitBuilder {
	tb.cmd.SetKeyUint32("f", uint32(format))
	return tb
}

// Dimensions sets the image dimensions (required for RGB/RGBA formats).
func (tb *TransmitBuilder) Dimensions(width, height int) *TransmitBuilder {
	tb.cmd.SetKeyInt("s", width)
	tb.cmd.SetKeyInt("v", height)
	return tb
}

// Compression enables ZLIB compression.
func (tb *TransmitBuilder) Compress() *TransmitBuilder {
	tb.cmd.SetKey("o", string(CompressionZlib))
	tb.compression = CompressionZlib
	return tb
}

// TransmitDirect embeds image data directly in the command (default).
func (tb *TransmitBuilder) TransmitDirect(data []byte) *TransmitBuilder {
	tb.cmd.SetKey("t", string(TransmitDirect))
	tb.imageData = data
	return tb
}

// TransmitFile reads image data from a file path.
func (tb *TransmitBuilder) TransmitFile(path string) *TransmitBuilder {
	tb.cmd.SetKey("t", string(TransmitFile))
	tb.imageData = []byte(path)
	return tb
}

// TransmitFileWithOffset reads image data from a file with offset and size.
func (tb *TransmitBuilder) TransmitFileWithOffset(path string, offset, size int) *TransmitBuilder {
	tb.cmd.SetKey("t", string(TransmitFile))
	tb.cmd.SetKeyInt("O", offset)
	tb.cmd.SetKeyInt("S", size)
	tb.imageData = []byte(path)
	return tb
}

// TransmitTemp reads image data from a temporary file (terminal deletes after reading).
// Path must contain "tty-graphics-protocol" for security.
func (tb *TransmitBuilder) TransmitTemp(path string) *TransmitBuilder {
	if _, err := tb.TryTransmitTemp(path); err != nil {
		panic(err.Error())
	}
	return tb
}

// ValidateTempPath validates that a temporary-file transmit path satisfies protocol requirements.
func ValidateTempPath(path string) error {
	if !strings.Contains(path, "tty-graphics-protocol") {
		return ErrInvalidTempPath
	}
	return nil
}

// TryTransmitTemp reads image data from a temporary file and returns an error on invalid paths.
func (tb *TransmitBuilder) TryTransmitTemp(path string) (*TransmitBuilder, error) {
	if err := ValidateTempPath(path); err != nil {
		return nil, err
	}
	tb.cmd.SetKey("t", string(TransmitTemp))
	tb.imageData = []byte(path)
	return tb, nil
}

// TransmitSharedMemory reads image data from POSIX shared memory.
func (tb *TransmitBuilder) TransmitSharedMemory(name string, size int) *TransmitBuilder {
	tb.cmd.SetKey("t", string(TransmitSharedMem))
	tb.cmd.SetKeyInt("S", size)
	tb.imageData = []byte(name)
	return tb
}

// PlacementID sets the placement ID for the initial placement.
func (tb *TransmitBuilder) PlacementID(id uint32) *TransmitBuilder {
	tb.cmd.SetKeyUint32("p", id)
	return tb
}

// VirtualPlacement creates an invisible placement for Unicode placeholder use.
func (tb *TransmitBuilder) VirtualPlacement() *TransmitBuilder {
	tb.cmd.SetKey("U", "1")
	return tb
}

// ResponseSuppression controls which responses the terminal sends.
func (tb *TransmitBuilder) ResponseSuppression(mode ResponseSuppression) *TransmitBuilder {
	tb.cmd.SetKeyUint32("q", uint32(mode))
	return tb
}

// CellOffset sets the pixel offset within the starting cell.
func (tb *TransmitBuilder) CellOffset(x, y int) *TransmitBuilder {
	tb.cmd.SetKeyInt("X", x)
	tb.cmd.SetKeyInt("Y", y)
	return tb
}

// DisplaySize sets the display size in terminal cells.
func (tb *TransmitBuilder) DisplaySize(columns, rows int) *TransmitBuilder {
	tb.cmd.SetKeyInt("c", columns)
	tb.cmd.SetKeyInt("r", rows)
	return tb
}

// SourceRect displays only a portion of the source image.
func (tb *TransmitBuilder) SourceRect(x, y, width, height int) *TransmitBuilder {
	tb.cmd.SetKeyInt("x", x)
	tb.cmd.SetKeyInt("y", y)
	tb.cmd.SetKeyInt("w", width)
	tb.cmd.SetKeyInt("h", height)
	return tb
}

// ZIndex sets the z-index (negative = below text, positive = above text).
func (tb *TransmitBuilder) ZIndex(z int) *TransmitBuilder {
	tb.cmd.SetKeyInt("z", z)
	return tb
}

// CursorMovement controls whether the cursor moves after placement.
func (tb *TransmitBuilder) CursorMovement(move bool) *TransmitBuilder {
	if move {
		tb.cmd.SetKey("C", "0")
	} else {
		tb.cmd.SetKey("C", "1")
	}
	return tb
}

// RelativeTo sets the parent placement for relative positioning.
func (tb *TransmitBuilder) RelativeTo(parentImageID, parentPlacementID uint32, offsetH, offsetV int) *TransmitBuilder {
	tb.cmd.SetKeyUint32("P", parentImageID)
	tb.cmd.SetKeyUint32("Q", parentPlacementID)
	tb.cmd.SetKeyInt("H", offsetH)
	tb.cmd.SetKeyInt("V", offsetV)
	return tb
}

// Build constructs the final command.
func (tb *TransmitBuilder) Build() *Command {
	if len(tb.imageData) > 0 {
		tb.cmd.SetPayload(tb.imageData)
	}
	return tb.cmd
}
