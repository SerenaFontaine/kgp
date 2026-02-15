---
title: PutBuilder
weight: 3
---

Creates placement commands for existing images. Use after `NewTransmit()` to display an image in multiple locations or with different display parameters.

## Constructor

### NewPut

```go
func NewPut(imageID uint32) *PutBuilder
```

Creates a put builder for the specified image. The image must already exist (transmitted previously).

## Methods

### Image & Placement Identification

| Method | Signature | Description |
|--------|-----------|-------------|
| `ImageNumber` | `(num uint32)` | Application-specific image number (non-unique) |
| `PlacementID` | `(id uint32)` | Set placement ID (auto-generated if omitted) |

### Display Options

| Method | Signature | Description |
|--------|-----------|-------------|
| `CellOffset` | `(x, y int)` | Pixel offset within starting cell |
| `DisplaySize` | `(columns, rows int)` | Display size in terminal cells |
| `SourceRect` | `(x, y, width, height int)` | Display only a region of the source image |
| `ZIndex` | `(z int)` | Z-index (negative=below text, positive=above) |
| `CursorMovement` | `(move bool)` | `true`=cursor advances, `false`=cursor stays |
| `RelativeTo` | `(parentImageID, parentPlacementID uint32, offsetH, offsetV int)` | Position relative to parent placement |
| `VirtualPlacement` | `()` | Invisible placement for Unicode placeholders |

### Response

| Method | Signature | Description |
|--------|-----------|-------------|
| `ResponseSuppression` | `(mode ResponseSuppression)` | Control which responses are sent |

## Build

```go
func (pb *PutBuilder) Build() *Command
```
