package players

import (
  "ttt/cli"
  "ttt/board"
  "strconv"
  "strings"
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
  spot := -1
  spotIsUnavailable := true
  for spotIsUnavailable{
    spotString := string(h.ui.Input.Read())
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
