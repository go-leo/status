package util

import (
	"testing"
)

func TestFindFunc(t *testing.T) {
	// Test case 1: Element found (int)
	t.Run("FoundInt", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		f := func(e int) bool { return e%2 == 0 } // Find first even number
		val, ok := FindFunc(s, f)
		if !ok {
			t.Error("Expected to find element, but got false")
		}
		if val != 2 {
			t.Errorf("Expected 2, got %v", val)
		}
	})

	// Test case 2: Element not found (string)
	t.Run("NotFoundString", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		f := func(e string) bool { return e == "d" } // Non-existent element
		val, ok := FindFunc(s, f)
		if ok {
			t.Error("Expected not to find element, but got true")
		}
		if val != "" {
			t.Errorf("Expected empty string, got %v", val)
		}
	})

	// Test case 3: Empty slice (custom struct)
	type Person struct{ Name string }
	t.Run("EmptySlice", func(t *testing.T) {
		s := []Person{}
		f := func(e Person) bool { return e.Name == "Alice" }
		val, ok := FindFunc(s, f)
		if ok {
			t.Error("Expected not to find element, but got true")
		}
		if val != (Person{}) {
			t.Errorf("Expected zero value, got %v", val)
		}
	})

	// Test case 4: First matching element returned (int)
	t.Run("FirstMatchReturned", func(t *testing.T) {
		s := []int{1, 2, 4, 6}
		f := func(e int) bool { return e%2 == 0 } // Multiple matches
		val, ok := FindFunc(s, f)
		if !ok {
			t.Error("Expected to find element, but got false")
		}
		if val != 2 { // Should return first match (2), not 4 or 6
			t.Errorf("Expected 2, got %v", val)
		}
	})
}
