package cmdbranch

import (
	"log"

	"github.com/manifoldco/promptui"
)

func promptType(state *branchState) (string, error) {
	typesPresentation := []string{}
	for _, v := range state.settings.branchTypes {
		typePresentation := v
		if v, ok := state.settings.emojis[v]; ok {
			typePresentation = v + " " + typePresentation
		}
		typesPresentation = append(typesPresentation, typePresentation)
	}
	prompt := promptui.Select{
		Label: "Branch type",
		Items: typesPresentation,
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
		Label: "Branch name",
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
