.PHONY: run build linux windows

run:
	go run .

build: linux windows

linux:
	@bash -c '\
	branch=$$(git rev-parse --abbrev-ref HEAD); \
	tag=$$(git describe --tags --abbrev=0 2>/dev/null); \
	commit=$$(git rev-parse --short HEAD); \
	env GOOS=linux GOARCH=amd64 go build -o build/lexilift-$$branch-$$tag-$$commit;'

windows:
	@bash -c '\
	branch=$$(git rev-parse --abbrev-ref HEAD); \
	tag=$$(git describe --tags --abbrev=0 2>/dev/null); \
	commit=$$(git rev-parse --short HEAD); \
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o build/lexilift-$$branch-$$tag-$$commit.exe;'

