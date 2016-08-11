package computer

import (
  "testing"
  "github.com/raluca8th/ttt/tttboard"
  "github.com/raluca8th/ttt/board"
  "github.com/raluca8th/ttt/mocks"
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
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  computerPlayer := NewComputerPlayer("Wallee", "W", ui)

  if name := computerPlayer.Name(); name != "Wallee" {
    t.Error("Expected name to be Wallee, but it was", name)
  }

  if marker := computerPlayer.Marker(); marker != "W" {
    t.Error("Expected marker to be W, but it was", marker)
  }
}

func TestAvailableSpots(t *testing.T){
  board := tttboard.NewBoard(9, [2]string{"W", "I"})
  fillSpots(board, []int{0, 1, 2, 4})

  if availableSpots := availableSpots(board); !reflect.DeepEqual(availableSpots, []int{3, 5, 6, 7, 8}) {
    t.Error("Expected available spots to be 3 5 6 7 8, but they were", availableSpots)
  }
}

func TestStopOpponentFromWinning(t *testing.T) {
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(9, [2]string{"W", "I"})
  fillSpots(board, []int{0, 1, 5, 4})

  if spot := computerPlayer.SelectSpot(board); spot != 7 {
    t.Error("Expected spot to be 7, but it was", spot)
  }
}

func TestStopOpponentFromWinning4X4(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping 4X4 test in short mode")
  }
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(16, [2]string{"W", "I"})
  fillSpots(board, []int{7, 0, 11, 1, 8, 2})

  if spot := computerPlayer.SelectSpot(board); spot != 3 {
    t.Error("Expected selection to be 3, but it was", spot)
  }
}

func TestSelectsSpotWithTheBestChanceOfWinningComputerMovesFirst(t *testing.T) {
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(9, [2]string{"W", "I"})

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
  stdin := new(mocks.TestSTDIN)
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  computerPlayer := ComputerPlayer{name: "Wallee", marker: "W", ui: &ui}
  board := tttboard.NewBoard(9, [2]string{"I", "W"})

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

func fillSpots(b board.Board, spots []int) {
  for _, index := range spots {
    b.FillSpot(index)
  }
}
