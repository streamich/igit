install:
	go install
	go mod download
	go get -u github.com/mitchellh/gox
	go get -u github.com/tcnksm/ghr

run:
	go run . $(ARGS)

test:
	go test cmdcommit

build:
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -output="build/{{.OS}}_{{.Arch}}"

clean:
	rm -rf ./{build,dist}/
