package quizz

import "strings"

func NewQuizz(questions []Question) Quizz {
	if questions == nil {
		questions = []Question{}
	}

	return &quizz{
		questions:  questions,
		score:      0,
		terminated: false,
		current:    0,
	}
}

func (q *quizz) Stop() {
	q.terminated = true
}

func (q quizz) Terminated() bool {
	return q.terminated
}

func (q quizz) Questions() []Question {
	return q.questions
}

func (q quizz) Current() string {
	if q.current < 0 || q.current >= len(q.questions) {
		return ""
	}
	return q.questions[q.current].Desc
}

func (q quizz) Score() int {
	return q.score
}

func (q *quizz) Check(answer string) bool {
	cleanAnswer := strings.ToLower(strings.Trim(strings.TrimSpace(answer), "\n"))

	if strings.ToLower(q.questions[q.current].Answer) == cleanAnswer {
		q.score += 1
		q.nextQuestion()
		return true
	}

	return false
}

func (q *quizz) nextQuestion() bool {
	if q.current == len(q.questions)-1 {
		q.Stop()
		return true
	}

	q.current += 1
	return false
}
