run:
	go run .

build:
	go build .

linux:
	go build -o build/lexilift .

windows:
	GOOS=windows GOARCH=amd64 go build -o build/lexilift.exe .
