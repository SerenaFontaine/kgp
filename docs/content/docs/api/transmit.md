---
title: TransmitBuilder
weight: 2
---

Builds transmit or transmit+display commands. Transmit uploads image data; transmit+display also creates the initial placement.

## Constructors

### NewTransmit

```go
func NewTransmit() *TransmitBuilder
```

Creates a transmit-only builder (action `t`). Image data is uploaded but not displayed. Use with `NewPut()` for placements.

### NewTransmitDisplay

```go
func NewTransmitDisplay() *TransmitBuilder
```

Creates a transmit-and-display builder (action `T`). Uploads and displays in one command.

## Methods

### Image Identification

| Method | Signature | Description |
|--------|-----------|-------------|
| `ImageID` | `(id uint32)` | Set image ID (auto-generated if omitted) |
| `ImageNumber` | `(num uint32)` | Set image number (application-specific, non-unique) |

### Format & Dimensions

| Method | Signature | Description |
|--------|-----------|-------------|
| `Format` | `(format Format)` | Set format: `FormatRGB`, `FormatRGBA`, `FormatPNG` |
| `Dimensions` | `(width, height int)` | **Required** for RGB/RGBA formats |

### Transmission

| Method | Signature | Description |
|--------|-----------|-------------|
| `TransmitDirect` | `(data []byte)` | Embed data in escape sequence (default) |
| `TransmitFile` | `(path string)` | Read from file path |
| `TransmitFileWithOffset` | `(path string, offset, size int)` | Read from file with byte range |
| `ValidateTempPath` | `(path string) error` | Validate temporary-file path requirement |
| `TryTransmitTemp` | `(path string) (*TransmitBuilder, error)` | Temporary file with error return on invalid path |
| `TransmitTemp` | `(path string)` | Temporary file; panics if path is invalid |
| `TransmitSharedMemory` | `(name string, size int)` | POSIX shared memory object |
| `Compress` | `()` | Enable ZLIB compression (use with pre-compressed data) |

### Placement Options

| Method | Signature | Description |
|--------|-----------|-------------|
| `PlacementID` | `(id uint32)` | Set placement ID for initial placement |
| `CellOffset` | `(x, y int)` | Pixel offset within starting cell |
| `DisplaySize` | `(columns, rows int)` | Display size in terminal cells |
| `SourceRect` | `(x, y, width, height int)` | Crop source to rectangle |
| `ZIndex` | `(z int)` | Z-index (negative=below text, positive=above) |
| `CursorMovement` | `(move bool)` | `true`=cursor moves after, `false`=stays |
| `RelativeTo` | `(parentImageID, parentPlacementID uint32, offsetH, offsetV int)` | Relative positioning |
| `VirtualPlacement` | `()` | Invisible placement for Unicode placeholders |

### Response

| Method | Signature | Description |
|--------|-----------|-------------|
| `ResponseSuppression` | `(mode ResponseSuppression)` | `ResponseAll`, `ResponseErrorsOnly`, `ResponseOKOnly` |

## Build

```go
func (tb *TransmitBuilder) Build() *Command
```

Constructs the final `*Command`. Call exactly once per builder.
