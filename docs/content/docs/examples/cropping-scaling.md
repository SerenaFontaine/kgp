---
title: Cropping and Scaling
weight: 3
---

## Source Rect (Crop)

Display only a portion of the source image.

```go
// Display top-left 400×300 region
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    SourceRect(0, 0, 400, 300).
    Build()
fmt.Print(cmd.Encode())
```

## Display Size (Scale)

Control how many terminal cells the image occupies.

```go
// Crop center 100×100, display in 10×10 cells
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(largePNG).
    SourceRect(50, 50, 100, 100).
    DisplaySize(10, 10).
    Build()
fmt.Print(cmd.Encode())
```

## Put with Crop and Scale

```go
// Already transmitted image ID 20
place := kgp.NewPut(20).
    SourceRect(100, 100, 200, 200).  // Crop
    DisplaySize(15, 15).              // Scale
    Build()
fmt.Print(place.Encode())
```

## Cell Offset

Position the image within the starting cell.

```go
place := kgp.NewPut(10).
    DisplaySize(10, 10).
    CellOffset(8, 4).  // Pixel offset in first cell
    Build()
fmt.Print(place.Encode())
```
