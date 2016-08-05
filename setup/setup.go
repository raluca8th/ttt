package setup

import (
  "ttt/game"
  "strings"
  "ttt/ui"
  "ttt/cli"
  "ttt/players"
)

type Setup struct{
  Ui ui.UI
  playersArray []game.Player
}

func (s *Setup) Welcome(){
  s.print(welcome)
}

func (s *Setup) GeneratePlayers() []game.Player{
  s.playersArray = append(s.playersArray, s.createPlayer())
  s.playersArray = append(s.playersArray, s.createPlayer())
  return s.playersArray
}

func (s *Setup) createPlayer() game.Player{
  name := s.GetPlayerName()
  marker := s.GetPlayerMarker()
  ui := cli.CLI{}
  return *players.NewHumanPlayer(name, marker, ui)
}

func (s *Setup) GetPlayerName() string{
  var name string
  for true {
    name := s.validateInput(playerName)
    playersArray := s.playersArray
    if len(playersArray) == 1  && playersArray[0].Name() == name{
      s.print(nameNotAvailable) } else {
      return name
    }
  }
  return name
}

func (s *Setup) GetPlayerMarker() string{
  var marker  string
  for true {
    marker := s.validateInput(playerMarker)
    playersArray := s.playersArray
    if len(playersArray) == 1  && playersArray[0].Marker() == marker{
      s.print(markerNotAvailable) } else {
      return marker
    }
  }
  return marker
}

func (s *Setup) GetGameSize() string{
  size := emptySelection
  s.print(gameSizeSelection)
  for true {
    size := s.getUserInput()
    if validGameSize(size) {
      return size
    }
    s.print(invalidSelection)
    s.print(gameSizeSelection)
  }
  return size
}

func (s *Setup) validateInput(message string) string{
  selection := emptySelection
  s.print(message)
  for true{
    selection := s.getUserInput()
    if s.validInput(selection){
      return selection
    }
    s.print(invalidSelection)
    s.print(message)
  }
  return selection
}

func (s *Setup) print(st string){
  s.Ui.Print(st)
}

func (s *Setup) getUserInput() string{
  return s.Ui.Read()
}

func validGameSize(size string) bool{
  return size == "1" || size == "2"
}

func (s *Setup) validInput(input string) bool{
  return strings.TrimSpace(input) != ""
}

const (
  emptySelection = ""
  welcome = "Welcome to GO TicTacToe\n"
  playerName = "Please enter player name\n"
  nameNotAvailable = "Name not available. Please enter another name\n"
  markerNotAvailable = "Marker not available. Please enter another name\n"
  playerMarker = "Please enter player marker\n"
  invalidSelection = "Invalid Selection\n"
  gameSizeSelection = "Please select game type\n\t 1. Select 1 for 3X3 board \n\t 2. Select 2 for 4X4 board\n"
)

