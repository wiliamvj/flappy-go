package main

import (
  "flag"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/wiliamvj/flappy-go/internal/game"
)

var flagCRT = flag.Bool("crt", false, "enable the CRT effect")

func main() {
  flag.Parse()
  ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
  ebiten.SetWindowTitle("Flappy Go")
  if err := ebiten.RunGame(game.NewGame(*flagCRT)); err != nil {
    panic(err)
  }
}
