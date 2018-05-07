build:
	go build

install: build
	go install

test:
	go test ./...

testacc:
	TF_ACC=1 go test ./... -v
