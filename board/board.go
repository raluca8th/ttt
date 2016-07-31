package board

type Board struct {
  size int
}

func (b Board) Size() int {
  b.setDefaultSize()
  return b.size
}

func (b *Board) setDefaultSize() {
  if b.size == 0{
    b.size = 9
  }
}
