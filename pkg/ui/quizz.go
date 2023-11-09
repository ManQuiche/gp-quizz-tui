package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gp-guizz-tui/pkg/quizz"
)

var (
	resInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type UI struct {
	quizz    quizz.Quizz
	resInput textinput.Model

	lastCorrect *bool
	timeout     int

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
	}
}

func (ui UI) Init() tea.Cmd {
	return ui.resInput.Focus()
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case TerminatedMsg:
		ui.quitting = true
		ui.timedOut = msg.timeout
		return ui, tea.Quit

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
			return fmt.Sprintf("Quizz done ! Score: %d", ui.quizz.Score())
		}
		return fmt.Sprintf("Quizz timed out ! Try again ... Score: %d", ui.quizz.Score())
	}
	correct := ""
	if ui.lastCorrect != nil && *ui.lastCorrect {
		correct = "Correct answer !\n"
	} else if ui.lastCorrect != nil {
		correct = "Wrong ! Try again...\n"
	}

	return fmt.Sprintf(`
%s

%s

%s
Score: %d
(ctrl+c to quit)
`, ui.quizz.Current(), resInputStyle.Render(ui.resInput.View()), correct, ui.quizz.Score())
}

func (ui UI) terminate(isTimeout bool) tea.Cmd {
	return func() tea.Msg {
		return TerminatedMsg{timeout: isTimeout}
	}
}

//
//func (ui UI) Run() {
//	select {
//	case <-ui.askMany():
//		fmt.Printf("\n\nQuizz done ! Score: %d", ui.quizz.Score())
//	case <-time.After(time.Duration(ui.timeout) * time.Second):
//		fmt.Printf("\n\nQuizz timed out ! Try again ... Score: %d", ui.quizz.Score())
//	}
//
//	return
//}
//
//func (ui UI) askMany() <-chan bool {
//	done := make(chan bool)
//
//	go func() {
//		for !ui.quizz.Terminated() {
//			ui.ask()
//		}
//
//		done <- true
//	}()
//
//	return done
//}
//
//func (ui UI) ask() {
//	reader := bufio.NewReader(os.Stdin)
//
//	correct := false
//
//	for !correct {
//		fmt.Printf("%s:\n", ui.quizz.Current())
//
//		read, err := reader.ReadString('\n')
//		if err != nil {
//			log.Panic(err)
//		}
//
//		if ui.quizz.Check(read) {
//			fmt.Printf("Correct answer !\n\n")
//			correct = true
//		} else {
//			fmt.Printf("Wrong ! Try again...\n\n")
//		}
//	}
//}
