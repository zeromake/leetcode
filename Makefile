test:
	go test ./...

cover:
	./scripts/cover.sh
cover-html:
	./scripts/cover.sh html
