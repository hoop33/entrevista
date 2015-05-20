package entrevista_test

import (
	"testing"

	"github.com/hoop33/entrevista"
	"github.com/stretchr/testify/assert"
)

func AnswerMirror(question *entrevista.Question) string {
	return question.Text
}

func QuietAnswer(question *entrevista.Question) string {
  return ""
}

func QuietOutput(message string) {
}

func QuietError(message string) {
}

func TestAnswersGetFilled(t *testing.T) {
	interview := entrevista.NewInterview()
	interview.ShowOutput = QuietOutput
	interview.ShowError = QuietError
	interview.ReadAnswer = AnswerMirror

	interview.Questions = []entrevista.Question{
		{
			Text: "One",
		},
		{
			Text: "Two",
		},
	}
	answers := interview.Run()
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
      Text: "One",
      DefaultAnswer: "First default",
    },
    {
      Text: "Two",
      DefaultAnswer: "Second default",
    },
  }
  answers := interview.Run()
  assert.Equal(t, answers[0], "First default")
  assert.Equal(t, answers[1], "Second default")
}
