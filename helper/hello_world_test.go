package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fahri")
	if result != "Hello_Fahri" {
		// error
		panic("Result is not Hello_Fahri, something wrong")
	}
}

func TestFailHelloWorld(t *testing.T) {
	result := HelloWorld("Fadil")
	if result != "Hello_Fahri" {
		// error
		t.Fatal("Result must be Hello_Fahri")
	}
}
