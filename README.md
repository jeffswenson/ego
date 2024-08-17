# Ego

Ego is an experiment in implementing exception handling in go.

## Catching and Throwing Exceptions

The `panic` keyword is used to throw exceptions. `ego.Errorf` should be used to
create errors with stack traces attached. Exceptions are caught using
`ego.Try`.

```
func ThisFunctionThrows() {
   panic(ego.Errorf("hello %s", "world"))
}

func ThisFunctionCatches() int {
    value, err := ego.Try(func() int {
        ThisFunctionThrows()
        return 42
    })
    if err != nil {
        // %+v prints the error and a stack trace
        fmt.Printf("%+v\n", err)
        return 0
    }
    return value
}
```

## Calling Idiomatic Go 

The `ego.Unwrap` and `ego.AssertNil` helpers should be used to convert
traditional go errors into exceptions.

```go
func UnwrapExample() net.Conn {
    return ego.Unwrap(net.Dial("tcp", "golang.org:80"))
}

func AssertNilExample(listener net.Listener) {
    ego.AssertNil(listener.Close())
}
```

## Goroutines

Exception handling with goroutines is tricky because an uncaught exception in
the goroutine will cause the program to crash. `ego.Go` should be used to
create goroutines. `ego.Go` returns a `ego.Future` that the caller can wait on.
Wait will throw if the goroutine throws, which allows the caller to handle the
exception.

```go
func DoSomethingAsync() {
    // Create 2 goroutines that run in parallel
    future1 := ego.Go(func() string {
        return "hello"
    })
    future2 := ego.Go(func() string {
        return "world"
    })
    // Wait for the goroutines to finish. Wait will throw an exception if the
    // goroutine failed.
    fmt.Printf("%s %s", future1.Wait(), future2.Wait())
```
