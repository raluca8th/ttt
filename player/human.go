package player

import (
  "strings"
  "strconv"
  "github.com/raluca8th/ttt/zen"
  "github.com/raluca8th/ttt/ui"
  "github.com/raluca8th/ttt/board"
)

type HumanPlayer struct {
  name, marker string
  ui ui.UI
}

func NewHumanPlayer(name, marker string, ui ui.UI) *HumanPlayer {
  return &HumanPlayer{name: name, marker: marker, ui: ui}
}

func (h *HumanPlayer) Name() string {
  return h.name
}

func (h *HumanPlayer) Marker() string {
  return h.marker
}

func (h *HumanPlayer) SelectSpot(board board.Board) int {
  spot := -1
  spotIsUnavailable := true
  for spotIsUnavailable{
    h.ui.Print(strings.ToUpper(h.name), spotSelection)
    spotString := h.ui.Read()
    spotsAvailable := zen.IntToStringSlice(board.AvailableSpots())
    if zen.ContainsString(spotsAvailable, spotString) {
      spot, _ := strconv.Atoi(spotString)
      return spot
    }
    h.ui.Print(unavailableSpot, zen.ToString(board.AvailableSpots()), newLine) }
    return spot
  }

  const (
    spotSelection = ", please select your spot:\n"
    unavailableSpot = "Not a valid spot. Please select one of: "
    newLine = "\n"
  )
