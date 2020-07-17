module github.com/streamich/igit

go 1.14

require gitutils v1.0.0

replace gitutils => ./internal/gitutils

require (
	cmdcommit v1.0.0
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/spf13/cobra v1.0.0
)

replace cmdcommit => ./internal/cmdcommit
