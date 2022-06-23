package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	//after
	fmt.Println("After Unit Test")
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on mac os")
	}

	result := HelloWorld("Fahri")
	require.Equal(t, "Fahri", result, "Result is not Hello_Fahri, something wrong")
}

func TestSubTest(t *testing.T) {
	t.Run("Fahri", func(t *testing.T) {
		result := HelloWorld("Fahri")
		require.Equal(t, "Hello_Fahri", result, "Result is not Hello_Fahri, something wrong")
	})

	t.Run("Fadil", func(t *testing.T) {
		result := HelloWorld("Fadil")
		require.Equal(t, "Hello_Fahri", result, "Result is not Hello_Fahri, something wrong")
	})
}
