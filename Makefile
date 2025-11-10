go/test:
	go test ./... -v

go/pkg/test:
	go test ./pkg/... -v

go/test/cover:
	go test -timeout 9000s -cover -a -v ./...

go/fmt:
	go fmt ./...