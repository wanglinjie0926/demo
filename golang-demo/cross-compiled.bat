set GOOS=linux
set GOARCH=arm64
 go build -o .\bin\client .\src\main.go .\src\strategy.go