package kgp

import (
	"strings"
	"testing"
)

// TestNewQuery tests creating a query builder
func TestNewQuery(t *testing.T) {
	qb := NewQuery()
	if qb == nil {
		t.Fatal("NewQuery returned nil")
	}
	if qb.cmd.controlData["a"] != string(ActionQuery) {
		t.Errorf("Action should be %s, got %s", ActionQuery, qb.cmd.controlData["a"])
	}
}

// TestQueryBuilder_Format tests setting format
func TestQueryBuilder_Format(t *testing.T) {
	qb := NewQuery()
	qb.Format(FormatPNG)
	if qb.cmd.controlData["f"] != "100" {
		t.Errorf("Expected format '100', got %s", qb.cmd.controlData["f"])
	}
}

// TestQueryBuilder_Dimensions tests setting dimensions
func TestQueryBuilder_Dimensions(t *testing.T) {
	qb := NewQuery()
	qb.Dimensions(800, 600)
	if qb.cmd.controlData["s"] != "800" {
		t.Errorf("Expected width '800', got %s", qb.cmd.controlData["s"])
	}
	if qb.cmd.controlData["v"] != "600" {
		t.Errorf("Expected height '600', got %s", qb.cmd.controlData["v"])
	}
}

// TestQueryBuilder_TransmitMedium tests setting transmission medium
func TestQueryBuilder_TransmitMedium(t *testing.T) {
	qb := NewQuery()
	qb.TransmitMedium(TransmitFile)
	if qb.cmd.controlData["t"] != string(TransmitFile) {
		t.Errorf("Expected transmit medium 'f', got %s", qb.cmd.controlData["t"])
	}
}

// TestQueryBuilder_TestData tests setting test data
func TestQueryBuilder_TestData(t *testing.T) {
	qb := NewQuery()
	testData := []byte("test")
	qb.TestData(testData)

	cmd := qb.Build()
	if string(cmd.payload) != "test" {
		t.Errorf("Expected payload 'test', got %s", string(cmd.payload))
	}
}

// TestQueryBuilder_Build tests building the command
func TestQueryBuilder_Build(t *testing.T) {
	qb := NewQuery()
	qb.Format(FormatRGB)
	qb.Dimensions(1, 1)

	cmd := qb.Build()
	if cmd == nil {
		t.Fatal("Build returned nil")
	}
}

// TestQuerySupport tests the helper function
func TestQuerySupport(t *testing.T) {
	cmd := QuerySupport()
	if cmd == nil {
		t.Fatal("QuerySupport returned nil")
	}

	encoded := cmd.Encode()

	// Should contain query action
	if !strings.Contains(encoded, "a=q") {
		t.Error("Should contain action=q")
	}

	// Should contain RGB format
	if !strings.Contains(encoded, "f=24") {
		t.Error("Should contain format=24 (RGB)")
	}

	// Should contain 1x1 dimensions
	if !strings.Contains(encoded, "s=1") {
		t.Error("Should contain width=1")
	}
	if !strings.Contains(encoded, "v=1") {
		t.Error("Should contain height=1")
	}

	// Should contain direct transmission
	if !strings.Contains(encoded, "t=d") {
		t.Error("Should contain transmit=d (direct)")
	}

	// Should have payload
	if len(cmd.payload) == 0 {
		t.Error("Query should have test data payload")
	}
}

// TestQueryBuilder_CompleteFlow tests a complete query flow
func TestQueryBuilder_CompleteFlow(t *testing.T) {
	testData := []byte{0, 0, 0}
	cmd := NewQuery().
		Format(FormatRGB).
		Dimensions(1, 1).
		TransmitMedium(TransmitDirect).
		TestData(testData).
		Build()

	encoded := cmd.Encode()

	if !strings.Contains(encoded, "a=q") {
		t.Error("Should contain action=q")
	}
	if !strings.Contains(encoded, "f=24") {
		t.Error("Should contain format RGB")
	}
	if !strings.Contains(encoded, "s=1") {
		t.Error("Should contain width 1")
	}
	if !strings.Contains(encoded, "v=1") {
		t.Error("Should contain height 1")
	}
	if !strings.Contains(encoded, "t=d") {
		t.Error("Should contain direct transmission")
	}
}
