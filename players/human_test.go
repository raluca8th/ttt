package players

import (
  "testing"
  "ttt/board"
  "ttt/cli"
  "bytes"
)

func TestName(t *testing.T) {
  humanPlayer := HumanPlayer{name: "Anda"}
  if name := humanPlayer.Name(); name != "Anda" {
    t.Error("Expected name to be Anda, but it was", name)
  }
}

func TestMarker(t *testing.T) {
  humanPlayer := HumanPlayer{name: "Anda", marker: "A"}
  if marker := humanPlayer.Marker(); marker != "A" {
    t.Error("Expected marker to be A, but it was", marker)
  }
}

func TestSelectSpot(t *testing.T) {
  board := board.NewBoard(board.Params{})
  stdinReader := new(testSTDINReader)
  stdinReader.buffer.WriteString("5")
  cliReader := cli.CLIInput{Reader: stdinReader}
  testPrinter := new(testSTDOUTprinter)
  cliPrinter := cli.CLIOutput{Printer: testPrinter}
  ui := cli.UI{Input: cliReader, Output: cliPrinter}
  humanPlayer := HumanPlayer{name: "Anda", marker: "A", ui: ui}

  if spot := humanPlayer.SelectSpot(board); spot != 5 {
    t.Error("Expected spot to be 5, but it was", spot)
  }
}

type testSTDINReader struct {
  buffer bytes.Buffer
}

func (r *testSTDINReader) Read() string{
  return r.buffer.String()
}

type testSTDOUTprinter struct {
  buffer bytes.Buffer
}

func (p *testSTDOUTprinter) Print(s string) {
  p.buffer.WriteString(s)
}