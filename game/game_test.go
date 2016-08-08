package game

import (
  "testing"
  "github.com/raluca8th/ttt/tttboard"
  "ttt/board"
  "strings"
  "strconv"
  "bytes"
  "reflect"
)

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
  board := tttboard.NewBoard(tttboard.Params{})
  g := Game{players: []Player{player1, player2}, board: board}

  if boardSize := g.Board().Size(); boardSize != 9{
    t.Error("Expected board size to be 9, but it was", boardSize)
  }
}

func TestNewGameGeneratesBoard(t *testing.T){
  player1 := testPlayer{name: "Anda", marker: "A"}
  player2 := testPlayer{name: "Eli", marker: "E"}
  players := []Player{player1, player2}
  ui := new(testUI)
  g := NewGame(players, 9, ui)

  if boardMarkers := g.Board().Markers(); boardMarkers != [2]string{"A", "E"}{
    t.Error("Expected markers to be 'A' and 'E', but they were", boardMarkers)
  }
}

func TestTakeTurnCurrentPlayer(t *testing.T){
  var mockInput bytes.Buffer
  mockInput.WriteString("4")
  player1 := testPlayer{name: "Anda", marker: "A", mockInput: &mockInput}
  player2 := testPlayer{name: "Eli", marker: "E", mockInput: &mockInput}
  players := []Player{player1, player2}
  ui := new(testUI)
  g := NewGame(players, 9, ui)
  g.takeTurn(player1)

  if spotIsAvailable := g.Board().SpotIsAvailable(4); spotIsAvailable != false{
    t.Error("Expected spot available to return false, but it returned", spotIsAvailable)
  }
}

func TestPlayGameUntilThereIsATie(t *testing.T){
  var mockInput bytes.Buffer
  mockInput.WriteString("0 1 2 5 3 6 4 8 7")
  player1 := testPlayer{name: "Anda", marker: "X", mockInput: &mockInput}
  player2 := testPlayer{name: "Eli", marker: "Y", mockInput: &mockInput}
  players := []Player{player1, player2}
  ui := new(testUI)
  g := NewGame(players, 9, ui)
  expectedBoardState := []string{"X", "Y", "X",
                                 "X", "X", "Y",
                                 "Y", "X", "Y"}
  g.PlayGame()

  if boardSurface := g.Board().Surface(); reflect.DeepEqual(boardSurface, expectedBoardState) != true{
    t.Error("Expected board to be tied, but it was", boardSurface)
  }
}

func TestPlayGameUntilGameIsWon(t *testing.T){
  var mockInput bytes.Buffer
  mockInput.WriteString("0 1 4 5 8 2 3 6 7")
  player1 := testPlayer{name: "Anda", marker: "X", mockInput: &mockInput}
  player2 := testPlayer{name: "Eli", marker: "Y", mockInput: &mockInput}
  players := []Player{player1, player2}
  ui := new(testUI)
  g := NewGame(players, 9, ui)
  expectedBoardState := []string{"X", "Y", "",
                                 "",  "X", "Y",
                                 "",   "", "X"}
  g.PlayGame()

  if boardSurface := g.Board().Surface(); reflect.DeepEqual(boardSurface, expectedBoardState) != true{
    t.Error("Expected board to be tied, but it was", boardSurface)
  }
}

func TestGameWinner(t *testing.T){
  var mockInput bytes.Buffer
  mockInput.WriteString("0 3 1 4 2")
  player1 := testPlayer{name: "Anda", marker: "X", mockInput: &mockInput}
  player2 := testPlayer{name: "Eli", marker: "Y", mockInput: &mockInput}
  players := []Player{player1, player2}
  ui := new(testUI)
  g := NewGame(players, 9, ui)
  g.PlayGame()

  if winnerName := g.winner().Name(); winnerName != "Anda"{
    t.Error("Expected winner name to be Anda, but it was", winnerName)
  }
}
type testPlayer struct{
  name, marker string
  mockInput *bytes.Buffer
}

func (p testPlayer) Name() string{
  return p.name
}

func (p testPlayer) Marker() string{
  return p.marker
}

func (p testPlayer) SelectSpot(board board.Board) int{
  input, _ := p.mockInput.ReadString(' ')
  spot, _ := strconv.Atoi(strings.TrimSpace(input))
  return spot
}

type testUI struct{
}

func (t testUI) Read() string {return ""}
func (t testUI) Print(s ...string){}
