package main

import (
	"fmt"

	"github.com/Styzex/GoOfLife/config"
	"github.com/Styzex/GoOfLife/simulation"
	"github.com/Styzex/GoOfLife/utils"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetLevel(log.DebugLevel)

	seed, _ := utils.GenSeed()
	log.Debug("The seed was generated", "seed", seed)
	log.Debug("The seeds lenght is", "len", len(seed.String()))

	grid := utils.GenGrid(seed)

	for i := range config.GridWidth {
		for j := range config.GridHeight {
			value, err := simulation.CheckLife(i, j, grid)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(value, i, j)
		}
	}

	tick := config.Tick
	for {
		log.Debug("Time step", "tick", tick)

		grid = Update(grid)

		if tick >= 20000 {
			break
		}
		tick += 1
	}
}

func Update(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return grid
	}

	newGrid := make([][]bool, config.GridHeight)
	for i := range newGrid {
		newGrid[i] = make([]bool, config.GridWidth)
	}

	for y := range config.GridHeight {
		for x := range config.GridWidth {
			isAlive, err := simulation.CheckLife(x, y, grid)
			if err != nil {
				log.Error("Check life failed", "err", err)
				continue
			}
			newGrid[y][x] = isAlive
		}
	}

	return newGrid
}
