---
title: DeleteBuilder
weight: 4
---

Builds delete commands. Delete modes determine what is removed and whether image data is freed.

## Constructor

### NewDelete

```go
func NewDelete(mode DeleteMode) *DeleteBuilder
```

Creates a delete builder. The mode must be one of the `DeleteBy*` constants.

## Methods

Mode-specific parameters (call only the ones relevant to the chosen mode):

| Method | Signature | Used With Modes |
|--------|-----------|-----------------|
| `ImageID` | `(id uint32)` | `DeleteByImageID`, `DeleteByImageIDFree` |
| `ImageNumber` | `(num uint32)` | `DeleteByImageNumber`, `DeleteByImageNumberFree` |
| `PlacementID` | `(id uint32)` | `DeleteByPlacementID`, `DeleteByPlacementIDFree` |
| `Cell` | `(x, y int)` | `DeleteByCell`, `DeleteByCellFree` |
| `Column` | `(x int)` | `DeleteByColumn`, `DeleteByColumnFree` |
| `Row` | `(y int)` | `DeleteByRow`, `DeleteByRowFree` |
| `ZIndex` | `(z int)` | `DeleteByZIndex`, `DeleteByZIndexFree` |
| `IDRange` | `(startID, endID int)` | `DeleteByIDRange`, `DeleteByIDRangeFree` |

| Method | Signature | Description |
|--------|-----------|-------------|
| `ResponseSuppression` | `(mode ResponseSuppression)` | Control responses |

**Note:** For `DeleteByImageID`, `DeleteByImageIDFree`, `DeleteByPlacementID`, and `DeleteByPlacementIDFree`, you must also set `ImageID` via `ImageID()` when the mode requires it. The protocol uses `i` for both image ID and (in some modes) other purposes; the builder maps these correctly.

## Build

```go
func (db *DeleteBuilder) Build() *Command
```
