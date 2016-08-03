package game

import "testing"
import "github.com/raluca8th/ttt/board"

func TestPlayers(t *testing.T){
  player1 := testPlayer{name: "Anda", marker: "A"}
  player2 := testPlayer{name: "Eli", marker: "E"}
  g := Game{players: []Player{player1, player2}}

  if player1Name := g.Players()[0].Name(); player1Name != "Anda"{
    t.Error("Expected name to be Anda, but it was", player1Name)
  }
}

func TestBoard(t *testing.T){
  player1 := testPlayer{name: "Anda", marker: "A"}
  player2 := testPlayer{name: "Eli", marker: "E"}
  board := board.NewBoard(board.Params{})
  g := Game{players: []Player{player1, player2}, board: board}

  if boardSize := g.Board().Size(); boardSize != 9{
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}

func TestNewGameGeneratesBoard(t *testing.T){
  player1 := testPlayer{name: "Anda", marker: "A"}
  player2 := testPlayer{name: "Eli", marker: "E"}
  players := []Player{player1, player2}
  g := NewGame(players, 9)

  if boardMarkers := g.Board().Markers(); boardMarkers != [2]string{"A", "E"}{
    t.Error("Expected markers to be 'A' and 'E', but they were", boardMarkers)
  }
}

type testPlayer struct{
  name, marker string
}

func (p testPlayer) Name() string{
  return p.name
}

func (p testPlayer) Marker() string{
  return p.marker
}
