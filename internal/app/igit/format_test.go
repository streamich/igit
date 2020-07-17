package igit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createState(info CommitInfo) *CommitState {
	state := CommitState{
		info:     info,
		settings: defaultSettings,
	}

	return &state
}

func TestSimpleStateWithTypeAndTitleIsFormatteedCorrectly(t *testing.T) {
	state := createState(CommitInfo{
		commitType: "feat",
		title:      "foo",
	})
	message, err := formatCommitMessage(state)

	assert.Nil(t, err)
	assert.Equal(t, message, "feat: 🎸 foo")
}

func TestFormatsMessageFullOfUserInput(t *testing.T) {
	state := createState(CommitInfo{
		commitType:     "fix",
		title:          "worked 2h on it, phew",
		body:           "This was some very complicated commit that solved x, y, and z.",
		breakingChange: "Breaks frontend API, completely!",
		closes:         "#123, #4, and #5",
		scope:          "backend",
		skipCi:         true,
	})
	message, err := formatCommitMessage(state)

	assert.Nil(t, err)
	assert.Equal(t, message,
		`fix(backend): 🐛 worked 2h on it, phew

This was some very complicated commit that solved x, y, and z.

BREAKING CHANGE: 💥 Breaks frontend API, completely!

[skip ci]

✅ Closes #123, #4, and #5`,
	)
}
