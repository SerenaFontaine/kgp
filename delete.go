package kgp

// DeleteBuilder builds a delete action command.
type DeleteBuilder struct {
	cmd *Command
}

// NewDelete creates a new delete action builder with the specified mode.
func NewDelete(mode DeleteMode) *DeleteBuilder {
	db := &DeleteBuilder{
		cmd: NewCommand(ActionDelete),
	}
	db.cmd.SetKey("d", string(mode))
	return db
}

// ImageID sets the image ID for deletion (used with DeleteByImageID/DeleteByImageIDFree).
func (db *DeleteBuilder) ImageID(id uint32) *DeleteBuilder {
	db.cmd.SetKeyUint32("i", id)
	return db
}

// ImageNumber sets the image number for deletion (used with DeleteByImageNumber/DeleteByImageNumberFree).
func (db *DeleteBuilder) ImageNumber(num uint32) *DeleteBuilder {
	db.cmd.SetKeyUint32("I", num)
	return db
}

// PlacementID sets the placement ID for deletion (used with DeleteByPlacementID/DeleteByPlacementIDFree).
func (db *DeleteBuilder) PlacementID(id uint32) *DeleteBuilder {
	db.cmd.SetKeyUint32("p", id)
	return db
}

// Cell sets the cell coordinates for deletion (used with DeleteByCell/DeleteByCellFree).
func (db *DeleteBuilder) Cell(x, y int) *DeleteBuilder {
	db.cmd.SetKeyInt("x", x)
	db.cmd.SetKeyInt("y", y)
	return db
}

// Column sets the column for deletion (used with DeleteByColumn/DeleteByColumnFree).
func (db *DeleteBuilder) Column(x int) *DeleteBuilder {
	db.cmd.SetKeyInt("x", x)
	return db
}

// Row sets the row for deletion (used with DeleteByRow/DeleteByRowFree).
func (db *DeleteBuilder) Row(y int) *DeleteBuilder {
	db.cmd.SetKeyInt("y", y)
	return db
}

// ZIndex sets the z-index for deletion (used with DeleteByZIndex/DeleteByZIndexFree).
func (db *DeleteBuilder) ZIndex(z int) *DeleteBuilder {
	db.cmd.SetKeyInt("z", z)
	return db
}

// ResponseSuppression controls which responses the terminal sends.
func (db *DeleteBuilder) ResponseSuppression(mode ResponseSuppression) *DeleteBuilder {
	db.cmd.SetKeyUint32("q", uint32(mode))
	return db
}

// Build constructs the final command.
func (db *DeleteBuilder) Build() *Command {
	return db.cmd
}

// Helper functions for common deletion operations

// DeleteAll deletes all placements and preserves image data.
func DeleteAll() *Command {
	return NewDelete(DeleteAllPlacements).Build()
}

// DeleteAllFree deletes all placements and frees image data.
func DeleteAllFree() *Command {
	return NewDelete(DeleteAllPlacementsFree).Build()
}

// DeleteImage deletes all placements of the specified image and preserves data.
func DeleteImage(imageID uint32) *Command {
	return NewDelete(DeleteByImageID).ImageID(imageID).Build()
}

// DeleteImageFree deletes all placements of the specified image and frees data.
func DeleteImageFree(imageID uint32) *Command {
	return NewDelete(DeleteByImageIDFree).ImageID(imageID).Build()
}

// DeleteAtCursor deletes images at the cursor position and preserves data.
func DeleteAtCursor() *Command {
	return NewDelete(DeleteByCursor).Build()
}

// DeleteAtCursorFree deletes images at the cursor position and frees data.
func DeleteAtCursorFree() *Command {
	return NewDelete(DeleteByCursorFree).Build()
}
