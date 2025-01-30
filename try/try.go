// Package try provides utility functions for handling functions that return errors
// by automatically logging and terminating execution upon encountering an error.
//
// Overview:
//   - `Try`: Executes a function returning a value and an error, logging and terminating on failure.
//   - `TryMultReturn`: Executes a function returning multiple values (including an error), logging and terminating on failure.
//
// Usage:
//   The functions in this package simplify error handling by eliminating the need for explicit error checking.
//   However, they terminate the program on error, making them suitable for cases where errors should be fatal.
//
// Performance Warning:
//   - `TryMultReturn` uses reflection to handle arbitrary function signatures, which incurs runtime overhead.
//   - Consider using direct error handling instead of `TryMultReturn` in performance-sensitive applications.

package try

import (
	"log"
)

// Try executes a function that returns a value of type T and an error.
// If the function returns an error, it logs the error message and terminates the program.
//
// Parameters:
//   - fn: A function that returns a value of type T and an error.
//   - msg: A formatted error message to log before terminating the program.
//
// Returns:
//   - The value of type T returned by fn if no error occurs.
//
// Note:
//   - This function calls log.Fatalf on error, which terminates the program.
//   - Ensure the function passed to Try does not return an error in normal execution flow.
func Try[T any](fn func() (T, error), msg string) T {
	result, err := fn()
	if err != nil {
		log.Fatalf(msg, err)
	}

	return result
}

// TryMultReturn executes a function that returns multiple values, including an error.
// If the function returns an error, it logs the error message and terminates the program.
//
// Parameters:
//   - fn: A function that returns at least one value, with the last return value being an error.
//   - msg: A formatted error message to log before terminating the program.
//
// Returns:
//   - The non-error return values of the function, either as a single value or a tuple.
//
// Notes:
//   - This function **uses reflection**, which introduces runtime overhead and may impact performance.
//   - Reflection-based type conversion can cause panics if type assertions fail.
//   - Ensure that the passed function strictly follows the expected return signature:
//     the last return value must be of type `error`.
//   - The program will **terminate** if an error occurs, as it calls `log.Fatalf`.
//
// Performance Considerations:
//   - Reflection-based function calls are significantly slower than direct function calls.
//   - This should only be used in scenarios where flexibility is required over performance.
// func TryMultReturn[T any](fn interface{}, msg string) T {
// 	fnValue := reflect.ValueOf(fn)
// 	if fnValue.Kind() != reflect.Func {
// 		panic("HandleError expects a function")
// 	}
//
// 	// Call the function dynamically
// 	results := fnValue.Call(nil)
//
// 	// Check if last return value is an error
// 	lastIndex := len(results) - 1
// 	if lastIndex < 0 {
// 		panic("HandleError expects a function that returns at least one value")
// 	}
//
// 	errValue := results[lastIndex]
// 	if errValue.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
// 		if !errValue.IsNil() {
// 			log.Fatalf(msg, errValue.Interface())
// 		}
// 	}
//
// 	// Extract the non-error return values
// 	var returnValue T
// 	if lastIndex == 1 {
// 		// Single non-error return value
// 		returnValue = results[0].Interface().(T)
// 	} else {
// 		// Multiple non-error return values packed as tuple
// 		var tuple []interface{}
// 		for i := 0; i < lastIndex; i++ {
// 			tuple = append(tuple, results[i].Interface())
// 		}
// 		returnValue = any(tuple).(T)
// 	}
//
// 	return returnValue
// }
