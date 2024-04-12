package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Mode int

var crtGo []byte

const (
  ScreenWidth      = 640
  ScreenHeight     = 480
  tileSize         = 32
  titleFontSize    = fontSize * 1.5
  fontSize         = 24
  smallFontSize    = fontSize / 2
  pipeWidth        = tileSize * 2
  pipeStartOffsetX = 8
  pipeIntervalX    = 8
  pipeGapY         = 5
)

var (
  gopherImage      *ebiten.Image
  tilesImage       *ebiten.Image
  backgroundImage  *ebiten.Image
  arcadeFaceSource *text.GoTextFaceSource
)

type GameWithCRTEffect struct {
  ebiten.Game
  crtShader *ebiten.Shader
}
