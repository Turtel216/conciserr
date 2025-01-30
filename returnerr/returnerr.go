// Package returnerr provides utilities for handling error returns in Go.
package returnerr

import (
	"fmt"
)

// ReturnErr executes the provided function and returns an error with a custom message if an error occurs.
//
// Parameters:
// - fn: A function that returns an error.
// - msg: A custom error message to wrap the returned error.
//
// Returns:
// - An error with the custom message if fn() returns an error, otherwise nil.
func ReturnErr(fn func() error, msg string) error {
	err := fn()
	if err != nil {
		// Use fmt.Errorf to properly format and return the error
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}
