module github.com/streamich/igit

go 1.14

require internal/app/igit v1.0.0

replace internal/app/igit => ../../internal/app/igit

require gitutils v1.0.0

replace gitutils => ../../internal/pkg/gitutils

require cmdcommit v1.0.0

replace cmdcommit => ../../internal/pkg/cmdcommit
