package game

import (
  "bytes"
  "log"
  "math/rand"
  "time"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
  "github.com/hajimehoshi/ebiten/v2/text/v2"
  "github.com/wiliamvj/flappy-go/assets"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

// init images
func init() {
  gopherImage = ebiten.NewImageFromImage(assets.PlayerSprite)
  tilesImage = ebiten.NewImageFromImage(assets.TilesSprite)
  backgroundImage = ebiten.NewImageFromImage(assets.BackgroundSprite)
}

// init fonts
func init() {
  s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
  if err != nil {
    log.Fatal(err)
  }
  arcadeFaceSource = s
}
