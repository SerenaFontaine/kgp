---
title: Response
weight: 7
---

Parsed terminal response to a graphics command.

## Struct

```go
type Response struct {
    ImageID     uint32  // Assigned image ID
    ImageNumber uint32  // Assigned image number (if requested)
    PlacementID uint32  // Assigned placement ID
    Success     bool    // true if OK
    ErrorCode   string  // e.g., "ENOSPC", "ENOENT", "EINVAL"
    Message     string  // Error message (if any)
}
```

## ParseResponse

```go
func ParseResponse(response string) (*Response, error)
```

Parses a terminal response string in permissive mode. This parser accepts responses with or without APC markers and tolerates unknown/malformed control-data pairs when possible.

## ParseResponseStrict

```go
func ParseResponseStrict(response string) (*Response, error)
```

Parses a terminal response string with strict validation. Requires APC markers and a valid status (`OK` or `ERROR_CODE:message`).

Strict format:

```
ESC_Gi=<id>[,I=<num>][,p=<pid>];[OK|ERROR_CODE:message]ESC\
```

**Returns:** `*Response` and `error`. If strict format validation fails, returns a non-nil error. If the format is valid but the status is an error (e.g., `ENOSPC:Storage full`), returns a `*Response` with `Success == false` and `ErrorCode`/`Message` set.

## Error Codes

| Code | Meaning |
|------|---------|
| `ENOSPC` | Storage quota exceeded — delete old images |
| `ENOENT` | Image not found |
| `EINVAL` | Invalid parameters |
| `EIO` | I/O error |
