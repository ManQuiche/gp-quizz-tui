package quizz

import (
	"reflect"
	"testing"
)

func TestNewQuizz(t *testing.T) {
	type args struct {
		questions []Question
	}

	questions := []Question{{Desc: "5+5", Answer: "10"}, {Desc: "8+2", Answer: "10"}}

	tests := []struct {
		name string
		args args
		want Quizz
	}{
		{
			"two questions quizz",
			args{
				questions: questions,
			},
			&quizz{questions: questions},
		},
		{
			"empty quizz",
			args{
				questions: nil,
			},
			&quizz{questions: nil},
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
