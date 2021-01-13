package gamemap

import (
	"candy/game/cell"
	"candy/game/cutter"
	"candy/game/direction"
	"candy/game/square"
)

var _ cutter.Range = (*candyRangeCutter)(nil)

type candyRangeCutter struct {
	maxRow int
	maxCol int
	grid   *[defaultRows][defaultCols]square.Square
}

func (c candyRangeCutter) CutRange(start cell.Cell, initialRange int, dir direction.Direction) int {
	for currRange := 1; currRange <= initialRange; currRange++ {
		nc := nextCell(start, currRange, dir)
		if !inGrid(nc, c.maxRow, c.maxCol) {
			return currRange - 1
		}
		sq := c.grid[nc.Row][nc.Col]
		if sq != nil && !sq.IsBreakable() {
			return currRange - 1
		}
	}
	return initialRange
}

func nextCell(start cell.Cell, offset int, dir direction.Direction) cell.Cell {
	switch dir {
	case direction.Up:
		return cell.Cell{
			Row: start.Row + offset,
			Col: start.Col,
		}
	case direction.Down:
		return cell.Cell{
			Row: start.Row - offset,
			Col: start.Col,
		}
	case direction.Left:
		return cell.Cell{
			Row: start.Row,
			Col: start.Col - offset,
		}
	case direction.Right:
		return cell.Cell{
			Row: start.Row,
			Col: start.Col + offset,
		}
	}
	return start
}