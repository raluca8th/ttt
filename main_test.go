package main

import "testing"

func TestHello(t *testing.T) {
  expected := "Welcome to my TTT"
  returned := welcome()
  if returned != expected {
    t.Fatalf("Expected %s, got %s", expected, returned)
  }
}
