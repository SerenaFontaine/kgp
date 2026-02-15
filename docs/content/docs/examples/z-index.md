---
title: Z-Index Layering
weight: 4
---

Control whether images appear above or below text.

- **Negative z-index**: below text (background)
- **Positive z-index**: above text (overlay)
- **Zero**: default (implementation-defined)

## Background Image

```go
bgCmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(bgData).
    ZIndex(-10).
    CursorMovement(false).
    Build()
fmt.Print(bgCmd.Encode())
fmt.Println("Text rendered on top of the image")
```

## Foreground Overlay

```go
fgCmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(iconData).
    ZIndex(10).
    Build()
fmt.Print(fgCmd.Encode())
```

## Layered Composition

```go
// Back layer
fmt.Print(kgp.NewTransmitDisplay().
    ImageID(1).
    Format(kgp.FormatPNG).
    TransmitDirect(bgData).
    ZIndex(-2).
    CursorMovement(false).
    Build().Encode())

// Middle layer
fmt.Print(kgp.NewTransmitDisplay().
    ImageID(2).
    Format(kgp.FormatPNG).
    TransmitDirect(midData).
    ZIndex(-1).
    CursorMovement(false).
    Build().Encode())

// Front layer
fmt.Print(kgp.NewTransmitDisplay().
    ImageID(3).
    Format(kgp.FormatPNG).
    TransmitDirect(fgData).
    ZIndex(1).
    Build().Encode())
```
