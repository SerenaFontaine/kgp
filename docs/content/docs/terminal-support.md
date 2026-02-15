---
title: Terminal Support
weight: 50
---

The Kitty Graphics Protocol is supported by:

| Terminal | Support | Notes |
|----------|---------|-------|
| **Kitty** | Full | Version 0.19.0+ |
| **WezTerm** | Partial | Core features supported |
| **Konsole** | Experimental | KDE Konsole |

## Verify Support

Before rendering graphics, check support:

```go
cmd := kgp.QuerySupport()
fmt.Print(cmd.Encode())
// Terminal responds with OK if supported
```

## Performance Tips

1. **Local applications**: Use `TransmitSharedMemory` for best performance
2. **Remote applications**: Use PNG format to reduce transmission size
3. **Large images**: Enable compression with `Compress()`
4. **Repeated images**: Use placements (`NewPut`) instead of re-transmitting
5. **Animations**: Use gap frames and frame composition efficiently

## References

- [Kitty Graphics Protocol Specification](https://sw.kovidgoyal.net/kitty/graphics-protocol/)
- [Kitty Terminal](https://sw.kovidgoyal.net/kitty/)
