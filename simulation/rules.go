package simulation

import (
	"fmt"

	"github.com/Styzex/GoOfLife/config"
)

// Rewritten CheckLive function

func CheckLife(x, y int, grid [][]bool) (bool, error) {
	if len(grid) == 0 {
		return false, fmt.Errorf("grid is empty")
	}

	if x > config.GridWidth || x < 0 {
		err := fmt.Errorf("invalid width  %d, input width is bigger than the height of the grid", x)
		return false, err
	} else if y > config.GridHeight || y < 0 {
		err := fmt.Errorf("invalid height %d, input height is bigger than the height of the grid", y)
		return false, err
	}

	// I should look at the Go structs
	pos := []struct{ dx, dy int }{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	if grid[x][y] != true {
		return false, nil
	}

	ln := 0
	for _, i := range pos {
		nx, ny := x+i.dx, y+i.dy

		if nx < config.GridWidth && nx >= 0 && ny < config.GridHeight && ny >= 0 {
			if grid[ny][nx] {
				ln++
			}
		}
	}

	return ln >= 2, nil
}
