package game

func (g *Game) score() int {
  x := floorDiv(g.x16, 16) / tileSize
  if (x - pipeStartOffsetX) <= 0 {
    return 0
  }
  return floorDiv(x-pipeStartOffsetX, pipeIntervalX)
}
