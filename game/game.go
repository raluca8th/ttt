package game

import "github.com/raluca8th/ttt/board"

type Player interface{
  Name() string
  Marker() string
}

type Game struct{
  players []Player
  board *board.Board
}

func (g Game) Players() []Player{
  return g.players
}

func (g Game) Board() *board.Board{
  return g.board
}
