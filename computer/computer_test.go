package computer

import (
  "testing"
  "bytes"
  "github.com/raluca8th/ttt/tttboard"
  "github.com/raluca8th/ttt/board"
  "math/rand"
  "reflect"
)

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

func TestAvailableSpots(t *testing.T){
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})
  fillSpots(board, []int{0, 1, 2, 4})

  if availableSpots := availableSpots(board); !reflect.DeepEqual(availableSpots, []int{3, 5, 6, 7, 8}) {
    t.Error("Expected available spots to be 3 5 6 7 8, but they were", availableSpots)
  }
}

func TestStopOpponentFromWinning(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})
  fillSpots(board, []int{0, 1, 5, 4})

  if spot := computerPlayer.SelectSpot(board); spot != 7 {
    t.Error("Expected spot to be 7, but it was", spot)
  }
}

func TestSelectsSpotWithTheBestChanceOfWinningComputerMovesFirst(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})

  for true {
    computerSpot := computerPlayer.SelectSpot(board)
    fillSpots(board, []int{computerSpot})
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
    availableSpots := board.AvailableSpots()
    opponentPick := availableSpots[rand.Intn(len(availableSpots))]
    fillSpots(board, []int{opponentPick})
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
  }

  if winnerMarker := board.WinningMarker(); winnerMarker == "I" {
    t.Error("Expected winner marker to be W, but it was", winnerMarker)
  }
}

func TestSelectsSpotWithTheBestChanceOfWinningComputerMovesSecond(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"I", "W"}})

  for true {
    availableSpots := board.AvailableSpots()
    opponentPick := availableSpots[rand.Intn(len(availableSpots))]
    fillSpots(board, []int{opponentPick})
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
    computerSpot := computerPlayer.SelectSpot(board)
    fillSpots(board, []int{computerSpot})
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
  }

  if winnerMarker := board.WinningMarker(); winnerMarker == "I" {
    t.Error("Expected winner marker to be W, but it was", winnerMarker)
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

func (ui testUI) PrintBoard(board []string){}

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
