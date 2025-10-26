# KGP - Kitty Graphics Protocol Go Bindings

Go bindings for the [Kitty Graphics Protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/), enabling terminal applications to render pixel-based graphics alongside text.

## Features

- **Type-safe API** with fluent builders for all actions
- **Complete protocol coverage** including:
  - Image transmission (direct, file, temporary file, shared memory)
  - Multiple placements of images
  - Animation support (frames, playback control, composition)
  - Z-index layering (above/below text)
  - Relative positioning (parent-child relationships)
  - Cropping and scaling
  - Response parsing
- **Chunked transmission** for large images
- **Helper functions** for common operations
- **Image format support**: PNG, RGB, RGBA
- **Compression support**: ZLIB
- **Zero dependencies** (beyond Go standard library)

## Installation

```bash
go get github.com/SerenaFontaine/kgp
```

## Quick Start

### Display a PNG Image

```go
package main

import (
    "fmt"
    "os"
    "github.com/SerenaFontaine/kgp"
)

func main() {
    // Read PNG file
    data, _ := os.ReadFile("image.png")

    // Create and send command
    cmd := kgp.NewTransmitDisplay().
        Format(kgp.FormatPNG).
        TransmitDirect(data).
        Build()

    fmt.Print(cmd.Encode())
}
```

### Display from image.Image

```go
import (
    "fmt"
    "image"
    "github.com/SerenaFontaine/kgp"
)

func displayImage(img image.Image) {
    cmd, _ := kgp.TransmitImage(img)
    fmt.Print(cmd.Encode())
}
```

## Usage Examples

### Basic Image Display

```go
// Transmit and display PNG
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    Build()
fmt.Print(cmd.Encode())

// Transmit RGBA with dimensions
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    TransmitDirect(rgbaData).
    Build()
fmt.Print(cmd.Encode())
```

### Multiple Placements

```go
// Transmit image once
transmitCmd := kgp.NewTransmit().
    ImageID(10).
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    Build()
fmt.Print(transmitCmd.Encode())

// Create multiple placements
placement1 := kgp.NewPut(10).
    DisplaySize(20, 15).
    ZIndex(1).
    Build()
fmt.Print(placement1.Encode())

placement2 := kgp.NewPut(10).
    DisplaySize(10, 7).
    CellOffset(5, 5).
    Build()
fmt.Print(placement2.Encode())
```

### Cropping and Scaling

```go
// Display top-left quarter of image at 2x size
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    SourceRect(0, 0, 400, 300).  // Crop to 400x300
    DisplaySize(20, 15).           // Display in 20x15 cells
    Build()
fmt.Print(cmd.Encode())
```

### Z-Index Layering

```go
// Background image below text
bgCmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(bgData).
    ZIndex(-10).  // Negative = below text
    Build()
fmt.Print(bgCmd.Encode())

// Foreground image above text
fgCmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(fgData).
    ZIndex(10).   // Positive = above text
    Build()
fmt.Print(fgCmd.Encode())
```

### Deletion

```go
// Delete all placements (preserve data)
fmt.Print(kgp.DeleteAll().Encode())

// Delete specific image and free memory
fmt.Print(kgp.DeleteImageFree(10).Encode())

// Delete at cursor position
fmt.Print(kgp.DeleteAtCursor().Encode())

// Custom deletion by z-index
cmd := kgp.NewDelete(kgp.DeleteByZIndex).
    ZIndex(5).
    Build()
fmt.Print(cmd.Encode())
```

### Animation

```go
// Transmit base image
baseCmd := kgp.NewTransmitDisplay().
    ImageID(20).
    Format(kgp.FormatPNG).
    TransmitDirect(frame0Data).
    Build()
fmt.Print(baseCmd.Encode())

// Add animation frames
frame1Cmd := kgp.NewFrame(20).
    FrameData(frame1Data).
    Gap(100).  // 100ms delay
    Build()
fmt.Print(frame1Cmd.Encode())

frame2Cmd := kgp.NewFrame(20).
    FrameData(frame2Data).
    Gap(100).
    BackgroundFrame(0).
    Composition(kgp.CompositionBlend).
    Build()
fmt.Print(frame2Cmd.Encode())

// Play with infinite looping
fmt.Print(kgp.PlayAnimationLoop(20).Encode())

// Stop animation
fmt.Print(kgp.StopAnimation(20).Encode())

// Reset to first frame
fmt.Print(kgp.ResetAnimation(20).Encode())
```

### Relative Positioning

```go
// Create parent image
parentCmd := kgp.NewTransmitDisplay().
    ImageID(100).
    PlacementID(1).
    Format(kgp.FormatPNG).
    TransmitDirect(parentData).
    Build()
fmt.Print(parentCmd.Encode())

// Create child image 10 cells to the right
childCmd := kgp.NewPut(100).
    PlacementID(2).
    RelativeTo(100, 1, 10, 0).  // parent image ID, parent placement ID, H offset, V offset
    Build()
fmt.Print(childCmd.Encode())
```

### Chunked Transmission

```go
// For large images or remote connections
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(largePNGData).
    Build()

// Split into 4096-byte chunks
chunks := cmd.EncodeChunked(4096)
for _, chunk := range chunks {
    fmt.Print(chunk)
}
```

### Compression

```go
// Compress RGBA data
rgbaData := kgp.SolidColorImage(800, 600, 255, 0, 0, 255)
compressed, _ := kgp.CompressZlib(rgbaData)

cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    Compress().
    TransmitDirect(compressed).
    Build()
fmt.Print(cmd.Encode())
```

### File Transmission

```go
// Load from file
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitFile("/path/to/image.png").
    Build()
fmt.Print(cmd.Encode())

// Load from file with offset and size
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitFileWithOffset("/path/to/image.png", 1024, 50000).
    Build()
fmt.Print(cmd.Encode())

// Temporary file (terminal deletes after reading)
// Path must contain "tty-graphics-protocol"
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitTemp("/tmp/tty-graphics-protocol-12345.png").
    Build()
fmt.Print(cmd.Encode())
```

### Shared Memory

```go
// POSIX shared memory (most efficient for local apps)
size := 1920000  // 800x600 RGBA
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    TransmitSharedMemory("/my-shm-object", size).
    Build()
fmt.Print(cmd.Encode())
```

### Query Support

```go
// Check if terminal supports graphics protocol
cmd := kgp.QuerySupport()
fmt.Print(cmd.Encode())

// Terminal responds with OK if supported
```

### Response Parsing

```go
// Parse terminal response
resp, err := kgp.ParseResponse("\x1b_Gi=10,p=1;OK\x1b\\")
if err != nil {
    fmt.Println("Error:", err)
    return
}

if resp.Success {
    fmt.Printf("Image ID: %d, Placement ID: %d\n", resp.ImageID, resp.PlacementID)
} else {
    fmt.Printf("Error: %s - %s\n", resp.ErrorCode, resp.Message)
}
```

## API Reference

### Builders

- **`NewTransmit()`** - Transmit image without displaying
- **`NewTransmitDisplay()`** - Transmit and display image
- **`NewPut(imageID)`** - Create new placement of existing image
- **`NewDelete(mode)`** - Delete images or placements
- **`NewFrame(imageID)`** - Add animation frame
- **`NewAnimate(imageID)`** - Control animation playback
- **`NewCompose(imageID)`** - Compose animation frames
- **`NewQuery()`** - Query terminal capabilities

### Helper Functions

- **`TransmitImage(img)`** - Convenience function to transmit image.Image
- **`TransmitImageWithID(img, id)`** - Transmit image.Image with specific ID
- **`TransmitImageRGBA(img, compress)`** - Transmit as raw RGBA
- **`DeleteAll()`** - Delete all placements
- **`DeleteAllFree()`** - Delete all and free memory
- **`DeleteImage(id)`** - Delete specific image
- **`DeleteImageFree(id)`** - Delete and free specific image
- **`DeleteAtCursor()`** - Delete at cursor position
- **`PlayAnimation(id)`** - Play animation
- **`PlayAnimationLoop(id)`** - Play with infinite looping
- **`StopAnimation(id)`** - Stop animation
- **`ResetAnimation(id)`** - Reset to first frame
- **`QuerySupport()`** - Check protocol support

### Image Utilities

- **`ImageToRGBA(img)`** - Convert image.Image to raw RGBA bytes
- **`ImageToRGB(img)`** - Convert image.Image to raw RGB bytes
- **`ImageToPNG(img)`** - Convert image.Image to PNG bytes
- **`CompressZlib(data)`** - Compress data with ZLIB
- **`SolidColorImage(w, h, r, g, b, a)`** - Create solid color image
- **`CreateRGBAColor(r, g, b, a)`** - Create 32-bit RGBA color

### Constants

#### Formats
- `FormatRGB` (24) - 24-bit RGB
- `FormatRGBA` (32) - 32-bit RGBA (default)
- `FormatPNG` (100) - PNG format

#### Transmission Methods
- `TransmitDirect` - Direct embedding (default)
- `TransmitFile` - Load from file
- `TransmitTemp` - Load from temporary file
- `TransmitSharedMem` - Load from shared memory

#### Delete Modes
- `DeleteAllPlacements` / `DeleteAllPlacementsFree`
- `DeleteByImageID` / `DeleteByImageIDFree`
- `DeleteByImageNumber` / `DeleteByImageNumberFree`
- `DeleteByCursor` / `DeleteByCursorFree`
- `DeleteByPlacementID` / `DeleteByPlacementIDFree`
- `DeleteByCell` / `DeleteByCellFree`
- `DeleteByColumn` / `DeleteByColumnFree`
- `DeleteByRow` / `DeleteByRowFree`
- `DeleteByZIndex` / `DeleteByZIndexFree`

(Lowercase = preserve data, Uppercase = free data)

#### Animation States
- `AnimationStop` (1) - Stop at current frame
- `AnimationLoading` (2) - Loading mode (waits for more frames)
- `AnimationLoop` (3) - Normal looping playback

#### Composition Modes
- `CompositionBlend` (0) - Alpha blend (default)
- `CompositionReplace` (1) - Replace pixels

#### Response Suppression
- `ResponseAll` (0) - Send all responses (default)
- `ResponseErrorsOnly` (1) - Suppress OK responses
- `ResponseOKOnly` (2) - Suppress error responses

## Terminal Support

The Kitty Graphics Protocol is supported by:

- **Kitty** (full support, version 0.19.0+)
- **WezTerm** (partial support)
- **Konsole** (experimental)

Check support with:

```go
cmd := kgp.QuerySupport()
fmt.Print(cmd.Encode())
// Terminal responds with OK if supported
```

## Performance Tips

1. **Local applications**: Use shared memory (`TransmitSharedMemory`) for best performance
2. **Remote applications**: Use PNG format to reduce transmission size
3. **Large images**: Enable compression with `Compress()`
4. **Repeated images**: Use placements (`NewPut`) instead of re-transmitting
5. **Animations**: Use gap frames and frame composition efficiently

## Error Handling

Parse terminal responses to handle errors:

```go
resp, err := kgp.ParseResponse(responseStr)
if err != nil {
    // Invalid response format
    return
}

if !resp.Success {
    switch resp.ErrorCode {
    case "ENOSPC":
        // Storage quota exceeded - delete old images
    case "ENOENT":
        // Image not found
    case "EINVAL":
        // Invalid parameters
    case "EIO":
        // I/O error
    default:
        fmt.Printf("Error: %s - %s\n", resp.ErrorCode, resp.Message)
    }
}
```

## License

[MIT License](LICENSE)

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md)

## References

- [Kitty Graphics Protocol Specification](https://sw.kovidgoyal.net/kitty/graphics-protocol/)
- [Kitty Terminal](https://sw.kovidgoyal.net/kitty/)
