package main

import "testing"

func testRune(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
