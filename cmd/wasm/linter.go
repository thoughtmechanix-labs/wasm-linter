package main

import (
	"errors"
	"fmt"
)

type ValidationData struct {
	RuleSetName string
	Description string
	ContentPath string
	RuleData    *RuleSet
}

type ValidationResult struct {
	SuccessResults *[]RuleResult `json:"success"`
	FailureResults *[]RuleResult `json:"failed"`
}

type RuleResult struct {
	Id             string           `json:"id"`
	Level          string           `json:"level"`
	Description    string           `json:"description"`
	IsSuccess      bool             `json:"-"`
	FileHighlights *[]FileHighlight `json:"fileHighlights,omitempty"`
	Error          *ValidationError `json:"error,omitempty"`
}

type ConditionResult struct {
	IsSuccess      bool
	FileHighlights *[]FileHighlight
	Error          error
}

type FileHighlight struct {
	Path        string `json:"path"`
	LineNumber  int    `json:"lineNumber"`
	LineCount   int    `json:"lineCount"`
	LineContent string `json:"lineContent"`
}

type ValidationError struct {
	RuleId string
	Err    error
}

type Validator interface {
	Validate() *ConditionResult
}

// Validate the content
func (input *ValidationData) Validate() (*ValidationResult, error) {
	markdown := input.ContentPath
	ruleData := input.RuleData
	finalResult := &ValidationResult{
		SuccessResults: &[]RuleResult{},
		FailureResults: &[]RuleResult{},
	}
	ch := make(chan *RuleResult)
	defer close(ch)

	rules := ruleData.RuleGroups.Content.Rules
	for id, rule := range rules {
		ruleIdFull := fmt.Sprintf("%s_%v", "ContentRules", id)
		ruleCpy := rule

		go func() {
			ch <- validateRule(&ruleCpy, ruleIdFull, markdown)
		}()
	}

	for i := 0; i < len(ruleData.RuleGroups.Content.Rules); i++ {
		ruleResult := <-ch

		if ruleResult.IsSuccess {
			*finalResult.SuccessResults = append(*finalResult.SuccessResults, *ruleResult)
		} else {
			*finalResult.FailureResults = append(*finalResult.FailureResults, *ruleResult)
		}
	}

	return finalResult, nil
}

// Evaluate the specific rule and get the RuleResult. Path is the root of
// content files
func validateRule(rule *Rule, ruleId string, markdown string) *RuleResult {
	ret := &RuleResult{
		Id:          ruleId,
		Level:       rule.Level,
		Description: rule.Description,
	}
	for _, condition := range rule.Conditions {

		condResult := validateCondition(&condition, markdown)

		if condResult == nil {
			ret.Error = &ValidationError{
				RuleId: ruleId,
				Err:    errors.New("unexpected error. No result from condition"),
			}
			break
		}

		ret.IsSuccess = condResult.IsSuccess
		ret.FileHighlights = condResult.FileHighlights

		if condResult.Error != nil {
			ret.Error = &ValidationError{
				RuleId: ruleId,
				Err:    condResult.Error,
			}
		}

		// Short circuit failing conditions
		if !ret.IsSuccess {
			break
		}
	}

	return ret
}

// Evaluate the condition. Any failure in any type of condition will short circuit the evaluation.
func validateCondition(condition *Condition, markdown string) *ConditionResult {
	var ret *ConditionResult
	var validator Validator

	contains := &condition.Contains
	// Contains Conditions

	if condition.Contains != nil {
		validator = &ContainsCondition{
			Path:        markdown,
			ContainsArr: contains,
		}
	}

	ret = validator.Validate()
	return ret
}
