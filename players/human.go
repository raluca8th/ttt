package players

import (
  "ttt/cli"
  "ttt/board"
)

type HumanPlayer struct {
  name, marker string
  ui cli.UI
}

func (h HumanPlayer) Name() string{
  return h.name
}

func (h HumanPlayer) Marker() string{
  return h.marker
}

func (h HumanPlayer) UI() cli.UI{
  return h.ui
}

func (h HumanPlayer) SelectSpot(board *board.Board) int {
  return 5
}
