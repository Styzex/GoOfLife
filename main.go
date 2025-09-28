package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

const width = 10
const height = 10

func main() {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	value, err := Check_live(1, 20, grid)
	if err != nil {
		log.Fatal(err)
	}

	print(value)
}

func Check_live(x int, y int, grid [][]bool) (bool, error) {
	if x > width {
		err := fmt.Errorf("invalid width  %d, input width is bigger than the height of the grid", x)
		return false, err
	} else if y > height {
		err := fmt.Errorf("invalid height %d, input height is bigger than the height of the grid", y)
		return false, err
	} else {
		if grid[x][y] {
			return true, nil
		}
		return false, nil
	}
}
