package controller

import (
	"snakego/snake"

	"github.com/eiannone/keyboard"
)

func KeyHandler(event keyboard.KeyEvent, snake *snake.Snake) {
	if event.Err != nil {
		panic(event.Err)
	}

	// Up
	if event.Key == 65517 || event.Rune == 119 {
		snake.Up()
	} else if event.Key == 65516 || event.Rune == 115 {
		snake.Down()
	} else if event.Key == 65515 || event.Rune == 97 {
		snake.Left()
	} else if event.Key == 65514 || event.Rune == 100 {
		snake.Rigth()
	}

	if event.Key == keyboard.KeyEsc {
		snake.Kill()
	}
}
