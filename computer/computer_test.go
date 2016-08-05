package players

import (
  "testing"
  "bytes"
  "ttt/tttBoard"
  "ttt/board"
  "math/rand"
//  "fmt"
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

func TestSelectWinningSpot(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})
  fillSpots(board, []int{0, 3, 1, 5})

  if spot := computerPlayer.SelectSpot(board); spot != 2 {
    t.Error("Expected spot to be 2, but it was", spot)
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

func TestSelectsSpotWithTheBestChanceOfWinning(t *testing.T) {
  stdin := new(testSTDIN)
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(tttboard.Params{Size: 9, Markers: [2]string{"W", "I"}})

  for true {
    computerSpot := computerPlayer.SelectSpot(board)
    fillSpots(board, []int{computerSpot})
    t.Log("computer", computerSpot)
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
    availableSpots := board.AvailableSpots()
    t.Log(availableSpots)
    opponentPick := availableSpots[rand.Intn(len(availableSpots))]
    t.Log("opponent", opponentPick)
    fillSpots(board, []int{opponentPick})
    t.Log(board)
    if board.IsTiedBoard() || board.IsBoardSolved() {
      break
    }
  }

  t.Log("FINAL STATE", board)
  t.Log("WinnerMarker", board.WinningMarker())

  if winnerMarker := board.WinningMarker(); winnerMarker == "I" {
    t.Error("Expected spot to be W, but it was", winnerMarker)
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
