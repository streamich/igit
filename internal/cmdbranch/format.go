package cmdbranch

func formatBranchName(state *branchState) string {
	name := state.info.branchType

	if name != "" {
		name = name + "/"
	}

	if state.info.name != "" {
		name = name + state.info.name
	}

	return name
}
