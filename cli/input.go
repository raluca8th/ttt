package cli

import "fmt"

type Reader interface {
  Read() string
}

type STDINReader struct {}

func (in *STDINReader) Read() string {
  var input string
  fmt.Scanf("%s", &input)
  return input
}
