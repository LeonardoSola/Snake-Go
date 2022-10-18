package keylistener

import (
	"github.com/eiannone/keyboard"
)

var (
	KeysEvents <-chan keyboard.KeyEvent
)

func NewKeyLisener() <-chan keyboard.KeyEvent {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	return keysEvents
}
