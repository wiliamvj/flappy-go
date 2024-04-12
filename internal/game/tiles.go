package game

import (
  "image"

  "github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawTiles(screen *ebiten.Image) {
  const (
    nx           = ScreenWidth / tileSize
    ny           = ScreenHeight / tileSize
    pipeTileSrcX = 128
    pipeTileSrcY = 192
  )

  op := &ebiten.DrawImageOptions{}
  for i := -2; i < nx+1; i++ {
    // ground
    op.GeoM.Reset()
    op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
      float64((ny-1)*tileSize-floorMod(g.cameraY, tileSize)))
    screen.DrawImage(tilesImage.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)

    // pipe
    if tileY, ok := g.pipeAt(floorDiv(g.cameraX, tileSize) + i); ok {
      for j := 0; j < tileY; j++ {
        op.GeoM.Reset()
        op.GeoM.Scale(1, -1)
        op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
          float64(j*tileSize-floorMod(g.cameraY, tileSize)))
        op.GeoM.Translate(0, tileSize)
        var r image.Rectangle
        if j == tileY-1 {
          r = image.Rect(pipeTileSrcX, pipeTileSrcY, pipeTileSrcX+tileSize*2, pipeTileSrcY+tileSize)
        } else {
          r = image.Rect(pipeTileSrcX, pipeTileSrcY+tileSize, pipeTileSrcX+tileSize*2, pipeTileSrcY+tileSize*2)
        }
        screen.DrawImage(tilesImage.SubImage(r).(*ebiten.Image), op)
      }
      for j := tileY + pipeGapY; j < ScreenHeight/tileSize-1; j++ {
        op.GeoM.Reset()
        op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
          float64(j*tileSize-floorMod(g.cameraY, tileSize)))
        var r image.Rectangle
        if j == tileY+pipeGapY {
          r = image.Rect(pipeTileSrcX, pipeTileSrcY, pipeTileSrcX+pipeWidth, pipeTileSrcY+tileSize)
        } else {
          r = image.Rect(pipeTileSrcX, pipeTileSrcY+tileSize, pipeTileSrcX+pipeWidth, pipeTileSrcY+tileSize+tileSize)
        }
        screen.DrawImage(tilesImage.SubImage(r).(*ebiten.Image), op)
      }
    }
  }
}
