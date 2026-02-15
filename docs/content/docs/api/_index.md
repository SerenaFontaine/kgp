---
title: API Reference
weight: 30
---

Exhaustive reference for all exported types, functions, and constants in the `kgp` package.

## Builders

Builders use a fluent interface. Call `Build()` to produce a `*Command`.

- [TransmitBuilder](/docs/api/transmit/) — `NewTransmit()`, `NewTransmitDisplay()`
- [PutBuilder](/docs/api/put/) — `NewPut(imageID)`
- [DeleteBuilder](/docs/api/delete/) — `NewDelete(mode)`
- [FrameBuilder](/docs/api/animation/) — `NewFrame(imageID)`
- [AnimateBuilder](/docs/api/animation/) — `NewAnimate(imageID)`
- [ComposeBuilder](/docs/api/animation/) — `NewCompose(imageID)`
- [QueryBuilder](/docs/api/query/) — `NewQuery()`

## Core Types

- [Command](/docs/api/command/) — Protocol command and encoding
- [Response](/docs/api/response/) — Parsed terminal response

## Constants

- [Formats & Transmit Mediums](/docs/api/constants/#formats)
- [Delete Modes](/docs/api/constants/#delete-modes)
- [Animation States & Composition](/docs/api/constants/#animation)
- [Response Suppression](/docs/api/constants/#response-suppression)

## Helper Functions

- [Transmit Helpers](/docs/api/helpers/#transmit)
- [Delete Helpers](/docs/api/helpers/#delete)
- [Animation Helpers](/docs/api/helpers/#animation)
- [Image Utilities](/docs/api/helpers/#image-utilities)
