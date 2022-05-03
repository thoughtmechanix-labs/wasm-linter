package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func GetMarkdownText() js.Func {

	return js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		markdown := js.ValueOf(inputs[0].String())
		results := validateContent(markdown.String())
		failureBytes, err := json.Marshal(results.FailureResults)

		if err != nil {
			fmt.Printf("An error occurred while trying to deserialize an object %s\n", failureBytes)
			return ""
		}
		return string(failureBytes)
	})
}

func validateContent(markdown string) *ValidationResult {
	validationData := &ValidationData{
		ContentPath: markdown,
		RuleData:    LoadRuleSet(),
	}

	result, err := validationData.Validate()
	if err != nil {
		fmt.Printf("Error while trying to validate date: %s\n", err)
	}

	return result
}

func main() {
	holdch := make(chan struct{}, 0)

	js.Global().Set("getMarkdownText", GetMarkdownText())
	<-holdch
}
