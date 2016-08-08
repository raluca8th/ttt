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

  if playerName := setUp.getPlayerName(); playerName != expectedMessage {
    t.Error("Expected player name to be 'Anda', but got", playerName)
  }
}

func TestValidNameSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("   Anda")
  setUp := Setup{Ui: testUI}
  expectedMessage := "Anda"

  if playerName := setUp.getPlayerName(); playerName != expectedMessage {
    t.Error("Expected player name to be 'Anda', but got", playerName)
  }
}

func TestGetGameType(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("5 7 3")
  setup := Setup{Ui: testUI}

  if gameType := setup.GetGameType(); gameType != "3" {
    t.Error("Expected game type to be 3, but got", gameType)
  }
}

func TestMarkerSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("A")
  setUp := Setup{Ui: testUI}
  expectedMessage := "A"

  if playerMarker := setUp.getPlayerMarker(); playerMarker != expectedMessage {
    t.Error("Expected marker to be 'A', but got", playerMarker)
  }
}

func TestValidMarkerSelection(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("    A")
  setUp := Setup{Ui: testUI}
  expectedMessage := "A"

  if playerMarker := setUp.getPlayerMarker(); playerMarker != expectedMessage {
    t.Error("Expected marker to be 'A', but got", playerMarker)
  }
}

func TestGameSize(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("8 1")
  setUp := Setup{Ui: testUI}
  expectedSize := 9

  if gameSize := setUp.GetGameSize(); gameSize != expectedSize {
    t.Error("Expected 9, but got", gameSize)
  }
}

func TestValidGameSize(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("5 0 g 2")
  setUp := Setup{Ui: testUI}
  expectedGameSize := 16

  if gameSize := setUp.GetGameSize(); gameSize != expectedGameSize {
    t.Error("Expected 2, but got", gameSize)
  }
}

func TestGameOrder(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda A Eli E 2")
  setup := Setup{Ui: testUI}
  players := setup.GeneratePlayers("2")
  player1 := players[0]

  if player1Name := player1.Name(); player1Name != "Eli " {
    t.Error("Expected player name to be Eli, but it was", player1Name)
  }

  if player1Marker := player1.Marker(); player1Marker != "E " {
    t.Error("Expected marker to be 'E', but got", player1Marker)
  }
}

func TestGenerateHumanPlayers(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda A Anda Eli A E 1")
  setup := Setup{Ui: testUI}
  players := setup.GeneratePlayers("1")
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

  if player2Marker := player2.Marker(); player2Marker != "E " {
    t.Error("Expected marker to be 'E', but got", player2Marker)
  }
}

func TestGenerateComputerPlayers(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Wallee W Wallee Eve W E 1")
  setup := Setup{Ui: testUI}
  players := setup.GeneratePlayers("3")
  player1 := players[0]
  player2 := players[1]

  if player1Name := player1.Name(); player1Name != "Wallee " {
    t.Error("Expected player name to be Wallee, but it was", player1Name)
  }

  if player1Marker := player1.Marker(); player1Marker != "W " {
    t.Error("Expected marker to be 'W', but got", player1Marker)
  }

  if player2Name := player2.Name(); player2Name != "Eve " {
    t.Error("Expected player name to be Eve, but it was", player2Name)
  }

  if player2Marker := player2.Marker(); player2Marker != "E " {
    t.Error("Expected marker to be 'E', but got", player2Marker)
  }
}

func TestGenerateHumanComputerPlayers(t *testing.T){
  testUI := TestUI{Input: new(testSTDIN), Output: new(testSTDOUT)}
  testUI.Populate("Anda A Anda Wallee A W 1")
  setup := Setup{Ui: testUI}
  players := setup.GeneratePlayers("2")
  player1 := players[0]
  player2 := players[1]

  if player1Name := player1.Name(); player1Name != "Anda " {
    t.Error("Expected player name to be Anda, but it was", player1Name)
  }

  if player1Marker := player1.Marker(); player1Marker != "A " {
    t.Error("Expected marker to be 'A', but got", player1Marker)
  }

  if player2Name := player2.Name(); player2Name != "Wallee " {
    t.Error("Expected player name to be Wallee, but it was", player2Name)
  }

  if player2Marker := player2.Marker(); player2Marker != "W " {
    t.Error("Expected marker to be 'W', but got", player2Marker)
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
