---
title: Command
weight: 1
---

`Command` represents a complete Kitty Graphics Protocol command.

## Creating Commands

Commands are created via builders. Do not instantiate `Command` directly.

```go
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(data).
    Build()  // Returns *Command
```

## Methods

### Encode

```go
func (c *Command) Encode() string
```

Returns the complete escape sequence as a string, including the APC delimiters `ESC_G` and `ESC\`. Payload data is base64-encoded.

### EncodeChunked

```go
func (c *Command) EncodeChunked(maxChunkSize int) []string
```

Splits the payload into multiple escape sequences for chunked transmission. Required when payload exceeds terminal limits (e.g., over SSH).

**Parameters:**
- `maxChunkSize` — Must be ≤ 4096 and divisible by 4 (base64 alignment)

**Returns:** Slice of strings, each a complete escape sequence. First chunk contains all control data; subsequent chunks contain only the `m` key (1=more, 0=last).

## Low-Level API

For custom commands, use `NewCommand` and the setters:

```go
cmd := kgp.NewCommand(kgp.ActionTransmit)
cmd.SetKey("f", "100")           // string value
cmd.SetKeyInt("s", 800)          // integer
cmd.SetKeyUint32("i", 10)        // uint32
cmd.SetPayload(pngData)           // binary payload
```
