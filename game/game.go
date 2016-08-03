package game

import "github.com/raluca8th/ttt/board"

type Player interface{
  Name() string
  Marker() string
  SelectSpot(board *board.Board) int
}

type Game struct{
  players []Player
  board *board.Board
}

func NewGame(players []Player, size int) *Game{
  markers := getMarkersFromPlayers(players)
  board := board.NewBoard(board.Params{Size: size, Markers: markers})
  return &Game{players: players, board: board}
}

func (g Game) Players() []Player{
  return g.players
}

func (g Game) Board() *board.Board{
  return g.board
}

func (g Game) TakeTurn(player Player){
  spot := player.SelectSpot(g.Board())
  g.board.FillSpot(spot)
}

func getMarkersFromPlayers(players []Player) [2]string{
  var markers [2]string
  markers[0] = players[0].Marker()
  markers[1] = players[1].Marker()
  return markers
}
