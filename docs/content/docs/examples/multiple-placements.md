---
title: Multiple Placements
weight: 2
---

Transmit once, display many times. Saves bandwidth and memory.

## Transmit Then Put

```go
// 1. Transmit image (no display)
transmitCmd := kgp.NewTransmit().
    ImageID(10).
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    Build()
fmt.Print(transmitCmd.Encode())

// 2. Create placements
placement1 := kgp.NewPut(10).
    DisplaySize(20, 15).
    ZIndex(1).
    Build()
fmt.Print(placement1.Encode())

placement2 := kgp.NewPut(10).
    DisplaySize(10, 7).
    CellOffset(5, 5).
    Build()
fmt.Print(placement2.Encode())
```

## Same Image, Different Sizes

```go
transmitCmd := kgp.NewTransmit().
    ImageID(100).
    Format(kgp.FormatPNG).
    TransmitDirect(pngData).
    Build()
fmt.Print(transmitCmd.Encode())

// Small
fmt.Print(kgp.NewPut(100).DisplaySize(5, 5).Build().Encode())
fmt.Print("  ")

// Medium
fmt.Print(kgp.NewPut(100).DisplaySize(8, 8).Build().Encode())
fmt.Print("  ")

// Large
fmt.Print(kgp.NewPut(100).DisplaySize(12, 12).Build().Encode())
fmt.Println()
```

## Placement IDs for Targeting

Use `PlacementID` when you need to delete or update specific placements.

```go
place1 := kgp.NewPut(100).
    PlacementID(1).
    DisplaySize(10, 10).
    Build()
fmt.Print(place1.Encode())

place2 := kgp.NewPut(100).
    PlacementID(2).
    DisplaySize(5, 5).
    Build()
fmt.Print(place2.Encode())

// Later: delete only placement 2
del := kgp.NewDelete(kgp.DeleteByPlacementID).
    ImageID(100).
    PlacementID(2).
    Build()
fmt.Print(del.Encode())
```
