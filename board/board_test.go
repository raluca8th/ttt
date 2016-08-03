package board

import (
  "testing"
  "reflect"
)

func TestDefaultSize(t *testing.T) {
  board := NewBoard(Params{})
  boardSize := board.Size()
  if boardSize != 9 {
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}

func TestSize(t *testing.T) {
  board := NewBoard(Params{Size: 16})
  boardSize := board.Size()
  if boardSize != 16 {
    t.Error("Expected board size to be 16, but it was", boardSize)
  }
}

func TestMarkers(t *testing.T) {
  board := NewBoard(Params{Markers: [2]string{"A", "B"}})
  boardMarkers := board.Markers()
  if boardMarkers != [2]string{"A", "B"} {
    t.Error("Expected markers to be 'A' and 'B', but they were", boardMarkers)
  }
}

func TestDefaultMarkers(t *testing.T) {
  board := NewBoard(Params{})
  boardMarkers := board.Markers()
  if boardMarkers != [2]string{"X", "Y"} {
    t.Error("Expected markers to be 'X' and 'Y', but they were", boardMarkers)
  }
}

func TestNewBoardSurface(t *testing.T) {
  board := NewBoard(Params{})
  boardSurface := board.Surface()
  if surfaceLenght := len(boardSurface); surfaceLenght != 9 {
    t.Error("Expected board surface to have lenght of 9, but it was", surfaceLenght)
  }

  if spot := boardSurface[3]; spot != "" {
    t.Error("Expected spot to be empty but it was", spot)
  }
}

func TestFillSpot(t *testing.T) {
  board := NewBoard(Params{})
  boardSurface := board.Surface()
  board.FillSpot(2)
  if spot := boardSurface[2]; spot != "X" {
    t.Error("Expected marker to be 'X'  but it was", spot)
  }
}

func TestSpotIsAvailable(t *testing.T) {
  board := NewBoard(Params{})
  board.FillSpot(2)
  if spotIsAvailable := board.SpotIsAvailable(2); spotIsAvailable != false {
    t.Error("Expected SpotIsAvailable to be false, but it was", spotIsAvailable)
  }
}

func TestAvailableSpots(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{2, 4, 8})
  expectedResult := make([]int, 0, 9)
  expectedResult = append(expectedResult, 0, 1, 3, 5, 6, 7)
  availableSpots := board.AvailableSpots()
  if !reflect.DeepEqual(availableSpots, expectedResult) {
    t.Error("Expected {0, 1, 3, 5, 6, 7}, but it was", availableSpots)
  }
}

func TestWinningMarkerWinningRow(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{0, 3, 1, 4, 2})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
  if isBoardSolved := board.IsBoardSolved(); isBoardSolved != true {
    t.Error("Expected IsBoardSolved to be true, but it was", isBoardSolved)
  }
}

func TestWinningMarkerWinningColumn(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{0, 1, 3, 4, 6, 7})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerWinningLeftDiagonal(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{0, 1, 4, 2, 8})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerWinningRightDiagonal(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{2, 1, 4, 3, 6})
  if winningMarker := board.WinningMarker(); winningMarker != "X" {
    t.Error("Expected Winning Marker to be 'X', but it was", winningMarker)
  }
}

func TestWinningMarkerBoardNotSolved(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{1, 4, 2, 5, 3})
  if winningMarker := board.WinningMarker(); winningMarker != "" {
    t.Error("Expected Winning Marker to be '', but it was", winningMarker)
  }
}

func TestIsBoardTied(t *testing.T) {
  board := NewBoard(Params{})
  fillSpots(board, []int{0, 1, 2, 3, 5, 4, 6, 8, 7})
  if isTiedBoard := board.IsTiedBoard(); isTiedBoard != true {
    t.Error("Expected IsTiedBoard to return true, but it returned", isTiedBoard)
  }
  if isBoardSolved := board.IsBoardSolved(); isBoardSolved != true {
    t.Error("Expected IsBoardSolved to be true, but it was", isBoardSolved)
  }
}

func TestNextMarker(t *testing.T) {
  board := NewBoard(Params{})
  if nextMarker := board.NextMarker(); nextMarker != "X" {
    t.Error("Expected marker to be X, but it was", nextMarker)
  }
  board.FillSpot(1)
  if nextMarker := board.NextMarker(); nextMarker != "Y" {
    t.Error("Expected marker to be Y, but it was", nextMarker)
  }
}

func TestNextMarker4X4Board(t *testing.T) {
  board := NewBoard(Params{Size: 16})
  if nextMarker := board.NextMarker(); nextMarker != "X" {
    t.Error("Expected marker to be X, but it was", nextMarker)
  }
  board.FillSpot(1)
  if nextMarker := board.NextMarker(); nextMarker != "Y" {
    t.Error("Expected marker to be Y, but it was", nextMarker)
  }
}

func fillSpots(b *Board, spots []int) {
  for _, index := range spots {
    b.FillSpot(index)
  }
}
