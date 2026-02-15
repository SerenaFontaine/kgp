---
title: Kitty Graphics Protocol
weight: 20
---

KGP implements the [Kitty Graphics Protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/), a terminal escape sequence specification for rendering pixel-based graphics.

## Escape Sequence Format

Commands use the APC (Application Program Command) sequence:

```
ESC_G<control-data>;<payload>ESC\
```

- **Start**: `\x1b_G` (ESC followed by `G`)
- **Control data**: Comma-separated `key=value` pairs
- **Payload**: Base64-encoded binary data (when applicable)
- **End**: `\x1b\` (ESC followed by backslash)

## Actions

| Action | Code | Description |
|--------|------|-------------|
| Transmit | `t` | Upload image data without displaying |
| Transmit & Display | `T` | Upload and display image |
| Put | `p` | Create placement of existing image |
| Delete | `d` | Remove images or placements |
| Frame | `f` | Add animation frame |
| Animate | `a` | Control animation playback |
| Compose | `c` | Compose animation frames |
| Query | `q` | Query terminal capabilities |

## Control Data Keys

### Common Keys

| Key | Description | Used By |
|-----|-------------|---------|
| `a` | Action type | All |
| `i` | Image ID | Transmit, Put, Delete, Frame, Animate, Compose |
| `I` | Image number (non-unique) | Transmit, Put, Delete |
| `p` | Placement ID | Transmit, Put, Delete |
| `f` | Format (24, 32, 100) | Transmit, Frame |
| `s`, `v` | Width, height (dimensions) | Transmit, Frame, Query |
| `t` | Transmit medium (d, f, t, s) | Transmit, Query |
| `o` | Compression (z = zlib) | Transmit |
| `q` | Response suppression (0, 1, 2) | All |
| `d` | Delete mode | Delete |

### Placement Keys

| Key | Description |
|-----|-------------|
| `c`, `r` | Display size (columns, rows) |
| `X`, `Y` | Cell offset within starting cell (pixels) |
| `x`, `y`, `w`, `h` | Source rectangle (crop) |
| `z` | Z-index (negative = below text) |
| `C` | Cursor movement (0=move, 1=no move) |
| `P`, `Q` | Parent image/placement for relative positioning |
| `H`, `V` | Horizontal/vertical offset for relative positioning |
| `U` | Virtual placement (Unicode placeholder) |

### Frame Keys (action=f)

| Key | Description |
|-----|-------------|
| `r` | Frame number to edit (omit to create new frame) |
| `z` | Frame gap (ms) |
| `c` | Background frame number |
| `X` | Composition mode (0=blend, 1=replace) |
| `Y` | Background color (32-bit RGBA) |

### Animate Keys (action=a)

| Key | Description |
|-----|-------------|
| `s` | Animation state (1=stop, 2=loading, 3=loop) |
| `v` | Loop count (0=ignored, 1=infinite, N>1=loop N-1 times) |
| `z` | Gap override (ms) |
| `c` | Frame number (used with stop state) |

### Compose Keys (action=c)

| Key | Description |
|-----|-------------|
| `r` | Source frame number |
| `c` | Destination frame number |
| `x`, `y`, `w`, `h` | Source rectangle |
| `X`, `Y` | Destination offset (pixels) |
| `C` | Composition mode (0=blend, 1=replace) |

## Response Format

Terminal responses follow:

```
ESC_Gi=<id>[,I=<num>][,p=<pid>];[OK|ERROR_CODE:message]ESC\
```

Parse with `kgp.ParseResponse()`.
