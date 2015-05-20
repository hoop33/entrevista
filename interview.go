package entrevista

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Interview struct {
	// The string to show at the end of questions. Default is ":".
	PromptTerminator string
	// The error message to display if the answer is required and not supplied.
	RequiredMessage string
	// The error message to display if the answer is invalid.
	InvalidMessage string
	// The function to use for normal output
	ShowOutput func(message string)
	// The function to use for error output
	ShowError func(message string)
	// The questions in the interview.
	Questions  []Question
	ReadAnswer func(question *Question) string
}

func showOutput(message string) {
	fmt.Print(message)
}

func showError(message string) {
	fmt.Print(message)
}

func (interview *Interview) displayPrompt(question *Question) {
	interview.ShowOutput(question.Text)
	if question.DefaultAnswer != "" {
		interview.ShowOutput(fmt.Sprintf(" (%s)", question.DefaultAnswer))
	}
	interview.ShowOutput(interview.PromptTerminator)
}

func isValid(text string, regex *regexp.Regexp) bool {
	if text == "" || regex == nil {
		return true
	}
	return regex.MatchString(text)
}

func readAnswer(question *Question) string {
	answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		// TODO something
	}
	// Strip off trailing newline
	return answer[0 : len(answer)-1]
}

func (interview *Interview) getAnswer(question *Question) string {
	for {
		interview.displayPrompt(question)
		answer := interview.ReadAnswer(question)

		// If they left answer blank and there's a default, set to default
		if answer == "" && question.DefaultAnswer != "" {
			answer = question.DefaultAnswer
		}

		// If it's still blank and it's required, show an error and loop back
		if answer == "" && question.Required {
			if question.RequiredMessage != "" {
				interview.ShowError(question.RequiredMessage)
			} else {
				interview.ShowError(interview.RequiredMessage)
			}
		} else if !isValid(answer, question.Validator) {
			// If answer isn't valid, show an error and loop back
			if question.InvalidMessage != "" {
				interview.ShowError(question.InvalidMessage)
			} else {
				interview.ShowError(interview.InvalidMessage)
			}
		} else {
			// We have a valid answer; return it
			return answer
		}
	}
}

func NewInterview() *Interview {
	return &Interview{
		PromptTerminator: ": ",
		RequiredMessage:  "You must provide an answer to this question.",
		InvalidMessage:   "Your answer is not valid.",
		ShowOutput:       showOutput,
		ShowError:        showError,
		ReadAnswer:       readAnswer,
	}
}

func (interview *Interview) Run() []string {
	answers := make([]string, len(interview.Questions))
	for index, question := range interview.Questions {
		if question.Text != "" {
			answer := interview.getAnswer(&question)
			answers[index] = answer
		}
	}
	return answers
}
