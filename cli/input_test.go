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
  cliReader := CLIInput{reader: reader}
  if input := cliReader.Read(); input != "A" {
    t.Error("Expected input to be 'A', but it was", input)
  }
}
