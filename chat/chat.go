package chat

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"strings"
	"time"
)

const initialScreen = "commands"

type post struct {
	username string
	message  string
	time     string
}

var posts = []post{
	{username: "john", message: "hi, what's up?", time: "14:41"},
	{username: "jane", message: "not much", time: "14:43"},
}

type chatScreen struct {
	channelListSidebar *tui.Box
	historyBoxArea     *tui.Box
	currentInput       *tui.Entry
	currentChannelName string
	ui                 tui.UI
	userName           string
}

func NewChatScreen(ui tui.UI, userName string) *chatScreen {
	channelListSidebar := tui.NewVBox(
		tui.NewLabel("CHANNELS"),
		tui.NewLabel(initialScreen),
		tui.NewSpacer(),
	)
	channelListSidebar.SetBorder(true)
	historyBoxArea := tui.NewVBox()

	return &chatScreen{
		channelListSidebar: channelListSidebar,
		historyBoxArea:     historyBoxArea,
		ui:                 ui,
		userName:           userName,
	}

}

func (g *chatScreen) Render() *tui.Box {
	return g.renderChatFromChannelName(initialScreen)
}

func (g *chatScreen) renderChatFromChannelName(channelName string) *tui.Box {
	g.clearOldBufferedChatHistory()
	historyBox := g.renderChatHistoryBoxFromChannelName(channelName)
	inputBox := g.renderInputBox()
	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)
	root := tui.NewHBox(g.channelListSidebar, chat)
	return root
}

func (g *chatScreen) clearOldBufferedChatHistory() {
	for g.historyBoxArea.Length() > 0 {
		g.historyBoxArea.Remove(0)
	}
}

func (g *chatScreen) renderInputBox() *tui.Box {
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)
	input.OnSubmit(func(e *tui.Entry) {
		message := e.Text()
		if isCommand(message) {
			g.handlerCommands(message)
		} else {
			g.historyBoxArea.Append(tui.NewHBox(
				tui.NewLabel(time.Now().Format("15:04")),
				tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", g.userName))),
				tui.NewLabel(e.Text()),
				tui.NewSpacer(),
			))
		}

		input.SetText("")
	})
	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	return inputBox
}

func (g *chatScreen) renderChatHistoryBoxFromChannelName(channelName string) *tui.Box {
	historyScroll := tui.NewScrollArea(g.historyBoxArea)
	historyScroll.SetAutoscrollToBottom(true)
	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)
	historyBox.SetTitle(channelName)
	g.renderChatHistoryMessagesFromChannelName(channelName)
	return historyBox
}

func (g *chatScreen) renderChatHistoryMessagesFromChannelName(channelName string) {
	for _, m := range posts {
		g.historyBoxArea.Append(tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),
		))
	}
}

func (g *chatScreen) handlerCommands(command string) {
	fragmentedCommand := strings.Split(command, " ")
	commandName := fragmentedCommand[1]
	switch commandName {
	case "chn":
		g.doChangeChannel(fragmentedCommand[2:])
	}
}

func (g *chatScreen) doChangeChannel(args []string) {
	channelName := args[0]
	g.ui.SetWidget(g.renderChatFromChannelName(channelName))
	g.currentChannelName = channelName
}

func isCommand(msg string) bool {
	return strings.Contains(msg, "/e")
}
