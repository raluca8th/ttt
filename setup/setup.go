package setup

import "github.com/raluca8th/ttt/players"

type setupUI interface{
  Read() string
  Print(s string)
}

type Setup struct{
  Ui setupUI
}

func (s *Setup) Welcome(){
  s.print(welcome)
}

func (s *Setup) GeneratePlayers() []*players.HumanPlayer{
  var playerArray []*players.HumanPlayer
  playerArray = append(playerArray, s.createPlayer(), s.createPlayer())
  return playerArray
}

func (s *Setup) createPlayer() *players.HumanPlayer{
  name := s.GetPlayerName()
  marker := s.GetPlayerMarker()
  return players.NewHumanPlayer(name, marker)
}

func (s *Setup) GetPlayerName() string{
  return s.validateInput(playerName)
}

func (s *Setup) GetPlayerMarker() string{
  return s.validateInput(playerMarker)
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
  return input != " "
}

const (
  emptySelection = ""
  welcome = "Welcome to GO TicTacToe"
  playerName = "Please enter player name"
  playerMarker = "Please enter player marker"
  invalidSelection = "Invalid Selection"
  gameSizeSelection = "Please select game type\n\t 1. Select 1 for 3X3 board \n\t 2. Select 2 for 4X4 board"
)

