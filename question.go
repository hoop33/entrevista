package entrevista

import (
	"fmt"
	"reflect"
	"regexp"
)

type Question struct {
	// The text of the question. Required.
	Text string
	// The type of the expected answer.
	AnswerKind reflect.Kind
	// Whether an answer is required.
	Required bool
	// The default answer.
	DefaultAnswer string
	// The regular expression used to validate the answer.
	RegularExpression *regexp.Regexp
	// Whether to validate minimum and maximum
	ValidateMinMax bool
	// The minimum (length for a string, value for a number)
	Minimum int
	// The maximum (length for a string, value for a number)
	Maximum int
	// The error message to display if the answer is required and not supplied.
	RequiredMessage string
	// The error message to display if the answer is invalid.
	InvalidMessage string
}

var YesNoRegularExpression = regexp.MustCompile("^[YyNn]")

func NewQuestion(text string) *Question {
	return &Question{
		Text:       text,
		AnswerKind: reflect.String,
	}
}

func NewStringQuestion(text string, minLength int, maxLength int) *Question {
	return &Question{
		Text:           text,
		AnswerKind:     reflect.String,
		ValidateMinMax: true,
		Minimum:        minLength,
		Maximum:        maxLength,
		InvalidMessage: fmt.Sprintf("Your answer must have a length between %d and %d characters.", minLength, maxLength),
	}
}

func NewBoolQuestion(text string) *Question {
	return &Question{
		Text:              text,
		AnswerKind:        reflect.Bool,
		DefaultAnswer:     "no",
		RegularExpression: YesNoRegularExpression,
		InvalidMessage:    "You must answer yes or no.",
	}
}

func NewNumberQuestion(text string) *Question {
	return &Question{
		Text:           text,
		AnswerKind:     reflect.Int,
		ValidateMinMax: false,
	}
}

func NewNumberInRangeQuestion(text string, min int, max int) *Question {
	return &Question{
		Text:           text,
		AnswerKind:     reflect.Int,
		ValidateMinMax: true,
		Minimum:        min,
		Maximum:        max,
		InvalidMessage: fmt.Sprintf("You must enter a number between %d and %d.", min, max),
	}
}
