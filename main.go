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

type gameInitializer struct{
  setup *setup.Setup
  game *game.Game
  ui ui.UI
}

func (g *gameInitializer) playGame(){
  g.setup = &setup.Setup{Ui: g.ui}
  g.setup.Welcome()
  sizeSelection := g.setup.GetGameSize()
  gameType := g.setup.GetGameType()
  players := g.setup.GeneratePlayers(gameType)
  g.game = game.NewGame(players, sizeSelection, g.ui)
  g.game.PlayGame()
}
