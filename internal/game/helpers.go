package game

func floorDiv(x, y int) int {
  d := x / y
  if d*y == x || x >= 0 {
    return d
  }
  return d - 1
}

func floorMod(x, y int) int {
  return x - floorDiv(x, y)*y
}

func (g *Game) pipeAt(tileX int) (tileY int, ok bool) {
  if (tileX - pipeStartOffsetX) <= 0 {
    return 0, false
  }
  if floorMod(tileX-pipeStartOffsetX, pipeIntervalX) != 0 {
    return 0, false
  }
  idx := floorDiv(tileX-pipeStartOffsetX, pipeIntervalX)
  return g.pipeTileYs[idx%len(g.pipeTileYs)], true
}
