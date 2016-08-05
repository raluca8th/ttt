package players

import (
  "ttt/ui"
  "ttt/board"
)

type ComputerPlayer struct {
  name, marker string
  ui ui.UI
}

func NewComputerPlayer(name, marker string, ui ui.UI) *ComputerPlayer{
  return &ComputerPlayer{name: name, marker: marker, ui: ui}
}

func (h ComputerPlayer) Name() string{
  return h.name
}

func (h ComputerPlayer) Marker() string{
  return h.marker
}

func (h ComputerPlayer) SelectSpot(board board.Board) int {
  if len(board.AvailableSpots()) == 9 {
    return 4
  } else {
    return 0
  }
}
