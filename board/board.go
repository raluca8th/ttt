package board

type Board struct {
  size int
  markers [2]string
}

func (b Board) Size() int {
  b.setDefaultSize()
  return b.size
}

func (b Board) Markers() [2]string {
  b.setDefaultMarkers()
  return b.markers
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
