---
title: Basic Image Display
weight: 1
---

## PNG from File

```go
package main

import (
    "fmt"
    "os"
    "github.com/SerenaFontaine/kgp"
)

func main() {
    data, err := os.ReadFile("image.png")
    if err != nil {
        panic(err)
    }

    cmd := kgp.NewTransmitDisplay().
        Format(kgp.FormatPNG).
        TransmitDirect(data).
        Build()

    fmt.Print(cmd.Encode())
}
```

## RGBA with Dimensions

PNG embeds dimensions; RGB and RGBA require explicit dimensions.

```go
rgbaData := kgp.SolidColorImage(800, 600, 255, 0, 0, 255) // Red

cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    TransmitDirect(rgbaData).
    Build()
fmt.Print(cmd.Encode())
```

## RGB (24-bit, no alpha)

```go
import "image"

img := loadYourImage()
rgbData := kgp.ImageToRGB(img)
width, height := img.Bounds().Dx(), img.Bounds().Dy()

cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGB).
    Dimensions(width, height).
    TransmitDirect(rgbData).
    Build()
fmt.Print(cmd.Encode())
```

## From image.Image (Convenience)

```go
img, _ := image.Decode(file)
cmd, err := kgp.TransmitImage(img)
if err != nil {
    panic(err)
}
fmt.Print(cmd.Encode())
```
