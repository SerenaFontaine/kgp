package kgp

// FrameBuilder builds a frame action command for animations.
type FrameBuilder struct {
	cmd       *Command
	frameData []byte
}

// NewFrame creates a new frame action builder for the specified image ID.
func NewFrame(imageID uint32) *FrameBuilder {
	fb := &FrameBuilder{
		cmd: NewCommand(ActionFrame),
	}
	fb.cmd.SetKeyUint32("i", imageID)
	return fb
}

// FrameData sets the frame image data.
func (fb *FrameBuilder) FrameData(data []byte) *FrameBuilder {
	fb.frameData = data
	return fb
}

// Format sets the image format for the frame.
func (fb *FrameBuilder) Format(format Format) *FrameBuilder {
	fb.cmd.SetKeyUint32("f", uint32(format))
	return fb
}

// Dimensions sets the frame dimensions (required for RGB/RGBA formats).
func (fb *FrameBuilder) Dimensions(width, height int) *FrameBuilder {
	fb.cmd.SetKeyInt("s", width)
	fb.cmd.SetKeyInt("v", height)
	return fb
}

// BackgroundFrame sets the frame number to use as background (0 = base image).
func (fb *FrameBuilder) BackgroundFrame(frameNum uint32) *FrameBuilder {
	fb.cmd.SetKeyUint32("c", frameNum)
	return fb
}

// Gap sets the frame delay in milliseconds.
func (fb *FrameBuilder) Gap(milliseconds uint32) *FrameBuilder {
	fb.cmd.SetKeyUint32("z", milliseconds)
	return fb
}

// Composition sets the composition mode.
func (fb *FrameBuilder) Composition(mode CompositionMode) *FrameBuilder {
	fb.cmd.SetKeyUint32("X", uint32(mode))
	return fb
}

// BackgroundColor sets the background color (32-bit RGBA).
func (fb *FrameBuilder) BackgroundColor(rgba uint32) *FrameBuilder {
	fb.cmd.SetKeyUint32("Y", rgba)
	return fb
}

// ResponseSuppression controls which responses the terminal sends.
func (fb *FrameBuilder) ResponseSuppression(mode ResponseSuppression) *FrameBuilder {
	fb.cmd.SetKeyUint32("q", uint32(mode))
	return fb
}

// Build constructs the final command.
func (fb *FrameBuilder) Build() *Command {
	if len(fb.frameData) > 0 {
		fb.cmd.SetPayload(fb.frameData)
	}
	return fb.cmd
}

// AnimateBuilder builds an animate action command for controlling animation playback.
type AnimateBuilder struct {
	cmd *Command
}

// NewAnimate creates a new animate action builder for the specified image ID.
func NewAnimate(imageID uint32) *AnimateBuilder {
	ab := &AnimateBuilder{
		cmd: NewCommand(ActionAnimate),
	}
	ab.cmd.SetKeyUint32("i", imageID)
	return ab
}

// State sets the animation state.
func (ab *AnimateBuilder) State(state AnimationState) *AnimateBuilder {
	ab.cmd.SetKeyUint32("s", uint32(state))
	return ab
}

// LoopCount sets the number of times to loop (0 = infinite, default = 1).
func (ab *AnimateBuilder) LoopCount(count uint32) *AnimateBuilder {
	ab.cmd.SetKeyUint32("v", count)
	return ab
}

// GapOverride overrides all frame gaps with the specified delay in milliseconds.
func (ab *AnimateBuilder) GapOverride(milliseconds uint32) *AnimateBuilder {
	ab.cmd.SetKeyUint32("z", milliseconds)
	return ab
}

// Frame sets the frame number to stop at (used with AnimationStopCurrent state).
func (ab *AnimateBuilder) Frame(frameNum uint32) *AnimateBuilder {
	ab.cmd.SetKeyUint32("c", frameNum)
	return ab
}

// ResponseSuppression controls which responses the terminal sends.
func (ab *AnimateBuilder) ResponseSuppression(mode ResponseSuppression) *AnimateBuilder {
	ab.cmd.SetKeyUint32("q", uint32(mode))
	return ab
}

// Build constructs the final command.
func (ab *AnimateBuilder) Build() *Command {
	return ab.cmd
}

// ComposeBuilder builds a compose action command for composing animation frames.
type ComposeBuilder struct {
	cmd *Command
}

// NewCompose creates a new compose action builder for the specified image ID.
func NewCompose(imageID uint32) *ComposeBuilder {
	cb := &ComposeBuilder{
		cmd: NewCommand(ActionCompose),
	}
	cb.cmd.SetKeyUint32("i", imageID)
	return cb
}

// BackgroundFrame sets the background frame number.
func (cb *ComposeBuilder) BackgroundFrame(frameNum uint32) *ComposeBuilder {
	cb.cmd.SetKeyUint32("c", frameNum)
	return cb
}

// Composition sets the composition mode.
func (cb *ComposeBuilder) Composition(mode CompositionMode) *ComposeBuilder {
	cb.cmd.SetKeyUint32("X", uint32(mode))
	return cb
}

// BackgroundColor sets the background color (32-bit RGBA).
func (cb *ComposeBuilder) BackgroundColor(rgba uint32) *ComposeBuilder {
	cb.cmd.SetKeyUint32("Y", rgba)
	return cb
}

// ResponseSuppression controls which responses the terminal sends.
func (cb *ComposeBuilder) ResponseSuppression(mode ResponseSuppression) *ComposeBuilder {
	cb.cmd.SetKeyUint32("q", uint32(mode))
	return cb
}

// Build constructs the final command.
func (cb *ComposeBuilder) Build() *Command {
	return cb.cmd
}

// Helper functions for common animation operations

// PlayAnimation plays an animation once (no looping).
func PlayAnimation(imageID uint32) *Command {
	return NewAnimate(imageID).State(AnimationLoop).LoopCount(2).Build()
}

// PlayAnimationLoop plays an animation with infinite looping.
func PlayAnimationLoop(imageID uint32) *Command {
	return NewAnimate(imageID).State(AnimationLoop).LoopCount(1).Build()
}

// StopAnimation stops an animation at the current frame.
func StopAnimation(imageID uint32) *Command {
	return NewAnimate(imageID).State(AnimationStop).Build()
}
