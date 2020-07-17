module cmdcommit

go 1.14

require (
	github.com/AlecAivazis/survey/v2 v2.0.8
	github.com/manifoldco/promptui v0.7.0
	github.com/spf13/cobra v1.0.0
)

require (
	github.com/stretchr/testify v1.2.2
	gitutils v1.0.0
)

replace gitutils => ../gitutils
