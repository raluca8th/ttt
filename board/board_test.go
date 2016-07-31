package board

import "testing"

func TestDefaultSize(t *testing.T) {
  board := new(Board)
  boardSize := board.Size()
  if boardSize != 9 {
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}
