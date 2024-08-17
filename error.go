package ego

// Try executes the function and catches thrown exceptions.
func Try[T any](do func() T) (t T, err Exception) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			switch recovered := recovered.(type) {
			case Exception:
				err = recovered
			case error:
				err = addStack(0, recovered)
			default:
				err = Errorf("caught a non-error type: %v", recovered)
			}
		}
	}()
	return do(), nil
}

// AssertNil will throw the error as an exception if err is non-nil.
func AssertNil(err error) {
	if err != nil {
		panic(addStack(0, err))
	}
}

// Unwrap will throw the error as an exception if err is non-nil. Otherwise, it
// returns t.
func Unwrap[T any](t T, err error) T {
	if err != nil {
		panic(addStack(0, err))
	}
	return t
}
