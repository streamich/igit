package cmdbranch

import "time"

func formatBranchName(state *branchState) string {
	name := state.info.branchType

	if name != "" {
		if emoji, ok := state.settings.emojis[name]; ok {
			name = emoji + "-" + name
		}
	}

	if name != "" && state.settings.useTime {
		name = name + "/"
	}

	if state.settings.useTime {
		timeFormatted := time.Now().Format(time.RFC3339)[0:10]
		name = name + timeFormatted
	}

	if name != "" && state.info.name != "" {
		name = name + "/"
	}

	if state.info.name != "" {
		name = name + state.info.name
	}

	return name
}
