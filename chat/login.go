package chat

import (
	"github.com/marcusolsson/tui-go"
)

const logo = `             
  _          _      _   _    ____   _   _      _      _____ 
 | |        / \    | \ | |  / ___| | | | |    / \    |_   _|
 | |       / _ \   |  \| | | |     | |_| |   / _ \     | |  
 | |___   / ___ \  | |\  | | |___  |  _  |  / ___ \    | |  
 |_____| /_/   \_\ |_| \_|  \____| |_| |_| /_/   \_\   |_|  
`

type loginScreen struct {
	rootWidgetBox *tui.Box
	user          string
	password      string
	authCallback  func(screen *loginScreen)
}

func NewLoginScreen(authCallback func(screen *loginScreen)) *loginScreen {
	return &loginScreen{
		authCallback: authCallback,
	}
}

func (s *loginScreen) Render() *tui.Box {
	user := tui.NewEntry()
	user.SetFocused(true)
	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)
	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	form.AppendRow(user, password)
	status := tui.NewStatusBar("Ready.")
	login := tui.NewButton("[Login]")
	login.OnActivated(func(b *tui.Button) {
		s.user = user.Text()
		s.password = password.Text()
		status.SetText("Logged in. with " + s.user + " " + s.password)
		s.authCallback(s)
	})
	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, login),
	)
	window := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("Welcome lan chat, fast and security internal chat.")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	window.SetBorder(true)
	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())
	root := tui.NewVBox(
		content,
		status,
	)
	tui.DefaultFocusChain.Set(user, password, login)
	s.rootWidgetBox = root
	return s.rootWidgetBox
}
