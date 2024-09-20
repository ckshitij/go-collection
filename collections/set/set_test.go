package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	// Test NewSet
	s := NewSet[int]()
	if s == nil {
		t.Fatal("NewSet() returned nil")
	}

	// Test Add and Contains
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(1) {
		t.Errorf("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Errorf("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Errorf("Set should contain 3")
	}
	if s.Contains(4) {
		t.Errorf("Set should not contain 4")
	}

	// Test Remove
	s.Remove(2)
	if s.Contains(2) {
		t.Errorf("Set should not contain 2 after removal")
	}

	// Test removing an element not in the set
	s.Remove(4) // Should not panic or throw errors
	if s.Contains(4) {
		t.Errorf("Set should not contain 4 after removal")
	}
}

func TestStringSet(t *testing.T) {
	// Test with string type
	s := NewSet[string]()
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	if !s.Contains("apple") {
		t.Errorf("Set should contain 'apple'")
	}
	if !s.Contains("banana") {
		t.Errorf("Set should contain 'banana'")
	}
	if s.Contains("grape") {
		t.Errorf("Set should not contain 'grape'")
	}

	// Remove an element and check
	s.Remove("banana")
	if s.Contains("banana") {
		t.Errorf("Set should not contain 'banana' after removal")
	}
}
