package chat

import (
	"context"
	"github.com/marcusolsson/tui-go"
	"log"
)

type Gui struct {
	screenContext   context.Context
	ui              tui.UI
	user            string
	password        string
	isAuthenticated bool
}

func NewGuiManager() *Gui {
	return &Gui{
		screenContext:   context.Background(),
		isAuthenticated: false,
	}

}

func (g *Gui) Render() {
	if g.isAuthenticated {
		panic("Error in auth on app bootstrap ")
	}
	ls := NewLoginScreen(g.authMiddleware)
	rootScreen := ls.Render()
	ui, err := tui.New(rootScreen)
	if err != nil {
		panic(err)
	}
	g.ui = ui
	g.ui.SetKeybinding("Esc", func() { ui.Quit() })
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func (g *Gui) authMiddleware(screen *loginScreen) {
	chatScreen := NewChatScreen(g.ui, screen.user)
	g.ui.SetWidget(chatScreen.Render())
}
