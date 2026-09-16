package tetromino

import (
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

type Tetromino struct {
	Shape []Point
}

func ReadAndValidate(path string) ([]Tetromino, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rawStr := string(content)
	rawStr = strings.ReplaceAll(rawStr, "\r\n", "\n")

	rawStr = strings.TrimSpace(rawStr)
	if rawStr == "" {
		return nil, nil
	}

	chunks := strings.Split(rawStr, "\n\n")
	var tetros []Tetromino

	for _, chunk := range chunks {
		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			return nil, nil
		}

		lines := strings.Split(chunk, "\n")
		if len(lines) != 4 {
			return nil, nil
		}

		var blocks []Point
		for y := 0; y < 4; y++ {
			line := strings.TrimSpace(lines[y])
			if len(line) != 4 {
				return nil, nil
			}
			for x := 0; x < 4; x++ {
				if line[x] == '#' {
					blocks = append(blocks, Point{X: x, Y: y})
				} else if line[x] != '.' {
					return nil, nil
				}
			}
		}

		if len(blocks) != 4 || !isConnected(blocks) {
			return nil, nil
		}

		tetros = append(tetros, Tetromino{Shape: normalize(blocks)})
	}

	return tetros, nil
}

func isConnected(blocks []Point) bool {
	connections := 0
	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			dx := blocks[i].X - blocks[j].X
			dy := blocks[i].Y - blocks[j].Y
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {
				connections++
			}
		}
	}
	return connections == 3 || connections == 4
}

func normalize(blocks []Point) []Point {
	minX, minY := 4, 4
	for _, b := range blocks {
		if b.X < minX {
			minX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
	}
	normalized := make([]Point, len(blocks))
	for i, b := range blocks {
		normalized[i] = Point{X: b.X - minX, Y: b.Y - minY}
	}
	return normalized
}
