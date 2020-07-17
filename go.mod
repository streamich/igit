module github.com/streamich/igit

go 1.14

require (
	cmdbranch v0.0.0
	cmdcommit v0.0.0
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/spf13/cobra v1.0.0
	gitutils v1.0.0
)

replace cmdcommit => ./internal/cmdcommit

replace cmdbranch => ./internal/cmdbranch

replace gitutils => ./internal/gitutils
