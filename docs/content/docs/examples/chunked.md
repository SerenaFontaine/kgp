---
title: Chunked Transmission
weight: 8
---

For large images or remote connections (e.g., SSH), split the payload into chunks. The terminal reassembles them.

## Basic Chunking

```go
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(largePNGData).
    Build()

chunks := cmd.EncodeChunked(4096)
for _, chunk := range chunks {
    fmt.Print(chunk)
}
```

## Chunk Size Requirements

- **Maximum**: 4096 bytes per chunk
- **Alignment**: Must be divisible by 4 (base64 encoding)

```go
// Valid
cmd.EncodeChunked(4096)
cmd.EncodeChunked(2048)
cmd.EncodeChunked(1024)

// Invalid — will panic
cmd.EncodeChunked(4097)  // > 4096
cmd.EncodeChunked(1002) // not divisible by 4
```

## When to Use Chunking

- Large PNG or RGBA payloads that exceed terminal limits
- Slow or unreliable connections (SSH, serial)
- Some terminal implementations enforce payload size limits
