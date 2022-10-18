package main

import (
	"fmt"
	"sync"
	"time"

	"snakego/controller"
	"snakego/garden"
	"snakego/keylistener"
	"snakego/snake"

	"github.com/gosuri/uilive"
)

type State struct {
	Running       bool
	Snake         *snake.Snake
	Garden        *garden.Garden
	ConsoleWriter *uilive.Writer
}

func (state *State) Render() {
	cobra := *state.Snake
	garden := *state.Garden
	fmt.Fprintf(state.ConsoleWriter, garden.ToString(cobra))
}

func main() {
	// Game State
	cobra := snake.NewSnake()
	jardin := garden.NewGarden()
	writer := uilive.New()
	game := State{
		Running:       true,
		Snake:         &cobra,
		Garden:        &jardin,
		ConsoleWriter: writer,
	}

	fmt.Print("\033[H\033[2J")
	jardin.GenerateFruit()

	fmt.Println(
		"██████████████████████████████████████████████████████████████████████████████████████\n" +
			"██████████████████████████████████████████████████████████████████████████████████████\n" +
			"█░░░░░░░░░░░░░░█░░░░░░██████████░░░░░░█░░░░░░░░░░░░░░█░░░░░░██░░░░░░░░█░░░░░░░░░░░░░░█\n" +
			"█░░\033[4;32m██████\033[0m  \033[0;31m██\033[0m░░█░░▄▀░░░░░░░░░░██░░▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀░░██░░▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█\n" +
			"█░░\033[4;32m██\033[0m░░░░░░░░░░█░░▄▀▄▀▄▀▄▀▄▀░░██░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░██░░▄▀░░░░█░░▄▀░░░░░░░░░░█\n" +
			"█░░\033[4;32m██\033[0m░░█████████░░▄▀░░░░░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░███░░▄▀░░█████████\n" +
			"█░░\033[4;32m██\033[0m░░░░░░░░░░█░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░░░░░▄▀░░███░░▄▀░░░░░░░░░░█\n" +
			"█░░\033[4;32m██████████\033[0m░░█░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░███░░▄▀▄▀▄▀▄▀▄▀░░█\n" +
			"█░░░░░░░░░░\033[4;32m██\033[0m░░█░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░░░░░▄▀░░███░░▄▀░░░░░░░░░░█\n" +
			"█████████░░\033[4;32m██\033[0m░░█░░▄▀░░██░░▄▀░░░░░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░███░░▄▀░░█████████\n" +
			"█░░░░░░░░░░\033[4;32m██\033[0m░░█░░▄▀░░██░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░░░█░░▄▀░░░░░░░░░░█\n" +
			"█░░\033[4;32m██████████\033[0m░░█░░▄▀░░██░░░░░░░░░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█\n" +
			"█░░░░░░░░░░░░░░█░░░░░░██████████░░░░░░█░░░░░░██░░░░░░█░░░░░░██░░░░░░░░█░░░░░░░░░░░░░░█\n" +
			"██████████████████████████████████████████████████████████████████████████████████████\n" +
			"███████████████████████████████Feito por Leonardo Sola████████████████████████████████\n" +
			"██████████████████████████████████████████████████████████████████████████████████████\n" +
			"	Controles:\n" +
			"		W ou ⬆ - Ir para cima\n" +
			"		D ou ➡️ - Ir para baixo\n" +
			"		A ou ⬅️ - Ir para esquerda\n" +
			"		S ou ⬇️ - Ir para direita\n" +
			"		ESC - Acabar com o jogo\n" +
			"	Precione enter para iniciar o jogo...")

	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	// Keyboard Handler
	go func() {
		defer waitGroup.Done()
		keyChanel := keylistener.NewKeyLisener()
		for game.Running {
			event := <-keyChanel
			controller.KeyHandler(event, &cobra)
		}
	}()

	// Game tick time
	go func() {
		defer waitGroup.Done()
		for game.Running {
			time.Sleep(100 * time.Millisecond)
			game.Snake.Move()

			if game.Snake.CheckCollision(jardin.Fruit.X, jardin.Fruit.Y) {
				game.Snake.Size++
				game.Garden.GenerateFruit()
			}

			if game.Snake.Death(game.Garden.X, game.Garden.Y) {
				game.Running = false
			}
		}
	}()

	// Game render
	go func() {
		defer waitGroup.Done()
		game.ConsoleWriter.Start()
		defer game.ConsoleWriter.Stop()
		for game.Running {
			time.Sleep(50 * time.Millisecond)
			game.Render()
		}
	}()

	// End
	waitGroup.Wait()
	fmt.Print("\033[H\033[2J")
}
