---
title: KGP — Kitty Graphics Protocol Go Bindings
---

Go bindings for the [Kitty Graphics Protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/), enabling terminal applications to render pixel-based graphics alongside text.

## Features

- **Type-safe API** with fluent builders for all actions
- **Complete protocol coverage** including:
  - Image transmission (direct, file, temporary file, shared memory)
  - Multiple placements of images
  - Animation support (frames, playback control, composition)
  - Z-index layering (above/below text)
  - Relative positioning (parent-child relationships)
  - Cropping and scaling
  - Response parsing
- **Chunked transmission** for large images
- **Helper functions** for common operations
- **Image format support**: PNG, RGB, RGBA
- **Compression support**: ZLIB
- **Zero dependencies** (beyond Go standard library)

## Quick Links

- [Getting Started](/docs/getting-started/) — Installation and first steps
- [Examples](/docs/examples/) — Practical usage examples
- [API Reference](/docs/api/) — Exhaustive technical reference
- [Protocol](/docs/protocol/) — Kitty Graphics Protocol details
