package setup

import (
  "github.com/raluca8th/ttt/game"
  "strings"
  "github.com/raluca8th/ttt/ui"
  "github.com/raluca8th/ttt/zen"
  "github.com/raluca8th/ttt/player"
  "github.com/raluca8th/ttt/computer"
)

type Setup struct{
  Ui ui.UI
  playersArray []game.Player
}

func (s *Setup) Welcome(){
  s.print(welcome)
}

func (s *Setup) GeneratePlayers(selection string) []game.Player{
  switch selection {
  case "1":
    s.playersArray = append(s.playersArray, s.createHumanPlayer())
    s.playersArray = append(s.playersArray, s.createHumanPlayer())
  case "2":
    s.playersArray = append(s.playersArray, s.createHumanPlayer())
    s.playersArray = append(s.playersArray, s.createComputerPlayer())
  case "3":
    s.playersArray = append(s.playersArray, s.createComputerPlayer())
    s.playersArray = append(s.playersArray, s.createComputerPlayer())
  default:
    s.playersArray = append(s.playersArray, s.createHumanPlayer())
    s.playersArray = append(s.playersArray, s.createHumanPlayer())
  }
    return s.getGameOrder(s.playersArray)
}

func (s *Setup) createHumanPlayer() game.Player{
  name := s.getPlayerName()
  marker := s.getPlayerMarker()
  return player.NewHumanPlayer(name, marker, s.Ui)
}

func (s *Setup) createComputerPlayer() game.Player{
  name := s.getPlayerName()
  marker := s.getPlayerMarker()
  return computer.NewComputerPlayer(name, marker, s.Ui)
}

func (s *Setup) getPlayerName() string{
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

func (s *Setup) getPlayerMarker() string{
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

func (s *Setup) GetGameSize() int{
  size := emptySelection
  for true {
    s.print(gameSizeSelection)
    size = s.getUserInput()
    if validGameSize(size) {
      break
    }
    s.print(invalidSelection)
  }
  return sizeFromSelection(size)
}

func sizeFromSelection(selection string) int{
  if selection == "1"{
    return 9
  } else {
    return 16
  }
}

func (s *Setup) GetGameType() string{
  selection := emptySelection
  for true {
    s.print(gameTypeSelection)
    selection := s.getUserInput()
    if validGameType(selection) {
      return selection
    }
    s.print(invalidSelection)
  }
  return selection
}

func (s *Setup) getGameOrder(players []game.Player) []game.Player{
  var orderedPlayers []game.Player
  selection := emptySelection
  for true {
    s.Ui.Print(gameOrderSelectionFirstPlayer, strings.ToUpper(players[0].Name()), gameOrderSelectionSecondPlayer, strings.ToUpper(players[1].Name()), "\n")
    selection = s.Ui.Read()
    if validGameOrder(strings.TrimSpace(selection)) {
      break
    }
    s.print(invalidSelection)
  }

  if selection == "1" {
    orderedPlayers = players
  } else if selection == "2" {
    orderedPlayers = append(orderedPlayers, players[1], players[0])
  }
  return orderedPlayers
}

func (s *Setup) validateInput(message string) string{
  selection := emptySelection
  s.print(message)
  for true{
    selection = s.getUserInput()
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
  return zen.IncludesString(size, []string{"1", "2"})
}

func validGameOrder(size string) bool{
  return zen.IncludesString(size, []string{"1", "2"})
}

func validGameType(size string) bool{
  return zen.IncludesString(size, []string{"1", "2", "3"})
}

func (s *Setup) validInput(input string) bool{
  return strings.TrimSpace(input) != emptySelection
}

const (
  emptySelection = ""
  welcome = "\nWelcome to GO TicTacToe with Minimax!\n\n"
  playerName = "Please enter player name\n"
  nameNotAvailable = "Name not available. Please enter another name\n"
  markerNotAvailable = "Marker not available. Please enter another name\n"
  playerMarker = "Please enter player marker\n"
  invalidSelection = "Invalid Selection\n"
  gameSizeSelection = "Please select board size\n\t 1. Select 1 for 3X3 board \n\t 2. Select 2 for 4X4 board\n"
  gameTypeSelection = "Please select game type\n\t 1. Select 1 for Human vs. Human \n\t 2. Select 2 for Human vs. Computer \n\t 3. Select 3 for Computer vs. Computer\n"
  gameOrderSelectionFirstPlayer = "Select '1' to start with player "
  gameOrderSelectionSecondPlayer = " or select '2' to start with player "
)
