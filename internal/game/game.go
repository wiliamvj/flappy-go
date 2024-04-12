package game

import (
  "fmt"
  "image/color"
  "math/rand"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/wiliamvj/flappy-go/assets"

  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
  ModeTitle Mode = iota
  ModeGame
  ModeGameOver
)

type Game struct {
  mode Mode

  // The gopher's position
  x16  int
  y16  int
  vy16 int

  // Camera
  cameraX int
  cameraY int

  // Pipes
  pipeTileYs []int

  gameoverCount int

  touchIDs   []ebiten.TouchID
  gamepadIDs []ebiten.GamepadID
}

func NewGame(c bool) ebiten.Game {
  g := &Game{}
  g.init()
  if c {
    return &GameWithCRTEffect{Game: g}
  }
  return g
}

func (g *Game) init() {
  g.x16 = 0
  g.y16 = 100 * 16
  g.cameraX = -240
  g.cameraY = 0
  g.pipeTileYs = make([]int, 256)
  for i := range g.pipeTileYs {
    g.pipeTileYs[i] = rand.Intn(6) + 2
  }
}

func (g *Game) isKeyJustPressed() bool {
  if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
    return true
  }
  if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
    return true
  }
  g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
  if len(g.touchIDs) > 0 {
    return true
  }
  g.gamepadIDs = ebiten.AppendGamepadIDs(g.gamepadIDs[:0])
  for _, g := range g.gamepadIDs {
    if ebiten.IsStandardGamepadLayoutAvailable(g) {
      if inpututil.IsStandardGamepadButtonJustPressed(g, ebiten.StandardGamepadButtonRightBottom) {
        return true
      }
      if inpututil.IsStandardGamepadButtonJustPressed(g, ebiten.StandardGamepadButtonRightRight) {
        return true
      }
    } else {
      // The button 0/1 might not be A/B buttons.
      if inpututil.IsGamepadButtonJustPressed(g, ebiten.GamepadButton0) {
        return true
      }
      if inpututil.IsGamepadButtonJustPressed(g, ebiten.GamepadButton1) {
        return true
      }
    }
  }
  return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
  return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
  switch g.mode {
  case ModeTitle:
    if g.isKeyJustPressed() {
      g.mode = ModeGame
    }
  case ModeGame:
    g.x16 += 32
    g.cameraX += 2
    if g.isKeyJustPressed() {
      g.vy16 = -96
    }
    g.y16 += g.vy16

    // Gravity
    g.vy16 += 4
    if g.vy16 > 96 {
      g.vy16 = 96
    }

    if g.hit() {
      g.mode = ModeGameOver
      g.gameoverCount = 30
    }
  case ModeGameOver:
    if g.gameoverCount > 0 {
      g.gameoverCount--
    }
    if g.gameoverCount == 0 && g.isKeyJustPressed() {
      g.init()
      g.mode = ModeTitle
    }
  }
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  backgroundImage := ebiten.NewImageFromImage(assets.BackgroundSprite)
  screenWidthDraw, screenHeightDraw := screen.Bounds().Max.X, screen.Bounds().Max.Y
  for y := 0; y < screenHeightDraw; y += backgroundImage.Bounds().Dy() {
    for x := 0; x < screenWidthDraw; x += backgroundImage.Bounds().Dx() {
      op := &ebiten.DrawImageOptions{}
      op.GeoM.Translate(float64(x), float64(y))
      screen.DrawImage(backgroundImage, op)
    }
  }
  g.drawTiles(screen)
  if g.mode != ModeTitle {
    g.drawGopher(screen)
  }

  var titleTexts string
  var texts string
  switch g.mode {
  case ModeTitle:
    titleTexts = "FLAPPY GOPHER"
    texts = "\n\n\n\n\n\nPRESS SPACE KEY\n\nOR A/B BUTTON\n\nOR TOUCH SCREEN"
  case ModeGameOver:
    texts = "\nGAME OVER!"
  }

  op := &text.DrawOptions{}
  op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
  op.ColorScale.ScaleWithColor(color.White)
  op.LineSpacing = titleFontSize
  op.PrimaryAlign = text.AlignCenter
  text.Draw(screen, titleTexts, &text.GoTextFace{
    Source: arcadeFaceSource,
    Size:   titleFontSize,
  }, op)

  op = &text.DrawOptions{}
  op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
  op.ColorScale.ScaleWithColor(color.White)
  op.LineSpacing = fontSize
  op.PrimaryAlign = text.AlignCenter
  text.Draw(screen, texts, &text.GoTextFace{
    Source: arcadeFaceSource,
    Size:   fontSize,
  }, op)

  if g.mode == ModeTitle {
    const msg = "Go Gopher by Renee French is\nlicenced under CC BY 3.0."

    op := &text.DrawOptions{}
    op.GeoM.Translate(ScreenWidth/2, ScreenHeight-smallFontSize/2)
    op.ColorScale.ScaleWithColor(color.White)
    op.LineSpacing = smallFontSize
    op.PrimaryAlign = text.AlignCenter
    op.SecondaryAlign = text.AlignEnd
    text.Draw(screen, msg, &text.GoTextFace{
      Source: arcadeFaceSource,
      Size:   smallFontSize,
    }, op)
  }

  op = &text.DrawOptions{}
  op.GeoM.Translate(ScreenWidth, 0)
  op.ColorScale.ScaleWithColor(color.White)
  op.LineSpacing = fontSize
  op.PrimaryAlign = text.AlignEnd
  text.Draw(screen, fmt.Sprintf("%04d", g.score()), &text.GoTextFace{
    Source: arcadeFaceSource,
    Size:   fontSize,
  }, op)

  ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) hit() bool {
  if g.mode != ModeGame {
    return false
  }
  const (
    gopherWidth  = 30
    gopherHeight = 60
  )
  w, h := gopherImage.Bounds().Dx(), gopherImage.Bounds().Dy()
  x0 := floorDiv(g.x16, 16) + (w-gopherWidth)/2
  y0 := floorDiv(g.y16, 16) + (h-gopherHeight)/2
  x1 := x0 + gopherWidth
  y1 := y0 + gopherHeight
  if y0 < -tileSize*4 {
    return true
  }
  if y1 >= ScreenHeight-tileSize {
    return true
  }
  xMin := floorDiv(x0-pipeWidth, tileSize)
  xMax := floorDiv(x0+gopherWidth, tileSize)
  for x := xMin; x <= xMax; x++ {
    y, ok := g.pipeAt(x)
    if !ok {
      continue
    }
    if x0 >= x*tileSize+pipeWidth {
      continue
    }
    if x1 < x*tileSize {
      continue
    }
    if y0 < y*tileSize {
      return true
    }
    if y1 >= (y+pipeGapY)*tileSize {
      return true
    }
  }
  return false
}

func (g *GameWithCRTEffect) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
  if g.crtShader == nil {
    s, err := ebiten.NewShader(crtGo)
    if err != nil {
      panic(fmt.Sprintf("flappy: failed to compiled the CRT shader: %v", err))
    }
    g.crtShader = s
  }

  os := offscreen.Bounds().Size()

  op := &ebiten.DrawRectShaderOptions{}
  op.Images[0] = offscreen
  op.GeoM = geoM
  screen.DrawRectShader(os.X, os.Y, g.crtShader, op)
}
