package entrevista

import "regexp"

type AnswerType int

const (
	Text AnswerType = 1 + iota
	YesNo
	Number
)

type Question struct {
	// The text of the question. Required.
	Text string
	// The type of the expected response. Defaults to Text.
	ExpectedAnswerType AnswerType
	// Whether a response is required.
	Required bool
	// The default answer. Defaults to blank.
	DefaultAnswer string
	// The regular expression used to validate the answer. Defaults to blank.
	Validator *regexp.Regexp
	// The error message to display if the answer is required and not supplied.
	RequiredMessage string
	// The error message to display if the answer is invalid.
	InvalidMessage string
}
