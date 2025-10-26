// Package kgp provides complete Go bindings for the Kitty Graphics Protocol.
//
// The Kitty Graphics Protocol enables terminal applications to render pixel-based
// graphics alongside text. This package provides a type-safe, idiomatic Go API
// for working with the protocol.
//
// Basic usage:
//
//	// Display a PNG image
//	cmd := kgp.NewTransmit().
//		Format(kgp.FormatPNG).
//		TransmitDirect(imageData).
//		Build()
//	fmt.Print(cmd.Encode())
//
//	// Create multiple placements of an image
//	transmitCmd := kgp.NewTransmit().
//		ImageID(10).
//		Format(kgp.FormatPNG).
//		TransmitDirect(pngData).
//		Build()
//	fmt.Print(transmitCmd.Encode())
//
//	placementCmd := kgp.NewPut(10).
//		DisplaySize(20, 15).
//		Build()
//	fmt.Print(placementCmd.Encode())
package kgp

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// Action represents the action type for a graphics command.
type Action string

const (
	// ActionTransmit uploads image data to the terminal (lowercase = transmit only)
	ActionTransmit Action = "t"
	// ActionTransmitDisplay uploads and displays image data (uppercase = transmit + display)
	ActionTransmitDisplay Action = "T"
	// ActionPut creates a new placement of an existing image
	ActionPut Action = "p"
	// ActionDelete removes images or placements
	ActionDelete Action = "d"
	// ActionFrame transmits animation frame data
	ActionFrame Action = "f"
	// ActionAnimate controls animation playback
	ActionAnimate Action = "a"
	// ActionCompose composes animation frames
	ActionCompose Action = "c"
	// ActionQuery queries terminal capabilities
	ActionQuery Action = "q"
)

// Format represents the image data format.
type Format uint32

const (
	// FormatRGB represents 24-bit RGB (3 bytes per pixel)
	FormatRGB Format = 24
	// FormatRGBA represents 32-bit RGBA (4 bytes per pixel, default)
	FormatRGBA Format = 32
	// FormatPNG represents PNG format with embedded dimensions
	FormatPNG Format = 100
)

// Compression represents the compression method for image data.
type Compression string

const (
	// CompressionZlib represents ZLIB deflate compression (RFC 1950)
	CompressionZlib Compression = "z"
)

// TransmitMedium represents how image data is transmitted.
type TransmitMedium string

const (
	// TransmitDirect embeds data directly in the escape code (default)
	TransmitDirect TransmitMedium = "d"
	// TransmitFile reads from a regular file
	TransmitFile TransmitMedium = "f"
	// TransmitTemp reads from a temporary file (terminal deletes after reading)
	TransmitTemp TransmitMedium = "t"
	// TransmitSharedMem reads from POSIX shared memory
	TransmitSharedMem TransmitMedium = "s"
)

// DeleteMode specifies what to delete and whether to free data.
type DeleteMode string

const (
	// DeleteAllPlacements deletes all placements (preserve data)
	DeleteAllPlacements DeleteMode = "a"
	// DeleteAllPlacementsFree deletes all placements and frees data
	DeleteAllPlacementsFree DeleteMode = "A"
	// DeleteByImageID deletes by image ID (preserve data)
	DeleteByImageID DeleteMode = "i"
	// DeleteByImageIDFree deletes by image ID and frees data
	DeleteByImageIDFree DeleteMode = "I"
	// DeleteByImageNumber deletes by image number (preserve data)
	DeleteByImageNumber DeleteMode = "n"
	// DeleteByImageNumberFree deletes by image number and frees data
	DeleteByImageNumberFree DeleteMode = "N"
	// DeleteByCursor deletes at cursor position (preserve data)
	DeleteByCursor DeleteMode = "c"
	// DeleteByCursorFree deletes at cursor position and frees data
	DeleteByCursorFree DeleteMode = "C"
	// DeleteByPlacementID deletes by placement ID (preserve data)
	DeleteByPlacementID DeleteMode = "p"
	// DeleteByPlacementIDFree deletes by placement ID and frees data
	DeleteByPlacementIDFree DeleteMode = "P"
	// DeleteByCell deletes by cell coordinates (preserve data)
	DeleteByCell DeleteMode = "q"
	// DeleteByCellFree deletes by cell coordinates and frees data
	DeleteByCellFree DeleteMode = "Q"
	// DeleteByColumn deletes by column (preserve data)
	DeleteByColumn DeleteMode = "x"
	// DeleteByColumnFree deletes by column and frees data
	DeleteByColumnFree DeleteMode = "X"
	// DeleteByRow deletes by row (preserve data)
	DeleteByRow DeleteMode = "y"
	// DeleteByRowFree deletes by row and frees data
	DeleteByRowFree DeleteMode = "Y"
	// DeleteByZIndex deletes by z-index (preserve data)
	DeleteByZIndex DeleteMode = "z"
	// DeleteByZIndexFree deletes by z-index and frees data
	DeleteByZIndexFree DeleteMode = "Z"
)

// AnimationState controls animation playback.
type AnimationState uint32

const (
	// AnimationStop stops the animation at current frame
	AnimationStop AnimationState = 1
	// AnimationLoading puts animation in loading mode (waits for more frames)
	AnimationLoading AnimationState = 2
	// AnimationLoop plays animation with looping
	AnimationLoop AnimationState = 3
)

// CompositionMode controls how frames blend.
type CompositionMode uint32

const (
	// CompositionBlend performs alpha blending (default)
	CompositionBlend CompositionMode = 0
	// CompositionReplace replaces pixels without blending
	CompositionReplace CompositionMode = 1
)

// ResponseSuppression controls which responses the terminal sends.
type ResponseSuppression uint32

const (
	// ResponseAll sends all responses (default)
	ResponseAll ResponseSuppression = 0
	// ResponseErrorsOnly suppresses OK responses
	ResponseErrorsOnly ResponseSuppression = 1
	// ResponseOKOnly suppresses error responses
	ResponseOKOnly ResponseSuppression = 2
)

// Command represents a complete Kitty Graphics Protocol command.
type Command struct {
	controlData map[string]string
	payload     []byte
}

// NewCommand creates a new command with the specified action.
func NewCommand(action Action) *Command {
	return &Command{
		controlData: map[string]string{
			"a": string(action),
		},
	}
}

// SetKey sets a control data key-value pair.
func (c *Command) SetKey(key, value string) *Command {
	c.controlData[key] = value
	return c
}

// SetKeyInt sets a control data key with an integer value.
func (c *Command) SetKeyInt(key string, value int) *Command {
	c.controlData[key] = strconv.Itoa(value)
	return c
}

// SetKeyUint32 sets a control data key with a uint32 value.
func (c *Command) SetKeyUint32(key string, value uint32) *Command {
	c.controlData[key] = strconv.FormatUint(uint64(value), 10)
	return c
}

// SetPayload sets the payload data (will be base64-encoded).
func (c *Command) SetPayload(data []byte) *Command {
	c.payload = data
	return c
}

// Encode generates the complete escape sequence for this command.
func (c *Command) Encode() string {
	var sb strings.Builder

	// Start APC sequence: ESC_G
	sb.WriteString("\x1b_G")

	// Write control data as comma-separated key=value pairs
	first := true
	for k, v := range c.controlData {
		if !first {
			sb.WriteString(",")
		}
		first = false
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
	}

	// Add payload if present
	if len(c.payload) > 0 {
		sb.WriteString(";")
		sb.WriteString(base64.StdEncoding.EncodeToString(c.payload))
	}

	// End APC sequence: ESC\
	sb.WriteString("\x1b\\")

	return sb.String()
}

// EncodeChunked generates multiple escape sequences for chunked transmission.
// The payload is split into chunks of maxChunkSize (must be ≤4096 and divisible by 4).
func (c *Command) EncodeChunked(maxChunkSize int) []string {
	if maxChunkSize > 4096 || maxChunkSize%4 != 0 {
		panic("maxChunkSize must be ≤4096 and divisible by 4")
	}

	if len(c.payload) == 0 {
		return []string{c.Encode()}
	}

	// Base64 encode the entire payload first
	encoded := base64.StdEncoding.EncodeToString(c.payload)

	var chunks []string

	for i := 0; i < len(encoded); i += maxChunkSize {
		end := i + maxChunkSize
		isLast := end >= len(encoded)
		if isLast {
			end = len(encoded)
		}

		chunk := encoded[i:end]

		var sb strings.Builder
		sb.WriteString("\x1b_G")

		if i == 0 {
			// First chunk: include all control data
			first := true
			for k, v := range c.controlData {
				if !first {
					sb.WriteString(",")
				}
				first = false
				sb.WriteString(k)
				sb.WriteString("=")
				sb.WriteString(v)
			}

			// Add m=1 if not last chunk
			if !isLast {
				sb.WriteString(",m=1")
			}
		} else {
			// Subsequent chunks: only m key
			if !isLast {
				sb.WriteString("m=1")
			} else {
				sb.WriteString("m=0")
			}
		}

		sb.WriteString(";")
		sb.WriteString(chunk)
		sb.WriteString("\x1b\\")

		chunks = append(chunks, sb.String())
	}

	return chunks
}

// Response represents a terminal response to a graphics command.
type Response struct {
	ImageID     uint32
	PlacementID uint32
	Success     bool
	ErrorCode   string
	Message     string
}

// ParseResponse parses a terminal response.
// Format: ESC_Gi=<id>[,p=<pid>];[OK|ERROR_CODE:message]ESC\
func ParseResponse(response string) (*Response, error) {
	// Strip APC markers
	response = strings.TrimPrefix(response, "\x1b_G")
	response = strings.TrimSuffix(response, "\x1b\\")

	// Split into control data and status
	parts := strings.SplitN(response, ";", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid response format")
	}

	controlData := parts[0]
	status := parts[1]

	resp := &Response{}

	// Parse control data
	for _, pair := range strings.Split(controlData, ",") {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			continue
		}

		switch kv[0] {
		case "i":
			if val, err := strconv.ParseUint(kv[1], 10, 32); err == nil {
				resp.ImageID = uint32(val)
			}
		case "p":
			if val, err := strconv.ParseUint(kv[1], 10, 32); err == nil {
				resp.PlacementID = uint32(val)
			}
		}
	}

	// Parse status
	if status == "OK" {
		resp.Success = true
	} else if strings.Contains(status, ":") {
		parts := strings.SplitN(status, ":", 2)
		resp.ErrorCode = parts[0]
		resp.Message = parts[1]
	}

	return resp, nil
}
