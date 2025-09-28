package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/Styzex/GoOfLife/config"
)

func GenSeed() (*big.Int, error) {
	min := new(big.Int)
	min.SetString(config.MinSeedValue, 10)
	max := new(big.Int)
	max.SetString(config.MaxSeedValue, 10)

	offset, err := rand.Int(rand.Reader, new(big.Int).Sub(max, min))

	seed := new(big.Int).Add(min, offset)
	return seed, err
}

func GenGrid(seed *big.Int) [][]bool {
	grid := make([][]bool, config.GridHeight)
	for y := 0; y < config.GridHeight; y++ {
		grid[y] = make([]bool, config.GridWidth)
	}

	seedStr := seed.String()
	seedLen := len(seedStr)

	for y := 0; y < config.GridHeight; y++ {
		for x := 0; x < config.GridWidth; x++ {
			idx := (y*config.GridWidth + x) % seedLen

			c, _ := strconv.Atoi(string(seedStr[idx]))
			n, _ := rand.Int(rand.Reader, big.NewInt(9))

			grid[y][x] = c >= int(n.Int64())
		}
	}

	return grid
}

// Old GenGrid
/*
func GenGrid(seed *big.Int) [][]bool {
	grid := make([][]bool, config.GridWidth)
	for i := range grid {
		grid[i] = make([]bool, config.GridWidth)
	}

	for i := range config.GridWidth {
		for j := range config.GridHeight {
			c, _ := strconv.Atoi(string(seed.String()[j]))
			n, _ := rand.Int(rand.Reader, big.NewInt(9))

			if c >= int(n.Int64()) {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	return grid
}
*/
