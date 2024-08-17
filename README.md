# Ego

Ego is an experimental library for exception handling. It models try-catch
exception handling using panic and recover.

## Catching and Throwing Exceptions

The `panic` keyword is used to throw exceptions. `ego.Errorf` should be used to
create errors with stack traces attached. Exceptions are caught using
`ego.Try`.

```go
func ThisFunctionThrows() {
   panic(ego.Errorf("hello %s", "world"))
}

func ThisFunctionCatches() int {
    value, err := ego.Try(func() int {
        ThisFunctionThrows()
        return 42
    })
    if err != nil {
        // '%+v' prints the error and a stack trace
        fmt.Printf("%+v\n", err)
        return 0
    }
    return value
}
```

## Calling Idiomatic Go

The `ego.Unwrap` and `ego.AssertNil` helpers should be used to convert
traditional Go errors into exceptions.

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
create goroutines. `ego.Go` returns an `ego.Future` that the caller can wait
on. `Wait` will panic if the goroutine panics, allowing the caller to handle
the exception.

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
}
```

## Stack Traces

Formatting an exception with `'%+v'` prints a stack trace.

```go
func Print(e ego.Exception) {
    fmt.Printf("%+v\n", e)
}
```

Here's an example stack trace.

```
this is the .Error() message
0: ego.frameA exception_test.go:11
1: ego.frameB exception_test.go:15
2: ego.frameC exception_test.go:19
3: ego.TestFormatException exception_test.go:23
```
