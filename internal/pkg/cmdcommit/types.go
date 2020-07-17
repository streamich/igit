package cmdcommit

// CommitSettings are settings provided by user or through a settings file
// before collecting CommitInfo interactive from the user.
type CommitSettings struct {
	// Prompt represents type of questions (aka prompts) that use will be
	// interactively asked.
	prompts []string

	// Whether to allow user to enter semantic commit scope in free text form.
	allowScopes bool

	// Semantic commit types, such as "feat", "test", etc..
	commitTypes []string

	// Map of commit types to emojis.
	emojis map[string]string

	// Map of commit type descriptions.
	descriptions map[string]string
}

// CommitInfo is information needed to format a commit message.
//
// Formatted semantic commit structure:
//
//     type(scope): title
//
//     body
//
//     [skip-ci]
//
//     BREAKING CHANGE: description
//
//     Closes #123
//
type CommitInfo struct {
	// Selected semantic commit type, such as "feat", "test", etc..
	commitType string

	// Semantic commit scope.
	scope string

	// User-entered semantic commit title.
	title string

	// Long-form description of commit.
	body string

	// Summary of a breaking change. If present, semantic commit will contain a
	// section with BREAKING CHANGE: ... block, which will result semantic semver
	// to be bumped to the next major.
	breakingChange string

	// If present, "[skip ci]" will be added to commit message which
	// informs most continuous integration tools to skip this commit.
	skipCi bool

	// List of GitHub issues to close.
	closes string
}

type commitFlags struct {
	args  []string
	scope bool
}

// CommitState is coplete global state of the CLI tool.
type CommitState struct {
	flags    *commitFlags
	info     CommitInfo
	settings CommitSettings
}
