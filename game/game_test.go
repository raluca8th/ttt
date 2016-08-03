package game

import "testing"

func TestPlayers(t *testing.T){
  player1 := testPlayer{name: "Anda", marker: "A"}
  player2 := testPlayer{name: "Eli", marker: "E"}
  g := Game{players: []Player{player1, player2}}

  if player1Name := g.Players()[0].Name(); player1Name != "Anda"{
    t.Error("Expected name to be Anda, but it was", player1Name)
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
