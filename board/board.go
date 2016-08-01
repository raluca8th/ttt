package board

import (
  "math"
)

type Board struct {
  size int
  markers [2]string
  surface []string
}

func NewBoard(params Params) *Board {
  board := &Board{size: params.size, markers: params.markers}
  board.setDefaultSize()
  board.setDefaultMarkers()
  board.setSurface()
  return board
}

func (b Board) Size() int {
  return b.size
}

func (b Board) Markers() [2]string {
  return b.markers
}

func (b Board) Surface() []string {
  return b.surface
}

func (b *Board) FillSpot(spot int, marker string) {
  b.surface[spot] = marker
}

func (b Board) SpotIsAvailable(spot int) bool{
  return b.surface[spot] == ""
}

func (b Board) AvailableSpots() []int{
  availableSpots := make([]int, 0, b.Size())
  for i, spot := range b.Surface() {
    if spot == "" {
      availableSpots = append(availableSpots, i)
    }
  }
  return availableSpots
}

func (b Board) WinningMarker() string{
  checkBoard := b.checkRows()
  checkBoard += b.checkColumns()
  checkBoard += b.checkLeftDiagonal()
  checkBoard += b.checkRightDiagonal()
  return checkBoard
}

func (b Board) IsTiedBoard() bool{
  return len(b.AvailableSpots()) == 0 && b.WinningMarker()== ""
}

func (b Board) checkRows() string {
  incrementor := b.incrementor()
  for i := 0; i < b.Size(); i += incrementor{
    boardSubSet := b.Surface()[i:i+incrementor]
    if identicalElements(boardSubSet) {
      return boardSubSet[0]
    }
  }
  return ""
}

func (b Board) checkColumns() string {
  b.surface = b.transposeBoard()
  return b.checkRows()
}

func (b Board) checkLeftDiagonal() string {
  leftDiagonal := make([]string, 0)
  for i := 0; i < b.Size(); i += b.incrementor() + 1 {
    leftDiagonal = append(leftDiagonal, b.Surface()[i])
  }
  return checkRow(leftDiagonal)
}

func (b Board) checkRightDiagonal() string {
  rightDiagonal := make([]string, 0)
  for i := 0; i < b.incrementor(); i++ {
    index := int(math.Abs(float64(i - ((b.incrementor() * (i + 1)) - 1))))
    rightDiagonal = append(rightDiagonal, b.Surface()[index])
  }
  return checkRow(rightDiagonal)
}

func (b *Board) setSurface() {
  b.surface = make([]string, b.Size())
}

func (b *Board) setDefaultSize() {
  if b.size == 0{
    b.size = 9
  }
}

func (b *Board) setDefaultMarkers() {
  if b.markers == [2]string{} {
    b.markers = [2]string{"X", "Y"}
  }
}

func identicalElements(boardSubSection []string) bool{
  elementsAreIdentical := false
  elementsMap := make(map[string]int)
  for _, element := range boardSubSection {
    if element != "" {
      elementsMap[element] = elementsMap[element] + 1
      if elementsMap[element] == 3 {
        elementsAreIdentical = true
      }
    }
  }
  return elementsAreIdentical
}

func checkRow(row []string) string {
  marker := ""
  if identicalElements(row) {
    marker = row[0]
  }
  return marker
}

func (b Board) transposeBoard() []string{
  transposedBoard := make([]string, 0)
  incrementor := b.incrementor()
  for i := 0; i < incrementor; i++ {
    for j:= 0; j < b.Size(); j += incrementor {
      transposedBoard = append(transposedBoard, b.Surface()[i + j])
    }
  }
  return transposedBoard
}

func (b Board) incrementor() int {
  return int(math.Sqrt(float64(b.Size())))
}

type Params struct {
  size int
  markers [2]string
}
