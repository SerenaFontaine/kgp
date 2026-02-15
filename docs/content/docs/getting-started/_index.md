---
title: Getting Started
weight: 10
---

## Installation

Add KGP to your Go module:

```bash
go get github.com/SerenaFontaine/kgp
```

## Requirements

- Go 1.23 or later
- A terminal that supports the Kitty Graphics Protocol (Kitty, WezTerm, or Konsole)

## Quick Start

### Display a PNG Image

```go
package main

import (
    "fmt"
    "os"
    "github.com/SerenaFontaine/kgp"
)

func main() {
    data, _ := os.ReadFile("image.png")

    cmd := kgp.NewTransmitDisplay().
        Format(kgp.FormatPNG).
        TransmitDirect(data).
        Build()

    fmt.Print(cmd.Encode())
}
```

### Display from image.Image

```go
import (
    "fmt"
    "image"
    "github.com/SerenaFontaine/kgp"
)

func displayImage(img image.Image) {
    cmd, _ := kgp.TransmitImage(img)
    fmt.Print(cmd.Encode())
}
```

## Verify Terminal Support

Before rendering graphics, check if the terminal supports the protocol:

```go
cmd := kgp.QuerySupport()
fmt.Print(cmd.Encode())
// Terminal responds with OK if supported
```
