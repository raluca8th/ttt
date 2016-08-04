package setup

import (
  "github.com/raluca8th/ttt/players"
  "strings"
)

type setupUI interface{
  Read() string
  Print(s string)
}

type Setup struct{
  Ui setupUI
  playersArray []*players.HumanPlayer
}

func (s *Setup) Welcome(){
  s.print(welcome)
}

func (s *Setup) GeneratePlayers() []*players.HumanPlayer{
  s.playersArray = append(s.playersArray, s.createPlayer())
  s.playersArray = append(s.playersArray, s.createPlayer())
  return s.playersArray
}

func (s *Setup) createPlayer() *players.HumanPlayer{
  name := s.GetPlayerName()
  marker := s.GetPlayerMarker()
  return players.NewHumanPlayer(name, marker)
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
  welcome = "Welcome to GO TicTacToe"
  playerName = "Please enter player name"
  nameNotAvailable = "Name not available. Please enter another name"
  markerNotAvailable = "Marker not available. Please enter another name"
  playerMarker = "Please enter player marker"
  invalidSelection = "Invalid Selection"
  gameSizeSelection = "Please select game type\n\t 1. Select 1 for 3X3 board \n\t 2. Select 2 for 4X4 board"
)

