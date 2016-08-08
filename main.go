package main

import (
  "github.com/raluca8th/ttt/setup"
  "github.com/raluca8th/ttt/cli"
  "github.com/raluca8th/ttt/game"
  "github.com/raluca8th/ttt/ui"
)

func main(){
  ui := new(cli.CLI)
  g := gameInitializer{ui: ui}
  g.playGame()
}

func size(selection string) int{
  if selection == "1"{
    return 9
  } else {
    return 16
  }
}

type gameInitializer struct{
  setup *setup.Setup
  game *game.Game
  ui ui.UI
}

func (g *gameInitializer) playGame(){
  g.setup = &setup.Setup{Ui: g.ui}
  g.setup.Welcome()
  sizeSelection := g.setup.GetGameSize()
  size := size(sizeSelection)
  gameType := g.setup.GetGameType()
  players := g.setup.GeneratePlayers(gameType)
  g.game = game.NewGame(players, size, g.ui)
  g.game.PlayGame()
}
