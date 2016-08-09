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
  c.ui.Print(computerThinking)
  var depth int
  alpha := -100
  beta := 100
  availableSpots := board.AvailableSpots()
  if len(availableSpots) > 13 {
    return availableSpots[rand.Intn(len(availableSpots))]
  }
  return c.minimax(board, depth, alpha, beta)
}

func (c *ComputerPlayer) minimax(board board.Board, depth, alpha, beta int) int{
  if c.maxNode(board) {
    return c.max(board, depth, alpha, beta)
  } else {
    return c.min(board, depth, alpha, beta)
  }
}

func (c *ComputerPlayer) max(board board.Board, depth, alpha, beta int) int{
  var bestScore = map[int]int{}

  if board.IsTiedBoard(){
    return 0
  } else if board.IsBoardSolved() {
    return -1
  }

  for _, spot := range board.AvailableSpots() {
    board.FillSpot(spot)
    bestScore[spot] = c.minimax(board, depth + 1, alpha, beta)
    if alpha >= bestScore[spot] {
      alpha = bestScore[spot]
    }
    if beta <= alpha{
      board.ResetSpot(spot)
      break
    }
    board.ResetSpot(spot)
  }

  bestSpot := zen.MaxValueKey(bestScore)
  maxScore := zen.MaxValue(bestScore)

  if depth == 0 {
    return bestSpot
  }
  return maxScore
}

func (c *ComputerPlayer) min(board board.Board, depth, alpha, beta int) int{
  var bestScore = map[int]int{}

  if board.IsTiedBoard(){
    return 0
  } else if board.IsBoardSolved() {
    return 1
  }

  for _, spot := range board.AvailableSpots() {
    board.FillSpot(spot)
    bestScore[spot] = c.minimax(board, depth + 1, alpha, beta)
    if beta <= bestScore[spot] {
      beta = bestScore[spot]
    }

    if beta <= alpha {
      board.ResetSpot(spot)
      break
    }
    board.ResetSpot(spot)
  }

  bestSpot := zen.MinValueKey(bestScore)
  minScore := zen.MinValue(bestScore)

  if depth == 0 {
    return bestSpot
  }
  return minScore
}

func (c *ComputerPlayer) maxNode(board board.Board) bool{
  return board.NextMarker() == c.marker
}

func availableSpots(board board.Board) []int{
  return board.AvailableSpots()
}

const (
  computerThinking = "Minimax is now working hard. Think happy thoughts...\n"
)
