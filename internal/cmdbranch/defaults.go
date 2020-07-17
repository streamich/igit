package cmdbranch

var emojis = map[string]string{
	"feat":    "🎸",
	"fix":     "🐛",
	"hotfix":  "🔥",
	"release": "🏹",
}

var defaultSettings = branchSettings{
	useTime:     true,
	branchTypes: []string{"feat", "fix", "hotfix", "release"},
	prompts:     []string{"type", "name"},
	emojis:      emojis,
}
