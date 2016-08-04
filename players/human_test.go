package players

import (
  "testing"
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

func TestNewHumanPlayer(t * testing.T){
  humanPlayer := NewHumanPlayer("Anda", "A")

  if name := humanPlayer.Name(); name != "Anda" {
    t.Error("Expected name to be Anda, but it was", name)
  }

  if marker := humanPlayer.Marker(); marker != "A" {
    t.Error("Expected marker to be A, but it was", marker)
  }
}

func TestSelectSpot(t *testing.T) {
  stdin := testSTDIN{}
  stdin.buffer.WriteString("5")
  stdout := testSTDOUT{}
  ui := testUI{input: stdin, output: stdout}
  humanPlayer := HumanPlayer{name: "Anda", marker: "A", ui: &ui}
  board := new(testBoard)
  board.availableSpots = []int{5}

  if spot := humanPlayer.SelectSpot(board); spot != 5 {
    t.Error("Expected spot to be 5, but it was", spot)
  }
}

func TestSelectAvailableSpot(t *testing.T) {
  stdin := testSTDIN{}
  stdin.buffer.WriteString("3 6")
  stdout := testSTDOUT{}
  ui := testUI{input: stdin, output: stdout}
  humanPlayer := HumanPlayer{name: "Anda", marker: "A", ui: &ui}
  board := new(testBoard)
  board.availableSpots = []int{6}

  if spot := humanPlayer.SelectSpot(board); spot != 6 {
    t.Error("Expected spot to be 6, but it was", spot)
  }
}

type testBoard struct {
  availableSpots []int
}

func (board *testBoard) AvailableSpots() []int {
  return board.availableSpots
}

type testSTDIN struct {
  buffer bytes.Buffer
}

func (r *testSTDIN) Read() string{
  input, _ := r.buffer.ReadString(' ')
  return input
}

type testSTDOUT struct {
  buffer bytes.Buffer
}

func (p *testSTDOUT) Print(s string) {
  p.buffer.WriteString(s)
}

type testUI struct{
  input testSTDIN
  output testSTDOUT
}

func (ui *testUI) Read() string{
  return ui.input.Read()
}

func (ui *testUI) Print(s string) {
  ui.output.Print(s)
}
