package quizz

import (
	"reflect"
	"testing"
)

// TODO: use github.com/stretchr/testify if it gets too verbose

var (
	threeQuestions = []Question{
		{Desc: "2+1", Answer: "3"},
		{Desc: "0+0", Answer: "0"},
		{Desc: "30+75", Answer: "105"},
	}
	twoQuestions = []Question{
		{Desc: "5+5", Answer: "10"},
		{Desc: "What is the capital city of France ? ", Answer: "Paris"},
	}
	oneQuestion = []Question{
		{Desc: "Hello ?", Answer: "World"},
	}
	emptyQuizz = &quizz{questions: []Question{}}
)

func TestNewQuizz(t *testing.T) {
	type args struct {
		questions []Question
	}

	tests := []struct {
		name string
		args args
		want Quizz
	}{
		{
			"two questions quizz",
			args{
				questions: twoQuestions,
			},
			&quizz{questions: twoQuestions},
		},
		{
			"empty quizz",
			args{
				questions: nil,
			},
			emptyQuizz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuizz(tt.args.questions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuizz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizz_Check(t *testing.T) {
	oneQuestionQuizz := &quizz{questions: oneQuestion}
	twoQuestionsQuizz := &quizz{questions: twoQuestions}

	t.Run("wrong answer", func(t *testing.T) {
		if oneQuestionQuizz.Check(oneQuestion[0].Answer+"mlkjsbhdf") == true {
			t.Errorf("Check() = %t, want %t", true, false)
		}
	})

	t.Run("last answer, good question", func(t *testing.T) {
		if oneQuestionQuizz.Check(oneQuestion[0].Answer) == false {
			t.Errorf("Check() = %t, want %t", false, true)
		}

		if oneQuestionQuizz.Terminated() == false {
			t.Errorf("Terminated() = %t, want %t", false, true)

		}
	})

	t.Run("two questions", func(t *testing.T) {
		// First question
		if twoQuestionsQuizz.Check(twoQuestions[0].Answer) == false {
			t.Errorf("Check() = %t, want %t", true, false)
		}

		// Last question but wrong answer (current is 1, not 0 anymore), can't be terminated
		if twoQuestionsQuizz.Check(twoQuestions[0].Answer) == true {
			t.Errorf("Check() = %t, want %t", true, false)
		}

		if twoQuestionsQuizz.Terminated() == true {
			t.Errorf("Terminated() = %t, want %t", false, true)
		}

		// Last question, needs to be terminated
		if twoQuestionsQuizz.Check(twoQuestions[1].Answer) == false {
			t.Errorf("Check() = %t, want %t", false, true)
		}

		if twoQuestionsQuizz.Terminated() == false {
			t.Errorf("Terminated() = %t, want %t", true, false)
		}
	})
}

func Test_quizz_Current(t *testing.T) {
	type fields struct {
		questions  []Question
		current    int
		score      int
		timeout    int
		terminated bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "two questions",
			fields: fields{
				questions:  twoQuestions,
				timeout:    0,
				terminated: false,
				current:    1,
			},
			want: twoQuestions[1].Desc,
		},
		{
			name: "nil questions",
			fields: fields{
				questions:  nil,
				timeout:    0,
				terminated: false,
				current:    1,
			},
			want: "",
		},
		{
			name: "empty questions",
			fields: fields{
				questions:  []Question{},
				timeout:    0,
				terminated: false,
				current:    1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := quizz{
				questions:  tt.fields.questions,
				current:    tt.fields.current,
				score:      tt.fields.score,
				timeout:    tt.fields.timeout,
				terminated: tt.fields.terminated,
			}
			if got := q.Current(); got != tt.want {
				t.Errorf("Current() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_quizz_nextQuestion(t *testing.T) {
	oneQuestionQuizz := &quizz{questions: oneQuestion}
	twoQuestionsQuizz := &quizz{questions: twoQuestions}

	t.Run("last question", func(t *testing.T) {
		if oneQuestionQuizz.nextQuestion() == false {
			t.Errorf("nextQuestion() = %t, want %t", false, true)
		}

		if oneQuestionQuizz.Terminated() == false {
			t.Errorf("Terminated() = %t, want %t", false, true)
		}
	})

	t.Run("two questions", func(t *testing.T) {
		// First question
		if twoQuestionsQuizz.nextQuestion() == true {
			t.Errorf("nextQuestion() = %t, want %t", true, false)
		}

		// Last question, needs to be terminated
		if twoQuestionsQuizz.nextQuestion() == false {
			t.Errorf("nextQuestion() = %t, want %t", false, true)
		}

		if twoQuestionsQuizz.Terminated() == false {
			t.Errorf("Terminated() = %t, want %t", false, true)
		}
	})
}
