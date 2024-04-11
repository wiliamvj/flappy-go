package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/wiliamvj/flappy-go/assets"
)

type Player struct {
  image    *ebiten.Image
  position Vector
  velocity float64
}

func NewPlayer() *Player {
  image := assets.PlayerSprite
  return &Player{
    image:    image,
    position: Vector{X: 20, Y: screenHeight / 2},
    velocity: 0,
  }
}

func (p *Player) Update() {
  gravity := 0.5
  jumpVelocity := -6.0

  if ebiten.IsKeyPressed(ebiten.KeySpace) {
    p.velocity = jumpVelocity
  }

  p.velocity += gravity
  p.position.Y += p.velocity
}

func (p *Player) Draw(screen *ebiten.Image) {
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(p.position.X, p.position.Y)
  screen.DrawImage(p.image, op)
}
