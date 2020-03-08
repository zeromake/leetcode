test:
	go test ./...

cover:
	./scripts/cover.sh count
cover-html:
	./scripts/cover.sh html
