# Conciserr

A libray to make error handling in Go easier

> [!CAUTION]
> The package is still in early development. Many features are not tested or yet to be implemented.

## Key Features

### Try

Package `try` provides utility functions for handling functions that return errors
by automatically logging and terminating execution upon encountering an error.

Overview:
- `Try`: Executes a function returning a value and an error, logging and terminating on failure.

Usage:
- The functions in this package simplify error handling by eliminating the need for explicit error checking. However, they terminate the program on error, making them suitable for cases where errors should be fatal.

### Catch

Package catch provides utility functions for handling errors in a concise and functional way. It allows executing functions that return a value and an error, providing a way to handle errors inline without interrupting the control flow.

Overview:
- `Catch`: Executes the provided function `fn`, which returns a value of type T and an error. If `fn` returns an error, the `orElse` function is called with the error as an argument. The function returns the result of `fn`, even if an error occurs.

### Errors

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
