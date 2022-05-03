package main

import (
	"encoding/json"
	"fmt"
)

var configRulesStr = `
{
	"name": "Blueprint Rules",
	"description": "Default rule configuration for Genesys Cloud Blueprints",
	"ruleGroups": {
		"CONTENT": {
			"description": "Content related validation",
			"rules": [{
					"description": "Overview image should be referred to in README.MD",
					"file": "./README.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "!\\[.*\\]\\(blueprint/images/overview\\.png *['|\"]*.*['|\"]*\\)"
						}]
					}],
					"level": "error"
				}, {
					"description": "The front matter must be defined in the file or the blueprint will not appear in the Developer Center",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "(?s)^---.*---"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md file's front matter must include the following fields: title, author, indextype, icon, image, category, and summary",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "title: *.*"
						}, {
							"type": "regex",
							"value": "author: *.*"
						}, {
							"type": "regex",
							"value": "indextype: *blueprint"
						}, {
							"type": "regex",
							"value": "icon: *blueprint"
						}, {
							"type": "regex",
							"value": "image: *.*"
						}, {
							"type": "regex",
							"value": "category: *.*"
						}, {
							"type": "regex",
							"value": "summary: *.*"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ## Scenario section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "## *Scenario *"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ## Solution section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "## *Solution *"
						}]
					}],
					"level": "error"
				},
				{
					"description": "The index.md must have a ## Prerequisites section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "## *Prerequisites *"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ### Specialized knowledge section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "### *Specialized knowledge *"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ## Implementation steps section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "## *Implementation steps *"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ### Download the repository containing the project files section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "### *Download the repository containing the project files *"
						}]
					}],
					"level": "error"
				}, {
					"description": "The index.md must have a ## Additional resources section describing the problem the blueprint is trying to solve.",
					"file": "./blueprint/index.md",
					"conditions": [{
						"contains": [{
							"type": "regex",
							"value": "## *Additional resources *"
						}]
					}],
					"level": "error"
				}
			]
		}
	}
}
`

type ContainsConditional struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Condition struct {
	Contains []ContainsConditional `json:"contains"`
}

type Rule struct {
	Description string      `json:"description"`
	File        string      `json:"file"`
	Conditions  []Condition `json:"conditions"`
	Level       string      `json:"level"`
}

type RuleSet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	RuleGroups  struct {
		Content struct {
			Description string `json:"description"`
			Rules       []Rule `json:"rules"`
		} `json:"CONTENT"`
	} `json:"ruleGroups"`
}

func LoadRuleSet() *RuleSet {
	configData := []byte(configRulesStr)

	ruleSet := &RuleSet{}
	err := json.Unmarshal(configData, ruleSet)

	if err != nil {
		fmt.Println(err)
	}

	return ruleSet
}
