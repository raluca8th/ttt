package cli

import (
  "testing"
  "bytes"
)

type testSTDINReader struct {
  buffer bytes.Buffer
}

func (r *testSTDINReader) Read() string{
  return r.buffer.String()
}

func TestCLIInput(t *testing.T) {
  reader := new(testSTDINReader)
  reader.buffer.WriteString("A")

  if input := reader.Read(); input != "A" {
    t.Error("Expected input to be 'A', but it was", input)
  }
}
