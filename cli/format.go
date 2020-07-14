package main

func formatCommitMessage(state *CommitState) (string, error) {
	var message = state.info.commitType

	if state.info.scope != "" {
		message = message + "(" + state.info.scope + ")"
	}

	message = message + ":"

	if emoji, ok := state.settings.emojis[state.info.commitType]; ok {
		message = message + " " + emoji
	}

	if state.info.title != "" {
		message = message + " " + state.info.title
	}

	if state.info.body != "" {
		message = message + "\n\n" + state.info.body
	}

	if state.info.breakingChange != "" {
		message = message + "\n\n" + "BREAKING CHANGE: 💥 " + state.info.breakingChange
	}

	if state.info.skipCi {
		message = message + "\n\n" + "[skip ci]"
	}

	if state.info.closes != "" {
		message = message + "\n\n" + "✅ Closes " + state.info.closes
	}

	return message, nil
}
