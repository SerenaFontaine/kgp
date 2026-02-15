---
title: Demo Application
weight: 12
---

The `examples/demo` package provides an interactive walkthrough of KGP features.

## Running the Demo

```bash
cd examples/demo
go run .
```

Place a file named `sf.png` in the demo directory for the external PNG loading demo.

## Demo Sequence

1. **Image Formats** — PNG, RGBA, RGB
2. **Compression** — ZLIB compressed RGBA
3. **External PNG** — Load from `sf.png`
4. **Cropping and Scaling** — Source rect and display size
5. **Multiple Placements** — One image, multiple displays
6. **Placement Sizes** — Same image at different cell sizes
7. **Z-Index** — Background image behind text
8. **Query Support** — Check terminal capabilities
9. **Animation** — Color-changing square with movement

## Code Structure

```go
// Clear screen
fmt.Print("\x1b[2J\x1b[H")

// Create image data
redSquare := createColoredSquare(100, 100, color.RGBA{R: 255, G: 0, B: 0, A: 255})
pngData, _ := kgp.ImageToPNG(redSquare)

// Transmit and display
cmd := kgp.NewTransmitDisplay().
    ImageID(1).
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    ResponseSuppression(kgp.ResponseErrorsOnly).
    Build()
fmt.Print(cmd.Encode())
```

## Helper Functions in Demo

- `createColoredSquare(width, height, color)` — Solid color image
- `createGradientSquare(width, height)` — Gradient for cropping demo
- `clearScreen()` — ANSI clear and home
- `waitForEnter(reader)` — Pause for user input
