download:
	go mod download

run:
	go run . $(ARGS)

test:
	go test

build:
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -output="build/{{.OS}}_{{.Arch}}"
