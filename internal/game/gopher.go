package game

import (
  "math"

  "github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawGopher(screen *ebiten.Image) {
  op := &ebiten.DrawImageOptions{}
  w, h := gopherImage.Bounds().Dx(), gopherImage.Bounds().Dy()
  op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
  op.GeoM.Rotate(float64(g.vy16) / 96.0 * math.Pi / 6)
  op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
  op.GeoM.Translate(float64(g.x16/16.0)-float64(g.cameraX), float64(g.y16/16.0)-float64(g.cameraY))
  op.Filter = ebiten.FilterLinear
  screen.DrawImage(gopherImage, op)
}
