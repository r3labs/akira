install:
	go install -v

build:
	go build -v ./...

lint:
	gometalinter --config .linter.conf

test:
	go test -v ./...

cover:
	go test -v ./... --cover

deps:
	go get github.com/nats-io/go-nats

dev-deps: deps
	go get github.com/alecthomas/gometalinter
	go get github.com/stretchr/testify/suite
	gometalinter --install

clean:
	go clean
