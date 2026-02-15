---
title: QueryBuilder
weight: 6
---

Queries terminal capabilities. Send a minimal graphics command; the terminal responds with `OK` if supported.

## Constructor

### NewQuery

```go
func NewQuery() *QueryBuilder
```

## Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `Format` | `(format Format)` | Format to test |
| `Dimensions` | `(width, height int)` | Test dimensions |
| `TransmitMedium` | `(medium TransmitMedium)` | Medium to test |
| `TestData` | `(data []byte)` | Minimal payload (e.g., 3 bytes for RGB) |

## Build

```go
func (qb *QueryBuilder) Build() *Command
```

## QuerySupport

Convenience function for the most common query:

```go
func QuerySupport() *Command
```

Returns a minimal query (1×1 RGB pixel) that tests basic protocol support. The terminal responds with `OK` if the Kitty Graphics Protocol is supported.
