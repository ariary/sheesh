before.build:
	go mod tidy && go mod download

build.sheesh:
	@echo "build in ${PWD}";go build sheesh.go