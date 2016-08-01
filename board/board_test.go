package board

import "testing"

func TestDefaultSize(t *testing.T) {
  board := NewBoard(Params{})
  boardSize := board.Size()
  if boardSize != 9 {
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}

func TestSize(t *testing.T) {
  board := NewBoard(Params{size: 16})
  boardSize := board.Size()
  if boardSize != 16 {
    t.Error("Expected board size to be 16, but it was", boardSize)
  }
}

func TestMarkers(t *testing.T) {
  board := NewBoard(Params{markers: [2]string{"A", "B"}})
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
  board.FillSpot(2, "X")
  if spot := boardSurface[2]; spot != "X" {
    t.Error("Expected marker to be 'X'  but it was", spot)
  }
}

func TestSpotIsAvailable(t *testing.T) {
  board := NewBoard(Params{})
  board.FillSpot(2, "X")
  if spotIsAvailable := board.SpotIsAvailable(2); spotIsAvailable != false {
    t.Error("Expected SpotIsAvailable to be false, but it was", spotIsAvailable)
  }
}
