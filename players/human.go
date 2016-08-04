package players

import (
  "strings"
  "strconv"
)

type HumanPlayer struct {
  name, marker string
  ui interface{
    Read() string
  }
}

func NewHumanPlayer(name, marker string) *HumanPlayer{
  return &HumanPlayer{name: name, marker: marker}
}

type GameBoard interface {
  AvailableSpots() []int
}

func (h HumanPlayer) Name() string{
  return h.name
}

func (h HumanPlayer) Marker() string{
  return h.marker
}

func (h HumanPlayer) SelectSpot(board GameBoard) int {
  spot := -1
  spotIsUnavailable := true
  for spotIsUnavailable{
    spotString := h.ui.Read()
    spot, _ := strconv.Atoi(strings.TrimSpace(spotString))
    if contains(board.AvailableSpots(), spot){
      return spot
    }
  }
  return spot
}

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
