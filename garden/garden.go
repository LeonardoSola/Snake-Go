package garden

import (
	"math/rand"
	"snakego/snake"
	"strconv"
	"strings"
	"time"
)

type Garden struct {
	X     int16
	Y     int16
	Fruit Fruit
}

type Fruit struct {
	X int16
	Y int16
}

func NewGarden() Garden {
	return Garden{
		X:     20,
		Y:     20,
		Fruit: Fruit{},
	}
}

func (garden *Garden) randomCord(max int16) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(int(max)-0) + 0
}

func (garden *Garden) GenerateFruit() {
	newFruit := Fruit{
		X: int16(garden.randomCord(garden.X)),
		Y: int16(garden.randomCord(garden.Y)),
	}

	garden.Fruit = newFruit
}

func (fruit *Fruit) CheckCollision(X int16, Y int16) bool {
	if fruit.X == X && fruit.Y == Y {
		return true
	}
	return false
}

func (garden *Garden) ToString(snake snake.Snake) string {
	var screen string = "\n"
	screen += strings.Repeat("█", int(garden.X)*2+6) + "\n"

	for y := 0; y <= int(garden.Y); y++ {
		linha := "██"
		for x := 0; x <= int(garden.X); x++ {
			if snake.CheckCollision(int16(x), int16(y)) {
				linha += "\033[4;32m" + "██" + "\033[0m"
			} else if garden.Fruit.CheckCollision(int16(x), int16(y)) {
				linha += "\033[0;31m" + "██" + "\033[0m"
			} else {
				linha += "░░"
			}
		}
		linha += "██\n"
		screen += linha
	}

	screen += strings.Repeat("█", int(garden.X)*2+6)
	screen += "\n"
	screen += "Pontos:" + strconv.Itoa(int(snake.Size))
	return screen
}
