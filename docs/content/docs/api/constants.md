---
title: Constants
weight: 8
---

## Formats

| Constant | Value | Description |
|----------|-------|-------------|
| `FormatRGB` | 24 | 24-bit RGB (3 bytes per pixel) |
| `FormatRGBA` | 32 | 32-bit RGBA (4 bytes per pixel, default) |
| `FormatPNG` | 100 | PNG with embedded dimensions |

## Transmit Mediums

| Constant | Value | Description |
|----------|-------|-------------|
| `TransmitDirect` | `"d"` | Embed data in escape sequence (default) |
| `TransmitFile` | `"f"` | Read from regular file |
| `TransmitTemp` | `"t"` | Temporary file (terminal deletes after read) |
| `TransmitSharedMem` | `"s"` | POSIX shared memory |

## Compression

| Constant | Value |
|----------|-------|
| `CompressionZlib` | `"z"` |

## Delete Modes

**Preserve data** (lowercase): Delete placement(s) but keep image data in memory.

**Free data** (uppercase): Delete and free memory.

| Preserve | Free | Description |
|----------|------|-------------|
| `DeleteAllPlacements` | `DeleteAllPlacementsFree` | All placements / all + free |
| `DeleteByImageID` | `DeleteByImageIDFree` | By image ID |
| `DeleteByImageNumber` | `DeleteByImageNumberFree` | By image number |
| `DeleteByCursor` | `DeleteByCursorFree` | At cursor position |
| `DeleteByPlacementID` | `DeleteByPlacementIDFree` | By placement ID |
| `DeleteByCell` | `DeleteByCellFree` | By cell (x, y) |
| `DeleteByColumn` | `DeleteByColumnFree` | By column |
| `DeleteByRow` | `DeleteByRowFree` | By row |
| `DeleteByZIndex` | `DeleteByZIndexFree` | By z-index |
| `DeleteByIDRange` | `DeleteByIDRangeFree` | By image ID range |
| `DeleteFrames` | `DeleteFramesFree` | Animation frames |

## Animation States

| Constant | Value | Description |
|----------|-------|-------------|
| `AnimationStop` | 1 | Stop at current frame |
| `AnimationLoading` | 2 | Loading mode (waits for more frames) |
| `AnimationLoop` | 3 | Looping playback |

## Composition Modes

| Constant | Value | Description |
|----------|-------|-------------|
| `CompositionBlend` | 0 | Alpha blending (default) |
| `CompositionReplace` | 1 | Replace pixels, no blend |

## Response Suppression

| Constant | Value | Description |
|----------|-------|-------------|
| `ResponseAll` | 0 | Send all responses (default) |
| `ResponseErrorsOnly` | 1 | Suppress OK responses |
| `ResponseOKOnly` | 2 | Suppress error responses |

## Actions (Internal)

| Constant | Value |
|----------|-------|
| `ActionTransmit` | `"t"` |
| `ActionTransmitDisplay` | `"T"` |
| `ActionPut` | `"p"` |
| `ActionDelete` | `"d"` |
| `ActionFrame` | `"f"` |
| `ActionAnimate` | `"a"` |
| `ActionCompose` | `"c"` |
| `ActionQuery` | `"q"` |
