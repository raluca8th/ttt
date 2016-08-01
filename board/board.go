package board

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
