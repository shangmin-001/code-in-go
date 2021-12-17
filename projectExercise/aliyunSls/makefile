TARGET := aliyunSLS


.PHONY: linux64 darwin clean build

build:
	go build -o $(TARGET) .

clean:
	rm -rf *.log

linux64:
	GOOS=linux GOARCH=amd64 go build -o $(TARGET) .

darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(TARGET) .

test:
	go generate; CGO_ENABLED=1 go test -race -cover `go list ./... | grep -v "tmp" `

gofmt:
	find . -name '*.go' | xargs -L1 go fmt
