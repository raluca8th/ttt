package tttboard

import (
  "math"
  "ttt/board"
  "ttt/ui"
)

type TTTBoard struct {
  size int
  markers [2]string
  surface []string
}

func NewBoard(params Params) board.Board {
  board := TTTBoard{size: params.Size, markers: params.Markers}
  board.setDefaultSize()
  board.setDefaultMarkers()
  board.setSurface()
  return board
}

func (b TTTBoard) Size() int {
  return b.size
}

func (b TTTBoard) Markers() [2]string {
  return b.markers
}

func (b TTTBoard) Surface() []string {
  return b.surface
}

func (b TTTBoard) FillSpot(spot int) {
  b.surface[spot] = b.NextMarker()
}

func (b TTTBoard) SpotIsAvailable(spot int) bool{
  return b.surface[spot] == ""
}

func (b TTTBoard) ResetSpot(spot int){
   b.surface[spot] = ""
}

func (b TTTBoard) AvailableSpots() []int{
  var availableSpots []int
  for i, spot := range b.Surface() {
    if spot == "" {
      availableSpots = append(availableSpots, i)
    }
  }
  return availableSpots
}

func (b TTTBoard) NextMarker() string{
  if odd(b.Size()) {
    if even(len(b.AvailableSpots())) {
      return b.markers[1]
    } else {
      return b.markers[0]
    }
  } else {
    if even(len(b.AvailableSpots())) {
      return b.markers[0]
    } else {
      return b.markers[1]
    }
  }
}

func (b TTTBoard) WinningMarker() string{
  checkBoard := b.checkRows()
  checkBoard += b.checkColumns()
  checkBoard += b.checkLeftDiagonal()
  checkBoard += b.checkRightDiagonal()
  return checkBoard
}

func (b TTTBoard) IsTiedBoard() bool{
  return len(b.AvailableSpots()) == 0 && b.WinningMarker()== ""
}

func (b TTTBoard) IsBoardSolved() bool{
  return b.WinningMarker() != "" || b.IsTiedBoard()
}

func (b TTTBoard) checkRows() string {
  incrementor := b.incrementor()
  for i := 0; i < b.Size(); i += incrementor{
    boardSubSet := b.Surface()[i:i+incrementor]
    if identicalElements(boardSubSet, incrementor) {
      return boardSubSet[0]
    }
  }
  return ""
}

func (b TTTBoard) checkColumns() string {
  b.surface = b.transposeBoard()
  return b.checkRows()
}

func (b TTTBoard) checkLeftDiagonal() string {
  var leftDiagonal []string
  for i := 0; i < b.Size(); i += b.incrementor() + 1 {
    leftDiagonal = append(leftDiagonal, b.Surface()[i])
  }
  return checkRow(leftDiagonal)
}

func (b TTTBoard) checkRightDiagonal() string {
  var rightDiagonal []string
  for i := 0; i < b.incrementor(); i++ {
    index := int(math.Abs(float64(i - ((b.incrementor() * (i + 1)) - 1))))
    rightDiagonal = append(rightDiagonal, b.Surface()[index])
  }
  return checkRow(rightDiagonal)
}

func (b *TTTBoard) setSurface() {
  b.surface = make([]string, b.Size())
}

func (b *TTTBoard) setDefaultSize() {
  if b.size == 0{
    b.size = 9
  }
}

func (b *TTTBoard) setDefaultMarkers() {
  if b.markers == [2]string{} {
    b.markers = [2]string{"X", "Y"}
  }
}

func (b TTTBoard) transposeBoard() []string{
  var transposedBoard []string
  incrementor := b.incrementor()
  for i := 0; i < incrementor; i++ {
    for j:= 0; j < b.Size(); j += incrementor {
      transposedBoard = append(transposedBoard, b.Surface()[i + j])
    }
  }
  return transposedBoard
}

func checkRow(row []string) string {
  marker := ""
  if identicalElements(row, len(row)) {
    marker = row[0]
  }
  return marker
}

func identicalElements(boardSubSection []string, incrementor int) bool{
  elementsAreIdentical := false
  elementsMap := map[string]int{}
  for _, element := range boardSubSection {
    if element != "" {
      elementsMap[element] = elementsMap[element] + 1
      if elementsMap[element] == incrementor {
        elementsAreIdentical = true
      }
    }
  }
  return elementsAreIdentical
}

func (b TTTBoard) incrementor() int {
  return int(math.Sqrt(float64(b.Size())))
}

func even(number int) bool{
  return number % 2 == 0
}

func odd(number int) bool{
  return !even(number)
}

type Params struct {
  Size int
  Markers [2]string
  UI ui.UI
}
