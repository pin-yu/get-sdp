GO_ENV := GOARCH=arm64 GOOS=linux CGO_ENABLED=0

.PHONY: clean
clean:
	rm -rf bin

.PHONY: fmt
fmt:
	test -z "$$(goimports -w -l .)"

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: build
build: clean fmt vet staticcheck
	$(GO_ENV) go build -ldflags "-w -s" -trimpath -o bin/get-sdp main.go
