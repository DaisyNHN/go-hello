// hello_test.go
package go_hello_test

import (
	"fmt"
	"github.com/DaisyNHN/go-hello"
	"testing"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := go_hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}

}
