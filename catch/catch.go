// Package catch provides utility functions for handling errors in a concise and functional way.
// It allows executing functions that return a value and an error, providing a way to handle
// errors inline without interrupting the control flow.
package catch

// Catch executes the provided function `fn`, which returns a value of type T and an error.
// If `fn` returns an error, the `orElse` function is called with the error as an argument.
// The function returns the result of `fn`, even if an error occurs.
//
// Example usage:
//
//	result := Catch(
//	    func() (int, error) {
//	        return someFunction()
//	    },
//	    func(err error) {
//	        log.Println("An error occurred:", err)
//	    },
//	)
//
// In this example, `someFunction` is executed, and if it returns an error,
// the provided error-handling function logs the error.
func Catch[T any](fn func() (T, error), orElse func(error)) T {
	result, err := fn()
	if err != nil {
		orElse(err)
	}

	return result
}
