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
- `TryMultReturn`: Executes a function returning multiple values (including an error), logging and terminating on failure.

Usage:
- The functions in this package simplify error handling by eliminating the need for explicit error checking. However, they terminate the program on error, making them suitable for cases where errors should be fatal.

Performance Warning:
- `TryMultReturn` uses reflection to handle arbitrary function signatures, which incurs runtime overhead.
- Consider using direct error handling instead of `TryMultReturn` in performance-sensitive applications.

### Catch

### Errors

---

## Installation
To use utils in your project, add it as a module dependency:

### Clone the Repository
```
go get github.com/kashifkhan0771/utils
```
Alternatively, include it directly in your go.mod file (use the latest release):
```
require github.com/kashifkhan0771/utils v0.3.0
```

---

# Contributions
Contributions to this project are welcome! If you would like to contribute, please feel free to open a PR.
