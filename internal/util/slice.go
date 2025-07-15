package util

import "golang.org/x/exp/slices"

// FindFunc searches for the first element in slice s that satisfies the predicate function f.
// It returns the found element and true if found, otherwise returns the zero value of type E and false.
//
// Parameters:
//   - s: the slice to search through
//   - f: predicate function that tests each element
//
// Returns:
//   - E: the found element (or zero value if not found)
//   - bool: true if element was found, false otherwise
func FindFunc[E any](s []E, f func(E) bool) (E, bool) {
    // Use slices.IndexFunc to find the first matching index
    if i := slices.IndexFunc(s, f); i != -1 {
        return s[i], true
    }
    
    // Return zero value and false if no element matches
    var e E
    return e, false
}