## How to run the examples
- `cd <directory> && go run main.go`
- Example: `cd concurrency && go run main.go`

## Go Programming Examples

### Concurrency
- [concurrency/](concurrency/): Show how to use goroutines and channels for concurrent programming in Go.
- [concurrency-buffered-channel/](concurrency-buffered-channel/): Demonstrates buffered channels in Go.
- [concurrency-context/](concurrency-context/): Shows context usage for managing goroutines.
- [concurrency-worker-pool/](concurrency-worker-pool/): Implements a worker pool pattern.

### Interfaces
- [interface/](interface/): Demonstrates basic interface usage with a `Shape` interface that defines behavior for `Circle` and `Rectangle` types.
- [interface-empty/](interface-empty/): Introduces the empty interface (`any` type), which can hold values of any type. Shows type information using `%v` and `%T` format verbs.
- [interface-empty-type-assertion/](interface-empty-type-assertion/): Shows type assertion with the empty interface. Demonstrates the safe type assertion pattern with the "comma ok" idiom to avoid panics.
- [interface-empty-type-switches/](interface-empty-type-switches/): Uses type switches with the empty interface to perform different actions based on the underlying type of a value.
- [interface-stringer/](interface-stringer/): Implements the `Stringer` interface from the `fmt` package to provide custom string representation for a `Person` struct.
- [interface-values/](interface-values/): Shows interface values with a `Describer` interface implemented by multiple types (`Person` and `Product`).
- [interface-values-with-nil/](interface-values-with-nil/): Demonstrates interface behavior with nil pointer receivers. Shows how an interface holding a nil pointer can still call methods if they handle nil properly.
