// Package errors defines common error variables.
package errors

import "errors"

// Common application errors.
var (
	// ErrNotFound indicates that the requested resource was not found.
	ErrNotFound = errors.New("resource not found")

	// ErrAlreadyExists occurs when attempting to create a resource that already exists.
	ErrAlreadyExists = errors.New("resource already exists")

	// ErrInvalidArgument is returned when an input argument is invalid.
	ErrInvalidArgument = errors.New("invalid argument")

	// ErrPermissionDenied indicates that the operation is not permitted due to insufficient privileges.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrUnauthenticated occurs when authentication is required but has not been provided or failed.
	ErrUnauthenticated = errors.New("unauthenticated")

	// ErrTimeout indicates that an operation exceeded the allowed time limit.
	ErrTimeout = errors.New("operation timed out")

	// ErrRateLimited occurs when too many requests have been made in a given time frame.
	ErrRateLimited = errors.New("rate limit exceeded")

	// ErrIOFailure signifies an input/output operation failure.
	ErrIOFailure = errors.New("input/output operation failed")

	// ErrDatabaseError represents a failure during a database operation.
	ErrDatabaseError = errors.New("database operation failed")

	// ErrNetworkFailure indicates a failure in a network request.
	ErrNetworkFailure = errors.New("network request failed")

	// ErrConfigInvalid occurs when the provided configuration is invalid.
	ErrConfigInvalid = errors.New("invalid configuration")

	// ErrConflict indicates that the requested operation conflicts with the current state.
	ErrConflict = errors.New("conflicting operation")

	// ErrPreconditionFailed occurs when a required condition for the operation is not met.
	ErrPreconditionFailed = errors.New("precondition failed")

	// ErrOutOfRange is returned when a value is outside the acceptable range.
	ErrOutOfRange = errors.New("value out of range")

	// ErrStateInvalid indicates that the operation cannot be performed due to an invalid state.
	ErrStateInvalid = errors.New("invalid state for operation")
)
