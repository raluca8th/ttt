package players

import (
  "ttt/ui"
  "ttt/board"
  "fmt"
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
    for _, spot := range availableSpots(board){
      board.FillSpot(spot)
      if board.IsBoardSolved(){
        return spot
      } else {
        board.FillSpot(spot)
        if board.IsBoardSolved(){
          fmt.Println(board)
          return spot
        } else{
          board.ResetSpot(spot)
        }
      }
      board.ResetSpot(spot)
    }
  }
  return -1
}

func availableSpots(board board.Board) []int{
  return board.AvailableSpots()
}
