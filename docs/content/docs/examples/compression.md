---
title: Compression
weight: 9
---

ZLIB compression reduces payload size for raw RGB/RGBA data. PNG is already compressed.

## Compress RGBA

```go
rgbaData := kgp.SolidColorImage(800, 600, 255, 0, 0, 255)
compressed, err := kgp.CompressZlib(rgbaData)
if err != nil {
    panic(err)
}

cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    Compress().
    TransmitDirect(compressed).
    Build()
fmt.Print(cmd.Encode())
```

## TransmitImageRGBA Helper

```go
// With compression
cmd, err := kgp.TransmitImageRGBA(img, true)
if err != nil {
    panic(err)
}
fmt.Print(cmd.Encode())

// Without compression
cmd, err = kgp.TransmitImageRGBA(img, false)
```

## When to Compress

- **Raw RGBA/RGB**: Often benefits from compression, especially for gradients or similar colors
- **PNG**: Already compressed; no need for `Compress()`
- **Solid colors**: Compression may not help much; test for your use case
