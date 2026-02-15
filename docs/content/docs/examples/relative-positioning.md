---
title: Relative Positioning
weight: 7
---

Place child images relative to a parent placement. Useful for overlays, badges, and anchored UI elements.

## Parent and Child

```go
// Parent image
parentCmd := kgp.NewTransmitDisplay().
    ImageID(100).
    PlacementID(1).
    Format(kgp.FormatPNG).
    TransmitDirect(parentData).
    Build()
fmt.Print(parentCmd.Encode())

// Child: 10 cells right, 0 cells down from parent
childCmd := kgp.NewPut(100).
    PlacementID(2).
    RelativeTo(100, 1, 10, 0).
    Build()
fmt.Print(childCmd.Encode())
```

## RelativeTo Parameters

```go
RelativeTo(parentImageID, parentPlacementID uint32, offsetH, offsetV int)
```

- `parentImageID` — Image ID of the parent
- `parentPlacementID` — Placement ID of the parent
- `offsetH` — Horizontal offset in cells (positive = right)
- `offsetV` — Vertical offset in cells (positive = down)

## Badge on Avatar

```go
// Avatar (parent)
avatarCmd := kgp.NewTransmitDisplay().
    ImageID(50).
    PlacementID(1).
    Format(kgp.FormatPNG).
    TransmitDirect(avatarData).
    Build()
fmt.Print(avatarCmd.Encode())

// Transmit badge (no display)
fmt.Print(kgp.NewTransmit().
    ImageID(51).
    Format(kgp.FormatPNG).
    TransmitDirect(badgeData).
    Build().Encode())

// Badge (child) — top-right of avatar
badgeCmd := kgp.NewPut(51).
    PlacementID(1).
    DisplaySize(2, 2).
    RelativeTo(50, 1, 8, -1).
    Build()
fmt.Print(badgeCmd.Encode())
```
