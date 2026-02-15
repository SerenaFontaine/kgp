---
title: Error Handling
weight: 11
---

Parse terminal responses to detect and handle errors.

## Basic Parsing

```go
resp, err := kgp.ParseResponse(responseStr)
if err != nil {
    fmt.Println("Invalid response format:", err)
    return
}

if resp.Success {
    fmt.Printf("Image ID: %d, Placement ID: %d\n", resp.ImageID, resp.PlacementID)
} else {
    fmt.Printf("Error: %s - %s\n", resp.ErrorCode, resp.Message)
}
```

## Error Code Handling

```go
resp, err := kgp.ParseResponse(responseStr)
if err != nil {
    return err
}

if !resp.Success {
    switch resp.ErrorCode {
    case "ENOSPC":
        // Storage quota exceeded — delete old images
        fmt.Print(kgp.DeleteAllFree().Encode())
        // Retry...
    case "ENOENT":
        // Image not found — may have been deleted
    case "EINVAL":
        // Invalid parameters — check your command
    case "EIO":
        // I/O error — terminal problem
    default:
        fmt.Printf("Unknown error: %s - %s\n", resp.ErrorCode, resp.Message)
    }
}
```

## Suppressing OK Responses

Reduce response traffic when you don't need confirmation:

```go
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitDirect(data).
    ResponseSuppression(kgp.ResponseErrorsOnly).
    Build()
```

- `ResponseErrorsOnly` — Terminal sends only errors, not OK
- `ResponseOKOnly` — Terminal sends only OK, not errors (rare)
- `ResponseAll` — Default; both OK and errors
