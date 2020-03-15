package main

import "testing"

func TestMain(t *testing.T) {
	isCalled := false
	outFn = func(a ...interface{}) (n int, err error) {
		isCalled = true
		return 0, nil
	}

	main()

	if !isCalled {
		t.Error("function should have been called but not")
	}
}
