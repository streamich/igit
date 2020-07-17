package igit

var prompts []string = []string{
	"type",
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
	"docs":     "📚",
	"refactor": "💡",
	"style":    "💄",
	"ci":       "🎡",
	"perf":     "⚡️",
	"release":  "🏹",
}

var descriptions = map[string]string{
	"test":     "Adding missing tests",
	"feat":     "A new feature",
	"fix":      "A bug fix",
	"chore":    "Build process or auxiliary tool changes",
	"docs":     "Documentation only changes",
	"refactor": "A code change that neither fixes a bug or adds a feature",
	"style":    "Markup, white-space, formatting, missing semi-colons...",
	"ci":       "Changes related to CI",
	"perf":     "A code change that improves performance",
	"release":  "Create a release commit",
}

var defaultSettings CommitSettings = CommitSettings{
	prompts:      prompts,
	commitTypes:  commitTypes,
	emojis:       emojis,
	descriptions: descriptions,
}
