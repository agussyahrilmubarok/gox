.PHONY: go/fmt
go/fmt:
	go fmt ./...

.PHONY: go/test
go/test:
	go test ./... -v

.PHONY: go/pkg/test
go/pkg/test:
	go test ./pkg/... -v

.PHONY: go/test/cover
go/test/cover:
	go test -timeout 9000s -cover -a -v ./...

.PHONY: go/build
go/build:
	go fmt ./...
	go vet ./...
	go test ./...
	go build -ldflags "-X github.com/agussyahrilmubarok/gox.Version=$(tag)" -o bin/gox .

.PHONY: go/push
go/push:
	git add .
	git commit -S -m "$(message)"
	git tag $(tag)
	git push origin main

.PHONY: clean
clean:
	rm -rf bin