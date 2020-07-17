package cmdbranch

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
)

func promptType(state *branchState) (string, error) {
	prompt := promptui.Select{
		Label: "Type",
		Items: state.settings.branchTypes,
	}
	index, _, err := prompt.Run()
	return state.settings.branchTypes[index], err
}

func collectType(state *branchState) error {
	result, err := promptType(state)
	state.info.branchType = result
	return err
}

func promptName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Name",
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("please enter name")
			}
			return nil
		},
	}
	return prompt.Run()
}

func collectName(state *branchState) error {
	result, err := promptName()
	state.info.name = result
	return err
}

// Collector is function that shows use a prompt and populates state
// with the result use entered.
type Collector func(state *branchState) error

var collectors map[string]Collector = map[string]Collector{
	"type": collectType,
	"name": collectName,
}

func collectUserInput(state *branchState) {
	for _, key := range state.settings.prompts {
		if collector, ok := collectors[key]; ok {
			err := collector(state)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
		}
	}
}
