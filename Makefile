run:
	go run .

build:
	go build .

build-linux:
	go build -o build/lexilift .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/lexilift.exe .
