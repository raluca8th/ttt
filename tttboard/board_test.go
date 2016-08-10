package tttboard

import (
  "testing"
  "reflect"
  "github.com/raluca8th/ttt/board"
)

func TestSize(t *testing.T) {
  board := NewBoard(16, [2]string{"X", "Y"})
  boardSize := board.Size()
  if boardSize != 16 {
    t.Error("Expected board size to be 16, but it was", boardSize)
  }
}

func TestMarkers(t *testing.T) {
  board := NewBoard(9, [2]string{"A", "B"})
  boardMarkers := board.Markers()
  if boardMarkers != [2]string{"A", "B"} {
    t.Error("Expected markers to be 'A' and 'B', but they were", boardMarkers)
  }
}

func TestNewBoardSurface(t *testing.T) {
  board := NewBoard(9, [2]string{"A", "B"})
  boardSurface := board.Surface()
  if surfaceLenght := len(boardSurface); surfaceLenght != 9 {
    t.Error("Expected board surface to have lenght of 9, but it was", surfaceLenght)
  }

  if spot := boardSurface[3]; spot != "" {
    t.Error("Expected spot to be empty but it was", spot)
  }
}

func TestFillSpot(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  boardSurface := board.Surface()
  board.FillSpot(2)
  if spot := boardSurface[2]; spot != "X" {
    t.Error("Expected marker to be 'X'  but it was", spot)
  }
}

func TestSpotIsAvailable(t *testing.T) {
  board := NewBoard(9, [2]string{"A", "B"})
  board.FillSpot(2)
  if spotIsAvailable := board.SpotIsAvailable(2); spotIsAvailable != false {
    t.Error("Expected SpotIsAvailable to be false, but it was", spotIsAvailable)
  }
}

func TestAvailableSpots(t *testing.T) {
  board := NewBoard(9, [2]string{"A", "B"})
  fillSpots(board, []int{2, 4, 8})
  expectedResult := make([]int, 0, 9)
  expectedResult = append(expectedResult, 0, 1, 3, 5, 6, 7)
  availableSpots := board.AvailableSpots()
  if !reflect.DeepEqual(availableSpots, expectedResult) {
    t.Error("Expected {0, 1, 3, 5, 6, 7}, but it was", availableSpots)
  }
}

func TestWinningMarkerWinningRow(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{0, 3, 1, 4, 2})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
  if isBoardSolved := board.IsBoardSolved(); isBoardSolved != true {
    t.Error("Expected IsBoardSolved to be true, but it was", isBoardSolved)
  }
}

func TestWinningMarkerWinningColumn(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{0, 1, 3, 4, 6, 7})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerWinningLeftDiagonal(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{0, 1, 4, 2, 8})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerWinningRightDiagonal(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{2, 1, 4, 3, 6})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerBoardNotSolved(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{1, 4, 2, 5, 3})
  if winningMarker := board.WinningMarker(); winningMarker != "" {
    t.Error("Expected Winning Marker to be '', but it was", winningMarker)
  }
}

func TestIsBoardTied(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  fillSpots(board, []int{0, 1, 2, 3, 5, 4, 6, 8, 7})
  if isTiedBoard := board.IsTiedBoard(); isTiedBoard != true {
    t.Error("Expected IsTiedBoard to return true, but it returned", isTiedBoard)
  }
  if isBoardSolved := board.IsBoardSolved(); isBoardSolved != true {
    t.Error("Expected IsBoardSolved to be true, but it was", isBoardSolved)
  }
}

func TestNextMarker(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  if nextMarker := board.NextMarker(); nextMarker != "X" {
    t.Error("Expected marker to be X, but it was", nextMarker)
  }
  board.FillSpot(1)
  if nextMarker := board.NextMarker(); nextMarker != "Y" {
    t.Error("Expected marker to be Y, but it was", nextMarker)
  }
}

func TestFillAvailableSpot(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  newBoard := board.FillAvailableSpot(2)

  if spots := len(newBoard.AvailableSpots()); spots != 8 {
    t.Error("Expected number of available spots to be 8, but they were", spots)
  }

  if spots := len(board.AvailableSpots()); spots != 9 {
    t.Error("Expected number of available spots to be 9, but they were", spots)
  }
}

func TestNextMarker4X4Board(t *testing.T) {
  board := NewBoard(9, [2]string{"X", "Y"})
  if nextMarker := board.NextMarker(); nextMarker != "X" {
    t.Error("Expected marker to be X, but it was", nextMarker)
  }
  board.FillSpot(1)
  if nextMarker := board.NextMarker(); nextMarker != "Y" {
    t.Error("Expected marker to be Y, but it was", nextMarker)
  }
}

func fillSpots(b board.Board, spots []int) {
  for _, index := range spots {
    b.FillSpot(index)
  }
}
