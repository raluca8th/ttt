package player

import (
  "testing"
  "github.com/raluca8th/ttt/board"
  "github.com/raluca8th/ttt/mocks"
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
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  humanPlayer := NewHumanPlayer("Anda", "A", ui)

  if name := humanPlayer.Name(); name != "Anda" {
    t.Error("Expected name to be Anda, but it was", name)
  }

  if marker := humanPlayer.Marker(); marker != "A" {
    t.Error("Expected marker to be A, but it was", marker)
  }
}

func TestSelectSpot(t *testing.T) {
  stdin := new(mocks.TestSTDIN)
  stdin.Buffer.WriteString("5")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  humanPlayer := HumanPlayer{name: "Anda", marker: "A", ui: &ui}
  board := new(testBoard)
  board.availableSpots = []int{5}

  if spot := humanPlayer.SelectSpot(board); spot != 5 {
    t.Error("Expected spot to be 5, but it was", spot)
  }
}

func TestSelectAvailableSpot(t *testing.T) {
  stdin := new(mocks.TestSTDIN)
  stdin.Buffer.WriteString("f 3 6")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  humanPlayer := HumanPlayer{name: "Anda", marker: "A", ui: &ui}
  board := new(testBoard)
  board.availableSpots = []int{0, 6}

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

func (board *testBoard) SpotIsAvailable(spot int) bool {return false}
func (board *testBoard) FillSpot(spot int){}
func (board *testBoard) WinningMarker() string{return ""}
func (board *testBoard) IsBoardSolved() bool{return false}
func (board *testBoard) Size() int{return 0}
func (board *testBoard) Markers() [2]string{return [2]string{}}
func (board *testBoard) Surface() []string{return []string{}}
func (board *testBoard) IsTiedBoard() bool{return false}
func (board *testBoard) NextMarker() string{return ""}
func (board *testBoard) ResetSpot(spot int) {}
func (board *testBoard) FillAvailableSpot(spot int) board.Board{return nil}
