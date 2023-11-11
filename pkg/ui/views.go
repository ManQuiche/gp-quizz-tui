package ui

import (
	"fmt"
	"time"
)

func (ui UI) viewPlaying() string {
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
Score: %d, remaining time: %d s

%s
`, ui.quizz.Current(), resInputStyle.Render(ui.resInput.View()), correct, ui.quizz.Score(), ui.timer.Timeout/time.Second, ui.help.View(ui.keymap))
}

func (ui UI) viewDone() string {
	return fmt.Sprintf(`

Quizz done ! Score: %d

%s
`, ui.quizz.Score(), ui.help.View(ui.keymap))
}

func (ui UI) viewTimedOut() string {
	return fmt.Sprintf(`

Quizz timed out ! Try again ... Score: %d

%s
`, ui.quizz.Score(), ui.help.View(ui.keymap))
}
