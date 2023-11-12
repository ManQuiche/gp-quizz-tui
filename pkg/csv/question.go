package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gp-guizz-tui/pkg/quizz"
	"io"
	"os"
	"strings"
)

var (
	errIncorectQuestionFormat = errors.New("incorrect question format")
)

func ReadQuestions(filename string) ([]quizz.Question, error) {
	content, err := read(filename)

	if err != nil {
		return nil, fmt.Errorf("questions: %w", err)
	}

	r := csv.NewReader(strings.NewReader(content))
	questions := make([]quizz.Question, 0)
	for {
		question, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read: %w", err)
		}
		if len(question) != 2 {
			return nil, fmt.Errorf("could not read: %w", errIncorectQuestionFormat)
		}

		questions = append(questions, quizz.Question{Desc: question[0], Answer: question[1]})
	}

	return questions, nil
}

func read(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}

	return string(content), nil
}
