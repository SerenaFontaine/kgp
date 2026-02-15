package kgp

// QueryBuilder builds a query action command for checking terminal capabilities.
type QueryBuilder struct {
	cmd *Command
}

// NewQuery creates a new query action builder.
func NewQuery() *QueryBuilder {
	return &QueryBuilder{
		cmd: NewCommand(ActionQuery),
	}
}

// Format sets the image format to test.
func (qb *QueryBuilder) Format(format Format) *QueryBuilder {
	qb.cmd.SetKeyUint32("f", uint32(format))
	return qb
}

// Dimensions sets test image dimensions.
func (qb *QueryBuilder) Dimensions(width, height int) *QueryBuilder {
	qb.cmd.SetKeyInt("s", width)
	qb.cmd.SetKeyInt("v", height)
	return qb
}

// TransmitMedium sets the transmission medium to test.
func (qb *QueryBuilder) TransmitMedium(medium TransmitMedium) *QueryBuilder {
	qb.cmd.SetKey("t", string(medium))
	return qb
}

// TestData sets minimal test data.
func (qb *QueryBuilder) TestData(data []byte) *QueryBuilder {
	qb.cmd.SetPayload(data)
	return qb
}

// Build constructs the final command.
func (qb *QueryBuilder) Build() *Command {
	return qb.cmd
}

// QuerySupport returns a query command to check if the terminal supports
// the Kitty Graphics Protocol. The terminal will respond with OK if supported.
func QuerySupport() *Command {
	// Minimal query with small test data
	return NewQuery().
		Format(FormatRGB).
		Dimensions(1, 1).
		TransmitMedium(TransmitDirect).
		TestData([]byte{0, 0, 0}).
		Build()
}
