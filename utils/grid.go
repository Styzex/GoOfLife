package utils

import (
	"fmt"

	"github.com/Styzex/GoOfLife/config"
)

func CountLiveCells(grid [][]bool) int {
	count := 0
	for y := 0; y < config.GridHeight; y++ {
		for x := 0; x < config.GridWidth; x++ {
			if grid[x][y] {
				count++
			}
		}
	}
	return count
}

func PrintCellMap(grid [][]bool) {
	for y := 0; y < config.GridHeight; y++ {
		for x := 0; x < config.GridWidth; x++ {
			if grid[x][y] {
				fmt.Print("■ ")
			} else {
				fmt.Print("· ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
