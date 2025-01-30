# Conciserr

A libray to make error handling in Go easier

> [!CAUTION]
> The package is still in early development. Many features are not tested or yet to be implemented.

## Key Features

### 1. Try

Package `try` provides utility functions for handling functions that return errors
by automatically logging and terminating execution upon encountering an error.

Overview:
- `Try`: Executes a function returning a value and an error, logging and terminating on failure.

Usage:
- The functions in this package simplify error handling by eliminating the need for explicit error checking. However, they terminate the program on error, making them suitable for cases where errors should be fatal.

### 2. Catch

Package `catch` provides utility functions for handling errors in a concise and functional way. It allows executing functions that return a value and an error, providing a way to handle errors inline without interrupting the control flow.

Overview:
- `Catch`: Executes the provided function `fn`, which returns a value of type T and an error. If `fn` returns an error, the `orElse` function is called with the error as an argument. The function returns the result of `fn`, even if an error occurs.

### 3. ReturnErr

Package `returnerr` provides utilities for handling error returns in Go.

Overview:

- `ReturnErr`: Executes the provided function and returns an error with a custom message if an error occurs. Returns nil on success.

### 3. Errors

Package errors defines common error variables.

- **ErrNotFound** indicates that the requested resource was not found.

- **ErrAlreadyExists** occurs when attempting to create a resource that already exists.

- **ErrInvalidArgument** is returned when an input argument is invalid.

- **ErrPermissionDenied** indicates that the operation is not permitted due to insufficient privileges.

- **ErrUnauthenticated** occurs when authentication is required but has not been provided or failed.

- **ErrTimeout** indicates that an operation exceeded the allowed time limit.

- **ErrRateLimited** occurs when too many requests have been made in a given time frame.

- **ErrIOFailure** signifies an input/output operation failure.

- **ErrDatabaseError** represents a failure during a database operation.

- **ErrNetworkFailure** indicates a failure in a network request.

- **ErrConfigInvalid** occurs when the provided configuration is invalid.

- **ErrConflict** indicates that the requested operation conflicts with the current state.

- **ErrPreconditionFailed** occurs when a required condition for the operation is not met.

- **ErrOutOfRange** is returned when a value is outside the acceptable range.

- **ErrStateInvalid** indicates that the operation cannot be performed due to an invalid state.

---

## Installation
To use utils in your project, add it as a module dependency:

### Clone the Repository
```
go get github.com/Turtel216/conciserr
```
Alternatively, include it directly in your go.mod file (use the latest release):
```
require github.com/Turtel216/conciserr v0.0.1
```

---

# Contributions
Contributions to this project are welcome! If you would like to contribute, please feel free to open a PR.
