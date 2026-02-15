---
title: Animation Builders
weight: 5
---

Animation support uses three builders: `FrameBuilder` (add frames), `AnimateBuilder` (control playback), and `ComposeBuilder` (compose frames).

## FrameBuilder

Adds animation frames to an existing image.

### NewFrame

```go
func NewFrame(imageID uint32) *FrameBuilder
```

### Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `FrameData` | `(data []byte)` | Raw frame image data |
| `Format` | `(format Format)` | Format for frame (required for RGB/RGBA) |
| `Dimensions` | `(width, height int)` | Required for RGB/RGBA |
| `FrameNumber` | `(frameNum uint32)` | Edit existing frame (replace instead of append) |
| `BackgroundFrame` | `(frameNum uint32)` | Frame to use as background (0 = base image) |
| `Gap` | `(milliseconds uint32)` | Delay before next frame in ms |
| `Composition` | `(mode CompositionMode)` | `CompositionBlend` or `CompositionReplace` |
| `BackgroundColor` | `(rgba uint32)` | 32-bit RGBA background color |
| `ResponseSuppression` | `(mode ResponseSuppression)` | Control responses |

---

## AnimateBuilder

Controls animation playback.

### NewAnimate

```go
func NewAnimate(imageID uint32) *AnimateBuilder
```

### Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `State` | `(state AnimationState)` | `AnimationStop`, `AnimationLoading`, `AnimationLoop` |
| `LoopCount` | `(count uint32)` | 0 = ignored, 1 = infinite, N>1 = loop N-1 times |
| `GapOverride` | `(milliseconds uint32)` | Override all frame gaps |
| `Frame` | `(frameNum uint32)` | Frame to stop at (with `AnimationStop`) |
| `ResponseSuppression` | `(mode ResponseSuppression)` | Control responses |

---

## ComposeBuilder

Composes rectangular regions between animation frames without adding new frame data.

### NewCompose

```go
func NewCompose(imageID uint32) *ComposeBuilder
```

### Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `SourceFrame` | `(frameNum uint32)` | Source frame to compose from |
| `DestFrame` | `(frameNum uint32)` | Destination frame to compose onto |
| `SourceRect` | `(x, y, width, height int)` | Source rectangle offset and size |
| `DestOffset` | `(x, y int)` | Destination offset for composed rectangle |
| `Composition` | `(mode CompositionMode)` | `CompositionBlend` or `CompositionReplace` |
| `ResponseSuppression` | `(mode ResponseSuppression)` | Control responses |
