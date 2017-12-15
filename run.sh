clear

# Run during Development
go run $(find . -name "*.go" -and -not -name "*_test.go" -maxdepth 1)

# Run via building binary
# go build && ./go-boilerplate
