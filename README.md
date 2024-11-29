# go-norm
CLI to normalise unicode

## Go Build

```
GOOS=darwin GOARCH=arm64 go build -o norm-arm main.go
GOOS=darwin GOARCH=amd64 go build -o norm-amd main.go
lipo -create norm-arm norm-amd -output norm
GOOS=windows GOARCH=amd64 go build -o norm.exe main.go
```
