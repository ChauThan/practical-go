## How to run the examples
- `go run <filename>.go`

## Go Programming Examples

### Concurrency
- `concurrency.go`: Show how to use goroutines and channels for concurrent programming in Go.

### Interfaces
- `interface.go`: Demonstrates basic interface usage with a `Shape` interface that defines behavior for `Circle` and `Rectangle` types.
- `interface-empty.go`: Introduces the empty interface (`any` type), which can hold values of any type. Shows type information using `%v` and `%T` format verbs.
- `interface-empty-type-assertion.go`: Shows type assertion with the empty interface. Demonstrates the safe type assertion pattern with the "comma ok" idiom to avoid panics.
- `interface-empty-type-switches.go`: Uses type switches with the empty interface to perform different actions based on the underlying type of a value.
- `interface-stringer.go`: Implements the `Stringer` interface from the `fmt` package to provide custom string representation for a `Person` struct.
- `interface-values.go`: Shows interface values with a `Describer` interface implemented by multiple types (`Person` and `Product`).
- `interface-values-with-nil.go`: Demonstrates interface behavior with nil pointer receivers. Shows how an interface holding a nil pointer can still call methods if they handle nil properly.
