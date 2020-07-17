package cmdbranch

type branchSettings struct {
	useTime     bool
	branchTypes []string
	prompts     []string
	emojis      map[string]string
}

// User provided information.
type branchInfo struct {
	branchType string
	name       string
}

type branchState struct {
	settings branchSettings
	info     branchInfo
}
