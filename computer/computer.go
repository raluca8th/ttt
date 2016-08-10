package computer

import (
  "github.com/raluca8th/ttt/ui"
  "github.com/raluca8th/ttt/board"
  "github.com/raluca8th/ttt/zen"
  "math/rand"
)

type ComputerPlayer struct {
  name, marker string
  ui ui.UI
}

func NewComputerPlayer(name, marker string, ui ui.UI) *ComputerPlayer{
  return &ComputerPlayer{name: name, marker: marker, ui: ui}
}

func (c *ComputerPlayer) Name() string{
  return c.name
}

func (c *ComputerPlayer) Marker() string{
  return c.marker
}

func (c *ComputerPlayer) SelectSpot(board board.Board) int{
  var depth int
  availableSpots := board.AvailableSpots()
  bestScore := make(map[int]int)
  c.ui.Print(computerThinking)

  if len(availableSpots) >= 12 {
    return availableSpots[rand.Intn(len(availableSpots))]
  }

  ch := make(chan *ChannelValues)
  for _, spot := range board.AvailableSpots() {
    go func(spot int) {
      negamaxChannel(board.FillAvailableSpot(spot), ch, depth, spot)
    }(spot)
  }

  for range board.AvailableSpots() {
    channelValue := <-ch
    bestScore[channelValue.spot] = channelValue.score
  }

  return zen.MaxValueKey(bestScore)
}

func negamaxChannel(board board.Board, ch chan *ChannelValues, depth, originalSpot int) {
  channelValue := new(ChannelValues)
  channelValue.spot = originalSpot
  channelValue.score = -1 * negamax(board, depth)
  ch <- channelValue
}

func negamax(board board.Board, depth int) int{
  var bestScore = map[int]int{}

  if board.IsTiedBoard(){
    return 0
  } else if board.IsBoardSolved() {
    return -1
  }

  for _, spot := range board.AvailableSpots() {
    board.FillSpot(spot)
    bestScore[spot] = -1 * negamax(board, depth + 1)
    board.ResetSpot(spot)
  }

  maxScore := zen.MaxValue(bestScore)
  return maxScore
}

func availableSpots(board board.Board) []int{
  return board.AvailableSpots()
}

type ChannelValues struct {
  spot int
  score int
}

const (
  computerThinking = "Minimax is now working hard. Think happy thoughts...\n"
)
