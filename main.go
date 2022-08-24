package main

import "github.com/mrlucca/lan-message/chat"

func main() {
	ls := chat.NewGuiManager()
	ls.Render()
}
