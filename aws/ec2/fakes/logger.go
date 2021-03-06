package fakes

import "fmt"

type Logger struct {
	PrintfCall struct {
		CallCount int
		Receives  struct {
			Message   string
			Arguments []interface{}
		}
		Messages []string
	}

	PromptWithDetailsCall struct {
		CallCount int
		Receives  struct {
			Type string
			Name string
		}
		Returns struct {
			Proceed bool
		}
	}
}

func (l *Logger) Printf(message string, a ...interface{}) {
	l.PrintfCall.CallCount++
	l.PrintfCall.Receives.Message = message
	l.PrintfCall.Receives.Arguments = a

	l.PrintfCall.Messages = append(l.PrintfCall.Messages, fmt.Sprintf(message, a...))
}

func (l *Logger) PromptWithDetails(resourceType, resourceName string) bool {
	l.PromptWithDetailsCall.CallCount++
	l.PromptWithDetailsCall.Receives.Type = resourceType
	l.PromptWithDetailsCall.Receives.Name = resourceName

	return l.PromptWithDetailsCall.Returns.Proceed
}
