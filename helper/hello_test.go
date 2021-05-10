package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEga(t *testing.T) {
	result := HelloWorld("Ega")
	if result != "Hello Ega" {
		// unit test failed
		// is not a recomended use a panic
		// panic("Result is not Hello Ega")

		t.Error("Result must be 'Hello Ega'")
	}

	fmt.Println("Unit test is done")
}

func TestHelloWorldBer(t *testing.T) {
	result := HelloWorld("Ber")
	if result != "Hello Ber" {
		// unit test failed
		// is not a recomended use a panic
		// panic("Result is not Hello Ber")

		t.Fatal("Result must be 'Hello Ber'")
	}

	fmt.Println("Unit test is done")
}
