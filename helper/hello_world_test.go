package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fahri")
	if result != "Hello_Fahri" {
		// error
		panic("Result is not Hello_Fahri, something wrong")
	}
	fmt.Println("[TestHelloWorld] : unit test successfully")
}

func TestFailHelloWorld(t *testing.T) {
	result := HelloWorld("Fadil")
	if result != "Hello_Fahri" {
		// error
		t.Fatal("Result must be Hello_Fahri")
	}
	fmt.Println("[TestFailHelloWorld] : unit test successfully")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fahri")
	assert.Equal(t, "Hello_Fahri", result, "Result is not Hello_Fahri, something wrong")
	fmt.Println("[TestHelloWorldAssert] : unit test successfully")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fahri")
	require.Equal(t, "Hello_Fahri", result, "Result is not Hello_Fahri, something wrong")
	fmt.Println("[TestHelloWorldRequire] : unit test successfully")
}
