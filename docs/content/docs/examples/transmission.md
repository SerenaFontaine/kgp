---
title: File and Shared Memory Transmission
weight: 10
---

Alternatives to embedding data directly in escape sequences.

## Transmit from File

```go
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitFile("/path/to/image.png").
    Build()
fmt.Print(cmd.Encode())
```

## File with Offset and Size

Read a byte range from the file:

```go
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitFileWithOffset("/path/to/image.png", 1024, 50000).
    Build()
fmt.Print(cmd.Encode())
```

## Temporary File

The terminal deletes the file after reading. **Path must contain `tty-graphics-protocol`** for security.

```go
path := "/tmp/tty-graphics-protocol-" + strconv.Itoa(os.Getpid()) + ".png"
_ = os.WriteFile(path, pngData, 0644)
defer os.Remove(path)

cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG).
    TransmitTemp(path). // panics if path is invalid
    Build()
fmt.Print(cmd.Encode())
```

## Temporary File (Non-Panicking)

```go
path := "/tmp/tty-graphics-protocol-" + strconv.Itoa(os.Getpid()) + ".png"
_ = os.WriteFile(path, pngData, 0644)
defer os.Remove(path)

tb := kgp.NewTransmitDisplay().
    Format(kgp.FormatPNG)

if _, err := tb.TryTransmitTemp(path); err != nil {
    log.Fatal(err)
}

cmd := tb.
    Build()
fmt.Print(cmd.Encode())
```

## Shared Memory (POSIX)

Most efficient for local applications. Create a shared memory object and pass its name and size.

```go
size := 800 * 600 * 4  // RGBA
cmd := kgp.NewTransmitDisplay().
    Format(kgp.FormatRGBA).
    Dimensions(800, 600).
    TransmitSharedMemory("/my-shm-object", size).
    Build()
fmt.Print(cmd.Encode())
```

**Note:** Your application must create and populate the shared memory object before sending the command. The terminal reads from it directly.
