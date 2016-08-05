package setup

import (
  "testing"
  "bytes"
)

func TestWelcome(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  setUp := Setup{Ui: testUI}
  expectedMessage := "Welcome to GO TicTacToe\n"
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

func TestCreatePlayer(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda A")
  setup := Setup{Ui: testUI}
  player := setup.createPlayer()

  if playerName := player.Name(); playerName != "Anda " {
    t.Error("Expected player name to be Anda, but it was", playerName)
  }

  if playerMarker := player.Marker(); playerMarker != "A" {
    t.Error("Expected marker to be 'A', but got", playerMarker)
  }
}

func TestGeneratePlayers(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda A Anda Eli A E")
  setup := Setup{Ui: testUI}
  players := setup.GeneratePlayers()
  player1 := players[0]
  player2 := players[1]

  if player1Name := player1.Name(); player1Name != "Anda " {
    t.Error("Expected player name to be Anda, but it was", player1Name)
  }

  if player1Marker := player1.Marker(); player1Marker != "A " {
    t.Error("Expected marker to be 'A', but got", player1Marker)
  }

  if player2Name := player2.Name(); player2Name != "Eli " {
    t.Error("Expected player name to be Eli, but it was", player2Name)
  }

  if player2Marker := player2.Marker(); player2Marker != "E" {
    t.Error("Expected marker to be 'E', but got", player2Marker)
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

func (ui TestUI) Print(strings ...string){
  for _, s := range strings {
    ui.Output.print(s)
  }
}

func (ui *TestUI) CheckOutput() string{
  input, _ := ui.Output.buffer.ReadString('#')
  return input
}

func (ui *TestUI) Populate(s string){
   ui.Input.buffer.WriteString(s)
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
