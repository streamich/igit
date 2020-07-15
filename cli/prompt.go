package main

import (
	"errors"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/manifoldco/promptui"
)

func promptType(state *CommitState) (string, error) {
	typesPresentation := []string{}
	for _, v := range state.settings.commitTypes {
		var emoji string
		if v, ok := state.settings.emojis[v]; ok {
			emoji = v
		}
		if emoji != "" {
			typesPresentation = append(typesPresentation, emoji+" "+v)
		} else {
			typesPresentation = append(typesPresentation, v)
		}
	}
	prompt := promptui.Select{
		Label: "Type",
		Items: typesPresentation,
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
	prompt := &survey.Editor{
		Message:  "Body",
		Default:  "Ctrl+C to skip",
		FileName: "*.md",
	}
	var result string
	err := survey.AskOne(prompt, &result)
	if err == terminal.InterruptErr {
		return "", nil
	}
	return result, err
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

func collectType(state *CommitState) error {
	commitType, err := promptType(state)
	state.info.commitType = commitType
	return err
}

func collectScope(state *CommitState) error {
	scope, err := promptScope()
	state.info.scope = scope
	return err
}

func collectTitle(state *CommitState) error {
	title, err := promptTitle()
	state.info.title = title
	return err
}

func collectBody(state *CommitState) error {
	body, err := promptBody()
	state.info.body = body
	return err
}

func collectBreakingChange(state *CommitState) error {
	breakingChange, err := promptBreakingChange()
	state.info.breakingChange = breakingChange
	return err
}

func collectIssues(state *CommitState) error {
	issues, err := promptIsses()
	state.info.closes = issues
	return err
}

func collectSkipCi(state *CommitState) error {
	skipCi, err := promptSkipCi()
	state.info.skipCi = skipCi == "Yes"
	return err
}

// Collector is function that shows use a prompt and populates state
// with the result use entered.
type Collector func(state *CommitState) error

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
			err := collector(state)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
		}
	}
}
