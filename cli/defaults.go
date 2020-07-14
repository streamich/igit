package main

var prompts []string = []string{
	"type",
	"scope",
	"title",
	"body",
	"breaking-change",
	"skip-ci",
	"issues",
}

var commitTypes []string = []string{
	"test",
	"feat",
	"fix",
	"chore",
	"docs",
	"refactor",
	"style",
	"ci",
	"perf",
}

var emojis map[string]string = map[string]string{
	"test":     "💍",
	"feat":     "🎸",
	"fix":      "🐛",
	"chore":    "🤖",
	"docs":     "✏️",
	"refactor": "💡",
	"style":    "💄",
	"ci":       "🎡",
	"perf":     "⚡️",
	"release":  "🏹",
}

var defaultSettings CommitSettings = CommitSettings{
	prompts:     prompts,
	commitTypes: commitTypes,
	emojis:      emojis,
}
