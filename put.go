package kgp

// PutBuilder builds a put (placement) action command.
type PutBuilder struct {
	cmd *Command
}

// NewPut creates a new put action builder for the specified image ID.
func NewPut(imageID uint32) *PutBuilder {
	pb := &PutBuilder{
		cmd: NewCommand(ActionPut),
	}
	pb.cmd.SetKeyUint32("i", imageID)
	return pb
}

// ImageNumber sets the image number (application-specific, non-unique).
func (pb *PutBuilder) ImageNumber(num uint32) *PutBuilder {
	pb.cmd.SetKeyUint32("I", num)
	return pb
}

// PlacementID sets the placement ID (auto-generated if not specified).
func (pb *PutBuilder) PlacementID(id uint32) *PutBuilder {
	pb.cmd.SetKeyUint32("p", id)
	return pb
}

// CellOffset sets the pixel offset within the starting cell.
func (pb *PutBuilder) CellOffset(x, y int) *PutBuilder {
	pb.cmd.SetKeyInt("X", x)
	pb.cmd.SetKeyInt("Y", y)
	return pb
}

// DisplaySize sets the display size in terminal cells.
func (pb *PutBuilder) DisplaySize(columns, rows int) *PutBuilder {
	pb.cmd.SetKeyInt("c", columns)
	pb.cmd.SetKeyInt("r", rows)
	return pb
}

// SourceRect displays only a portion of the source image.
func (pb *PutBuilder) SourceRect(x, y, width, height int) *PutBuilder {
	pb.cmd.SetKeyInt("x", x)
	pb.cmd.SetKeyInt("y", y)
	pb.cmd.SetKeyInt("w", width)
	pb.cmd.SetKeyInt("h", height)
	return pb
}

// ZIndex sets the z-index (negative = below text, positive = above text).
func (pb *PutBuilder) ZIndex(z int) *PutBuilder {
	pb.cmd.SetKeyInt("z", z)
	return pb
}

// CursorMovement controls whether the cursor moves after placement.
func (pb *PutBuilder) CursorMovement(move bool) *PutBuilder {
	if move {
		pb.cmd.SetKey("C", "0")
	} else {
		pb.cmd.SetKey("C", "1")
	}
	return pb
}

// VirtualPlacement creates an invisible placement for Unicode placeholder use.
func (pb *PutBuilder) VirtualPlacement() *PutBuilder {
	pb.cmd.SetKey("U", "1")
	return pb
}

// RelativeTo sets the parent placement for relative positioning.
func (pb *PutBuilder) RelativeTo(parentImageID, parentPlacementID uint32, offsetH, offsetV int) *PutBuilder {
	pb.cmd.SetKeyUint32("P", parentImageID)
	pb.cmd.SetKeyUint32("Q", parentPlacementID)
	pb.cmd.SetKeyInt("H", offsetH)
	pb.cmd.SetKeyInt("V", offsetV)
	return pb
}

// ResponseSuppression controls which responses the terminal sends.
func (pb *PutBuilder) ResponseSuppression(mode ResponseSuppression) *PutBuilder {
	pb.cmd.SetKeyUint32("q", uint32(mode))
	return pb
}

// Build constructs the final command.
func (pb *PutBuilder) Build() *Command {
	return pb.cmd
}
