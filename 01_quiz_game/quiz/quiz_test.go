package quiz

import (
	"os"
	"strings"
	"testing"
)

func TestQuizLoadQuestions(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		q := &Quiz{}

		data, err := os.Open("testdata/quiz.csv")
		if err != nil {
			t.Fatal(err)
		}

		err = q.LoadQuestions(data)

		if err != nil {
			t.Fatal(err)
		}

		if len(q.questions) != 13 {
			t.Fatalf("expected 13 questions, got %d questions", len(q.questions))
		}
	})
}

func TestGetInput(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		got, err := getInput(strings.NewReader("22\n"))
		if err != nil {
			t.Fatal(err)
		}

		if got != 22 {
			t.Fatalf("expected 22, got %d", got)
		}
	})
}
