package quizz

import "math/rand"

type Question struct {
	Desc   string
	Answer string
}

type quizz struct {
	questions  []Question
	current    int
	score      int
	timeout    int
	terminated bool
}

type Quizz interface {
	Stop()

	Terminated() bool
	Questions() []Question
	Current() string
	Score() int
	Check(answer string) bool
}

func Shuffle() func(*quizz) {
	return func(q *quizz) {
		rand.Shuffle(len(q.questions), func(i, j int) { q.questions[i], q.questions[j] = q.questions[j], q.questions[i] })
	}
}
