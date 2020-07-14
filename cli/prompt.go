package main

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func promptType() (string, error) {
	prompt := promptui.Select{
		Label: "Type",
		Items: commitTypes,
	}
	_, result, err := prompt.Run()
	return result, err
}

func promptScope() (string, error) {
	prompt := promptui.Prompt{
		Label: "Scope",
	}
	return prompt.Run()
}

func promptTitle() (string, error) {
	prompt := promptui.Prompt{
		Label: "Title",
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("please enter message")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptBody() (string, error) {
	prompt := promptui.Prompt{
		Label: "Body",
	}
	return prompt.Run()
}

func promptBreakingChange() (string, error) {
	prompt := promptui.Prompt{
		Label: "💥 BREAKING CHANGE",
	}
	return prompt.Run()
}

func promptIsses() (string, error) {
	prompt := promptui.Prompt{
		Label: "✅ Closed issues",
	}
	return prompt.Run()
}

func promptSkipCi() (string, error) {
	prompt := promptui.Select{
		Label: "Skip CI",
		Items: []string{"No", "Yes"},
	}
	_, result, err := prompt.Run()
	return result, err
}

func collectType(state *CommitState) {
	commitType, _ := promptType()
	state.info.commitType = commitType
}

func collectScope(state *CommitState) {
	scope, _ := promptScope()
	state.info.scope = scope
}

func collectTitle(state *CommitState) {
	title, _ := promptTitle()
	state.info.title = title
}

func collectBody(state *CommitState) {
	body, _ := promptBody()
	state.info.body = body
}

func collectBreakingChange(state *CommitState) {
	breakingChange, _ := promptBreakingChange()
	state.info.breakingChange = breakingChange
}

func collectIssues(state *CommitState) {
	issues, _ := promptIsses()
	state.info.closes = issues
}

func collectSkipCi(state *CommitState) {
	skipCi, _ := promptSkipCi()
	state.info.skipCi = skipCi == "Yes"
}

// Collector is function that shows use a prompt and populates state
// with the result use entered.
type Collector func(state *CommitState)

var collectors map[string]Collector = map[string]Collector{
	"type":            collectType,
	"scope":           collectScope,
	"title":           collectTitle,
	"body":            collectBody,
	"breaking-change": collectBreakingChange,
	"issues":          collectIssues,
	"skip-ci":         collectSkipCi,
}

func collectUserInput(state *CommitState) {
	for _, key := range state.settings.prompts {
		if collector, ok := collectors[key]; ok {
			collector(state)
		}
	}
}
