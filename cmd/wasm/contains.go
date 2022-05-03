package main

import (
	"errors"
	"regexp"
	"strings"
)

type ContainsCondition struct {
	Path        string
	ContainsArr *[]ContainsConditional
}

func (condition *ContainsCondition) Validate() *ConditionResult {
	ret := &ConditionResult{
		FileHighlights: &[]FileHighlight{},
		IsSuccess:      true,
	}

	dataString := condition.Path

	for _, contains := range *condition.ContainsArr {
		if strings.TrimSpace(contains.Value) == "" {
			ret.Error = errors.New("value of contains is empty")
			ret.IsSuccess = false
			break
		}

		switch contains.Type {

		case "regex":
			re, err := regexp.Compile(contains.Value)
			if err != nil {
				ret.Error = err
				ret.IsSuccess = false
				break
			}

			loc := re.FindStringIndex(dataString)
			if loc == nil {
				ret.IsSuccess = false
				break
			}

			match := dataString[loc[0]:loc[1]]
			lineIndex := strings.Count(dataString[:loc[0]], "\n") + 1
			lineCount := strings.Count(dataString[loc[0]:loc[1]], "\n") + 1

			*ret.FileHighlights = append(*ret.FileHighlights, FileHighlight{
				Path:        "",
				LineNumber:  lineIndex,
				LineContent: strings.TrimSpace(match),
				LineCount:   lineCount,
			})
		default:
			ret.Error = errors.New("unknown contains type")
			ret.IsSuccess = false
		}
	}

	return ret
}
