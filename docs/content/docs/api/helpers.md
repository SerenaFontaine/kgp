---
title: Helper Functions
weight: 9
---

Convenience functions for common operations.

## Transmit Helpers

### TransmitImage

```go
func TransmitImage(img image.Image) (*Command, error)
```

Transmits and displays an `image.Image` as PNG. Uses `image/png` for encoding.

### TransmitImageWithID

```go
func TransmitImageWithID(img image.Image, imageID uint32) (*Command, error)
```

Same as `TransmitImage` but with a specific image ID.

### TransmitImageRGBA

```go
func TransmitImageRGBA(img image.Image, compress bool) (*Command, error)
```

Transmits as raw RGBA. If `compress` is true, applies ZLIB compression.

---

## Delete Helpers

| Function | Description |
|----------|-------------|
| `DeleteAll()` | Delete all placements, preserve data |
| `DeleteAllFree()` | Delete all and free memory |
| `DeleteImage(imageID uint32)` | Delete placements of image, preserve data |
| `DeleteImageFree(imageID uint32)` | Delete image and free memory |
| `DeleteAtCursor()` | Delete at cursor, preserve data |
| `DeleteAtCursorFree()` | Delete at cursor and free |

---

## Animation Helpers

| Function | Description |
|----------|-------------|
| `PlayAnimation(imageID uint32)` | Play using protocol `LoopCount(2)` |
| `PlayAnimationLoop(imageID uint32)` | Play with infinite looping |
| `PlayAnimationWithLoopCount(imageID, count uint32)` | Play with explicit protocol loop-count value |
| `StopAnimation(imageID uint32)` | Stop at current frame |
| `ResetAnimation(imageID uint32)` | Stop and reset to first frame |

---

## Image Utilities

### Conversion

| Function | Description |
|----------|-------------|
| `ImageToRGBA(img image.Image) []byte` | Convert to raw RGBA bytes |
| `ImageToRGB(img image.Image) []byte` | Convert to raw RGB bytes |
| `ImageToPNG(img image.Image) ([]byte, error)` | Encode as PNG |

### Compression

```go
func CompressZlib(data []byte) ([]byte, error)
```

Compresses data with ZLIB (RFC 1950). Use with `TransmitDirect()` and `Compress()`.

### Color Utilities

```go
func CreateRGBAColor(r, g, b, a uint8) uint32
```

Creates a 32-bit RGBA value for `BackgroundColor` / `BackgroundFrame` composition.

```go
func SolidColorImage(width, height int, r, g, b, a uint8) []byte
```

Creates raw RGBA bytes for a solid color image. Useful for demos and placeholders.
