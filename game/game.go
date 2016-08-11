package game

import (
  "github.com/raluca8th/ttt/tttboard"
  "github.com/raluca8th/ttt/board"
  "github.com/raluca8th/ttt/ui"
  "strconv"
  "strings"
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
  board := tttboard.NewBoard(size, markers)
  return &Game{players: players, board: board, ui: ui}
}

func (g *Game) Players() []Player{
  return g.players
}

func (g *Game) Board() board.Board{
  return g.board
}

func (g *Game) PlayGame(){
  g.printBoard(g.Board().Surface())
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
    stringSpot := strconv.Itoa(spot)
    g.ui.Print(strings.ToUpper(player.Name()), " selected spot ", stringSpot, "\n")
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
    g.ui.Print(strings.ToUpper(g.winner().Name()), congrats)
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
  g.ui.PrintBoard(board)
}

const (
  tie = "Game ended in a tie.\n"
  congrats = " won the game.\n"
)
