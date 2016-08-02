package main

import "testing"

func ExampleWelcome() {
  main()
  // Output: Welcome to my TTT
}

func TestHello(t *testing.T) {
  expected := "Welcome to my TTT"
  returned := welcome()
  if returned != expected {
    t.Fatalf("Expected %s, got %s", expected, returned)
  }
}
