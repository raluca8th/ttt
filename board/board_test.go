package board

import "testing"

func TestDefaultSize(t *testing.T) {
  board := new(Board)
  boardSize := board.Size()
  if boardSize != 9 {
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}

func TestCustomSize(t *testing.T) {
  board := Board{size: 16}
  boardSize := board.Size()
  if boardSize != 16 {
    t.Error("Expected board size to be 16, but it was", boardSize)
  }
}

func TestMarkers(t *testing.T) {
  board := Board{markers: [2]string{"X", "Y"}}
  boardMarkers := board.Markers()
  if boardMarkers != [2]string{"X", "Y"} {
    t.Error("Expected markers to be 'X' and 'Y', but they were", boardMarkers)
  }
}
