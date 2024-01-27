package quiz

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Quiz struct {
	questions []*Question
	r         io.Reader
	correct   int
	incorrect int
}

func (q *Quiz) LoadQuestions(r io.Reader) error {
	reader := csv.NewReader(r)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("quiz: %v", err)
		}

		if len(record) != 2 {
			return fmt.Errorf("quiz: invalid question input: %v", record)
		}

		question, err := NewQuestion(record[0], record[1])
		if err != nil {
			return err
		}

		q.questions = append(q.questions, question)
	}

	return nil
}

func (q *Quiz) Run(ctx context.Context) error {
	for _, question := range q.questions {
		// Print question
		fmt.Printf("%s", question.Name)

		userAnswer, err := getInput(q.r)
		if err != nil {
			fmt.Printf("error processing question: %v: skipping\n", err)
			continue
		}

		// Increment correct/incorrect counters
		if userAnswer == question.Answer {
			q.correct++
		} else {
			q.incorrect++
		}
	}

	return nil
}

type Question struct {
	Name   string
	Answer int
}

func NewQuestion(question, answer string) (*Question, error) {
	q := &Question{
		Name: question,
	}

	if q.Name == "" {
		return nil, fmt.Errorf("quiz: question cannot be empty")
	}

	aInt, err := strconv.Atoi(answer)
	if err != nil {
		return nil, fmt.Errorf("quiz: %v is not a valid number", answer)
	}

	q.Answer = aInt

	return q, nil
}

func (q *Question) IsCorrect(answer int) bool {
	return q.Answer == answer
}

func getInput(r io.Reader) (int, error) {
	var textInt int
	var err error

	// Get input from user
	fmt.Printf("Enter Answer: ")
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Check input against correct answer
		textInt, err = strconv.Atoi(scanner.Text())

		if err != nil {
			return -1, err
		}
	}

	if scanner.Err() != nil {
		return -1, scanner.Err()
	}

	return textInt, nil
}
