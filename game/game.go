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
  if g.Board().SpotIsAvailable(spot) {
    g.board.FillSpot(spot)
  }
}

func (g Game) PlayGame(){
  gameOngoing := true
  for gameOngoing {
    for _, player := range g.Players() {
      g.TakeTurn(player)
    }
    gameOngoing = !g.gameOver()
  }
}

func (g Game) gameOver() bool{
  return g.Board().IsTiedBoard()
}

func getMarkersFromPlayers(players []Player) [2]string{
  var markers [2]string
  markers[0] = players[0].Marker()
  markers[1] = players[1].Marker()
  return markers
}
