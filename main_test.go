package main

import "testing"
import "reflect"
import "bytes"
import "strings"

func TestHumanVsHumanGame(t *testing.T){
  stdin := new(testSTDIN)
  stdin.populateBuffer("1 1 Anda X Eli Y 1 4 2 3")
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()
  finalBoardState := g.game.Board().Surface()
  expectedBoardState := []string{"X", "X", "X", "Y", "Y", "", "", "", ""}

  if !reflect.DeepEqual(finalBoardState, expectedBoardState) {
    t.Error("Expected X to win, but final board state was", finalBoardState)
  }
}

func TestHumanVsComputer(t *testing.T){
  stdin := new(testSTDIN)
  stdin.populateBuffer("1 2 Anda X Wallee W 1 4 2 3 5 6")
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()
  humanMarker := "X"

  if g.game.Board().WinningMarker() == humanMarker {
    t.Error("Expected computer to win or tie, but winning marker was", humanMarker)
  }
}

func TestComputerVsComputer(t *testing.T){
  stdin := new(testSTDIN)
  stdin.populateBuffer("1 3 Eve E Wallee W")
  stdout := new(testSTDOUT)
  ui := testUI{input: stdin, output: stdout}
  g := &gameInitializer{ui: ui}
  g.playGame()

  if g.game.Board().IsTiedBoard() != true {
    t.Error("Expected computer vs. computer to always end in a tie")
  }
}

type testSTDIN struct {
  buffer bytes.Buffer
}

func (r *testSTDIN) Read() string{
  input, _ := r.buffer.ReadString(' ')
  return strings.TrimSpace(input)
}

func (p *testSTDIN) populateBuffer(s string){
  p.buffer.WriteString(s)
}

type testSTDOUT struct {
  buffer bytes.Buffer
}

func (p testSTDOUT) Print(s string) {
  p.buffer.WriteString(s)
}

type testUI struct{
  input *testSTDIN
  output *testSTDOUT
}

func (ui testUI) Read() string{
  return ui.input.Read()
}

func (ui testUI) Print(out ...string) {
  for _, s := range out {
    ui.output.Print(s)
  }
}
