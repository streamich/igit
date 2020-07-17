module igit

go 1.14

require (
	github.com/AlecAivazis/survey/v2 v2.0.8
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/manifoldco/promptui v0.7.0
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/tcnksm/ghr v0.13.0 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

require gitutils v1.0.0

replace gitutils => ../../pkg/gitutils

require cmdcommit v1.0.0

replace cmdcommit => ../../pkg/cmdcommit
