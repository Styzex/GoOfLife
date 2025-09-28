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

	if grid[y][x] != true {
		return false, nil
	}

	ln := 0
	for _, i := range pos {
		nx, ny := x+i.dx, x+i.dy

		if nx < config.GridWidth && nx >= 0 || ny < config.GridHeight && ny >= 0 {
			if grid[ny][nx] { // Again, grid[y][x] format
				ln++
			}
		}
	}

	return ln > 1, nil
}

// OLD CheckLive function
/*
func CheckLive(x int, y int, grid [][]bool) (bool, error) {
    if x > config.GridWidth {
        err := fmt.Errorf("invalid width  %d, input width is bigger than the height of the grid", x)
        return false, err
    } else if y > config.GridHeight {
        err := fmt.Errorf("invalid height %d, input height is bigger than the height of the grid", y)
        return false, err
    } else {
        if grid[x][y] && grid[x][y+1] {
            return true, nil
        } else if grid[x][y] && grid[x+1][y] {
            return true, nil
        } else if grid[x][y] && grid[x][y-1] {
            return true, nil
        } else if grid[x][y] && grid[x-1][y] {
            return true, nil
        }
        return false, nil
    }
}
*/
