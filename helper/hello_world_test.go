package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fahri")
	if result != "Hello_Fahri" {
		// error
		panic("Result is not Hello_Fahri, something wrong")
	}
}
