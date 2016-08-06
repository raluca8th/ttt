package game

import (
  "math"
  "ttt/tttboard"
  "ttt/board"
  "ttt/ui"
  "strconv"
)

type Player interface{
  Name() string
  Marker() string
  SelectSpot(board board.Board) int
}

type Game struct{
  players []Player
  board board.Board
  ui ui.UI
}

func NewGame(players []Player, size int, ui ui.UI) *Game{
  markers := getMarkersFromPlayers(players)
  board := tttboard.NewBoard(tttboard.Params{Size: size, Markers: markers, UI: ui})
  return &Game{players: players, board: board, ui: ui}
}

func (g *Game) Players() []Player{
  return g.players
}

func (g *Game) Board() board.Board{
  return g.board
}

func (g *Game) PlayGame(){
  for true {
    for _, player := range g.Players() {
      g.takeTurn(player)
      g.printBoard(g.Board().Surface())
      if g.gameOver() {
        g.gameOverMessage()
        return
      }
    }
  }
}

func (g *Game) takeTurn(player Player){
  spot := player.SelectSpot(g.Board())
  if g.Board().SpotIsAvailable(spot) {
    g.board.FillSpot(spot)
  }
}

func (g *Game) winner() Player{
  var winner Player
  winnerMarker := g.Board().WinningMarker()
  for _, player := range g.Players() {
    if player.Marker() == winnerMarker {
      return player
    }
  }
  return winner
}

func (g *Game) gameOverMessage(){
  if g.winner() == nil {
    g.ui.Print(tie)
  } else {
    g.ui.Print(g.winner().Name(), congrats)
  }
}

func (g *Game) gameOver() bool{
  return g.Board().IsBoardSolved()
}

func getMarkersFromPlayers(players []Player) [2]string{
  var markers [2]string
  markers[0] = players[0].Marker()
  markers[1] = players[1].Marker()
  return markers
}

func (g *Game) printBoard(board []string){
  for i, value := range board{
    if value != "" {
      g.ui.Print("__", value, "_")
    } else {
      g.ui.Print("__", strconv.Itoa(i), "_")
    }

    if ((i + 1) % int(math.Sqrt(float64(len(board)))) == 0){
      g.ui.Print("\n\n")
    }
  }
}

const (
  tie = "Game ended in a tie\n"
  congrats = "! Congrats, you won the game!"
)
