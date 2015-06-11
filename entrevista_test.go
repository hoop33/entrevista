package entrevista_test

import (
	"fmt"

	"github.com/hoop33/entrevista"
)

func Example() {
	interview := entrevista.NewInterview()
	interview.Questions = []entrevista.Question{
		{
			Text:     "Enter your name",
			Required: true,
		},
		{
			Text:          "Enter your email address",
			DefaultAnswer: "john.doe@example.com",
		},
	}
	answers, err := interview.Run()
	if err != nil {
		for _, answer := range answers {
			fmt.Println(answer)
		}
	}
}
