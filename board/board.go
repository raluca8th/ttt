package board

import "math"

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

func (b Board) WinningMarker() string {
  return b.checkRows()
}

func (b Board) checkRows() string {
  incrementor := b.incrementor()
  for i := 0; i < b.Size(); i += incrementor {
    boardSubSet := b.Surface()[i:i+incrementor]
    if identicalElements(boardSubSet) {
      return boardSubSet[0]
    }
  }
  return ""
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
  elementsMap := make(map[string]bool)
  for _, element := range boardSubSection {
    if element != "" {
      elementsMap[element] = true
    }
  }
  if len(elementsMap) == 1 {
    elementsAreIdentical = true
  }
  return elementsAreIdentical
}

func (b Board) incrementor() int {
  return int(math.Sqrt(float64(b.Size())))
}

type Params struct {
  size int
  markers [2]string
}
