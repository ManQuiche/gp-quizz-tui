package quizz

import (
	"reflect"
	"testing"
)

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
	emptyQuizz = &quizz{questions: nil}
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
	type fields struct {
		questions  []Question
		current    int
		score      int
		timeout    int
		terminated bool
	}
	type args struct {
		answer string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "wrong answer",
			fields: fields{
				questions:  oneQuestion,
				timeout:    0,
				terminated: false,
				current:    0,
			},
			args: args{answer: oneQuestion[0].Answer + "hihi"},
			want: false,
		},
		{
			name: "right answer",
			fields: fields{
				questions:  threeQuestions,
				timeout:    0,
				terminated: false,
				current:    1,
			},
			args: args{answer: twoQuestions[1].Answer},
			want: true,
		},
		{
			name: "right answer, last question",
			fields: fields{
				questions:  twoQuestions,
				timeout:    0,
				terminated: false,
				current:    1,
			},
			args: args{answer: twoQuestions[1].Answer},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quizz{
				questions:  tt.fields.questions,
				current:    tt.fields.current,
				score:      tt.fields.score,
				timeout:    tt.fields.timeout,
				terminated: tt.fields.terminated,
			}
			if got := q.Check(tt.args.answer); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
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
				t.Errorf("Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizz_Questions(t *testing.T) {
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
		want   []Question
	}{
		// TODO: Add test cases.
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
			if got := q.Questions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Questions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizz_Score(t *testing.T) {
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
		want   int
	}{
		// TODO: Add test cases.
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
			if got := q.Score(); got != tt.want {
				t.Errorf("Score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizz_Stop(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quizz{
				questions:  tt.fields.questions,
				current:    tt.fields.current,
				score:      tt.fields.score,
				timeout:    tt.fields.timeout,
				terminated: tt.fields.terminated,
			}
			q.Stop()
		})
	}
}

func Test_quizz_Terminated(t *testing.T) {
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
		want   bool
	}{
		// TODO: Add test cases.
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
			if got := q.Terminated(); got != tt.want {
				t.Errorf("Terminated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizz_nextQuestion(t *testing.T) {
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
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quizz{
				questions:  tt.fields.questions,
				current:    tt.fields.current,
				score:      tt.fields.score,
				timeout:    tt.fields.timeout,
				terminated: tt.fields.terminated,
			}
			if got := q.nextQuestion(); got != tt.want {
				t.Errorf("nextQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
