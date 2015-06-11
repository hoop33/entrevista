package entrevista_test

import (
	"reflect"
	"testing"

	"github.com/hoop33/entrevista"
	"github.com/stretchr/testify/assert"
)

func MirrorAnswer(question *entrevista.Question) (string, error) {
	return question.Text, nil
}

func QuietAnswer(question *entrevista.Question) (string, error) {
	return "", nil
}

func QuietOutput(message string) {
}

func QuietError(message string) {
}

func TestAnswersGetFilled(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = MirrorAnswer

	interview.Questions = []entrevista.Question{
		{
			Text: "One",
		},
		{
			Text: "Two",
		},
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], "One")
	assert.Equal(t, answers[1], "Two")
}

func TestDefaultAnswerIsReturnedForBlank(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = QuietAnswer

	interview.Questions = []entrevista.Question{
		{
			Text:          "One",
			DefaultAnswer: "First default",
		},
		{
			Text:          "Two",
			DefaultAnswer: "Second default",
		},
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], "First default")
	assert.Equal(t, answers[1], "Second default")
}

func TestBooleansAreReturnedForBooleans(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = MirrorAnswer

	interview.Questions = []entrevista.Question{
		{
			Text:       "Yes",
			AnswerKind: reflect.Bool,
		},
		{
			Text:       "No",
			AnswerKind: reflect.Bool,
		},
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], true)
	assert.Equal(t, answers[1], false)
}

func TestBooleanAnswersActAsBooleans(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = MirrorAnswer

	interview.Questions = []entrevista.Question{
		*entrevista.NewBoolQuestion("yes"),
		*entrevista.NewBoolQuestion("no"),
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], true)
	assert.Equal(t, answers[1], false)
}

func TestNumbersAreReturnedForNumbers(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = MirrorAnswer

	interview.Questions = []entrevista.Question{
		{
			Text:       "12345",
			AnswerKind: reflect.Int,
		},
		{
			Text:       "43",
			AnswerKind: reflect.Int,
		},
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], 12345)
	assert.Equal(t, answers[1], 43)
}

func TestMinAndMaxWorkForNumbers(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = MirrorAnswer

	interview.Questions = []entrevista.Question{
		{
			Text:       "12345",
			AnswerKind: reflect.Int,
			Maximum:    99999,
		},
		{
			Text:       "-43",
			AnswerKind: reflect.Int,
			Minimum:    -50,
		},
	}
	answers, err := interview.Run()
	assert.Equal(t, err, nil)
	assert.Equal(t, answers[0], 12345)
	assert.Equal(t, answers[1], -43)
}

func TestStringQuestionValidatesLength(t *testing.T) {
}
