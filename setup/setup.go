package setup

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

func (s *Setup) GetPlayerName() string{
  s.print(playerName)
  return s.getUserInput()
}

func (s *Setup) GetPlayerMarker() string{
  s.print(playerMarker)
  return s.getUserInput()
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

func (s *Setup) print(st string){
  s.Ui.Print(st)
}

func (s *Setup) getUserInput() string{
  return s.Ui.Read()
}

func validGameSize(size string) bool{
  return size == "1" || size == "2"
}

const (
  emptySelection = ""
  welcome = "Welcome to GO TicTacToe"
  playerName = "Please enter player name"
  playerMarker = "Please enter player marker"
  invalidSelection = "Invalid Selection"
  gameSizeSelection = "Please select game type\n\t 1. Select 1 for 3X3 board \n\t 2. Select 2 for 4X4 board"
)

