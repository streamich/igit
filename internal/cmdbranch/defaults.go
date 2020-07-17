package cmdbranch

var defaultSettings = branchSettings{
	useTime:     true,
	branchTypes: []string{"feature", "bug", "hotfix", "release"},
	prompts:     []string{"type", "name"},
}
