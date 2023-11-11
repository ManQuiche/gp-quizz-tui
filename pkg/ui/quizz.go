package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gp-guizz-tui/pkg/quizz"
	"time"
)

var (
	resInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type keyMap struct {
	Answer key.Binding
	Exit   key.Binding
}

type UI struct {
	quizz    quizz.Quizz
	resInput textinput.Model
	timer    timer.Model

	keymap keyMap
	help   help.Model

	timeout     int
	lastCorrect *bool

	quitting bool
	timedOut bool
}

type TerminatedMsg struct {
	// timeout tells if the message originates from a timeout or a completion of the quizz
	timeout bool
}

func NewUI(q quizz.Quizz, timeout int) UI {
	ti := textinput.New()
	ti.Placeholder = "Your answer goes here..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	lipgloss.NewStyle()

	return UI{
		quizz:    q,
		timeout:  timeout,
		resInput: ti,
		timer:    timer.New(time.Duration(timeout) * time.Second),
		keymap: keyMap{
			Answer: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "Submit your answer"),
			),
			Exit: key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "Exit the quizz"),
			),
		},
		help: help.New(),
	}
}

func (ui UI) Init() tea.Cmd {
	ui.help.ShowAll = true
	return tea.Sequence(tea.EnterAltScreen, ui.resInput.Focus(), ui.timer.Init())
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case timer.TickMsg:
		ui.timer, cmd = ui.timer.Update(msg)
		return ui, cmd

	case timer.TimeoutMsg:
		return ui, ui.terminate(true)

	case TerminatedMsg:
		ui.quitting = true
		ui.timedOut = msg.timeout
		ui.help.ShowAll = false
		return ui, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if ui.lastCorrect == nil {
				ui.lastCorrect = new(bool)
			}

			*ui.lastCorrect = ui.quizz.Check(ui.resInput.Value())
			ui.resInput.SetValue("")

			if ui.quizz.Terminated() {
				return ui, ui.terminate(false)
			}

			return ui, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return ui, tea.Quit
		}
	}

	ui.resInput, cmd = ui.resInput.Update(msg)
	return ui, cmd
}

func (ui UI) View() string {
	if ui.quitting {
		if ui.timedOut {
			return ui.viewTimedOut()
		}
		return ui.viewDone()
	}

	return ui.viewPlaying()
}

func (ui UI) terminate(isTimeout bool) tea.Cmd {
	return func() tea.Msg {
		return TerminatedMsg{timeout: isTimeout}
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Exit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Answer},
		{k.Exit},
	}
}
