package main

import (
  "testing"
  "reflect"
  "github.com/raluca8th/ttt/mocks"
)

func TestHumanVsHumanGame(t *testing.T){
  stdin := new(mocks.TestSTDIN)
  stdin.PopulateBuffer("1 1 Anda X Eli Y 1 0 3 1 4 2")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()
  finalBoardState := g.game.Board().Surface()
  expectedBoardState := []string{"X", "X", "X", "Y", "Y", "", "", "", ""}

  if !reflect.DeepEqual(finalBoardState, expectedBoardState) {
    t.Error("Expected X to win, but final board state was", finalBoardState)
  }
}

func TestHumanVsHumanGame4X4Game(t *testing.T){
  stdin := new(mocks.TestSTDIN)
  stdin.PopulateBuffer("2 1 Anda X Eli Y 1 0 3 5 4 10 7 15")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()
  finalBoardState := g.game.Board().Surface()
  expectedBoardState := []string{"X", "", "", "Y", "Y",  "X", "", "Y", "", "", "X", "", "", "", "", "X"}

  if !reflect.DeepEqual(finalBoardState, expectedBoardState) {
    t.Error("Expected X to win, but final board state was", finalBoardState)
  }
}

func TestHumanVsComputer4X4Game(t *testing.T){
  if testing.Short() {
    t.Skip("Skipping 4X4 test in short mode.Can take up to 5 minutes")
  }
  stdin := new(mocks.TestSTDIN)
  stdin.PopulateBuffer("2 2 Anda X Wallee W 1 0 1 2 3 4 5 6 7 8 9")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()
  humanMarker := "X"

  if g.game.Board().WinningMarker() == humanMarker {
    t.Error("Expected computer to win or tie, but winning marker was", humanMarker)
  }
}

func TestComputerVsComputer(t *testing.T){
  stdin := new(mocks.TestSTDIN)
  stdin.PopulateBuffer("1 3 Eve E Wallee W 1")
  stdout := new(mocks.TestSTDOUT)
  ui := mocks.TestUI{Input: stdin, Output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()

  if g.game.Board().IsTiedBoard() != true {
    t.Error("Expected computer vs. computer to always end in a tie")
  }
}
