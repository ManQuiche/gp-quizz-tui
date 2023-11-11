package main

import (
	"flag"
	"gp-guizz-tui/pkg/csv"
	"gp-guizz-tui/pkg/quizz"
	"gp-guizz-tui/pkg/ui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	timeout := flag.Int("t", 30, "Quizz timeout in seconds.")
	flag.Parse()

	questions, err := csv.ReadQuestions("./problems_full.csv")
	if err != nil {
		log.Panic(err)
	}

	qz := quizz.NewQuizz(questions)
	qui := ui.NewUI(qz, *timeout)

	p := tea.NewProgram(qui, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
