package stringset

import "fmt"

var testKeys *StringSet

func ExampleNew() {
	// Initialize new stringset
	testKeys = New()
}

func ExampleStringSet_Set() {
	// Set foo key
	testKeys.Set("foo")
	// Set bar key
	testKeys.Set("bar")
}

func ExampleStringSet_Unset() {
	// Remove bar key
	testKeys.Unset("bar")
}

func ExampleStringSet_Has() {
	if testKeys.Has("foo") {
		fmt.Println("We have foo!")
	}

	if !testKeys.Has("bar") {
		fmt.Println("We do not have bar!")
	}
}

func ExampleStringSet_Slice() {
	keys := testKeys.Slice()
	for _, key := range keys {
		fmt.Printf("Iterating key: %s\n", key)
	}
}
