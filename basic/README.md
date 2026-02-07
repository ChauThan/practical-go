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
- [interface-errors/](interface-errors/): Implements the error interface with a custom error type. Shows how to create custom errors by implementing the `Error() string` method.
- [interface-images/](interface-images/): Demonstrates using the standard library image interface. Creates an RGBA image and shows how to access its bounds and pixel data using the image interface methods.
- [interface-images-exercise/](interface-images-exercise/): Exercise that generates images using the `golang.org/x/tour/pic` package. Creates a 2D slice and generates a pattern based on the XOR of coordinates.
- [interface-reader/](interface-reader/): Demonstrates the `io.Reader` interface, a core abstraction for reading data in Go. Shows how to read data in chunks from a `strings.Reader` until reaching `io.EOF`.

### Slices
- [slice-exercice/](slice-exercice/): Demonstrates slice operations including creating 2D slices and using the `golang.org/x/tour/pic` package for image generation. Shows how to generate a 2D pixel array using the XOR operation.
