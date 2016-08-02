package players

type HumanPlayer struct {
  name, marker string
}

func (h HumanPlayer) Name() string{
  return h.name
}

func (h HumanPlayer) Marker() string{
  return h.marker
}
