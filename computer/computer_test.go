package players

import (
  "testing"
  "bytes"
  "ttt/tttBoard"
  "ttt/board"
//  "fmt"
)
/*
func TestName(t *testing.T) {
  computerPlayer := ComputerPlayer{name: "Walle"}
  if name := computerPlayer.Name(); name != "Walle" {
    t.Error("Expected name to be Walle, but it was", name)
  }
}

func TestMarker(t *testing.T) {
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W"}
  if marker := computerPlayer.Marker(); marker != "W" {
    t.Error("Expected marker to be W, but it was", marker)
  }
}

func TestNewComputerPlayer(t * testing.T){
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := NewComputerPlayer("Wallee", "W", ui)

  if name := computerPlayer.Name(); name != "Wallee" {
    t.Error("Expected name to be Wallee, but it was", name)
  }

  if marker := computerPlayer.Marker(); marker != "W" {
    t.Error("Expected marker to be W, but it was", marker)
  }
}

func TestSelectFirstSpot(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})

  if spot := computerPlayer.SelectSpot(board); spot != 4 {
    t.Error("Expected spot to be 4, but it was", spot)
  }
}
*/
func TestSelectWinningSpot(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})
  fillSpots(board, []int{0, 3, 4, 5})

  if spot := computerPlayer.SelectSpot(board); spot != 8 {
    t.Error("Expected spot to be 8, but it was", spot)
  }
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
  input *testSTDIN
  output *testSTDOUT
}

func (ui testUI) Read() string{
  return ui.input.Read()
}

func (ui testUI) Print(strings ...string) {
  for _, s := range strings {
    ui.output.Print(s)
  }
}

func fillSpots(b board.Board, spots []int) {
  for _, index := range spots {
    b.FillSpot(index)
  }
}
