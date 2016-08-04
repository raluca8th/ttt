package setup

import (
  "testing"
  "bytes"
)

func TestWelcome(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  setUp := Setup{Ui: testUI}
  expectedMessage := "Welcome to GO TicTacToe"
  setUp.Welcome()

  if welcomeMessage := testUI.CheckOutput(); welcomeMessage != expectedMessage {
    t.Error("Expected welcome message, but got", welcomeMessage)
  }
}

func TestNameSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda")
  setUp := Setup{Ui: testUI}
  expectedMessage := "Anda"

  if playerName := setUp.GetPlayerName(); playerName != expectedMessage {
    t.Error("Expected player name to be 'Anda', but got", playerName)
  }
}

func TestValidNameSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("   Anda")
  setUp := Setup{Ui: testUI}
  expectedMessage := "Anda"

  if playerName := setUp.GetPlayerName(); playerName != expectedMessage {
    t.Error("Expected player name to be 'Anda', but got", playerName)
  }
}

func TestMarkerSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("A")
  setUp := Setup{Ui: testUI}
  expectedMessage := "A"

  if playerMarker := setUp.GetPlayerMarker(); playerMarker != expectedMessage {
    t.Error("Expected marker to be 'A', but got", playerMarker)
  }
}

func TestValidMarkerSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("    A")
  setUp := Setup{Ui: testUI}
  expectedMessage := "A"

  if playerMarker := setUp.GetPlayerMarker(); playerMarker != expectedMessage {
    t.Error("Expected marker to be 'A', but got", playerMarker)
  }
}

func TestGameSize(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("1")
  setUp := Setup{Ui: testUI}
  expectedMessage := "1"

  if gameSize := setUp.GetGameSize(); gameSize != expectedMessage {
    t.Error("Expected 1, but got", gameSize)
  }
}

func TestValidGameSize(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("5 0 g 2")
  setUp := Setup{Ui: testUI}
  expectedMessage := "2"

  if gameSize := setUp.GetGameSize(); gameSize != expectedMessage {
    t.Error("Expected 2, but got", gameSize)
  }
}

type testSTDIN struct {
  buffer bytes.Buffer
}

func (r *testSTDIN) read() string{
  input, _ := r.buffer.ReadString(' ')
  return input
}

type testSTDOUT struct {
  buffer bytes.Buffer
}

func (p *testSTDOUT) print(s string){
  p.buffer.WriteString(s)
}

type TestUI struct{
  Input *testSTDIN
  Output *testSTDOUT
}

func (ui TestUI) Read() string{
  return ui.Input.read()
}

func (ui TestUI) Print(s string){
  ui.Output.print(s)
}

func (ui *TestUI) CheckOutput() string{
  input, _ := ui.Output.buffer.ReadString('#')
  return input
}

func (ui *TestUI) Populate(s string){
   ui.Input.buffer.WriteString(s)
}
