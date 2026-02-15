---
title: Animation
weight: 6
---

## Basic Animation Flow

1. Transmit base frame (or first frame)
2. Add frames with `NewFrame`
3. Create placement
4. Control playback with `NewAnimate` or helpers

## Transmit Base, Add Frames

```go
// Base image (frame 0)
baseCmd := kgp.NewTransmit().
    ImageID(20).
    Format(kgp.FormatPNG).
    TransmitDirect(frame0Data).
    Build()
fmt.Print(baseCmd.Encode())

// Frame 1
frame1Cmd := kgp.NewFrame(20).
    FrameData(frame1Data).
    Gap(100).
    Build()
fmt.Print(frame1Cmd.Encode())

// Frame 2 with composition
frame2Cmd := kgp.NewFrame(20).
    Format(kgp.FormatRGBA).
    Dimensions(80, 80).
    FrameData(frame2Data).
    Gap(100).
    BackgroundFrame(0).
    Composition(kgp.CompositionBlend).
    Build()
fmt.Print(frame2Cmd.Encode())
```

## Playback Control

```go
// Play once
fmt.Print(kgp.PlayAnimation(20).Encode())

// Loop infinitely
fmt.Print(kgp.PlayAnimationLoop(20).Encode())

// Stop at current frame
fmt.Print(kgp.StopAnimation(20).Encode())

// Reset to first frame
fmt.Print(kgp.ResetAnimation(20).Encode())
```

## Custom Loop Count

Loop count follows the Kitty protocol: 0 is ignored, 1 means infinite, and N>1 loops N-1 times.

```go
// Loop 4 times (LoopCount 5 = play 5-1 = 4 loops)
cmd := kgp.NewAnimate(20).
    State(kgp.AnimationLoop).
    LoopCount(5).
    Build()
fmt.Print(cmd.Encode())
```

## Color-Changing Animation (RGBA Frames)

```go
width, height := 80, 80
redData := kgp.SolidColorImage(width, height, 255, 0, 0, 255)
orangeData := kgp.SolidColorImage(width, height, 255, 165, 0, 255)
yellowData := kgp.SolidColorImage(width, height, 255, 255, 0, 255)
greenData := kgp.SolidColorImage(width, height, 0, 255, 0, 255)

// Transmit base frame as raw RGBA
fmt.Print(kgp.NewTransmit().
    ImageID(300).
    Format(kgp.FormatRGBA).
    Dimensions(width, height).
    TransmitDirect(redData).
    Build().Encode())

// Add frames
for _, data := range [][]byte{orangeData, yellowData, greenData} {
    fmt.Print(kgp.NewFrame(300).
        Format(kgp.FormatRGBA).
        Dimensions(width, height).
        FrameData(data).
        Gap(100).
        Composition(kgp.CompositionReplace).
        Build().Encode())
}

// Place and play
fmt.Print(kgp.NewPut(300).PlacementID(1).Build().Encode())
fmt.Print(kgp.PlayAnimationLoop(300).Encode())
```
