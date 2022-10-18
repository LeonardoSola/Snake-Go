package snake

type Snake struct {
	Size          uint16
	Cords         [][2]int16
	Direction     string
	LastDirection string
}

func NewSnake() Snake {
	return Snake{
		Direction:     "R",
		LastDirection: "R",
		Size:          3,
		Cords:         [][2]int16{{3, 0}, {2, 0}, {1, 0}}, // First value X, Secound value Y
	}
}

func (snake *Snake) Up() {
	snake.Direction = "U"
}

func (snake *Snake) Down() {
	snake.Direction = "D"
}

func (snake *Snake) Left() {
	snake.Direction = "L"
}

func (snake *Snake) Rigth() {
	snake.Direction = "R"
}

func (snake *Snake) Kill() {
	cords := snake.Cords
	cords = append(cords, [2]int16{-1, -1})
	cords = append(cords, [2]int16{-1, -1})
	snake.Cords = cords
}

func (snake *Snake) CheckCollision(x int16, y int16) bool {
	for _, corpo := range snake.Cords {
		if corpo[0] == x && corpo[1] == y {
			return true
		}
	}
	return false
}

func (snake *Snake) Move() {
	lastPlace := snake.Cords[0]
	nextPlace := [2]int16{lastPlace[0], lastPlace[1]}

	if snake.LastDirection == "U" && snake.Direction == "D" {
		snake.Direction = snake.LastDirection
	} else if snake.LastDirection == "D" && snake.Direction == "U" {
		snake.Direction = snake.LastDirection
	} else if snake.LastDirection == "L" && snake.Direction == "R" {
		snake.Direction = snake.LastDirection
	} else if snake.LastDirection == "R" && snake.Direction == "L" {
		snake.Direction = snake.LastDirection
	}

	if snake.Direction == "U" {
		nextPlace[1] -= 1
	} else if snake.Direction == "D" {
		nextPlace[1] += 1
	} else if snake.Direction == "L" {
		nextPlace[0] -= 1
	} else if snake.Direction == "R" {
		nextPlace[0] += 1
	}

	snake.Cords = append([][2]int16{nextPlace}, snake.Cords...)

	length := len(snake.Cords)
	if length > int(snake.Size) {
		snake.Cords = (snake.Cords)[:length-1]
	}

	snake.LastDirection = snake.Direction
}
func (snake *Snake) Duplicated() bool {
	limpo := [][2]int16{}

	for _, peca := range snake.Cords {
		if contains(limpo, peca) {
			return true
		} else {
			limpo = append(limpo, peca)
		}
	}

	return false
}

func contains(array [][2]int16, value [2]int16) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

func (snake *Snake) WallCollision(x int16, y int16) bool {
	for _, peca := range snake.Cords {
		if peca[0] < 0 || peca[1] < 0 || peca[0] > x || peca[1] > y {
			return true
		}
	}
	return false
}

func (snake *Snake) Death(x int16, y int16) bool {
	if snake.Duplicated() {
		return true
	} else if snake.WallCollision(x, y) {
		return true
	}
	return false
}
