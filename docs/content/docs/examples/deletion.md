---
title: Deletion
weight: 5
---

## Delete All Placements

Keep image data in memory for reuse:

```go
fmt.Print(kgp.DeleteAll().Encode())
```

Free everything:

```go
fmt.Print(kgp.DeleteAllFree().Encode())
```

## Delete by Image ID

```go
fmt.Print(kgp.DeleteImage(10).Encode())      // Preserve data
fmt.Print(kgp.DeleteImageFree(10).Encode())  // Free memory
```

## Delete at Cursor

```go
fmt.Print(kgp.DeleteAtCursor().Encode())
```

## Delete by Z-Index

```go
cmd := kgp.NewDelete(kgp.DeleteByZIndex).
    ZIndex(5).
    Build()
fmt.Print(cmd.Encode())
```

## Delete by Placement ID

```go
cmd := kgp.NewDelete(kgp.DeleteByPlacementID).
    ImageID(100).
    PlacementID(2).
    Build()
fmt.Print(cmd.Encode())
```

## Delete by Cell Coordinates

```go
cmd := kgp.NewDelete(kgp.DeleteByCell).
    Cell(10, 5).
    Build()
fmt.Print(cmd.Encode())
```

## Delete by Image ID Range

```go
cmd := kgp.NewDelete(kgp.DeleteByIDRange).
    IDRange(10, 50).
    Build()
fmt.Print(cmd.Encode())
```

Free memory for a range of image IDs:

```go
cmd := kgp.NewDelete(kgp.DeleteByIDRangeFree).
    IDRange(10, 50).
    Build()
fmt.Print(cmd.Encode())
```

## Delete Animation Frames

```go
cmd := kgp.NewDelete(kgp.DeleteFrames).
    ImageID(20).
    Build()
fmt.Print(cmd.Encode())
```
