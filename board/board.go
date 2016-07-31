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
  return b.markers
}

func (b *Board) setDefaultSize() {
  if b.size == 0{
    b.size = 9
  }
}
