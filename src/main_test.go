package main

import (
	"testing"
	"fmt"
)

func TestHelloWorld(t *testing.T) {
	// if helloWorld("") != "Hello, World!" {
	// 	t.Error("helloWorld(\"\") nao esta retornando \"Hello, World!\"")
	// }

	fmt.Println(helloWorld("fulano"))
	if helloWorld("fulano") != "Hello, fulano!" {
		t.Error("helloWorld(\"fulano\") nao esta retornando \"Hello, <name>!\"")
	}
}