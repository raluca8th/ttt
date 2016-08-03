package game

type Player interface{
  Name() string
  Marker() string
}

type Game struct{
  players []Player
}

func (g Game) Players() []Player{
  return g.players
}
