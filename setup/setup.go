package setup

type setupUI interface{
  Read() string
  Print(s string)
}

type Setup struct{
  Ui setupUI
}

func (s *Setup) Welcome(){
  s.print("Welcome to GO TicTacToe")
}

func (s *Setup) GetPlayerName() string{
  s.print("Please enter player name")
  input := s.read()
  return input
}

func (s *Setup) GetPlayerMarker() string{
  s.print("Please enter player marker")
  input := s.read()
  return input
}

func (s *Setup) print(st string){
  s.Ui.Print(st)
}

func (s *Setup) read() string{
  return s.Ui.Read()
}
